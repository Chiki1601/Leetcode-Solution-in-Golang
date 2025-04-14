
func countGoodTriplets(arr []int, a int, b int, c int) int {
	i := 0
	j := 1
	k := 2

	res := 0

	for {

		val1 := Abs(arr[i] - arr[j])
		val2 := Abs(arr[j] - arr[k])
		val3 := Abs(arr[i] - arr[k])

		if val1 <= a && val2 <= b && val3 <= c {
			res++
		}

		if i == len(arr)-3 {
			break
		}

		k++

		if k == len(arr) {
			j++

			k = j + 1
		}

		if j == len(arr)-1 {
			i++

			j = i + 1

			k = j + 1
		}

	}

	return res
}

func Abs(v int) int {
	if v < 0 {
		return -v
	}

	return v
}
