package api

import (
	"errors"
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

func MakeLoginEndpoint(api Api) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		if r, ok := request.(ReqLogin); ok {
			return api.Login(r)
		}
		return nil, errors.New("Error Req Type")
	}
}
