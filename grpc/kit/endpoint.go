package kit

import (
	"gitee.com/godY/gokit-inaction/grpc/pb"
	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

func MakeLoginEndpoint(server pb.UserServer) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		if r, ok := request.(*pb.LoginReq); ok {
			return server.Login(ctx, r)
		}

		return &pb.LoginRes{}, errors.New("Request Type Error")
	}
}

type UserHandler struct {
	LoginHandler grpctransport.Handler
}

func (us UserHandler) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRes, error) {

	_, i, err := us.LoginHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return i.(*pb.LoginRes), nil
}
