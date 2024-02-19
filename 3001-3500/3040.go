func maxOperations(nums []int) int {
    n := len(nums)
    memo := make(map[string]int)

    var dfs func(int,int,int) int

    dfs = func(left,right,val int) int{
        if left >= right{
            return 0
        }

        pos := strconv.Itoa(left)+"-"+strconv.Itoa(right)+"-"+strconv.Itoa(val)
        if val,ok := memo[pos]; ok{
            return val
        }
        
        ans := 0
        if nums[left]+nums[right] == val{
            ans = max(ans,dfs(left+1,right-1,val)+1)
        }
        if nums[left]+nums[left+1] == val{
            ans = max(ans,dfs(left+2,right,val)+1)
        }
        if nums[right]+nums[right-1] == val{
            ans = max(ans,dfs(left,right-2,val)+1)
        }
        
        memo[pos] = ans
        return ans
    }
    
    answer := 0
    answer = max(answer,dfs(2,n-1,nums[0]+nums[1]))
    answer = max(answer,dfs(1,n-2,nums[0]+nums[n-1]))
    answer = max(answer,dfs(0,n-3,nums[n-2]+nums[n-1]))
    
    return answer+1
}
