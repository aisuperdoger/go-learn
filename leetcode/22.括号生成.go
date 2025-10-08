/*
 * @lc app=leetcode.cn id=22 lang=golang
 * @lcpr version=30203
 *
 * [22] 括号生成
 */

// @lc code=start

func generateParenthesis(n int) []string {
	result := []string{}
	current := make([]byte, 0, n*2) // 预分配容量，优化性能
	backtrack(&result, current, 0, 0, n)
	return result
}

func backtrack(result *[]string, current []byte, open int, close int, n int) {
	if len(current) == n*2 {
		// 将字节切片转换为字符串并添加到结果中
		*result = append(*result, string(current))
		return
	}

	if open < n {
		// 添加左括号并递归
		backtrack(result, append(current, '('), open+1, close, n)
	}

	if close < open {
		// 添加右括号并递归
		backtrack(result, append(current, ')'), open, close+1, n)
	}
}

// 我的本方法
// var ans []string
// var path []byte
// var lcount, rcount int
// var l, r int

// func generateParenthesis(n int) []string {
// 	l, r = n, n
// 	n = 2 * n
// 	rcount,lcount = 0,0
// 	path = make([]byte, 0, n)
// 	ans = make([]string, 0)
// 	dfs(n)
// 	return ans
// }

// func dfs(n int) {
// 	if len(path) == n {
// 		tmp := make([]byte, n)
// 		copy(tmp, path)
// 		ans = append(ans, string(tmp))
// 		return
// 	}

// 	ls := []byte{'(', ')'}
// 	for _, v := range ls {
// 		if v == '(' && l != 0 {
// 			lcount++
// 			l--
// 		} else if v == ')' && lcount > rcount {
// 			rcount++
// 			r--
// 		} else {
// 			continue
// 		}
// 		path = append(path, v)

// 		dfs(n)
// 		path = path[:len(path)-1]
// 		if v == '(' {
// 			lcount--
// 			l++
// 		} else if v == ')' {
// 			r++
// 			rcount--
// 		}
// 	}
// }

// @lc code=end

/*
// @lcpr case=start
// 3\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

*/

