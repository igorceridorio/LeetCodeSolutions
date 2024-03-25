package containsduplicate

func containsDuplicate(nums []int) bool {
	numsMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		_, ok := numsMap[nums[i]]
		if ok {
			return true
		} else {
			numsMap[nums[i]] = nums[i]
		}
	}

	return false
}
