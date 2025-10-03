/*
 * @lc app=leetcode.cn id=209 lang=golang
 * @lcpr version=30203
 *
 * [209] 长度最小的子数组
 */

// @lc code=start
func minSubArrayLen(target int, nums []int) int {
	// 思路：如果累加大于target，那么就移动左指针
	ans := len(nums) + 1
	tmp := 0
	j := 0
	for i, v := range nums {
		tmp += v
		if tmp >= target {
			for tmp >= target {
				tmp -= nums[j]
				j++
			}
			if (i - j + 2) < ans {
				ans = i - j + 2
			}
		}
	}

	if ans == len(nums)+1 {
		return 0
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// 7\n[2,3,1,2,4,3]\n
// @lcpr case=end

// @lcpr case=start
// 4\n[1,4,4]\n
// @lcpr case=end

// @lcpr case=start
// 11\n[1,1,1,1,1,1,1,1]\n
// @lcpr case=end

*/

