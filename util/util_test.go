package util

import (
	"testing"
)

func TestCountToken(t *testing.T) {
	input := "Hello, world!"
	count := CountToken(input)
	if count <= 0 {
		t.Errorf("Expected token count > 0, got %d", count)
	}

	// Test with a known empty string
	input = ""
	count = CountToken(input)
	if count != 0 {
		t.Errorf("Expected token count 0 for empty string, got %d", count)
	}
}
