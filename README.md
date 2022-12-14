## 项目文档
本项目开发环境推荐docker-compose，使用直链方式，放弃服务注册发现中间件（etcd、nacos、consul等）带来的麻烦
测试、线上部署使用k8s+istio（也不需要etcd、nacos、consul等）
## 项目简介
整个项目使用了go-zero开发的微服务，基本包含了go-zero以及相关go-zero作者开发的一些中间件，所用到的技术栈基本是go-zero项目组的自研组件，基本是go-zero全家桶了 另外，前端是小程序，本项目已经对接好了小程序授权登录 以及 微信支付了 ，前端看看后面是否能开源吧
项目目录结构如下：

- app：所有业务代码包含api、rpc以及mq（消息队列、延迟队列、定时任务）
- common：通用组件 error、middleware、interceptor、tool、ctxdata等
- data：该项目包含该目录依赖所有中间件(mysql、es、redis、grafana等)产生的数据，此目录下的所有内容应该在git忽略文件中，不需要提交。
- deploy：
    - filebeat: docker部署filebeat配置
    - go-stash：go-stash配置
    - nginx: nginx网关配置
    - prometheus ： prometheus配置
    - script：
        - gencode：生成api、rpc，以及创建kafka语句，复制粘贴使用
        - mysql：生成model的sh工具
    - goctl: 该项目goctl的template，goctl生成自定义代码模版，tempalte用法可参考go-zero文档，复制到家目录下.goctl即可
- doc : 该项目系列文档
- modd.conf : modd热加载配置文件，不要怕～它用起来很简单，关于modd更多用法可以去这里了解 ： [https://github.com/cortesi/modd](https://github.com/cortesi/modd) ， 本项目镜像只是将golang-1.17.7-alpine作为基础镜像安装了modd在内部，如果你想把goctl、protoc、golint等加进去，不用我的镜像直接制作一个镜像也一样的哈
## 思维导图
![go-zero微服务电商系统 (1).png](https://cdn.nlark.com/yuque/0/2022/png/21644241/1661847508555-2f6d7926-ce49-4c0a-81dd-dee937d95c4d.png#clientId=u73b4c332-6a55-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=1500&id=u44e42430&margin=%5Bobject%20Object%5D&name=go-zero%E5%BE%AE%E6%9C%8D%E5%8A%A1%E7%94%B5%E5%95%86%E7%B3%BB%E7%BB%9F%20%281%29.png&originHeight=1875&originWidth=924&originalType=binary&ratio=1&rotation=0&showTitle=false&size=143500&status=done&style=none&taskId=uda4d9f9e-31ed-4d2b-9d47-e91a53b8aa3&title=&width=739.2)
## 系统架构图
![未命名文件 (1).png](https://cdn.nlark.com/yuque/0/2022/png/21644241/1662394469443-72ac852f-ec50-4224-a1b4-24d9d0ae4644.png#clientId=u7034512a-07ae-4&crop=0&crop=0&crop=1&crop=1&from=ui&id=u4089ff2c&margin=%5Bobject%20Object%5D&name=%E6%9C%AA%E5%91%BD%E5%90%8D%E6%96%87%E4%BB%B6%20%281%29.png&originHeight=595&originWidth=1534&originalType=binary&ratio=1&rotation=0&showTitle=false&size=64077&status=done&style=none&taskId=u6aa23c16-d8f4-4c3a-9b59-cf0d80fc303&title=)
## 业务架构图
待上传
## 网关
nginx做对外网关，网关前面是slb，另外，很多同学觉得nginx做网关不太好，这块原理基本一样，可以自行替换成apisix、kong等
## 开发模式
本项目使用的是微服务开发，api （http） + rpc（grpc） ， api充当聚合服务，复杂、涉及到其他业务调用的统一写在rpc中，如果一些不会被其他服务依赖使用的简单业务，可以直接写在api的logic中
## 日志
关于日志，统一使用filebeat收集，上报到kafka中，由于logstash懂得都懂，资源占用太夸张了，这里使用了go-stash替换了logstash
链接：[https://github.com/kevwan/go-stash](https://github.com/kevwan/go-stash) ， go-stash是由go-zero开发团队开发的，性能很高不占资源，主要代码量没多少，只需要配置就可以使用，很简单。它是吧kafka数据源同步到elasticsearch中，默认不支持elasticsearch账号密码，我fork了一份修改了一下，很简单支持了账号、密码
## 监控
监控采用prometheus，这个go-zero原生支持，只需要配置就可以了，这里可以看项目中的配置
## 链路追踪（isito中插件替代）
go-zero默认jaeger、zipkin支持，只需要配置就可以了，可以看配置
```
#链路追踪
Telemetry:
  Name: goods-api
  Endpoint: http://127.0.0.1:14268/api/traces
  Sampler: 1.0
  Batcher: jaeger
```

## 熔断限流（istio 管理 ）
## 分布式锁
使用 redis **redsync分布式锁**
goredislib **"github.com/go-redis/redis/v8"
"github.com/go-redsync/redsync/v4"
**goredis **"github.com/go-redsync/redsync/v4/redis/goredis/v8"**
测试用例：
```
func newTestMutexes(pools []redis.Pool, name string, n int) []*Mutex {
	mutexes := make([]*Mutex, n)
	for i := 0; i < n; i++ {
		mutexes[i] = &Mutex{
			name:          name,
			expiry:        8 * time.Second,
			tries:         32,
			delayFunc:     func(tries int) time.Duration { return 500 * time.Millisecond },
			genValueFunc:  genValue,
			driftFactor:   0.01,
			timeoutFactor: 0.05,
			quorum:        len(pools)/2 + 1,
			pools:         pools,
		}
	}
	return mutexes
}
```
## 发布订阅
kafka ， 发布订阅使用的是go-zero开发团队开发的go-queue，链接：[https://github.com/zeromicro/go-queue](https://github.com/zeromicro/go-queue)
这里使用kq，kq是基于kafka做的高性能发布订阅

## 消息队列
使用rocketmq 普通消息，同一个消息的topic需要指定相同，
生产者：[https://github.com/jackjie2016/zero-mall2/blob/main/examples/rocketmq_test/producer/main.go](https://github.com/jackjie2016/zero-mall2/blob/main/examples/rocketmq_test/producer/main.go)
消费者：[https://github.com/jackjie2016/zero-mall2/blob/main/examples/rocketmq_test/Consumer/main.go](https://github.com/jackjie2016/zero-mall2/blob/main/examples/rocketmq_test/Consumer/main.go)
## 分布式事务
使用rocketmq 事务消息
![050afdf9d4aecdb5fef2116d89a1b4ad.png](https://cdn.nlark.com/yuque/0/2022/png/21644241/1661842337366-b8b64ddb-5548-47b9-ac08-7d138fef6e78.png#clientId=u3eb1c3c7-23d9-4&crop=0&crop=0&crop=1&crop=1&from=ui&id=u0d0d872f&margin=%5Bobject%20Object%5D&name=050afdf9d4aecdb5fef2116d89a1b4ad.png&originHeight=477&originWidth=1006&originalType=binary&ratio=1&rotation=0&showTitle=false&size=35963&status=done&style=none&taskId=u380867d1-0efd-4bb6-ace7-db0f661afbc&title=)
RocketMQ事务消息（Transactional Message）是指应用本地事务和发送消息操作可以被定义到全局事务中，要么同时成功，要么同时失败。RocketMQ的事务消息提供类似 X/Open XA 的分布式事务功能，通过事务消息能达到分布式事务的最终一致。

Apache RocketMQ在4.3.0版中已经支持分布式事务消息，采用了2PC（两阶段提交）+ 补偿机制（事务状态回查）的思想来实现了提交事务消息，同时增加一个补偿逻辑来处理二阶段超时或者失败的消息，如下图所示。
事务消息主要分为两个流程:
(1)、正常事务消息的发送及提交

- a、生产者发送half消息到Broker服务端（半消息）；

半消息是一种特殊的消息类型，该状态的消息暂时不能被Consumer消费。当一条事务消息被成功投递到Broker上，但是Broker并没有接收到Producer发出的二次确认时，该事务消息就处于"暂时不可被消费"状态，该状态的事务消息被称为半消息。

- b、Broker服务端将消息持久化之后，给生产者响应消息写入结果（ACK响应）；

- c、生产者根据发送结果执行本地事务逻辑（如果写入失败，此时half消息对业务不可见，本地逻辑不执行）；

- d、生产者根据本地事务执行结果向Broker服务端提交二次确认（Commit 或是 Rollback），Broker服务端收到 Commit 状态则将半事务消息标记为可投递，订阅方最终将收到该消息；Broker服务端收到 Rollback 状态则删除半事务消息，订阅方将不会接收该消息；

(2)、事务消息的补偿流程

- a、在网络闪断或者是应用重启的情况下，可能导致生产者发送的二次确认消息未能到达Broker服务端，经过固定时间后，Broker服务端将会对没有Commit/Rollback的事务消息（pending状态的消息）进行“回查”；

- b、生产者收到回查消息后，检查回查消息对应的本地事务执行的最终结果；

- c、生产者根据本地事务状态，再次提交二次确认给Broker，然后Broker重新对半事务消息Commit或者Rollback；

其中，补偿阶段用于解决消息Commit或者Rollback发生超时或者失败的情况。

事务消息共有三种状态，提交状态、回滚状态、中间状态：
TransactionStatus.CommitTransaction：提交事务，它允许消费者消费此消息。
TransactionStatus.RollbackTransaction：回滚事务，它代表该消息将被删除，不允许被消费。
TransactionStatus.Unknown：中间状态，它代表需要回查本地事务状态来决定是提交还是回滚事务。
下面我们通过示例演示如何使用RocketMQ的事务消息。
链接:[ https://github.com/jackjie2016/zero-mall2/blob/main/examples/rocketmq_test/transaction/main.go](https://github.com/jackjie2016/zero-mall2/blob/main/examples/rocketmq_test/transaction/main.go)
## 延迟消息
使用rocketmq 延迟消息
当消息写入到Broker后，不能立刻被消费者消费，需要等待指定的时长后才可被消费处理的消息，称为延时消息。
RocketMQ延时消息的延迟时长不支持随意时长的延迟，是通过特定的延迟等级来指定的。默认支持18个等级的延迟消息，延时等级定义在RocketMQ服务端的MessageStoreConfig类中的如下变量中：
```
// MessageStoreConfig.java
private String messageDelayLevel = "1s 5s 10s 30s 1m 2m 3m 4m 5m 6m 7m 8m 9m 10m 20m 30m 1h 2h";
 
//发消息时，设置delayLevel等级即可：msg.setDelayLevel(level)。level有以下三种情况：
level == 0，消息为非延迟消息
1<=level<=maxLevel，消息延迟特定时间，例如level==1，延迟1s
level > maxLevel，则level== maxLevel，例如level==20，延迟2h
```
例如指定的延时等级为3，则表示延迟时长为10s，即延迟等级是从1开始计数的
链接：[https://github.com/jackjie2016/zero-mall2/blob/main/examples/rocketmq_test/delay/main.go](https://github.com/jackjie2016/zero-mall2/blob/main/examples/rocketmq_test/delay/main.go)
## 基于Es的模糊搜索
中文搜索需要安装ik分词器，ik分词器插件版本需要跟es的版本一致
ik分词器有ik_smart 和 ik_max_word，需要使用ik_max_word分词力度比较大，ik_smart 的分词力度小
复合查询
```
{
    "query":{
        "bool":{
            "must":[
            ],
            "should":[
            ],
            "must_not":[
            ],
            "filter":[
            ],
        }
    }
}

```
### must: 必须匹配,查询上下文,加分
should: 应该匹配,查询上下文,加分
must_not: 必须不匹配,过滤上下文,过滤
filter: 必须匹配,过滤上下文,过滤
## 测试用例
表格驱动测试，目前实现了几个模板
```
func TestSetinv(t *testing.T) {
	tests := []struct {
		name      string
		GoodsInfo []*proto.GoodsInvInfo
		OrderSn   string
		err       error
	}{
		{
			name: "ok",
			GoodsInfo: []*proto.GoodsInvInfo{
				{GoodsId: 421, Num: 1},
				//{GoodsId: 2, Num: 1},
			},
			OrderSn: GenerateOrderSn(),
			err:     nil,
		},
		{
			name: "Inventory_Internal",
			GoodsInfo: []*proto.GoodsInvInfo{
				{GoodsId: 421, Num: 1},
				//{GoodsId: 2, Num: 1},
			},
			OrderSn: GenerateOrderSn(),
			err:     status.Errorf(codes.Internal, "保存库存扣减历史失败"),
		},
		{
			name: "order_sn_inserted",
			GoodsInfo: []*proto.GoodsInvInfo{
				{GoodsId: 421, Num: 1},
				//{GoodsId: 2, Num: 1},
			},
			OrderSn: GenerateOrderSn(),
			err:     status.Errorf(codes.NotFound, "没有库存信息"),
		},
		{
			name: "goods_NotFound",
			GoodsInfo: []*proto.GoodsInvInfo{
				{GoodsId: 421, Num: 1},
				//{GoodsId: 2, Num: 1},
			},
			OrderSn: GenerateOrderSn(),
			err:     status.Errorf(codes.NotFound, "没有库存信息"),
		},
		{
			name: "ResourceExhausted",
			GoodsInfo: []*proto.GoodsInvInfo{
				{GoodsId: 421, Num: 1000},
				//{GoodsId: 2, Num: 1},
			},
			OrderSn: GenerateOrderSn(),
			err:     status.Errorf(codes.ResourceExhausted, "库存不足"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			_, err := InventoryClient.Sell(context.Background(), &proto.SellInfo{
				GoodsInfo: test.GoodsInfo,
				OrderSn:   test.OrderSn,
			})

			if err != test.err {
				t.Errorf("%v", err)
			}
		})
	}
}
```
基准测试（benchmark）
```
func BenchmarkSetinv(b *testing.B) {
	for k := 0; k < b.N; k++ {

	}
}
```
链接：[https://github.com/jackjie2016/zero-mall2/tree/main/test](https://github.com/jackjie2016/zero-mall2/tree/main/test)
## 部署
本项目开发环境推荐k8s+istio，使用直链方式，放弃服务注册发现中间件（etcd、nacos、consul等）带来的麻烦,配置文件（例子）：
```
UserRpc:
  Endpoints:
   - user-svc-svc:80
```
go-zero 推荐使用target：//k8s/user-svc-svc:80

gitee+ jenkins + 阿里云镜像仓库（harbor ）+ k8s+istio
在jenkins中点击部署对应的服务，会去gitee拉取代码-->再去拉取线上配置（线上配置单独一个git库，为什么不用配置中心，部署文档中有介绍）---->自动构建镜像-->推送到镜像仓库--->使用kubectl自动发布到k8s中---->前面要挂一个nginx做网关统一入口
