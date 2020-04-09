package main

import (
	"gitee.com/godY/gokit-inaction/grpc/biz"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"net"
	"os"

	"fmt"
	"gitee.com/godY/gokit-inaction/grpc/kit"
	"gitee.com/godY/gokit-inaction/grpc/pb"
	"google.golang.org/grpc"
)

func main() {

	us := biz.UserServer{}

	loginEndpoint := kit.MakeLoginEndpoint(us)

	loginServer := kitgrpc.NewServer(loginEndpoint, kit.NoDecodeRequestFunc, kit.NoEncodeResponseFunc)

	userHandler := kit.UserHandler{LoginHandler: loginServer}

	grpcListener, err := net.Listen("tcp", ":8082")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))

	pb.RegisterUserServer(grpcServer, userHandler)

	err = grpcServer.Serve(grpcListener)

	if err != nil {
		fmt.Println(err)
	}
}
