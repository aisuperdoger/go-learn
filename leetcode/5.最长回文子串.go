/*
 * @lc app=leetcode.cn id=5 lang=golang
 * @lcpr version=30203
 *
 * [5] 最长回文子串
 */

// @lc code=start

func longestPalindrome(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}

	// 创建DP表，dp[i][j]代表以i和j为开头和结尾的回文子串长度
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	ans := s[:1] // 初始化为第一个字符

	// 初始化：所有单个字符都是回文，长度为1
	// dp表达式dp[i][j] = dp[i+1][j-1] + 2，dp[i][j]没有考虑到i和j指向同一个位置时的情况，所以需要特殊处理
	// 并且dp[i+1][j-1]中两个指针向里缩，那么最终可能是i+1==j-1，初始时了，此才能让dp可以走下去
	for i := 0; i < n; i++ {
		dp[i][i] = 1

		// dp[i+1][j-1]中两个指针向里缩，那么最终可能是i+1>j-1，初始时了，此才能让dp可以走下去
		if i <= n-2 && s[i] == s[i+1] {
			dp[i][i+1] = 2
			if dp[i][i+1] > len(ans) {
				ans = s[i : i+2]
			}
		}

	}

	//  dp[i][j]依赖dp[i+1][j-1]，即
	// i依赖i+1，所以必须先知道结尾
	for i := n - 1; i >= 0; i-- {
		// j==i和j==i+1的情况前面考虑过了，这里不用也不能再遍历了
		// 如果遍历了元素j==i+1，那么dp[i][i+1]会被重新赋值为0
		for j := i + 2; j < n; j++ { 
			if s[i] == s[j] {
				if dp[i+1][j-1] == 0 {
					// 中间子串不是回文，当前子串也不是回文
					dp[i][j] = 0
				} else {
					// 中间是回文，当前也是回文，长度加2
					dp[i][j] = dp[i+1][j-1] + 2
				}

				// 更新最长回文子串
				if dp[i][j] > len(ans) {
					ans = s[i : j+1]
				}
			} else {
				dp[i][j] = 0
			}
		}
	}

	return ans
}

// @lc code=end

/*
// @lcpr case=start
// "babad"\n
// @lcpr case=end

// @lcpr case=start
// "cbbd"\n
// @lcpr case=end

*/

