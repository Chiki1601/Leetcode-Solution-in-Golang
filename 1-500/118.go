func generate(numRows int) [][]int {
    ans := [][]int{}
    ans = append(ans, []int{1})
    if numRows == 1 {
        return ans
    }
    
    ans = append(ans, []int{1, 1})
    if numRows == 2 {
        return ans
    }

    for i := 3; i <= numRows; i++ {
        if i > len(ans) {
            ans = append(ans, make([]int, i))
            ans[i - 1][0] = 1
            ans[i - 1][i - 1] = 1
        }

        for j := 1; j < i - 1; j++ {
            ans[i - 1][j] = ans[i - 2][j - 1] + ans[i - 2][j]
        }
    }

    return ans
}
