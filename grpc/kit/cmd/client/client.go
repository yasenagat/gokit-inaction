package main

import (
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
	"fmt"
	"os"
	"gitee.com/godY/gokit-inaction/grpc/kit"
	"gitee.com/godY/gokit-inaction/grpc/pb"
	"golang.org/x/net/context"
	"github.com/go-kit/kit/endpoint"
)

func main() {
	conn, e := grpc.Dial(":8082", grpc.WithInsecure())
	if e != nil {
		fmt.Println(e)
		os.Exit(-1)
	}

	us := NewUserClient(conn)

	req := pb.LoginReq{}
	req.Username = "abc"
	req.Pwd = "abc"

	res, e := us.Login(context.Background(), &req)

	if e != nil {
		fmt.Println(e)
		return
	}

	fmt.Println(res.Code)
	fmt.Println(res.Err)
}

type UserClient struct {
	LoginEndpoint endpoint.Endpoint
}

func NewUserClient(conn *grpc.ClientConn) pb.UserServer {

	loginEndpoint := kitgrpc.NewClient(conn, "pb.User", "Login", kit.NoEncodeRequestFunc, kit.NoDecodeResponseFunc, pb.LoginRes{}).Endpoint()

	return &UserClient{LoginEndpoint: loginEndpoint}
}

func (uc UserClient) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRes, error) {
	res, e := uc.LoginEndpoint(ctx, req)
	if e != nil {
		return nil, e
	}
	return res.(*pb.LoginRes), nil
}
