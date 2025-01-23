func countServers(grid [][]int) int {
    Rows := make([]int, len(grid))
    Col := make([]int, len(grid[0]))
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            Rows[i] += grid[i][j]
            Col[j] += grid[i][j]
        }
    }
    ans := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == 1 && (Rows[i] > 1 || Col[j] > 1) {
                ans++
            }
        }
    }
    return ans
}
