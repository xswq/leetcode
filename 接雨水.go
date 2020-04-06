package leetcode

func trap(height []int) int {
	left := make([]int, len(height))
	right := make([]int, len(height))
	maxH := 0
	for i := 0; i < len(left); i++ {
		if height[i] > maxH {
			maxH = height[i]
		}
		left[i] = maxH
	}
	maxH = 0
	for i := len(right) - 1; i >= 0; i-- {
		if height[i] > maxH {
			maxH = height[i]
		}
		right[i] = maxH
	}
	ans := 0
	for i := range height {
		hold := min(left[i], right[i])
		if hold > height[i] {
			ans += hold - height[i]
		}
	}
	return ans
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
