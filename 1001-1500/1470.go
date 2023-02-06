
func shuffle(nums []int, n int) []int {
	result := make([]int, n * 2)

	for i := 0; i < n; i++ {
		result[2 * i] = nums[i]
		result[2 * i + 1] = nums[i + n]
	}

	return result
}
