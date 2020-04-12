package leetcode

// 回溯算法
func permute(nums []int) [][]int {
	var path []int
	var ans [][]int
	backtrack(nums, 0, path, &ans)
	return ans
}

func backtrack(nums []int, pos int, path []int, ans *[][]int) {
	if pos == len(nums) {
		temp := make([]int, len(path))
		copy(temp, path)
		*ans = append(*ans, temp)
		return
	}

	for i := pos; i < len(nums); i++ {
		nums[i], nums[pos] = nums[pos], nums[i]
		path = append(path, nums[pos])
		backtrack(nums, pos+1, path, ans)
		path = path[:len(path)-1]
		nums[i], nums[pos] = nums[pos], nums[i]
	}
}
