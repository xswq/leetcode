package leetcode

func subsets(nums []int) [][]int {
	var ans [][]int
	var path []int
	dfs(nums, path, 0, &ans)
	return ans
}

func dfs(nums []int, path []int, pos int, ans *[][]int) {
	if pos == len(nums) {
		temp := make([]int, len(path))
		copy(temp, path)
		*ans = append(*ans, temp)
		return
	}

	path = append(path, nums[pos])
	dfs(nums, path, pos+1, ans)

	path = path[:len(path)-1]
	dfs(nums, path, pos+1, ans)
}
