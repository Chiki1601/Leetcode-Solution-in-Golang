/*
Recap:
arr nums
find total way split arr
-  sums a[0] -> a[i] >= sum a[i+1] -> a[n-1]

Solution:
- calc sum for arr
- loop arr
- count case sum left >= sum-sum(left)
*/

func waysToSplitArray(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	result, sumLeft := 0, 0
	for i := 0; i < len(nums)-1; i++ {
		sumLeft += nums[i]
		if sumLeft >= (sum - sumLeft) {
			result++
		}
	}
	return result
}
