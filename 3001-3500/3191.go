func minOperations(nums []int) int {
    count := 0
    n := len(nums)

    for i := 0; i < n-2; i++ {
        if nums[i] == 0 {
            nums[i] ^= 1
            nums[i+1] ^= 1
            nums[i+2] ^= 1
            count++
        }
    }

    if nums[n-2] == 1 && nums[n-1] == 1 {
        return count
    }
    return -1
}
