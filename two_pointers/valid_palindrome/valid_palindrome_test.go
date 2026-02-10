package validpalindrome

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	// Define a struct for test cases
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "Example 1: Complex sentence with symbols",
			input: "Was it a car or a cat I saw?",
			want:  true,
		},
		{
			name:  "Example 2: Not a palindrome",
			input: "tab a cat",
			want:  false,
		},
		{
			name:  "Simple Palindrome",
			input: "aba",
			want:  true,
		},
		{
			name:  "Empty string (Edge Case)",
			input: "",
			want:  true,
		},
		{
			name:  "Single character (Edge Case)",
			input: "x",
			want:  true,
		},
		{
			name:  "Only non-alphanumeric characters",
			input: "!!!",
			want:  true, // effectively empty after cleaning
		},
		{
			name:  "Case sensitivity check",
			input: "Aa",
			want:  true,
		},
		{
			name:  "Numbers check",
			input: "12321",
			want:  true,
		},
		{
			name:  "Space Check",
			input: "12321",
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validPalindrome(tt.input)
			if got != tt.want {
				t.Errorf("IsPalindrome(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
