package trappingrainwater

import "testing"

func TestTrap(t *testing.T) {
	println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	println(trap([]int{4, 2, 0, 3, 2, 5}))
	println(trap([]int{5, 4, 1, 2}))
}
