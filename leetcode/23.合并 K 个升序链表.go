/*
 * @lc app=leetcode.cn id=23 lang=golang
 * @lcpr version=30203
 *
 * [23] 合并 K 个升序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


type ListHeap []*ListNode

func (h ListHeap) Len() int {
    return len(ListHeap)
}

func (h ListHeap) Swap(i,j int) {
  h[i],h[j] =   h[j],h[i] 
}

func (h ListHeap) Less(i, j int) bool {
    return h[i] < h[j]
}

func (h *ListHeap) Push(x any){
    *h = append(*h,x.(*ListNode))
}

func (h *ListHeap) Pop() any{
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0:n-1]
    return x
}


func mergeKLists(lists []*ListNode) *ListNode{
    h := &ListHeap{}
    heap.Init(h)
    
    for _,v := range lists {
        if v!=nil {
             heap.Push(h,v)
        }
    }
    
    dummy := &ListNode{}
    cur := dummy
    
    for h.Len() > 0 {
        tmp :=heap.Pop(h)
        cur.Next = tmp
        cur = cur.Next
        if tmp.Next!=nil{
            heap.Push(h,tmp.Next)
        }
    }
    
    return dummy.Next
}

// @lc code=end



/*
// @lcpr case=start
// [[1,4,5],[1,3,4],[2,6]]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

// @lcpr case=start
// [[]]\n
// @lcpr case=end

 */

