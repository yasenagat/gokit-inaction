package main

import (
	"fmt"
	"net/http"
	transporthttp "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
	"encoding/json"
	"log"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"io"
)

type Number struct {
	N int
}

//curl -X POST "http://localhost:7777/n2" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"N\": 4}"
func main() {

	errc := make(chan error)

	r := mux.NewRouter()

	svr := transporthttp.NewServer(func(_ context.Context, request interface{}) (response interface{}, err error) {

		if r, ok := request.(Number); ok {
			if r.N >= 20 {
				i, e := bizN2(r.N)
				if e != nil {
					return nil, e
				}
				return Number{N: i}, nil
			} else {
				return Number{N: -1}, nil
			}
		}

		return nil, errors.New("Type Error")
	}, func(_ context.Context, req *http.Request) (request interface{}, err error) {
		n := Number{}
		e := json.NewDecoder(req.Body).Decode(&n)
		//if e != nil {
		//	return nil, e
		//}
		return n, e
	}, func(_ context.Context, w http.ResponseWriter, i interface{}) error {

		fmt.Println("i", i)
		if n, ok := i.(Number); ok {
			if n.N == -1 {
				w.WriteHeader(http.StatusBadRequest)
				io.WriteString(w, "[N2] Server Req Error[ 不能小于20！]")
				return nil
			}
		}

		return json.NewEncoder(w).Encode(i)
	})

	r.Handle("/n2", svr)
	go func() {
		log.Println("[N*2] Service Start On :7777")
		errc <- http.ListenAndServe(":7777", r)
	}()

	fmt.Println(<-errc)
}

func bizN2(num int) (int, error) {

	res := num * 2

	if res <= 100 {
		return res, nil
	}

	return -1, errors.New("[N2] Server Biz Error")
}
