# atomic.AddInt64 vs 普通操作对比

本示例展示了使用 `atomic.AddInt64` 和不使用原子操作的区别，帮助理解原子操作在并发编程中的重要性。

## 运行示例

```bash
go run addint64_comparison.go
```

## 核心区别对比

### 1. 竞态条件 (Race Condition)

#### ❌ 普通操作 - 存在竞态条件
```go
var counter int64

func increment() {
    counter++  // 这不是原子操作！
    // 实际上包含三个步骤：
    // 1. 读取 counter 的值
    // 2. 将值加 1  
    // 3. 将结果写回 counter
}
```

**问题**: 多个 goroutine 同时执行时，可能会读取到相同的值，导致计数丢失。

#### ✅ 原子操作 - 线程安全
```go
var counter int64

func increment() {
    atomic.AddInt64(&counter, 1)  // 原子操作，线程安全
}
```

**优势**: 整个操作是原子的，不会被其他 goroutine 中断。

### 2. 性能对比

| 操作类型 | 性能 | 安全性 | 使用场景 |
|---------|------|--------|----------|
| `atomic.AddInt64` | 🚀 高性能 | ✅ 线程安全 | 简单的数值操作 |
| 普通操作 | 🚀 最快 | ❌ 不安全 | 单线程环境 |
| 互斥锁保护 | 🐌 较慢 | ✅ 线程安全 | 复杂的临界区 |

### 3. 使用场景

#### 适合使用 `atomic.AddInt64` 的场景：
- **计数器**: 请求计数、错误计数等
- **ID生成器**: 生成唯一递增ID
- **状态统计**: 并发统计数据
- **性能监控**: 实时性能指标

#### 不适合使用原子操作的场景：
- **复杂逻辑**: 需要多步操作的业务逻辑
- **条件判断**: 基于当前值做复杂判断
- **批量操作**: 需要保证多个变量的一致性

## 代码示例解析

### 竞态条件演示
```go
// 危险操作 - 存在竞态条件
func incrementRegularCounterUnsafe() {
    regularCounter++  // 多个goroutine可能同时读取相同值
}

// 安全操作 - 原子递增
func incrementAtomicCounter() {
    atomic.AddInt64(&atomicCounter, 1)  // 原子操作，线程安全
}
```

### 性能测试结果示例
```
预期结果: 100000
atomic.AddInt64 结果: 100000 (耗时: 15.2ms)
普通操作结果: 87234 (耗时: 12.8ms) ⚠️  存在竞态条件！
互斥锁保护结果: 100000 (耗时: 45.6ms)

atomic操作比互斥锁快 3.00x
```

## 最佳实践

### 1. 何时使用 atomic.AddInt64
- ✅ 简单的数值递增/递减
- ✅ 计数器类操作
- ✅ 高并发场景
- ✅ 性能敏感的代码

### 2. 何时避免使用
- ❌ 复杂的业务逻辑
- ❌ 需要条件判断的操作
- ❌ 多个变量需要同时更新

### 3. 代码规范
```go
// 好的实践
var counter int64

func increment() {
    atomic.AddInt64(&counter, 1)
}

func getCount() int64 {
    return atomic.LoadInt64(&counter)
}

// 避免的做法
var counter int64
var mu sync.Mutex

func increment() {
    mu.Lock()
    counter++  // 简单操作不需要互斥锁
    mu.Unlock()
}
```

## 常见错误

### 1. 混合使用原子操作和普通操作
```go
// ❌ 错误：混合使用
var counter int64

func badIncrement() {
    atomic.AddInt64(&counter, 1)  // 原子操作
}

func badRead() {
    return counter  // 普通读取 - 危险！
}

// ✅ 正确：统一使用原子操作
func goodRead() {
    return atomic.LoadInt64(&counter)  // 原子读取
}
```

### 2. 对指针使用原子操作
```go
// ❌ 错误
func badPointerOp() {
    var p *int64
    atomic.AddInt64(p, 1)  // p 为 nil，会panic
}

// ✅ 正确
func goodPointerOp() {
    var counter int64
    atomic.AddInt64(&counter, 1)  // 使用变量地址
}
```

## 总结

`atomic.AddInt64` 提供了一种高性能、线程安全的方式来进行简单的数值操作。相比互斥锁，它具有更好的性能；相比普通操作，它提供了并发安全性。在合适的场景下使用原子操作，可以写出既高效又安全的并发代码。