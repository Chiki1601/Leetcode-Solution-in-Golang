func replaceNonCoprimes(nums []int) []int {
	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}
	lcm := func(a, b int) int {
		return a * b / gcd(a, b)
	}

	stack := make([]int, 0, len(nums))
	var n int

	for _, num := range nums {
		stack = append(stack, num)
		n++
		for n > 1 && gcd(stack[n-2], stack[n-1]) != 1 {
			stack[n-2] = lcm(stack[n-2], stack[n-1])
			stack = stack[:n-1]
			n--
		}
	}
	return stack
}
