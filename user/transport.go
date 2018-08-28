package user

import (
	"net/http"
	"golang.org/x/net/context"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type BaseRes struct {
	Code int
	Msg  string
}

type LoginReq struct {
	Username string
	Pwd      string
}

type LoginRes struct {
	Body User
	Code int
	Msg  string
}

type UpdatePhoneReq struct {
	Username string
	Phone    string
}

type UpdatePhoneRes struct {
	Code int
	Msg  string
}

type GetUserReq struct {
	Username string
}

type GetUserRes struct {
	Body User
	Code int
	Msg  string
}

func DecodeLoginReq(_ context.Context, r *http.Request) (request interface{}, err error) {
	req := LoginReq{}
	bytes, e := ioutil.ReadAll(r.Body)
	if e != nil {
		return nil, e
	}
	fmt.Println(string(bytes))
	if e := json.Unmarshal(bytes, &req); e != nil {
		fmt.Println(e)
		return nil, e
	}
	return req, nil
}

func DefaultEncodeResponse(_ context.Context, res http.ResponseWriter, i interface{}) error {
	return json.NewEncoder(res).Encode(i)
}

func DecodeUpdatePhoneReq(_ context.Context, r *http.Request) (request interface{}, err error) {
	req := UpdatePhoneReq{}
	bytes, e := ioutil.ReadAll(r.Body)
	if e != nil {
		return nil, e
	}
	fmt.Println(string(bytes))
	if e := json.Unmarshal(bytes, &req); e != nil {
		fmt.Println(e)
		return nil, e
	}
	return req, nil
}

func DecodeGetUserReq(_ context.Context, r *http.Request) (request interface{}, err error) {
	req := GetUserReq{}
	bytes, e := ioutil.ReadAll(r.Body)
	if e != nil {
		return nil, e
	}
	fmt.Println(string(bytes))
	if e := json.Unmarshal(bytes, &req); e != nil {
		fmt.Println(e)
		return nil, e
	}
	return req, nil
}
