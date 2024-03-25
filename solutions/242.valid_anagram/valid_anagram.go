package validanagram

func isAnagram(s string, t string) bool {
	lettersFrequencyMap := make(map[string]int)

	if len(s) != len(t) {
		return false
	}

	for i := 0; i < len(s); i++ {
		_, ok := lettersFrequencyMap[(string(s[i]))]
		if !ok {
			lettersFrequencyMap[string(s[i])] = 1
		} else {
			lettersFrequencyMap[string(s[i])] += 1
		}
	}

	for i := 0; i < len(t); i++ {
		_, ok := lettersFrequencyMap[string(t[i])]
		if !ok {
			return false
		} else {
			lettersFrequencyMap[string(t[i])] -= 1
		}
	}

	for _, value := range lettersFrequencyMap {
		if value > 0 {
			return false
		}
	}

	return true
}
