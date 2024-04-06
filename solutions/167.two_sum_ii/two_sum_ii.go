package twosumii

func twoSum(numbers []int, target int) []int {
	fwIdx := 0
	backIdx := len(numbers) - 1

	for {
		if numbers[fwIdx]+numbers[backIdx] == target {
			break
		} else if numbers[fwIdx]+numbers[backIdx] > target {
			backIdx--
		} else {
			fwIdx++
		}
	}

	return []int{fwIdx + 1, backIdx + 1}
}
