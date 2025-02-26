func maxAbsoluteSum(nums []int) int {
    min_sum, max_sum, ans := 0, 0, 0
    for _, v := range nums {
        min_sum = min(v, min_sum + v)
        max_sum = max(v, max_sum + v)
        ans = max(ans, max(abs(min_sum), abs(max_sum)))
    }
    return ans
}

func abs(a int) int  {
    if a < 0 {
        return -a
    }
    return a
}
