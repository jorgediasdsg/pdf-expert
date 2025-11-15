package pdfanalyzer

import "testing"

func TestCountWords(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"simple", "hello world", 2},
		{"multiple spaces", "hello   world", 2},
		{"tabs and newlines", "hello\tworld\nagain", 3},
		{"empty", "", 0},
		{"spaces only", "     ", 0},
		{"unicode text", "olá mundo maravilhoso", 3},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := countWords(tc.input)
			if got != tc.expected {
				t.Errorf("countWords(%q) = %d; want %d", tc.input, got, tc.expected)
			}
		})
	}
}
