package main

import (
	"fmt"
	"sort"
)

// Person 用于演示结构体排序
type Person struct {
	Name string
	Age  int
}

func function(str string) (one, zero int) {
	one,zero =0,0

	for _,v :=range str {
		if v== '0' {zero++}else{one++}
	}
	return 
}

func function(str string) (one, zero int) {
	
	return 1,1 
}
// sort.Slice()
func slice() {
	Person := &Person{
		Name:"sfdfs",
	}
	nums := []int{5, 2, 8, 1, 9, 3}

	// 使用 sort.Slice() 进行降序排序
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j] // 左边一个比右边大
	})
	fmt.Println("降序排序后:", nums)

	// 2. 对字符串切片进行排序
	fmt.Println("\n=== 字符串排序 ===")
	names := []string{"Charlie", "Alice", "Bob", "David"}
	fmt.Println("排序前:", names)

	sort.Slice(names, func(i, j int) bool {
		return names[i] < names[j]
	})
	fmt.Println("按字母顺序排序后:", names)

	// 3. 对结构体切片进行排序
	fmt.Println("\n=== 结构体排序 ===")
	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
		{"David", 20},
	}

	fmt.Println("排序前:")
	for _, p := range people {
		fmt.Printf("  %s (%d岁)\n", p.Name, p.Age)
	}

	// 按年龄升序排序
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})

	fmt.Println("按年龄升序排序后:")
	for _, p := range people {
		fmt.Printf("  %s (%d岁)\n", p.Name, p.Age)
	}

	// 按姓名排序
	sort.Slice(people, func(i, j int) bool {
		return people[i].Name < people[j].Name
	})

	fmt.Println("按姓名排序后:")
	for _, p := range people {
		fmt.Printf("  %s (%d岁)\n", p.Name, p.Age)
	}
}

func main() {

	f := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	sort.Sort(sort.Reverse(sort.Float64Slice(f)))
	fmt.Println(f)
	// sort.Reverse(sort.Float64s(f))为错误
	// sort.Float64s(f) 是一个排序操作，不是返回一个 sort.Interface 对象，而是直接排序并返回 ()（无返回值）。
	// sort.Float64Slice(f) 将 f 转换为可排序接口，可用于 sort.Sort 和 sort.Reverse

	sort.Sort(sort.Float64Slice(f))
	fmt.Println(f)

	sort.Float64s(f)
	fmt.Println(f)


	n := []int{1,3,23,4,4,5,5}
	sort.Ints(n)
	fmt.Println(n)

}
func findMaxForm(strs []string, m int, n int) int {

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for _, v := range strs {
		one, zero := function(v)
		for i := m; i >=zero zero; i-- {
			for j := n; j >= one; j-- {
				dp[i][j] = max(dp[i-zero][j-one]+1,dp[i][j])
			}
		}
	}

	return dp[m][n]
}

func function(str string) (one, zero int) {
	one,zero =0,0

	for _,v :=range str {
		if v== '0' {
			zero++
		}else{
			one++
		}
	}
	return 
}