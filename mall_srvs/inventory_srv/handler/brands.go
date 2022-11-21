package handler

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
	"mall_srvs/inventory_srv/global"
	"mall_srvs/inventory_srv/model"
	"mall_srvs/inventory_srv/proto"
)

// BrandList 品牌和轮播图
func (g *GoodsServer) BrandList(ctx context.Context, req *proto.BrandFilterRequest) (*proto.BrandListResponse, error) {
	var resp proto.BrandListResponse

	var brands []model.Brands
	res := global.DB.WithContext(ctx).Where("is_deleted = ?", 0).
		Scopes(model.Paginate(int(req.Pages), int(req.PagePerNums))).Find(&brands)
	if res.Error == gorm.ErrRecordNotFound {
		return nil, status.Error(codes.NotFound, "cannot find brand")
	}
	if res.Error != nil {
		zap.S().Errorw("get brands error", "msg", res.Error.Error())
		return nil, status.Error(codes.Internal, res.Error.Error())
	}

	var total int64
	global.DB.WithContext(ctx).Where("is_deleted = ?", 0).Model(&model.Brands{}).Count(&total)

	var data []*proto.BrandInfoResponse
	for _, v := range brands {
		data = append(data, &proto.BrandInfoResponse{
			Id:   v.ID,
			Name: v.Name,
			Logo: v.Logo,
		})
	}
	resp.Data = data
	resp.Total = int32(total)

	return &resp, nil
}

func (g *GoodsServer) CreateBrand(ctx context.Context, req *proto.BrandRequest) (*proto.BrandInfoResponse, error) {
	if result := global.DB.WithContext(ctx).Where(&model.Brands{BaseModel: model.BaseModel{IsDeleted: 0}, Name: req.Name}).First(&model.Brands{}); result.RowsAffected == 1 {
		return nil, status.Error(codes.InvalidArgument, "the brand already exists")
	} else if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		zap.S().Errorw("find brands by name error", "msg", result.Error.Error())
		return nil, status.Error(codes.Internal, result.Error.Error())
	}
	brand := &model.Brands{
		Name: req.Name,
		Logo: req.Logo,
	}
	result := global.DB.WithContext(ctx).Save(brand)
	if result.Error != nil {
		zap.S().Errorw("create brand error", "msg", result.Error.Error())
		return nil, status.Error(codes.Internal, result.Error.Error())
	}

	return &proto.BrandInfoResponse{
		Id: brand.ID,
	}, nil
}

func (g *GoodsServer) DeleteBrand(ctx context.Context, req *proto.BrandRequest) (*emptypb.Empty, error) {
	delSql := fmt.Sprint("update brands set is_deleted = ? where id = ? and is_deleted = ?")
	if result := global.DB.WithContext(ctx).Exec(delSql, 1, req.Id, 0); result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "the brand not exists")
	} else if result.Error != nil {
		zap.S().Errorw("delete brand err", "msg", result.Error.Error())
		return nil, status.Error(codes.Internal, result.Error.Error())
	}
	return &emptypb.Empty{}, nil
}

func (g *GoodsServer) UpdateBrand(ctx context.Context, req *proto.BrandRequest) (*emptypb.Empty, error) {
	brands := model.Brands{BaseModel: model.BaseModel{ID: req.Id, IsDeleted: 0}}
	if result := global.DB.WithContext(ctx).First(&brands); result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "cannot find brand for update")
	} else if result.Error != nil {
		zap.S().Errorw("update brand err", "msg", result.Error.Error())
		return nil, status.Error(codes.Internal, result.Error.Error())
	}
	if req.Name != "" {
		brands.Name = req.Name
	}
	if req.Logo != "" {
		brands.Logo = req.Logo
	}

	if result := global.DB.WithContext(ctx).Save(&brands); result.Error != nil {
		zap.S().Errorw("find brand err", "msg", result.Error.Error())
		return nil, status.Error(codes.Internal, result.Error.Error())
	}

	return &emptypb.Empty{}, nil
}
