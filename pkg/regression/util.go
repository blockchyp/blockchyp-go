package regression

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomStr() string {
	rand.Seed(time.Now().UnixNano())
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, 24)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}

func randomSMSNum() string {
	rand.Seed(time.Now().UnixNano())
	const charset = "0123456789"

	b := make([]byte, 10)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}

// None of the amount stuff is concurrency safe. Never ever use it for anything
// except serial tests.

const replacementPrefix = "generated-amount"

// amount is a shorthand helper for amountRange with a default range.
func amount(n int) string {
	return amountRange(n, 1, 120000)
}

// amountRange returns a string instructing the interpreter to generate a
// random amount at runtime. Each subsequent request for an amount, given the
// same token (n) will return the initially generated amount. This way, a test
// can generate an amount at runtime and then use the same amount for
// assertions.
func amountRange(n int, start, stop int) string {
	return fmt.Sprintf("%s:%d:%d:%d", replacementPrefix, n, start, stop)
}

func getAmount(name, token string) string {
	return getMoney(name, token).String()
}

// getMoney interprets a string to generate or return an existing amount. The
// initial request for a given token (n) will generate a new amount within the
// range. Subsequent requests with the same token (n) will return the
// initlially generated amount.
//
// An optional addition or multiplication token can be included in the string.
func getMoney(name, token string) money {
	tokens := strings.Split(token, ":")
	n := tokens[1]
	start, _ := strconv.Atoi(tokens[2])
	stop, _ := strconv.Atoi(tokens[3])

	var feesOnly, cashDiscount bool
	var add, txfeeRate int
	var mult, flatfeeRate float64
	for i := range tokens {
		switch tokens[i] {
		case "add":
			add, _ = strconv.Atoi(tokens[i+1])
		case "mult":
			mult, _ = strconv.ParseFloat(tokens[i+1], 64)
		case "txfee":
			txfeeRate, _ = strconv.Atoi(tokens[i+1])
		case "flatfee":
			flatfeeRate, _ = strconv.ParseFloat(tokens[i+1], 64)
		case "fees":
			feesOnly = true
		case "cashDiscount":
			cashDiscount = true
		}
	}

	if _, ok := testCache[name]; !ok {
		testCache[name] = map[string]money{}
	}

	var result money
	if m, ok := testCache[name][n]; ok {
		result = m
	} else {
		result = newMoney(start, stop)
		testCache[name][n] = result
	}

	result = result.mult(mult).add(add)

	txfee := money(txfeeRate)

	var flatfee money
	if flatfeeRate > 0 {
		flatfee = result.mult(flatfeeRate)
	}

	if feesOnly {
		return txfee.add(int(flatfee))
	}
	if cashDiscount {
		flatfee = -flatfee
		txfee = -txfee
	}

	return result.add(int(txfee)).add(int(flatfee))
}

var testCache = map[string]map[string]money{}

var amountCache = map[money]bool{}

func randomAmount() string {
	return newMoney(1, 10000).String()
}

// money represents an amount in cents.
type money int

// newMoney creates a new instance of a money, randomly generated within the
// given range, unique within the test run, and excluding
// known trigger amounts.
func newMoney(start, end int) money {
	for {
		amount := rand.Intn(end-start) + start

		m := money(amount)

		if isTrigger(m) {
			continue
		}

		if !amountCache[m] {
			amountCache[m] = true
			return m
		}
	}
}

// String formats the money using the same format that the BlockChyp API server
// uses for currency amounts.
func (m money) String() string {
	s := strconv.Itoa(int(m))
	var cents, dollars string
	if len(s) > 2 {
		dollars = s[0 : len(s)-2]
		cents = s[len(dollars):len(s)]
	} else {
		dollars = "0"
		cents = s
	}

	// Add a comma at the thousanths place in reverse
	chars := make([]byte, 0, len(dollars)+((len(dollars)-1)/3))
	for i := len(dollars) - 1; i >= 0; i-- {
		chars = append(chars, dollars[i])
		if n := len(dollars) - i; i > 0 && n >= 3 && n%3 == 0 {
			chars = append(chars, byte(','))
		}
	}

	// Swap the slice
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}

	return string(chars) + "." + cents
}

func (m money) add(n int) money {
	m += money(n)

	return m
}

func (m money) mult(n float64) money {
	if n != 0 {
		result := math.Ceil(float64(m) * n)
		m = money(result)

	}

	return m
}

// add returns a string that instructs the interpreter to add an amount to the
// base amount at assertion time.
func add(s string, money int) string {
	return fmt.Sprintf("%s:add:%d", s, money)
}

// mult returns a string that instructs the interpreter to multiply the
// base amount at assertion time.
func mult(s string, multiplier float64) string {
	return fmt.Sprintf("%s:mult:%f", s, multiplier)
}

// addFees adds default test fees.
func addFees(s string) string {
	return fmt.Sprintf("%s:flatfee:%f:txfee:%d", s, defaultFlatRateFee, defaultPerTransactionFee)
}

// fees returns only the fee portion using default test fees.
func fees(s string) string {
	return addFees(s) + ":fees"
}

// cashDiscount discounts the amount by the default test fees.
func cashDiscount(s string) string {
	return addFees(s) + ":cashDiscount"
}

func isTrigger(m money) bool {
	if !strings.ContainsAny(m.String(), "123456890") {
		return true
	}

	return triggers[m.String()]
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

func getAddr() string {
	for i := 8080; i < 9000; i++ {
		addr := fmt.Sprintf("localhost:%d", i)
		ln, err := net.Listen("tcp", addr)
		if err != nil {
			continue
		}
		ln.Close()

		return addr
	}

	panic("could not open port")
}

func showInBrowser(path string) {
	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:    getAddr(),
		Handler: mux,
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, path)

		srv.Shutdown(context.Background())
	})

	go func() {
		srv.ListenAndServe()
	}()

	u := (&url.URL{
		Scheme: "http",
		Host:   srv.Addr,
		Path:   path,
	}).String()

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open", u)
	case "windows":
		cmd = exec.Command("rundll32", u)
	case "darwin":
		cmd = exec.Command("open", u)
	default:
		panic("unsupported platform")
	}

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

type testGroup uint8

// The order that tests will be grouped in.
const (
	testGroupNonInteractive testGroup = iota + 1
	testGroupNoCVM
	testGroupSignature
	testGroupMSR
	testGroupDebit
	testGroupInteractive
)
