package api

import (
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
	"github.com/pkg/errors"
)

func MakeLoginEndpoint(api Api) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		if r, ok := request.(ReqLogin); ok {
			return api.Login(r)
		}

		return ResLogin{}, errors.New("Error Req Type")
	}
}
