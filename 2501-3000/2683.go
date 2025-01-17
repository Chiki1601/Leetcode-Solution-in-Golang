func doesValidArrayExist(derived []int) bool {
    s := 0
    for _, d := range derived {
        s = s^d
    }
    return s==0
}
