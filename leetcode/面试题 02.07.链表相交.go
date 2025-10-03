/*
 * @lc app=leetcode.cn id=面试题 02.07 lang=golang
 * @lcpr version=30203
 *
 * [面试题 02.07] 链表相交
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA, curB := headA, headB

	Alen := 0
	for curA != nil {
		curA = curA.Next
		Alen++
	}
	Blen := 0
	for curB != nil {
		curB = curB.Next
		Blen++
	}

	if Blen > Alen {
		curA = headB
		curB = headA
	}else {
		curA = headA
		curB = headB
	}

	diff := int(math.Abs(float64(Blen) - float64(Alen)))
	for i := diff; i >= 1; i-- {
		curA = curA.Next
	}

	for curA !=nil {
		if curA==curB {
			return curA
		}
		curA = curA.Next
		curB = curB.Next
	}

	return nil
}

// @lc code=end

/*
// @lcpr case=start
// 8\n[4,1,8,4,5]\n[5,0,1,8,4,5]\n2\n3\n
// @lcpr case=end

// @lcpr case=start
// 2\n[0,9,1,2,4]\n[3,2,4]\n3\n1\n
// @lcpr case=end

// @lcpr case=start
// 0\n[2,6,4]\n[1,5]\n3\n2\n
// @lcpr case=end

*/

