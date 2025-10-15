/*
 * @lc app=leetcode.cn id=101 lang=golang
 * @lcpr version=30203
 *
 * [101] 对称二叉树
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
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    left := root.Left
    right := root.Right

    ans:= true

    ans = function(left,right)

    return ans
}

func function(left,right *TreeNode) bool{
    if (left==nil && right!=nil)||(left!=nil && right==nil) {
        return false
    }

    if left==nil && right==nil {
        return true
    }


    if left.Val!= right.Val{
        return false
    } 

    ans_l := function(left.Left,right.Right)
    ans_r := function(left.Right,right.Left)


    return ans_l &&ans_r
}
// @lc code=end



/*
// @lcpr case=start
// [1,2,2,3,4,4,3]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,2,null,3,null,3]\n
// @lcpr case=end

 */

