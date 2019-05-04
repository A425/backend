# backend

## 教程
https://micro.mu/docs/cn/toolkit.html

## 基本上
最微型的一个环境是
- 用consul做服务注册服务发现，以下的所有服务都要通过传入 --registry=consul 注册到consul集群
- `micro --registry=consul api` 起一个api gateway，作为面向小程序前端唯一入口，把http请求转发给对应的api服务，也就是后端服务的第一层
- api gateway 后面是具体的api服务，也就是通过 `micro new --type "api"` 生成的服务，通过api gateway的请求，调用api服务后边的rpc服务进行业务处理，也就是后端服务第二层
- 第三层的 rpc 服务，也就是通过 `micro new ` 生成的服务，是经过业务细分后的功能相对简单独立的系统，纯内部系统，不对外开放，值暴露给第二层的api服务
- `micro --registry=consul --enable_stats --web_address 127.0.0.1:8585 web` 可以起一个web服务，用来观察当前微服务集群的基本状态
