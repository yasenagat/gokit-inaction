package user

import (
	"github.com/go-kit/kit/log"
	"time"
)

type logService struct {
	logger log.Logger
	Service
}

func NewLogService(logger log.Logger, s Service) Service {
	return &logService{logger, s}
}

func (ls logService) Login(username, pwd string) (u User, e error) {

	defer func(start time.Time) {

		ls.logger.Log("method", "login", "username", username, "pwd", pwd, "user", u, "err", e, "time", time.Since(start))

	}(time.Now())

	return ls.Service.Login(username, pwd)
}

func (ls logService) UpdatePhone(username, phone string) (e error) {

	defer func(start time.Time) {

		ls.logger.Log("method", "updatephone", "username", username, "phone", phone, "err", e, "time", time.Since(start))

	}(time.Now())
	return ls.Service.UpdatePhone(username, phone)
}
