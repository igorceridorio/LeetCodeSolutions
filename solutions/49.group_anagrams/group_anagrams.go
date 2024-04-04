package groupanagrams

func groupAnagrams(strs []string) [][]string {

	anagramsMap := make(map[[26]int][]string)

	for i := 0; i < len(strs); i++ {
		var count [26]int

		for j := 0; j < len(strs[i]); j++ {
			count[strs[i][j]-'a'] += 1
		}

		anagramsMap[count] = append(anagramsMap[count], strs[i])
	}

	var response [][]string
	for _, value := range anagramsMap {
		response = append(response, value)
	}

	return response
}
