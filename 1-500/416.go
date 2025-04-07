func canPartition(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%2 == 1 {
		return false
	}
	target := sum / 2
	dp := make([]bool, target+1)
	dp[0] = true
	for _, num := range nums {
		for i := target; i >= 1; i-- {
			if i-num >= 0 && dp[i-num] {
				dp[i] = true
				if i == target {
					return true
				}
			}
		}
	}

	return false
}
