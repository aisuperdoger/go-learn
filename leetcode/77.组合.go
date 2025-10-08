/*
 * @lc app=leetcode.cn id=77 lang=golang
 * @lcpr version=30203
 *
 * [77] 组合
 */

// @lc code=start
var ans [][]int
var path []int

func combine(n int, k int) [][]int {
	path, ans = make([]int, 0, k), make([][]int, 0) // 不知道为什么没有这行，无法通过所有测试用例
	track(n, k, 1)
	return ans
}


func track(n, k, i int) {
	if len(path) == k {
		tmp := make([]int,k)
		copy(tmp,path)
		ans = append(ans, tmp)
		return
	}
	// if i>n {return}

	for j := i; j <= n; j++ {
		path = append(path, j)
		track(n, k, j+1)
		path = path[:len(path)-1]
	}
}
// @lc code=end

/*
// @lcpr case=start
// 4\n2\n
// @lcpr case=end

// @lcpr case=start
// 1\n1\n
// @lcpr case=end

*/

