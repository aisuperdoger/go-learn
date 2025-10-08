/*
 * @lc app=leetcode.cn id=349 lang=golang
 * @lcpr version=30203
 *
 * [349] 两个数组的交集
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]struct{})
	ans := make([]int,0)

	for _, v := range nums1 {
		set[v] = struct{}{}
	}

	for _, v := range nums2 {
		if _,ok := set[v];ok {
			delete(set,v)
			ans =append(ans,v)
		}
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,2,1]\n[2,2]\n
// @lcpr case=end

// @lcpr case=start
// [4,9,5]\n[9,4,9,8,4]\n
// @lcpr case=end

*/

