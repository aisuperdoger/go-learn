package main

import (
	"fmt"
	"runtime/debug"
)

// level3 第三层函数
func level3() {
	fmt.Println("=== 在 level3 函数中打印调用栈 ===")
	stack := debug.Stack()
	fmt.Printf("调用栈:\n%s\n", stack)
}

// level2 第二层函数
func level2() {
	fmt.Println("进入 level2 函数")
	level3()
}

// level1 第一层函数
func level1() {
	fmt.Println("进入 level1 函数")
	level2()
}

// panicExample 演示panic时的调用栈
func panicExample() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("捕获到panic: %v\n", r)
			fmt.Println("=== panic时的调用栈 ===")
			stack := debug.Stack()
			fmt.Printf("%s\n", stack)
		}
	}()

	fmt.Println("即将触发panic...")
	panic("这是一个测试panic")
}

// deepCall 深层调用示例
func deepCall(depth int) {
	if depth <= 0 {
		fmt.Println("=== 深层调用的调用栈 ===")
		stack := debug.Stack()
		fmt.Printf("%s\n", stack)
		return
	}
	deepCall(depth - 1)
}

// logWithStack 带调用栈的日志函数
func logWithStack(message string) {
	fmt.Printf("日志: %s\n", message)
	fmt.Println("=== 调用位置 ===")
	stack := debug.Stack()
	fmt.Printf("%s\n", stack)
}

// businessLogic 业务逻辑函数
func businessLogic() {
	fmt.Println("执行业务逻辑...")
	logWithStack("业务逻辑执行完成")
}

func main() {
	fmt.Println("debug.Stack() 使用示例")
	fmt.Println("========================")

	// 示例1: 基本调用栈
	fmt.Println("\n1. 基本调用栈:")
	level1()

	fmt.Println("\n==================================================")

	// 示例2: panic时的调用栈
	fmt.Println("\n2. panic时的调用栈:")
	panicExample()

	fmt.Println("\n==================================================")

	// 示例3: 深层递归调用栈
	fmt.Println("\n3. 深层递归调用栈:")
	deepCall(3)

	fmt.Println("\n==================================================")

	// 示例4: 日志中的调用栈
	fmt.Println("\n4. 日志中的调用栈:")
	businessLogic()
}
