package account

import (
	"fmt"
	"gitee.com/godY/gokit-inaction/aio/pro"
	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

func MakeGetAccountEndpoint(api AccountApi) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		if r, ok := request.(*pro.GetAccountReq); ok {
			fmt.Println("msgid", r.Header.Msgid)
			balance, e := api.GetAccount(r.Sid)
			if e != nil {
				return &pro.GetAccountRes{Header: &pro.ResponseHeader{Code: -1}}, e
			}
			return &pro.GetAccountRes{Header: &pro.ResponseHeader{Code: 0}, Balance: balance}, nil
		}

		return &pro.GetAccountRes{Header: &pro.ResponseHeader{Code: -1}}, errors.New("Request Type Error")
	}
}

type Handler struct {
	GetAccountHandler grpctransport.Handler
}

func (h Handler) GetAccount(ctx context.Context, req *pro.GetAccountReq) (*pro.GetAccountRes, error) {

	_, i, err := h.GetAccountHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return i.(*pro.GetAccountRes), nil
}
