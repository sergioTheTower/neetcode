// Package validparentheses provides functionality to validate parentheses strings.
package validparentheses

func validParentheses(s string) bool {
	stack := []rune{}
	// Store the closing paren as the key and open as the value.
	closeToOpen := map[rune]rune{')': '(', ']': '[', '}': '{'}
	for _, c := range s {
		if open, ok := closeToOpen[c]; ok {
			// If there is a close paren in the stack but the stack is empty it's not balanced.
			if len(stack) == 0 {
				return false
			}
			// Stack is not empty, pop off first element and update stack.
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// If the popped element doesn't match the corresponding open bracket, it's not balanced.
			if top != open {
				return false
			}
			continue
		}
		// If 'c' is an open bracket, push it onto the stack
		stack = append(stack, c)
	}
	// The string is valid if the stack is completely empty at the end
	return len(stack) == 0
}
