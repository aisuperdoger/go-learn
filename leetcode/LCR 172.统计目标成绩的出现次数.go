/*
 * @lc app=leetcode.cn id=LCR 172 lang=golang
 * @lcpr version=30203
 *
 * [LCR 172] 统计目标成绩的出现次数
 */



// @lc code=start
func countTarget(scores []int, target int) int {
	l := find(scores,target,true)
	r := find(scores,target,false)

	return r-l
}


func find(scores []int, target int, left bool) int {
	l,r := 0,len(scores)-1

	for r>=l {
		mid := l +(r-l)/2
		if target == scores[mid] {
			if left {
				r= mid-1
			}else{
				l=mid+1
			}
			continue
		}

		if  target > scores[mid] {
			l  = mid+1
		}else{
			r = mid- 1
		}
	}

	return l
}

// @lc code=end



/*
// @lcpr case=start
// [2, 2, 3, 4, 4, 4, 5, 6, 6, 8]\n4\n
// @lcpr case=end

// @lcpr case=start
// [1, 2, 3, 5, 7, 9]\n6\n
// @lcpr case=end

 */

