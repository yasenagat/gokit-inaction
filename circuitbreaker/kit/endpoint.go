package kit

import (
	"github.com/afex/hystrix-go/hystrix"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/sony/gobreaker"
	"golang.org/x/net/context"
	"log"
	"time"
)

func NewCbEndpoint(name string, endpoint endpoint.Endpoint) endpoint.Endpoint {
	s := gobreaker.Settings{}
	s.Timeout = time.Second * 10
	s.Name = name
	s.ReadyToTrip = func(counts gobreaker.Counts) bool {

		if counts.TotalFailures > 5 || counts.ConsecutiveFailures > 2 {
			return true
		}
		return false
	}
	s.OnStateChange = func(name string, from gobreaker.State, to gobreaker.State) {

		if to == gobreaker.StateOpen {
			log.Println("WARN=======", name, "=>", to)
		} else {
			log.Println(name, "=>", to)
		}
	}
	return circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(s))(endpoint)
}

func NewHystrixEndpoint(name string, end endpoint.Endpoint) endpoint.Endpoint {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			var resp interface{}
			if err := hystrix.Do(name, func() (err error) {
				resp, err = next(ctx, request)
				return err
			}, func(e error) error {
				log.Println("e", e)
				return e
			}); err != nil {
				return nil, err
			}
			return resp, nil
		}
	}(end)
}
