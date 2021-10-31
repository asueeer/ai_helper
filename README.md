## go版本需要是1.15(1.14会报错)

## references:

https://silenceper.com/wechat/

## 代码分层

### 主要参考了https://tech.meituan.com/2017/12/22/ddd-in-practice.html

-handler // 控制器,用来接收用户请求 \
-service // 应用服务层, \
-domain // 领域层 \
-dal // 数据持久化层

-domain \
--aggregate // 聚合根, 一个由多个实体聚合起来的概念 \
--entity // 实体, 有一个唯一id标识 \
--val_obj // 值对象, 没有一个唯一id标识 \
--model // 一些req,resp定义在这里 \
---vo // view_object, 展示给前端的 \
--service // 领域服务层

## id生成器

参考 https://chai2010.cn/advanced-go-programming-book/ch6-cloud/ch6-01-dist-id.html

但我们没有做成一个分布式的id生成器, 而是把workerId写死在了代码里, 这个是不太优雅的地方

### todo 密码都是明文写在项目里的，之后走配置中心会比较好

###   

### redis-server运行命令，redis-run:

redis-server /etc/redis.conf

### 统计代码量

find . -name "*.go"  -print | xargs wc -l

### 消息系统设计

为了防止读扩散，采用发件箱/收件箱的方式进行设计。

当Sender发送一条消息后，gateway要做的事：\
1.将该消息发送给消息队列进行后续的处理。（实现解耦）

当消息队列收到一条消息后，要做的是：

1. 将该消息存到发件箱. 持久化消息。
2. 进行写扩散，将这条消息分成若干份发到群聊中的每个人的收件箱。

当接收方读消息的时候, 要做的是：\
从收件箱里找到自己要读的若干消息id；\
收件箱不存消息内容，消息ids找到之后，从发件箱读取消息内容。

### 技术优化点

日志应该用开源的日志组件 \
而不是自带的log包

客服与单聊没有做解耦 

### 有一个很难受的点: 

路由表的控制只是在path上，没有在方法的粒度上进行权限控制。

导致每次删除的时候，都需要校验一下用户是否存在...

但是暂时没想清楚怎么优雅地解决这个问题。