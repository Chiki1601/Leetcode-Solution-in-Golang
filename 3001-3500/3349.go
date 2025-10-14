// Complexity:
// Time O(N) and Space O(1) where N is the length of nums.
func hasIncreasingSubarrays(nums []int, k int) bool {
	currLen := 1
	isFirst := true

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			currLen++
		} else if currLen < k {
			currLen = 1
			isFirst = true
		} else if isFirst && currLen < 2*k {
			currLen = 1
			isFirst = false
		} else {
			return true
		}
	}
	return currLen >= 2*k || (!isFirst && currLen >= k)
}
