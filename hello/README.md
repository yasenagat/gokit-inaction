## 最简单的http server

#### Hello World 实现

#### ServerBefore方法
>ServerBefore方法，在http request decode方法前执行。

#### Request中的信息
>因为gokit的设计原因，Endpoint的接口的输入参数中，没有标准的http request，但是有些时候，我们需要在业务处理的Endpoint中后区http request的相关信息，比如ip地址，method等
gokit提供了PopulateRequestContext，把http request中的常用基本属性，通过context传递到Endpoint中。
PopulateRequestContext就是在http request decode前执行，这样就能保证Decode,Endpoint,Encode中都可以获取到http request的相关信息

#### 内部流水号
>一个请求从接受开始，到处理结束，中间无论经过多少个流程和服务，都应该可以通过一个唯一且不重复的值(msgId)串联起来。
可以用来查询日志，监控数据等等。
实现方式和PopulateRequestContext类似


```shell

$ curl -X POST "http://localhost:8080"
Jack T%                                                                                        

$ go run main.go
req msgid 634f25cb-a0ea-45ce-99f3-209e4336f8dc
exe msgid 634f25cb-a0ea-45ce-99f3-209e4336f8dc
exe user id 1726ed22-8a89-474b-bf20-7d687feba5e6
Method POST
RequestPath /
RequestURI /
X-Request-ID 
res msgid 634f25cb-a0ea-45ce-99f3-209e4336f8dc
res user id 1726ed22-8a89-474b-bf20-7d687feba5e6


```
