/*
 * @lc app=leetcode.cn id=1 lang=golang
 * @lcpr version=30203
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	tmp := make(map[int]int, len(nums))

	for i, v := range nums {
		tmp[target-v] = i
	}

	for i, v := range nums {
		if vv, ok := tmp[v]; ok && vv != i {
			return []int{vv, i}
		}
	}

	return  []int{} 

}

// 最优解
// func twoSum(nums []int, target int) []int {
//     m := make(map[int]int)
//     for i, num := range nums {
//         complement := target - num
//         if index, found := m[complement]; found {
//             return []int{index, i}
//         }
//         m[num] = i  // 只需要一个循环，是因为即使当前错过了，后面也会找到的
//     }
//     return nil  // 返回空数组 nil 代替空切片
// }

// @lc code=end

/*
// @lcpr case=start
// [2,7,11,15]\n9\n
// @lcpr case=end

// @lcpr case=start
// [3,2,4]\n6\n
// @lcpr case=end

// @lcpr case=start
// [3,3]\n6\n
// @lcpr case=end

*/

