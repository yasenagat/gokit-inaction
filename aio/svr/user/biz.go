package user

type UserApi interface {
	Login(username, pwd string) (uid, sid string, err error)
}

type UserSvr struct {
}

func (UserSvr) Login(username, pwd string) (uid, sid string, err error) {

	r, e := Login(username, pwd)

	if e != nil {
		return "", "", e
	}

	return r.Uid, r.Sid, nil
}
