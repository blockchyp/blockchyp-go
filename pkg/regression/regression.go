package regression

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/blockchyp/blockchyp-go"
)

// TestRunner runs the regression test suite.
type TestRunner struct {
	CLI        string
	Acquirer   bool
	CloudRelay bool
	TestCase   string
	Terminal   string

	log *log.Logger
}

// NewTestRunner creates a new RegressionTestRunner by parsing
// command line flags.
func NewTestRunner() *TestRunner {
	var app TestRunner
	app.log = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

	flag.StringVar(&app.CLI, "cli", "go run ./cmd/blockchyp", "CLI executable")
	flag.BoolVar(&app.Acquirer, "acquirer", false, "execute in acquirer-only mode")
	flag.BoolVar(&app.CloudRelay, "cloud-relay", false, "execute in cloud-relay-only mode")
	flag.StringVar(&app.TestCase, "run", "", "regex pattern matching the test cases to run")
	flag.StringVar(&app.Terminal, "terminal", "Test Terminal", "name of the terminal to test against")

	flag.Parse()

	return &app
}

// Run runs the test cases.
func (app *TestRunner) Run() error {
	var tests testCases
	tests = append(tests, asyncTests...)
	tests = append(tests, batchTests...)
	tests = append(tests, cashDiscountTests...)
	tests = append(tests, chargeTests...)
	tests = append(tests, customerTests...)
	tests = append(tests, duplicateTests...)
	tests = append(tests, ebtTests...)
	tests = append(tests, failureModeTests...)
	tests = append(tests, giftCardTests...)
	tests = append(tests, interactionTests...)
	tests = append(tests, preauthTests...)
	tests = append(tests, refundTests...)
	tests = append(tests, reversalTests...)
	tests = append(tests, terminalPingTests...)
	tests = append(tests, terminalStatusTests...)
	tests = append(tests, tipTests...)
	tests = append(tests, tokenTests...)
	tests = append(tests, voidTests...)
	tests.sort()

	os.MkdirAll("/tmp/blockchyp-regression-test", 0755)

	var last testGroup

TEST:
	for _, test := range tests {
		if match, _ := regexp.MatchString(app.TestCase, test.name); !match {
			continue
		}

		pauseForSetup := shouldPauseForSetup(test, last)

		// Deep copy to re-generate args for every test run.
		ops := copyOps(test.operations)

	RETRY:
		for {
			test.operations = copyOps(ops)

			if err := app.runTest(test, pauseForSetup); err != nil {
				app.log.Printf("%sFAIL: %s%s: %+v", format(red), test.name, format(), err)

				switch multiprompt([]string{"[R]entry", "(S)kip", "(A)bort"}, 0) {
				case 0:
					pauseForSetup = false
					continue RETRY
				case 1:
					continue TEST
				case 2:
					return errors.New("test run failed")
				}
			}

			break RETRY
		}

		last = test.group
	}

	return nil
}

func copyOps(v []operation) []operation {
	ops := append([]operation{}, v...)
	for i := range ops {
		ops[i].args = append([]string{}, ops[i].args...)
	}

	return ops
}

func (app *TestRunner) runTest(test testCase, pauseForSetup bool) error {
	if test.local && app.CloudRelay {
		app.log.Printf("SKIP: %s", test.name)
		return nil
	}
	if test.sim && app.Acquirer {
		app.log.Printf("SKIP: %s", test.name)
		return nil
	}

	app.log.Printf("%sRUN:  %s%s", format(cyan), test.name, format())

	delete(testCache, test.name)

	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i := range test.operations {
		app.substituteConstants(&test, i)

		if test.operations[i].background {
			wg.Add(1)
			go func(i int) {
				app.runOperation(ctx, &test, i, false, false)
				wg.Done()
			}(i)
			continue
		}

		if i > 0 {
			// Always show prompts for steps after the first one.
			pauseForSetup = true
		}

		if _, err := app.runOperation(ctx, &test, i, pauseForSetup, false); err != nil {
			cancel()
			return err
		}
	}

	wg.Wait()

	app.log.Printf("%sPASS: %s%s", format(green), test.name, format())

	return nil
}

func (app *TestRunner) runOperation(ctx context.Context, test *testCase, i int, pauseForSetup, recurse bool) (interface{}, error) {
	op := test.operations[i]

	if op.timeout > 0 && !recurse {
		var res interface{}
		var err error
		done := make(chan struct{})
		go func() {
			res, err = app.runOperation(ctx, test, i, pauseForSetup, true)
			close(done)
		}()

		timer := time.NewTimer(op.timeout)
		select {
		case <-timer.C:
			return nil, errors.New("timed out")
		case <-done:
			return res, err
		}
	}

	if op.wait > 0 {
		wait(op.wait)
	}

	if op.validation != nil {
		if err := app.validate(*op.validation); err != nil {
			return nil, err
		}
	}

	if op.msg != "" {
		app.setup(op.msg, pauseForSetup)
	}

	var result interface{}
	if op.expect != nil {
		result = reflect.New(reflect.TypeOf(op.expect)).Interface()
	}

	if len(op.args) > 0 {
		app.exec(ctx, test.name, op.args, result)
	}

	if op.fn != nil {
		if err := op.fn(); err != nil {
			return nil, err
		}
	}

	if result != nil {
		if err := cmp(test.name, op.expect, result); err != nil {
			return nil, err
		}

		if txID := reflect.ValueOf(result).Elem().FieldByName("TransactionID"); txID.IsValid() {
			test.operations[i].txID = txID.String()
		}
		if batchID := reflect.ValueOf(result).Elem().FieldByName("BatchID"); batchID.IsValid() {
			test.operations[i].batchID = batchID.String()
		}
		if customer := reflect.ValueOf(result).Elem().FieldByName("Customer"); customer.IsValid() && !customer.IsZero() {
			if customer.Kind() == reflect.Ptr {
				customer = customer.Elem()
			}
			if id := customer.FieldByName("ID"); id.IsValid() {
				test.operations[i].customerID = id.String()
			}
		}
		if token := reflect.ValueOf(result).Elem().FieldByName("Token"); token.IsValid() {
			test.operations[i].token = token.String()
		}
	}

	return result, nil
}

// testCase contains a single set of transactions and assertions.
type testCase struct {
	name  string
	group testGroup
	local bool
	sim   bool

	operations []operation
}

type testCases []testCase

func (t testCases) sort() {
	sort.SliceStable(t, func(i, j int) bool {
		return t[i].group < t[j].group
	})
}

// operation contains a single test operation and its assertions.
type operation struct {
	// timeout is a timeout for the entire operation.
	timeout time.Duration

	// wait is a delay before performing the operation.
	wait time.Duration

	// args are passed to the CLI.
	args []string

	// fn is an arbitrary function to call.
	fn func() error

	// returned by responses and used in subsequent tests.
	txID       string
	batchID    string
	customerID string
	token      string

	// expect is compared to the response from the CLI.
	expect interface{}

	// validation prompts the user for manual assertions.
	validation *validation

	// msg prints a msg for the user and pauses.
	msg string

	// background runs the operation without blocking.
	background bool
}

func (o *operation) setTXID(s string) {
	o.setArg("tx", s)
}

func (o *operation) setTXRef(s string) {
	o.setArg("txRef", s)
}

func (o *operation) setBatchID(s string) {
	o.setArg("batchId", s)
}

func (o *operation) setArg(arg, value string) {
	if arg[0] != '-' {
		arg = "-" + arg
	}

	for i := range o.args {
		if o.args[i] == arg || strings.HasPrefix(o.args[i], arg+"=") {
			o.args[i] = arg + "=" + value
		}
	}
}

func cmp(name string, expect, result interface{}) error {
	resultVal := reflect.ValueOf(result)
	if resultVal.Kind() == reflect.Ptr {
		resultVal = resultVal.Elem()
	}

	// The expected value needs to be a pointer so we can swap in the
	// calculated amounts.
	expectVal := reflect.ValueOf(expect)
	if expectVal.Kind() == reflect.Ptr {
		expectVal = expectVal.Elem()
	} else {
		expectVal = reflect.New(expectVal.Type()).Elem()
		expectVal.Set(reflect.ValueOf(expect))
	}

	for i := 0; i < expectVal.NumField(); i++ {
		if !expectVal.Field(i).CanInterface() {
			continue
		}

		switch expectVal.Field(i).Kind() {
		case reflect.Bool:
		// Check booleans even if they are falsey
		case reflect.String:
			if expectVal.Field(i).IsZero() {
				continue
			}
			s := expectVal.Field(i).Interface().(string)
			if s == notEmpty {
				if resultVal.Field(i).IsZero() {
					return fmt.Errorf("%s should not be empty", expectVal.Type().Field(i).Name)
				}
				continue
			}
			if strings.HasPrefix(s, replacementPrefix) {
				expectVal.Field(i).SetString(getAmount(name, s))
			}
		case reflect.Struct:
			if err := cmp(name, expectVal.Field(i).Interface(), resultVal.Field(i).Interface()); err != nil {
				return err
			}
			continue
		case reflect.Ptr:
			if expectVal.Field(i).IsNil() {
				continue
			}
			if expectVal.Field(i).Elem().Kind() == reflect.Struct {
				if err := cmp(name, expectVal.Field(i).Interface(), resultVal.Field(i).Interface()); err != nil {
					return err
				}
				continue
			}
		case reflect.Slice:
			if expectVal.Field(i).Len() == 0 {
				continue
			}
			if ln := resultVal.Field(i).Len(); ln != 1 {
				return fmt.Errorf("%s: expected 1 element, got %d", expectVal.Type().Field(i).Name, ln)
			}
			// This only handles cases where we compare a slice of structs with
			// one element in the assertion and the result. If we need to test
			// more than this, it will require modification.
			if expectVal.Field(i).Index(0).Kind() == reflect.Struct {
				if err := cmp(name, expectVal.Field(i).Index(0), resultVal.Field(i).Index(0)); err != nil {
					return err
				}
				continue
			}
		default:
			// For all other types, only check them if they are set
			if expectVal.Field(i).IsZero() {
				continue
			}
		}

		if !reflect.DeepEqual(expectVal.Field(i).Interface(), resultVal.Field(i).Interface()) {
			return fmt.Errorf("%s should be %+v", expectVal.Type().Field(i).Name, expectVal.Field(i).Interface())
		}
	}

	return nil
}

func (app *TestRunner) exec(ctx context.Context, name string, args []string, v interface{}) {
	app.printCmd(app.CLI, args)

	bin := strings.Split(app.CLI, " ")
	bin = append(bin, args...)

	cmd := exec.Command(bin[0], bin[1:]...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	cmd.Start()

	done := make(chan struct{})
	go func() {
		cmd.Wait()
		close(done)
	}()
	select {
	case <-ctx.Done():
		return
	case <-done:
	}

	if s := stdout.String(); s != "" {
		app.log.Print(s)
	}
	if s := stderr.String(); s != "" {
		app.log.Print(s)
	}

	if v == nil {
		return
	}
	if err := json.Unmarshal(stdout.Bytes(), v); err != nil {
		app.log.Fatalf("Failed to unmarshal: %+v; Raw output: %s", err, string(stdout.Bytes()))
	}
}

func (app *TestRunner) printCmd(bin string, args []string) {
	cpy := make([]string, len(args))
	copy(cpy, args)
	for i := range cpy {
		if strings.Contains(cpy[i], " ") {
			cpy[i] = "'" + cpy[i] + "'"
		}
	}
	app.log.Printf("+ %s %s", app.CLI, strings.Join(cpy, " "))
}

func (app *TestRunner) substituteConstants(test *testCase, i int) {
	for j, arg := range test.operations[i].args {
		switch {
		case arg == terminalName:
			test.operations[i].args[j] = app.Terminal
		case arg == previousAmount:
			if i > 0 {
				test.operations[i].args[j] = argFrom("-amount", test.operations[i-1].args)
			}
		case arg == newTxRef:
			test.operations[i].args[j] = randomStr()
		case strings.HasPrefix(arg, txRefPrefix):
			op := opN(test.operations[i].args[j], i, test.operations)
			test.operations[i].args[j] = argFrom("-txRef", op.args)
		case strings.HasPrefix(arg, txIDPrefix):
			op := opN(test.operations[i].args[j], i, test.operations)
			test.operations[i].args[j] = op.txID
		case strings.HasPrefix(arg, batchIDPrefix):
			op := opN(test.operations[i].args[j], i, test.operations)
			test.operations[i].args[j] = op.batchID
		case strings.HasPrefix(arg, tokenPrefix):
			op := opN(test.operations[i].args[j], i, test.operations)
			test.operations[i].args[j] = op.token
		case strings.HasPrefix(arg, customerIDPrefix):
			op := opN(test.operations[i].args[j], i, test.operations)
			test.operations[i].args[j] = op.customerID
		case strings.HasPrefix(arg, replacementPrefix):
			test.operations[i].args[j] = getAmount(test.name, arg)
		}
	}
}

func opN(arg string, i int, ops []operation) operation {
	idx := strings.Index(arg, "[")
	s := arg[idx+1 : len(arg)-1]
	n, _ := strconv.Atoi(s)
	if n >= 0 {
		return ops[n]
	}

	return ops[i+n]
}

func argFrom(arg string, args []string) string {
	for i, s := range args {
		if s == arg {
			if i < len(args)-1 {
				return args[i+1]
			}
		}
		if strings.HasPrefix(s, arg+"=") {
			return strings.TrimPrefix(s, arg+"=")
		}
	}

	return ""
}

func (app *TestRunner) isCloudRelay() bool {
	// Make sure the cache has actually been generated first.
	app.exec(context.Background(), "", []string{"-type", "ping", "-terminal", app.Terminal, "-test"}, nil)

	path := filepath.Join(os.TempDir(), ".blockchyp_routes")

	f, err := os.Open(path)
	if err != nil {
		app.log.Fatal(err)
	}
	defer f.Close()

	var cache blockchyp.RouteCache
	if err := json.NewDecoder(f).Decode(&cache); err != nil {
		app.log.Fatal(err)
	}

	for k, v := range cache.Routes {
		if strings.HasSuffix(k, app.Terminal) && v.Route.CloudRelayEnabled {
			return true
		}
	}

	return false
}

func shouldPauseForSetup(test testCase, last testGroup) bool {
	// For the following test categories:
	// testGroupNonInteractive, testGroupNoCVM, testGroupSignature
	// Skip showing setup instructions for tests after the first in a series.
	//
	// For all other test groups:
	// Show setup steps for every test.

	switch test.group {
	case testGroupNonInteractive, testGroupNoCVM, testGroupSignature:
	default:
		return true
	}

	return last == 0 || test.group != last
}
