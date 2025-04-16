func countGood(nums []int, k int) int64 {
    mp := make(map[int]int)
    l, r := 0, 0
    count := 0
    var ans int64 = 0
    for r < len(nums) {
        count += mp[nums[r]]
        mp[nums[r]]++
        for l < len(nums) && (count - (mp[nums[l]] - 1)) >= k {
            mp[nums[l]]--
            count -= mp[nums[l]]
            l++
        }
        if count >= k {
            ans += int64(l + 1)
        }
        r++
    }
    return ans
}
