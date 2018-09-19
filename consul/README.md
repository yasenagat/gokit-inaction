## consul

> consul提供的功能非常多，这里仅讨论service discover

> sd 主要提供2个功能，一直注册服务，二是获取服务。

#### 服务发现

>注册服务的两种模式

* 通过配置文件读取服务列表
* 通过http api动态注册服务

注册服务一般都会对服务配置健康检查，这样可以监控服务的状态，最简单的检查方式就是http请求方式，通过向服务的health接口发送一条http请求，检查服务状态。当然还有脚本方式。

>获取服务

* 通过consul的client lib，官方是go语言的，还有其他各种语言的lib
* 通过http api获取

#### 用途

>当把业务拆分成微服务后，某一个微服务可能会部署N份，可能会动态部署或者下线某服务。如果没有sd，就不能做到动态扩容和服务的伸缩。

* 例子

>比如一个order service 订单服务，目前部署3份，用户的手机客户端调用api gateway，api gateway 通过sd获取可用的 order service，调用order service。

>假如用户增多请求量变大，需要增加order service，则可以部署一个新的order service，并且注册到sd，这样api gateway，就可以在下次获取可用的order service的时候，获取到新增的order service(其中一种策略，也可以用sd通知api gateway，或者api gateway订阅sd中order service的新增event等等各种策略，超时缓存？)

>如果某一个order service 挂掉，sd就会通过health check动态感知service的状态，api gateway后续(策略同上)就不会把请求发送到这个order service。


#### 架构

* consul是以`agent`的方式运行
* agent分为两种模式
    * `server` 保存数据持久化
    * `client` 作为client把请求发送到server
* 结合经典的开发模式，简单理解为(实际复杂的多)
    * server 对应我们日常中的db
    * client 对应我们日常开发的程序,比如web server
    * web server 访问 db，consule client 访问 consule -> server
    * 这种比喻主要是为了理解，为什么有server和client两种模式
    * 客户端程序(比如手机)会直接接入我们自己写的web server，不会直接访问db。
    * 同理，对于consul，我们要接入的是`client`而不是`server`。
* server数量，建议奇数个数，且>=3。
* client数量,每个server(这个server可以理解为服务器)一个
    * 机器A，上面有2个service ,user svr,order svr
    * 机器B，上面有1个service ,user svr,account svr
    * A,B上面分别启动consul client，client join到 consul server
    * A上的user svr,order svr，通过A本地的consul client注册。
    * B上的user svr,account svr，通过B本地的consul client注册。


#### 部署

* 2 consul client
    * A 
        * A join D,E,F或者其中任意一个。(如果是任意一个，启动时候要可用，比如join E，E要可用。启动后如果E挂了，并不影响，因为已经通过E join上了D,F)
    * B 同上
* 3 consul server 也就是cluster
    * D join E,F
    * E join D,F
    * F join D,E
    
>一共5台机器，A,B上部署我们自己service，比如登录，注册各种服务都在,A,B上跑，同时A,B上启动consul client提供sd的功能。
>D,E,F以cluster的方式运行 consul server。

