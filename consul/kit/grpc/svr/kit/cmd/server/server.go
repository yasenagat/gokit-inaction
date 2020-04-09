package main

import (
	"gitee.com/godY/gokit-inaction/consul/kit/grpc/svr/biz"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"net"
	"os"

	"flag"
	"fmt"
	"gitee.com/godY/gokit-inaction/consul/kit/grpc/svr/kit"
	"gitee.com/godY/gokit-inaction/consul/kit/grpc/svr/pb"
	"github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	hv1 "google.golang.org/grpc/health/grpc_health_v1"
	"log"
)

//nohup ./usersvr -svr.reg.check=192.168.10.210:15555 -svr.reg.addr=192.168.10.210 > usersvr.out &
func main() {

	addr := flag.String("grpc.addr", ":15555", "")
	consuladdr := flag.String("consul.addr", "localhost:8500", "")

	name := flag.String("svr.reg.name", "GrpcUserSvr", "")
	address := flag.String("svr.reg.addr", "192.168.10.37", "")
	port := flag.Int("svr.reg.port", 15555, "")
	check := flag.String("svr.reg.check", "192.168.10.37:15555", "")

	flag.Parse()

	us := biz.UserServer{}

	loginEndpoint := kit.MakeLoginEndpoint(us)
	healthCheckEndpoint := kit.MakeHealthCheckEndpoint(us)

	loginServer := kitgrpc.NewServer(loginEndpoint, kit.NoDecodeRequestFunc, kit.NoEncodeResponseFunc)
	healthCheckServer := kitgrpc.NewServer(healthCheckEndpoint, kit.NoDecodeRequestFunc, kit.NoEncodeResponseFunc)

	userHandler := kit.UserHandler{LoginHandler: loginServer, CheckHandler: healthCheckServer}

	grpcListener, err := net.Listen("tcp", *addr)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))

	pb.RegisterUserServer(grpcServer, userHandler)

	hv1.RegisterHealthServer(grpcServer, userHandler)

	cfg := api.DefaultConfig()
	cfg.Address = *consuladdr
	c, e := api.NewClient(cfg)

	if e != nil {
		log.Println(e)
		os.Exit(-1)
	}

	kitc := consul.NewClient(c)
	r := &api.AgentServiceRegistration{Name: *name, Port: *port, Address: *address, Check: &api.AgentServiceCheck{GRPC: *check, Interval: "10s"}}
	e = kitc.Register(r)
	if e != nil {
		log.Println(e)
		os.Exit(-1)
	}

	err = grpcServer.Serve(grpcListener)

	if err != nil {
		fmt.Println(err)
	}

}
