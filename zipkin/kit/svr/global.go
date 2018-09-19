package svr

import (
	"os"
	"github.com/go-kit/kit/log"
	"github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
	transporthttp "github.com/go-kit/kit/transport/http"
	transportgrpc "github.com/go-kit/kit/transport/grpc"
	kitzipkin "github.com/go-kit/kit/tracing/zipkin"
	"gitee.com/godY/tcgetway/time"
)

var MsgSvrName = "msgsvr"
var MsgSvrAddress = "localhost:9988"
var UserSvrAddress = "localhost:9977"
var UserSvrName = "usersvr"

//var Zipkinhttpurl = "http://localhost:9411/api/v2/spans"
var Zipkinhttpurl = "http://192.168.3.125:9411/api/v2/spans"

func NewLogger(servicename string) log.Logger {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.WithPrefix(logger, "service", servicename)
	logger = log.With(logger, "ts", time.Now())
	logger = log.With(logger, "caller", log.DefaultCaller)
	return logger
}

func NewZipkinTracer(serviceName, hostPort, zipkinhttpurl string, logger log.Logger) (*zipkin.Tracer, error) {

	reporter := zipkinhttp.NewReporter(zipkinhttpurl)
	defer func() {
		reporter.Close()
	}()
	zEP, err := zipkin.NewEndpoint(serviceName, hostPort)
	if err != nil {
		logger.Log("err", err)
		os.Exit(1)
	}
	return zipkin.NewTracer(
		reporter, zipkin.WithLocalEndpoint(zEP), zipkin.WithNoopTracer(zipkinhttpurl == ""),
	)
	//if err != nil {
	//	logger.Log("err", err)
	//	os.Exit(1)
	//}
	//if !(zipkinhttpurl == "") {
	//	logger.Log("hostPort", hostPort, "serviceName", serviceName, "tracer", "Zipkin", "type", "Native", "URL", zipkinhttpurl)
	//}
}

func NewHttpServerOptions(tracer *zipkin.Tracer, name string, logger log.Logger) []transporthttp.ServerOption {

	var zipkinServer transporthttp.ServerOption
	if name == "" {
		zipkinServer = kitzipkin.HTTPServerTrace(tracer, kitzipkin.Logger(logger))
	} else {
		zipkinServer = kitzipkin.HTTPServerTrace(tracer, kitzipkin.Logger(logger), kitzipkin.Name(name))
	}
	options := []transporthttp.ServerOption{
		transporthttp.ServerErrorLogger(logger),
		zipkinServer,
	}
	return options
}

func NewGrpcServerOptions(tracer *zipkin.Tracer, name string, logger log.Logger) []transportgrpc.ServerOption {

	var zipkinServer transportgrpc.ServerOption
	if name == "" {
		zipkinServer = kitzipkin.GRPCServerTrace(tracer, kitzipkin.Logger(logger))
	} else {
		zipkinServer = kitzipkin.GRPCServerTrace(tracer, kitzipkin.Logger(logger), kitzipkin.Name(name))
	}
	options := []transportgrpc.ServerOption{
		transportgrpc.ServerErrorLogger(logger),
		zipkinServer,
	}
	return options
}

func NewGrpcClientOptions(tracer *zipkin.Tracer, name string, logger log.Logger) []transportgrpc.ClientOption {

	var zipkinClient transportgrpc.ClientOption
	if name == "" {
		zipkinClient = kitzipkin.GRPCClientTrace(tracer, kitzipkin.Logger(logger))
	} else {
		zipkinClient = kitzipkin.GRPCClientTrace(tracer, kitzipkin.Logger(logger), kitzipkin.Name(name))
	}
	options := []transportgrpc.ClientOption{
		zipkinClient,
	}
	return options
}
