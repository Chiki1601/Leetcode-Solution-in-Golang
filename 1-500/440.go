func findKthNumber(n int, k int) int {
    current := 1
    k--

    for k > 0 {
        count := countSteps(n, current, current+1)
        if count <= k {
            current++
            k -= count
        } else {
            current *= 10
            k--
        }
    }
    
    return current
}

func countSteps(n int, curr int, next int) int {
    steps := 0
    for curr <= n {
        steps += min(n+1, next) - curr
        curr *= 10
        next *= 10
    }
    return steps
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
