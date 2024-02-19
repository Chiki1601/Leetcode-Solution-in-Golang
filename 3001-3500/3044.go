func isPrime(num int) bool {
    if num < 2 {
        return false
    }
    for i := 2; i*i <= num; i++ {
        if num%i == 0 {
            return false
        }
    }
    return true
}

func mostFrequentPrime(mat [][]int) int {
    m, n := len(mat), len(mat[0])
    dirs := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}}
    freq := make(map[int]int)
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            for _, dir := range dirs {
                x, y := i, j
                num := 0
                for x >= 0 && x < m && y >= 0 && y < n {
                    num = num*10 + mat[x][y]
                    if num > 10 && isPrime(num) {
                        freq[num]++
                    }
                    x += dir[0]
                    y += dir[1]
               }
            }
        }
    }

    maxFreq := 0
    mostFreqPrime := -1
    for prime, f := range freq {
        if f > maxFreq || (f == maxFreq && prime > mostFreqPrime) {
            maxFreq = f
            mostFreqPrime = prime
        }
    }

    return mostFreqPrime
}
