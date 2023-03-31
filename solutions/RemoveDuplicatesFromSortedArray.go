func removeDuplicates(nums []int) int {
	array := make([]int, 0)
	array = append(array, nums[0])
	k := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] != array[k] {
			array = append(array, nums[i])
			k++
		}
	}

	copy(nums, array)
	return len(array)
}
