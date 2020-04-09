package leetcode

func hammingDistance(x int, y int) int {
	temp := x ^ y
	ans := 0
	for temp > 0 {
		if temp&1 == 1 {
			ans++
		}
		temp >>= 1
	}
	return ans
}
