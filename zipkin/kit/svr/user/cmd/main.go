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
	"github.com/openzipkin/zipkin-go"
	openzipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
)

func main() {

	logger := svr.NewLogger("user")

	//client start
	var zipkinTracerMsg *zipkin.Tracer
	{
		var (
			err           error
			hostPort      = svr.UserSvrAddress
			serviceName   = svr.UserSvrName
			useNoopTracer = false
			reporter      = openzipkinhttp.NewReporter(svr.Zipkinhttpurl)
		)
		defer reporter.Close()
		zEP, _ := zipkin.NewEndpoint(serviceName, hostPort)
		zipkinTracerMsg, err = zipkin.NewTracer(
			reporter, zipkin.WithLocalEndpoint(zEP), zipkin.WithNoopTracer(useNoopTracer),
		)
		if err != nil {
			logger.Log("err", err)
			os.Exit(1)
		}
		if !useNoopTracer {
			logger.Log("tracer", "Zipkin", "type", "Native", "URL", svr.Zipkinhttpurl)
		}
	}

	var msgClient pb.MsgServer

	conn, e := grpc.Dial(svr.MsgSvrAddress, grpc.WithInsecure())
	defer conn.Close()
	if e != nil {
		logger.Log("err", e)
	}

	msgClient = biz.NewMsgClient(conn, zipkinTracerMsg, logger)

	//client end

	var zipkinTracerUser *zipkin.Tracer
	{
		var (
			err           error
			hostPort      = svr.UserSvrAddress
			serviceName   = svr.UserSvrName
			useNoopTracer = false
			reporter      = openzipkinhttp.NewReporter(svr.Zipkinhttpurl)
		)
		defer reporter.Close()
		zEP, _ := zipkin.NewEndpoint(serviceName, hostPort)
		zipkinTracerUser, err = zipkin.NewTracer(
			reporter, zipkin.WithLocalEndpoint(zEP), zipkin.WithNoopTracer(useNoopTracer),
		)
		if err != nil {
			logger.Log("err", err)
			os.Exit(1)
		}
		if !useNoopTracer {
			logger.Log("tracer", "Zipkin", "type", "Native", "URL", svr.Zipkinhttpurl)
		}
	}

	service := biz.UserSvr{Logger: logger, MsgClient: msgClient}

	endpoint := user.MakeLoginEndpoint(service)

	opts := svr.NewGrpcServerOptions(zipkinTracerUser, "", logger)

	server := kitgrpc.NewServer(endpoint, svr.NoDecodeRequestFunc, svr.NoEncodeResponseFunc, opts...)

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
