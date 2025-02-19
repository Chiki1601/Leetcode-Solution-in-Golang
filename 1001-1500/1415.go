func getHappyString(n int, k int) string {
	var (
		backtrack func() bool
		seq       = make([]rune, 0, n)
	)

	backtrack = func() bool {
		if len(seq) == n {
			k--
			if k == 0 {
				return true
			}
			return false
		}

		for choice := 'a'; choice <= 'c'; choice++ {
			if len(seq) > 0 && seq[len(seq)-1] == choice {
				continue
			}
			seq = append(seq, choice)
			if backtrack() {
				return true
			}
			seq = seq[:len(seq)-1]
		}
		return false
	}
	backtrack()
	return string(seq)
}

