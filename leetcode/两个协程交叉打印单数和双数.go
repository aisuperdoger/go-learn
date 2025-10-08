package leetcode


import (
	"fmt"
	"sync"
)

func main() {
	const object = 20
	ochan := make(chan int)
	schan := make(chan int)

	var wg sync.WaitGroup

	wg.Add(1) // 可以写到协程里吗？可能导致Add()和Wait()同时运行
	go func() {
		defer wg.Done()

		for i := 1; i <= object; i += 2 {
			<-schan
			fmt.Println(i)
			ochan <- 1
		}

	}()

	wg.Add(1) 
	go func() {
		defer wg.Done()

		for i := 0; i <= object; i += 2 {
			<-ochan
			fmt.Println(i)
			if i == 20 { // i==20时，不能再发送消息到schan，因为第一个协程中已经不接收信息了。
				break
			}
			schan <- 1
		}

	}()

	ochan <- 1
	wg.Wait()
}
