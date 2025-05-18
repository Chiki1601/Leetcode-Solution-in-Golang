func colorTheGrid(m int, n int) int {
    const MOD = 1000000007
    var states []int
    var dfs func(pos, prev_color, mask int)
    dfs = func(pos, prev_color, mask int) {
        if pos == m {
            states = append(states, mask)
            return
        }
        for color := 0; color < 3; color++ {
            if color != prev_color {
                dfs(pos+1, color, mask*3+color)
            }
        }
    }
    dfs(0, -1, 0)
    S := len(states)
    compat := make([][]int, S)
    for i := 0; i < S; i++ {
        for j := 0; j < S; j++ {
            x, y := states[i], states[j]
            ok := true
            for k := 0; k < m; k++ {
                if x%3 == y%3 {
                    ok = false
                    break
                }
                x /= 3; y /= 3
            }
            if ok {
                compat[i] = append(compat[i], j)
            }
        }
    }
    dp := make([]int, S)
    for i := range dp {
        dp[i] = 1
    }
    var new_dp []int
    for t := 0; t < n-1; t++ {
        new_dp = make([]int, S)
        for i := 0; i < S; i++ {
            if dp[i] != 0 {
                for _, j := range compat[i] {
                    new_dp[j] = (new_dp[j] + dp[i]) % MOD
                }
            }
        }
        dp = new_dp
    }
    ans := 0
    for _, v := range dp {
        ans = (ans + v) % MOD
    }
    return ans
}
