# go-zero开源项目介绍

> 部门：晓黑板研发部
>
> git地址：

## 1. go-zero框架背景

18年初，晓黑板后端在经过频繁的宕机后，决定从`Java+MongoDB`的单体架构迁移到微服务架构，经过仔细思考和对比，我们决定：

* 基于Go语言
  * 高效的性能
  * 简洁的语法
  * 广泛验证的工程效率
  * 极致的部署体验
  * 极低的服务端资源成本
* 自研微服务框架
  * 个人有过很多微服务框架自研经验
  * 需要有更快速的问题定位能力
  * 更便捷的增加新特性

## 2. go-zero框架设计思考

对于微服务框架的设计，我们期望保障微服务稳定性的同时，也要特别注重研发效率。所以设计之初，我们就有如下一些准则：

* 保持简单
* 高可用
* 高并发
* 易扩展
* 弹性设计，面向故障编程
* 尽可能对业务开发友好，封装复杂度
* 尽可能约束做一件事只有一种方式

我们经历不到半年时间，彻底完成了从`Java+MongoDB`到`Golang+MySQL`为主的微服务体系迁移，并于18年8月底完全上线，稳定保障了晓黑板后续增长，确保了整个服务的高可用。

## 3. go-zero项目实现和特点

go-zero是一个集成了各种工程实践的包含web和rpc框架，有如下主要特点：

* 强大的工具支持，尽可能少的代码编写
* 极简的接口
* 完全兼容net/http
* 支持中间件，方便扩展
* 高性能
* 面向故障编程，弹性设计
* 内建服务发现、负载均衡
* 内建限流、熔断、降载，且自动触发，自动恢复
* API参数自动校验
* 超时级联控制
* 自动缓存控制
* 链路跟踪、统计报警等
* 高并发支撑，稳定保障了晓黑板疫情期间每天的流量洪峰

如下图，我们从多个层面保障了整体服务的高可用：

![弹性设计](/Users/kevin/Develop/go/opensource/go-zero/doc/images/resilience.jpg)

## 4. go-zero框架收益

* 保障大并发服务端的稳定性，经受了充分的实战检验
* 极简的API定义
* 一键生成Go, iOS, Android, Dart, TypeScript, JavaScript代码，并可直接运行
* 服务端自动校验参数合法性

## 5. go-zero近期开发计划

* 升级grpc到最新版
* 加入Power of Two Choices Peak Exponentially Weighted Moving Average负载均衡算法
* 自动生成API mock server，便于客户端开发
* 自动生成服务端功能测试