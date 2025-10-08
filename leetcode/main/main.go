package main

import (
	"fmt"
	"sort"
)

func main() {
	s := "fsdf"
	for i, v := range s {
		fmt.Printf("%T %T", s[i], v)
	}
	rune

	eraseOverlapIntervals([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}})
	return
}

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	cur := intervals[0][1]
	ans := 0

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < cur {
			ans++
			if intervals[i][1] < cur {
				cur = intervals[i][1]
			}
		} else {
			cur = intervals[i][1]
		}
	}
	return ans
}
