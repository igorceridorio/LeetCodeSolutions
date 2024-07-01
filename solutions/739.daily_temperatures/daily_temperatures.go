package dailytemperatures

func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0)
	result := make([]int, len(temperatures))

	for i := len(temperatures) - 1; i >= 0; i-- {

		// while stack is not empty and its top index is of a temperature less or equal than the current one
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] <= temperatures[i] {
			stack = stack[:len(stack)-1] // pop
		}

		// result to the current temperature is the higher stored index minus the current one
		if len(stack) > 0 {
			result[i] = stack[len(stack)-1] - i
		}

		stack = append(stack, i) // push
	}

	return result
}
