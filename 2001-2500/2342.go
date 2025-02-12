func maximumSum(nums []int) int {
	maxNums := [81]int{}
	result := -1
	var sum int
	for _, num := range nums {
		sum = 0
		for tmpNum := num; tmpNum != 0; tmpNum /= 10 {
			sum += tmpNum % 10
		}
		if maxNums[sum-1] != 0 {
			result = max(result, num+maxNums[sum-1])
		}
		maxNums[sum-1] = max(maxNums[sum-1], num)
	}
	return result
}
