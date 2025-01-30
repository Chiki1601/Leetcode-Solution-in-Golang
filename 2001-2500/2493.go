func magnificentSets(n int, edges [][]int) int {
	res := 0
	graph := make([][]int, n)

	for _, edge := range edges {
		a, b := edge[0]-1, edge[1]-1
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	dp := make([]int, n)

	for start := range dp {
		queue := []int{start}
		dist := make([]int, n)
		dist[start] = 1
		maxDepth, root := 1, start

		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]
			root = min(root, node)

			for _, neighbor := range graph[node] {
				if dist[neighbor] == 0 {
					dist[neighbor] = dist[node] + 1
					maxDepth = max(maxDepth, dist[neighbor])
					queue = append(queue, neighbor)
				} else if abs(dist[neighbor]-dist[node]) != 1 {
					return -1
				}
			}
		}
		dp[root] = max(dp[root], maxDepth)
	}

	for _, value := range dp {
		res += value
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
