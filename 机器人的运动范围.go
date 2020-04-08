package leetcode

func movingCount(m int, n int, k int) int {
	if m < 0 || n < 0 || k < 0 {
		return 0
	}
	visited := make([]int, m*n)
	return dfs(m, n, k, 0, 0, visited)
}

func dfs(m, n, k, row, col int, visited []int) int {
	ans := 0
	if check(m, n, k, row, col, visited) {
		visited[row*n+col] = 1
		ans = 1 + dfs(m, n, k, row-1, col, visited) + dfs(m, n, k, row+1, col, visited) + dfs(m, n, k, row, col-1, visited) + dfs(m, n, k, row, col+1, visited)
	}
	return ans
}

func check(m, n, k, row, col int, visited []int) bool {
	if row >= 0 && row < m && col >= 0 && col <= n && visited[row*n+col] == 0 && digitalSum(row)+digitalSum(col) <= k {
		return true
	}
	return false
}

func digitalSum(n int) int {
	ans := 0
	for n > 0 {
		ans += n % 10
		n = n / 10
	}
	return ans
}
