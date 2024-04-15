package minstack

import (
	"fmt"
	"testing"
)

func TestMinStack(t *testing.T) {

	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Printf("Min: %d\n", minStack.GetMin())
	minStack.Pop()
	fmt.Printf("Top: %d\n", minStack.Top())
	fmt.Printf("Min: %d\n", minStack.GetMin())
}
