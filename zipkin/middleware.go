package zipkin

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
)

func newLogMiddleware(next endpoint.Endpoint, logger log.Logger) endpoint.Endpoint {
	return func(i endpoint.Endpoint) endpoint.Endpoint {
		logger.Log()
		return i
	}(next)
}
