package main

import (
	"fmt"
	"gitee.com/godY/gokit-inaction/zipkin/kit/svr"
	"gitee.com/godY/gokit-inaction/zipkin/kit/svr/msg"
	"gitee.com/godY/gokit-inaction/zipkin/kit/svr/msg/biz"
	"gitee.com/godY/gokit-inaction/zipkin/kit/svr/pro"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/openzipkin/zipkin-go"
	openzipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
	"google.golang.org/grpc"
	"net"
	"os"
)

func main() {

	logger := svr.NewLogger("msg")

	service := biz.MsgSvr{Logger: logger}

	endpoint := msg.MakeUnReadEndpoint(service)

	var zipkinTracer *zipkin.Tracer
	{
		var (
			err           error
			hostPort      = svr.MsgSvrAddress
			serviceName   = svr.MsgSvrName
			useNoopTracer = false
			reporter      = openzipkinhttp.NewReporter(svr.Zipkinhttpurl)
		)
		defer reporter.Close()
		zEP, _ := zipkin.NewEndpoint(serviceName, hostPort)
		zipkinTracer, err = zipkin.NewTracer(
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

	opts := svr.NewGrpcServerOptions(zipkinTracer, "", logger)

	server := kitgrpc.NewServer(endpoint, svr.NoDecodeRequestFunc, svr.NoEncodeResponseFunc, opts...)

	handler := msg.Handler{GetUnReadHandler: server}

	grpcListener, err := net.Listen("tcp", svr.MsgSvrAddress)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))

	pb.RegisterMsgServer(grpcServer, handler)

	err = grpcServer.Serve(grpcListener)

	if err != nil {
		fmt.Println(err)
	}
}

//
//func newTracer(hostPort, serviceName string) (zipkintracer.Collector, opentracing.Tracer, error) {
//	// 收集器
//	collector, err := zipkin.NewHTTPCollector(zipkinhttpurl)
//	if err != nil {
//		return nil, nil, err
//	}
//	//log.Println(collector)
//	// 记录器
//	recorder := zipkin.NewRecorder(collector, true, hostPort, serviceName)
//	//log.Println(recorder)
//
//	// 追踪器
//	tracer, err := zipkin.NewTracer(
//		recorder,
//		zipkin.ClientServerSameSpan(true),
//		zipkin.TraceID128Bit(true),
//	)
//	if err != nil {
//		return nil, nil, err
//	}
//	//全局追踪器，一般一个服务用一个就可以，使用全局追踪器可以使用很多包装过的方法，创建span
//	opentracing.InitGlobalTracer(tracer)
//	return collector, tracer, nil
//}
