package api

import (
	"net/http"
	"golang.org/x/net/context"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

func DecodeLoginReq(ctx context.Context, r *http.Request) (request interface{}, err error) {

	req := ReqLogin{}
	bytes, e := ioutil.ReadAll(r.Body)
	defer func() {
		r.Body.Close()
	}()
	if e != nil {
		return nil, e
	}
	if e := json.Unmarshal(bytes, &req); e != nil {
		fmt.Println(e)
		return nil, e
	}
	return req, nil
}

func EncodeRes(ctx context.Context, writer http.ResponseWriter, i interface{}) error {

	bytes, e := json.Marshal(i)

	if e != nil {
		return e
	}

	writer.Write(bytes)

	return nil
}
