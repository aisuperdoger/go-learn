# sync.Once 简单使用示例

这个包演示了 Go 语言中 `sync.Once` 的基本用法。

## 什么是 sync.Once？

`sync.Once` 是 Go 标准库中的一个同步原语，用于确保某个函数只被执行一次，即使在多个 goroutine 并发调用的情况下也是如此。



## 主要特性

- **线程安全**: 多个 goroutine 可以安全地同时调用 `once.Do()`
- **只执行一次**: 无论调用多少次，传入的函数只会执行一次
- **阻塞等待**: 如果协程正在执行once.Do(initDatabase)，其他协程调用once.Do(initDatabase)时，其他协程会阻塞等待第一个协程执行完成once.Do(initDatabase)。

## 常见使用场景

- 单例模式实现
- 数据库连接初始化
- 配置文件加载
- 缓存预热
- 程序关闭清理

简单、高效、安全，是 Go 并发编程中的重要工具。
