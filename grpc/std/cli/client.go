package main

import (
	"gitee.com/godY/gokit-inaction/grpc/pb"
	"google.golang.org/grpc"
	"fmt"
	"golang.org/x/net/context"
)

func main() {

	conn, e := grpc.Dial(":8082", grpc.WithInsecure())
	if e != nil {
		fmt.Println(e)
		return
	}

	userClient := pb.NewUserClient(conn)

	req := pb.LoginReq{}
	req.Username = "abc"
	req.Pwd = "123"

	res, e := userClient.Login(context.Background(), &req)

	if e != nil {
		fmt.Println(e)
		return
	}

	fmt.Println(res.Code)
	fmt.Println(res.Msg)
}
