package main

import "sort"

// 找出前k个出现次数最多的字符串
// {"i", "love", "leetcode", "i", "love", "coding"} 就是
// {"i", "love"}
// 出现频率相同的就按照字典序进行排列

func main() {
	inputs := []string{"i", "love", "leetcode", "i", "love", "coding"}
	k := 2

	function(inputs,k)

}
func function(inputs []string, k int)  []string {

	sort.Slice(inputs, func(i, j int) bool {

		return inputs[i] < inputs[j]
	})

	ss := make([]string,0)
	ii := make([][]int, 0)
	i := 0

	cur := inputs[0]
	count := 0
	j:=0
	for  {
		if j< len(inputs) &&  cur == inputs[j] {
			count++
			j++
		} else {
			ss = append(ss, cur)
			tmp := make([]int, 2)
			tmp[0] = count
			tmp[1] = i
			ii = append(ii,tmp)

			i++
			count = 0
			if j>= len(inputs) {
				break
				// 
			}
			cur = inputs[j]
		}
	}

	
	sort.Slice(ii, func(i, j int) bool {
		return ii[i][0] > ii[j][0]
	})

	ans := make([]string,0)
	// ans = append(ans,) 

	for j:=0;j<k ;j++{
		ans = append(ans,ss[ii[j][1]])
	}

	return ans

}
