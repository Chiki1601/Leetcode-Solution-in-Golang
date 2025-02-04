func maxAscendingSum(nums []int) int {
	sum := 0
	maximum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		maximum = Max(maximum, sum)

		if i+1 < len(nums) && nums[i] >= nums[i+1] {
			sum = 0
		}
	}

	return maximum
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
