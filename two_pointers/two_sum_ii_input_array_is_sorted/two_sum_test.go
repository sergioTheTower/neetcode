package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	// Define a struct for our test cases
	tests := []struct {
		name    string // Description of the test case
		numbers []int  // Input array
		target  int    // Target sum
		want    []int  // Expected output indices
	}{
		{
			name:    "Example 1: Basic positive numbers",
			numbers: []int{1, 2, 3, 4},
			target:  3,
			want:    []int{1, 2},
		},
		{
			name:    "Example 2: Indices are not adjacent",
			numbers: []int{2, 3, 4},
			target:  6,
			want:    []int{1, 3}, // 2 + 4 = 6
		},
		{
			name:    "Example 3: Negative numbers",
			numbers: []int{-1, 0},
			target:  -1,
			want:    []int{1, 2},
		},
		{
			name:    "Example 4: Mixed negative and positive",
			numbers: []int{-5, -2, 0, 3, 7},
			target:  2,
			want:    []int{1, 5}, // -5 + 7 = 2
		},
		{
			name:    "Example 5: Target is zero",
			numbers: []int{-3, 1, 2, 3},
			target:  0,
			want:    []int{1, 4}, // -3 + 3 = 0
		},
	}

	// Iterate over the test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.numbers, tt.target)

			// Use DeepEqual to compare slices
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
