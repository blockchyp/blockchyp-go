package regression

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

const notEmpty = "NOT_EMPTY"
const terminalName = "TERMINAL_NAME"
const previousAmount = "PREVIOUS_AMOUNT"
const newTxRef = "TX_REF"

var previousTxRef = txRefN(-1)

func txRefN(n int) string {
	return refString(txRefPrefix, n)
}

var previousTxID = txIDN(-1)

func txIDN(n int) string {
	return refString(txIDPrefix, n)
}

var previousBatchID = batchIDN(-1)

func batchIDN(n int) string {
	return refString(batchIDPrefix, n)
}

func tokenN(n int) string {
	return refString(tokenPrefix, n)
}

func customerIDN(n int) string {
	return refString(customerIDPrefix, n)
}

const (
	txRefPrefix      = "TX_REF["
	txIDPrefix       = "TX_ID["
	batchIDPrefix    = "BATCH_ID["
	tokenPrefix      = "TOKEN["
	customerIDPrefix = "CUSTOMER_ID["
)

func refString(prefix string, n int) string {
	return fmt.Sprintf(prefix+"%d]", n)
}

const defaultFlatRateFee = 0.035
const defaultPerTransactionFee = 50

type cli struct {
	path string
	t    *testing.T
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

func (app *TestRunner) setup(instructions string, pause bool) {
	if instructions == "" {
		return
	}

	fmt.Println("\nSteps: " + format(magenta) + instructions + format() + "\n")

	if pause {
		fmt.Println(format(green) + "Press 'Enter' to continue" + format())

		f, err := os.Open("/dev/tty")
		if err != nil {
			app.log.Fatal(err)
		}
		defer f.Close()
		bufio.NewReader(f).ReadBytes('\n')
	}
}

func (app *TestRunner) validate(v validation) error {
	if v.prompt == "" {
		return nil
	}

	if v.serve != "" {
		showInBrowser(v.serve)
	}

	if prompt(v.prompt, !v.expect) != v.expect {
		return fmt.Errorf("expected: %v", v.expect)
	}

	return nil
}

func prompt(msg string, def bool) bool {
	var s string
	if def {
		s = "\n%s%s Y/n:%s "
	} else {
		s = "\n%s%s y/N:%s "
	}

	fmt.Printf(s, format(yellow), msg, format())

	f, err := os.Open("/dev/tty")
	if err != nil {
		return false
	}
	defer f.Close()
	res, _ := bufio.NewReader(f).ReadBytes('\n')

	if def {
		return !strings.HasPrefix(strings.ToLower(string(res)), "n")
	}

	return strings.HasPrefix(strings.ToLower(string(res)), "y")
}

func multiprompt(choices []string, def int) int {
	fmt.Printf("\n%s%s: %s", format(yellow), strings.Join(choices, ", "), format())

	f, err := os.Open("/dev/tty")
	if err != nil {
		return def
	}
	defer f.Close()

	for {
		res, _ := bufio.NewReader(f).ReadBytes('\n')

		if len(strings.TrimSpace(string(res))) == 0 {
			return def
		}

		for i, choice := range choices {
			for _, c := range strings.ToLower(choice) {
				if 'a' <= c && c <= 'z' {
					if rune(strings.ToLower(string(res))[0]) == c {
						return i
					}
					break
				}
			}
		}
	}
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
