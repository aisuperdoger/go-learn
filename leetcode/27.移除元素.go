/*
 * @lc app=leetcode.cn id=27 lang=golang
 * @lcpr version=30203
 *
 * [27] 移除元素
 */
import "fmt"

// @lc code=start
func removeElement(nums []int, val int) int {
	s, f := 0, 0
	n := len(nums)
	for f < n {
		if nums[f] != val {
			nums[s] = nums[f]
			s++
			// fmt.Println(nums[s])
		}
		f++
	}
	fmt.Println(nums)
	return s
}

// @lc code=end

/*
// @lcpr case=start
// [3,2,2,3]\n3\n
// @lcpr case=end

// @lcpr case=start
// [0,1,2,2,3,0,4,2]\n2\n
// @lcpr case=end

*/

