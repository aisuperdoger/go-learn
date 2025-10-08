package main

import (
	"fmt"
	"reflect"
)



// Calculator 示例方法调用
type Calculator struct{}

func (c Calculator) Add(a, b int) int {
	return a + b
}

func (c Calculator) Multiply(a, b int) int {
	return a * b
}

func main2() {
	// 示例1: 基本类型反射
	fmt.Println("=== 基本类型反射 ===")
	num := 42
	showTypeAndValue(num)

	str := "Hello, Go!"
	showTypeAndValue(str)

	// 示例2: 结构体反射
	fmt.Println("\n=== 结构体反射 ===")
	person := Person{
		Name: "张三",
		Age:  25,
		City: "北京",
	}
	showStructInfo(person)

	// 示例3: 通过反射修改值
	fmt.Println("\n=== 通过反射修改值 ===")
	modifyValue()

	// 示例4: 反射调用方法
	fmt.Println("\n=== 反射调用方法 ===")
	callMethod()

	// 示例5: TypeOf 和 ValueOf 的区别
	fmt.Println("\n=== TypeOf 和 ValueOf 的区别 ===")
	TypeOfValueOfComparison()
}

// 显示变量的类型和值
func showTypeAndValue(v interface{}) {
	t := reflect.TypeOf(v)
	val := reflect.ValueOf(v)

	fmt.Printf("类型: %v, 值: %v\n", t, val)
}

// 显示结构体信息
func showStructInfo(v interface{}) {
	t := reflect.TypeOf(v)
	val := reflect.ValueOf(v)

	// 确保是结构体
	if t.Kind() != reflect.Struct {
		fmt.Println("不是结构体类型")
		return
	}

	fmt.Printf("结构体名称: %v\n", t.Name())
	fmt.Printf("字段数量: %v\n", t.NumField())

	// 遍历结构体字段
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := val.Field(i)
		fmt.Printf("字段名: %s, 类型: %v, 值: %v, Tag: %s\n",
			field.Name, field.Type, value, field.Tag)
	}
}

// 通过反射修改值
func modifyValue() {
	num := 10
	fmt.Printf("原始值: %d\n", num)

	// 获取num的反射值对象
	val := reflect.ValueOf(&num)
	// 确保是ptr类型并可以修改
	if val.Kind() == reflect.Ptr && val.Elem().CanSet() {
		// 获取指针指向的值
		valElem := val.Elem()
		// 修改值
		valElem.SetInt(20)
		fmt.Printf("修改后值: %d\n", num)
	} else {
		fmt.Println("无法修改值")
	}
}

// 反射调用方法
func callMethod() {
	calc := Calculator{}

	// 获取类型信息
	t := reflect.TypeOf(calc)
	v := reflect.ValueOf(calc)

	// 查找方法
	method, exists := t.MethodByName("Add")
	if exists {
		// 准备调用参数
		args := []reflect.Value{
			reflect.ValueOf(calc), // receiver
			reflect.ValueOf(5),    // a
			reflect.ValueOf(3),    // b
		}

		// 调用方法
		result := method.Func.Call(args)
		fmt.Printf("Add方法调用结果: %v\n", result[0].Int())
	}

	// 直接通过Value调用方法
	multiplyMethod := v.MethodByName("Multiply")
	if multiplyMethod.IsValid() {
		args := []reflect.Value{
			reflect.ValueOf(4), // a
			reflect.ValueOf(6), // b
		}

		result := multiplyMethod.Call(args)
		fmt.Printf("Multiply方法调用结果: %v\n", result[0].Int())
	}
}

// TypeOfValueOfComparison 简单对比 TypeOf 和 ValueOf 的区别
func TypeOfValueOfComparison() {
	fmt.Println("=== reflect.TypeOf 和 reflect.ValueOf 的核心区别 ===")

	// 示例数据
	var num int = 42
	var str string = "Hello"

	// reflect.TypeOf - 获取类型信息
	numType := reflect.TypeOf(num)
	strType := reflect.TypeOf(str)

	// reflect.ValueOf - 获取值信息
	numValue := reflect.ValueOf(num)
	strValue := reflect.ValueOf(str)

	fmt.Println("\n1. TypeOf (获取类型信息):")
	fmt.Printf("   int 类型信息: %v\n", numType)
	fmt.Printf("   string 类型信息: %v\n", strType)

	fmt.Println("\n2. ValueOf (获取值信息):")
	fmt.Printf("   int 值信息: %v\n", numValue)
	fmt.Printf("   string 值信息: %v\n", strValue)

	fmt.Println("\n3. TypeOf 提供的类型方法:")
	fmt.Printf("   int 类型名称: %s\n", numType.Name())
	fmt.Printf("   int 类型种类: %s\n", numType.Kind())
	fmt.Printf("   string 类型名称: %s\n", strType.Name())
	fmt.Printf("   string 类型种类: %s\n", strType.Kind())

	fmt.Println("\n4. ValueOf 提供的值方法:")
	fmt.Printf("   int 值: %d\n", numValue.Int())
	fmt.Printf("   string 值: %s\n", strValue.String())
	fmt.Printf("   int 值种类: %s\n", numValue.Kind())
	fmt.Printf("   string 值种类: %s\n", strValue.Kind())

	fmt.Println("\n5. 关键区别总结:")
	fmt.Println("   - TypeOf 返回 reflect.Type 接口，用于获取类型相关信息")
	fmt.Println("   - ValueOf 返回 reflect.Value 结构体，用于获取和操作值")
	fmt.Println("   - TypeOf 关注类型结构（字段、方法等）")
	fmt.Println("   - ValueOf 关注值的内容和操作（获取值、设置值、调用方法等)")
}
