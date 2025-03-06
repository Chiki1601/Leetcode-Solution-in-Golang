func findMissingAndRepeatedValues(grid [][]int) []int {
	n := len(grid)
	vector := make(map[int]int)

	for i := 1; i <= n*n; i++ {
		vector[i] = 0
	}

	for _, item := range grid {
		for _, num := range item {
			vector[num]++
		}
	}

	var a, b int
	for num, frequency := range vector {
		if frequency > 1 {
			a = num
		}
		if frequency == 0 {
			b = num
		}
		if a != 0 && b != 0 {
			break
		}
	}

	return []int{a, b}
}
