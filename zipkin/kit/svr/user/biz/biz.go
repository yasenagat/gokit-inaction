package biz

import (
	"fmt"
	"gitee.com/godY/gokit-inaction/zipkin/kit/svr"
	"gitee.com/godY/gokit-inaction/zipkin/kit/svr/pro"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/openzipkin/zipkin-go"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type UserSvr struct {
	Logger    log.Logger
	MsgClient pb.MsgServer
}

func (u UserSvr) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRes, error) {

	res := pb.LoginRes{}
	u.Logger.Log("req.Username", req.Username, "req.Pwd", req.Pwd)
	if req.Username == "admin" && req.Pwd == "123" {
		body := pb.LoginResBody{}
		body.Userid = "1"

		//call rpc start
		unreadreq := pb.UnReadReq{}
		unreadreq.Userid = body.Userid
		unReadRes, e := u.MsgClient.GetUnRead(ctx, &unreadreq)

		if e != nil {
			fmt.Println(e)
			body.UnreadCount = 0
		} else {
			if unReadRes.Code == 0 {
				body.UnreadCount = unReadRes.Body.Count
			} else {
				fmt.Println("unReadRes.Msg", unReadRes.Msg)
				body.UnreadCount = 0
			}
		}
		//call rpc end

		res.Body = &body
		res.Code = 0
		res.Msg = "登录成功"
	} else {
		res.Code = -1
		res.Msg = "登录失败"
	}

	return &res, nil
}

type MsgClient struct {
	GetUnReadEndpoint endpoint.Endpoint
}

func (c MsgClient) GetUnRead(ctx context.Context, req *pb.UnReadReq) (*pb.UnReadRes, error) {

	res, e := c.GetUnReadEndpoint(ctx, req)
	if e != nil {
		return nil, e
	}
	return res.(*pb.UnReadRes), nil
}

func NewMsgClient(conn *grpc.ClientConn, zipkinTracer *zipkin.Tracer, logger log.Logger) pb.MsgServer {

	opts := svr.NewGrpcClientOptions(zipkinTracer, "", logger)

	GetUnReadEndpoint := kitgrpc.NewClient(conn, "pb.Msg", "GetUnRead", svr.NoEncodeRequestFunc, svr.NoDecodeResponseFunc, pb.UnReadRes{}, opts...).Endpoint()

	return &MsgClient{GetUnReadEndpoint: GetUnReadEndpoint}
}
