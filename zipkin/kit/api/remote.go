package api

import (
	"github.com/go-kit/kit/endpoint"
	"gitee.com/godY/gokit-inaction/zipkin/kit/svr/pro"
	"google.golang.org/grpc"
	"gitee.com/godY/gokit-inaction/zipkin/kit/svr"
	"golang.org/x/net/context"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"fmt"
)

type UserClient struct {
	LoginEndpoint endpoint.Endpoint
}

func (c UserClient) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRes, error) {

	res, e := c.LoginEndpoint(ctx, req)
	if e != nil {
		return nil, e
	}
	return res.(*pb.LoginRes), nil
}

func NewUserClient() (pb.UserServer, error) {

	conn, e := grpc.Dial(svr.UserSvrAddress, grpc.WithInsecure())
	if e != nil {
		fmt.Println(e)
		return &UserClient{}, e
	}
	LoginEndpoint := kitgrpc.NewClient(conn, "pb.User", "Login", svr.NoEncodeRequestFunc, svr.NoDecodeResponseFunc, pb.LoginRes{}).Endpoint()

	return &UserClient{LoginEndpoint: LoginEndpoint}, nil
}
