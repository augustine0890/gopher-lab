package main

import (
	"flag"
	"log"
)

// Given a string mathematical expression implement a function that outputs if the expression has balanced brackets.
// A stack is a last in, last out (LIFO) data structure

// isBalanced returns whether the given expression has balanced brackets.
func isBalanced(expr string) bool {
	bracketsPairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	var stack []rune
	for _, r := range expr {
		switch r {
		case '(', '{', '[':
			stack = append(stack, r)
		case ')', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != bracketsPairs[r] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

// printResult prints whether the expression is balanced.
func printResult(expr string, balanced bool) {
	if balanced {
		log.Printf("%s is balanced.\n", expr)
		return
	}
	log.Printf("%s is not balanced.\n", expr)
}

// go run balanced_brackets.go -expr="1+(2*3)"
func main() {
	expr := flag.String("expr", "", "The expression to validate brackets on.")
	flag.Parse()
	printResult(*expr, isBalanced(*expr))
}
