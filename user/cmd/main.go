
// Package user User API.
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
// This should demonstrate all the possible comment annotations
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http, https
//     Host: localhost
//     BasePath: /v2
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: John Doe<john.doe@example.com> http://john.doe.com
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
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

//go:generate swagger generate spec -o swagger.json
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
