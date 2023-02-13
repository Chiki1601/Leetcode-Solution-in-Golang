func minimumScore(s string, t string) int {
    ns, nt := len(s), len(t)
    rcnt := make([]int, ns + 1)
    for i, j := ns - 1, nt - 1; i >= 0; i-- {
        if (j >= 0 && s[i] == t[j]) {
            rcnt[i] = rcnt[i + 1] + 1
            j--
        } else {
            rcnt[i] = rcnt[i + 1]
        }
    }
    if rcnt[0] >= nt {
        return 0
    }
    ret := nt - rcnt[0]
    for i, j := 0, 0; j < nt; i, j = i + 1, j + 1 {
        for i < ns && s[i] != t[j] {
            i++
        }
        if i >= ns {
            break
        }
        if subret := nt - (rcnt[i + 1] + j + 1); subret < ret {
            ret = subret
        }
    }
    return ret
}
