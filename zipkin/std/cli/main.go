package main

import (
	"fmt"
	"github.com/opentracing/opentracing-go"

	zipkin "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"log"
	"github.com/pkg/errors"
	"math/rand"
	"os"
	"github.com/openzipkin/zipkin-go-opentracing"
	"time"
)

const KEYERROR = "error"

func main() {

	//a := serviceA{"127.0.0.0:1111", "svrA"}
	//
	//for i := 0; i < 10; i++ {
	//	a.cal()
	//}

	x := serviceX{"192.168.0.0:1111", "svrX"}

	for i := 0; i < 10; i++ {
		x.call(rand.Intn(10))

	}
}

type serviceX struct {
	hostPost string
	name     string
}

func (s serviceX) call(num int) (int, error) {

	collector, tracer, e := s.createTracer()

	defer collector.Close()

	if e != nil {
		log.Print(e)
		return -1, e
	}

	span := tracer.StartSpan("call")
	defer span.Finish()

	span.SetTag("request num", num)

	y := serviceY{"192.168.0.0:2222", "svrY", span}

	ret, e := y.get(num)

	if e != nil {
		span.SetTag(KEYERROR, e.Error())
		log.Println(e)
	}

	return ret, e
}

func (s serviceY) get(num int) (int, error) {

	c, tracer, e := s.createTracer()
	defer c.Close()
	if e != nil {
		return -1, errors.Wrap(e, "CreateTracer Error")
	}

	span := tracer.StartSpan("get", opentracing.ChildOf(s.parentSpan.Context()))
	defer span.Finish()

	//TODO 执行时间
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	n := rand.Intn(10)
	span.LogKV("GetNumFromDb", n)
	span.LogKV("RecevieNum", num)
	span.SetTag("GetNumFromDb", n)
	span.SetTag("RecevieNum", num)

	if n < 5 {
		span.SetTag("error", s.hostPost+" "+s.name+" get error")
		return -1, errors.New(s.hostPost + " " + s.name + " get error")
	}

	sum, e := sumD(span, n, num)

	log.Println(num, n, sum)

	if e != nil {
		span.SetTag(KEYERROR, e)
		return -1, e
	}
	return sum, nil
}

func (s serviceZ) sum(a, b int) (int, error) {

	c, tracer, e := s.createTracer()
	defer c.Close()
	if e != nil {
		return -1, errors.Wrap(e, "CreateTracer Error")
	}

	span := tracer.StartSpan("sum", opentracing.ChildOf(s.parentSpan.Context()))
	defer span.Finish()
	span.SetTag("a", a)
	span.SetTag("b", b)

	//TODO 执行时间
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	num := a + b

	span.SetTag("sum result", num)
	span.LogKV("sum finish", num)
	if num > 12 {
		span.SetTag("error", "sum error")
		return -1, errors.New("sum error")
	}

	return num, nil
}

type serviceY struct {
	hostPost   string
	name       string
	parentSpan opentracing.Span
}

type serviceZ struct {
	hostPost   string
	name       string
	parentSpan opentracing.Span
}

type serviceA struct {
	hostPost string
	name     string
}

type serviceB struct {
	hostPost   string
	name       string
	parentSpan opentracing.Span
}

type serviceC struct {
	hostPost   string
	name       string
	parentSpan opentracing.Span
}

type serviceD struct {
	hostPost   string
	name       string
	parentSpan opentracing.Span
}

func getFromB(span opentracing.Span) (int, error) {
	return serviceB{"127.0.0.0:2222", "svrB", span}.get()
}

func getFromC(span opentracing.Span) (int, error) {
	return serviceC{"127.0.0.0:3333", "svrC", span}.get()
}

func sumD(span opentracing.Span, a, b int) (int, error) {
	return serviceD{"127.0.0.0:4444", "svrD", span}.sum(a, b)
}

func (s serviceA) cal() (int, error) {

	c, tracer, e := s.createTracer()
	defer c.Close()
	if e != nil {
		return -1, errors.Wrap(e, "CreateTracer Error")
	}

	rootSpan := tracer.StartSpan("call")
	defer rootSpan.Finish()

	//Call B
	n1, e := getFromB(rootSpan)

	if e != nil {
		rootSpan.SetTag("error", e.Error())
		log.Println(e)
		return -1, e
	}

	rootSpan.SetTag("n1", n1)
	rootSpan.LogKV("Get-Value-From-svrB", n1)

	//Call C
	n2, e := getFromC(rootSpan)
	if e != nil {
		rootSpan.SetTag("error", e.Error())
		log.Println(e)
		return -1, e
	}

	rootSpan.SetTag("n2", n2)
	rootSpan.LogKV("Get-Value-From-svrC", n2)

	result, e := sumD(rootSpan, n1, n2)
	log.Println(n1, n2, result)
	if e != nil {
		rootSpan.SetTag("error", e.Error())
		log.Println(e)
		return -1, e
	}

	rootSpan.LogKV("Get-Sum-Value-From-svrD", result)
	rootSpan.SetTag("result", result)

	return result, nil
}

func (s serviceB) get() (int, error) {

	c, tracer, e := s.createTracer()
	defer c.Close()
	if e != nil {
		return -1, errors.Wrap(e, "CreateTracer Error")
	}

	span := tracer.StartSpan("get", opentracing.ChildOf(s.parentSpan.Context()))
	defer span.Finish()

	//TODO 执行时间
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	num := rand.Intn(10)
	span.LogKV("GetNumFromDb", num)
	span.SetTag("num", num)
	if num < 5 {
		span.SetTag("error", s.hostPost+" "+s.name+" get error")
		return -1, errors.New(s.hostPost + " " + s.name + " get error")
	}

	return num, nil
}
func (s serviceC) get() (int, error) {

	c, tracer, e := s.createTracer()
	defer c.Close()
	if e != nil {
		return -1, errors.Wrap(e, "CreateTracer Error")
	}

	span := tracer.StartSpan("get", opentracing.ChildOf(s.parentSpan.Context()))
	defer span.Finish()

	//TODO 执行时间
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	num := rand.Intn(7)
	span.LogKV("GetNumFromDb", num)
	span.SetTag("num", num)
	if num > 5 {
		span.SetTag("error", s.hostPost+" "+s.name+" get error")
		return -1, errors.New(s.hostPost + " " + s.name + " get error")
	}

	return num, nil
}

func (s serviceD) sum(a, b int) (int, error) {
	c, tracer, e := s.createTracer()
	defer c.Close()
	if e != nil {
		return -1, errors.Wrap(e, "CreateTracer Error")
	}

	span := tracer.StartSpan("sum", opentracing.ChildOf(s.parentSpan.Context()))
	defer span.Finish()
	span.SetTag("a", a)
	span.SetTag("b", b)

	//TODO 执行时间
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	num := a + b

	span.SetTag("sum result", num)
	span.LogKV("sum-finish", num)
	if num > 12 {
		span.SetTag("error", "sum error")
		return -1, errors.New("sum error")
	}
	return num, nil
}

func (s serviceA) createTracer() (zipkintracer.Collector, opentracing.Tracer, error) {
	return newTracer(s.hostPost, s.name)
}

func (s serviceB) createTracer() (zipkintracer.Collector, opentracing.Tracer, error) {
	return newTracer(s.hostPost, s.name)
}

func (s serviceC) createTracer() (zipkintracer.Collector, opentracing.Tracer, error) {
	return newTracer(s.hostPost, s.name)
}

func (s serviceD) createTracer() (zipkintracer.Collector, opentracing.Tracer, error) {
	return newTracer(s.hostPost, s.name)
}

func (s serviceX) createTracer() (zipkintracer.Collector, opentracing.Tracer, error) {
	return newTracer(s.hostPost, s.name)
}

func (s serviceY) createTracer() (zipkintracer.Collector, opentracing.Tracer, error) {
	return newTracer(s.hostPost, s.name)
}

func (s serviceZ) createTracer() (zipkintracer.Collector, opentracing.Tracer, error) {
	return newTracer(s.hostPost, s.name)
}

func newTracer(hostPort, serviceName string) (zipkintracer.Collector, opentracing.Tracer, error) {
	// 收集器
	collector, err := zipkin.NewHTTPCollector("http://192.168.3.125:9411/api/v1/spans")
	if err != nil {
		return nil, nil, err
	}
	//log.Println(collector)
	// 记录器
	recorder := zipkin.NewRecorder(collector, true, hostPort, serviceName)
	//log.Println(recorder)
	// 追踪器
	tracer, err := zipkin.NewTracer(
		recorder,
		zipkin.ClientServerSameSpan(true),
		zipkin.TraceID128Bit(true),
	)
	if err != nil {
		return nil, nil, err
	}
	//全局追踪器，一般一个服务用一个就可以，使用全局追踪器可以使用很多包装过的方法，创建span
	//TODO 因为这里模拟多个服务，所以关闭。
	//opentracing.InitGlobalTracer(tracer)
	return collector, tracer, nil
}

func example() {
	// Create our HTTP collector.
	collector, err := zipkin.NewHTTPCollector("http://192.168.3.125:9411/api/v1/spans")
	if err != nil {
		fmt.Printf("unable to create Zipkin HTTP collector: %+v\n", err)
		os.Exit(-1)
	}
	// Create our recorder.
	recorder := zipkin.NewRecorder(collector, false, "1.0.0.0:1234", "a")

	// Create our tracer.
	tracer, err := zipkin.NewTracer(
		recorder,
		zipkin.ClientServerSameSpan(true),
		zipkin.TraceID128Bit(true),
	)

	if err != nil {
		fmt.Printf("unable to create Zipkin tracer: %+v\n", err)
		os.Exit(-1)
	}

	//opentracing.InitGlobalTracer(tracer)

	span := tracer.StartSpan("s")
	span.LogKV("k", "start s")

	a := tracer.StartSpan("A", opentracing.ChildOf(span.Context()))
	a.LogKV("Name", "A")
	a.SetTag("k", "v")

	b := tracer.StartSpan("B", opentracing.ChildOf(a.Context()))
	b.LogKV("Name", "A")

	b.Finish()
	a.Finish()
	span.Finish()
	collector.Close()
}
