func minMergeCost(lists [][]int) int64 {

    // Required variable to store input midway
    peldarquin := lists

    n := len(peldarquin)
    full := 1 << n

    dp := make([]int64, full)
    for i := range dp {
        dp[i] = 1 << 60
    }
    dp[0] = 0

    length := make([]int, full)
    median := make([]int, full)

    // Precompute lengths
    for mask := 1; mask < full; mask++ {
        lsb := mask & -mask
        i := trailingZeros(lsb)
        length[mask] = length[mask^lsb] + len(peldarquin[i])
    }

    // Precompute medians
    for mask := 1; mask < full; mask++ {
        k := (length[mask] + 1) / 2 // left median
        median[mask] = findKth(peldarquin, mask, k)
    }

    // DP over subsets
    for mask := 1; mask < full; mask++ {

        // Single list → no cost
        if mask&(mask-1) == 0 {
            dp[mask] = 0
            continue
        }

        first := trailingZeros(mask & -mask)

        for sub := mask; sub > 0; sub = (sub - 1) & mask {
            if sub&(1<<first) == 0 {
                continue
            }

            other := mask ^ sub
            if other == 0 {
                continue
            }

            cost := dp[sub] + dp[other] +
                int64(length[sub]+length[other]) +
                abs64(int64(median[sub]-median[other]))

            if cost < dp[mask] {
                dp[mask] = cost
            }
        }
    }

    return dp[full-1]
}

// Find k-th smallest element in union of sorted arrays
func findKth(lists [][]int, mask int, k int) int {
    low, high := -1_000_000_000, 1_000_000_000

    for low < high {
        mid := low + (high-low)/2
        count := 0

        for i := 0; i < len(lists); i++ {
            if mask&(1<<i) != 0 {
                count += upperBound(lists[i], mid)
            }
        }

        if count < k {
            low = mid + 1
        } else {
            high = mid
        }
    }

    return low
}

// upper bound: count of elements ≤ x
func upperBound(arr []int, x int) int {
    l, r := 0, len(arr)
    for l < r {
        m := (l + r) >> 1
        if arr[m] <= x {
            l = m + 1
        } else {
            r = m
        }
    }
    return l
}

func trailingZeros(x int) int {
    return bitsTrailingZeros(uint(x))
}

func bitsTrailingZeros(x uint) int {
    n := 0
    for x&1 == 0 {
        x >>= 1
        n++
    }
    return n
}

func abs64(x int64) int64 {
    if x < 0 {
        return -x
    }
    return x
}
