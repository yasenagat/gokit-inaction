package main

import (
	transporthttp "github.com/go-kit/kit/transport/http"
	"net/http"
	"golang.org/x/net/context"
	"fmt"
	"github.com/google/uuid"
)

type User struct {
	Id string
}

//curl -X POST "http://localhost:8080"
func main() {

	server := transporthttp.NewServer(func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("exe msgid", ctx.Value("msgid"))
		fmt.Println("exe user id", ctx.Value("user").(User).Id)
		fmt.Println("Method", ctx.Value(transporthttp.ContextKeyRequestMethod).(string))
		fmt.Println("RequestPath", ctx.Value(transporthttp.ContextKeyRequestPath).(string))
		fmt.Println("RequestURI", ctx.Value(transporthttp.ContextKeyRequestURI).(string))
		fmt.Println("X-Request-ID", ctx.Value(transporthttp.ContextKeyRequestXRequestID).(string))
		if r, ok := request.(string); ok {
			return r, nil
		}
		return "hell world", nil
	}, func(ctx context.Context, request *http.Request) (interface{}, error) {
		fmt.Println("req msgid", ctx.Value("msgid"))
		return "Jack T", nil
	}, func(ctx context.Context, writer http.ResponseWriter, i interface{}) error {
		fmt.Println("res msgid", ctx.Value("msgid"))
		fmt.Println("res user id", ctx.Value("user").(User).Id)
		if r, ok := i.(string); ok {
			writer.Write([]byte(r))
		}
		return nil
	}, transporthttp.ServerBefore(transporthttp.PopulateRequestContext, CreateMsgId))
	http.Handle("/", server)
	http.ListenAndServe(":8080", nil)
}

func CreateMsgId(i context.Context, request *http.Request) context.Context {
	i = context.WithValue(i, "msgid", uuid.New().String())
	i = context.WithValue(i, "user", User{Id: uuid.New().String()})
	return i
}
