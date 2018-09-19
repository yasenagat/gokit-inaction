## consul 服务注册、发现


#### 简单使用

1. server
* 启动yuser服务，向consul注册服务
* 启动服务失败，从consul中删除服务

2. client
* 从consul中查询yuser的可用服务，可能多个。如果无可用，返回空数组。
* 选择第一个(其他策略，比如轮询，权重等)
* 执行add操作
* 执行query操作

>client不需要和server接口地址绑定，仅和sd server绑定。方便server接口的动态增加和删除，解耦server和client。

>client配合sd，可以做lb和retry相关的策略。

>虽然client和server解耦了，但是和sd server紧耦合，并且和service的name紧耦合，比如要从sd中查询订单服务,就必须知道订单服务的名称(也有可能是其他tag)，比如ordersvr。就要求sd server必须是高可用的，容易扩展的。

```shell
2018/09/14 21:44:38 0 yuser localhost 30000
2018/09/14 21:44:38 1 yuser localhost 30000
2018/09/14 21:44:38 http://localhost:30000/users
2018/09/14 21:44:38 success
2018/09/14 21:44:38 [{"username":"yasenagat","id":"8162fa88-6e34-4694-a327-95e29679493d"},{"username":"yasenagat","id":"da8a37c6-b2df-410e-8a8f-94ba7d0b588a"},{"username":"yasenagat","id":"bdbc213c-6d00-4ddc-89c9-72e7a74df0d8"},{"username":"yasenagat","id":"07214a5d-76f2-4bf9-bfc9-02a7c8d2f7d1"},{"username":"yasenagat","id":"8f0a2d52-8995-47a9-b23e-65caf9814d44"},{"username":"yasenagat","id":"9e8c064a-23ab-4ac1-a047-ea85cedb7ea2"}]

```