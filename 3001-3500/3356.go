func minZeroArray(nums []int, queries [][]int) int {
	queryIndex := 0        
	curSum := 0            
	q := priorityqueue.NewWith(func(a, b [2]int) int {
		return a[0] - b[0] 
	})

	for i := 0; i < len(nums); i++ {
		processScheduledEffects(q, &curSum, i)

		for curSum < nums[i] && queryIndex < len(queries) {
			start, end, delta := queries[queryIndex][0], queries[queryIndex][1], queries[queryIndex][2]

			if start <= i {
				curSum += delta
			} else {
				q.Enqueue([2]int{start, delta}) 
			}

			if end+1 <= i {
				curSum -= delta
			} else {
				q.Enqueue([2]int{end + 1, -delta}) 
			}

			queryIndex++
		}

		if queryIndex == len(queries) && curSum < nums[i] {
			return -1
		}
	}

	return queryIndex
}

func processScheduledEffects(q *priorityqueue.Queue[[2]int], curSum *int, currentIndex int) {
	for !q.Empty() {
		top, _ := q.Peek()
		if top[0] > currentIndex {
			break
		}

		*curSum += top[1]

		q.Dequeue()
	}
}
