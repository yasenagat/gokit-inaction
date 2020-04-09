package api

type ReqLogin struct {
	Username string `json:"username"`
	Pwd      string `json:"pwd"`
}

type ResLogin struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	UID     string `json:"uid"`
	SID     string `json:"sid"`
	Balance int64  `json:"balance"`
}
