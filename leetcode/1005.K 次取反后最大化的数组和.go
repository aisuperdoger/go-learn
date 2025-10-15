/*
 * @lc app=leetcode.cn id=1005 lang=golang
 * @lcpr version=30203
 *
 * [1005] K 次取反后最大化的数组和
 */

// @lc code=start
func largestSumAfterKNegations(nums []int, k int) int {
	// 最高效的方法，一个for搞定
	ans := 0
	sort.Ints(nums)
	index := 0 // nums中最小值所在位置
	for i := range nums {
		if nums[i] < 0 && k != 0 {
			k--
			nums[i] = -nums[i]
		}
		if nums[i] <nums[index] {
			index = i
		}
		ans+=nums[i]
	}

	if k%2==1{
		ans -=2*nums[index]
	}

	return ans
}

// 简单的方法：1005.K次取反后最大化的数组和：先将负数取反，然后不断取反最小正数
// func largestSumAfterKNegations(nums []int, k int) int {
// 	ans := 0
// 	sort.Ints(nums)
// 	for i := range nums {
// 		if nums[i] < 0 && k != 0 {
// 			k--
// 			nums[i] = -nums[i]
// 		}
// 	}

// 	sort.Ints(nums)

// 	for k != 0 {
// 		k--
// 		nums[0] = -nums[0]
// 	}

// 	for _,v:= range nums {
// 		ans+=v
// 	}

// 	return ans
// }

// @lc code=end



/*
// @lcpr case=start
// [4,2,3]\n1\n
// @lcpr case=end

// @lcpr case=start
// [3,-1,0,2]\n3\n
// @lcpr case=end

// @lcpr case=start
// [2,-3,-1,5,-4]\n2\n
// @lcpr case=end

 */

