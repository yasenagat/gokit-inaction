package main

import (
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
	"fmt"
	"os"
	"gitee.com/godY/gokit-inaction/consul/kit/grpc/svr/kit"
	"gitee.com/godY/gokit-inaction/consul/kit/grpc/svr/pb"
	"golang.org/x/net/context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"flag"
	"github.com/hashicorp/consul/api"
	"github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd"
	"io"
	"github.com/go-kit/kit/sd/lb"
	"time"
)

func main() {

	us := NewUserClient()

	for i := 0; i < 10; i++ {

		req := pb.LoginReq{}
		req.Username = "abc"
		req.Pwd = "123"

		res, e := us.Login(context.Background(), &req)

		if e != nil {
			fmt.Println(e)
			return
		}

		fmt.Println(res.Code, res.Msg)
	}
}

type UserClient struct {
	LoginEndpoint endpoint.Endpoint
}

func NewUserClient() pb.UserServer {

	loginEndpoint := newEndpoint()

	return &UserClient{LoginEndpoint: loginEndpoint}
}

func (uc UserClient) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRes, error) {
	res, e := uc.LoginEndpoint(ctx, req)
	if e != nil {
		return nil, e
	}
	return res.(*pb.LoginRes), nil
}

func newEndpoint() endpoint.Endpoint {

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	consuladdr := flag.String("consul.addr", "192.168.10.210:8500", "")

	flag.Parse()

	cfg := api.DefaultConfig()
	cfg.Address = *consuladdr
	c, e := api.NewClient(cfg)

	if e != nil {
		logger.Log("err", e)
		os.Exit(-1)
	}

	kitc := consul.NewClient(c)

	instancer := consul.NewInstancer(kitc, logger, "GrpcUserSvr", nil, true)

	endpointer := sd.NewEndpointer(instancer, func(instance string) (endpoint.Endpoint, io.Closer, error) {
		logger.Log("instance", instance)

		conn, e := grpc.Dial(instance, grpc.WithInsecure())

		//defer conn.Close()

		if e != nil {
			fmt.Println(e)
			os.Exit(-1)
		}

		return kitgrpc.NewClient(conn, "pb.User", "Login", kit.NoEncodeRequestFunc, kit.NoDecodeResponseFunc, pb.LoginRes{}).Endpoint(), conn, nil

	}, logger)

	balancer := lb.NewRoundRobin(endpointer)

	retry := lb.Retry(3, 5000*time.Millisecond, balancer)

	return retry
}
