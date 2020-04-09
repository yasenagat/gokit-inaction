package main

import (
	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func main() {

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	tgt, e := url.Parse("http://localhost:7788/login")

	if e != nil {
		logger.Log("err", e)
		os.Exit(-1)
	}
	endpoint := kithttp.NewClient("POST", tgt, func(context context.Context, request *http.Request, i interface{}) error {

		request.Body = ioutil.NopCloser(strings.NewReader("{ \"username\": \"abc\", \"Pwd\": \"123\"}"))
		return nil
	}, func(i context.Context, res *http.Response) (response interface{}, err error) {

		defer res.Body.Close()

		byte, e := ioutil.ReadAll(res.Body)

		if e != nil {
			logger.Log("err", e)
		}

		return string(byte), nil
	}).Endpoint()

	for i := 0; i < 10; i++ {
		res, e := endpoint(context.Background(), i)

		if e != nil {
			logger.Log("[err]", e)
		}

		logger.Log("index", strconv.Itoa(i), "res", res)
	}

}
