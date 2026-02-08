package encodedecode

import (
	"reflect"
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	sol := &Solution{}

	tests := []struct {
		name  string
		input []string
	}{
		{
			name:  "Standard case",
			input: []string{"hello", "world"},
		},
		{
			name:  "Contains delimiter",
			input: []string{"hello#world", "test#123"},
		},
		{
			name:  "Empty strings",
			input: []string{"", "", ""},
		},
		{
			name:  "Empty slice",
			input: []string{},
		},
		{
			name:  "Long strings and spaces",
			input: []string{"The quick brown fox", "   ", "1234567890"},
		},
		{
			name:  "Unicode characters",
			input: []string{"你好", "🌍", "coding"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoded := sol.Encode(tt.input)
			decoded := sol.Decode(encoded)

			if !reflect.DeepEqual(tt.input, decoded) {
				t.Errorf("Decode() = %v, want %v", decoded, tt.input)
			}
		})
	}
}
