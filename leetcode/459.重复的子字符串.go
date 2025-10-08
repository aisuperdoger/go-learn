/*
 * @lc app=leetcode.cn id=459 lang=golang
 * @lcpr version=30203
 *
 * [459] 重复的子字符串
 */

// @lc code=start
func repeatedSubstringPattern(s string) bool {
	// 最长相等前后缀不包含的子串就是最小重复子串。
	// 所以只要获取到kmp算法中的next[n-1]
	// 重复子串只能是s[0:(n-next[n-1])]
	n := len(s)
	next := make([]int, len(s))
	pre_len := 0
	i := 1
	for i <= n-1 {
		if s[i] == s[pre_len] {
			pre_len++
			// 包含元素i在内，后缀中最长的匹配次数
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

	ss := s[0 : n-next[n-1]]

	ssn := len(ss)
	if ssn == n { // 相等代表子串为原串
		return false
	}

	j := 0
	for i, _ := range s {
		j = i % ssn
		if s[i] != ss[j] {
			return false
		}
	}
	if j != ssn-1 { //代表没有遍历到最后
		return false
	}
	return true
}

// @lc code=end

/*
// @lcpr case=start
// "abab"\n
// @lcpr case=end

// @lcpr case=start
// "aba"\n
// @lcpr case=end

// @lcpr case=start
// "abcabcabcabc"\n
// @lcpr case=end

*/

