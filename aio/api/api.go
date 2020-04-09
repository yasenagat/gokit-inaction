package api

import (
	"context"
	"gitee.com/godY/gokit-inaction/aio/pro"
	"github.com/google/uuid"
)

type Api interface {
	Login(login ReqLogin) (ResLogin, error)
}

type ApiSvr struct {
	UserClient    pro.UserServer
	AccountClient pro.AccountServer
}

func (api ApiSvr) Login(login ReqLogin) (ResLogin, error) {

	ctx := context.Background()

	req := pro.LoginReq{}
	req.Username = login.Username
	req.Pwd = login.Pwd
	req.Header = &pro.RequestHeader{Msgid: uuid.New().String()}
	res, err := api.UserClient.Login(ctx, &req)

	if err != nil {
		return ResLogin{Code: -1}, err
	}

	actReq := pro.GetAccountReq{Sid: res.Sid}
	actReq.Header = &pro.RequestHeader{Msgid: uuid.New().String()}
	ress, err := api.AccountClient.GetAccount(ctx, &actReq)

	if err != nil {
		return ResLogin{Code: -1}, err
	}

	return ResLogin{Code: 0, Balance: ress.Balance, SID: res.Sid, UID: res.Userid}, nil
}
