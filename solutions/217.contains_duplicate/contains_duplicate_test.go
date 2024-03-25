package containsduplicate

import (
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	println(containsDuplicate([]int{1, 2, 3, 1}))
	println(containsDuplicate([]int{1, 2, 3, 4}))
	println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}
