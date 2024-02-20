func missingNumber(nums []int) int {
	n := len(nums)
	sum := (n + 1) * n / 2
	for _, x := range nums {
		sum -= x
	}
	return sum
}
