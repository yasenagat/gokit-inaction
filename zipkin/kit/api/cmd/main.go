package main

import (
	"gitee.com/godY/gokit-inaction/zipkin/kit/api"
	_ "net/http/pprof"
	"gitee.com/godY/gokit-inaction/zipkin/kit/svr"
	"github.com/openzipkin/zipkin-go"
	openzipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
	"os"
	"github.com/go-kit/kit/log"
	"net/http"
	"fmt"
	transporthttp "github.com/go-kit/kit/transport/http"
	kitzipkin "github.com/go-kit/kit/tracing/zipkin"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

//curl -X POST "http://localhost:8888/login" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"username\": \"admin\",\"pwd\":\"123\"}"

var Address = "localhost:8888"
var ServiceName = "apiSvr"

func main() {

	logger := svr.NewLogger("api")



	var zipkinTracerUserSvr *zipkin.Tracer
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
		zipkinTracerUserSvr, err = zipkin.NewTracer(
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

	conn, e := grpc.Dial(svr.UserSvrAddress, grpc.WithInsecure())
	defer conn.Close()
	if e != nil {
		logger.Log("err", e)
	}

	userClient := api.NewRemote(logger, zipkinTracerUserSvr).NewUserClient(conn)

	ser := api.ApiSvr{Logger: logger, UserClient: userClient}


	var zipkinTracerApiSvr *zipkin.Tracer
	{
		var (
			err           error
			hostPort      = Address
			serviceName   = ServiceName
			useNoopTracer = false
			reporter      = openzipkinhttp.NewReporter(svr.Zipkinhttpurl)
		)
		defer reporter.Close()
		zEP, _ := zipkin.NewEndpoint(serviceName, hostPort)
		zipkinTracerApiSvr, err = zipkin.NewTracer(
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
	//zipkinTracer = createTrace(logger)

	opt := svr.NewHttpServerOptions(zipkinTracerApiSvr, "", logger)

	login := api.MakeLoginEndpoint(ser)

	login = kitzipkin.TraceEndpoint(zipkinTracerApiSvr, "LoginEndpoint")(login)

	loginSvr := transporthttp.NewServer(login, api.DecodeLoginReq, api.EncodeRes, opt...)

	r := mux.NewRouter()
	r.Handle("/login", loginSvr)
	errc := make(chan error)
	go func() {
		errc <- http.ListenAndServe(Address, r)
	}()
	fmt.Println(<-errc)
}

func createTrace(logger log.Logger) *zipkin.Tracer {

	var zipkinTracer *zipkin.Tracer
	{
		var (
			err           error
			hostPort      = "localhost:80"
			serviceName   = "addsvc"
			useNoopTracer = false
			reporter      = openzipkinhttp.NewReporter(svr.Zipkinhttpurl)
		)
		defer reporter.Close()
		zEP, err := zipkin.NewEndpoint(serviceName, hostPort)
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
			logger.Log("tracer", "Zipkin", "type", "Native", "URL", svr.Zipkinhttpurl)
		}
	}

	return zipkinTracer
}
