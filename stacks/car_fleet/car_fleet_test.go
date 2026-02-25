package carfleet

import "testing"

func TestCarFleet(t *testing.T) {
	// Define the table of test cases
	tests := []struct {
		name     string
		target   int
		position []int
		speed    []int
		expected int
	}{
		{
			name:     "Example 1",
			target:   10,
			position: []int{1, 4},
			speed:    []int{3, 2},
			expected: 1,
		},
		{
			name:     "Example 2",
			target:   10,
			position: []int{4, 1, 0, 7},
			speed:    []int{2, 2, 1, 1},
			expected: 3,
		},
		{
			name:     "Single Car",
			target:   10,
			position: []int{3},
			speed:    []int{3},
			expected: 1,
		},
		{
			name:     "Already at Target Speed 0 (Edge Case)",
			target:   100,
			position: []int{0, 2, 4},
			speed:    []int{4, 2, 1},
			expected: 1,
		},
	}

	// Iterate over the test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := carFleet(tt.target, tt.position, tt.speed)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}
