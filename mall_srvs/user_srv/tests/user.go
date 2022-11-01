package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/credentials/insecure"
	"time"

	"google.golang.org/grpc"

	"mall_srvs/user_srv/proto"
)

var conn *grpc.ClientConn
var userClient proto.UserClient

func Init() {
	var err error
	conn, err = grpc.Dial("0.0.0.0:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	userClient = proto.NewUserClient(conn)
}

// 测试 user 内所有接口
func main() {
	Init()
	defer conn.Close()

	//TestUserList()
	TestUserById()
	//TestCreateUser()
}

func TestUserList() {
	fmt.Println("获取用户列,校验密码...")
	// 获取用户列
	rsp, err := userClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    1,
		PSize: 3,
	})
	if err != nil {
		panic(err)
	}
	for _, data := range rsp.Data {
		fmt.Println(data)
		//验证密码
		resCheck, err := userClient.CheckPassWord(context.Background(), &proto.PasswordCheckInfo{
			PassWord:          "admin123",
			EncryptedPassword: data.PassWord,
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(resCheck.Success)
	}
}

func TestUserById() {
	fmt.Println("获取用户通过id,通过手机号获取用户")
	resp, err := userClient.GetUserById(context.Background(), &proto.IdRequest{
		Id: "10",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

	respMobile, err := userClient.GetUserByMobile(context.Background(), &proto.MobileRequest{
		Mobile: resp.Mobile,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("the mobile is ", resp.Mobile)
	fmt.Println(respMobile)
}

func TestCreateUser() {
	fmt.Println("创建用户,修改用户...")
	resp, err := userClient.CreateUser(context.Background(), &proto.CreateUserInfo{
		NickName: "wsAA",
		PassWord: "admin123123",
		Mobile:   "15221889900",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
	user, err := userClient.UpdateUser(context.Background(), &proto.UpdateUserInfo{
		Id:       resp.Id,
		NickName: "wsBB",
		Gender:   "female",
		BirthDay: uint64(time.Now().Unix()),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
}
