package kit

import (
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
	"gitee.com/godY/gokit-inaction/consul/kit/grpc/svr/pb"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/pkg/errors"
	hv1 "google.golang.org/grpc/health/grpc_health_v1"
)

func MakeLoginEndpoint(server pb.UserServer) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		if r, ok := request.(*pb.LoginReq); ok {
			return server.Login(ctx, r)
		}

		return &pb.LoginRes{}, errors.New("Request Type Error")
	}
}

func MakeHealthCheckEndpoint(server hv1.HealthServer) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		if r, ok := request.(*hv1.HealthCheckRequest); ok {
			return server.Check(ctx, r)
		}

		return &hv1.HealthCheckResponse{}, errors.New("Request Type Error")
	}
}

type UserHandler struct {
	LoginHandler grpctransport.Handler
	CheckHandler grpctransport.Handler
}

func (us UserHandler) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRes, error) {

	_, i, err := us.LoginHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return i.(*pb.LoginRes), nil
}

func (us UserHandler) Check(ctx context.Context, req *hv1.HealthCheckRequest) (*hv1.HealthCheckResponse, error) {
	_, i, err := us.CheckHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return i.(*hv1.HealthCheckResponse), nil
}
