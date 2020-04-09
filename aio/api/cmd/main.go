package main

import (
	"fmt"
	"gitee.com/godY/gokit-inaction/aio/api"
	"gitee.com/godY/gokit-inaction/aio/pro"
	"github.com/go-kit/kit/endpoint"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"net/http"
	"os"

	"gitee.com/godY/gokit-inaction/aio/svr/global"
	"golang.org/x/net/context"
)

//curl -X POST "http://localhost:6666/login" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"username\": \"admin\",\"pwd\":\"123\"}"

var UserAddress = ":19911"
var AccountAddress = ":19922"

func main() {

	userConn, e := grpc.Dial(UserAddress, grpc.WithInsecure())
	defer userConn.Close()
	if e != nil {
		fmt.Println(e)
		os.Exit(-1)
	}

	accountConn, e := grpc.Dial(AccountAddress, grpc.WithInsecure())
	defer accountConn.Close()
	if e != nil {
		fmt.Println(e)
		os.Exit(-1)
	}

	userclient := NewUserClient(userConn)
	accountclient := NewAccountClient(accountConn)

	apiSvr := api.ApiSvr{UserClient: userclient, AccountClient: accountclient}

	loginEndpoint := api.MakeLoginEndpoint(apiSvr)

	loginSvr := kithttp.NewServer(loginEndpoint, api.DecodeLoginRequest, api.EncodeResponse)

	r := mux.NewRouter()
	r.Handle("/login", loginSvr)
	errc := make(chan error)
	go func() {
		errc <- http.ListenAndServe(":6666", r)
	}()
	fmt.Println(<-errc)
}

type UserClient struct {
	LoginEndpoint endpoint.Endpoint
}

func NewUserClient(conn *grpc.ClientConn) pro.UserServer {

	loginEndpoint := kitgrpc.NewClient(conn, "pro.User", "Login", global.NoEncodeRequestFunc, global.NoDecodeResponseFunc, pro.LoginRes{}).Endpoint()

	return &UserClient{LoginEndpoint: loginEndpoint}
}

func (uc UserClient) Login(ctx context.Context, req *pro.LoginReq) (*pro.LoginRes, error) {
	res, e := uc.LoginEndpoint(ctx, req)
	if e != nil {
		return nil, e
	}
	return res.(*pro.LoginRes), nil
}

type AccontClient struct {
	GetAccountEndpoint endpoint.Endpoint
}

func NewAccountClient(conn *grpc.ClientConn) pro.AccountServer {

	getAccountEndpoint := kitgrpc.NewClient(conn, "pro.Account", "GetAccount", global.NoEncodeRequestFunc, global.NoDecodeResponseFunc, pro.GetAccountRes{}).Endpoint()

	return &AccontClient{GetAccountEndpoint: getAccountEndpoint}
}

func (client AccontClient) GetAccount(ctx context.Context, req *pro.GetAccountReq) (*pro.GetAccountRes, error) {
	res, e := client.GetAccountEndpoint(ctx, req)
	if e != nil {
		return nil, e
	}
	return res.(*pro.GetAccountRes), nil
}
