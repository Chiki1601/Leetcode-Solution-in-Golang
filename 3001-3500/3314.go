func minBitwiseArray(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)

	for idx := 0; idx < n; idx++ {
		for num := 1; num <= nums[idx]; num++ {
			if num|(num+1) == nums[idx] {
				ans[idx] = num
				break
			}
		}
		if ans[idx] == 0 {
			ans[idx] = -1
		}
	}

	return ans
}s
