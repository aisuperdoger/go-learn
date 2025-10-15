/*
 * @lc app=leetcode.cn id=739 lang=golang
 * @lcpr version=30203
 *
 * [739] 每日温度
 */

// @lc code=start

func dailyTemperatures(temperatures []int) []int {
    stack :=list.New()

	ans := make([]int,len(temperatures))
	for i,v :=range temperatures {
		for stack.Len()!=0 && v > temperatures[stack.Front().Value.(int)] {
			ans[ stack.Front().Value.(int)] = i- stack.Front().Value.(int)
			stack.Remove(stack.Front())
		}
		stack.PushBack(i)
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [73,74,75,71,69,72,76,73]\n
// @lcpr case=end

// @lcpr case=start
// [30,40,50,60]\n
// @lcpr case=end

// @lcpr case=start
// [30,60,90]\n
// @lcpr case=end

*/

