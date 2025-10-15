/*
 * @lc app=leetcode.cn id=718 lang=golang
 * @lcpr version=30203
 *
 * [718] 最长重复子数组
 */

// @lc code=start
func findLength(nums1 []int, nums2 []int) int {
	n1 := len(nums1)
	n2 := len(nums2)
	dp := make([][]int, n1+1) // 必须要使用n+1的dp，否则会使dp初始化很复杂
	for i := 0; i <= n1; i++ {
		dp[i] = make([]int, n2+1)
	}

	ans := 0
	
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				ans = max(dp[i][j],ans)
			} else {
				dp[i][j] = 0
			}
		}
	}

	return ans
}

func max(a,b int) int {
	if a>b {
		return a
	}else{
		return b
	}
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,2,1]\n[3,2,1,4,7]\n
// @lcpr case=end

// @lcpr case=start
// [0,0,0,0,0]\n[0,0,0,0,0]\n
// @lcpr case=end

*/

