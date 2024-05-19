package util

import (
	"testing"
)

// TestHexStringLength checks if the hex string generated is of the correct length
func TestHexStringLength(t *testing.T) {
	hexStr := RandString(20) // Expecting 20 characters
	if len(hexStr) != 20 {
		t.Errorf("Expected hex string length of 20, got %d", len(hexStr))
	}
}

// TestIsHexStringValid checks that the string contains only hexadecimal characters.
func TestIsHexStringValid(t *testing.T) {
	hexStr := RandString(20)
	for _, char := range hexStr {
		if !isValidHexChar(char) {
			t.Errorf("Invalid hex character found: %q", char)
		}
	}
}

// isValidHexChar checks if a character is a valid hexadecimal digit.
func isValidHexChar(c rune) bool {
	return ('0' <= c && c <= '9') || ('A' <= c && c <= 'F')
}

// BenchmarkRandHexString measures the performance of hexadecimal string generation
func BenchmarkRandHexString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = RandString(20)
	}
}
