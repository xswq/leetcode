package leetcode

func numberOfSubArrays(nums []int, k int) int {
	ans := 0
	first, last := 0, 0
	count := 0
	for last < len(nums) {
		if nums[last]%2 != 0 {
			count++
		}
		last++
		if count == k {
			temp := last
			for last < len(nums) && nums[last]%2 == 0 {
				last++
			}
			right := last - temp
			left := 0
			for first < len(nums) && nums[first]%2 == 0 {
				left++
				first++
			}
			ans += (left + 1) * (right + 1)
			first++
			count--
		}
	}
	return ans
}
