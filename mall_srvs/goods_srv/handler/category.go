package handler

import (
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"mall_srvs/goods_srv/global"
	"mall_srvs/goods_srv/model"
	"mall_srvs/goods_srv/proto"
)

func (g *GoodsServer) GetAllCategorysList(ctx context.Context, req *emptypb.Empty) (*proto.CategoryListResponse, error) {
	var categorys []model.Category
	if res := global.DB.Where(&model.Category{Level: 1}).
		Preload("SubCategory.SubCategory", "is_deleted", 0).Find(&categorys, "is_deleted", 0); res.Error != nil {
		zap.S().Errorw("get categorys err", "msg", res.Error.Error())
		return nil, status.Error(codes.Internal, res.Error.Error())
	}

	data, err := json.Marshal(&categorys)
	if err != nil {
		zap.S().Errorw("json marshal categorys err", "msg", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.CategoryListResponse{
		JsonData: string(data), Data: nil,
	}, nil
}

func (g *GoodsServer) GetSubCategory(ctx context.Context, req *proto.CategoryListRequest) (*proto.SubCategoryListResponse, error) {
	resp := proto.SubCategoryListResponse{}
	var category model.Category
	if res := global.DB.WithContext(ctx).First(&category, req.Id); res.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "cannot found category")
	} else if res.Error != nil {
		zap.S().Errorw("get category by id err", "msg", res.Error.Error())
		return nil, status.Error(codes.Internal, res.Error.Error())
	}

	resp.Info = &proto.CategoryInfoResponse{
		Id:             category.ID,
		Name:           category.Name,
		Level:          category.Level,
		IsTab:          category.IsTab,
		ParentCategory: category.ParentCategoryID,
	}

	var subCategorys []model.Category
	//preloads := "SubCategory"
	//if category.Level == 1 {
	//	preloads = "SubCategory.SubCategory"
	//}

	if res := global.DB.Where(&model.Category{ParentCategoryID: req.Id}).Find(&subCategorys, "is_deleted", 0); res.Error != nil {
		zap.S().Errorw("get sub_categorys err", "msg", res.Error.Error())
		return nil, status.Error(codes.Internal, res.Error.Error())
	}

	for _, v := range subCategorys {
		resp.SubCategorys = append(resp.SubCategorys, &proto.CategoryInfoResponse{
			Id:             v.ID,
			Name:           v.Name,
			ParentCategory: v.ParentCategoryID,
			Level:          v.Level,
			IsTab:          v.IsTab,
		})
	}

	return &resp, nil
}

func (g *GoodsServer) CreateCategory(ctx context.Context, req *proto.CategoryInfoRequest) (*proto.CategoryInfoResponse, error) {
	panic("implement me")
}

func (g *GoodsServer) DeleteCategory(ctx context.Context, req *proto.DeleteCategoryRequest) (*emptypb.Empty, error) {
	panic("implement me")
}

func (g *GoodsServer) UpdateCategory(ctx context.Context, req *proto.CategoryInfoRequest) (*emptypb.Empty, error) {
	panic("implement me")
}
