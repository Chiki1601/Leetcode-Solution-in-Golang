func maxIncreasingSubarrays(nums []int) int {
    n, up, preUp, res := len(nums), 1, 0, 0
    for i := 1; i < n; i++ {
        if nums[i] > nums[i-1] {
            up++
        } else {
            preUp = up
            up = 1
        }
        half := up >> 1
        m := preUp
        if up < preUp {
            m = up
        }
        candidate := half
        if m > half {
            candidate = m
        }
        if candidate > res {
            res = candidate
        }
    }
    return res
}
