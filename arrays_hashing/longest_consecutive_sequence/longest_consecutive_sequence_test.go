package lcs

import (
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	// Table-driven tests definition
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{2, 20, 4, 10, 3, 4, 5},
			expected: 4, // Sequence: [2, 3, 4, 5]
		},
		{
			name:     "Example 2 (with duplicates and zero)",
			nums:     []int{0, 3, 2, 5, 4, 6, 1, 1},
			expected: 7, // Sequence: [0, 1, 2, 3, 4, 5, 6]
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: 0,
		},
		{
			name:     "Single element",
			nums:     []int{100},
			expected: 1,
		},
		{
			name:     "No consecutive elements",
			nums:     []int{10, 30, 50},
			expected: 1,
		},
		{
			name:     "Negative numbers",
			nums:     []int{-1, -2, -3, 0, 1},
			expected: 5, // Sequence: [-3, -2, -1, 0, 1]
		},
		{
			name:     "Unordered simple sequence",
			nums:     []int{100, 4, 200, 1, 3, 2},
			expected: 4, // Sequence: [1, 2, 3, 4]
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := longestConsecutive(tc.nums)
			if result != tc.expected {
				t.Errorf("longestConsecutive(%v) = %d; want %d",
					tc.nums, result, tc.expected)
			}
		})
	}
}
