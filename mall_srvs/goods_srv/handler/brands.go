package handler

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
	"mall_srvs/goods_srv/global"
	"mall_srvs/goods_srv/model"
	"mall_srvs/goods_srv/proto"
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
		return nil, res.Error
	}

	var total int64
	global.DB.Model(&model.Brands{}).Count(&total)

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
	panic("implement me")
}

func (g *GoodsServer) DeleteBrand(ctx context.Context, req *proto.BrandRequest) (*emptypb.Empty, error) {
	panic("implement me")
}

func (g *GoodsServer) UpdateBrand(ctx context.Context, req *proto.BrandRequest) (*emptypb.Empty, error) {
	panic("implement me")
}
