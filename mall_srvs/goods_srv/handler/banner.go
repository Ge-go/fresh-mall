package handler

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"mall_srvs/goods_srv/global"
	"mall_srvs/goods_srv/model"
	"mall_srvs/goods_srv/proto"
)

func (g *GoodsServer) BannerList(ctx context.Context, req *emptypb.Empty) (*proto.BannerListResponse, error) {
	var data []model.Banner
	result := global.DB.WithContext(ctx).Find(&data, "is_deleted", 0)
	if result.RowsAffected == 0 {
		return &proto.BannerListResponse{}, nil
	} else if result.Error != nil {
		zap.S().Errorw("get banner list err", "msg", result.Error.Error())
		return nil, status.Error(codes.Internal, result.Error.Error())
	}

	var res []*proto.BannerResponse
	for _, v := range data {
		res = append(res, &proto.BannerResponse{
			Id:    v.ID,
			Index: v.Index,
			Image: v.Image,
			Url:   v.Url,
		})
	}

	return &proto.BannerListResponse{
		Total: int32(result.RowsAffected),
		Data:  res,
	}, nil
}

func (g *GoodsServer) CreateBanner(ctx context.Context, req *proto.BannerRequest) (*proto.BannerResponse, error) {
	banner := &model.Banner{
		Url:   req.Url,
		Image: req.Image,
		Index: req.Index,
	}
	if result := global.DB.WithContext(ctx).Create(banner); result.Error != nil {
		zap.S().Errorw("create banner err", "msg", result.Error.Error())
		return nil, status.Error(codes.Internal, result.Error.Error())
	}

	return &proto.BannerResponse{
		Id:    banner.ID,
		Image: banner.Image,
		Index: banner.Index,
		Url:   banner.Url,
	}, nil
}

func (g *GoodsServer) DeleteBanner(ctx context.Context, req *proto.BannerRequest) (*emptypb.Empty, error) {
	if result := global.DB.WithContext(ctx).Where("id", req.Id).Delete(&model.Banner{}, "is_deleted", 0); result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "not found banner for delete")
	} else if result.Error != nil {
		zap.S().Errorw("del banner err", "msg", result.Error.Error())
		return nil, status.Error(codes.Internal, result.Error.Error())
	}

	return &emptypb.Empty{}, nil
}

func (g *GoodsServer) UpdateBanner(ctx context.Context, req *proto.BannerRequest) (*emptypb.Empty, error) {
	var banner model.Banner
	if result := global.DB.WithContext(ctx).Where("id", req.Id).Find(&banner, "is_deleted", 0); result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "not found banner for update")
	} else if result.Error != nil {
		zap.S().Errorw("find banner err", "msg", result.Error.Error())
		return nil, status.Error(codes.Internal, result.Error.Error())
	}
	if req.Url != "" {
		banner.Url = req.Url
	}
	if req.Index != 0 {
		banner.Index = req.Index
	}
	if req.Image != "" {
		banner.Image = req.Image
	}
	if res := global.DB.WithContext(ctx).Save(&banner); res.Error != nil {
		zap.S().Errorw("update banner err", "msg", res.Error.Error())
		return nil, status.Error(codes.Internal, res.Error.Error())
	}

	return &emptypb.Empty{}, nil
}
