func findDifferentBinaryString(nums []string) string {
    n := len(nums[0])
    for i := 0; i < (1 << n); i++ {
        binaryString := fmt.Sprintf("%0*b", n, i)
        if !contains(nums, binaryString) {
            return binaryString
        }
    }
    return ""
}



func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
