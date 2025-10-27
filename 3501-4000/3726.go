func removeZeros(n int64) int64 {
    var res int64 = 0
    var mult int64 = 1
    for n > 0 {
        rem := n % 10
        if rem != 0 {
            res += rem * mult
            mult *= 10
        }
        n /= 10
    }
    return res
}
