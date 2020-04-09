## gokit 微服务

### HelloWorld
> 简单的http server
* [概述](https://gitee.com/godY/gokit-inaction/tree/master/helloword)

### http 标准服务

> 通过context传递流水号(msgid)，用户信息(user)

* [概述](https://gitee.com/godY/gokit-inaction/tree/master/hello)

### 服务发现

* [概述](https://gitee.com/godY/gokit-inaction/tree/master/consul)

* [架构](https://gitee.com/godY/gokit-inaction/tree/master/consul)

* [部署 (2 client，3 server)](https://gitee.com/godY/gokit-inaction/tree/master/consul)

* [标准使用](https://gitee.com/godY/gokit-inaction/tree/master/consul/std/user)
    
    * http网关
    * http微服务
    
* [kit http服务发现](https://gitee.com/godY/gokit-inaction/tree/master/consul/kit/http)
    
    * http网关
    * http微服务
    
* [kit grpc服务发现](https://gitee.com/godY/gokit-inaction/tree/master/consul/kit/grpc)
    
    * http网关
    * grpc微服务
    
### 服务认证

* [概述](https://gitee.com/godY/gokit-inaction/tree/master/auth)

  `http标准认证`

### 服务熔断

* [概述](https://gitee.com/godY/gokit-inaction/tree/master/circuitbreaker)
    * 使用场景
    * 三种状态：闭合，半开，断开
    * 失败判断
    * 状态转移

* [标准使用](https://gitee.com/godY/gokit-inaction/tree/master/circuitbreaker/std)
    * gobreaker
    * handybreker
    * hystrix

* [kit集成](https://gitee.com/godY/gokit-inaction/tree/master/circuitbreaker/kit)
    * Client->Api Server->Remote Server(输入一个数N，返回N*2)

### 服务限流
* [概述](https://gitee.com/godY/gokit-inaction/tree/master/limit)

* [kit集成](https://gitee.com/godY/gokit-inaction/tree/master/limit)

### 链路追踪

* [概述](https://gitee.com/godY/gokit-inaction/tree/master/zipkin)

    * 使用场景
    * 关键点
    * zipkin使用
    
 * [标准使用](https://gitee.com/godY/gokit-inaction/tree/master/zipkin/std)
    * span
    * zipkin
 
* [kit集成](https://gitee.com/godY/gokit-inaction/tree/master/zipkin/std)
    * http网关
    * grpc微服务
    * 无业务侵入模式
    * trace业务
   
### GRPC
* [概述](https://gitee.com/godY/gokit-inaction/tree/master/grpc)

* [标准使用](https://gitee.com/godY/gokit-inaction/tree/master/grpc/std)

* [kit集成](https://gitee.com/godY/gokit-inaction/tree/master/grpc/kit)

### 系统监控

* [标准使用](https://gitee.com/godY/gokit-inaction/tree/master/prometheus)
    * prometheus
    
### 集成demo

[ALL IN ONE](https://gitee.com/godY/gokit-inaction/tree/master/aio)

* 上面所有组件都在一个demo里
* sd,circuitbreaker,ratelimit,trace,metrics
* 模拟用户登录
* api网关
* grpc用户服务
* grpc账户服务


### 用户服务demo
[模拟用户服务](https://gitee.com/godY/gokit-inaction/tree/master/user)
* http api server
* http post body
* request and response 都是json字符串
* /login 用户登录
* /phone 修改手机号
* /user 获取用户信息

