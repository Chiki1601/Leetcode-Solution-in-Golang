func highestPeak(isWater [][]int) [][]int {
    R := len(isWater)
    C := len(isWater[0])
    height := make([][]int, R)

    for i := range height {
        height[i] = make([]int, C)
        for j := range height[i] {
            height[i][j] = 1<<31 - 1
        }
    }

    for i := 0; i < R; i++ {
        for j := 0; j < C; j++ {
            if isWater[i][j] == 1 {
                height[i][j] = 0
            } else {
                if i > 0 {
                    height[i][j] = min(height[i][j], height[i-1][j] + 1)
                }
                if j > 0 {
                    height[i][j] = min(height[i][j], height[i][j-1] + 1)
                }
            }
        }
    }

    for i := R - 1; i >= 0; i-- {
        for j := C - 1; j >= 0; j-- {
            if i < R - 1 {
                height[i][j] = min(height[i][j], height[i+1][j] + 1)
            }
            if j < C - 1 {
                height[i][j] = min(height[i][j], height[i][j+1] + 1)
            }
        }
    }

    return height
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
