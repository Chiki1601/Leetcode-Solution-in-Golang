func wordSubsets(words1 []string, words2 []string) []string {
    // Initialize the maximum frequency required for each character
    maxFrequencies := make([]int, 26) // For letters 'a' to 'z'
    lettersNeeded := make(map[int]struct{}) // Track letters needed

    // Function to count character frequencies in a word
    countFrequencies := func(word string) []int {
        frequencies := make([]int, 26)
        for _, c := range word {
            idx := c - 'a'
            frequencies[idx]++
        }
        return frequencies
    }

    // Compute maximum frequencies from words2
    for _, word := range words2 {
        wordFrequencies := countFrequencies(word)
        for i := 0; i < 26; i++ {
            if wordFrequencies[i] > maxFrequencies[i] {
                maxFrequencies[i] = wordFrequencies[i]
                lettersNeeded[i] = struct{}{}
            }
        }
    }

    // Convert lettersNeeded to a slice for faster iteration
    lettersNeededSlice := make([]int, 0, len(lettersNeeded))
    for i := range lettersNeeded {
        lettersNeededSlice = append(lettersNeededSlice, i)
    }

    // Slice to store universal words
    universalWords := []string{}

    // Check each word in words1
    for _, word := range words1 {
        wordFrequencies := countFrequencies(word)
        isUniversal := true
        for _, i := range lettersNeededSlice {
            if wordFrequencies[i] < maxFrequencies[i] {
                isUniversal = false
                break
            }
        }
        if isUniversal {
            universalWords = append(universalWords, word)
        }
    }

    return universalWords
}
