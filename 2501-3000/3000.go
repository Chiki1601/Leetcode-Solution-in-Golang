func areaOfMaxDiagonal(dimensions [][]int) int {
    maxArea, maxDiag := 0, 0

    for _, d := range dimensions {
        l, w := d[0], d[1]
        currDiag := l*l + w*w

        if currDiag > maxDiag || (currDiag == maxDiag && l*w > maxArea) {
            maxDiag = currDiag
            maxArea = l * w
        }
    }
    return maxArea
}
