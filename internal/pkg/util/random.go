package util

import (
	"math/rand"
	"strings"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz1234567890"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

// RandomInt64 generates a random integer between min and max
func RandomInt64(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(charset)
	for i := 0; i < n; i++ {
		c := charset[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomBool generates a random bool
func RandomBool() bool {
	if RandomInt64(0, 1) == 1 {
		return true
	}
	return false
}

// RandomUsername generates a random username
func RandomUsername() string {
	return RandomString(6)
}
