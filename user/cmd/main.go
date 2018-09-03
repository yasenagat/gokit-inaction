package main

import (
	"net/http"
	"gitee.com/godY/gokit-inaction/user"
	"github.com/go-kit/kit/log"
	"os"
	"github.com/gorilla/mux"
	"os/signal"
	"syscall"
	"fmt"
)

//curl -X POST "http://localhost:8080/login" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"username\": \"Tom\", \"Pwd\": \"123456\"}"

//curl -X POST "http://localhost:8080/phone" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"username\": \"Tom\", \"Phone\": \"911\"}"

//curl -X POST "http://localhost:8080/user" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"username\": \"Tom\"}"

// Package classification User API.
//
// The purpose of this service is to provide an application
// that is using plain go code to define an API
//
//      Host: localhost
//      Version: 0.0.1
//
// swagger:meta
func main() {

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stdout)
	logger = log.With(logger, "caller", log.DefaultCaller)

	var svr user.Service

	svr = user.UserService{}
	svr = user.NewLogService(log.With(logger, "component", "user"), svr)
	svr = user.NewmetricsMW(svr)

	userEndpoint := user.MakeEndPoint(logger, svr)

	r := mux.NewRouter()

	r = user.MakeHandler(logger, userEndpoint, r)

	errc := make(chan error, 1)

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
		errc <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		errc <- http.ListenAndServe(":8080", r)
	}()

	logger.Log("err", <-errc)

}
