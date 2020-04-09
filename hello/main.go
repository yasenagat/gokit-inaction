package main

import (
	"fmt"
	transporthttp "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	"golang.org/x/net/context"
	"net/http"
)

//curl -X POST "http://localhost:8081/concat" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"A\": \"admin\",\"B\":\"123\"}"
type User struct {
	Id string
}

//var zipkinV2URL = "http://192.168.3.125:9411/api/v2/spans"
//curl -X POST "http://localhost:8080"
func main() {
	//
	//var zipkinTracer *zipkin.Tracer
	//{
	//	var (
	//		err           error
	//		hostPort      = "localhost:80"
	//		serviceName   = "hello"
	//		useNoopTracer = (zipkinV2URL == "")
	//		reporter      = zipkinhttp.NewReporter(zipkinV2URL)
	//	)
	//	defer reporter.Close()
	//	zEP, err := zipkin.NewEndpoint(serviceName, hostPort)
	//	if err != nil {
	//		logger.Log("err1", err)
	//		os.Exit(1)
	//	}
	//	zipkinTracer, err = zipkin.NewTracer(
	//		reporter, zipkin.WithLocalEndpoint(zEP), zipkin.WithNoopTracer(useNoopTracer),
	//	)
	//	if err != nil {
	//		logger.Log("err2", err)
	//		os.Exit(1)
	//	}
	//	if !useNoopTracer {
	//		logger.Log("tracer", "Zipkin", "type", "Native", "URL", zipkinV2URL)
	//	}
	//}
	////zipkinTracer.StartSpan("abc")
	//fmt.Println(zipkinTracer)

	//span.Finish()
	server := transporthttp.NewServer(func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("exe msgid", ctx.Value("msgid"))
		fmt.Println("exe user id", ctx.Value("user").(User).Id)
		fmt.Println("Method", ctx.Value(transporthttp.ContextKeyRequestMethod).(string))
		fmt.Println("RequestPath", ctx.Value(transporthttp.ContextKeyRequestPath).(string))
		fmt.Println("RequestURI", ctx.Value(transporthttp.ContextKeyRequestURI).(string))
		fmt.Println("X-Request-ID", ctx.Value(transporthttp.ContextKeyRequestXRequestID).(string))

		//var (
		//	//spanContext model.SpanContext
		//	serviceName = "MySQL"
		//	serviceHost = "mysql.example.com:3306"
		//	queryLabel  = "GetExamplesByParam"
		//	query       = "select * from example where param = :value"
		//	//parentSpan  zipkin.Span
		//)
		//
		//// retrieve the parent span from context to use as parent if available.
		////if parentSpan := zipkin.SpanFromContext(ctx); parentSpan != nil {
		////	spanContext = parentSpan.Context()
		////}
		//
		//// create the remote Zipkin endpoint
		//ep, _ := zipkin.NewEndpoint(serviceName, serviceHost)
		//
		//// create a new span to record the resource interaction
		//span := zipkinTracer.StartSpan(
		//	queryLabel,
		//	//zipkin.Parent(parentSpan.Context()),
		//	zipkin.RemoteEndpoint(ep),
		//)
		//
		//// add interesting key/value pair to our span
		//span.Tag("query", query)
		//
		//// add interesting timed event to our span
		//span.Annotate(time.Now(), "query:start")
		//
		//// do the actual query...
		//
		//// let's annotate the end...
		//span.Annotate(time.Now(), "query:end")
		//
		//// we're done with this span.
		//span.Finish()

		//return nil, errors.New("Custom Error")
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
	http.ListenAndServe(":8081", nil)
}

func CreateMsgId(i context.Context, request *http.Request) context.Context {
	i = context.WithValue(i, "msgid", uuid.New().String())
	i = context.WithValue(i, "user", User{Id: uuid.New().String()})
	return i
}
