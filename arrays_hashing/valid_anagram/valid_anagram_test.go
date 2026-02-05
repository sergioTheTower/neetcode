package validanagram

import "testing"

func TestValidAnagram(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		t        string
		expected bool
	}{
		{
			name:     "Example 1: racecar and carrace",
			s:        "racecar",
			t:        "carrace",
			expected: true,
		},
		{
			name:     "Example 2: jar and jam",
			s:        "jar",
			t:        "jam",
			expected: false,
		},
		{
			name:     "Different lengths",
			s:        "apple",
			t:        "aple",
			expected: false,
		},
		{
			name:     "Empty strings",
			s:        "",
			t:        "",
			expected: true,
		},
		{
			name:     "Same characters different counts",
			s:        "aabb",
			t:        "abab",
			expected: true,
		},
		{
			name:     "One string is a subset of another",
			s:        "a",
			t:        "ab",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidAnagram(tt.s, tt.t)
			if result != tt.expected {
				t.Errorf("ValidAnagram(%q, %q) = %v; want %v", tt.s, tt.t, result, tt.expected)
			}
		})
	}
}
