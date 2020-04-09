package user

import (
	"gitee.com/godY/gokit-inaction/aio/svr/global"
	"github.com/pkg/errors"
)

func Login(username, pwd string) (global.UserTable, error) {

	r, err := global.GetUserByUsername(username)
	if err != nil {
		return global.UserTable{}, err
	}

	if pwd == r.Pwd {
		return r, nil
	}
	return global.UserTable{}, errors.New("Username or Pwd Error!")
}
