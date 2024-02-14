// 165ms
func rearrangeArray(nums []int) []int {
	result := make([]int, len(nums))
	pos, neg := 0, 1
	ptr := getPointer(&pos, &neg)
	for _, n := range nums {
		result[*ptr(n)] = n
		*ptr(n) += 2
	}
	return result
}

func getPointer(pos, neg *int) func(int) *int {
	return func(n int) *int {
		if n < 0 {
			return neg
		} else {
			return pos
		}
	}
}
