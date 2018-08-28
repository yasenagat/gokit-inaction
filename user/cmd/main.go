package main

import (
	"net/http"
	"gitee.com/godY/gokit-inaction/user"
	"github.com/go-kit/kit/log"
	"os"
	transporthttp "github.com/go-kit/kit/transport/http"
	"encoding/json"
	"golang.org/x/net/context"
)

//curl -X POST "http://localhost:8080/login" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"username\": \"Tom\", \"Pwd\": \"123456\"}"

//curl -X POST "http://localhost:8080/phone" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"username\": \"Tom\", \"Phone\": \"911\"}"

//curl -X POST "http://localhost:8080/user" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"username\": \"Tom\"}"
func main() {

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stdout)
	logger = log.With(logger, "caller", log.DefaultCaller)

	var svr user.Service

	svr = user.UserService{}
	svr = user.LogService{Logger: logger, Service: svr}

	userEndpoint := user.MakeUserEndPoint(svr)

	opts := []transporthttp.ServerOption{
		transporthttp.ServerErrorLogger(logger),
		transporthttp.ServerErrorEncoder(encodeError),
	}

	loginHandler := transporthttp.NewServer(userEndpoint.Login, user.DecodeLoginReq, user.DefaultEncodeResponse, opts...)
	updatePhoneHandler := transporthttp.NewServer(userEndpoint.UpdatePhone, user.DecodeUpdatePhoneReq, user.DefaultEncodeResponse, opts...)
	getUserHandler := transporthttp.NewServer(userEndpoint.GetUser, user.DecodeGetUserReq, user.DefaultEncodeResponse, opts...)

	http.Handle("/login", loginHandler)
	http.Handle("/phone", updatePhoneHandler)
	http.Handle("/user", getUserHandler)
	http.ListenAndServe(":8080", nil)

}

// encode errors from business-logic
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	//case cargo.ErrUnknown:
	//	w.WriteHeader(http.StatusNotFound)
	//case ErrInvalidArgument:
	//	w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
