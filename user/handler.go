package user

import (
	"encoding/json"
	"github.com/go-kit/kit/log"
	transporthttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/net/context"
	"net/http"
)

//type UserHandler struct {
//	LoginHandler       *transporthttp.Server
//	UpdatePhoneHandler *transporthttp.Server
//	GetUserHandler     *transporthttp.Server
//}

func MakeHandler(logger log.Logger, userEndpoint *UserEndpoint, r *mux.Router) *mux.Router {

	opts := []transporthttp.ServerOption{
		transporthttp.ServerErrorLogger(logger),
		transporthttp.ServerErrorEncoder(encodeError),
		transporthttp.ServerBefore(transporthttp.PopulateRequestContext, InnerMsgRequestContext),
	}
	der := DefaultEncodeResponse(logger)
	loginHandler := transporthttp.NewServer(userEndpoint.Login, DecodeLoginReq(logger), der, opts...)
	updatePhoneHandler := transporthttp.NewServer(userEndpoint.UpdatePhone, DecodeUpdatePhoneReq(logger), der, opts...)
	getUserHandler := transporthttp.NewServer(userEndpoint.GetUser, DecodeGetUserReq(logger), der, opts...)
	r.Handle("/login", loginHandler)
	r.Handle("/phone", updatePhoneHandler)
	r.Handle("/user", getUserHandler)
	r.Handle("/metrics", promhttp.Handler())
	return r
}

// encode errors from business-logic
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	//case cargo.ErrUnknown:
	//	w.WriteHeader(http.StatusNotFound)
	//case ErrInvalidArgument:
	//	w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
