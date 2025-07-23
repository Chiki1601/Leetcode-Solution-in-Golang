func processAB(s string, x int, y int) int {
	var stk []byte
	ans := 0
	for _, c := range s {
		if len(stk) > 0 && stk[len(stk)-1] == 'a' && c == 'b' {
			stk = stk[:len(stk)-1]
			ans += x
		} else {
			stk = append(stk, byte(c))
		}
	}
	a, b := 0, 0
	for _, c := range stk {
		if c == 'a' {
			a++
		} else {
			b++
		}
	}
	return min(a, b) * y + ans
}

func processBA(s string, x int, y int) int {
	var stk []byte
	ans := 0
	for _, c := range s {
		if len(stk) > 0 && stk[len(stk)-1] == 'b' && c == 'a' {
			stk = stk[:len(stk)-1]
			ans += y
		} else {
			stk = append(stk, byte(c))
		}
	}
	a, b := 0, 0
	for _, c := range stk {
		if c == 'a' {
			a++
		} else {
			b++
		}
	}
	return min(a, b) * x + ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maximumGain(s string, x int, y int) int {
    var ans int
	for i := 0; i < len(s); i++ {
		if s[i] != 'a' && s[i] != 'b' {
			continue
		}
		j := i
		var t strings.Builder
		for j < len(s) && (s[j] == 'a' || s[j] == 'b') {
			t.WriteByte(s[j])
			j++
		}
		ans += max(processAB(t.String(), x, y), processBA(t.String(), x, y))
		i = j - 1
	}
	return ans
}
