package mostwater

import "testing"

func TestMaxArea(t *testing.T) {
	println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	println(maxArea([]int{1, 1}))
	println(maxArea([]int{2, 3, 4, 5, 18, 17, 6}))
}
