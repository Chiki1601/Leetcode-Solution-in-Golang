func stringMatching(words []string) []string {
    tmp := make(map[string]bool)
    res := []string{}
    for _, v := range words {
        for _, V := range words {
            if v == V || len(V) > len(v) { continue }
            if _, ok := tmp[V]; !ok && strings.Contains(v, V) {
                tmp[V] = true
                res = append(res, V)
            }
        }
    }
    return res
}
