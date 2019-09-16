# GO DDD 最佳实践
## 分层架构
### 领域模型设计
#### presentation 展现层

* ui 

#### gateway 网关层

-----------------------------------------------------------------------------------

#### interface 接口层
接口服务位于用户接口层，用于处理用户发送的 Restful 请求和解析用户输入的配置文件等，并将信息传递给应用层。
**可考虑直接使用facade命名**
* dto
* facade接口

#### application 应用层

* event 应用事件
应用层事件和领域事件都统一放这里
  
    * 应用事件处理
    * 应用事件发布
  
* service 应用服务

  应用服务对多个领域服务或外部应用服务进行封装、编排和组合，对外提供粗粒度的服务。

* repository 仓储接口

#### domain 领域层

* aggregate 聚合根
  * entity 实体
  * vaule object 值对象
  * command
  * query
  * projector
  * saga
  
* event 领域事件

* service 领域服务

  对多个不同实体对象操作的领域服务

* repository 仓储接口
          存放聚合对应的查询或持久化领域对象的代码，通常包括仓储接口和仓储实现方法。为了方便聚合的拆分和组合，我们设定一个原则：一个聚合对应一个仓储

   按照 DDD 分层原则，仓储实现本应属于基础层代码，但为了微服务代码拆分和重组的便利性，我们把聚合的仓储实现代码放到了领域层对应的聚合代码包内

#### infrastructure 基础设施层
* config 应用配置
* repository 仓储实现
* cache 缓存
* mq 消息队列
* eventbus 事件总线
* 第三方包

-----------------------------------------------------------------------------------

## 文件结构

### 微服务
#### ddd/micro.v1
├── identity
│   ├── application
│   │   ├── event
│   │   │   ├── publish
│   │   │   └── subscribe
│   │   ├── repository
│   │   └── service
│   ├── domain
│   │   ├── role
│   │   │   ├── event
│   │   │   ├── model
│   │   │   │   ├── aggregate.go
│   │   │   │   ├── command.go
│   │   │   │   ├── entity.go
│   │   │   │   └── value_object.go
│   │   │   ├── repository
│   │   │   └── service
│   │   └── user
│   │       ├── event
│   │       ├── model
│   │       │   ├── aggregate.go
│   │       │   ├── command.go
│   │       │   ├── entity.go
│   │       │   └── value_object.go
│   │       ├── repository
│   │       └── service
│   ├── infrastructure
│   └── interface
│       ├── dto
│       └── facade
└── shop

#### ddd/micro.v2
├── identity
│   ├── application
│   │   ├── event
│   │   │   ├── publish
│   │   │   └── subscribe
│   │   ├── repository
│   │   └── service
│   │       └── auth
│   ├── domain
│   │   ├── event
│   │   ├── model
│   │   │   ├── role
│   │   │   │   ├── event
│   │   │   │   ├── model
│   │   │   │   │   ├── aggregate.go
│   │   │   │   │   ├── command.go
│   │   │   │   │   ├── entity.go
│   │   │   │   │   └── value_object.go
│   │   │   │   ├── repository
│   │   │   │   └── service
│   │   │   └── user
│   │   │       ├── aggregate.go
│   │   │       ├── command.go
│   │   │       ├── entity.go
│   │   │       └── value_object.go
│   │   ├── repository
│   │   └── service
│   ├── infrastructure
│   │   ├── cache
│   │   ├── config
│   │   ├── eventstore
│   │   ├── lib
│   │   ├── mq
│   │   └── repository
│   └── interface
│       ├── dto
│       └── facade
│           ├── grpc
│           ├── rest
│           └── rpc
└── shop

### 单体应用
#### ddd/monolith.v1
├── application
│   ├── eventstore
│   │   └── event.go
│   └── service
│       ├── domain
│       │   ├── model
│       │   │   ├── role.go
│       │   │   └── user.go
│       │   └── service
│       │       └── user_service.go
│       ├── repository
│       │   ├── memory
│       │   │   └── user.go
│       │   ├── mysql
│       │   │   └── user_repository.go
│       │   ├── po
│       │   └── user_repository.go
│       └── user_service.go
└── rest
    └── user_handler.go

#### ddd/monolith/v2
├── application 应用层
│   ├── event 事件处理
│   │   └── event.go
│   └── service 应用服务
│       └── user_service.go
├── domain 领域层
│   └── model 领域模型
│       ├── account 聚合, 以实际业务属性的名称命名
│       │   ├── entity 聚合根/实体/值对象
│       │   │   ├── role.go
│       │   │   └── user.go
│       │   ├── repository 资源库
│       │   │   └── user_repository.go
│       │   └── service 领域服务
│       │       └── user_service.go
│       ├── blog 聚合
│       └── event 领域事件
│           └── service 领域服务
├── port 端口层
│   └── adapter 适配器
│       ├── event 事件处理器
│       │   └── event.go
│       ├── persistence 持久化处理器
│       │   └── repository 
│       │       └── mysql 资源库实现
│       │           └── user_repository.go
│       ├── rest REST 服务
│       │   ├── config
│       │   ├── dto/vo 数据迁移对象或视图模型
│       │   ├── handler http请求处理器
│       │   │   └── user_handler.go
│       │   └── server
│       │       └── server.go
│       └── service 服务实现 
│           └── log_service.go
└── presentation 展现层
    └── rest REST 应用
        ├── config.yml
        └── main.go

#### 调用关系
presentation layer -> port layer -> application layer -> domain layer

## 引用包
### EventStore 事件存储
- [eventhorizon](https://github.com/looplab/eventhorizon) CQRS/ES toolkit for Go

### DI 依赖注入
- [dig](https://godoc.org/go.uber.org/dig) dig provides an opinionated way of resolving object dependencies
- [fx](https://godoc.org/go.uber.org/fx) A dependency injection based application framework for Go

## 参考

## todo
- 应用层自治
- 应用层事务管理

## 问题



