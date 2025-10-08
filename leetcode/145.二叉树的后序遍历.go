/*
 * @lc app=leetcode.cn id=145 lang=golang
 * @lcpr version=30203
 *
 * [145] 二叉树的后序遍历
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	traversal(root, &ans)
	return ans
}

func traversal(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}

	traversal(root.Left, ans)
	traversal(root.Right, ans)
	*ans = append(*ans, root.Val)
	return
}

// @lc code=end

/*
// @lcpr case=start
// [1,null,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5,null,8,null,null,6,7,9]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

*/

