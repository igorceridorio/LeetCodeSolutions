package reversepolishnotation

import (
	"strconv"
)

var stack []int

var ops = map[string]func(int, int) int{
	"+": func(a, b int) int { return a + b },
	"-": func(a, b int) int { return a - b },
	"*": func(a, b int) int { return a * b },
	"/": func(a, b int) int { return a / b },
}

func evalRPN(tokens []string) int {
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "+" || tokens[i] == "-" || tokens[i] == "*" || tokens[i] == "/" {
			b := pop()
			a := pop()
			result := ops[tokens[i]](a, b)
			push(result)
		} else {
			val, _ := strconv.Atoi(tokens[i])
			push(val)
		}
	}

	return pop()
}

func push(val int) {
	stack = append(stack, val)
}

func pop() int {
	value := stack[len(stack)-1]
	stack = stack[:len(stack)-1]

	return value
}
