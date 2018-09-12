package svr

import (
	"os"
	"github.com/go-kit/kit/log"
	"github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
	transporthttp "github.com/go-kit/kit/transport/http"
	kitzipkin "github.com/go-kit/kit/tracing/zipkin"
)

var MsgSvrName = "MsgSvr"
var MsgSvrAddress = ":9988"
var UserSvrAddress = ":9977"
var UserSvrName = "UserSvr"

//var Zipkinhttpurl = "http://localhost:9411/api/v1/spans"
var Zipkinhttpurl = "http://192.168.3.125:9411/api/v1/spans"

func NewLogger(servicename string) log.Logger {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.WithPrefix(logger, "service", servicename)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)
	return logger
}

func NewZipkinTracer(serviceName, hostPort, zipkinhttpurl string, logger log.Logger) *zipkin.Tracer {

	var zipkinTracer *zipkin.Tracer
	{
		var (
			err           error
			useNoopTracer = (zipkinhttpurl == "")
			reporter      = zipkinhttp.NewReporter(zipkinhttpurl)
		)
		defer reporter.Close()
		zEP, err := zipkin.NewEndpoint(serviceName, "localhost:80")
		if err != nil {
			logger.Log("err", err)
			os.Exit(1)
		}
		zipkinTracer, err = zipkin.NewTracer(
			reporter, zipkin.WithLocalEndpoint(zEP), zipkin.WithNoopTracer(useNoopTracer),
		)
		if err != nil {
			logger.Log("err", err)
			os.Exit(1)
		}
		if !useNoopTracer {
			logger.Log("hostPort", hostPort, "serviceName", serviceName, "tracer", "Zipkin", "type", "Native", "URL", zipkinhttpurl)
		}
	}

	return zipkinTracer
}

func NewServerOptions(tracer *zipkin.Tracer, logger log.Logger) []transporthttp.ServerOption {

	zipkinServer := kitzipkin.HTTPServerTrace(tracer)
	options := []transporthttp.ServerOption{
		transporthttp.ServerErrorLogger(logger),
		zipkinServer,
	}
	return options
}
