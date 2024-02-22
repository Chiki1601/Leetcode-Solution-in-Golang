type Person struct {
	id        int
	trusts    int
	trustedBy int
}

func findJudge(n int, trust [][]int) int {

	if len(trust) == 0 && n > 1 {
		return -1
	}

	if len(trust) == 0 && n == 1 {
		return 1
	}
	var p []Person

	for i := 0; i <= n; i++ {
		p = append(p, Person{
			id:        i,
			trusts:    0,
			trustedBy: 0,
		})
	}

	for i := 0; i < len(trust); i++ {
		p[trust[i][1]].trustedBy++
		p[trust[i][0]].trusts++
	}

	for i := 0; i < len(p); i++ {
		if p[i].trustedBy == n-1 && p[i].trusts == 0 {
			return i
		}
	}

	return -1
}
