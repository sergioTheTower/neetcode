package topkfrequent

import (
	"reflect"
	"sort"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name     string
		nums     []string
		numsInt  []int
		k        int
		expected []int
	}{
		{
			name:     "Example 1",
			numsInt:  []int{1, 2, 2, 3, 3, 3},
			k:        2,
			expected: []int{2, 3},
		},
		{
			name:     "Example 2",
			numsInt:  []int{7, 7},
			k:        1,
			expected: []int{7},
		},
		{
			name:     "Single Element",
			numsInt:  []int{1},
			k:        1,
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := topKFrequent(tt.numsInt, tt.k)

			// Since order doesn't matter, sort both before comparing
			sort.Ints(actual)
			sort.Ints(tt.expected)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("topKFrequent() = %v, want %v", actual, tt.expected)
			}
		})
	}
}
