package trapwater

import (
	"testing"
)

func TestTrap(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "Example 1: Standard case",
			height: []int{0, 2, 0, 3, 1, 0, 1, 3, 2, 1},
			want:   9,
		},
		{
			name:   "Example 2: Classic bowl shape",
			height: []int{4, 2, 0, 3, 2, 5},
			want:   9,
		},
		{
			name:   "Ascending stairs (no trap)",
			height: []int{0, 1, 2, 3, 4},
			want:   0,
		},
		{
			name:   "Descending stairs (no trap)",
			height: []int{4, 3, 2, 1, 0},
			want:   0,
		},
		{
			name:   "Flat surface",
			height: []int{5, 5, 5, 5},
			want:   0,
		},
		{
			name:   "W shape",
			height: []int{5, 2, 5, 2, 5},
			want:   6, // 3 units in first dip + 3 units in second dip
		},
		{
			name:   "Empty input",
			height: []int{},
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := trapWater(tt.height)
			if got != tt.want {
				t.Errorf("trap(%v) = %d; want %d", tt.height, got, tt.want)
			}
		})
	}
}
