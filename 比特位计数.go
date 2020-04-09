package leetcode

func countBits(num int) []int {
	dp := make([]int, num+1)
	dp[0] = 0
	for i := 1; i <= num; i++ {
		if i%2 != 0 {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = dp[i/2]
		}
	}
	return dp
}
