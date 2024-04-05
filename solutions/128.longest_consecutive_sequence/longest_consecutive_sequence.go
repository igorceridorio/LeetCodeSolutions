package longestconsecutivesequence

func longestConsecutive(nums []int) int {
	longest := 0
	numsSet := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		numsSet[nums[i]] = true
	}

	for key := range numsSet {
		if _, ok := numsSet[key-1]; !ok {
			sequence := 1
			for {
				if _, ok := numsSet[key+sequence]; ok {
					sequence++
					continue
				}
				longest = max(sequence, longest)
				break
			}
		}
	}

	return longest
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
