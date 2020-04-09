package account

import "gitee.com/godY/gokit-inaction/aio/svr/global"

func GetAccount(sid string) (global.AccountTable, error) {

	username, err := global.GetUsernameBySid(sid)
	if err != nil {
		return global.AccountTable{0}, err
	}
	act, err := global.GetAccountByUsername(username)
	if err != nil {
		return global.AccountTable{0}, err
	}
	return act, nil
}
