package global

import (
	"errors"
	"github.com/google/uuid"
	"math/rand"
)

var userTableData = make(map[string]UserTable)

type UserTable struct {
	Uid string
	Pwd string
	Sid string
}

var accountTableData = make(map[string]AccountTable)

type AccountTable struct {
	Balance int64
}

var sessionData = make(map[string]string)

func init() {
	userTableData["admin"] = UserTable{Pwd: "123456", Uid: uuid.New().String(), Sid: "1"}
	userTableData["jack"] = UserTable{Pwd: "123456", Uid: uuid.New().String(), Sid: "2"}
	userTableData["tom"] = UserTable{Pwd: "123456", Uid: uuid.New().String(), Sid: "3"}
	userTableData["cat"] = UserTable{Pwd: "123456", Uid: uuid.New().String(), Sid: "4"}

	sessionData["1"] = "admin"
	sessionData["2"] = "jack"
	sessionData["3"] = "tom"
	sessionData["4"] = "cat"

	accountTableData["admin"] = AccountTable{rand.Int63n(100)}
	accountTableData["jack"] = AccountTable{rand.Int63n(100)}
	accountTableData["tom"] = AccountTable{rand.Int63n(100)}
	accountTableData["cat"] = AccountTable{rand.Int63n(100)}
}

func GetUsernameBySid(sid string) (string, error) {
	if username, ok := sessionData[sid]; ok {
		return username, nil
	}
	return "", errors.New("no this sid")
}

func GetUserByUsername(username string) (UserTable, error) {

	if u, ok := userTableData[username]; ok {
		return u, nil
	}
	return UserTable{}, errors.New("no this user")
}

func GetAccountByUsername(username string) (AccountTable, error) {
	if a, ok := accountTableData[username]; ok {
		return a, nil
	}
	return AccountTable{}, errors.New("no this user")
}
