package main

import "github.com/openzipkin/zipkin-go-opentracing"

func main() {
	zipkintracer.NewHTTPCollector("")
}
