package api

import (
	"log"
	"fmt"
	"golang.org/x/net/context"
	"gitee.com/godY/gokit-inaction/zipkin/kit/svr/pro"
)

type ApiSvr struct {
}

func (svr ApiSvr) Login(req ReqLogin) (ResLogin, error) {

	log.Println(req.Username, req.Pwd)

	res := ResLogin{}

	//call rpc start

	userClient, e := NewUserClient()

	if e != nil {
		fmt.Println(e)
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
		fmt.Println(e)
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
