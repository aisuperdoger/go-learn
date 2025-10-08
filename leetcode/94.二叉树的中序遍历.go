/*
 * @lc app=leetcode.cn id=94 lang=golang
 * @lcpr version=30203
 *
 * [94] 二叉树的中序遍历
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
func inorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	traversal(root, &ans)
	return ans
}

func traversal(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}
	
	traversal(root.Left, ans)
	*ans = append(*ans, root.Val)
	traversal(root.Right, ans)
	return
}
// @lc code=end



/*
// @lcpr case=start
// [1,null,2,3]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

 */

