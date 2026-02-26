package largestrectangleinhistogram

import "testing"

// TestLargestRectangleArea tests the largestRectangleArea function.
func TestLargestRectangleArea(t *testing.T) {
	// Define the test cases
	tests := []struct {
		name     string
		heights  []int
		expected int
	}{
		{
			name:     "Example 1 from problem description",
			heights:  []int{7, 1, 7, 2, 2, 4},
			expected: 8,
		},
		{
			name:     "Example 2 from problem description",
			heights:  []int{1, 3, 7},
			expected: 7,
		},
		{
			name:     "Single element",
			heights:  []int{5},
			expected: 5,
		},
		{
			name:     "All same heights",
			heights:  []int{2, 2, 2, 2},
			expected: 8,
		},
		{
			name:     "Ascending heights",
			heights:  []int{1, 2, 3, 4},
			expected: 6, // 3 * 2 = 6
		},
		{
			name:     "Descending heights",
			heights:  []int{4, 3, 2, 1},
			expected: 6, // 3 * 2 = 6
		},
	}

	// Iterate over all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := largestRectangleArea(tt.heights)

			if result != tt.expected {
				t.Errorf("got %d, want %d for heights %v", result, tt.expected, tt.heights)
			}
		})
	}
}
