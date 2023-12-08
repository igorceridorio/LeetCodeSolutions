package longestsubstringwithoutrepeatingcharacters

// Using Sliding Window Technique
func lengthOfLongestSubstring(s string) int {
	result := 0
	start := 0
	charIndexMap := make(map[uint8]int)

	for end := 0; end < len(s); end++ {
		duplicateIndex, isDuplicate := charIndexMap[s[end]]
		if isDuplicate {
			result = max(result, end-start)

			for i := start; i <= duplicateIndex; i++ {
				delete(charIndexMap, s[i])
			}

			start = duplicateIndex + 1
		}
		charIndexMap[s[end]] = end
	}

	return max(result, len(s)-start)
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
