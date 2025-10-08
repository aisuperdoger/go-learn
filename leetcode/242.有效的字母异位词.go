/*
 * @lc app=leetcode.cn id=242 lang=golang
 * @lcpr version=30203
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}
	counts := [26]int{}
	n := len(s)
	for i := 0; i < n; i++ {
		counts[s[i]-'a']++
		counts[t[i]-'a']--

	}
	// 数组可以直接比较，切片和map都是引用类型不能直接比较
	return counts == [26]int{}
}


// map实现
// func isAnagram(s string, t string) bool {

// 	if len(s) != len(t) {
// 		return false
// 	}

// 	mm := make(map[rune]int, 26)

// 	n := len(s)
// 	for i := 0; i < n; i++ {
// 		mm[rune(s[i])]++
// 		mm[rune(t[i])]--
// 	}
// 	for _, v := range mm {
// 		if v != 0 {
// 			return false
// 		}
// 	}
// 	// 数组可以直接比较，切片和map都是引用类型不能直接比较
// 	return true
// }


// @lc code=end

/*
// @lcpr case=start
// "anagram"\n"nagaram"\n
// @lcpr case=end

// @lcpr case=start
// "rat"\n"car"\n
// @lcpr case=end

*/

