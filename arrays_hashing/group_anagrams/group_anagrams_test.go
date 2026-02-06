package groupanagrams

import (
	"reflect"
	"sort"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected [][]string
	}{
		{
			name:  "Standard Anagrams",
			input: []string{"act", "pots", "tops", "cat", "stop", "hat"},
			expected: [][]string{
				{"act", "cat"},
				{"hat"},
				{"pots", "stop", "tops"},
			},
		},
		{
			name:     "Single Character",
			input:    []string{"x"},
			expected: [][]string{{"x"}},
		},
		{
			name:     "Empty String",
			input:    []string{""},
			expected: [][]string{{""}},
		},
		{
			name:  "Multiple Empty and Single Strings",
			input: []string{"", "b", ""},
			expected: [][]string{
				{"", ""},
				{"b"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := GroupAnagrams(tt.input)

			// Helper to sort sub-slices and the main slice for comparison
			sortResult(actual)
			sortResult(tt.expected)

			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("groupAnagrams(%v) = %v; want %v", tt.input, actual, tt.expected)
			}
		})
	}
}

// sortResult sorts each inner slice and then the outer slice to ensure
// we can compare the result regardless of map iteration order.
func sortResult(res [][]string) {
	for i := range res {
		sort.Strings(res[i])
	}
	sort.Slice(res, func(i, j int) bool {
		if len(res[i]) != len(res[j]) {
			return len(res[i]) < len(res[j])
		}
		return res[i][0] < res[j][0]
	})
}
