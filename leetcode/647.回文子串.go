/*
 * @lc app=leetcode.cn id=647 lang=golang
 * @lcpr version=30203
 *
 * [647] 回文子串
 */

// @lc code=start
	func countSubstrings(s string) int {
		ans := 0

		n := len(s)

		dp := make([][]bool, n)
		for i := range dp {
			dp[i] = make([]bool, n)
		}

		for i := 0; i < n; i++ {
			ans++
			dp[i][i] = true
			if i <= n-2 && s[i] == s[i+1] {
				ans++
				dp[i][i+1] = true

			}
		}

		for i := n - 1; i >= 0; i-- {
			for j := i + 2; j < n; j++ {
				if s[i] == s[j] {
					if dp[i+1][j-1] ==false {
						dp[i][j] = false
					} else {
						dp[i][j] = true
						ans++
					}
				} else {
					dp[i][j] = false
				}
			}
		}

		return ans

	}

// @lc code=end

/*
// @lcpr case=start
// "abc"\n
// @lcpr case=end

// @lcpr case=start
// "aaa"\n
// @lcpr case=end

*/

