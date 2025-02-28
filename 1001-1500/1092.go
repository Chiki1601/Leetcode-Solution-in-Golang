func shortestCommonSupersequence(str1 string, str2 string) string {
    // Step 1: Find the longest common subsequence using dynamic programming
    m, n := len(str1), len(str2)
    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }
    
    // Fill the dp table
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if str1[i-1] == str2[j-1] {
                dp[i][j] = 1 + dp[i-1][j-1]
            } else {
                dp[i][j] = max(dp[i-1][j], dp[i][j-1])
            }
        }
    }
    
    // Step 2: Construct the shortest common supersequence
    // Start from the bottom right of the dp table
    i, j := m, n
    var result []byte
    
    for i > 0 && j > 0 {
        if str1[i-1] == str2[j-1] {
            // If the characters are the same, add it once
            result = append(result, str1[i-1])
            i--
            j--
        } else if dp[i-1][j] > dp[i][j-1] {
            // If coming from top has higher value, take character from str1
            result = append(result, str1[i-1])
            i--
        } else {
            // Otherwise, take character from str2
            result = append(result, str2[j-1])
            j--
        }
    }
    
    // Add remaining characters from str1 (if any)
    for i > 0 {
        result = append(result, str1[i-1])
        i--
    }
    
    // Add remaining characters from str2 (if any)
    for j > 0 {
        result = append(result, str2[j-1])
        j--
    }
    
    // Reverse the result to get the final supersequence
    reverseBytes(result)
    
    return string(result)
}

// Helper function to reverse a byte slice
func reverseBytes(s []byte) {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
}

// Helper function for max of two integers
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
