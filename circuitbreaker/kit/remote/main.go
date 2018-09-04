package main

import (
	"fmt"
	"net/http"
	transporthttp "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
	"encoding/json"
	"log"
)

type Number struct {
	N int
}

//curl -X POST "http://localhost:6666" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"N\": 4}"
func main() {

	errc := make(chan error)

	svr := transporthttp.NewServer(func(_ context.Context, request interface{}) (response interface{}, err error) {

		if r, ok := request.(Number); ok {
			if r.N < 20 {
				return Number{N: r.N * 2}, nil
			}
		}

		return Number{N: -1}, nil
	}, func(_ context.Context, req *http.Request) (request interface{}, err error) {
		n := Number{}
		e := json.NewDecoder(req.Body).Decode(&n)
		return n, e
	}, func(_ context.Context, w http.ResponseWriter, i interface{}) error {

		fmt.Println("i", i)
		if n, ok := i.(Number); ok {
			if n.N == -1 {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(i)
				return nil
			}
		}

		return json.NewEncoder(w).Encode(i)
	})

	go func() {
		log.Println("[N*2] Service Start On :6666")
		errc <- http.ListenAndServe(":6666", svr)
	}()

	fmt.Println(<-errc)
}
