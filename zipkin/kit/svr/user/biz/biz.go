package biz

import (
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"golang.org/x/net/context"
	"gitee.com/godY/gokit-inaction/zipkin/kit/svr/pro"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"google.golang.org/grpc"
	"gitee.com/godY/gokit-inaction/zipkin/kit/svr"
	"github.com/go-kit/kit/log"
)

type UserSvr struct {
	Logger log.Logger
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
		unReadRes, e := GetUnRead(context.Background(), &unreadreq)

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

func GetUnRead(ctx context.Context, req *pb.UnReadReq) (*pb.UnReadRes, error) {

	conn, e := grpc.Dial(svr.MsgSvrAddress, grpc.WithInsecure())
	if e != nil {
		fmt.Println(e)
	}

	svr := NewMsgClient(conn)

	return svr.GetUnRead(ctx, req)
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

func NewMsgClient(conn *grpc.ClientConn) pb.MsgServer {

	//logger := svr.NewLogger()
	//
	//zipkinTracer := svr.NewZipkinTracer(svr.MsgSvrName, svr.MsgSvrAddress, svr.Zipkinhttpurl, logger)
	//
	//zipkinServer := kitzipkin.GRPCClientTrace(zipkinTracer)
	//
	//options := []kitgrpc.ClientOption{
	//	zipkinServer,
	//}

	GetUnReadEndpoint := kitgrpc.NewClient(conn, "pb.Msg", "GetUnRead", svr.NoEncodeRequestFunc, svr.NoDecodeResponseFunc, pb.UnReadRes{}, ).Endpoint()

	return &MsgClient{GetUnReadEndpoint: GetUnReadEndpoint}
}
