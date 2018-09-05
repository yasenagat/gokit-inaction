## gokit中实现circuitbreaker

#### 模拟场景
>Client->Api Server->Remote Server(输入一个数N，返回N*2)

#### 熔断位置
>Api Server对外提供N*2的接口服务`/n`,在`/n`的服务中需要调用远程接口`/n2`,所以要在Api Server的`/n`接口中增加熔断器。

#### 熔断判断
>是否启用熔断机制的判断依据是`/n`接口是否正常执行，并不直接对`/n`所依赖的远程接口`/n2`进行判断。但是`/n`是否能正常执行，依赖远程接口的响应，所以要对远程接口是否正常返回进行判断，如果`/n`中依赖多个远程接口，那么对远程接口是否正常响应的判断依据，可能会不同，建立统一的响应规则就十分重要。

注意：
熔断是针对`/n`接口，无论`/n`自己的内部出错，还是调用远程接口出错，都会启动熔断机制的判断。

#### 代码实现

##### 模拟场景(内部、远程错误)

* /n 自己内部执行异常，会触发熔断，调用远程接口异常，也会出发熔断，异常数量会累加，目前不能区分。

* /mock 模拟了只有内部逻辑，无远程调用的情况。

* /n和/mock的熔断器各自独立，互不影响。

##### Server

* Api Server :8888
    * /n num<10 会报错 `[Api] Server Error`
    * /mock num>20 会报错
* Remote Server(N*2) :7777
    * /n2 num<20 错误请求 400 `[N2] Server Req Error[ 不能小于20！]`
    * /n2 num*2>100 模拟内部调用报错或其他远程错误 500 `[N2] Server Biz Error`
    * `错误定义` 如果出现错误，返回对应的错误码,Response Body 直接返回一个错误描述(也可定义成json)。
    
* 熔断器
    * 超时时间 10秒。
    * 熔断阈值 总错误数>5或者连续错误>2。
    * 监控 状态变为断开发送警告日志。


#### 测试
>通过下面的命令，更改N的值，测试结果。
注意观察Api Server的`状态变化和响应内容`
```
$ curl -X POST "http://localhost:8888/n" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"N\": 4}"

$ curl -X POST "http://localhost:8888/mock" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"N\": 4}"

$ curl -X POST "http://localhost:7777/n2" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"N\": 4}"
```



