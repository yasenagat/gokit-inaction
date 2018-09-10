package api

type ReqLogin struct {
	Username string `json:"username"`
	Pwd      string `json:"pwd"`
}

type ResLogin struct {
	Code   int    `json:"code"`
	Unread int    `json:"unread"`
	Msg    string `json:"msg"`
	UID    string `json:"uid"`
}
