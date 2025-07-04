func kthCharacter(k int64, operations []int) byte {
    k--
	c := 0
	for i := len(operations) - 1; i >= 0; i-- {
		if k>>i&1 == 1 {
			c += operations[i]
		}
	}
	return 'a' + byte(c%26)
}
