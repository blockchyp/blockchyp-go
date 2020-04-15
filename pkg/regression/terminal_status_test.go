// +build regression

package regression

import (
	"sync"
	"testing"
	"time"

	"github.com/blockchyp/blockchyp-go"
)

func TestStatus(t *testing.T) {
	tests := map[string]struct {
		args   [][]string
		assert []interface{}
	}{
		"Idle": {
			args: [][]string{
				{
					"-type", "terminal-status", "-terminal", terminalName,
				},
			},
			assert: []interface{}{
				blockchyp.TerminalStatusResponse{
					Success: true,
					Idle:    true,
				},
			},
		},
		"Charge": {
			args: [][]string{
				{
					"-type", "charge", "-terminal", terminalName, "-test", "-amount", "1.00",
				},
				{
					"-type", "terminal-status", "-terminal", terminalName,
				},
				{
					"-type", "clear", "-terminal", terminalName,
				},
			},
			assert: []interface{}{
				nil,
				blockchyp.TerminalStatusResponse{
					Success: true,
					Idle:    false,
					Status:  "charge",
				},
				blockchyp.Acknowledgement{
					Success: true,
				},
			},
		},
		"TC": {
			args: [][]string{
				{
					"-type", "tc", "-terminal", terminalName, "-test",
					"-tcName", "Contract Title", "-tcContent", "Blah Blah Blah",
				},
				{
					"-type", "terminal-status", "-terminal", terminalName,
				},
				{
					"-type", "clear", "-terminal", terminalName,
				},
			},
			assert: []interface{}{
				nil,
				blockchyp.TerminalStatusResponse{
					Success: true,
					Idle:    false,
					Status:  "terms-and-conditions-prompt",
				},
				blockchyp.Acknowledgement{
					Success: true,
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			cli := newCLI(t)

			setup(t, "", false)

			var wg sync.WaitGroup

			for i := range test.args {
				wg.Add(1)
				go func(i int) {
					cli.run(test.args[i], test.assert[i])
					wg.Done()
				}(i)
				time.Sleep(1 * time.Second)
			}

			wg.Wait()
		})
	}
}
