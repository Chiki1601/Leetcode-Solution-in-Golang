
func maxFreeTime(ev int, k int, st []int, end []int) int {
    dist := make([]int, 0)
    sl := len(st)

    // init first meeting gap value, (e.g. first meeting start at 18, instead of 0). 
    dist = append(dist, st[0])
    max := st[0]


    // count gap, between next start meeting (i) with prev end meeting (i-1).
    for i := 1; i < sl; i++ {
        tmp := st[i] - end[i-1]
        dist = append(dist, tmp)
        if tmp > max {
            max = tmp
        }
    }
    // init last meeting gap
    tmdif := ev - end[sl-1]
    dist = append(dist, tmdif)
    if tmdif > max {
        max = tmdif
    }

    // we can't move the schedule, so we return the biggest value
    if k == 0 {
        return max
    }

    s0 := 0
    s1 := k

    slide := 0
    dl := len(dist)
    for i := s0; i <= s1; i++ { // initial window value.
        slide += dist[i]
    }
    if slide > max {
        max = slide
    }

    for s1 < dl-1 {
        slide -= dist[s0] // slowly remove first value and add new value
        s0++
        s1++
        slide += dist[s1]

        if slide > max {
            max = slide
        }
    }

    return max
}
