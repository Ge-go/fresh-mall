package main

import (
	"flag"
	"fmt"
	"net"

	"google.golang.org/grpc"

	"mall_srvs/user_srv/handler"
	"mall_srvs/user_srv/proto"
)

var (
	IP   = flag.String("ip", "0.0.0.0", "ip地址")
	Port = flag.String("port", ":50051", "端口号")
)

func main() {
	flag.Parse()

	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})
	lis, err := net.Listen("tcp", *Port)
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	fmt.Printf("user_srv is running,the ip is %v port is %v\n", *IP, *Port)
	err = server.Serve(lis)
	if err != nil {
		panic("failed to start user_srv." + err.Error())
	}
}
