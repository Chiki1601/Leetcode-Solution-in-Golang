func largestPerimeter(nums []int) int {
    sort.Slice(nums, func(i, j int) bool {
        return nums[i] > nums[j]
    })
    
    fmt.Println(nums)
    
    for i := 0; i < len(nums)-2; i++ {
        if nums[i] < nums[i+1] + nums[i+2] {
            return nums[i] + nums[i+1] + nums[i+2]
        }
    }
    
    return 0
}
