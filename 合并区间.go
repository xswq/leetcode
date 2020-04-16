package leetcode

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var ans [][]int
	var i int
	if len(intervals) == 1 {
		ans = append(ans, intervals[0])
	}
	for i < len(intervals)-1 {
		first := intervals[i][0]
		last := intervals[i][1]
		next := intervals[i+1]
		for last >= next[0] && i < len(intervals)-1 {
			last = max(last, next[1])
			i++
			if i == len(intervals)-1 {
				ans = append(ans, []int{first, last})
				break
			}
			next = intervals[i+1]
		}
		if last < next[0] {
			ans = append(ans, []int{first, last})
			if i == len(intervals)-2 {
				ans = append(ans, []int{next[0], next[1]})
			}
		} else if i == len(intervals)-2 {
			ans = append(ans, []int{first, max(last, next[1])})
		}
		i++
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
