/*
 * @lc app=leetcode.cn id=707 lang=golang
 * @lcpr version=30203
 *
 * [707] 设计链表
 */

// @lc code=start
type MyLinkedList struct {
	Val  int
	Next *MyLinkedList
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		Val:  0,
		Next: nil,
	}
}

func (this *MyLinkedList) Get(index int) int {
	// ?
	// cur := &MyLinkedList{
	// 	val: 0,
	// 	next: nil
	// }

	for index != 0 && this.Next!=nil {
		this = this.Next
		index--
	}

	if this.Next!=nil {
		return this.Next.Val
	}

	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	item := &MyLinkedList{
		Val:  val,
		Next: this.Next,
	}
	this.Next = item
}

func (this *MyLinkedList) AddAtTail(val int) {
	for this.Next != nil {
		this = this.Next
	}

	item := &MyLinkedList{
		Val:  val,
		Next: nil,
	}

	this.Next = item
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	for i := 1; i <= index && this != nil; i++ {
		this = this.Next
	}

	if this == nil {
		return
	}

	item := &MyLinkedList{
		Val:  val,
		Next: this.Next,
	}

	this.Next = item
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	for i := 1; i <= index && this != nil; i++ {
		this = this.Next
	}

	if this==nil || this.Next == nil {
		return
	}

	this.Next = this.Next.Next
}



/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
// @lc code=end



