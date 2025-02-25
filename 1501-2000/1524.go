func numOfSubarrays(arr []int) int {
	r := 0
	o := 0
	e := 0
	for _, n := range arr {
		if n%2 == 0 {
			e++
		} else {
			e, o = o, e+1
		}
		r += o
	}
	return r%1000000007
}
