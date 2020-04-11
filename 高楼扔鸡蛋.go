package leetcode

func superEggDrop(K int, N int) int {
	dp := make([][]int, K+1)
	for i := range dp {
		dp[i] = make([]int, N+1)
	}
	m := 0
	for dp[K][m] < N {
		m++
		for i := 1; i <= K; i++ {
			dp[i][m] = dp[i][m-1] + dp[i-1][m-1] + 1
		}
	}
	return m
}
