package api

import "golang.org/x/net/context"

type Api interface {
	Login(ctx context.Context, req ReqLogin) (ResLogin, error)
}


