package main

import (
	"errors"
	"fmt"
	"log"
)

// 自定义错误类型
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("验证错误 - 字段: %s, 信息: %s", e.Field, e.Message)
}

// divide 安全除法函数，返回错误而不是触发panic
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为零")
	}
	return a / b, nil
}

// validateUser 验证用户信息
func validateUser(name, email string) error {
	if name == "" {
		return &ValidationError{
			Field:   "name",
			Message: "用户名不能为空",
		}
	}

	if email == "" {
		return &ValidationError{
			Field:   "email",
			Message: "邮箱不能为空",
		}
	}

	return nil
}

// riskyOperation 模拟可能出错的操作
func riskyOperation(shouldPanic bool) {
	if shouldPanic {
		panic("这是一个模拟的严重错误")
	}
	fmt.Println("操作成功完成")
}

// safeOperation 安全包装函数，使用recover捕获panic
func safeOperation(shouldPanic bool) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("捕获到panic: %v", r)
		}
	}()

	riskyOperation(shouldPanic)
	return nil
}

// readFile 模拟文件读取操作
func readFile(filename string) (string, error) {
	if filename == "" {
		return "", errors.New("文件名不能为空")
	}

	// 模拟文件不存在的情况
	if filename == "notfound.txt" {
		return "", fmt.Errorf("文件未找到: %s", filename)
	}

	return "文件内容", nil
}

func main() {
	fmt.Println("Go语言异常处理示例")
	fmt.Println("==================")

	// 1. 错误处理示例
	fmt.Println("\n1. 错误处理示例:")

	// 正常除法
	if result, err := divide(10, 2); err != nil {
		log.Printf("除法错误: %v", err)
	} else {
		fmt.Printf("10 / 2 = %.1f\n", result)
	}

	// 除零错误
	if result, err := divide(10, 0); err != nil {
		log.Printf("除法错误: %v", err)
	} else {
		fmt.Printf("10 / 0 = %.1f\n", result)
	}

	// 2. 自定义错误类型示例
	fmt.Println("\n2. 自定义错误类型示例:")

	if err := validateUser("", "test@example.com"); err != nil {
		// 类型断言检查是否为特定错误类型
		if validationErr, ok := err.(*ValidationError); ok {
			fmt.Printf("验证失败: %v\n", validationErr)
		} else {
			fmt.Printf("其他错误: %v\n", err)
		}
	}

	// 3. panic和recover示例
	fmt.Println("\n3. panic和recover示例:")

	// 不安全的操作会panic
	fmt.Println("不安全操作:")
	if err := safeOperation(true); err != nil {
		fmt.Printf("操作失败: %v\n", err)
	}

	// 安全的操作
	fmt.Println("安全操作:")
	if err := safeOperation(false); err != nil {
		fmt.Printf("操作失败: %v\n", err)
	} else {
		fmt.Println("操作成功")
	}

	// 4. 多层错误处理
	fmt.Println("\n4. 多层错误处理示例:")

	content, err := readFile("notfound.txt")
	if err != nil {
		fmt.Printf("读取文件失败: %v\n", err)

		// 可以进一步处理错误
		if errors.Is(err, fmt.Errorf("文件未找到: notfound.txt")) {
			fmt.Println("提示: 请检查文件路径是否正确")
		}
	} else {
		fmt.Printf("文件内容: %s\n", content)
	}

	// 读取正常文件
	content, err = readFile("example.txt")
	if err != nil {
		fmt.Printf("读取文件失败: %v\n", err)
	} else {
		fmt.Printf("文件内容: %s\n", content)
	}

	fmt.Println("\n程序正常结束")
}
