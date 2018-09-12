package api

import (
	"github.com/go-kit/kit/endpoint"
	"gitee.com/godY/gokit-inaction/zipkin/kit/svr/pro"
	"google.golang.org/grpc"
	"gitee.com/godY/gokit-inaction/zipkin/kit/svr"
	"golang.org/x/net/context"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/kit/log"
	"github.com/openzipkin/zipkin-go"
)

func NewRemote(logger log.Logger, zipkinTracer *zipkin.Tracer) Remote {
	return Remote{logger: logger, zipkinTracer: zipkinTracer}
}

type Remote struct {
	logger       log.Logger
	zipkinTracer *zipkin.Tracer
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

func (r Remote) NewUserClient(conn *grpc.ClientConn) (pb.UserServer) {

	opts := svr.NewGrpcClientOptions(r.zipkinTracer, "", r.logger)

	LoginEndpoint := kitgrpc.NewClient(conn, "pb.User", "Login", svr.NoEncodeRequestFunc, svr.NoDecodeResponseFunc, pb.LoginRes{}, opts...).Endpoint()

	return &UserClient{LoginEndpoint: LoginEndpoint}
}
