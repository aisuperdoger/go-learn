/*
 * @lc app=leetcode.cn id=98 lang=golang
 * @lcpr version=30203
 *
 * [98] 验证二叉搜索树
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
var cur int
var sign bool


func isValidBST(root *TreeNode) bool {
	cur = math.MinInt64
	sign = true

	function(root)

	return sign
}

func function(root *TreeNode) {
	if root==nil {return}
	function(root.Left)
	if cur >= root.Val {
		sign =false
		return
	}
	cur = root.Val
	function(root.Right)

}


// @lc code=end

/*
// @lcpr case=start
// [2,1,3]\n
// @lcpr case=end

// @lcpr case=start
// [5,1,4,null,null,3,6]\n
// @lcpr case=end

*/

