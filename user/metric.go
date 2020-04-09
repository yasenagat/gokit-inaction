package user

import (
	"github.com/go-kit/kit/metrics"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"time"
)

type metricsMW struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	Service
}

func NewmetricsMW(s Service) Service {
	return &metricsMW{rc, rl, s}
}

const METHOD = "method"

var fieldKeys = []string{METHOD}

var rc = kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
	Namespace: "api",
	Subsystem: "user_service",
	Name:      "request_count",
	Help:      "Number of requests received.",
}, fieldKeys)

var rl = kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
	Namespace: "api",
	Subsystem: "user_service",
	Name:      "request_latency_microseconds",
	Help:      "Total duration of requests in microseconds.",
}, fieldKeys)

func (m metricsMW) Login(username, pwd string) (User, error) {
	defer func(start time.Time) {
		m.requestCount.With(METHOD, "user_Login").Add(1)
		m.requestLatency.With(METHOD, "user_Login").Observe(time.Since(start).Seconds())
	}(time.Now())
	return m.Service.Login(username, pwd)
}
func (m metricsMW) UpdatePhone(username, phone string) error {
	defer func(start time.Time) {
		m.requestCount.With(METHOD, "user_UpdatePhone").Add(1)
		m.requestLatency.With(METHOD, "user_UpdatePhone").Observe(time.Since(start).Seconds())
	}(time.Now())
	return m.Service.UpdatePhone(username, phone)
}
func (m metricsMW) GetUser(username string) (User, error) {
	defer func(start time.Time) {
		m.requestCount.With(METHOD, "user_GetUser").Add(1)
		m.requestLatency.With(METHOD, "user_GetUser").Observe(time.Since(start).Seconds())
	}(time.Now())
	return m.Service.GetUser(username)
}
