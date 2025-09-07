func sumZero(n int) []int {
    ans := make([]int, n)
    
    for i := 1; i < n; i += 2 {
        if (i + 1) > n {
            ans[i - 1] = 0
        } else {
            var c = n - i
            ans[i] = -c
            ans[i - 1] = c
        }
    }
    
    return ans
} 
