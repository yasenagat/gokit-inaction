package api

import (
	"encoding/json"
	"golang.org/x/net/context"
	"net/http"
)

func DecodeLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request ReqLogin
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
