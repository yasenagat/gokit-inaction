package main

import (
	"fmt"
	"net/http"
	transporthttp "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
	"encoding/json"
	"github.com/pkg/errors"
	"log"
	"net/url"
	"gitee.com/godY/gokit-inaction/circuitbreaker/kit"
	"io/ioutil"
	"bytes"
	"github.com/gorilla/mux"
)

type Number struct {
	N int
}

//curl -X POST "http://localhost:8888/n" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"N\": 4}"

//curl -X POST "http://localhost:8888/mock" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"N\": 4}"
func main() {

	errc := make(chan error)

	r := mux.NewRouter()
	svr := transporthttp.NewServer(kit.NewHystrixEndpoint("/n", NewRemoteEndPoint), func(_ context.Context, req *http.Request) (request interface{}, err error) {
		n := Number{}
		e := json.NewDecoder(req.Body).Decode(&n)
		return n, e
	}, func(_ context.Context, writer http.ResponseWriter, i interface{}) error {
		return json.NewEncoder(writer).Encode(i)
	})

	mocksvr := transporthttp.NewServer(kit.NewCbEndpoint("/mock", NewMockEndPoint), func(_ context.Context, req *http.Request) (request interface{}, err error) {
		n := Number{}
		e := json.NewDecoder(req.Body).Decode(&n)
		return n, e
	}, func(_ context.Context, writer http.ResponseWriter, i interface{}) error {
		return json.NewEncoder(writer).Encode(i)
	})

	r.Handle("/n", svr)
	r.Handle("/mock", mocksvr)
	go func() {
		log.Println("[Api] Service Start On :8888")
		errc <- http.ListenAndServe(":8888", r)
	}()

	fmt.Println(<-errc)
}

func NewRemoteEndPoint(ctx context.Context, request interface{}) (response interface{}, err error) {
	log.Println(request.(Number).N)
	return bizN(ctx, request)
}

func NewMockEndPoint(ctx context.Context, request interface{}) (response interface{}, err error) {

	log.Println("mock", request.(Number).N)

	if r, ok := request.(Number); ok {
		if r.N <= 20 {
			r.N = r.N * 2
			return r, nil
		}
	}
	return nil, errors.New("[Api] Server Error Mock")
}

func bizN(ctx context.Context, request interface{}) (Number, error) {

	//模拟Api Server Error
	if request.(Number).N < 10 {
		return Number{}, errors.New("[Api] Server Error N")
	}

	return RemoteCall(ctx, request)

}

func RemoteCall(ctx context.Context, request interface{}) (Number, error) {

	parseUrl, _ := url.Parse("http://localhost:7777/n2")

	client := transporthttp.NewClient("POST", parseUrl, func(ctx context.Context, request *http.Request, i interface{}) error {
		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(i); err != nil {
			return err
		}
		request.Body = ioutil.NopCloser(&buf)
		return nil
	}, func(ctx context.Context, res *http.Response) (response interface{}, err error) {

		b, e := ioutil.ReadAll(res.Body)

		log.Println("b", string(b))
		log.Println("StatusCode", res.StatusCode)

		if res.StatusCode != http.StatusOK {

			return nil, errors.New(string(b) + " => " + res.Status)
		}

		n := Number{}
		e = json.Unmarshal(b, &n)
		return n, e

	})

	ret, e := client.Endpoint()(ctx, request)

	if e != nil {
		return Number{}, e
	}

	return ret.(Number), nil
}
