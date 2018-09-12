package api

import (
	"github.com/go-kit/kit/endpoint"
	"gitee.com/godY/gokit-inaction/zipkin/kit/svr/pro"
	"google.golang.org/grpc"
	"gitee.com/godY/gokit-inaction/zipkin/kit/svr"
	"golang.org/x/net/context"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/kit/log"
)

func NewRemote(logger log.Logger) Remote {
	return Remote{logger}
}

type Remote struct {
	logger log.Logger
}
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

func (r Remote) NewUserClient() (pb.UserServer, error) {

	conn, e := grpc.Dial(svr.UserSvrAddress, grpc.WithInsecure())
	if e != nil {
		r.logger.Log("error", e)
		return &UserClient{}, e
	}
	LoginEndpoint := kitgrpc.NewClient(conn, "pb.User", "Login", svr.NoEncodeRequestFunc, svr.NoDecodeResponseFunc, pb.LoginRes{}).Endpoint()

	return &UserClient{LoginEndpoint: LoginEndpoint}, nil
}
