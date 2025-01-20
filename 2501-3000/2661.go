func firstCompleteIndex(arr []int, mat [][]int) int {
    rows, cols := len(mat), len(mat[0])
    positionMap := make(map[int][2]int)
    rowCount := make([]int, rows)
    colCount := make([]int, cols)
    
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            positionMap[mat[i][j]] = [2]int{i, j}
        }
    }
    
    for i := 0; i < rows; i++ {
        rowCount[i] = cols
    }
    
    for i := 0; i < cols; i++ {
        colCount[i] = rows
    }

    for idx, val := range arr {
        row, col := positionMap[val][0], positionMap[val][1]
        rowCount[row]--
        colCount[col]--
        if rowCount[row] == 0 || colCount[col] == 0 {
            return idx
        }
    }
    
    return -1
}
