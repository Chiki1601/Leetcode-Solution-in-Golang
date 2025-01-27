func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	parents := make([][]int8, numCourses)
	for _, prerequisite := range prerequisites {
		parents[prerequisite[1]] = append(parents[prerequisite[1]], int8(prerequisite[0]))
	}
	var visited [100][100]bool
	var leaf int
	var dfs func(child int8)
	dfs = func(child int8) {
		for _, parent := range parents[child] {
			if visited[leaf][parent] {
				continue
			}
			visited[leaf][parent] = true
			dfs(parent)
		}
	}
	for leaf = range numCourses {
		dfs(int8(leaf))
	}
	result := make([]bool, len(queries))
	for i, query := range queries {
		result[i] = visited[query[1]][query[0]]
	}
	return result
}
