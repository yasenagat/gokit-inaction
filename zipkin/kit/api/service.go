package api

type Api interface {
	Login(ReqLogin) (ResLogin, error)
}


