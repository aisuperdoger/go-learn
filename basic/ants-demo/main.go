package main

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/panjf2000/ants/v2"
)

func main() {
	fmt.Println("=== ants 协程池简单使用示例 ===")
	fmt.Println("ants 是一个高性能的 goroutine 池，用于管理和复用 goroutine")

	// 示例1: 基本协程池使用
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("示例1: 基本协程池使用")
	fmt.Println(strings.Repeat("=", 50))
	basicPoolExample()

	// 示例2: 带参数的协程池
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("示例2: 带参数的协程池")
	fmt.Println(strings.Repeat("=", 50))
	poolWithFuncExample()

	// 示例4: 协程池配置和监控
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("示例4: 协程池配置和监控")
	fmt.Println(strings.Repeat("=", 50))
	poolConfigExample()

	// 第三部分：综合示例
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("第三部分：综合并发安全性验证")
	fmt.Println(strings.Repeat("=", 50))
	testConcurrentSafety()

	// 提示用户可以运行高级示例
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("💡 想要查看更多高级用法示例？")
	fmt.Println("请运行: go run advanced_example.go")
	fmt.Println("或者参考 README.md 中的详细说明")
	fmt.Println(strings.Repeat("=", 50))
}

// basicPoolExample 演示基本的协程池使用
func basicPoolExample() {
	fmt.Println("📝 创建一个容量为10的协程池...")

	// 创建一个容量为10的协程池
	pool, err := ants.NewPool(10)
	if err != nil {
		log.Fatalf("创建协程池失败: %v", err)
	}
	defer pool.Release() // 确保协程池被释放

	var wg sync.WaitGroup
	taskCount := 20

	fmt.Printf("🚀 提交 %d 个任务到协程池...\n", taskCount)

	startTime := time.Now()
	for i := 0; i < taskCount; i++ {
		wg.Add(1)
		taskID := i // 避免闭包问题

		// 提交任务到协程池
		err := pool.Submit(func() {
			defer wg.Done()
			// 模拟工作
			fmt.Printf("   任务 %d 开始执行...\n", taskID)
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("   任务 %d 完成\n", taskID)
		})

		if err != nil {
			log.Printf("提交任务失败: %v", err)
			wg.Done()
		}
	}

	wg.Wait()
	duration := time.Since(startTime)

	fmt.Printf("✅ 所有任务完成，耗时: %v\n", duration)
	fmt.Printf("📊 协程池状态 - 正在运行: %d, 可用: %d\n",
		pool.Running(), pool.Free())
}

// poolWithFuncExample 演示带参数的协程池（PoolWithFunc）
func poolWithFuncExample() {
	fmt.Println("📝 创建一个带固定函数的协程池...")

	var wg sync.WaitGroup

	// 定义要执行的任务函数
	taskFunc := func(i interface{}) {
		defer wg.Done() // 在任务函数内部调用 Done()

		taskID := i.(int)
		fmt.Printf("   处理任务 %d...\n", taskID)

		// 模拟不同的工作负载
		workTime := time.Duration(50+taskID*10) * time.Millisecond
		time.Sleep(workTime)

		fmt.Printf("   任务 %d 完成 (耗时: %v)\n", taskID, workTime)
	}

	// 创建带函数的协程池
	poolWithFunc, err := ants.NewPoolWithFunc(5, taskFunc)
	if err != nil {
		log.Fatalf("创建协程池失败: %v", err)
	}
	defer poolWithFunc.Release()

	taskCount := 15
	fmt.Printf("🚀 向协程池提交 %d 个带参数的任务...\n", taskCount)

	startTime := time.Now()
	for i := 0; i < taskCount; i++ {
		wg.Add(1) // 在提交前增加计数

		// 提交带参数的任务
		err := poolWithFunc.Invoke(i)
		if err != nil {
			log.Printf("提交任务失败: %v", err)
			wg.Done() // 提交失败时手动调用 Done()
			continue
		}
	}

	wg.Wait() // 等待所有任务完成
	duration := time.Since(startTime)

	fmt.Printf("✅ 所有任务完成，耗时: %v\n", duration)
}

// poolConfigExample 演示协程池的配置和监控
func poolConfigExample() {
	fmt.Println("⚙️  协程池配置和监控示例...")

	// 创建带配置的协程池
	pool, err := ants.NewPool(
		20, // 协程池大小
		ants.WithOptions(ants.Options{
			ExpiryDuration:   3 * time.Second, // 空闲goroutine过期时间
			PreAlloc:         true,            // 预分配goroutine
			MaxBlockingTasks: 100,             // 最大阻塞任务数
			Nonblocking:      false,           // 允许阻塞
		}),
	)
	if err != nil {
		log.Fatalf("创建协程池失败: %v", err)
	}
	defer pool.Release()

	// 启动监控goroutine
	go func() {
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()

		for i := 0; i < 10; i++ {
			<-ticker.C
			fmt.Printf("📊 [监控] 运行中: %d, 空闲: %d, 容量: %d\n",
				pool.Running(), pool.Free(), pool.Cap())
		}
	}()

	// 分批提交任务
	var wg sync.WaitGroup
	batchCount := 3
	tasksPerBatch := 10

	for batch := 0; batch < batchCount; batch++ {
		fmt.Printf("\n🚀 提交第 %d 批任务 (%d 个)...\n", batch+1, tasksPerBatch)

		for i := 0; i < tasksPerBatch; i++ {
			wg.Add(1)
			taskID := batch*tasksPerBatch + i

			err := pool.Submit(func() {
				defer wg.Done()
				fmt.Printf("   执行任务 %d\n", taskID)
				time.Sleep(time.Duration(100+taskID*10) * time.Millisecond)
			})

			if err != nil {
				log.Printf("提交任务失败: %v", err)
				wg.Done()
			}
		}

		// 等待一段时间再提交下一批
		time.Sleep(1 * time.Second)
	}

	wg.Wait()

	// 最终状态
	fmt.Printf("\n✅ 所有任务完成\n")
	fmt.Printf("📊 最终状态 - 运行中: %d, 空闲: %d, 容量: %d\n",
		pool.Running(), pool.Free(), pool.Cap())

	// 等待一段时间让空闲的goroutine过期
	fmt.Println("⏰ 等待空闲goroutine过期...")
	time.Sleep(4 * time.Second)
	fmt.Printf("📊 过期后状态 - 运行中: %d, 空闲: %d, 容量: %d\n",
		pool.Running(), pool.Free(), pool.Cap())
}

// testConcurrentSafety 测试并发安全性
func testConcurrentSafety() {
	const (
		goroutines = 20
		operations = 1000
	)

	// 创建一个小的协程池用于测试
	pool, err := ants.NewPool(5)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Release()

	var wg sync.WaitGroup
	var counter int64

	fmt.Printf("🚀 启动 %d 个goroutine，每个执行 %d 次操作...\n", goroutines, operations)

	startTime := time.Now()

	// 使用协程池进行并发操作
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		goroutineID := i

		err := pool.Submit(func() {
			defer wg.Done()
			for j := 0; j < operations; j++ {
				// 模拟原子操作
				counter++
				time.Sleep(time.Microsecond)
			}
			fmt.Printf("   Goroutine %d 完成\n", goroutineID)
		})

		if err != nil {
			log.Printf("提交任务失败: %v", err)
			wg.Done()
		}
	}

	wg.Wait()
	duration := time.Since(startTime)

	expectedValue := int64(goroutines * operations)
	fmt.Printf("\n📊 综合测试结果:\n")
	fmt.Printf("预期值: %d\n", expectedValue)
	fmt.Printf("实际值: %d\n", counter)
	fmt.Printf("总耗时: %v\n", duration)

	if counter == expectedValue {
		fmt.Printf("✅ 并发安全测试通过！\n")
	} else {
		fmt.Printf("⚠️  注意：本示例中的counter++不是原子操作，在高并发下可能丢失更新\n")
		fmt.Printf("实际业务中应使用 sync/atomic 包或互斥锁保证并发安全\n")
	}

	fmt.Printf("\n💡 ants 协程池的价值:\n")
	fmt.Printf("• 控制并发数量：最多同时运行 5 个 goroutine\n")
	fmt.Printf("• 资源复用：减少 goroutine 创建和销毁开销\n")
	fmt.Printf("• 监控能力：可以实时查看协程池状态\n")
	fmt.Printf("• 内存稳定：避免无限制创建 goroutine 导致的内存问题\n")
}
