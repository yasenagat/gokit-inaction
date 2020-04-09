package biz

import (
	"fmt"
	"gitee.com/godY/gokit-inaction/grpc/pb"
	"golang.org/x/net/context"
)

type UserServer struct {
}

func (UserServer) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRes, error) {
	res := pb.LoginRes{}
	fmt.Println(req.Username)
	fmt.Println(req.Pwd)
	if req.Username == "abc" && req.Pwd == "123" {
		res.Code = 0
		res.Msg = "登录成功"
	} else {
		res.Code = -1
		res.Msg = "登录失败"
	}
	return &res, nil
}
