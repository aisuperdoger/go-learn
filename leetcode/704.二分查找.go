/*
 * @lc app=leetcode.cn id=704 lang=golang
 * @lcpr version=30203
 *
 * [704] 二分查找
 */

// @lc code=start
func search(nums []int, target int) int {
	n := len(nums)
    l,r := 0,n-1
	
	for r>=l {
		mid := l + (r-l)/2
		if nums[mid]==target {
			return mid
		}
		if nums[mid] <target {
			l = mid+1
		}else{
			r= mid-1
		}
	}
	return -1
}
// @lc code=end



/*
// @lcpr case=start
// [-1,0,3,5,9,12]\n9\n
// @lcpr case=end

// @lcpr case=start
// [-1,0,3,5,9,12]\n2\n
// @lcpr case=end

 */

