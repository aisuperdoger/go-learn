// map操作不是原子的，这意味着多个协程同时操作map时有可能产生读写冲突，读写冲突会触发panic从而导致程序退出。
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