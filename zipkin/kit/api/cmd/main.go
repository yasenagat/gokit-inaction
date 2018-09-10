package main

import (
	"gitee.com/godY/gokit-inaction/zipkin/kit/api"
	transporthttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
)

//curl -X POST "http://localhost:8080/login" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"username\": \"admin\",\"pwd\":\"123\"}"

func main() {

	svr := api.ApiSvr{}

	login := api.MakeLoginEndpoint(svr)

	loginSvr := transporthttp.NewServer(login, api.DecodeLoginReq, api.EncodeRes)

	r := mux.NewRouter()
	r.Handle("/login", loginSvr)

	errc := make(chan error)
	go func() {
		errc <- http.ListenAndServe(":8080", r)
	}()
	fmt.Println(<-errc)
}
