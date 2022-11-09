package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"mall_srvs/goods_srv/proto"
)

var conn *grpc.ClientConn
var goodsClient proto.GoodsClient

func Init() {
	var err error
	conn, err = grpc.Dial("192.168.48.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	goodsClient = proto.NewGoodsClient(conn)
}

// 测试 user 内所有接口
func main() {
	Init()
	defer conn.Close()

	GetBrand()
}

func GetBrand() {
	brandLis, err := goodsClient.BrandList(context.Background(), &proto.BrandFilterRequest{
		Pages:       1,
		PagePerNums: 20,
	})
	if err != nil {
		panic(err)
	}
	for _, v := range brandLis.Data {
		fmt.Println("获取到的brand", v)
	}
	fmt.Println("brand 总条数", brandLis.Total)
}
