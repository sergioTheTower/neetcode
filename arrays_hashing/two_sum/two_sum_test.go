package twosum

import (
	"slices"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "Example 1: Basic case",
			nums:     []int{3, 4, 5, 6},
			target:   7,
			expected: []int{0, 1},
		},
		{
			name:     "Example 2: Gap in indices",
			nums:     []int{4, 5, 6},
			target:   10,
			expected: []int{0, 2},
		},
		{
			name:     "Example 3: Duplicate numbers",
			nums:     []int{5, 5},
			target:   10,
			expected: []int{0, 1},
		},
		{
			name:     "Large numbers",
			nums:     []int{1000000, 500, 2000000},
			target:   3000000,
			expected: []int{0, 2},
		},
		{
			name:     "Target is first two elements",
			nums:     []int{1, 2, 3, 4, 5},
			target:   3,
			expected: []int{0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TwoSum(tt.nums, tt.target)
			if !slices.Equal(got, tt.expected) {
				t.Errorf("TwoSum() = %v, want %v", got, tt.expected)
			}
		})
	}
}
