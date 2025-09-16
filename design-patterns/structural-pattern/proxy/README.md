# 代理模式 (Proxy Pattern)

## 概述

代理模式是一种结构型设计模式，让你能够提供对象的替代品或其占位符。代理控制着对于原对象的访问，并允许在将请求提交给对象前后进行一些处理。

## 实现说明

本示例模拟了Nginx作为Web服务器代理的场景：

### 组件说明

1. **server.go** - 服务接口
   - 定义了`server`接口，声明了`handleRequest`方法
   - 代理和真实服务都必须实现这个接口

2. **application.go** - 真实服务
   - `Application`结构体实现了实际的业务逻辑
   - 处理具体的HTTP请求并返回相应的状态码和响应体

3. **nginx.go** - 代理实现
   - `Nginx`结构体作为代理，实现了`server`接口
   - 提供访问控制和限流功能
   - 在转发请求给真实服务前进行预处理

4. **main.go** - 客户端代码
   - 演示如何使用代理服务器
   - 展示限流功能的效果

### 代理模式的优势

1. **访问控制** - 可以控制对服务对象的访问
2. **缓存** - 可以缓存请求结果
3. **限流** - 可以限制请求频率
4. **日志记录** - 可以记录请求日志
5. **延迟初始化** - 可以延迟创建重量级对象

## 运行示例

```bash
cd design-patterns/structural-pattern/proxy
go run .
```

## 预期输出

```
Url: /app/status
HttpCode: 200
Body: Ok

Url: /app/status
HttpCode: 200
Body: Ok

Url: /app/status
HttpCode: 403
Body: Not Allowed

Url: /create/user
HttpCode: 201
Body: User Created

Url: /create/user
HttpCode: 404
Body: Not Ok
```

## 应用场景

- **虚拟代理** - 延迟初始化重量级对象
- **保护代理** - 控制对敏感对象的访问
- **远程代理** - 为远程对象提供本地代表
- **缓存代理** - 缓存昂贵操作的结果
- **智能引用** - 在访问对象时执行额外操作
