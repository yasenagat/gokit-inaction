package user

import (
	"github.com/go-kit/kit/log"
	"time"
)

type LogService struct {
	Logger log.Logger
	Service
}

func (ls LogService) Login(username, pwd string) (u User, e error) {

	defer func(start time.Time) {

		ls.Logger.Log("method", "login", "username", username, "pwd", pwd, "user", u, "err", e, "time", time.Since(start))

	}(time.Now())

	return ls.Service.Login(username, pwd)
}

func (ls LogService) UpdatePhone(username, phone string) (e error) {

	defer func(start time.Time) {

		ls.Logger.Log("method", "updatephone","username", username, "phone", phone, "err", e, "time", time.Since(start))

	}(time.Now())
	return ls.Service.UpdatePhone(username, phone)
}
