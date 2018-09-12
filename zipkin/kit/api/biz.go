package api

import (
	"golang.org/x/net/context"
	"gitee.com/godY/gokit-inaction/zipkin/kit/svr/pro"
	"github.com/go-kit/kit/log"
)

type ApiSvr struct {
	Logger     log.Logger
	UserClient pb.UserServer
}

func (svr ApiSvr) Login(ctx context.Context, req ReqLogin) (ResLogin, error) {

	svr.Logger.Log(req.Username, req.Pwd)

	res := ResLogin{}

	//call rpc start

	r := pb.LoginReq{}
	r.Username = req.Username
	r.Pwd = req.Pwd
	loginRes, e := svr.UserClient.Login(ctx, &r)

	if e != nil {
		svr.Logger.Log("error", e)
		res.Msg = e.Error()
		res.Code = -1
		res.UID = "-1"
		res.Unread = -1
		return res, e
	}

	res.Msg = loginRes.Msg
	res.Code = int(loginRes.Code)

	if loginRes.Code == 0 {
		res.UID = loginRes.Body.Userid
		res.Unread = int(loginRes.Body.UnreadCount)
	}

	//call rpc end

	return res, nil
}
