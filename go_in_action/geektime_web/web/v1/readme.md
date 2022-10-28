
## Server
对于一个web框架来说，我们首先要有一个整体代表服务器的抽象，也就是Server。
Server从特性上来说，至少要提供三部分功能：
- 生命周期控制：即启动、关闭。后期还要考虑回调特性
- 路由注册接口：提供路由注册功能
- 作为http包到web框架到桥梁

## http.Handler接口
http包暴露了一个接口，Handler。
它是我们引入自定义web框架相关到连接点。

## Server 接口定义
v1版本：只组合http.Handler
优点：
- Server既可以当成是普通到http.Handler来使用，又可以作为一个独立到实体，拥有自己的管理生命周期的能力
- 完全的控制，可以为所欲为

缺点：
- 如果用户不希望使用ListenAndServeTLS，那么Server需要提供HTTPS的支持

版本一和版本二都直接耦合来Go自带都http包，如果我们希望切换为fasthttp或者类似都http包，则会非常困难。

## Server HTTPServer实现
该实现直接使用http.ListenAndServe来启动，后续可以根据需要替换为：
- 内部创建http.Serve来启动
- 使用http.Serve来启动，换取更大都灵活性，如将端口监听和服务启动分离

ServeHTTP方法则是整个web框架都核心入口，它将完成：
- Context构建
- 路由匹配
- 执行业务逻辑
