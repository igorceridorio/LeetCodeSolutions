package generateparentheses

import (
	"strings"
)

var stack []string
var response []string

func generateParenthesis(n int) []string {
	stack = nil
	response = nil

	backtracking(0, 0, n)

	return response
}

func backtracking(open, close, n int) {
	if (open == close) && (close == n) {
		response = append(response, strings.Join(stack[:], ""))
		return
	}

	if open < n {
		stack = append(stack, "(")
		backtracking(open+1, close, n)
		stack = stack[:len(stack)-1]
	}

	if close < open {
		stack = append(stack, ")")
		backtracking(open, close+1, n)
		stack = stack[:len(stack)-1]
	}
}
