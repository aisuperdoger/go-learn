/*
 * @lc app=leetcode.cn id=151 lang=golang
 * @lcpr version=30203
 *
 * [151] 反转字符串中的单词
 */

// @lc code=start
func reverseWords(s string) string {
	s = strings.Trim(s, " ")

	bs := []byte(s) // 要修改string需要先转换为[]byte
	slow := 0
	for i, v := range bs { // 这一步就是27. 移除元素的特殊处理
		if v != ' ' || (v == ' ' && i > 0 && bs[i-1] != ' ') {
			bs[slow] = bs[i]
			slow++
		}
	}

	// 只处理0到slow
	reverse(bs[0:slow])

	left := 0
	for i, v := range bs[0:slow] {
		if v == ' ' || i == len(bs[0:slow])-1 {
			if i == len(bs[0:slow])-1 {
				reverse(bs[left : i+1])
			} else {
				reverse(bs[left:i])
			}
			left = i + 1
		}
	}

	return string(bs[0:slow]) // 这种临时返回，编译器优化以后，是不会产生拷贝的
}

func reverse(bs []byte) {

	l, r := 0, len(bs)-1

	for r > l {
		bs[r], bs[l] = bs[l], bs[r]
		r--
		l++
	}

}

// @lc code=end

/*
// @lcpr case=start
// "the sky is blue"\n
// @lcpr case=end

// @lcpr case=start
// "  hello world  "\n
// @lcpr case=end

// @lcpr case=start
// "a good   example"\n
// @lcpr case=end

*/

