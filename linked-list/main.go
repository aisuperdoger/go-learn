package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
    Val  int
    Next *ListNode
}

// 移除链表中所有值为 val 的节点
func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}

	cur := dummy

	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}

	}

	return dummy.Next
}

// 辅助函数：创建链表（方便测试）
func createList(vals []int) *ListNode {
    if len(vals) == 0 {
        return nil
    }
    head := &ListNode{Val: vals[0]}
    cur := head
    for i := 1; i < len(vals); i++ {
        cur.Next = &ListNode{Val: vals[i]}
        cur = cur.Next
    }
    return head
}

// 辅助函数：打印链表
func printList(head *ListNode) {
    cur := head
    for cur != nil {
        fmt.Printf("%d", cur.Val)
        if cur.Next != nil {
            fmt.Print(" -> ")
        }
        cur = cur.Next
    }
    fmt.Println()
}

// 测试函数
func main() {
    arr := []int{1, 2, 6, 3, 4, 5, 6}
	// head1 := createList(arr)


	dummy := &ListNode{}
	cur := dummy
	for _,v := range arr {
		cur.Next = &ListNode{
			Val:v,
		}
		cur = cur.Next
	}

   	removeElements(dummy.Next, 6)
	printList(dummy.Next)
}