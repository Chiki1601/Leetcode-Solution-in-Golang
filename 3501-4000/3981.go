func interleaveCharacters(a string, b string, target string) int {
    const MOD int64 = 1e9 + 7
    n, n2, m := len(a), len(b), len(target)
    newTable := func() [][][2][2]int64 {
        tbl := make([][][2][2]int64, n+1)
        for i := range tbl {
            tbl[i] = make([][2][2]int64, n2+1)
        }
        return tbl
    }
    dp := newTable()
    dp[0][0][0][0] = 1

    for i := 1; i <= m; i++ {
        dpnxt := newTable()
        x := target[i-1]

        for prevb := 0; prevb <= n2; prevb++ {
            for t := 0; t < 2; t++ {
                pref0 := make([]int64, n+1)
                pref1 := make([]int64, n+1)
                pref0[0] = dp[0][prevb][0][t]
                pref1[0] = dp[0][prevb][1][t]
                for preva := 1; preva <= n; preva++ {
                    pref0[preva] = (pref0[preva-1] + dp[preva][prevb][0][t]) % MOD
                    pref1[preva] = (pref1[preva-1] + dp[preva][prevb][1][t]) % MOD
                }
                for preva := 1; preva <= n; preva++ {
                    if a[preva-1] != x {
                        continue
                    }
                    dpnxt[preva][prevb][1][t] =
                        (dpnxt[preva][prevb][1][t] + pref0[preva-1] + pref1[preva-1]) % MOD
                }
            }
        }

        for preva := 0; preva <= n; preva++ {
            for k := 0; k < 2; k++ {
                pref0 := make([]int64, n2+1)
                pref1 := make([]int64, n2+1)
                pref0[0] = dp[preva][0][k][0]
                pref1[0] = dp[preva][0][k][1]
                for prevb := 1; prevb <= n2; prevb++ {
                    pref0[prevb] = (pref0[prevb-1] + dp[preva][prevb][k][0]) % MOD
                    pref1[prevb] = (pref1[prevb-1] + dp[preva][prevb][k][1]) % MOD
                }
                for prevb := 1; prevb <= n2; prevb++ {
                    if b[prevb-1] != x {
                        continue
                    }
                    dpnxt[preva][prevb][k][1] =
                        (dpnxt[preva][prevb][k][1] + pref0[prevb-1] + pref1[prevb-1]) % MOD
                }
            }
        }

        dp = dpnxt
    }

    var ans int64 = 0
    for preva := 0; preva <= n; preva++ {
        for prevb := 0; prevb <= n2; prevb++ {
            ans = (ans + dp[preva][prevb][1][1]) % MOD
        }
    }
    return int(ans)
}
