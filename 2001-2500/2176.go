func countPairs(nums []int, k int) int {
	count := 0
	m := make(map[int][]int, 0)
	for i, v := range nums {
		if _, ok := m[v]; !ok {
			m[v] = append(m[v], i)
			continue
		}
		for _, j := range m[v] {
			if i*j%k == 0 {
				count++
			}
		}
		m[v] = append(m[v], i)
	}
	return count
}
