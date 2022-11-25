package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"mall_srvs/order_srv/proto"
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

	// brand
	//GetBrand()
	//createBrand()
	//deleteBrand()
	//saveBrand()

	//banner
	//getBanner()
	//createBanner()
	//delBanner()
	//upBanner()

	//category
	//getCategorys()
	//getSubCategorys()
	//createCategory()

	//category brand
	categoryBrandList()
	//getCategoryBrandList()
}

func getCategoryBrandList() {
	list, err := goodsClient.GetCategoryBrandList(context.Background(), &proto.CategoryInfoRequest{
		Id: 135475,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(list)
}

func categoryBrandList() {
	list, err := goodsClient.CategoryBrandList(context.Background(), &proto.CategoryBrandFilterRequest{
		Pages:       1,
		PagePerNums: 10,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(list)
}

func createCategory() {
	category, err := goodsClient.CreateCategory(context.Background(), &proto.CategoryInfoRequest{
		Name:           "特殊水果",
		ParentCategory: 123456789,
		Level:          2,
		IsTab:          true,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(category)
}

func getSubCategorys() {
	category, err := goodsClient.GetSubCategory(context.Background(), &proto.CategoryListRequest{
		Id: 135487,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(category)
}

func getCategorys() {
	list, err := goodsClient.GetAllCategorysList(context.Background(), &emptypb.Empty{})
	if err != nil {
		panic(err)
	}
	fmt.Println(list.JsonData)
}

func upBanner() {
	banner, err := goodsClient.UpdateBanner(context.Background(), &proto.BannerRequest{
		Id:    1,
		Index: 3,
		Image: "123",
		Url:   "456",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(banner)
}

func delBanner() {
	banner, err := goodsClient.DeleteBanner(context.Background(), &proto.BannerRequest{
		Id: 2,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(banner)
}

func createBanner() {
	banner, err := goodsClient.CreateBanner(context.Background(), &proto.BannerRequest{
		Url:   "www:test",
		Index: 1,
		Image: "123",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(banner)
}

func saveBrand() {
	brand, err := goodsClient.UpdateBrand(context.Background(), &proto.BrandRequest{
		Id:   112233,
		Name: "就这???",
		Logo: "test",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(brand)
}

func deleteBrand() {
	brand, err := goodsClient.DeleteBrand(context.Background(), &proto.BrandRequest{
		Id: 1111,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(brand)
}

func createBrand() {
	brand, err := goodsClient.CreateBrand(context.Background(), &proto.BrandRequest{
		Name: "谁tm买小米",
		Logo: "test",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(brand)
}

func GetBrand() {
	brandLis, err := goodsClient.BrandList(context.Background(), &proto.BrandFilterRequest{
		Pages:       0,
		PagePerNums: 0,
	})
	if err != nil {
		panic(err)
	}
	for _, v := range brandLis.Data {
		fmt.Println("获取到的brand", v)
	}
	fmt.Println("brand 总条数", brandLis.Total)
}

// bananer
func getBanner() {
	list, err := goodsClient.BannerList(context.Background(), &emptypb.Empty{})
	if err != nil {
		panic(err)
	}
	fmt.Println(list)
}
