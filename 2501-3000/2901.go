func check(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var count int
	for i := range s {
		if s[i] != t[i] {
			count++
		}
	}

	return count == 1
}

func getWordsInLongestSubsequence(words []string, groups []int) []string {
	n := len(groups)
	f := make([]int, n)
	g := make([]int, n)
	for i := range f {
		f[i] = 1
		g[i] = -1
	}

	mx := 1
	for i, x := range groups {
		for j, y := range groups[:i] {
			if x != y && f[i] < f[j]+1 && check(words[i], words[j]) {
				f[i] = f[j] + 1
				g[i] = j
				if mx < f[i] {
					mx = f[i]
				}
			}
		}
	}

	result := make([]string, 0, mx)
	for i, x := range f {
		if x == mx {
			for j := i; j >= 0; j = g[j] {
				result = append(result, words[j])
			}
			break
		}
	}
    
	slices.Reverse(result)
	return result
}
