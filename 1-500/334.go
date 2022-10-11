func increasingTriplet(nums []int) bool {
    low := nums[0]
    mid := math.MaxInt32
    for _, v := range nums {
        if v > mid { return true }
        if v < low { low = v }
        if v > low && v < mid { mid = v }
    }
    return false
}
