package main

import (
	"net"
	"fmt"
	"os"
	"google.golang.org/grpc"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"gitee.com/godY/gokit-inaction/zipkin/kit/svr/msg/biz"
	"gitee.com/godY/gokit-inaction/zipkin/kit/svr/msg"
	"gitee.com/godY/gokit-inaction/zipkin/kit/svr"
	"gitee.com/godY/gokit-inaction/zipkin/kit/svr/pro"
	kitzipkin "github.com/go-kit/kit/tracing/zipkin"
)

func main() {

	logger := svr.NewLogger()

	service := biz.MsgSvr{}

	endpoint := msg.MakeUnReadEndpoint(service)

	//_, _, e := newTracer("", "")
	//
	//if e != nil {
	//
	//}
	zipkinTracer := svr.NewZipkinTracer(svr.MsgSvrName, svr.MsgSvrAddress, svr.Zipkinhttpurl, logger)

	zipkinServer := kitzipkin.GRPCServerTrace(zipkinTracer)

	options := []kitgrpc.ServerOption{
		kitgrpc.ServerErrorLogger(logger),
		zipkinServer,
	}

	server := kitgrpc.NewServer(endpoint, svr.NoDecodeRequestFunc, svr.NoEncodeResponseFunc, append(options)...)

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
