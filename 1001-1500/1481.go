func findLeastNumOfUniqueInts(arr []int, k int) int {
    m := make(map[int]int)
    for _, v := range arr {
        m[v]++
    }
    var count []int
    for _, v := range m {
        count = append(count, v)
    }
    sort.Ints(count)
    var idx int
    for k > 0 {
        k -= count[idx]
        idx++
    }
    if k == 0 { return len(count) - idx }
    return len(count) - idx + 1
}
