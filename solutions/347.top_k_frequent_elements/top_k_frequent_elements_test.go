package topkfrequentelements

import (
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	println(topKFrequent([]int{1,1,1,2,2,3}, 2))
	println(topKFrequent([]int{1}, 1))
}
