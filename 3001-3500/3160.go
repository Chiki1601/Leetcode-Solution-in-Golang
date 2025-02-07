func queryResults(limit int, queries [][]int) []int {
	ans := make([]int, len(queries))
	// `ballToColor` could be an array but the max val of limit is 10^9 which is too large which could lead to memory exhaustion
	// on the other hand, the number of queries is limited to 10^5.
	// So it's reasonable to make `ballToColor` a map.
	ballToColor, colorToBall := make(map[int]int), make(map[int]int)
	for i, query := range queries {
		if c, ok := ballToColor[query[0]]; ok { // ball was already colored
			cc := colorToBall[c]
			if cc == 1 { // no more balls with this color
				delete(colorToBall, c)
			} else { // there are more balls with this color. So decrease the count.
				colorToBall[c] = cc - 1
			}
		}
		ballToColor[query[0]] = query[1]
		colorToBall[query[1]] = colorToBall[query[1]] + 1
		ans[i] = len(colorToBall)
	}

	return ans
}
