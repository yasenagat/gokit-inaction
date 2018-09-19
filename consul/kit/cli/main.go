package main

import (
	"github.com/go-kit/kit/log"
	"os"
	"github.com/hashicorp/consul/api"
	"github.com/go-kit/kit/sd/consul"
	"flag"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/endpoint"
	"io"
	kithttp "github.com/go-kit/kit/transport/http"
	"net/url"
	"net/http"
	"golang.org/x/net/context"
	"github.com/go-kit/kit/sd/lb"
	"io/ioutil"
	"strings"
	"time"
)

func main() {

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

	instancer := consul.NewInstancer(kitc, logger, "time server", nil, true)

	endpointer := sd.NewEndpointer(instancer, func(instance string) (endpoint.Endpoint, io.Closer, error) {
		logger.Log("instance", instance)
		tgt, e := url.Parse(instance)

		if e != nil {
			logger.Log("err", e)
		}

		return kithttp.NewClient("GET", tgt, func(context context.Context, request *http.Request, i interface{}) error {

			request.Body = ioutil.NopCloser(strings.NewReader(""))
			return nil
		}, func(i context.Context, res *http.Response) (response interface{}, err error) {

			defer res.Body.Close()

			byte, e := ioutil.ReadAll(res.Body)

			if e != nil {
				logger.Log("err", e)
			}

			return string(byte), nil
		}).Endpoint(), nil, nil

	}, logger)

	balancer := lb.NewRoundRobin(endpointer)

	retry := lb.Retry(3, 5000*time.Millisecond, balancer)

	for i := 0; i < 3; i++ {
		//end, e := balancer.Endpoint()
		//if e != nil {
		//	logger.Log("err", e)
		//}
		//res, e := end(context.Background(), nil)
		res, e := retry(context.Background(), i)

		if e != nil {
			logger.Log("[err]", e)
		}

		logger.Log("res", res)
	}

}
