func walk(grid [][]int, x, y int) int {
	if y == len(grid) {
		return x
	}
	if grid[y][x] == 1 {
		if x == len(grid[y])-1 || grid[y][x+1] == -1 {
			return -1
		}
		return walk(grid, x+1, y+1)
	}
	if x == 0 || grid[y][x-1] == 1 {
		return -1
	}
	return walk(grid, x-1, y+1)
}

func findBall(grid [][]int) []int {
	var (
		answer = make([]int, len(grid[0]))
	)
	for i := 0; i < len(answer); i++ {
		answer[i] = walk(grid, i, 0)
	}
	return answer
}
