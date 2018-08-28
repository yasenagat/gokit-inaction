package main

import (
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/net/context"
	transporthttp "github.com/go-kit/kit/transport/http"
	"time"
	"math/rand"
	"fmt"
)

func main() {

	rc := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{Name: "RequestCount", Help: "hello world request count"}, nil)
	rt := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{Name: "RequestTime", Help: "hello world request time"}, nil)
	server := transporthttp.NewServer(func(ctx context.Context, request interface{}) (interface{}, error) {
		if r, ok := request.(string); ok {
			defer func(start time.Time) {
				rc.Add(1)
				rt.Observe(time.Since(start).Seconds())
			}(time.Now())
			t := rand.Intn(10)
			fmt.Println(t)
			time.Sleep(time.Duration(t) * time.Second)
			return r, nil
		}

		return "hell world", nil
	}, func(i context.Context, request *http.Request) (interface{}, error) {

		return "hell world", nil
	}, func(_ context.Context, writer http.ResponseWriter, i interface{}) error {
		if r, ok := i.(string); ok {
			writer.Write([]byte(r))
		}
		return nil
	})
	http.Handle("/hello", server)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":37000", nil)
}
