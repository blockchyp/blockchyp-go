// +build regression

package regression

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const notEmpty = "NOT EMPTY"

type cli struct {
	path string
	t    *testing.T
}

func newCLI(t *testing.T) cli {
	c := cli{
		t: t,
	}
	flag.StringVar(&c.path, "cli", "go run ../../cmd/blockchyp", "CLI executable to invoke")
	flag.Parse()

	return c
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
	expectVal := reflect.ValueOf(expect)
	if expectVal.Kind() == reflect.Ptr {
		expectVal = expectVal.Elem()
	}

	resultVal := reflect.ValueOf(result)
	if resultVal.Kind() == reflect.Ptr {
		resultVal = resultVal.Elem()
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
			if s := expectVal.Field(i).Interface().(string); s == notEmpty {
				assert.NotEmpty(t, resultVal.Field(i).Interface(),
					fmt.Sprintf("%s should not be empty", expectVal.Type().Field(i).Name),
				)
				continue
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
	c.t.Logf("+ %s %s", c.path, strings.Join(args, " "))

	bin := strings.Split(c.path, " ")
	bin = append(bin, args...)

	cmd := exec.Command(bin[0], bin[1:]...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	cmd.Run()

	c.t.Log(string(stdout.String()))
	c.t.Log(string(stderr.String()))

	if v == nil {
		return
	}
	if err := json.Unmarshal(stdout.Bytes(), v); err != nil {
		c.t.Fatalf("Failed to unmarshal: %+v; Raw output: %s", err, string(stdout.Bytes()))
	}
}

const (
	red     = "\x1b[31m"
	green   = "\x1b[32m"
	yellow  = "\x1b[33m"
	blue    = "\x1b[34m"
	magenta = "\x1b[35m"
	cyan    = "\x1b[36m"
	noColor = "\x1b[0m"
)

func setup(t *testing.T, instructions string, pause bool) {
	if instructions == "" {
		return
	}

	fmt.Println("\nSteps: " + magenta + instructions + noColor + "\n")

	if pause {
		fmt.Println(green + "Press 'Enter' to continue" + noColor)

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

	fmt.Printf("\n%s%s y/N:%s ", "\x1b[33m", v.prompt, "\x1b[0m")

	f, err := os.Open("/dev/tty")
	if err != nil {
		t.Fatal(err)
	}
	res, _ := bufio.NewReader(f).ReadBytes('\n')

	assert.Equal(t, v.expect, strings.HasPrefix(strings.ToLower(string(res)), "y"))
}

type validation struct {
	prompt string
	expect bool
}
