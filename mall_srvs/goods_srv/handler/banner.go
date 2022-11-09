package handler

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"mall_srvs/goods_srv/proto"
)

func (g *GoodsServer) BannerList(ctx context.Context, req *emptypb.Empty) (*proto.BannerListResponse, error) {
	panic("implement me")
}

func (g *GoodsServer) CreateBanner(ctx context.Context, req *proto.BannerRequest) (*proto.BannerResponse, error) {
	panic("implement me")
}

func (g *GoodsServer) DeleteBanner(ctx context.Context, req *proto.BannerRequest) (*emptypb.Empty, error) {
	panic("implement me")
}

func (g *GoodsServer) UpdateBanner(ctx context.Context, req *proto.BannerRequest) (*emptypb.Empty, error) {
	panic("implement me")
}
