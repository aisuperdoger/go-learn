/*
 * @lc app=leetcode.cn id=102 lang=golang
 * @lcpr version=30203
 *
 * [102] 二叉树的层序遍历
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
func levelOrder(root *TreeNode) [][]int {
    res := [][]int{}
    if root == nil{//防止为空
        return res
    }
    queue := list.New()
    queue.PushBack(root)

    var tmpArr []int

    for queue.Len() > 0 {
        length := queue.Len()               //保存当前层的长度，然后处理当前层（十分重要，防止添加下层元素影响判断层中元素的个数）
        for i := 0; i < length; i++ {
            node := queue.Remove(queue.Front()).(*TreeNode)    //出队列
            if node.Left != nil {
                queue.PushBack(node.Left)
            }
            if node.Right != nil {
                queue.PushBack(node.Right)
            }
            tmpArr = append(tmpArr, node.Val)    //将值加入本层切片中
        }
        res = append(res, tmpArr)          //放入结果集
        tmpArr = []int{}                  //清空层的数据
    }

    return res
}

// @lc code=end



/*
// @lcpr case=start
// [3,9,20,null,null,15,7]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

 */

