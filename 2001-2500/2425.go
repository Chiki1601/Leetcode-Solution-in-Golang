func xorAllNums(nums1 []int, nums2 []int) int {
    c1 := len(nums1)
    c2 := len(nums2)
    x1, x2 := 0, 0

    if c1%2 != 0 {
        for _, num := range nums2 {
            x2 ^= num
        }
    }
    if c2%2 != 0 {
        for _, num := range nums1 {
            x1 ^= num
        }
    }
    return x1 ^ x2
}
