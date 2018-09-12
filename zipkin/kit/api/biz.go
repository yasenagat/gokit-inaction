package api

import (
	"golang.org/x/net/context"
	"gitee.com/godY/gokit-inaction/zipkin/kit/svr/pro"
	"github.com/go-kit/kit/log"
)

type ApiSvr struct {
	Logger log.Logger
}

func (svr ApiSvr) Login(req ReqLogin) (ResLogin, error) {

	svr.Logger.Log(req.Username, req.Pwd)

	res := ResLogin{}

	//call rpc start

	userClient, e := NewRemote(svr.Logger).NewUserClient()

	if e != nil {
		svr.Logger.Log("error", e)
		res.Msg = e.Error()
		res.Code = -1
		res.UID = "-1"
		res.Unread = -1
		return res, e
	}

	r := pb.LoginReq{}
	r.Username = req.Username
	r.Pwd = req.Pwd
	loginRes, e := userClient.Login(context.Background(), &r)

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
