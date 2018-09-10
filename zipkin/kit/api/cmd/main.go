package main

import (
	"gitee.com/godY/gokit-inaction/zipkin/kit/api"
	transporthttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"gitee.com/godY/gokit-inaction/zipkin/kit/svr"
)

//curl -X POST "http://localhost:8888/login" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"username\": \"admin\",\"pwd\":\"123\"}"

var Address = ":8888"

func main() {

	logger := svr.NewLogger()
	ser := api.ApiSvr{}

	login := api.MakeLoginEndpoint(ser)

	svr.NewServerOptions("api", Address, svr.Zipkinhttpurl, logger)

	loginSvr := transporthttp.NewServer(login, api.DecodeLoginReq, api.EncodeRes)

	r := mux.NewRouter()
	r.Handle("/login", loginSvr)

	errc := make(chan error)
	go func() {
		errc <- http.ListenAndServe(Address, r)
	}()
	fmt.Println(<-errc)
}
