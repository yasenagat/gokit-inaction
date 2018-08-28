package user

import (
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
	"github.com/pkg/errors"
)

type UserEndpoint struct {
	Login       endpoint.Endpoint
	UpdatePhone endpoint.Endpoint
	GetUser     endpoint.Endpoint
}

func MakeUserEndPoint(service Service) *UserEndpoint {

	return &UserEndpoint{
		Login:       LoginEndpoint(service),
		UpdatePhone: UpdatePhoneEndpoint(service),
		GetUser:     GetUserEndpoint(service)}
}

func LoginEndpoint(service Service) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		if r, ok := request.(LoginReq); ok {
			user, e := service.Login(r.Username, r.Pwd)
			if e != nil {
				return LoginRes{Code: -1, Msg: e.Error()}, e
			} else {
				return LoginRes{Code: 0, Body: user, Msg: "登录成功"}, nil
			}
		}
		return LoginRes{Code: -1, Msg: "LoginEndpoint Error Type"}, errors.New("LoginEndpoint Error Type")
	}
}

func UpdatePhoneEndpoint(service Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		if r, ok := request.(UpdatePhoneReq); ok {
			e := service.UpdatePhone(r.Username, r.Phone)
			if e != nil {
				return UpdatePhoneRes{Code: -1, Msg: e.Error()}, e
			} else {
				return UpdatePhoneRes{Code: 0, Msg: "修改手机成功"}, nil
			}
		}
		return UpdatePhoneRes{Code: -1, Msg: "UpdatePhoneEndpoint Error Type"}, errors.New("UpdatePhoneEndpoint Error Type")
	}
}

func GetUserEndpoint(service Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		if r, ok := request.(GetUserReq); ok {
			user, e := service.GetUser(r.Username)
			if e != nil {
				return GetUserRes{Code: -1, Msg: e.Error()}, e
			} else {
				return GetUserRes{Code: 0, Body: user, Msg: "获取用户成功"}, nil
			}
		}
		return UpdatePhoneRes{Code: -1, Msg: "GetUserEndpoint Error Type"}, errors.New("GetUserEndpoint Error Type")
	}
}
