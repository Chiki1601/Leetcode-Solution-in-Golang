func minCost(basket1 []int, basket2 []int) int64 {
    all, b1, b2 := make(map[int]int), make(map[int]int), make(map[int]int)

    for _, fruit := range basket1 {
        all[fruit]++
        b1[fruit]++
    }

    for _, fruit := range basket2 {
        all[fruit]++
        b2[fruit]++
    }

    smallest := 1000000001

    for fruit, count := range all {
        if count % 2 != 0 {
            return -1
        }
        smallest = min(smallest, fruit)
    }

    lack1, lack2 := []int{}, []int{}

    for fruit, count := range b1 {
        if count > b2[fruit] {
            diff := count - b2[fruit]
            diff >>= 1
            b1[fruit] -= diff
            b2[fruit] += diff
            
            for i := 0; i < diff; i++ {
                lack2 = append(lack2, fruit)
            }
        }
    }

    for fruit, count := range b2 {
        if count > b1[fruit] {
            diff := count - b1[fruit]
            diff >>= 1
            b2[fruit] -= diff
            b1[fruit] += diff

            for i := 0; i < diff; i++ {
                lack1 = append(lack1, fruit)
            }
        }
    }

    sort.Ints(lack1)
    sort.Sort(sort.Reverse(sort.IntSlice(lack2)))

    res := 0

    for i := 0; i < len(lack1); i++ {
        res += min(min(lack1[i], lack2[i]), smallest * 2)
    }
    
    return int64(res)
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
