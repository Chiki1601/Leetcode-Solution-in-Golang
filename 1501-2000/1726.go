func tupleSameProduct(nums []int) int {
	h := make(map[int32]uint8)
	var result int
	for i, num1 := range nums[1:] {
		for _, num2 := range nums[:i+1] {
			product := int32(num1 * num2)
			count := h[product]
			result += int(count)
			h[product] = count + 1
		}
	}
	return result << 3
}
