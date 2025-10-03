/*
 * @lc app=leetcode.cn id=206 lang=golang
 * @lcpr version=30203
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var l, r, next *ListNode
	l = nil
	if head == nil {
		return head
	} else {
		r = head
	}
	next = r.Next

	for r != nil {
		r.Next = l
		l = r
		r = next
		if next != nil {
			next = next.Next
		}
	}
	return l

}
// @lc code=end



/*
// @lcpr case=start
// [1,2,3,4,5]\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

 */

