package leetcode

func canJump(nums []int) bool {
	canReach := 0
	for i := range nums {
		if i > canReach {
			return false
		}
		if i+nums[i] > canReach {
			canReach = i + nums[i]
		}
	}
	return true
}
