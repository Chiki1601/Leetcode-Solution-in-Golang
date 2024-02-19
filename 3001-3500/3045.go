const (
    b int64 = 26 // Base for polynomial rolling hash
    m int64 = 1e9 + 9 // Large prime number for modulo operation
)

// ComputeHash calculates the hash of a string.
func zfunction(s string) []int {
  z := make([]int, len(s))
  l, r := 0, 0
  for i := 1; i < len(s); i++ {
    if i < r {
      z[i] = min(r - i, z[i - l])
    }
    for i+z[i] < len(s) && s[z[i]] == s[i+z[i]] {
      z[i]++
    }
    if i + z[i] > r {
      l = i
      r = i+z[i]
    }
  }
  return z
}

// ComputeHash calculates the hash of a string.
func ComputeHash(s string) int64 {
    hash := int64(0)
    power := int64(1) // b^0
    for i := 0; i < len(s); i++ {
        hash = (hash + int64(s[i]-'a'+1)*power) % m
        power = (power * b) % m
    }
    return hash
}

func countPrefixSuffixPairs(words []string) int64 {
  n := len(words)
  var res int64
  fixes := make(map[int64]int64) // both prefix and suffix
  for i := n-1; i >= 0; i-- {
    word := words[i]
    h := ComputeHash(word)
    res += fixes[h]
    z := zfunction(word)
    z[0] = len(word)
    wordlen := len(word)
    var hash, power int64
    hash, power = 0, 1
    for l := 1; l <= wordlen; l++ {
      hash = (hash + int64(word[l-1]-'a'+1)*power)%m
      power = (power * b) % m
      if z[wordlen-l] == l {
        fixes[hash]++
      }
    }
  }
  
  return res
}
