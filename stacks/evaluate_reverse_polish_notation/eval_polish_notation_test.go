package rpn

import (
	"testing"
)

func TestEvalRPN(t *testing.T) {
	tests := []struct {
		name     string
		tokens   []string
		expected int
	}{
		{
			name:     "Basic addition and multiplication",
			tokens:   []string{"2", "1", "+", "3", "*"},
			expected: 9, // ((2 + 1) * 3)
		},
		{
			name:     "Substraction and division",
			tokens:   []string{"4", "13", "5", "/", "+"},
			expected: 6, // (4 + (13 / 5)) -> 4 + 2
		},
		{
			name:     "Complex expression with negative results",
			tokens:   []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			expected: 22,
		},
		{
			name:     "Truncation toward zero (positive)",
			tokens:   []string{"1", "2", "/"},
			expected: 0,
		},
		{
			name:     "Truncation toward zero (negative)",
			tokens:   []string{"-1", "2", "/"},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EvalRPN(tt.tokens)
			if result != tt.expected {
				t.Errorf("EvalRPN(%v) = %d; want %d", tt.tokens, result, tt.expected)
			}
		})
	}
}
