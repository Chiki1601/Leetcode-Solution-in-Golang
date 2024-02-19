func maxOperations(nums []int) int {
    n := len(nums)
    val := nums[0] + nums[1]
    answer := 1
    for i:=3;i<n;i+=2{
        if nums[i]+nums[i-1] == val{
            answer++
        }else{
            break   
        }   
    }
    return answer
}
