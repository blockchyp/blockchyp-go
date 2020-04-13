// +build regression

package regression

import (
	"fmt"
	"testing"
	"time"

	"github.com/blockchyp/blockchyp-go"
)

func TestReverse(t *testing.T) {
	tests := map[string]struct {
		instructions string
		args         [][]string
		assert       []interface{}
		txRef        string
		wait         time.Duration
	}{
		"Charge": {
			instructions: "Insert an EMV test card when prompted.",
			args: [][]string{
				{
					"-type", "charge", "-terminal", "Test Terminal", "-test",
					"-amount", amount(0),
					"-txRef",
				},
				{
					"-type", "reverse", "-test",
					"-txRef",
				},
				{
					"-type", "reverse", "-test",
					"-txRef",
				},
			},
			assert: []interface{}{
				blockchyp.AuthorizationResponse{
					Success:         true,
					Approved:        true,
					Test:            true,
					TransactionType: "charge",
				},
				blockchyp.AuthorizationResponse{
					Success:         true,
					Approved:        true,
					Test:            true,
					TransactionType: "reverse",
				},
				blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					TransactionType:     "reverse",
					ResponseDescription: "Reversed",
				},
			},
		},
		"Preauth": {
			instructions: "Insert an EMV test card when prompted.",
			args: [][]string{
				{
					"-type", "preauth", "-terminal", "Test Terminal", "-test",
					"-amount", amount(0),
					"-txRef",
				},
				{
					"-type", "reverse", "-test",
					"-txRef",
				},
			},
			assert: []interface{}{
				blockchyp.AuthorizationResponse{
					Success:         true,
					Approved:        true,
					Test:            true,
					TransactionType: "preauth",
				},
				blockchyp.AuthorizationResponse{
					Success:         true,
					Approved:        true,
					Test:            true,
					TransactionType: "reverse",
				},
			},
		},
		"Capture": {
			instructions: "Insert an EMV test card when prompted.",
			args: [][]string{
				{
					"-type", "preauth", "-terminal", "Test Terminal", "-test",
					"-amount", amount(0),
					"-txRef",
				},
				{
					"-type", "capture", "-test",
					"-tx",
				},
				{
					"-type", "reverse", "-test",
					"-txRef",
				},
			},
			assert: []interface{}{
				blockchyp.AuthorizationResponse{
					Success:         true,
					Approved:        true,
					Test:            true,
					TransactionType: "preauth",
				},
				blockchyp.CaptureResponse{
					Success:         true,
					Approved:        true,
					Test:            true,
					TransactionType: "capture",
				},
				blockchyp.AuthorizationResponse{
					Success:         true,
					Approved:        true,
					Test:            true,
					TransactionType: "reverse",
				},
			},
		},
		"TimeLimit": {
			instructions: "Insert an EMV test card when prompted.",
			wait:         125 * time.Second,
			args: [][]string{
				{
					"-type", "charge", "-terminal", "Test Terminal", "-test",
					"-amount", amount(0),
					"-txRef",
				},
				{
					"-type", "reverse", "-test",
					"-txRef",
				},
			},
			assert: []interface{}{
				blockchyp.AuthorizationResponse{
					Success:         true,
					Approved:        true,
					Test:            true,
					TransactionType: "charge",
				},
				blockchyp.AuthorizationResponse{
					Success:             false,
					Approved:            false,
					Test:                true,
					TransactionType:     "reverse",
					ResponseDescription: "Reverse Time Limit Exceeded. Use Void Instead.",
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			cli := newCLI(t)

			setup(t, test.instructions, true)

			var txID string
			if test.txRef == "" {
				test.txRef = randomStr()
			}

			for i := range test.args {
				if test.args[i][len(test.args[i])-1] == "-txRef" {
					test.args[i] = append(test.args[i], test.txRef)
				} else if test.args[i][len(test.args[i])-1] == "-tx" {
					test.args[i] = append(test.args[i], txID)
				}

				if res, ok := cli.run(test.args[i], test.assert[i]).(*blockchyp.AuthorizationResponse); ok && txID == "" {
					txID = res.TransactionID
				}

				if test.wait > 0 && i == 0 {
					fmt.Println("\n" + yellow + "Wait " + test.wait.String() + " for the test to complete." + noColor)
					time.Sleep(test.wait)
				}
			}
		})
	}
}
