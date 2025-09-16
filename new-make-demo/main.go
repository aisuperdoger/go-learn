package main

import (
	"fmt"
)

// Person 示例结构体
type Person struct {
	Name string
	Age  int
}

// Company 示例结构体
type Company struct {
	Name      string
	Employees []*Person
}

func main() {
	fmt.Println("Go语言中new和make的区别")
	fmt.Println("========================")

	// 1. new的使用示例
	fmt.Println("\n1. new的使用示例:")

	// 使用new创建基本类型
	intPtr := new(int)
	fmt.Printf("new(int) 返回指针: %v, 值: %d\n", intPtr, *intPtr)

	// 使用new创建结构体
	personPtr := new(Person)
	fmt.Printf("new(Person) 返回指针: %v, 值: %+v\n", personPtr, *personPtr)

	// 修改new创建的值
	*intPtr = 42
	personPtr.Name = "张三"
	personPtr.Age = 30
	fmt.Printf("修改后 - int值: %d, Person值: %+v\n", *intPtr, *personPtr)

	// 2. make的使用示例
	fmt.Println("\n2. make的使用示例:")

	// 使用make创建slice
	slice := make([]int, 5)
	fmt.Printf("make([]int, 5): %+v, 长度: %d, 容量: %d\n", slice, len(slice), cap(slice))

	// 使用make创建slice并指定容量
	sliceWithCap := make([]int, 3, 10)
	fmt.Printf("make([]int, 3, 10): %+v, 长度: %d, 容量: %d\n", sliceWithCap, len(sliceWithCap), cap(sliceWithCap))

	// 使用make创建map
	mapExample := make(map[string]int)
	mapExample["apple"] = 5
	mapExample["banana"] = 3
	fmt.Printf("make(map[string]int): %+v\n", mapExample)

	// 使用make创建channel
	ch := make(chan string, 2)
	ch <- "消息1"
	ch <- "消息2"
	close(ch)

	// 从channel读取数据
	for msg := range ch {
		fmt.Printf("从channel读取: %s\n", msg)
	}

	// 3. new和make在slice上的区别
	fmt.Println("\n3. new和make在slice上的区别:")

	// 使用make创建slice
	makeSlice := make([]int, 3)
	makeSlice[0] = 1
	fmt.Printf("make([]int, 3): %+v\n", makeSlice)

	// 使用new创建slice（注意区别）
	newSlicePtr := new([]int)
	fmt.Printf("new([]int) 返回指针: %v, 指向的值: %+v, 长度: %d\n", newSlicePtr, *newSlicePtr, len(*newSlicePtr))

	// 4. new和make在map上的区别
	fmt.Println("\n4. new和make在map上的区别:")

	// 使用make创建map
	makeMap := make(map[string]int)
	makeMap["key1"] = 100
	fmt.Printf("make(map[string]int): %+v\n", makeMap)

	// 使用new创建map指针
	newMapPtr := new(map[string]int)
	fmt.Printf("new(map[string]int) 返回指针: %v, 指向的值: %+v\n", newMapPtr, *newMapPtr)
	// 注意：newMapPtr指向的是nil map，不能直接赋值
	// (*newMapPtr)["key1"] = 100 // 这会引发panic

	// 5. 实际应用示例
	fmt.Println("\n5. 实际应用示例:")

	// 创建公司结构体
	company := &Company{
		Name:      "科技公司",
		Employees: make([]*Person, 0),
	}

	// 添加员工
	employee1 := &Person{Name: "张三", Age: 30}
	employee2 := &Person{Name: "李四", Age: 25}

	company.Employees = append(company.Employees, employee1, employee2)

	fmt.Printf("公司信息: %+v\n", company)
	fmt.Printf("员工数量: %d\n", len(company.Employees))
	for i, emp := range company.Employees {
		fmt.Printf("  员工%d: %+v\n", i+1, *emp)
	}

	fmt.Println("\n程序结束")
}
