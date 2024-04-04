package productofarrayexceptself

func productExceptSelf(nums []int) []int {

	L := len(nums)
	result := make([]int, L)

	result[0] = 1
	for i := 1; i < L; i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	product := 1
	for i := L - 1; i >= 0; i-- {
		result[i] = result[i] * product
		product = product * nums[i]
	}

	return result
}
