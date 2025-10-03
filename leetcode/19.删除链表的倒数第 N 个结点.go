/*
 * @lc app=leetcode.cn id=19 lang=golang
 * @lcpr version=30203
 *
 * [19] 删除链表的倒数第 N 个结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	l, r := dummy, dummy
	for i:=n; i >= 1; i-- {
		r = r.Next
	}

	for r.Next != nil {
		l = l.Next
		r = r.Next
	}

	l.Next = l.Next.Next
	return dummy.Next
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1]\n1\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n1\n
// @lcpr case=end

*/

