package exercises

import "fmt"

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// An input string is valid if:

// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.

func isValidParentheses(s string) bool {
	bracketMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	// Create a stack to hold the opening brackets
	var stack []rune

	// if it is an opening bracket, add it to the stack
	// if it is a closing bracket, if the last element on the stack is the correct opening bracket, pop it off. If it is the incorrect opening bracket, return false.
	for _, char := range s {
		_, ok := bracketMap[char]
		// if ok means it is a closing bracket
		if ok {
			// if the last element in the stack is the correct opening bracket pop it off, otherwise return false
			if len(stack) > 0 && stack[len(stack)-1] == bracketMap[char] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
			// if !ok, then it is an opening parentheses and we append it to the stack
		} else {
			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}

func IsValidParenthesesTester() bool {
	fmt.Print(isValidParentheses("()"))     //true
	fmt.Print(isValidParentheses("()[]{}")) //true
	fmt.Print(isValidParentheses("(]"))     //false
	fmt.Print(isValidParentheses("]"))      //false

	return true
}
