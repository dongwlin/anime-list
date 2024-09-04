package random

import (
	"math/rand"
	"strings"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz1234567890"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

// Int64 generates a random integer between min and max
func Int64(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// String generates a random string of length n
func String(n int) string {
	var sb strings.Builder
	k := len(charset)
	for i := 0; i < n; i++ {
		c := charset[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}
