func addToArrayForm(num []int, k int) []int {
    num[len(num) - 1] += k
    index := len(num) - 1
    for index >= 0 {
        if num[index] < 10 {
            return num
        }
        if index == 0 {
            num = append([]int{num[index] / 10}, num...)
            num[1] = num[1] % 10
        } else {
            num[index - 1] += num[index] / 10
            num[index] = num[index] % 10
            index--
        }
    }
    return num
}
