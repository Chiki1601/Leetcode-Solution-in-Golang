func passThePillow(n int, time int) int {
    time %= (n * 2) - 2

    if time > n - 1 {
        time -= n - 1
        return n - time
    }

    return time + 1
}
