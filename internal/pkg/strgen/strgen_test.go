package strgen

import (
	"goapp/pkg/util"
	"testing"
)

// TestHexStringLength checks if the hex string generated is of the correct length
func TestHexStringLength(t *testing.T) {
	hexStr := util.RandHexString(20) // Expecting 20 characters
	if len(hexStr) != 20 {
		t.Errorf("Expected hex string length of 20, got %d", len(hexStr))
	}
}

// BenchmarkRandHexString measures the performance of hexadecimal string generation
func BenchmarkRandHexString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = util.RandHexString(20)
	}
}
