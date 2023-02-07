// length of the longest subarray with 2 different number

func totalFruit(tree []int) int {
    counts := make([]int, len(tree))
    unique := 0
    i := 0
    res := 0
    for j := 0; j < len(tree); j++ {
        counts[tree[j]]++
        if counts[tree[j]] == 1 {
            unique++
        }
        
        for unique > 2 {
            counts[tree[i]]--
            if counts[tree[i]] == 0 {
                unique--
            }
            i++
        }
        res = max(res, j-i+1)
    }
    
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
