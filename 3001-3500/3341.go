func minTimeToReach(moveTime [][]int) int {
	// Problem constraints
	const (
		maxMoveTime = 1_000_000_000
		maxN        = 50
		maxM        = 50
	)

	const infMoveTime = maxMoveTime + maxM + maxN
	n := len(moveTime)
	m := len(moveTime[0])

	minTime := make([][]int, len(moveTime))
	for i := range n {
		minTime[i] = make([]int, m)
		for j := range minTime[i] {
			minTime[i][j] = infMoveTime
		}
	}

	minTime[0][0] = 0
	h := &minHeap{}
	heap.Push(h, item{0, 0, 0})

	for h.Len() > 0 {
		curr := heap.Pop(h).(item)
		if curr.row == n-1 && curr.col == m-1 {
			break
		}
		if curr.minTime > minTime[curr.row][curr.col] {
			continue
		}

		for _, neighbour := range getNeighbours(curr, n, m, minTime) {
			time := 1 + max(curr.minTime, moveTime[neighbour.row][neighbour.col])
			if time < neighbour.minTime {
				minTime[neighbour.row][neighbour.col] = time
				neighbour.minTime = time
				heap.Push(h, neighbour)
			}
		}
	}
	return minTime[n-1][m-1]
}

type item struct {
	row     int
	col     int
	minTime int
}

type minHeap []item

func (h minHeap) Len() int { return len(h) }

func (h minHeap) Less(i, j int) bool { return h[i].minTime < h[j].minTime }

func (h minHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x any) { *h = append(*h, x.(item)) }

func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	result := old[n-1]
	*h = old[:n-1]
	return result
}

func getNeighbours(cur item, n, m int, minTime [][]int) []item {
	result := make([]item, 0, 4)
	if cur.row > 0 {
		result = append(result, item{cur.row - 1, cur.col, minTime[cur.row-1][cur.col]})
	}
	if cur.row < n-1 {
		result = append(result, item{cur.row + 1, cur.col, minTime[cur.row+1][cur.col]})
	}
	if cur.col > 0 {
		result = append(result, item{cur.row, cur.col - 1, minTime[cur.row][cur.col-1]})
	}
	if cur.col < m-1 {
		result = append(result, item{cur.row, cur.col + 1, minTime[cur.row][cur.col+1]})
	}
	return result
}
