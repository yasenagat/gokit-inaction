package account

type AccountApi interface {
	GetAccount(sid string) (balance int64, err error)
}

type AccountSvr struct {
}

func (AccountSvr) GetAccount(sid string) (balance int64, err error) {

	r, e := GetAccount(sid)

	if e != nil {
		return 0, e
	}

	return r.Balance, nil
}
