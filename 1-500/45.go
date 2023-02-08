func jump(nums []int) int {
    ans, n, curEnd, curFar := 0, len(nums), 0, 0
    for i := 0; i < n-1; i++ {
        curFar = max(curFar, i + nums[i])
        if i == curEnd {
            ans++
            curEnd = curFar
        }
    }
    return ans
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
