


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