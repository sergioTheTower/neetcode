package threesum

import (
	"reflect"
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name: "Example 1",
			nums: []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			name:     "Example 2 - No solution",
			nums:     []int{0, 1, 1},
			expected: [][]int{},
		},
		{
			name:     "Example 3 - All zeros",
			nums:     []int{0, 0, 0},
			expected: [][]int{{0, 0, 0}},
		},
		{
			name:     "Empty input",
			nums:     []int{},
			expected: [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call your function here
			got := threeSum(tt.nums)

			// Normalize both results to ensure comparison works regardless of order
			sortResult(got)
			sortResult(tt.expected)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("threeSum() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// sortResult is a helper to sort the list of triplets and the triplets themselves
// so that reflect.DeepEqual can verify the contents regardless of order.
func sortResult(res [][]int) {
	// 1. Sort each individual triplet
	for _, triplet := range res {
		sort.Ints(triplet)
	}

	// 2. Sort the list of triplets based on the first, then second, then third element
	sort.Slice(res, func(i, j int) bool {
		for k := 0; k < 3; k++ {
			if res[i][k] != res[j][k] {
				return res[i][k] < res[j][k]
			}
		}
		return false
	})
}
