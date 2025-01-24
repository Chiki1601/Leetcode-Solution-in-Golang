type color uint8

const (
	white color = iota
	gray
	black
)

func eventualSafeNodes(graph [][]int) []int {
	v := len(graph)
	safe := make([]int, 0)
	colors := make([]color, v)

	for i := 0; i < v; i++ {
		if ok := dfs(graph, i, colors); ok {
			safe = append(safe, i)
		}
	}

	return safe
}

func dfs(g [][]int, i int, colors []color) bool {
	if colors[i] != white {
		return colors[i] == gray
	}

	colors[i] = black
	for _, v := range g[i] {
		switch colors[v] {
		case white:
			if ok := dfs(g, v, colors); !ok {
				return false
			}
		case gray:
			continue
		case black:
			return false
		}
	}
	colors[i] = gray
	return true
}
