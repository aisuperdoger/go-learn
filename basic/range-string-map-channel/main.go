package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. range操作字符串
	fmt.Println("=== range操作字符串 ===")
	str := "Hello 世界"
	fmt.Printf("字符串: %s\n", str)

	// 使用range遍历字符串，正确处理Unicode字符
	for i, char := range str {
		fmt.Printf("字节索引: %d, 字符: %c, Unicode码点: %U\n", i, str[i], char)
	}

	// 对比使用传统for循环遍历字符串（按字节）
	fmt.Println("\n使用传统for循环按字节遍历:")
	for i := 0; i < len(str); i++ {
		fmt.Printf("字节索引: %d, 字节: %c\n", i, str[i])
	}

	// 2. range操作映射（map）
	fmt.Println("\n=== range操作映射（map） ===")
	// 创建一个map
	students := map[string]int{
		"张三": 85,
		"李四": 92,
		"王五": 78,
		"赵六": 96,
	}

	fmt.Println("学生成绩:")
	// 使用range遍历map
	for name, score := range students {
		fmt.Printf("姓名: %s, 成绩: %d\n", name, score)
	}

	// 只遍历键
	fmt.Println("\n只遍历学生姓名:")
	for name := range students {
		fmt.Printf("姓名: %s\n", name)
	}

	// 只遍历值
	fmt.Println("\n只遍历成绩:")
	for _, score := range students {
		fmt.Printf("成绩: %d\n", score)
	}

	// 注意：map的遍历顺序是不确定的
	fmt.Println("\n多次遍历map观察顺序:")
	for i := 0; i < 3; i++ {
		fmt.Printf("第%d次遍历: ", i+1)
		for name := range students {
			fmt.Printf("%s ", name)
		}
		fmt.Println()
	}

	// 3. range操作通道
	fmt.Println("\n=== range操作通道 ===")
	// 创建一个带缓冲的通道
	ch := make(chan string, 5)

	// 启动一个goroutine向通道发送数据
	go func() {
		defer close(ch) // 发送完数据后关闭通道
		for i := 1; i <= 5; i++ {
			ch <- fmt.Sprintf("消息%d", i)
			time.Sleep(100 * time.Millisecond) // 模拟处理时间
		}
	}()

	// 使用range从通道接收数据，直到通道关闭
	fmt.Println("接收通道中的消息:")
	for msg := range ch {
		fmt.Printf("接收到: %s\n", msg)
	}

	fmt.Println("通道已关闭，range循环结束")

	// 4. range操作通道的实际应用示例：生产者-消费者模式
	fmt.Println("\n=== 生产者-消费者模式示例 ===")
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	// 启动3个worker
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 发送5个任务
	go func() {
		defer close(jobs)
		for j := 1; j <= 5; j++ {
			jobs <- j
			fmt.Printf("发送任务 %d\n", j)
		}
	}()

	// 收集结果
	for a := 1; a <= 5; a++ {
		result := <-results
		fmt.Printf("收到结果: %d\n", result)
	}

	fmt.Println("所有任务完成")
}

// worker函数处理任务
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d 开始处理任务 %d\n", id, j)
		time.Sleep(time.Second) // 模拟工作时间
		fmt.Printf("Worker %d 完成任务 %d\n", id, j)
		results <- j * 2 // 返回结果
	}
}
