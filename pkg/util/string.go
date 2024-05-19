package util

import (
	"encoding/hex"
	"math/rand"
)

var randx = rand.NewSource(42)

// RandString returns a random string of length n.
func RandString(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const (
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)

	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, randx.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = randx.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func RandHexString(n int) string {
	bytes := make([]byte, n/2) // Generate half the length because each byte becomes two hex characters
	if _, err := rand.Read(bytes); err != nil {
		panic(err) // For simplicity in this example; handle error appropriately in production code
	}
	return hex.EncodeToString(bytes)
}
