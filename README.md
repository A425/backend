# backend

## 教程
https://micro.mu/docs/cn/toolkit.html

## 基本上
最微型的一个环境是
- 用consul做服务注册服务发现和配置管理，以下的所有服务都要通过启动时传入 --registry=consul参数 注册到consul集群
- `micro --registry=consul api --handler=http` 起一个api gateway，作为面向小程序前端唯一入口，把http请求转发给对应的api服务，也就是后端服务的第一层
- api gateway 后面是具体的api服务，所有通过 `micro new --type "api"` 生成的服务，也就是后端服务第二层，使用gin处理http请求，转发请求给后边的rpc服务进行业务处理
- 第三层的 rpc 服务，也就是通过 `micro new ` 生成的服务，是经过业务细分后的功能相对简单独立的系统，纯内部系统，不对外开放，只暴露给第二层的api服务
- `micro --registry=consul --enable_stats --web_address 127.0.0.1:8585 web` 可以起一个web服务，用来观察当前微服务集群的基本状态

## 主要的研发内容
consul和micro的api gateway 是现成工具无需研发，直接使用。
第二层的api服务和第三层的rpc服务是研发重点。
以及后续可能要用到的日志收集或监控打点等

## 部署
### 生产环境

草稿：最好的情况
- 配置线上域名解析
- 配置负载均衡器 指到API gateway
- 所有服务配置自动扩容
- 关系型数据库 1个
- consul用于服务发现和配置管理，min 2 max 5（机器数量）
- api gateway 代理前端请求，min 2 max 5
- 检测状态用的web服务，节约成本的话可以考虑和consul一起部署。
- 每组RPC微服务，min 2 max 5，假如两组微服务就要 min 2 * 2，max 5 * 2

### 测试环境
- 关系型数据库 1个
- consul，api gateway，Web monitor服务部署一起，min max 2
- RPC 服务 2个
