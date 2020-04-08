// +build regression

package regression

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
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

// None of the amount stuff is concurrency safe. Never ever use it for anything
// except serial tests.

const replacementPrefix = "generated-amount-"

// amount generates a token that is unique within a test. It is substituted for
// an amount drawn from the test's pool at runtime.
func amount(n int) string {
	return amountRange(n, 0.01, 1200)
}

func amountRange(n int, start, stop float64) string {
	return fmt.Sprintf("%s:%d:%.2f:%.2f", replacementPrefix, n, start, stop)
}

func getAmount(name, token string) string {
	tokens := strings.Split(token, ":")
	n := tokens[1]
	start, _ := strconv.ParseFloat(tokens[2], 64)
	stop, _ := strconv.ParseFloat(tokens[3], 64)

	if _, ok := testCache[name]; !ok {
		testCache[name] = map[string]string{}
	}

	if amount, ok := testCache[name][n]; ok {
		return amount
	}

	testCache[name][n] = randomRange(start, stop)

	return testCache[name][n]
}

var testCache = map[string]map[string]string{}

var amountCache = map[string]bool{}

func randomAmount() string {
	return randomRange(0.01, 100)
}

func randomRange(start, end float64) string {
	for {
		s := fmt.Sprintf("%.2f", (rand.Float64()*(end-start))+start)

		// Add a comma at the thousanths place in reverse
		chars := make([]byte, 0, len(s)+(len(s)-4)/3)
		for i := len(s) - 1; i >= 0; i-- {
			chars = append(chars, s[i])
			if n := len(s) - i; i > 0 && n > 3 && n%3 == 0 {
				chars = append(chars, byte(','))
			}
		}

		// Swap the slice
		for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
			chars[i], chars[j] = chars[j], chars[i]
		}

		s = string(chars)

		if isTrigger(s) {
			continue
		}

		if !amountCache[s] {
			amountCache[s] = true
			return s
		}
	}
}

func isTrigger(amt string) bool {
	if !strings.ContainsAny(amt, "123456890") {
		return true
	}

	return triggers[amt]
}

// Trigger amounts to avoid.
var triggers = map[string]bool{
	declineTriggerAmount:        true,
	partialAuthTriggerAmount:    true,
	partialAuthAuthorizedAmount: true,
	fraudWarningTriggerAmount:   true,
	errorTriggerAmount:          true,
	timeOutTriggerAmount:        true,
	noResponseTriggerAmount:     true,
}

// Default trigger amounts.
const (
	declineTriggerAmount        = "201.00"
	partialAuthTriggerAmount    = "55.00"
	partialAuthAuthorizedAmount = "25.00"
	fraudWarningTriggerAmount   = "66.00"
	errorTriggerAmount          = "0.11"
	timeOutTriggerAmount        = "68.00"
	noResponseTriggerAmount     = "72.00"
)
