package main

import (
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/kit/auth/basic"
	"net/http"
	"golang.org/x/net/context"
	"io"
	"time"
	"encoding/base64"
	"fmt"
	"strings"
	"bytes"
)

//Basic eWFzZToxMjM0NTY=

//curl -X POST "http://localhost:9977" -H "Authorization: Basic eWFzZToxMjM0NTY="

func main() {

	//auth := base64.StdEncoding.EncodeToString([]byte("yase:123456"))
	//fmt.Println(auth)
	//username, password, ok := parseBasicAuth("Basic " + auth)
	//fmt.Println(string(username))
	//fmt.Println(string(password))
	//fmt.Println(ok)

	h := kithttp.NewServer(basic.AuthMiddleware("yase", "123456", "Auth Error")(func(ctx context.Context, request interface{}) (response interface{}, err error) {

		return nil, nil
	}), func(context context.Context, req *http.Request) (request interface{}, err error) {
		fmt.Println(context.Value(kithttp.ContextKeyRequestAuthorization))
		//u, s, b := req.BasicAuth()
		//fmt.Println(u, s, b)
		return nil, nil
	}, func(context context.Context, writer http.ResponseWriter, i interface{}) error {

		io.WriteString(writer, time.Now().Format(time.RFC3339))

		return nil
	}, kithttp.ServerBefore(kithttp.PopulateRequestContext))

	http.ListenAndServe(":9977", h)
}

func parseBasicAuth(auth string) (username, password []byte, ok bool) {
	const prefix = "Basic "
	if !strings.HasPrefix(auth, prefix) {
		return
	}
	c, err := base64.StdEncoding.DecodeString(auth[len(prefix):])
	if err != nil {
		return
	}

	s := bytes.IndexByte(c, ':')
	if s < 0 {
		return
	}
	return c[:s], c[s+1:], true
}
