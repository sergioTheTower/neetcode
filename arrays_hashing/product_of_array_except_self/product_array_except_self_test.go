package productofarrayexceptself

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "Example NC",
			nums:     []int{1, 2, 3, 4},
			expected: []int{24, 12, 8, 6},
		},
		{
			name:     "Example 1",
			nums:     []int{1, 2, 4, 6},
			expected: []int{48, 24, 12, 8},
		},
		{
			name:     "Example 2",
			nums:     []int{-1, 0, 1, 2, 3},
			expected: []int{0, -6, 0, 0, 0},
		},
		{
			name:     "Two Zeros",
			nums:     []int{1, 0, 3, 0},
			expected: []int{0, 0, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ProductExceptSelf(tt.nums)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
