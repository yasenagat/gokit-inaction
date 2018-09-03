package user

import (
	"net/http"
	"golang.org/x/net/context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"github.com/go-kit/kit/log"
	"git.oschina.net/janpoem/go-logger.git"
	"github.com/google/uuid"
	"strings"
)

type InnerMsg struct {
	MsgId string
}

type BaseRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type LoginReq struct {
	Username string `json:"username"`
	Pwd      string `json:"pwd"`
}

type LoginRes struct {
	Body User   `json:"body"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type UpdatePhoneReq struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
}

type UpdatePhoneRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type GetUserReq struct {
	Username string `json:"username"`
}

type GetUserRes struct {
	Body User   `json:"body"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func GenerateMsgId() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}

func GetMsgId(ctx context.Context) string {
	return ctx.Value(INNERMSG).(InnerMsg).MsgId
}

const MSGID = "MsgId"
const INNERMSG = "INNERMSG"

func DecodeLoginReq(logger log.Logger) func(context.Context, *http.Request) (interface{}, error) {

	return func(ctx context.Context, r *http.Request) (request interface{}, err error) {
		req := LoginReq{}
		bytes, e := ioutil.ReadAll(r.Body)
		if e != nil {
			return nil, e
		}
		reqlog(logger, r, bytes, ctx)
		if e := json.Unmarshal(bytes, &req); e != nil {
			fmt.Println(e)
			return nil, e
		}
		return req, nil
	}
}

//func DecodeLoginReq_(_ context.Context, r *http.Request) (request interface{}, err error) {
//	req := LoginReq{}
//	bytes, e := ioutil.ReadAll(r.Body)
//	if e != nil {
//		return nil, e
//	}
//	fmt.Println(string(bytes))
//	if e := json.Unmarshal(bytes, &req); e != nil {
//		fmt.Println(e)
//		return nil, e
//	}
//	return req, nil
//}

func DefaultEncodeResponse(logger log.Logger) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, writer http.ResponseWriter, i interface{}) error {

		bytes, e := json.Marshal(i)

		if e != nil {
			return e
		}

		logger.Log(MSGID, GetMsgId(ctx), "Res", string(bytes))
		writer.Write(bytes)

		return nil
	}

}

func DecodeUpdatePhoneReq(logger log.Logger) func(context.Context, *http.Request) (interface{}, error) {

	return func(ctx context.Context, r *http.Request) (interface{}, error) {

		req := UpdatePhoneReq{}
		bytes, e := ioutil.ReadAll(r.Body)
		if e != nil {
			return nil, e
		}
		reqlog(logger, r, bytes, ctx)
		if e := json.Unmarshal(bytes, &req); e != nil {
			fmt.Println(e)
			return nil, e
		}
		return req, nil
	}
}

//func DecodeUpdatePhoneReq_(_ context.Context, r *http.Request) (request interface{}, err error) {
//	req := UpdatePhoneReq{}
//	bytes, e := ioutil.ReadAll(r.Body)
//	if e != nil {
//		return nil, e
//	}
//	fmt.Println(string(bytes))
//	if e := json.Unmarshal(bytes, &req); e != nil {
//		fmt.Println(e)
//		return nil, e
//	}
//	return req, nil
//}

func DecodeGetUserReq(logger log.Logger) func(context.Context, *http.Request) (interface{}, error) {

	return func(ctx context.Context, r *http.Request) (interface{}, error) {

		req := GetUserReq{}

		bytes, e := ioutil.ReadAll(r.Body)
		if e != nil {
			return nil, e
		}
		reqlog(logger, r, bytes, ctx)
		if e := json.Unmarshal(bytes, &req); e != nil {
			fmt.Println(e)
			return nil, e
		}
		return req, nil
	}
}

func reqlog(logger log.Logger, r *http.Request, bytes []byte, ctx context.Context) error {
	return logger.Log(MSGID, GetMsgId(ctx), "RequestURI", r.RequestURI, "Req", string(bytes))
}

//func DecodeGetUserReq_(_ context.Context, r *http.Request) (request interface{}, err error) {
//	req := GetUserReq{}
//	bytes, e := ioutil.ReadAll(r.Body)
//	if e != nil {
//		return nil, e
//	}
//	fmt.Println(string(bytes))
//	if e := json.Unmarshal(bytes, &req); e != nil {
//		fmt.Println(e)
//		return nil, e
//	}
//	return req, nil
//}

func DecodeReq(r *http.Request, i interface{}) (interface{}, error) {

	bytes, e := ioutil.ReadAll(r.Body)
	if e != nil {
		return nil, e
	}
	logger.Log("RequestURI", r.RequestURI, "Req", string(bytes))
	if e := json.Unmarshal(bytes, &i); e != nil {
		fmt.Println(e)
		return nil, e
	}
	return i, nil
}

func InnerMsgRequestContext(ctx context.Context, r *http.Request) context.Context {
	im := InnerMsg{}
	im.MsgId = GenerateMsgId()
	ctx = context.WithValue(ctx, INNERMSG, im)
	return ctx
}
