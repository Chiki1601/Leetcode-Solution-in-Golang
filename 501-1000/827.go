var moves = [5]int16{0, 1, 0, -1, 0}

func largestIsland(grid [][]int) int {
	n := int16(len(grid))
	var islandId uint32
	var i, j, nextI, nextJ int16
	var islandCells [][2]int16
	var cell [2]int16
	var islandMark int
	var result int
	markIsland := func() {
		islandCells = islandCells[:0]
		islandCells = append(islandCells, [2]int16{i, j})
		grid[i][j] = 2
		for islandCellI := 0; islandCellI < len(islandCells); islandCellI++ {
			cell = islandCells[islandCellI]
			for moveI := range 4 {
				nextI, nextJ = cell[0]+moves[moveI], cell[1]+moves[moveI+1]
				if nextI == -1 || nextI == n || nextJ == -1 || nextJ == n || grid[nextI][nextJ] != 1 {
					continue
				}
				grid[nextI][nextJ] = 2
				islandCells = append(islandCells, [2]int16{nextI, nextJ})
			}
		}
		result = max(result, len(islandCells))
		islandMark = int(islandId)<<32 | len(islandCells)
		for _, cell = range islandCells {
			grid[cell[0]][cell[1]] = islandMark
		}
	}
	for i = 0; i < n; i++ {
		for j = 0; j < n; j++ {
			if grid[i][j] == 1 {
				islandId++
				markIsland()
			}
		}
	}
	joinedIslandIds := make([]uint32, 0, 4)
	var subResult int
	for i = 0; i < n; i++ {
		for j = 0; j < n; j++ {
			if grid[i][j] == 0 {
				subResult = 1
				joinedIslandIds = joinedIslandIds[:0]
			MoveLoop:
				for moveI := range 4 {
					nextI, nextJ = i+moves[moveI], j+moves[moveI+1]
					if nextI == -1 || nextI == n || nextJ == -1 || nextJ == n || grid[nextI][nextJ] == 0 {
						continue
					}
					islandId = uint32(grid[nextI][nextJ] >> 32)
					for _, joinedIslandId := range joinedIslandIds {
						if islandId == joinedIslandId {
							continue MoveLoop
						}
					}
					joinedIslandIds = append(joinedIslandIds, islandId)
					subResult += grid[nextI][nextJ] & 0xFFFFFFFF
				}
				result = max(result, subResult)
			}
		}
	}
	return result
}
