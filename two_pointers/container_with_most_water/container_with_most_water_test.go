package containswater

import (
	"testing"
)

// TestMaxArea is the test function
func TestMaxArea(t *testing.T) {
	// Table-driven tests
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "Example 1: Mixed heights",
			height: []int{1, 7, 2, 5, 4, 7, 3, 6},
			want:   36,
		},
		{
			name:   "Example 2: Equal heights",
			height: []int{2, 2, 2},
			want:   4,
		},
		{
			name:   "Minimum length (2 bars)",
			height: []int{1, 1},
			want:   1,
		},
		{
			name:   "Large difference in heights",
			height: []int{1, 100, 100, 1},
			want:   100,
		},
		{
			name:   "Descending order",
			height: []int{5, 4, 3, 2, 1},
			want:   6, // Bar 5 and Bar 3 (dist 2) -> min(5,3)*2 = 6
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := containsWater(tt.height)
			if got != tt.want {
				t.Errorf("maxArea(%v) = %d; want %d", tt.height, got, tt.want)
			}
		})
	}
}
