package kit

import (
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/sony/gobreaker"
	"github.com/go-kit/kit/endpoint"
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
