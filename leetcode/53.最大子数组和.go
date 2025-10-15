/*
 * @lc app=leetcode.cn id=53 lang=golang
 * @lcpr version=30203
 *
 * [53] 最大子数组和
 */

// @lc code=start
func maxSubArray(nums []int) int {

	n := len(nums)
	dp := make([]int, n)
	for i:=range dp{
		dp[i] = nums[i]
	}
	ans := dp[0]

	for i := 1; i < n; i++ {
		dp[i] = max(dp[i],dp[i-1]+nums[i])
		ans = 	max(dp[i],ans)
	}
	return ans

}


func max( a,b int) int{
	if a>b {
		return a
	}else{
		return b
	}
}
// @lc code=end

/*
// @lcpr case=start
// [-2,1,-3,4,-1,2,1,-5,4]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

// @lcpr case=start
// [5,4,-1,7,8]\n
// @lcpr case=end

*/

