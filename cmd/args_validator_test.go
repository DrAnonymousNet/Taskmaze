package cmd

import "testing"

func TestValidateTime(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "valid time",
			input:    "3:30PM",
			expected: true,
		},
		{
			name:     "invalid time",
			input:    "3:30",
			expected: false,
		},
		{
			name:     "invalid format",
			input:    "3:30 AM",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateDueTime(tt.input); got != tt.expected {
				t.Errorf("ValidateTime(%q) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}
