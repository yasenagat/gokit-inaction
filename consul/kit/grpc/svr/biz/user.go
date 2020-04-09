package biz

import (
	"fmt"
	"gitee.com/godY/gokit-inaction/consul/kit/grpc/svr/pb"
	"golang.org/x/net/context"
	hv1 "google.golang.org/grpc/health/grpc_health_v1"
	"net"
)

type UserServer struct {
}

func (UserServer) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRes, error) {
	res := pb.LoginRes{}
	fmt.Println(req.Username)
	fmt.Println(req.Pwd)

	ip, _ := getIp()
	fmt.Println("ip", ip)
	if req.Username == "abc" && req.Pwd == "123" {
		res.Code = 0
		res.Msg = ip + "=登录成功"
	} else {
		res.Code = -1
		res.Msg = ip + "=登录失败"
	}
	return &res, nil
}
func (UserServer) Check(ctx context.Context, req *hv1.HealthCheckRequest) (*hv1.HealthCheckResponse, error) {
	res := hv1.HealthCheckResponse{}
	res.Status = hv1.HealthCheckResponse_SERVING
	return &res, nil
}
func getIp() (string, error) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", nil
}
