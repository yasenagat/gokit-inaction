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
	"github.com/openzipkin/zipkin-go/reporter/http"
	opzipkin "github.com/openzipkin/zipkin-go"
	"github.com/go-kit/kit/tracing/zipkin"
	"github.com/go-kit/kit/tracing/opentracing"
	opentracing2 "github.com/opentracing/opentracing-go"
	"github.com/go-kit/kit/log"
)

func main() {

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)

	us := biz.UserServer{}

	reporter := http.NewReporter("http://192.168.3.125:9411/api/v2/spans")
	defer reporter.Close()

	zEP, _ := opzipkin.NewEndpoint("login", "localhost:8082")
	zkTracer, err := opzipkin.NewTracer(reporter, opzipkin.WithLocalEndpoint(zEP))
	if err != nil {
		fmt.Println(err)
	}

	zkServerTrace := zipkin.GRPCServerTrace(zkTracer, zipkin.Name("login"))

	before := kitgrpc.ServerBefore(opentracing.GRPCToContext(opentracing2.GlobalTracer(), "Sum", logger))

	options := []kitgrpc.ServerOption{
		kitgrpc.ServerErrorLogger(logger),
		zkServerTrace,
		before,
	}

	loginEndpoint := kit.MakeLoginEndpoint(us)

	loginServer := kitgrpc.NewServer(loginEndpoint, kit.NoDecodeRequestFunc, kit.NoEncodeResponseFunc, options...)

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
