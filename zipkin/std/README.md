## zipkin standard

#### span
>一个span理解为一个调用单元，A->B->C，则需要启动3个span。
启动A的时候，是root span，没有parent span，
启动B的时候，b的span的parent是A的span，
C同理，需要一个parent。
没有parent的span视作root span。

#### 场景

* 在a.call()中，同步调用b,c,d远程接口。
* 在x.call()中，调用y,在y中调用d。
* example()方法为最基本的使用方式，模拟了3个span，注意set parent span
* 远程调用的时候，parent span一般会放到requet的header中。