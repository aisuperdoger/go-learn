
# 总结
总结自：https://lxblog.com/qianwen/share?shareId=6e16021a-45a7-4364-ad02-8c2c8dace65b

M运行逻辑：
- 核心循环 (schedule())，不断地从“任务池”（P 的本地队列、全局队列、其他 P 的队列）里领取任务（G），当一个 G 的函数执行完毕，重新调用schedule()获取G
- 如果没有任务则自旋或睡眠等待被唤醒：当一个 P 的 G 处理完后，会从本地队列、全局队列获取，或其他 P 窃取（工作窃取）。如果长时间找不到 G，M 会进入自旋或休眠状态，释放P。在休眠后被唤醒时，它可能会获取到一个不同的 P。
- M的自旋和休眠
    - 如果运行时认为“系统可能很快就有新工作”，它会让这个 M 进入自旋状态，而不是立即休眠。
    - 休眠会从用户态陷入内核态，唤醒需要上下文切换的开销。
    - 自旋的 M 在自旋一段时间后，如果仍然无法获取到 P，会自动停止自旋并进入休眠状态（stopm）。
- 其他：
    - M不会被销毁只会睡眠，最多只有GOMAXPROCS个M同时运行。但是存在阻塞在系统调用或垃圾回收等任务的M，所以M总数一般超过GOMAXPROCS。
    - P 空闲但无 M 可用时，会创建M。
    - 阻塞 (非系统调用): 如果 G 因 channel 操作等阻塞，它会被标记为等待状态，M 继续从本地队列取其他 G。

G数据结构
- 保存的程序计数器（PC）、栈指针（SP）用于恢复协程等

P数据结构：
- P 拥有自己的本地运行队列 (runq)、内存分配缓存 (mcache)、定时器堆等资源。这些资源让绑定到它的 M 能够高效地执行 G。
- GOMAXPROCS 设置的是 P 的数量，初始化时就会创建GOMAXPROCS个P，并将第一个P与当前执行初始化的M绑定。


系统调用：
- M 上运行的 G 发起了一个阻塞式系统调用，则P 和 M 解绑，此时其他M获取这个空闲P，其他M执行 schedule()，获取G。
- 系统调用时：获取自旋的M处理这个空闲P，没有自旋的M，唤醒或新建M处理这个空闲的P。
- 系统调用结束后
    - 如果P没有其他M抢走，则 M立即重新绑定 P 
    - 获取其他空闲的P。


全局队列（Global Run Queue）是 Go 调度器中的一个重要组件，它存放的是处于就绪状态（_Grunnable）但暂时无法放入任何 P 的本地运行队列的 Goroutine。


**sysmon 监控程序：**
- **自旋M数量调整：** 如果发现系统负载高，npidle 经常大于 0，但 nmspinning 太低，它可能会间接促使更多的 M 进入自旋状态。
    - npidle：处于空闲状态的P，最大值为GOMAXPROCS
    - nmspinning：处于自旋状态的M
- **系统调用P状态改变：** 进入系统调用以后P的状态会变成_Psyscall ，此时是不会被其他M获取的。
    - 处于`_Psyscall`的P会被`sysmon` 线程释放成空闲 （_Pidle）：P在系统调用后变为`_Psyscall`状态，会被系统监控线程定期检查。当满足超时或工作需求条件时，P会被强制转为`_Pidle`状态，从而被其他M获取执行任务。
    - 如果系统调用很快结束，那么此时P可能还是`_Psyscall`状态，那么M会直接重新获取到此P。
- **触发GC**：检查是否需要强制触发垃圾回收（`forcegcperiod` 控制，默认2分钟）。
- **网络轮询**：定期检查网络事件（`netpoll`），唤醒阻塞的协程。
- **死锁检测**：通过 `checkdead` 函数检测是否所有协程都处于阻塞状态（死锁）。
---


# GMP模型

### **1. G（Goroutine，协程）**

G 是 Go 中的轻量级线程，数据结构定义在 `runtime/runtime2.go` 等文件中（部分字段）。核心字段包括：

```go
type g struct {
    // 状态相关
    status     uint32 // 状态（_Grunning/_Gwaiting/_Gdead 等）
    waitreason waitReason // 等待原因（如通道操作、IO 等待）

    // 栈信息
    stack      stack    // 协程栈（lo: 栈底地址, hi: 栈顶地址）
    stackguard0 uintptr // 栈保护标记（用于栈溢出检查）

    // 调度上下文
    sched      gobuf    // 调度恢复时的上下文（pc: 程序计数器, sp: 栈指针）

    // 关联关系
    m          *m       // 当前绑定的 M（可能为 nil）
    p          puintptr // 关联的 P（通过 M 间接关联）

    // 其他
    parentGoid int64    // 创建该协程的父协程 ID
    startpc    uintptr  // 协程入口函数地址
    // ... 其他字段（如定时器、同步组等）
}

```


---

### **2. M（Machine，操作系统线程）**

M 是实际运行的操作系统线程，数据结构定义在 `runtime/runtime2.go` 中。核心字段包括：

```go
type m struct {
    // 核心关联
    g0        *g       // 调度专用的协程（持有调度栈，用于执行调度逻辑）
    curg      *g       // 当前运行的业务协程（G）
    p         puintptr // 当前绑定的 P（逻辑处理器）
    oldp      puintptr // 系统调用前的 P（用于恢复）

    // 状态与控制
    spinning  bool     // 是否在自旋寻找可运行的 G（减少线程切换开销）
    blocked   bool     // 是否被阻塞（如等待系统调用）
    locks     int32    // 持有的锁数量（锁未释放时不能调度）

    // 其他
    id        int64    // 线程 ID（用于调试）
    mallocing int32    // 是否正在执行内存分配（防止并发分配）
    // ... 其他字段（如信号处理、线程本地存储等）
}

```


---

### **3. P（Processor，逻辑处理器）**

P 是逻辑处理器，负责管理 G 的运行队列和资源，数据结构定义在 `runtime/proc.go` 等文件中（部分字段）。核心字段包括：

```go
type p struct {
    // 运行队列
    runqhead uint32     // 本地运行队列头部
    runqtail uint32     // 本地运行队列尾部
    runq     [256]*g    // 本地可运行的 G 队列（环形缓冲区）
    runnext  *g         // 下一个优先运行的 G（用于抢占调度）

    // 状态与资源
    status   uint32     // 状态（_Pidle/_Prunning/_Psyscall 等）
    gFree    gQueue     // 空闲 G 的缓存（减少 G 对象的频繁分配）

    // 关联关系
    m        muintptr   // 当前绑定的 M（可能为 nil）
    // ... 其他字段（如系统调用计数、GC 相关标记等）
}

```

---

### **补充：全局调度器（schedt）**

调度器全局结构体 `schedt`（定义在 `runtime/runtime2.go`）负责协调 M、P、G 的全局状态，关键字段包括：

```go
type schedt struct {
    // 全局队列
    runq      gQueue    // 全局可运行的 G 队列（各 P 本地队列满时使用）
    runqsize  int32     // 全局队列长度

    // 空闲资源
    midle     muintptr  // 空闲 M 的链表
    nmidle    int32     // 空闲 M 的数量
    pidle     puintptr  // 空闲 P 的链表
    npidle    atomic.Int32 // 空闲 P 的数量

    // 其他控制
    goidgen   atomic.Uint64 // 协程 ID 生成器
    // ... 其他字段（如最大线程数、调试标记等）
}
```

### **总结**

- **G**：存储协程的状态、栈、调度上下文及关联的 M/P。
- **M**：代表操作系统线程，持有调度栈（g0）和当前运行的 G，通过 P 获取可运行的 G。
- **P**：管理本地运行队列和资源，是 M 与 G 之间的桥梁（M 必须绑定 P 才能运行 G）。
- **schedt**：全局调度器，管理空闲资源和全局队列，协调 M、P、G 的动态平衡。