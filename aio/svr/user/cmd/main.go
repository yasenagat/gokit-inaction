package main

import (
	"fmt"
	"gitee.com/godY/gokit-inaction/aio/pro"
	"gitee.com/godY/gokit-inaction/aio/svr/global"
	"gitee.com/godY/gokit-inaction/aio/svr/user"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
	"net"
	"os"
)

var Address = "localhost:19911"
var servicename = "usersvr"

var Zipkinhttpurl = "http://192.168.3.125:9411/api/v2/spans"

func main() {

	logger := global.NewLogger(servicename)

	logger.Log("start")

	api := user.UserSvr{}

	//var zipkinTracer *zipkin.Tracer
	//{
	//	var (
	//		err           error
	//		hostPort      = Address
	//		serviceName   = servicename
	//		useNoopTracer = false
	//		reporter      = openzipkinhttp.NewReporter("")
	//	)
	//	defer reporter.Close()
	//	zEP, _ := zipkin.NewEndpoint(serviceName, hostPort)
	//	zipkinTracer, err = zipkin.NewTracer(
	//		reporter, zipkin.WithLocalEndpoint(zEP), zipkin.WithNoopTracer(useNoopTracer),
	//	)
	//	if err != nil {
	//		logger.Log("err", err)
	//		os.Exit(1)
	//	}
	//	if !useNoopTracer {
	//		logger.Log("tracer", "Zipkin", "type", "Native", "URL", Zipkinhttpurl)
	//	}
	//}
	//
	//opts := global.NewGrpcServerOptions(zipkinTracer, "", logger)

	loginServer := kitgrpc.NewServer(user.MakeLoginEndpoint(api), global.NoDecodeRequestFunc, global.NoEncodeResponseFunc)

	userHandler := user.Handler{LoginHandler: loginServer}

	grpcListener, err := net.Listen("tcp", Address)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))

	pro.RegisterUserServer(grpcServer, userHandler)

	err = grpcServer.Serve(grpcListener)

	if err != nil {
		fmt.Println(err)
	}

}
