/*
 * @lc app=leetcode.cn id=24 lang=golang
 * @lcpr version=30203
 *
 * [24] 两两交换链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	var l,r,next,pre *ListNode
	dummy := &ListNode{
		Next:head,
	}

	pre =dummy
    if head!=nil && head.Next!=nil {
		l =head
		r = head.Next
		next = r.Next
	}else{
		return head
	}

	for l!=nil && r!=nil {
		l.Next = next
		r.Next = l
		pre.Next = r

		pre = l
		l = next
		if l!=nil && l.Next !=nil {
			r =  l.Next
		}else{
			break
		}
		next = r.Next
	}

	return dummy.Next
}
// @lc code=end



/*
// @lcpr case=start
// [1,2,3,4]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

 */

