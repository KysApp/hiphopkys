# hiphopkys-server::horse
<p align="center">
<img align="left" width="125" src="https://github.com/KysApp/hiphopkys/blob/master/doc/icon60@3x.png?raw=true">
<ul>
<li><strong>赛马游戏PC端</strong>: <a href="https://github.com/KysApp/hiphopkys/tree/master/clients/horse_pc">赛马PC端口</a>
<li><strong>赛马游戏客户端</strong>: <a href="">我的BAZIRIM</a>
<li><strong>赛马游戏Server</strong>: <a href="https://github.com/KysApp/hiphopkys/tree/master/servers">赛马游戏Server</a>
<li><strong>赛马游戏消息队列消费者</strong>: <a href="https://github.com/KysApp/hiphopkys/tree/master/servers/mq/consumer/horse">赛马游戏消息队列消费者</a>
<li><strong>问题反馈</strong>: <a href="">问题反馈</a>
</ul>
</p>
<br>

## 准备
- [赛马游戏PC端](https://github.com/KysApp/hiphopkys/tree/master/servers/horse)
- [BAZIRIM客户端PC端](https://github.com/KysApp/hiphopkys/tree/master/servers/horse)


## 介绍
赛马游戏服务端,golang编写。


### 工具清单
- 缓存使用 [redis](http://redis.io), golang驱动使用[redigo](https://github.com/garyburd/redigo)。
- 消息队列 [nsq](http://nsq.io/), golang驱动使用[go-nsq](https://github.com/nsqio/go-nsq)
- 与游戏端通讯采用[websocket](https://github.com/gorilla/websocket)
- 客户端通讯消息格式[protobuf](https://github.com/google/protobuf)
- 消息队列通讯消息格式[gencode](https://github.com/andyleap/gencode)
- 是哟的web框架[beego](https://beego.me/)
