func longestCommonPrefix(arr1 []int, arr2 []int) int {
    arr1Map := make(map[int]bool)

    for _, v := range arr1 {
        for v != 0 {
            arr1Map[v] = true
            v = v / 10
        }
    }

    currLen := 0
    for _, v := range arr2 {
        for v != 0 {
            if _, ok := arr1Map[v]; ok {
                l := intLen(v)
                if l > currLen {
                    currLen = l
                }
            }
            v = v / 10
        }
    }
    return currLen
}

func intLen(v int) int {
    res := 0
    for v != 0 {
        res++
        v = v / 10
    }
    return res
}
