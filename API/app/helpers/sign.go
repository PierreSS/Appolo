package helpers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// Sign signs provided payload and returns encoded string sum.
func Sign(payload []byte, key []byte) string {
	mac := hmac.New(sha256.New, key)
	mac.Write(payload)

	return hex.EncodeToString(mac.Sum(nil))
}
