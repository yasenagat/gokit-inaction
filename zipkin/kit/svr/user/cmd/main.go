package main

import (
	"net"
	"fmt"
	"os"
	"google.golang.org/grpc"
	kitzipkin "github.com/go-kit/kit/tracing/zipkin"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"gitee.com/godY/gokit-inaction/zipkin/kit/svr/user/biz"
	"gitee.com/godY/gokit-inaction/zipkin/kit/svr/user"
	"gitee.com/godY/gokit-inaction/zipkin/kit/svr"
	"gitee.com/godY/gokit-inaction/zipkin/kit/svr/pro"
)

func main() {

	logger := svr.NewLogger()

	service := biz.UserSvr{}

	endpoint := user.MakeLoginEndpoint(service)

	zipkinTracer := svr.NewZipkinTracer(svr.UserSvrName, svr.UserSvrAddress, svr.Zipkinhttpurl, logger)

	zipkinServer := kitzipkin.GRPCServerTrace(zipkinTracer)

	options := []kitgrpc.ServerOption{
		kitgrpc.ServerErrorLogger(logger),
		zipkinServer,
	}

	server := kitgrpc.NewServer(endpoint, svr.NoDecodeRequestFunc, svr.NoEncodeResponseFunc, options...)

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
