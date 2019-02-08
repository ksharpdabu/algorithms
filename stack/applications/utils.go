package applications

import (
	"strconv"
)

// isDigit returns true if the char is digit, else false
func isDigit(char string) bool {
	_, err := strconv.Atoi(char)
	return err == nil
}

// isOperator returns true if the char is an operator, else false
func isOperator(char string) bool {
	// return strings.ContainsAny(string(c), "+ & - & * & /")
	switch char {
	case "+", "-", "*", "/", "%":
		return true
	}
	return false
}

// getPrecedence returns precedence value of the operator
func getPrecedence(char string) int {
	switch char {
	case "^":
		return 3
	case "*", "/", "%":
		return 2
	case "+", "-":
		return 1
	}
	return 0
}

// reverse returns the reverse string
func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}