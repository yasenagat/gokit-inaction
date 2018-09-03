package biz

import (
	"golang.org/x/net/context"
	"gitee.com/godY/gokit-inaction/grpc/pb"
	"fmt"
)

type UserServer struct {
}

func (UserServer) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRes, error) {
	res := pb.LoginRes{}
	fmt.Println(req.Username)
	fmt.Println(req.Pwd)
	if (req.Username == "abc" && req.Pwd == "123") {
		res.Code = 0
		res.Err = "登录成功"
	} else {
		res.Code = -1
		res.Err = "登录失败"
	}
	return &res, nil
}


