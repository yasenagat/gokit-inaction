## http service 注册 和 发现

### 注意

> 所有的注册和发现都应该操作`本地的consul client`
,由于没启动本地的consul client,所以代码中直接访问了remote consul server。

#### 实例

* api 模拟api gateway，对外提供http接口服务。
* svr 模拟一个微服务，此微服务中只有一个接口，通常有N个接口
* cli 模拟web或者手机客户端，通过api gateway调用接口
* cli->api->svr
    * api 查询可用服务
    * svr 注册服务

#### 部署

* 启动3个consul agent server，互相join，成为cluster。
* 启动3个svr，3个consul agent client。每个svr对应自己本地的consul client。
* 启动一个api。
* 运行多个cli，模拟测试多请求。

>一共需要部署7个node

* sd 服务发现

    * Node A 部署 consul server
    * Node B 部署 consul server
    * Node C 部署 consul server
* svr 微服务
    * Node D 部署 consul client,svr
    * Node E 部署 consul client,svr
    * Node F 部署 consul client,svr
* api 网关
    * Node G 部署 api gateway
    
>cli可以运行在本地


#### 测试

* 启动多个cli,不停的发起请求，模拟用户请求。
* 请求会发送到D,E,F(目前是轮询策略)
* 关闭D的svr。
* cli仍然可以正常响应,请求发到到E,F。
* 恢复D的svr。
* 新请求会发送到D,E,F。
* 新增一个Node N，启动svr和consul client，注册服务。
* 后续请求会发送到D,E,F,N




----------

下面是脚本，可以忽略

```shell

nohup ./consul agent -data-dir consuldata -client 0.0.0.0 -ui -join 192.168.10.208  -bind 192.168.10.37 -node sddev-4 > consul.out &

nohup ./consul agent -data-dir=/tmp/consul-new -server -client 0.0.0.0 -ui -retry-join "192.168.10.208" -bind 192.168.3.125 -node sddev-1 > consul.out &

nohup ./consul agent -data-dir=/tmp/consul-new -server -bootstrap -client 0.0.0.0 -ui -retry-join 192.168.3.125 -bind 192.168.10.208 -node sddev-2 > consul.out &

nohup ./consul agent -data-dir=/tmp/consul-new -server -client 0.0.0.0 -ui -join 192.168.10.208  -bind 192.168.10.210 -node sddev-3 > consul.out &

[root@localhost ~]# ./consul members
Node     Address              Status  Type    Build  Protocol  DC   Segment
sddev-1  192.168.3.125:8301   alive   server  1.2.2  2         dc1  <all>
sddev-2  192.168.10.208:8301  failed  server  1.2.2  2         dc1  <all>
sddev-3  192.168.10.210:8301  alive   server  1.2.2  2         dc1  <all>


[root@localhost ~]# ./consul operator raft list-peers
Node     ID                                    Address              State     Voter  RaftProtocol
sddev-2  561ee619-fc31-eff4-021d-543283241e8d  192.168.10.208:8300  leader    true   3
sddev-1  d72dde6d-c632-8016-fcfa-977774ff416a  192.168.3.125:8300   follower  true   3
sddev-3  c314bbd2-ed53-3a23-4119-9d8d503e5460  192.168.10.210:8300  follower  true   3
```