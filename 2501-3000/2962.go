func countSubarrays(nums []int, k int) int64 {
    max := 0
    var left int64 = 0
    var res int64 = 0
    count := 0

    for _, val := range nums {
        if max < val {
            max = val;
        }
    }

    for _, val := range nums {
        if val == max {
            count++
        }

        for count >= k {
            if (nums[left] == max) {
                count--
            } 
            left++
        }

        res += left
    }

    return res
}
