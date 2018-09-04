## gokit中实现circuitbreaker

#### 使用场景
>本地Server的Endpoint的MiddleWare，Endpoint中有远程调用方法，比如调用远程RPC接口或http接口，并且远程调用很有可能会失败。

#### 远程调用失败
>http response StateCode
>RPC消息体中的err或者RPC调用返回error

#### 本地的Endpoint
>远程调用失败，直接在本地的Endpoint返回error或包装的error


