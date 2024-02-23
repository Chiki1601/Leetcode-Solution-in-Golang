func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
    dp, tmp := make([][]int, K + 2), 1 << 31
    for i := range dp { dp[i] = make([]int, n) }
    for i := 0; i < K + 2; i++ { for j := 0; j < n; j++ { dp[i][j] = tmp } }
    dp[0][src] = 0
    for i := 1; i < K + 2; i++ {
        dp[i][src] = 0
        for _, flight := range flights {
            dp[i][flight[1]] = min(dp[i][flight[1]], dp[i - 1][flight[0]] + flight[2])
        }
    }
    if dp[K + 1][dst] < tmp { return dp[K + 1][dst] }
    return -1
}

func min(a, b int) int {
    if a < b { return a }
    return b
}
