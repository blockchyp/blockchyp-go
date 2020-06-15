// +build regression

package regression

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/blockchyp/blockchyp-go"
	"github.com/stretchr/testify/assert"
)

const terminalName = "Test Terminal"
const notEmpty = "NOT EMPTY"

const defaultFlatRateFee = 0.035
const defaultPerTransactionFee = 50

type cli struct {
	path string
	t    *testing.T
}

var cliExecutable string
var acquirerMode bool

func init() {
	if env := os.Getenv("CLI"); env != "" {
		cliExecutable = env
	} else {
		cliExecutable = "go run ../../cmd/blockchyp"
	}

	if env := os.Getenv("MODE"); env != "" {
		switch env {
		case "acquirer":
			acquirerMode = true
		}
	}
}

func newCLI(t *testing.T) cli {
	return cli{
		t:    t,
		path: cliExecutable,
	}
}

func (c cli) run(args []string, expect interface{}) interface{} {
	var result interface{}
	if expect != nil {
		result = reflect.New(reflect.TypeOf(expect)).Interface()
	}

	c.exec(args, result)

	if result != nil {
		cmp(c.t, expect, result)
	}

	return result
}

func cmp(t *testing.T, expect, result interface{}) {
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
				assert.NotEmpty(t, resultVal.Field(i).Interface(),
					fmt.Sprintf("%s should not be empty", expectVal.Type().Field(i).Name),
				)
				continue
			}
			if strings.HasPrefix(s, replacementPrefix) {
				expectVal.Field(i).SetString(getAmount(t.Name(), s))
			}
		case reflect.Struct:
			cmp(t, expectVal.Field(i).Interface(), resultVal.Field(i).Interface())
			continue
		case reflect.Ptr:
			if expectVal.Field(i).IsNil() {
				continue
			}
			if expectVal.Field(i).Elem().Kind() == reflect.Struct {
				cmp(t, expectVal.Field(i).Interface(), resultVal.Field(i).Interface())
				continue
			}
		case reflect.Slice:
			if expectVal.Field(i).Len() == 0 {
				continue
			}
			if ln := resultVal.Field(i).Len(); ln != 1 {
				t.Fatalf("%s: expected 1 element, got %d", expectVal.Type().Field(i).Name, ln)
			}
			// This only handles cases where we compare a slice of structs with
			// one element in the assertion and the result. If we need to test
			// more than this, it will require modification.
			if expectVal.Field(i).Index(0).Kind() == reflect.Struct {
				cmp(t, expectVal.Field(i).Index(0), resultVal.Field(i).Index(0))
				continue
			}
		default:
			// For all other types, only check them if they are set
			if expectVal.Field(i).IsZero() {
				continue
			}
		}
		assert.Equal(t, expectVal.Field(i).Interface(), resultVal.Field(i).Interface(),
			fmt.Sprintf("%s should be %+v", expectVal.Type().Field(i).Name, expectVal.Field(i).Interface()))
	}
}

func (c cli) exec(args []string, v interface{}) {
	c.substituteAmounts(args)

	c.t.Logf("+ %s %s", c.path, strings.Join(args, " "))

	bin := strings.Split(c.path, " ")
	bin = append(bin, args...)

	cmd := exec.Command(bin[0], bin[1:]...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	cmd.Run()

	if s := stdout.String(); s != "" {
		c.t.Log(s)
	}
	if s := stderr.String(); s != "" {
		c.t.Log(s)
	}

	if v == nil {
		return
	}
	if err := json.Unmarshal(stdout.Bytes(), v); err != nil {
		c.t.Fatalf("Failed to unmarshal: %+v; Raw output: %s", err, string(stdout.Bytes()))
	}
}

func (c cli) substituteAmounts(args []string) {
	for i := range args {
		if strings.HasPrefix(args[i], replacementPrefix) {
			args[i] = getAmount(c.t.Name(), args[i])
		}
	}
}

func (c cli) skipCloudRelay() {
	// Make sure the cache has actually been generated first.
	c.exec([]string{"-type", "ping", "-terminal", terminalName, "-test"}, nil)

	path := filepath.Join(os.TempDir(), ".blockchyp_routes")

	f, err := os.Open(path)
	if err != nil {
		c.t.Fatal(err)
	}
	defer f.Close()

	var cache blockchyp.RouteCache
	if err := json.NewDecoder(f).Decode(&cache); err != nil {
		c.t.Fatal(err)
	}

	for k, v := range cache.Routes {
		if strings.HasSuffix(k, terminalName) && v.Route.CloudRelayEnabled {
			c.t.Skip("skipping local mode test in cloud relay mode")
			return
		}
	}
}

type style int

func (s style) String() string {
	return strconv.Itoa(int(s))
}

const (
	noColor style = iota + 30
	red
	green
	yellow
	blue
	magenta
	cyan
)

const (
	normal style = iota
	bold
	underline
	blink
)

func format(styles ...style) string {
	result := "\x1b["

	if len(styles) == 0 {
		return result + "0m"
	}

	for i, s := range styles {
		result += s.String()
		if i+1 < len(styles) {
			result += ";"
		}
	}

	return result + "m"
}

func setup(t *testing.T, instructions string, pause bool) {
	if instructions == "" {
		return
	}

	fmt.Println("\nSteps: " + format(magenta) + instructions + format() + "\n")

	if pause {
		fmt.Println(format(green) + "Press 'Enter' to continue" + format())

		f, err := os.Open("/dev/tty")
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()
		bufio.NewReader(f).ReadBytes('\n')
	}
}

func validate(t *testing.T, v validation) {
	if v.prompt == "" {
		return
	}

	if v.serve != "" {
		showInBrowser(v.serve)
	}

	fmt.Printf("\n%s%s y/N:%s ", format(yellow), v.prompt, format())

	f, err := os.Open("/dev/tty")
	if err != nil {
		t.Fatal(err)
	}
	res, _ := bufio.NewReader(f).ReadBytes('\n')

	assert.Equal(t, v.expect, strings.HasPrefix(strings.ToLower(string(res)), "y"))
}

func wait(duration time.Duration) {
	for i := duration; i > 0; i -= time.Second {
		fmt.Printf("\x1b[2K\r" + format(yellow) + "Wait " + i.String() + format())
		time.Sleep(1 * time.Second)
	}
	fmt.Printf("\n")
}

type validation struct {
	prompt string
	serve  string
	expect bool
}
