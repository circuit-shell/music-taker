package greeting

import "testing"

func TestGetGreeting(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "basic greeting",
			input:    "Developer",
			expected: "Hello, Developer! Welcome to Go programming!",
		},
		{
			name:     "empty name",
			input:    "",
			expected: "Hello, ! Welcome to Go programming!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetGreeting(tt.input)

			if result != tt.expected {
				t.Errorf("GetGreeting(%s) = %s; want %s", tt.input, result, tt.expected)
			}
		})
	}
}
