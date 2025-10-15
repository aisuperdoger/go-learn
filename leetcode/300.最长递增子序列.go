/*
 * @lc app=leetcode.cn id=300 lang=golang
 * @lcpr version=30203
 *
 * [300] 最长递增子序列
 */

// @lc code=start

func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	ans := 1

	for i := 1; i < n; i++ { // i := 1防止j := i - 1为零
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		ans = max(ans, dp[i])
	}

	return ans

}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end

/*
// @lcpr case=start
// [10,9,2,5,3,7,101,18]\n
// @lcpr case=end

// @lcpr case=start
// [0,1,0,3,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [7,7,7,7,7,7,7]\n
// @lcpr case=end

*/

