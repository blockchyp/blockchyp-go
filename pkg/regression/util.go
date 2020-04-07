// +build regression

package regression

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomStr() string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, 24)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}

func randomSMSNum() string {
	const charset = "0123456789"

	b := make([]byte, 10)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}

var amtCache = map[string]bool{}

func randomAmt() string {
	for {
		n := fmt.Sprintf("%.2f", rand.Float64())
		if !amtCache[n] {
			amtCache[n] = true
			return n
		}
	}
}
