// +build regression

package regression

import (
	"testing"

	"github.com/blockchyp/blockchyp-go"
)

func TestVoid(t *testing.T) {
	tests := map[string]struct {
		instructions string
		args         [][]string
		assert       []interface{}
		txID         string
	}{
		"Charge": {
			instructions: "Insert an EMV test card when prompted.",
			args: [][]string{
				{
					"-type", "charge", "-terminal", "Test Terminal", "-test",
					"-amount", "62.00",
				},
				{
					"-type", "void", "-test",
					"-tx",
				},
			},
			assert: []interface{}{
				blockchyp.AuthorizationResponse{
					Success:         true,
					Approved:        true,
					Test:            true,
					TransactionType: "charge",
				},
				blockchyp.VoidResponse{
					Success:         true,
					Approved:        true,
					Test:            true,
					TransactionType: "void",
				},
			},
		},
		"Preauth": {
			instructions: "Insert an EMV test card when prompted.",
			args: [][]string{
				{
					"-type", "preauth", "-terminal", "Test Terminal", "-test",
					"-amount", "63.00",
				},
				{
					"-type", "void", "-test",
					"-tx",
				},
			},
			assert: []interface{}{
				blockchyp.AuthorizationResponse{
					Success:         true,
					Approved:        true,
					Test:            true,
					TransactionType: "preauth",
				},
				blockchyp.VoidResponse{
					Success:         true,
					Approved:        true,
					Test:            true,
					TransactionType: "void",
				},
			},
		},
		"Capture": {
			instructions: "Insert an EMV test card when prompted.",
			args: [][]string{
				{
					"-type", "preauth", "-terminal", "Test Terminal", "-test",
					"-amount", "63.00",
				},
				{
					"-type", "capture", "-test",
					"-tx",
				},
				{
					"-type", "void", "-test",
					"-tx",
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
				blockchyp.VoidResponse{
					Success:         true,
					Approved:        true,
					Test:            true,
					TransactionType: "void",
				},
			},
		},
		"Unknown": {
			args: [][]string{
				{
					"-type", "void", "-test",
					"-tx", "NOT A REAL TRANSACTION",
				},
			},
			assert: []interface{}{
				blockchyp.VoidResponse{
					Success:             false,
					Approved:            false,
					Test:                true,
					TransactionType:     "void",
					ResponseDescription: "Invalid Transaction",
				},
			},
		},
		"Double": {
			instructions: "Insert an EMV test card when prompted.",
			args: [][]string{
				{
					"-type", "charge", "-terminal", "Test Terminal", "-test",
					"-amount", "103.00",
				},
				{
					"-type", "void", "-test",
					"-tx",
				},
				{
					"-type", "void", "-test",
					"-tx",
				},
			},
			assert: []interface{}{
				blockchyp.AuthorizationResponse{
					Success:         true,
					Approved:        true,
					Test:            true,
					TransactionType: "charge",
				},
				blockchyp.VoidResponse{
					Success:         true,
					Approved:        true,
					Test:            true,
					TransactionType: "void",
				},
				blockchyp.VoidResponse{
					Success:             false,
					Approved:            false,
					Test:                true,
					TransactionType:     "void",
					ResponseDescription: "Already Voided",
				},
			},
		},
	}

	cli := newCLI(t)

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			setup(t, test.instructions, true)

			for i := range test.args {
				if test.txID != "" {
					test.args[i] = append(test.args[i], test.txID)
				}

				res, ok := cli.run(test.args[i], test.assert[i]).(*blockchyp.AuthorizationResponse)
				if ok && test.txID == "" {
					test.txID = res.TransactionID
				}
			}
		})
	}
}
