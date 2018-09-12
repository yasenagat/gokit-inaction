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
	"github.com/openzipkin/zipkin-go/reporter/http"
	"github.com/go-kit/kit/tracing/zipkin"
	opzipkin "github.com/openzipkin/zipkin-go"
	"github.com/go-kit/kit/tracing/opentracing"
	opentracing2 "github.com/opentracing/opentracing-go"
	"github.com/go-kit/kit/log"
)

func main() {
	conn, e := grpc.Dial(":8082", grpc.WithInsecure())
	defer conn.Close()
	if e != nil {
		fmt.Println(e)
		os.Exit(-1)
	}

	reporter := http.NewReporter("http://192.168.3.125:9411/api/v2/spans")
	defer reporter.Close()
	zEP, _ := opzipkin.NewEndpoint("login", "localhost:8082")
	zkTracer, err := opzipkin.NewTracer(reporter, opzipkin.WithLocalEndpoint(zEP))
	if err != nil {
		fmt.Println(err)
	}

	us := NewUserClient(conn, zkTracer)

	req := pb.LoginReq{}
	req.Username = "abc"
	req.Pwd = "123"

	//span := zkTracer.StartSpan("call login")
	//defer span.Finish()
	res, e := us.Login(context.Background(), &req)
	//span.Tag("res", res.String())
	if e != nil {
		fmt.Println(e)
		//span.Tag("err", e.Error())
		return
	}

	fmt.Println(res.Code)
	fmt.Println(res.Msg)
}

type UserClient struct {
	LoginEndpoint endpoint.Endpoint
}

func NewUserClient(conn *grpc.ClientConn, trace *opzipkin.Tracer) pb.UserServer {

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)

	zipkinClient := zipkin.GRPCClientTrace(trace)

	zkClientTrace := zipkin.GRPCClientTrace(trace, zipkin.Name("Login"))

	before := kitgrpc.ClientBefore(opentracing.ContextToGRPC(opentracing2.GlobalTracer(), logger))

	// global client middlewares
	options := []kitgrpc.ClientOption{
		zipkinClient,
		before,
		zkClientTrace,
	}

	loginEndpoint := kitgrpc.NewClient(conn, "pb.User", "Login", kit.NoEncodeRequestFunc, kit.NoDecodeResponseFunc, pb.LoginRes{}, options...).Endpoint()

	return &UserClient{LoginEndpoint: loginEndpoint}
}

func (uc UserClient) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRes, error) {
	res, e := uc.LoginEndpoint(ctx, req)
	if e != nil {
		return nil, e
	}
	return res.(*pb.LoginRes), nil
}
