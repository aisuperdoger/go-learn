/*
 * @lc app=leetcode.cn id=474 lang=golang
 * @lcpr version=30203
 *
 * [474] 一和零
 */

// @lc code=start
func findMaxForm(strs []string, m int, n int) int {

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for _, v := range strs {
		one, zero := function(v)
		for i := m; i >=zero ; i-- {
			for j := n; j >= one; j-- {
				dp[i][j] = max(dp[i-zero][j-one]+1,dp[i][j])
			}
		}
	}

	return dp[m][n]
}

func function(str string) (one, zero int) {
	one,zero =0,0

	for _,v :=range str {
		if v== '0' {
			zero++
		}else{
			one++
		}
	}
	return 
}

// @lc code=end

/*
// @lcpr case=start
// ["10", "0001", "111001", "1", "0"]\n5\n3\n
// @lcpr case=end

// @lcpr case=start
// ["10", "0", "1"]\n1\n1\n
// @lcpr case=end

*/

