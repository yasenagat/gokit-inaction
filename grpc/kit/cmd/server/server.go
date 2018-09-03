package main

import (
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"gitee.com/godY/gokit-inaction/grpc/biz"
	"net"
	"os"

	"google.golang.org/grpc"
	"gitee.com/godY/gokit-inaction/grpc/pb"
	"fmt"
	"gitee.com/godY/gokit-inaction/grpc/kit"
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
