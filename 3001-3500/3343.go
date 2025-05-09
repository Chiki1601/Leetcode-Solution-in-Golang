
var mod int = int(1e9 + 7)

func countBalancedPermutations(num string) int {
    cnt := make([]int, 10)
	total := 0
	for i := 0; i < len(num); i++ {
		digit := int(num[i] - '0')
		cnt[digit]++
		total += digit
	}

	if total%2 != 0 {
		return 0
	}

	halfSum := total / 2
	n := len(num)
	evenCount := n / 2
	oddCount := n - evenCount

	var comb = func(n, k int) int {
		if k > n {
			return 0
		}
		res := 1
		for i := 0; i < k; i++ {
			res = res * (n - i) / (i + 1)
		}
		return res % mod
	}
	memo := make(map[[4]int]int)

	var dfs func(digit, odd, even, balance int) int
	dfs = func(digit, odd, even, balance int) int {
		if odd == 0 && even == 0 && balance == 0 {
			return 1
		}
		if digit < 0 || odd < 0 || even < 0 || balance < 0 {
			return 0
		}

		key := [4]int{digit, odd, even, balance}
		if val, ok := memo[key]; ok {
			return val
		}

		res := 0
		for j := 0; j <= cnt[digit]; j++ {
			oddUsed := j
			evenUsed := cnt[digit] - j
			if oddUsed > odd || evenUsed > even {
				continue
			}
			chooseOdd := comb(odd, oddUsed)
			chooseEven := comb(even, evenUsed)
			totalWays := (chooseOdd * chooseEven) % mod
			res += (totalWays * dfs(digit-1, odd-oddUsed, even-evenUsed, balance-digit*oddUsed)) % mod
			res %= mod
		}

		memo[key] = res
		return res
	}

	return dfs(9, oddCount, evenCount, halfSum)
}
