package main

import (
	"fmt"
	"gitee.com/godY/gokit-inaction/aio/pro"
	"gitee.com/godY/gokit-inaction/aio/svr/account"
	"gitee.com/godY/gokit-inaction/aio/svr/global"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
	"net"
	"os"
)

var Address = "localhost:19922"
var servicename = "accountsvr"

var Zipkinhttpurl = "http://192.168.3.125:9411/api/v2/spans"

func main() {

	logger := global.NewLogger(servicename)

	logger.Log("start")

	api := account.AccountSvr{}

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

	getAccountServer := kitgrpc.NewServer(account.MakeGetAccountEndpoint(api), global.NoDecodeRequestFunc, global.NoEncodeResponseFunc)

	handler := account.Handler{GetAccountHandler: getAccountServer}

	grpcListener, err := net.Listen("tcp", Address)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))

	pro.RegisterAccountServer(grpcServer, handler)

	err = grpcServer.Serve(grpcListener)

	if err != nil {
		fmt.Println(err)
	}

}
