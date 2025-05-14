func lengthAfterTransformations(s string, t int, nums []int) int {
    const MOD = 1000000007
    var freq [26]int64
    for _, ch := range s {
        freq[ch-'a']++
    }
    var baseM [26][26]int64
    for i := 0; i < 26; i++ {
        for k := 1; k <= nums[i]; k++ {
            j := (i + k) % 26
            baseM[i][j]++
        }
    }
    multiply := func(A, B [26][26]int64) [26][26]int64 {
        var C [26][26]int64
        for i := 0; i < 26; i++ {
            for k := 0; k < 26; k++ {
                aik := A[i][k]
                if aik == 0 {
                    continue
                }
                for j := 0; j < 26; j++ {
                    C[i][j] = (C[i][j] + aik*B[k][j]) % MOD
                }
            }
        }
        return C
    }
    matrixPower := func(M [26][26]int64, exp int) [26][26]int64 {
        var res [26][26]int64
        for i := 0; i < 26; i++ {
            res[i][i] = 1
        }
        base := M
        for exp > 0 {
            if exp&1 == 1 {
                res = multiply(res, base)
            }
            base = multiply(base, base)
            exp >>= 1
        }
        return res
    }
    mt := matrixPower(baseM, t)
    var ans int64
    for i := 0; i < 26; i++ {
        fi := freq[i]
        if fi == 0 {
            continue
        }
        for j := 0; j < 26; j++ {
            ans = (ans + fi*mt[i][j]) % MOD
        }
    }
    return int(ans)
}
