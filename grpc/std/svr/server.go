package main

import (
	"google.golang.org/grpc"
	"gitee.com/godY/gokit-inaction/grpc/pb"
	"net"
	"fmt"
	"gitee.com/godY/gokit-inaction/grpc/biz"
)

func main() {

	grpcSvr := grpc.NewServer()

	pb.RegisterUserServer(grpcSvr, biz.UserServer{})

	c := make(chan error)

	listen, e := net.Listen("tcp", ":8082")

	if e != nil {
		c <- e
	}

	go func() {
		c <- grpcSvr.Serve(listen)
	}()

	fmt.Println(<-c)

}
