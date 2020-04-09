package main

import (
	transhttp "github.com/go-kit/kit/transport/http"
	"net/http"
	"github.com/go-kit/kit/endpoint"
	"context"
	"github.com/pkg/errors"
	"io/ioutil"
)

//curl -X POST "http://localhost:9777" -H "accept: application/json" -H "Content-Type: application/json" -d "张三"
//curl -X POST "http://localhost:9777" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"A\": \"admin\",\"B\":\"123\"}"
func main() {

	server := transhttp.NewServer(MakeHelloWorldEndpoint(), MakeHelloDeReqFunc(), MakeHelloEnResFunc(), )

	http.Handle("/", server)
	http.ListenAndServe(":9777", nil)
}

func MakeHelloWorldEndpoint() endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		if v, ok := request.(string); ok {
			return v + " => hello world", nil
		}
		return nil, errors.New("HelloWorld Error")
	}
}

func MakeHelloDeReqFunc() transhttp.DecodeRequestFunc {

	return func(i context.Context, request *http.Request) (req interface{}, err error) {
		b, err := ioutil.ReadAll(request.Body)
		if err != nil {
			return nil, err
		}
		return string(b), nil
	}
}

func MakeHelloEnResFunc() transhttp.EncodeResponseFunc {
	return func(ctx context.Context, writer http.ResponseWriter, data interface{}) error {
		if r, ok := data.(string); ok {
			writer.Write([]byte(r))
		}
		return nil
	}
}
