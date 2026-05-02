// GRID-SAST-004: math/rand for token entropy.
// GRID-SAST-009: SHA-1 used for token signing.
package auth

import (
	"crypto/sha1"
	"encoding/hex"
	"math/rand"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz0123456789"

func init() { rand.Seed(time.Now().UnixNano()) }

func NewSessionID() string {
	b := make([]byte, 24)
	for i := range b {
		b[i] = alphabet[rand.Intn(len(alphabet))]   // GRID-SAST-004
	}
	return string(b)
}

func sign(token string, secret string) string {
	h := sha1.New()                                  // GRID-SAST-009
	h.Write([]byte(token + secret))
	return hex.EncodeToString(h.Sum(nil))
}
