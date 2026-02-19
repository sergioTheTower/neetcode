package validparentheses

import "testing"

func TestIsValid(t *testing.T) {
	// Table-driven tests are the standard way to write tests in Go
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Example 1 - Simple Pair",
			input:    "[]",
			expected: true,
		},
		{
			name:     "Example 2 - Nested and Sequential",
			input:    "([{}])",
			expected: true,
		},
		{
			name:     "Example 3 - Wrong Order",
			input:    "[(])",
			expected: false,
		},
		{
			name:     "Edge Case - Single Bracket",
			input:    "[",
			expected: false,
		},
		{
			name:     "Edge Case - Only Open Brackets",
			input:    "(((",
			expected: false,
		},
		{
			name:     "Edge Case - Starts with Closing",
			input:    "]()",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validParentheses(tt.input)
			if result != tt.expected {
				t.Errorf("isValid(%q) = %v; expected %v", tt.input, result, tt.expected)
			}
		})
	}
}
