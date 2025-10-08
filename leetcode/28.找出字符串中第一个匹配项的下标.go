/*
 * @lc app=leetcode.cn id=28 lang=golang
 * @lcpr version=30203
 *
 * [28] 找出字符串中第一个匹配项的下标
 */

// @lc code=start

func strStr(haystack string, needle string) int {
	j := 0
	i := 0
	n := len(needle)
	next := getNext(needle)
	for i < len(haystack) {
		if needle[j] == haystack[i] {
			j++
			i++
		} else {
			if j > 0 {
				j = next[j-1]
			} else {
				i++
			}
		}

		if n == j {
			return i - n 
		}
	}
	return -1
}

func getNext(needle string) []int {
	n := len(needle)
	next := make([]int, n)
	pre_len := 0
	i := 1
	for i <= n-1 {
		if needle[pre_len] == needle[i] {
			pre_len++
			next[i] = pre_len
			i++
		} else {
			if pre_len == 0 {
				i++
			} else {
				pre_len = next[pre_len-1]
			}
		}
	}
	return next
}


// @lc code=end

/*
// @lcpr case=start
// "sadbutsad"\n"sad"\n
// @lcpr case=end

// @lcpr case=start
// "leetcode"\n"leeto"\n
// @lcpr case=end

*/

