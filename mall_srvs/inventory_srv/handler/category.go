package handler

import (
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
	"mall_srvs/inventory_srv/global"
	"mall_srvs/inventory_srv/model"
	"mall_srvs/inventory_srv/proto"
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
	category := &model.Category{
		Name:  req.Name,
		Level: req.Level,
		IsTab: req.IsTab,
		Url:   "test.com", // 仅用作测试url
	}

	if req.Level > 1 {
		// 做一个父节点是否存在的动作
		if res := global.DB.WithContext(ctx).Find(&model.Category{}, req.ParentCategory); res.RowsAffected == 0 {
			return nil, status.Errorf(codes.NotFound, "cannot found parent category %d", req.ParentCategory)
		} else if res.Error != nil {
			zap.S().Errorw("find parent category for id", "msg", res.Error.Error(), "id", req.ParentCategory)
			return nil, status.Error(codes.Internal, res.Error.Error())
		}
		category.ParentCategoryID = req.ParentCategory
	}
	if res := global.DB.WithContext(ctx).Save(category); res.Error != nil {
		zap.S().Errorw("create category err", "msg", res.Error.Error())
		return nil, status.Error(codes.Internal, res.Error.Error())
	}

	return &proto.CategoryInfoResponse{
		Id:             category.ID,
		Name:           category.Name,
		ParentCategory: category.ParentCategoryID,
		Level:          category.Level,
		IsTab:          category.IsTab,
	}, nil
}

func (g *GoodsServer) DeleteCategory(ctx context.Context, req *proto.DeleteCategoryRequest) (*emptypb.Empty, error) {
	if res := global.DB.Where("id", req.Id).Update("is_deleted", 1); res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil, status.Error(codes.NotFound, "not found category for id")
		}
		zap.S().Errorw("cannot found category by id", "msg", res.Error.Error())
		return nil, status.Error(codes.Internal, res.Error.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *GoodsServer) UpdateCategory(ctx context.Context, req *proto.CategoryInfoRequest) (*emptypb.Empty, error) {
	var category model.Category

	if result := global.DB.First(&category, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品分类不存在")
	}

	if req.Name != "" {
		category.Name = req.Name
	}
	if req.ParentCategory != 0 {
		category.ParentCategoryID = req.ParentCategory
	}
	if req.Level != 0 {
		category.Level = req.Level
	}
	if req.IsTab {
		category.IsTab = req.IsTab
	}

	global.DB.Save(&category)

	return &emptypb.Empty{}, nil
}
