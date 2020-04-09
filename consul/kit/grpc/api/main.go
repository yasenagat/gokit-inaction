package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"gitee.com/godY/gokit-inaction/consul/kit/grpc/svr/kit"
	"gitee.com/godY/gokit-inaction/consul/kit/grpc/svr/pb"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/lb"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/hashicorp/consul/api"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"net/http"
	"os"
	"time"
)

type ReqLogin struct {
	Username string `json:"username"`
	Pwd      string `json:"pwd"`
}

type Res struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

//curl -X POST "http://localhost:7788/login" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"username\": \"Tom\", \"Pwd\": \"123456\"}"

//curl -X POST "http://localhost:7788/login" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"username\": \"abc\", \"Pwd\": \"123\"}"
func main() {

	endpoint := newEndpoint()

	s := kithttp.NewServer(func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := pb.LoginReq{}
		if r, ok := request.(ReqLogin); ok {
			req.Username = r.Username
			req.Pwd = r.Pwd
		} else {
			return nil, errors.New("Err Req Type")
		}

		return endpoint(ctx, &req)
	}, func(i context.Context, req *http.Request) (request interface{}, err error) {
		u := ReqLogin{}
		e := json.NewDecoder(req.Body).Decode(&u)

		return u, e
	}, func(i context.Context, writer http.ResponseWriter, v interface{}) error {
		res := Res{}
		if r, ok := v.(*pb.LoginRes); ok {
			res.Msg = r.Msg
			res.Code = r.Code
		}
		return json.NewEncoder(writer).Encode(&res)
	})
	r := mux.NewRouter()
	r.MethodNotAllowedHandler = kithttp.NewServer(func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return nil, nil
	}, func(i context.Context, req *http.Request) (request interface{}, err error) {
		return nil, nil
	}, func(i context.Context, writer http.ResponseWriter, v interface{}) error {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return errors.New("http.StatusMethodNotAllowed")
	})
	r.NotFoundHandler = kithttp.NewServer(func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return nil, nil
	}, func(i context.Context, req *http.Request) (request interface{}, err error) {
		return nil, nil
	}, func(i context.Context, writer http.ResponseWriter, v interface{}) error {
		writer.WriteHeader(http.StatusNotFound)
		return errors.New("http.StatusNotFound")
	})
	r.Path("/login").Methods(http.MethodPost).Handler(s)
	http.ListenAndServe(":7788", r)
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
