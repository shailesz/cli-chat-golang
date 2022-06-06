package helpers

import (
	"crypto/sha256"
	"encoding/hex"
)

// Sha256 generates sha256 hashed string from given string.
func Sha256(x string) string {
	h := sha256.New()

	h.Write([]byte(x))

	bs := h.Sum(nil)

	hs := hex.EncodeToString(bs)

	return hs
}
