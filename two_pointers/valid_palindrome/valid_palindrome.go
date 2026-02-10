// Package validpalindrome provides functionality to validate if a string is a valid palindrome.
package validpalindrome

import (
	"strings"
	"unicode"
)

func validPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	s = strings.ToLower(s)
	left, right := 0, len(s)-1
	for left < right {
		if !isCharOrNumber(rune(s[left])) {
			left++
			continue
		}
		if !isCharOrNumber(rune(s[right])) {
			right--
			continue
		}
		if s[left] != s[right] {
			return false

		}
		left++
		right--
	}
	return true
}

func isCharOrNumber(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}
