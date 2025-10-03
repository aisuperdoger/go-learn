/*
 * @lc app=leetcode.cn id=977 lang=golang
 * @lcpr version=30203
 *
 * [977] 有序数组的平方
 */

// @lc code=start
func sortedSquares(nums []int) []int {
	n := len(nums)
	l := 0
	r := n - 1
	ans := make([]int, n)

	for r >= l {
		if nums[r]*nums[r] > nums[l]*nums[l] {
			ans[n-1] = nums[r] * nums[r]
			r--
		} else {
			ans[n-1] = nums[l] * nums[l]
			l++
		}
		n--
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [-4,-1,0,3,10]\n
// @lcpr case=end

// @lcpr case=start
// [-7,-3,2,3,11]\n
// @lcpr case=end

*/

