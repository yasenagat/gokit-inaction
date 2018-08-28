package main

import (
	transporthttp "github.com/go-kit/kit/transport/http"
	"net/http"
	"golang.org/x/net/context"
	"github.com/go-kit/kit/ratelimit"
	"golang.org/x/time/rate"
	"time"
	"math/rand"
	"github.com/pkg/errors"
	rl "github.com/juju/ratelimit"
	"github.com/go-kit/kit/endpoint"
)

func main() {

	helloHandler := transporthttp.NewServer(ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Minute), 5))(func(ctx context.Context, request interface{}) (interface{}, error) {

		if rand.Intn(10) > 5 {
			if r, ok := request.(string); ok {
				return r, nil
			}

			return "default data", nil
		} else {
			return nil, errors.New("service error")
		}

	}), func(i context.Context, request *http.Request) (interface{}, error) {

		return "hell world", nil
	}, func(_ context.Context, writer http.ResponseWriter, i interface{}) error {
		if r, ok := i.(string); ok {
			writer.Write([]byte(r))
		}
		return nil
	})

	b := rl.NewBucket(time.Minute, 5)

	limitHandler := transporthttp.NewServer(NewBucketLimit(b)(func(ctx context.Context, request interface{}) (interface{}, error) {

		if r, ok := request.(string); ok {
			return r, nil
		}

		return "default data", nil

	}), func(i context.Context, request *http.Request) (interface{}, error) {

		return "hell world", nil
	}, func(_ context.Context, writer http.ResponseWriter, i interface{}) error {
		if r, ok := i.(string); ok {
			writer.Write([]byte(r))
		}
		return nil
	})
	http.Handle("/hello", helloHandler)
	http.Handle("/limit", limitHandler)
	http.ListenAndServe(":8080", nil)
}

func NewBucketLimit(b *rl.Bucket) endpoint.Middleware {

	return func(i endpoint.Endpoint) endpoint.Endpoint {

		return func(ctx context.Context, request interface{}) (response interface{}, err error) {

			if b.TakeAvailable(1) == 0 {
				return nil, errors.New("访问速率过快")
			}
			return i(ctx, request)
		}
	}
}
