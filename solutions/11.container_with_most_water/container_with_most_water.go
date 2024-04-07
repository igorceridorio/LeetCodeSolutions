package mostwater

func maxArea(height []int) int {
	var maxArea int

	fwIdx := 0
	backIdx := len(height) - 1

	for {
		if fwIdx == backIdx || fwIdx > backIdx {
			break
		}

		minHeight := min(height[fwIdx], height[backIdx])
		maxArea = max(maxArea, minHeight*(backIdx-fwIdx))

		if height[fwIdx] <= height[backIdx] {
			fwIdx++
		} else {
			backIdx--
		}
	}

	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
