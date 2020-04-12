package leetcode

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	var path []int
	var ans [][]int
	sort.Ints(candidates)
	backtrack(candidates, 0, target, path, &ans)
	return ans
}

func backtrack(candidates []int, pos int, target int, path []int, ans *[][]int) {
	if target == 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		*ans = append(*ans, temp)
		return
	}

	for i := pos; i < len(candidates); i++ {
		if target < candidates[i] {
			break
		}
		path = append(path, candidates[i])
		backtrack(candidates, i, target-candidates[i], path, ans)
		path = path[:len(path)-1]
	}
}
