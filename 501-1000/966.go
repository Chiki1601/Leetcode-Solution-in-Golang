// bitmask: A=0, E=4, I=8, O=14, U=20
// (1<<pos) marks vowels in uppercase alphabet positions
const vowelMask = (1 << 0) | (1 << 4) | (1 << 8) | (1 << 14) | (1 << 20)

// normalize converts a word into a pattern where all vowels are replaced by 'A'.
// Example: "Leet" -> "LAAT"
// This helps handle vowel errors during matching.
func normalize(s string) string {
	b := []byte(s)
	for i := 0; i < len(s); i++ {
		if (1<<(s[i]-'A'))&vowelMask != 0 {
			b[i] = 'A'
		}
	}
	return string(b)
}

// spellchecker matches queries against the given wordlist with three rules:
// 1. Exact match
// 2. Case-insensitive match
// 3. Vowel-error match (all vowels considered interchangeable)
func spellchecker(wordlist []string, queries []string) []string {
	// Rule 1: Exact match lookup
	dictExact := make(map[string]struct{})

	// Rule 2: Case-insensitive lookup (store first occurrence)
	dictUpper := make(map[string]string)

	// Rule 3: Vowel-pattern lookup (store first occurrence)
	dictPattern := make(map[string]string)

	// Build dictionaries in one pass
	for _, w := range wordlist {
		// exact word
		dictExact[w] = struct{}{}

		// case-insensitive
		up := strings.ToUpper(w)
		if _, ok := dictUpper[up]; !ok {
			dictUpper[up] = w
		}

		// vowel-normalized
		ptn := normalize(up)
		if _, ok := dictPattern[ptn]; !ok {
			dictPattern[ptn] = w
		}
	}

	// Answer queries in order
	ans := make([]string, len(queries))
	for i, q := range queries {
		// Rule 1: Exact match
		if _, found := dictExact[q]; found {
			ans[i] = q
			continue
		}

		// Rule 2: Case-insensitive match
		up := strings.ToUpper(q)
		if orig, found := dictUpper[up]; found {
			ans[i] = orig
			continue
		}

		// Rule 3: Vowel-error match
		ptn := normalize(up)
		if orig, found := dictPattern[ptn]; found {
			ans[i] = orig
			continue
		}

		// No match
		ans[i] = ""
	}

	return ans
}
