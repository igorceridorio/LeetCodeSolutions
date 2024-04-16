package reversepolishnotation

import "testing"

func TestEvalRPN(t *testing.T) {
	println(evalRPN([]string{"2", "1", "+", "3", "*"}))
	println(evalRPN([]string{"4", "13", "5", "/", "+"}))
	println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}
