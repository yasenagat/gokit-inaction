package user

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	kitlimit "github.com/go-kit/kit/ratelimit"
	"github.com/juju/ratelimit"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"golang.org/x/time/rate"
	"time"
)

const RATELIMIT = 100
const ERRORLIMIT = 10

var rateLimitBucket = ratelimit.NewBucket(time.Second, RATELIMIT)

type UserEndpoint struct {
	Login       endpoint.Endpoint
	UpdatePhone endpoint.Endpoint
	GetUser     endpoint.Endpoint
}

func MakeEndPoint(logger log.Logger, service Service) *UserEndpoint {

	return &UserEndpoint{
		Login:       NewLogginMiddelware(logger)(NewErrorLimitMiddelware()(NewRateLimitMiddelware(rateLimitBucket)(LoginEndpoint(service)))),
		UpdatePhone: NewErrorLimitMiddelware()(NewRateLimitMiddelware(rateLimitBucket)(UpdatePhoneEndpoint(service))),
		GetUser:     NewErrorLimitMiddelware()(NewRateLimitMiddelware(rateLimitBucket)(GetUserEndpoint(service)))}
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

func NewRateLimitMiddelware(b *ratelimit.Bucket) endpoint.Middleware {

	return func(i endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {

			if b.TakeAvailable(1) == 0 {
				return nil, errors.New("访问限制")
			}
			return i(ctx, request)
		}
	}
}

func NewErrorLimitMiddelware() endpoint.Middleware {
	return kitlimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), ERRORLIMIT))
}

func NewDelayingLimitMiddelware() endpoint.Middleware {
	return kitlimit.NewDelayingLimiter(rate.NewLimiter(rate.Every(time.Second), 10))
}

func NewLogginMiddelware(logger log.Logger) endpoint.Middleware {
	return func(i endpoint.Endpoint) endpoint.Endpoint {

		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(start time.Time) {
				logger.Log("time", time.Since(start))
			}(time.Now())
			return i(ctx, request)
		}
	}
}
