package trappingrainwater

func trap(height []int) int {
	water := 0

	pL := 0
	pR := len(height) - 1
	maxL := height[pL]
	maxR := height[pR]

	for pL < pR {
		if height[pL] <= height[pR] {
			pL++
			if height[pL] < maxL {
				water += maxL - height[pL]
			} else {
				maxL = height[pL]
			}
		} else {
			pR--
			if height[pR] < maxR {
				water += maxR - height[pR]
			} else {
				maxR = height[pR]
			}
		}
	}

	return water
}
