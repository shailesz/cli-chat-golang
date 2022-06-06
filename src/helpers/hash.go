package helpers

import (
	"crypto/sha256"
	"encoding/hex"
)

// Sha256 converts a string to sha256 hash string.
func Sha256(x string) string {
	h := sha256.New()

	h.Write([]byte(x))

	bs := h.Sum(nil)

	hs := hex.EncodeToString(bs)

	return hs
}
