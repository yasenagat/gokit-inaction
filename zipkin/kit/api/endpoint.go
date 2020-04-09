package api

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

func MakeLoginEndpoint(api Api) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		if r, ok := request.(ReqLogin); ok {
			return api.Login(ctx, r)
		}

		return nil, errors.New("Error Req Type")
	}
}
