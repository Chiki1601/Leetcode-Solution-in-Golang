func largestPerimeter(nums []int) int64 {
    sort.Ints(nums)
    var ans, preSum int64
    ans = -1
    for _, num := range nums {
        n := int64(num)
        if preSum > n {
            ans = preSum + n
        }
        preSum += n
    }
    return ans
}
