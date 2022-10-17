func checkIfPangram(sentence string) bool {
    arr := make([]int, 26)
	for i := 0; i < 26; i++ {
		arr[i] = 0
	}
	for i := 0; i < len(sentence); i++ {
		if sentence[i] >= 97 && sentence[i] <= 122 {
			arr[sentence[i]-97] = 1
		}
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			return false
		}
	}
	return true
}
