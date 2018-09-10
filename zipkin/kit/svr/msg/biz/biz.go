package biz

import (
	"golang.org/x/net/context"
	"gitee.com/godY/gokit-inaction/zipkin/kit/svr/pro"
	"strconv"
	"log"
)

type MsgSvr struct {
}

func (msg MsgSvr) GetUnRead(ctx context.Context, req *pb.UnReadReq) (*pb.UnReadRes, error) {

	log.Println("req.Userid", req.Userid)
	res := pb.UnReadRes{}

	body := pb.UnReadResBody{}

	body.Count = DB[req.Userid]
	log.Println("body.Count", body.Count)
	res.Msg = "ok"
	res.Code = 0
	res.Body = &body

	return &res, nil
}

var DB = make(map[string]int64)

func init() {
	for i := 0; i < 100; i++ {
		DB[strconv.Itoa(i)] = int64(i * 2)
	}
}
