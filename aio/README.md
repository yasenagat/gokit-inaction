## all in one demo

>sd,circuitbreaker,ratelimit,trace,metrics

几乎所有组件全都在一个demo里面

#### 

#### 模拟用户登录场景

1. client -> 用户登录
2. api <- 收到请求
3. api -> 校验用户名密码
4. usersvr <- 收到请求
5. usersvr -> `校验`-> 发送消息(校验结果和用户信息)
6. api -> 收到消息 
7. api -> 判断是否成功
8. 
    * 失败 ; api -> 返回
        * . client <- 收到消息
    * 成功 ; api -> 查询用户账户信息
        * . account <- 收到请求
        * . account -> `查询` -> 发送消息(账户信息)
        * . api <- 收到消息
        * . api -> 组合`用户信息`和`账户信息` -> 发送消息
        * . client <- 收到消息
    
