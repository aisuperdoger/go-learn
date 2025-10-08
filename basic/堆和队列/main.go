package main

import (
	"container/heap"
	"container/list"
	"fmt"
)

// 1. 定义一个类型（这里使用int类型的切片）
type IntHeap []int

// 2. 实现sort.Interface（Len, Less, Swap）
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // 最小堆
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// 3. 实现heap.Interface的Push和Pop方法
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	h := &IntHeap{2, 1, 5}
	heap.Init(h) // 初始化堆

	heap.Push(h, 3)                  // 插入新元素
	fmt.Printf("最小值: %d\n", (*h)[0]) // 查看最小元素

	// 依次弹出所有元素（从小到大）
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h)) // 输出: 1 2 3 5
	}
}

func listTest() {
	l := list.New() // 创建一个新链表

	// 添加元素
	e1 := l.PushBack(1) // 在尾部添加：1
	l.PushFront(0)      // 在头部添加：0
	e2 := l.PushBack(2) // 在尾部添加：2

	// 在指定元素前后插入
	l.InsertBefore(3, e1) // 在元素1前插入3
	l.InsertAfter(4, e2)  // 在元素2后插入4

	// 遍历链表 (从头到尾)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ") // 输出: 0 3 1 2 4
	}

	// 移除元素
	a := l.Remove(e1).(int)
	fmt.Printf("%T\n", a)

	if l.Len() != 0 {
		x := l.Front().Value.(int)
		fmt.Printf("%T\n", x)
	}

	l.Remove(l.Front())
}
