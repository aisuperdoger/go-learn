package main

import (
	"fmt"
	"reflect"
)

// printValue 使用空接口接收任意类型的值并打印
func printValue(v interface{}) {
	fmt.Printf("值: %v, 类型: %T\n", v, v)
}

// printValueGeneric 使用泛型接收任意类型的值并打印（Go 1.18+）
func printValueGeneric[T any](v T) {
	fmt.Printf("泛型值: %v, 类型: %T\n", v, v)
}

// getTypeInfo 获取值的类型信息
func getTypeInfo(v interface{}) string {
	return fmt.Sprintf("类型: %T, 值: %v", v, v)
}

// processValue 根据类型处理不同的值
func processValue(v interface{}) {
	switch val := v.(type) {
	case int:
		fmt.Printf("这是一个整数: %d, 平方是: %d\n", val, val*val)
	case string:
		fmt.Printf("这是一个字符串: %s, 长度是: %d\n", val, len(val))
	case bool:
		fmt.Printf("这是一个布尔值: %t\n", val)
	case float64:
		fmt.Printf("这是一个浮点数: %.2f\n", val)
	default:
		fmt.Printf("未知类型: %T, 值: %v\n", val, val)
	}
}

// processValueGeneric 使用泛型和类型约束处理不同类型的值（Go 1.18+）
func processValueGeneric[T any](v T) {
	switch val := any(v).(type) {
	case int:
		fmt.Printf("泛型整数: %d, 平方是: %d\n", val, val*val)
	case string:
		fmt.Printf("泛型字符串: %s, 长度是: %d\n", val, len(val))
	case bool:
		fmt.Printf("泛型布尔值: %t\n", val)
	case float64:
		fmt.Printf("泛型浮点数: %.2f\n", val)
	default:
		fmt.Printf("泛型未知类型: %T, 值: %v\n", val, val)
	}
}

// Container 通用容器示例
type Container struct {
	data map[string]interface{}
}

// NewContainer 创建新的容器
func NewContainer() *Container {
	return &Container{
		data: make(map[string]interface{}),
	}
}

// Set 在容器中存储值
func (c *Container) Set(key string, value interface{}) {
	c.data[key] = value
}

// Get 从容器中获取值
func (c *Container) Get(key string) interface{} {
	return c.data[key]
}

// GetInt 获取整数值（类型安全的访问方式）
func (c *Container) GetInt(key string) (int, bool) {
	val, exists := c.data[key]
	if !exists {
		return 0, false
	}

	// 使用类型断言
	if intVal, ok := val.(int); ok {
		return intVal, true
	}
	return 0, false
}

// GetString 获取字符串值
func (c *Container) GetString(key string) (string, bool) {
	val, exists := c.data[key]
	if !exists {
		return "", false
	}

	if strVal, ok := val.(string); ok {
		return strVal, true
	}
	return "", false
}

// PrintAll 打印所有存储的值
func (c *Container) PrintAll() {
	fmt.Println("容器中的所有值:")
	for key, value := range c.data {
		fmt.Printf("  %s: %v (%T)\n", key, value, value)
	}
}

// calculateSum 计算数值切片的和
func calculateSum(numbers []interface{}) float64 {
	var sum float64
	for _, num := range numbers {
		// 使用反射处理不同数值类型
		switch v := reflect.ValueOf(num); v.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			sum += float64(v.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			sum += float64(v.Uint())
		case reflect.Float32, reflect.Float64:
			sum += v.Float()
		default:
			fmt.Printf("警告: 跳过非数值类型 %T\n", num)
		}
	}
	return sum
}

// calculateSumGeneric 使用泛型计算数值切片的和（Go 1.18+）
func calculateSumGeneric[T interface{ int | float64 }](numbers []T) T {
	var sum T
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func main() {
	fmt.Println("Go语言空接口(interface{})示例")
	fmt.Println("============================")

	// 1. 基本用法：接收任意类型
	fmt.Println("\n1. 基本用法 - 接收任意类型:")
	printValue(42)
	printValue("Hello, World!")
	printValue(true)
	printValue(3.14159)
	printValue([]int{1, 2, 3})

	// 2. 泛型用法：接收任意类型（Go 1.18+）
	fmt.Println("\n2. 泛型用法 - 接收任意类型:")
	printValueGeneric(42)
	printValueGeneric("Hello, Go!")
	printValueGeneric(true)
	printValueGeneric(3.14159)
	printValueGeneric([]int{1, 2, 3})

	// 3. 类型处理
	fmt.Println("\n3. 根据类型处理不同值:")
	values := []interface{}{42, "Go语言", true, 3.14, []int{1, 2, 3}}
	for _, v := range values {
		processValue(v)
	}

	// 4. 泛型类型处理
	fmt.Println("\n4. 泛型类型处理:")
	processValueGeneric(100)
	processValueGeneric("Go泛型")
	processValueGeneric(false)
	processValueGeneric(2.718)

	// 5. 通用容器示例
	fmt.Println("\n5. 通用容器示例:")
	container := NewContainer()
	container.Set("name", "张三")
	container.Set("age", 25)
	container.Set("active", true)
	container.Set("score", 95.5)

	container.PrintAll()

	// 类型安全的访问
	if name, ok := container.GetString("name"); ok {
		fmt.Printf("姓名: %s\n", name)
	}

	if age, ok := container.GetInt("age"); ok {
		fmt.Printf("年龄: %d\n", age)
	}

	// 6. 数值计算示例
	fmt.Println("\n6. 数值计算示例:")
	numbers := []interface{}{10, 20, 30, 40}
	sum := calculateSum(numbers)
	fmt.Printf("整数切片和: %.2f\n", sum)

	mixedNumbers := []interface{}{10, 20.5, 30, 40.8}
	sum = calculateSum(mixedNumbers)
	fmt.Printf("混合数值切片和: %.2f\n", sum)

	// 7. 泛型数值计算示例
	fmt.Println("\n7. 泛型数值计算示例:")
	intNumbers := []int{10, 20, 30, 40}
	intSum := calculateSumGeneric(intNumbers)
	fmt.Printf("整数切片和(泛型): %d\n", intSum)

	floatNumbers := []float64{10.5, 20.5, 30.5, 40.5}
	floatSum := calculateSumGeneric(floatNumbers)
	fmt.Printf("浮点数切片和(泛型): %.2f\n", floatSum)

	// 8. 类型断言示例
	fmt.Println("\n8. 类型断言示例:")
	var unknown interface{} = "这是一个字符串"

	// 类型断言
	if str, ok := unknown.(string); ok {
		fmt.Printf("成功断言为字符串: %s, 长度: %d\n", str, len(str))
	}

	// 类型断言失败的情况
	if num, ok := unknown.(int); ok {
		fmt.Printf("成功断言为整数: %d\n", num)
	} else {
		fmt.Println("断言为整数失败")
	}

	// 使用switch进行类型判断
	switch v := unknown.(type) {
	case int:
		fmt.Printf("这是一个整数: %d\n", v)
	case string:
		fmt.Printf("这是一个字符串: %s\n", v)
	case bool:
		fmt.Printf("这是一个布尔值: %t\n", v)
	default:
		fmt.Printf("未知类型: %T\n", v)
	}

	fmt.Println("\n程序结束")
}
