package evaluate_reverse_polish_notation

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	l := len(tokens)
	stack := make([]int, 0, l)

	for _, v := range tokens {
		if !isSymbol(v) {
			i, _ := strconv.Atoi(v)
			stack = append(stack, i)
		} else {
			l := len(stack)
			b := stack[l-1]
			a := stack[l-2]
			stack = stack[:l-2]

			switch v {
			case "+":
				stack = append(stack, a+b)
			case "-":
				stack = append(stack, a-b)
			case "/":
				stack = append(stack, a/b)
			case "*":
				stack = append(stack, a*b)
			}
		}
	}

	return stack[0]
}

func isSymbol(s string) bool {
	if s == "+" || s == "-" || s == "*" || s == "/" {
		return true
	} else {
		return false
	}
}
