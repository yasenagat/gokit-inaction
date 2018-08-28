package cmd

import (
	transporthttp "github.com/go-kit/kit/transport/http"
	"net/http"
	"golang.org/x/net/context"
)

func main() {

	server := transporthttp.NewServer(func(ctx context.Context, request interface{}) (interface{}, error) {
		if r, ok := request.(string); ok {
			return r, nil
		}
		return "hell world", nil
	}, func(i context.Context, request *http.Request) (interface{}, error) {

		return "Jack T", nil
	}, func(_ context.Context, writer http.ResponseWriter, i interface{}) error {
		if r, ok := i.(string); ok {
			writer.Write([]byte(r))
		}
		return nil
	})
	http.Handle("/", server)
	http.ListenAndServe(":8080", nil)
}
