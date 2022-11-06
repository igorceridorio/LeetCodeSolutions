func romanToInt(s string) int {
	romanToIntMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	sum := 0

	for i := 0; i < len(s); {
		if i+1 < len(s) {
			if romanToIntMap[string(s[i])] >= romanToIntMap[string(s[i+1])] {
				sum += romanToIntMap[string(s[i])]
				i = i + 1
			} else {
				sum += romanToIntMap[string(s[i+1])] - romanToIntMap[string(s[i])]
				i = i + 2
			}
		} else {
			sum += romanToIntMap[string(s[i])]
			i = i + 1
		}
	}

	return sum
}
