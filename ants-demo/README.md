# ants 协程池简单使用示例

本示例演示如何使用 [ants](https://github.com/panjf2000/ants) 高性能协程池库。

## 什么是 ants？

ants 是一个高性能的 goroutine 池，实现了对大规模 goroutine 的调度管理、goroutine 复用，允许使用者在开发并发程序的时候限制 goroutine 数量，复用资源，达到更高效执行任务的效果。

## 主要特性

- 🚀 **高性能**: 通过复用 goroutine 减少创建和销毁开销
- 📊 **资源控制**: 限制并发 goroutine 数量，避免资源耗尽
- 🛡️ **内存安全**: 自动管理 goroutine 生命周期
- 📈 **监控支持**: 提供丰富的运行时统计信息
- ⚙️ **灵活配置**: 支持多种配置选项

## 运行示例

```bash
cd ants-demo
go run main.go
```

## 示例内容

### 1. 基本协程池使用

演示如何创建协程池并提交任务：

```go
// 创建容量为10的协程池
pool, err := ants.NewPool(10)
if err != nil {
    log.Fatal(err)
}
defer pool.Release()

// 提交任务
err = pool.Submit(func() {
    // 你的任务代码
    fmt.Println("任务执行中...")
})
```

### 2. 带参数的协程池（PoolWithFunc）

使用固定函数的协程池，适合重复执行相同逻辑的场景：

```go
// 定义任务函数
taskFunc := func(i interface{}) {
    taskID := i.(int)
    fmt.Printf("处理任务 %d\n", taskID)
}

// 创建带函数的协程池
poolWithFunc, err := ants.NewPoolWithFunc(5, taskFunc)
if err != nil {
    log.Fatal(err)
}
defer poolWithFunc.Release()

// 提交带参数的任务
err = poolWithFunc.Invoke(123) // 传递参数123
```

### 3. 性能对比测试

对比 ants 协程池与原生 goroutine 的性能差异：

- 大量短任务：ants 通常更优（减少创建销毁开销）
- 长时间任务：差异较小，但 ants 提供更好的资源控制
- 内存使用：ants 更稳定，避免 goroutine 泄漏

### 4. 协程池配置和监控

展示如何配置协程池参数并监控运行状态：

```go
pool, err := ants.NewPool(
    20, // 协程池大小
    ants.WithOptions(ants.Options{
        ExpiryDuration:   3 * time.Second, // 空闲goroutine过期时间
        PreAlloc:         true,            // 预分配goroutine
        MaxBlockingTasks: 100,             // 最大阻塞任务数
        Nonblocking:      false,           // 允许阻塞
    }),
)

// 监控协程池状态
fmt.Printf("运行中: %d, 空闲: %d, 容量: %d\n",
    pool.Running(), pool.Free(), pool.Cap())
```

## 使用场景

### 适合使用 ants 的场景：

- ✅ **高并发任务处理**: 如批量数据处理、文件操作
- ✅ **资源受限环境**: 需要控制并发数量的场景
- ✅ **长期运行的服务**: 需要稳定内存使用的应用
- ✅ **性能敏感应用**: 需要减少 goroutine 创建开销

### 不适合的场景：

- ❌ **任务数量很少**: 原生 goroutine 更简单
- ❌ **任务执行时间很长**: 协程池优势不明显
- ❌ **不需要并发控制**: 简单应用场景

## 配置选项说明

| 配置项 | 类型 | 默认值 | 说明 |
|-------|------|--------|------|
| `ExpiryDuration` | `time.Duration` | 1秒 | 空闲worker过期时间 |
| `PreAlloc` | `bool` | false | 是否预分配worker |
| `MaxBlockingTasks` | `int` | 0 | 最大阻塞任务数，0表示无限制 |
| `Nonblocking` | `bool` | false | 是否非阻塞模式 |
| `PanicHandler` | `func(interface{})` | nil | panic处理函数 |

## 最佳实践

### 1. 合理设置池大小
```go
// 根据系统CPU核心数设置
poolSize := runtime.NumCPU() * 2

// 或根据业务需求设置
poolSize := 100 // 适合I/O密集型任务
```

### 2. 及时释放资源
```go
pool, err := ants.NewPool(10)
if err != nil {
    log.Fatal(err)
}
defer pool.Release() // 确保释放资源
```

### 3. 错误处理
```go
err := pool.Submit(func() {
    // 任务代码
})
if err != nil {
    log.Printf("提交任务失败: %v", err)
    // 处理错误，可能需要重试或其他逻辑
}
```

### 4. 监控和调优
```go
// 定期监控协程池状态
go func() {
    ticker := time.NewTicker(30 * time.Second)
    defer ticker.Stop()
    
    for {
        <-ticker.C
        log.Printf("协程池状态 - 运行:%d 空闲:%d 容量:%d",
            pool.Running(), pool.Free(), pool.Cap())
    }
}()
```

## 总结

ants 是一个优秀的 Go 协程池库，特别适合需要大量并发处理的场景。通过合理使用 ants，可以：

- 🎯 **提升性能**: 减少 goroutine 创建销毁开销
- 🛡️ **控制资源**: 避免无限制创建 goroutine 导致的资源耗尽
- 📊 **便于监控**: 提供丰富的运行时信息
- 🔧 **灵活配置**: 适应不同的业务场景

在实际项目中，建议根据具体的业务需求和系统资源情况来选择合适的协程池配置。