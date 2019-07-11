# 运行
## 启动服务
* make love 运行
* make  编译
* 关闭服务 请用 kill -2,以便服务可以优雅的关闭
# 架构
## go
* 版本1.12以上
* 框架**gin**
* 采用**restful**风格编写
* go modules 管理pkg
* [redisgo](https://godoc.org/github.com/gomodule/redigo/redis) (Redis包,适合写原生语句)
* [go-redis](https://github.com/go-redis/redis) (Redis包,高度封装,本项目默认使用,推荐使用！！！)
* [gorm](http://gorm.book.jasperxu.com/crud.html#u) (MySQL包)
* [govalidator](https://godoc.org/github.com/asaskevich/govalidator) (验证包)
* [jaeger-client-go](https://github.com/jaegertracing/jaeger-client-go) (Jaeger包)
* [sarama](https://github.com/Shopify/sarama) (Kafka包)
* [amqp](https://github.com/streadway/amqp) (RabbitMQ包)
* [tollbooth](https://github.com/didip/tollbooth) (Token Bucket算法限流,支持ip，token，用户，方法)
* [limiter](https://github.com/ulule/limiter) (简单的限流)
* [viper](github.com/spf13/viper) (配置文件)
* [zap](https://godoc.org/go.uber.org/zap) (log包)
## redis
版本 5.0以上,如果不用streams可以不用
## mysql
版本 5.7以上
```
├── admin.sh                     # 进程的start|stop|status|restart控制文件
├── conf                         # 配置文件统一存放目录
│   ├── config.yaml              # 配置文件
│   ├── server.crt               # TLS配置文件
│   └── server.key
├── config                       # 专门用来处理配置和配置文件的Go package
├── domain                       # 领域层 负责写业务逻辑，领域层不能相互调用形成依赖，依赖反转到service层
│   ├── entity                   # 领域实体
│   └── value-object             # 值对象
├── handler                      # 类似MVC架构中的C，用来读取输入，验证数据，并将处理流程转发给service层，最后返回结果  
├── main.go                      # Go程序唯一入口
├── repositories                 # 基础设施层 础设施层是系统中的技术密集部分,它为领域层、应用层的业务服务（例如持久化、消息通信等等）提供具体的技术支持
│   ├── data-mappers             # 数据存储交互转化
│   └── util                     # 存放自定义函数
├── log                          # 日志目录
├── pkg                          # 第三方包
├── router                       # 路由相关处理
│   └── middleware               # API服务器用的是Gin Web框架，Gin中间件存放位置
├── service                      # 应用层 DDD的application层 应用层定义系统功能，并指挥领域层中的领域对象实现这些功能。简单的来书只控制顺序，所有业务逻辑有domain实现，调度他们
├── util                         # 工具类函数存放目录 不在新增 新增全部在 infrastructure层 得repositories目录下
└── Makefile                     # 包含打包运行等命令
```
# 简单讲一下DDD
## DDD是什么
> DDD是一种设计思想，它本身不绑定到任何一种具体的架构风格，可以应用在多种不同的架构风格中。本文探讨在经典的分层架构中如何应用DDD，以及在DDD的语境下，分层结构每一层的具体职责。
## DDD的组成
整个系统划分为基础设施层（Infrastructure）、领域层（Domain）、应用层（Application）和用户接口层（User Interface，也称为表示层）。下面从各个维度分别讨论之。
### DDD的职责分配
#### 领域层（Domain Layer）
**领域层实现业务逻辑。**
什么是业务逻辑？业务逻辑就是存在于问题域即业务领域中的实体、概念、规则和策略等，与软件实现无关
#### 应用层（Application Layer）
**应用层定义系统功能，并指挥领域层中的领域对象实现这些功能。**
应用层是整个系统的功能外观，封装了领域层的复杂性并隐藏了其内部实现机制。
#### 基础设施层（Infrastructure Layer）
**基础设施层为其余各层提供技术支持。**
基础设施层是系统中的技术密集部分。它为领域层、应用层的业务服务（例如持久化、消息通信等等）提供具体的技术支持
#### 用户接口层（User Interface）
**用户接口层为外部用户访问底层系统提供访问界面和数据表示。**
用户接口层在底层系统之上封装了一层可访问外壳，为特定类型的外部用户（人或计算机程序）访问底层系统提供访问入口，并将底层系统的状态数据以该类型客户需要的形式呈现给它们。
### 关键点
* 基础设施层和其他各层的编译时依赖关系和运行时调用关系是相反的：在运行时，其他各层中的对象调用基础设施层中的对象实例，使用后者提供的服务；而在编译时，基础设施层中的类依赖于其他各层（主要是领域层）中的类。这是通过运用面向对象原则中的依赖倒置原则达到的，在领域层中定义服务接口，而在基础设施层中实现领域层定义的接口。
* 代表业务的层（领域层和应用层）不依赖于代表技术的层（基础设施层和用户接口层），代表技术的层依赖于代表业务的层。
* 领域层处于整个系统的核心位置，它在概念上不依赖于其他各层，其他各层都直接或间接依赖于它。领域层是整个系统的核心引擎，直接实现业务目标，攸关业务正确性、可靠性、灵活性和扩展性。
* 领域层应该是整个系统中最“胖”的一层，因为它实现了全部业务逻辑并且通过各种校验手段保证业务正确性，其余各层相对都较“瘦”。如果你的代码中不是如此，你肯定走错了路。胖用户接口层是“以数据库为中心的增删改查”模式的典型症状，胖应用层是事务脚本模式的典型症状。
## 流程图
![](https://i.loli.net/2019/04/29/5cc669fb177db.png)
