package user

import (
	"fmt"
	"gitee.com/godY/gokit-inaction/aio/pro"
	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

func MakeLoginEndpoint(api UserApi) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		if r, ok := request.(*pro.LoginReq); ok {
			fmt.Println("msgid", r.Header.Msgid)
			uid, sid, e := api.Login(r.Username, r.Pwd)
			if e != nil {
				return &pro.LoginRes{}, e
			}
			return &pro.LoginRes{Userid: uid, Sid: sid}, nil
		}

		return &pro.LoginRes{}, errors.New("Request Type Error")
	}
}

type Handler struct {
	LoginHandler grpctransport.Handler
}

func (h Handler) Login(ctx context.Context, req *pro.LoginReq) (*pro.LoginRes, error) {

	_, i, err := h.LoginHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return i.(*pro.LoginRes), nil
}
