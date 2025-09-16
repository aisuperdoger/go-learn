package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	once     sync.Once
	instance *Database
)

type Database struct {
	Name string
}

// 初始化函数，只会执行一次
func initDatabase() {
	fmt.Println("正在初始化数据库...")
	time.Sleep(1 * time.Second) // 模拟初始化耗时
	instance = &Database{Name: "MyDB"}
	fmt.Println("数据库初始化完成")
}

// 获取数据库实例（线程安全）
func GetDatabase() *Database {
	once.Do(initDatabase) // 确保 initDatabase 只执行一次
	return instance
}

func main() {
	fmt.Println("sync.Once 简单示例")
	fmt.Println("==================")

	var wg sync.WaitGroup

	// 启动多个 goroutine 同时获取数据库实例
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d 开始获取数据库\n", id)
			db := GetDatabase()
			fmt.Printf("Goroutine %d 获取到数据库: %s\n", id, db.Name)
		}(i)
	}

	wg.Wait()
	fmt.Println("所有 goroutine 完成")
}
