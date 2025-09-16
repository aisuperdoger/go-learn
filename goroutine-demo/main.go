package main

import (
	"fmt"
	"sync"
	"time"
)

// 示例1: 父协程结束，子协程也会结束
func example1() {
	fmt.Println("=== 示例1: 父协程结束，子协程被强制结束 ===")

	// 启动子协程
	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Printf("子协程工作中... %d\n", i)
			time.Sleep(500 * time.Millisecond)
		}
		fmt.Println("子协程完成工作") // 这行可能不会执行
	}()

	// 父协程只等待2秒就结束
	fmt.Println("父协程等待2秒...")
	time.Sleep(2 * time.Second)
	fmt.Println("父协程结束")
	// 程序退出，子协程被强制终止
}

// 示例2: 使用channel等待子协程完成
func example2() {
	fmt.Println("\n=== 示例2: 使用channel等待子协程完成 ===")

	done := make(chan bool)

	// 启动子协程
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Printf("子协程工作中... %d\n", i)
			time.Sleep(500 * time.Millisecond)
		}
		fmt.Println("子协程完成工作")
		done <- true // 通知父协程
	}()

	fmt.Println("父协程等待子协程完成...")
	<-done // 等待子协程完成
	fmt.Println("父协程收到完成信号，程序结束")
}

// 示例3: 多个子协程的情况
func example3() {
	fmt.Println("\n=== 示例3: 多个子协程，父协程提前结束 ===")

	// 启动多个子协程
	for i := 1; i <= 3; i++ {
		go func(id int) {
			for j := 1; j <= 5; j++ {
				fmt.Printf("子协程%d 工作中... %d\n", id, j)
				time.Sleep(300 * time.Millisecond)
			}
			fmt.Printf("子协程%d 完成工作\n", id) // 可能不会全部执行
		}(i)
	}

	// 父协程只等待1秒
	fmt.Println("父协程等待1秒...")
	time.Sleep(1 * time.Second)
	fmt.Println("父协程结束，所有子协程被终止")
}

// 示例4: 使用WaitGroup等待所有子协程
func example4() {
	fmt.Println("\n=== 示例4: 使用WaitGroup等待所有子协程 ===")

	var wg sync.WaitGroup

	// 启动多个子协程
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 1; j <= 3; j++ {
				fmt.Printf("子协程%d 工作中... %d\n", id, j)
				time.Sleep(300 * time.Millisecond)
			}
			fmt.Printf("子协程%d 完成工作\n", id)
		}(i)
	}

	fmt.Println("父协程等待所有子协程完成...")
	wg.Wait() // 等待所有子协程完成
	fmt.Println("所有子协程完成，程序结束")
}

func main() {
	// 运行示例1
	example1()

	time.Sleep(1 * time.Second) // 分隔符

	// 运行示例2
	example2()

	time.Sleep(1 * time.Second) // 分隔符

	// 运行示例3
	example3()

	time.Sleep(1 * time.Second) // 分隔符

	// 运行示例4
	example4()
}
