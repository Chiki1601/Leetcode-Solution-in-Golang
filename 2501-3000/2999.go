func numberOfPowerfulInt(start int64, finish int64, limit int, s string) int64 {
    // Convert the range bounds to strings and subtract start by 1 to make it inclusive
    start_ := strconv.FormatInt(start - 1, 10)
    finish_ := strconv.FormatInt(finish, 10)
    // The count is the difference between finish and start-1
    return calculate(finish_, s, limit) - calculate(start_, s, limit)
}

func calculate(x string, s string, limit int) int64 {
    // If x has fewer digits than s, no numbers can have s as suffix
    if len(x) < len(s) {
        return 0
    }
    // If x has same digits as s, just compare directly
    if len(x) == len(s) {
        if x >= s {  // If x is ≥ s, then s itself is the only possible number
            return 1
        }
        return 0
    }

    // Get the suffix part that must match s
    suffix := x[len(x) - len(s):]
    var count int64
    preLen := len(x) - len(s)  // Length of prefix before suffix

    // Process each digit in the prefix
    for i := 0; i < preLen; i++ {
        digit := int(x[i] - '0')  // Current digit being examined
        
        // If digit exceeds limit, all numbers with this prefix are invalid
        if limit < digit {
            // Add all possible numbers with remaining digits ≤ limit
            count += int64(math.Pow(float64(limit + 1), float64(preLen - i)))
            return count
        }
        
        // Add numbers where current digit is less than x's digit
        // (and remaining digits can be anything ≤ limit)
        count += int64(digit) * int64(math.Pow(float64(limit + 1), float64(preLen - 1 - i)))
    }
    
    // After processing all prefix digits, check if suffix is valid
    if suffix >= s {
        count++  // The number formed by prefix + s is valid
    }
    return count
}
