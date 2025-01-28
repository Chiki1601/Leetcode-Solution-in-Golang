var moves = [5]int8{0, 1, 0, -1, 0}

func findMaxFish(grid [][]int) int {
	m := int8(len(grid))
	n := int8(len(grid[0]))
	var result, subResult int
	var stack [][2]int8
	var cur [2]int8
	for i, row := range grid {
		for j, cell := range row {
			if cell == 0 {
				continue
			}
			subResult = cell
			row[j] = 0
			stack = append(stack, [2]int8{int8(i), int8(j)})
			for len(stack) != 0 {
				cur = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				for moveI := range 4 {
					nextI, nextJ := cur[0]+moves[moveI], cur[1]+moves[moveI+1]
					if nextI == -1 || nextI == m || nextJ == -1 || nextJ == n || grid[nextI][nextJ] == 0 {
						continue
					}
					subResult += grid[nextI][nextJ]
					grid[nextI][nextJ] = 0
					stack = append(stack, [2]int8{nextI, nextJ})
				}
			}
			result = max(result, subResult)
		}
	}
	return result
}
