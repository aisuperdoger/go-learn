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
	for i,_ :=range dp{
		dp[i]=1
	}
	ans := 1
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				dp[i] =int(math.Max(float64( dp[j]+1), float64(dp[i])))
			}
		
		}
		ans = int(math.Max(float64(ans), float64(dp[i])))
	}
	return ans
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

