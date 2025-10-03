package main

import (
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

var (
	// 使用atomic操作的计数器
	atomicCounter int64

	// 不使用atomic操作的计数器
	regularCounter int64

	// 保护regularCounter的互斥锁
	mu sync.Mutex
)

// 使用atomic.AddInt64的安全递增函数
func incrementAtomicCounter() {
	atomic.AddInt64(&atomicCounter, 1)
}

// 不使用atomic的普通递增函数（存在竞态条件）
func incrementRegularCounterUnsafe() {
	regularCounter++ // 这里存在竞态条件！
}

// 使用互斥锁保护的递增函数
func incrementRegularCounterSafe() {
	mu.Lock()
	regularCounter++
	mu.Unlock()
}

// 读取atomic计数器的值
func getAtomicCounter() int64 {
	return atomic.LoadInt64(&atomicCounter)
}

// 读取普通计数器的值（需要加锁保护）
func getRegularCounter() int64 {
	mu.Lock()
	defer mu.Unlock()
	return regularCounter
}

// 重置所有计数器
func resetCounters() {
	atomic.StoreInt64(&atomicCounter, 0)
	mu.Lock()
	regularCounter = 0
	mu.Unlock()
}

func runComparison() {
	fmt.Println("=== atomic 原子操作对比示例 ===")
	fmt.Println("本示例演示 atomic.AddInt64 和 atomic.LoadInt64 与普通操作的区别")

	// 第一部分：atomic.AddInt64 对比
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("第一部分：atomic.AddInt64 写操作对比")
	fmt.Println(strings.Repeat("=", 50))
	testAtomicAddInt64()
}

// testAtomicAddInt64 测试 atomic.AddInt64 与普通写操作的区别
func testAtomicAddInt64() {
	fmt.Println("\n1. 竞态条件演示 - 写操作对比:")
	fmt.Println("多个goroutine同时写入计数器...")

	resetCounters()
	var wg sync.WaitGroup
	goroutineCount := 100
	incrementCount := 1000

	// 使用atomic.AddInt64的并发操作（安全）
	startTime := time.Now()
	for i := 0; i < goroutineCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < incrementCount; j++ {
				incrementAtomicCounter()
			}
		}()
	}
	wg.Wait()
	atomicDuration := time.Since(startTime)

	// 不使用atomic的并发操作（不安全，会产生竞态条件）
	var unsafeCounter int64
	startTime = time.Now()
	for i := 0; i < goroutineCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < incrementCount; j++ {
				unsafeCounter++ // 危险操作！
			}
		}()
	}
	wg.Wait()
	unsafeDuration := time.Since(startTime)

	// 使用互斥锁保护的操作（安全但较慢）
	var mutexCounter int64
	var testMu sync.Mutex
	startTime = time.Now()
	for i := 0; i < goroutineCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < incrementCount; j++ {
				testMu.Lock()
				mutexCounter++
				testMu.Unlock()
			}
		}()
	}
	wg.Wait()
	mutexDuration := time.Since(startTime)

	// 显示结果
	expectedValue := int64(goroutineCount * incrementCount)
	fmt.Printf("\n📊 写操作测试结果:\n")
	fmt.Printf("预期结果: %d\n", expectedValue)
	fmt.Printf("✅ atomic.AddInt64 结果: %d (耗时: %v)\n", getAtomicCounter(), atomicDuration)
	fmt.Printf("❌ 普通操作结果: %d (耗时: %v) ⚠️  存在竞态条件！\n", unsafeCounter, unsafeDuration)
	fmt.Printf("🔒 互斥锁保护结果: %d (耗时: %v)\n", mutexCounter, mutexDuration)

	// 性能对比
	fmt.Printf("\n📈 性能分析:\n")
	fmt.Printf("atomic.AddInt64 比互斥锁快 %.2fx\n", float64(mutexDuration)/float64(atomicDuration))
	fmt.Printf("atomic.AddInt64 比普通操作慢 %.2fx（但普通操作不安全）\n", float64(atomicDuration)/float64(unsafeDuration))

	// atomic.AddInt64 的其他用法
	fmt.Println("\n2. atomic.AddInt64 的灵活用法:")

	resetCounters()

	// 递增不同的值
	atomic.AddInt64(&atomicCounter, 5) // 增加5
	fmt.Printf("📝 增加5后: %d\n", getAtomicCounter())

	atomic.AddInt64(&atomicCounter, -3) // 减少3
	fmt.Printf("📝 减少3后: %d\n", getAtomicCounter())

	atomic.AddInt64(&atomicCounter, 10) // 再增加10
	fmt.Printf("📝 再增加10后: %d\n", getAtomicCounter())

	// 用法总结
	fmt.Println("\n💡 atomic.AddInt64 用法总结:")
	fmt.Println("   ✅ atomic.AddInt64(&counter, 1)   // 递增1")
	fmt.Println("   ✅ atomic.AddInt64(&counter, -1)  // 递减1")
	fmt.Println("   ✅ atomic.AddInt64(&counter, n)   // 增加任意值")
	fmt.Println("   ❌ counter++                      // 不安全，存在竞态条件")
}

func main() {
	runComparison()
}
