package user

import (
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
	"gitee.com/godY/gokit-inaction/zipkin/kit/svr/pro"
	"github.com/pkg/errors"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

func MakeLoginEndpoint(svr pb.UserServer) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		if r, ok := request.(*pb.LoginReq); ok {
			return svr.Login(ctx, r)
		}

		return &pb.LoginRes{}, errors.New("Error Req Type")
	}
}

type Handler struct {
	LoginEndpoint grpctransport.Handler
}

func (h Handler) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRes, error) {

	_, i, err := h.LoginEndpoint.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return i.(*pb.LoginRes), nil
}
