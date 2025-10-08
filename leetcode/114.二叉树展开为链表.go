/*
 * @lc app=leetcode.cn id=114 lang=golang
 * @lcpr version=30203
 *
 * [114] 二叉树展开为链表
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
func flatten(root *TreeNode) {
	// 因为要保持先序遍历，所以核心思路就是将右子树放在左子树的最右边
	cur := root

	for cur != nil {
		if cur.Left ==nil {
			cur = cur.Right
			continue
		}
		rt := cur.Left
		for rt.Right != nil {
			rt = rt.Right
		}
		rt.Right = cur.Right
		cur.Right = cur.Left
		cur.Left = nil
		cur = cur.Right
	}
	
}


// @lc code=end

/*
// @lcpr case=start
// [1,2,5,3,4,null,6]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

// @lcpr case=start
// [0]\n
// @lcpr case=end

*/

