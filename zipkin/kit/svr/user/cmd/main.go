package main

import (
	"net"
	"fmt"
	"os"
	"google.golang.org/grpc"
	_ "net/http/pprof"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"gitee.com/godY/gokit-inaction/zipkin/kit/svr/user/biz"
	"gitee.com/godY/gokit-inaction/zipkin/kit/svr/user"
	"gitee.com/godY/gokit-inaction/zipkin/kit/svr"
	"gitee.com/godY/gokit-inaction/zipkin/kit/svr/pro"
)

func main() {

	logger := svr.NewLogger("user")

	service := biz.UserSvr{Logger: logger}

	endpoint := user.MakeLoginEndpoint(service)

	//svr.NewServerOptions(svr.UserSvrName, svr.UserSvrAddress, svr.Zipkinhttpurl, logger)

	server := kitgrpc.NewServer(endpoint, svr.NoDecodeRequestFunc, svr.NoEncodeResponseFunc)

	handler := user.Handler{LoginEndpoint: server}

	grpcListener, err := net.Listen("tcp", svr.UserSvrAddress)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))

	pb.RegisterUserServer(grpcServer, handler)

	err = grpcServer.Serve(grpcListener)

	if err != nil {
		fmt.Println(err)
	}
}
