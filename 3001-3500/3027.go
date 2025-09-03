import (
    "sort"
)

func numberOfPairs(points [][]int) int {
    // Step 1: Sort points by x ascending, then by y descending
    sort.Slice(points, func(i, j int) bool {
        if points[i][0] == points[j][0] {
            return points[i][1] > points[j][1]
        }
        return points[i][0] < points[j][0]
    })
    pairCount := 0
    n := len(points)
    // Step 2: Iterate through all points as potential "upper-left" points
    for i := 0; i < n; i++ {
        upperY := points[i][1]
        lowerYLimit := -1 << 31 // math.MinInt32
        // Step 3: Check subsequent points as potential "bottom-right" points
        for j := i + 1; j < n; j++ {
            currentY := points[j][1]
            if currentY <= upperY && currentY > lowerYLimit {
                pairCount++
                lowerYLimit = currentY
                if currentY == upperY {
                    break
                }
            }
        }
    }
    return pairCount
}
