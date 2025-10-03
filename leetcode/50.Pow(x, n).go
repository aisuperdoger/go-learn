/*
 * @lc app=leetcode.cn id=50 lang=golang
 * @lcpr version=30203
 *
 * [50] Pow(x, n)
 */

// @lc code=start
func myPow(x float64, n int) float64 {
	if n <0 {
		n = -n
		x =1/x
	}

	ans :=1.0
	for n!=0 {
		if n%2 ==1 {
			ans *= x
		}
		x*=x
		n = n/2
	}

	
	return ans
}
// @lc code=end



/*
// @lcpr case=start
// 2.00000\n10\n
// @lcpr case=end

// @lcpr case=start
// 2.10000\n3\n
// @lcpr case=end

// @lcpr case=start
// 2.00000\n-2\n
// @lcpr case=end

 */

