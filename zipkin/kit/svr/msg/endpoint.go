package msg

import (
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"

	"gitee.com/godY/gokit-inaction/zipkin/kit/svr/pro"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/pkg/errors"
)

func MakeUnReadEndpoint(server pb.MsgServer) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		if r, ok := request.(*pb.UnReadReq); ok {
			return server.GetUnRead(ctx, r)
		}

		return &pb.UnReadRes{}, errors.New("Request Type Error")
	}
}

type Handler struct {
	GetUnReadHandler grpctransport.Handler
}

func (us Handler) GetUnRead(ctx context.Context, req *pb.UnReadReq) (*pb.UnReadRes, error) {

	_, i, err := us.GetUnReadHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return i.(*pb.UnReadRes), nil
}
