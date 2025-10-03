package main

import (
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

// SimpleJob 是一个简单的任务结构体，实现了 cron.Job 接口
type SimpleJob struct {
	Name    string
	Message string
	Counter int
}

// Run 实现 cron.Job 接口的 Run 方法
// 这是唯一需要实现的方法
func (j *SimpleJob) Run() {
	j.Counter++
	fmt.Printf("[%s] 任务 '%s' 第 %d 次执行: %s\n", 
		time.Now().Format("15:04:05"), 
		j.Name, 
		j.Counter, 
		j.Message)
}

// LogJob 记录日志的任务
type LogJob struct {
	LogLevel string
	Message  string
}

func (l *LogJob) Run() {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] [%s] %s\n", timestamp, l.LogLevel, l.Message)
}

// CounterJob 计数器任务
type CounterJob struct {
	Name  string
	count int
}

func (c *CounterJob) Run() {
	c.count++
	fmt.Printf("计数器 %s: %d\n", c.Name, c.count)
}

func main() {
	// 创建 cron 调度器
	c := cron.New(cron.WithSeconds())

	// 创建任务实例
	job1 := &SimpleJob{
		Name:    "问候任务",
		Message: "Hello, World!",
	}

	job2 := &LogJob{
		LogLevel: "INFO",
		Message:  "系统健康检查完成",
	}

	job3 := &CounterJob{
		Name: "访问计数",
	}

	// 添加任务到调度器
	// 每5秒执行一次 job1
	entryID1, err := c.AddJob("*/5 * * * * *", job1)
	if err != nil {
		log.Fatalf("添加任务1失败: %v", err)
	}
	fmt.Printf("任务1添加成功，ID: %d\n", entryID1)

	// 每10秒执行一次 job2
	entryID2, err := c.AddJob("*/10 * * * * *", job2)
	if err != nil {
		log.Fatalf("添加任务2失败: %v", err)
	}
	fmt.Printf("任务2添加成功，ID: %d\n", entryID2)



	// 启动调度器
	fmt.Println("\n启动调度器...")
	fmt.Println("程序将运行30秒后自动停止")
	fmt.Println("========================")
	
	c.Start()

	// c.Start()后面的AddJob任务是会被执行的，
	// 每3秒执行一次 job3
	entryID3, err := c.AddJob("*/3 * * * * *", job3)
	if err != nil {
		log.Fatalf("添加任务3失败: %v", err)
	}
	fmt.Printf("任务3添加成功，ID: %d\n", entryID3)

	// 运行30秒
	time.Sleep(30 * time.Second)

	// 停止调度器
	fmt.Println("\n停止调度器...")
	c.Stop()
	fmt.Println("程序结束")
}
