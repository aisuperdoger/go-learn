/*
 * @lc app=leetcode.cn id=20 lang=golang
 * @lcpr version=30203
 *
 * [20] 有效的括号
 */

// @lc code=start

func isValid(s string) bool {
	stack := make([]int, len(s))
	top := -1

	for _, v := range s {
		n := getInt(v)
		if n%2 == 1 {
			top++
			stack[top] = n
		} else {
			if top >= 0 && stack[top] == n-1 {
				top--
			} else {
				return false
			}
		}
	}

	return top == -1

}

func getInt(r rune) int {
	switch r {
	case '[':
		return 1
	case ']':
		return 2
	case '(':
		return 3
	case ')':
		return 4
	case '{':
		return 5
	case '}':
		return 6
	default:
		return 0
	}
}

// func isValid(s string) bool {
//     stack := make([]int,len(s))
// 	top := -1

// 	for _,v :=range s {
// 		if v== ')'{
// 			if top >=0 && stack[top]== '('{
// 				top--
// 			}else{
// 				return false
// 			}
// 		}else if v== '}'{
// 			if top >=0 && stack[top]== '{'{
// 				top--
// 			}else{
// 				return false
// 			}
// 		}else if v== ']'{
// 			if top >=0 &&stack[top]== '['{
// 				top--
// 			}else{
// 				return false
// 			}
// 		}else{
// 			top++
// 			stack[top] = v
// 		}
// 	}

// 	return top==-1

// }

// @lc code=end

/*
// @lcpr case=start
// "()"\n
// @lcpr case=end

// @lcpr case=start
// "()[]{}"\n
// @lcpr case=end

// @lcpr case=start
// "(]"\n
// @lcpr case=end

// @lcpr case=start
// "([])"\n
// @lcpr case=end

// @lcpr case=start
// "([)]"\n
// @lcpr case=end

*/

