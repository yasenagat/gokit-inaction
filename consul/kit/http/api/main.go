package main

import (
	kithttp "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
	"net/http"
	"github.com/hashicorp/consul/api"
	"github.com/go-kit/kit/log"
	"os"
	"github.com/go-kit/kit/sd/consul"
	"flag"
	"io/ioutil"
	"github.com/go-kit/kit/endpoint"
	"strings"
	"github.com/go-kit/kit/sd/lb"
	"github.com/go-kit/kit/sd"
	"io"
	"net/url"
	"time"
)

func main() {

	endpoint := newEndpoint()
	s := kithttp.NewServer(func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return endpoint(ctx, request)
	}, func(i context.Context, req *http.Request) (request interface{}, err error) {
		return nil, nil
	}, func(i context.Context, writer http.ResponseWriter, v interface{}) error {

		if r, ok := v.(string); ok {
			io.WriteString(writer, r)
		}
		return nil
	})

	http.ListenAndServe(":7777", s)
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

	instancer := consul.NewInstancer(kitc, logger, "TimeSvr", nil, true)

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
	return retry
}
