### map

map操作不是原子的，这意味着多个协程同时操作map时有可能产生读写冲突，读写冲突会触发panic从而导致程序退出。

```go
package main

import (
	"fmt"
	"sync"
)

var m = make(map[int]int)
var wg sync.WaitGroup

func main() {
	// 增加计数，表示有2个goroutine要等待
	wg.Add(2)

	go func() {
		for i := 0; i < 1000; i++ {
			m[i] = i // 对map进行写入操作
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			m[i]++ // 对map进行修改操作
		}
		wg.Done()
	}()

	// 等待所有goroutine完成
	wg.Wait()

	fmt.Println("Finished:", m[123]) // 打印一个示例结果
}
```

运行上面代码有概率出现panic。原因：写入操作不仅仅包括简单的键值赋值，背后包含了一系列操作，还可能触发`map`的增长（即扩容），这意味着需要分配新的内存空间并迁移现有的键值对到新的位置。如果同时有多个goroutine尝试进行这样的操作，可能会导致数据结构的一致性被破坏。

为了避免这种情况，可以使用互斥锁（`sync.Mutex`或`sync.RWMutex`）来保护对`map`的操作，或者使用Go提供的线程安全的map替代方案如`sync.Map`。例如，使用互斥锁保护的代码如下所示：

```
go深色版本
var mu sync.Mutex

// 在每个涉及map操作的地方：
mu.Lock()
m[i] = i
mu.Unlock()

mu.Lock()
m[i]++
mu.Unlock()
```