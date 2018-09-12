## zipkin in gokit

#### 场景

* client 调用 http api
* http api 调用 rpc A
* rpc A 调用 rpc B
* 返回结果

#### 类型

1. #### 业务无侵入

* ###### Http
>`HTTPClientTrace`和
`HTTPServerTrace`
是一组分别对应client和server的trace方法，帮我们封装了span的相关调用。直接在生成client和server的时候，当做options参入传入即可。

span的name默认为http的method，比如post，get等，可以通过`Name() TracerOption`自定义name

* ###### GRPC
>`GRPCClientTrace`和
`GRPCServerTrace`
是一组分别对应client和server的trace方法，帮我们封装了span的相关调用。直接在生成client和server的时候，当做options参入传入即可。

span的name默认为grpc的/service/method，比如user/login，可以通过`Name() TracerOption`自定义name

* ###### Endpoint

>`TraceEndpoint`提供了通用的针对endpoint的middleware,主要是针对各种endpoint的执行情况进行追踪，比如自定义了一个endpoint获取分布式锁，打算追踪这个endpoint的情况，又不想修改原来的代码，可以通过`TraceEndpoint`来实现。
代码如下:

```go

myEndpoint := MakeMyEndpoint(svc)
myEndpoint = zipkin.TraceEndpoint(zipkinTracer, "log")(myEndpoint)
myEndpoint = Log(myEndpoint)
myEndpoint = zipkin.TraceEndpoint(zipkinTracer, "getlock")(myEndpoint)
myEndpoint = GetLock(myEndpoint)

````
注意:可以只有http或者grpc的trace，没有endpoint的trace，也就是说endpoint不是必须的。

2. #### trace业务

>如果使用了gokit的http或者grpc的trace功能，具体业务的endpoint的ctx中已经包含了span的相关信息，直接使用就可以了。
>比如trace sql query的执行时间,cache的load时间等等。

```go

span := opzipkin.SpanFromContext(ctx)
if span != nil {
    span.Annotate(time.Now(), "Start biz")
    defer func() {
        span.Annotate(time.Now(), "End biz")
    }()
}

if span != nil {
	span.Tag("k","v")
}

//do biz

if span != nil {
	span.Tag("k","v")
}
```

####注意

>业务接口中一定要带入ctx
否则无法连接父span，也不能做业务的trace了。

#### 测试

>* 分别启动3个服务

```shell
$ go run kit/api/cmd/main.go
$ go run kit/svr/user/cmd/main.go
$ go run kit/svr/msg/cmd/main.go
```

>* curl test

```
$ curl -X POST "http://localhost:8888/login" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"username\": \"admin\",\"pwd\":\"123\"}"
{"code":0,"unread":2,"msg":"登录成功","uid":"1"}
```




