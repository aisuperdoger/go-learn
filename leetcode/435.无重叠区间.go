/*
 * @lc app=leetcode.cn id=435 lang=golang
 * @lcpr version=30203
 *
 * [435] 无重叠区间
 */

// @lc code=start
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals,func(i,j int) bool{
		return intervals[i][0] <  intervals[j][0]
	})

	cur := intervals[0][1]
	ans := 0
	
	for i:=1;i<len(intervals);i++  {
		if intervals[i][0] < cur {
			ans ++
			if intervals[i][1] < cur {
				cur =intervals[i][1]
			}
		}else {
			cur =intervals[i][1]
		}
	}
	return ans
}

// @lc code=end



/*
// @lcpr case=start
// [[1,2],[2,3],[3,4],[1,3]]\n
// @lcpr case=end

// @lcpr case=start
// [ [1,2], [1,2], [1,2] ]\n
// @lcpr case=end

// @lcpr case=start
// [ [1,2], [2,3] ]\n
// @lcpr case=end

 */

