package main

import (
	"gitee.com/godY/gokit-inaction/zipkin/kit/api"
	transporthttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	_ "net/http/pprof"
	"gitee.com/godY/gokit-inaction/zipkin/kit/svr"
	"github.com/go-kit/kit/tracing/zipkin"
)

//curl -X POST "http://localhost:8888/login" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"username\": \"admin\",\"pwd\":\"123\"}"

var Address = "localhost:8888"

func main() {

	logger := svr.NewLogger("api")

	ser := api.ApiSvr{logger}

	tracer := svr.NewZipkinTracer("api", Address, svr.Zipkinhttpurl, logger)
	opt := svr.NewServerOptions(tracer, logger)

	login := api.MakeLoginEndpoint(ser)
	login = zipkin.TraceEndpoint(tracer, "Login")(login)

	loginSvr := transporthttp.NewServer(login, api.DecodeLoginReq, api.EncodeRes)

	r := mux.NewRouter()
	r.Handle("/login", loginSvr)
	//r.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
	//r.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
	//r.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
	//r.Handle("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
	//r.Handle("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))
	http.Handle("/", r)
	errc := make(chan error)
	go func() {
		errc <- http.ListenAndServe(Address, nil)
	}()
	fmt.Println(<-errc)
}
