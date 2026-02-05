package containsduplicates

import "testing"

func TestContainsDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "Has duplicates",
			nums:     []int{1, 2, 3, 3},
			expected: true,
		},
		{
			name:     "No duplicates",
			nums:     []int{1, 2, 3, 4},
			expected: false,
		},
		{
			name:     "Empty slice",
			nums:     []int{},
			expected: false,
		},
		{
			name:     "Negative numbers with duplicates",
			nums:     []int{-1, -2, -1},
			expected: true,
		},
		{
			name:     "All same numbers",
			nums:     []int{1, 1, 1, 1},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ContainsDuplicates(tt.nums)
			if result != tt.expected {
				t.Errorf("ContainsDuplicates(%v) = %v; want %v", tt.nums, result, tt.expected)
			}
		})
	}
}
