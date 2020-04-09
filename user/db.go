package user

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"sync"
	"time"
)

var Users = make(map[string]User, 10)
var lock = &sync.Mutex{}

func init() {
	Users["Tom"] = User{Id: uuid.New().String(), Username: "Tom", Age: 11, Phone: "13333331111", Balance: "99", Sid: uuid.New().String(), Pwd: "123456"}

	Users["Jack"] = User{Id: uuid.New().String(), Username: "Jack", Age: 11, Phone: "18899997777", Balance: "11", Sid: uuid.New().String(), Pwd: "123456"}
}

func CheckLogin(username, pwd string) (User, error) {

	for _, v := range Users {

		if v.Username == username && v.Pwd == pwd {
			return v, nil
		}
	}
	time.Sleep(time.Millisecond * 300)
	return User{}, errors.New("用户名或密码错误")
}

func GetUser(username string) (User, error) {
	time.Sleep(time.Millisecond * 200)
	if u, ok := Users[username]; ok {
		return u, nil
	} else {
		return User{}, errors.New("无此用户")
	}
}

func UpdatePhone(username, phone string) error {
	lock.Lock()

	defer func() {
		lock.Unlock()
	}()

	time.Sleep(time.Millisecond * 100)
	user, e := GetUser(username)

	if e != nil {
		return e
	}

	user.Phone = phone
	Users[username] = user
	return nil
}
