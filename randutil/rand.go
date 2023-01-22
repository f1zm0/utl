package randutil

import "math/rand"

const (
	asciiLowercase = "abcdefghijklmnopqrstuvwxyz"
	asciiUppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits         = "0123456789"
	special        = "-_"
)

var safeChars = []rune(asciiLowercase + asciiUppercase + digits + special)

// GetRandomStringN returns a randomly generated string of length n,
// composed of ASCII letters (lowercase, uppercase), digits, dash and underscore.
// Important: This function is not cryptographically secure, and it should not be used
// to perform any cryptographic operations (generating passwords, secrets, etc.)
func GetRandomStringN(slen int) string {
	b := make([]rune, slen)
	for i := range b {
		// rand.Intn() is not safe for crypto, but it's good enough for our use case.
		// #nosec G404
		b[i] = safeChars[rand.Intn(len(safeChars))]
	}
	return string(b)
}
