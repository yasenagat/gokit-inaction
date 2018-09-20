## GPRC Service 注册 和 发现

>模式和http service 注册和发现`几乎`一样。

* svr提供服务的协议改为grpc

* api网关调用svr的协议改为grpc

* 部署测试方式几乎完全一样
 
* 项目结构略有不同

注意

* GRPC的health check方式和http方式不同

* GRPC 需要实现GRPC标准的health check,标准的message如下

* google.golang.org/grpc/health/grpc_health_v1

    