/*
 * @lc app=leetcode.cn id=59 lang=golang
 * @lcpr version=30203
 *
 * [59] 螺旋矩阵 II
 */

// @lc code=start

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	x := 1
	count := 1
	for count <= int(math.Ceil(float64(n)/2.0))+1 {

		i := count - 1
		j := count - 1
		for ; j <= n-count-1; j++ {
			res[i][j] = x
			x++
		}

		for ; i <= n-count-1; i++ {
			res[i][j] = x
			x++
		}

		for ; j >= count; j-- {
			res[i][j] = x
			x++
		}

		for ; i >= count; i-- {
			res[i][j] = x
			x++
		}
		count++
	}

	if n%2 ==1 {
		res[n/2][n/2] = x
	}

	return res
}


// @lc code=end

/*
// @lcpr case=start
// 3\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

*/

