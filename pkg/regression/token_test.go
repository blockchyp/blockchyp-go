// +build regression

package regression

import (
	"testing"

	"github.com/blockchyp/blockchyp-go"
)

func TestToken(t *testing.T) {
	tests := map[string]struct {
		instructions string
		args         [][]string
		assert       []interface{}
		token        string
	}{
		"Direct": {
			instructions: "Insert an EMV test card when prompted.",
			args: [][]string{
				{
					"-type", "enroll", "-terminal", "Test Terminal", "-test",
				},
				{
					"-type", "charge", "-test",
					"-amount", "41.11", "-token",
				},
				{
					"-type", "preauth", "-test",
					"-amount", "41.12", "-token",
				},
				{
					"-type", "refund", "-test",
					"-amount", "41.09", "-token",
				},
			},
			assert: []interface{}{
				blockchyp.AuthorizationResponse{
					Success:         true,
					Approved:        true,
					Test:            true,
					TransactionType: "enroll",
					Token:           notEmpty,
				},
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "charge",
					EntryMethod:      "TOKEN",
					RequestedAmount:  "41.11",
					AuthorizedAmount: "41.11",
				},
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "preauth",
					EntryMethod:      "TOKEN",
					RequestedAmount:  "41.12",
					AuthorizedAmount: "41.12",
				},
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "refund",
					EntryMethod:      "TOKEN",
					RequestedAmount:  "41.09",
					AuthorizedAmount: "41.09",
				},
			},
		},
		"Charge": {
			instructions: "Insert an EMV test card when prompted.",
			args: [][]string{
				{
					"-type", "charge", "-terminal", "Test Terminal", "-test",
					"-amount", "25.90", "-enroll",
				},
				{
					"-type", "charge", "-test",
					"-amount", "25.91", "-token",
				},
				{
					"-type", "preauth", "-test",
					"-amount", "25.92", "-token",
				},
				{
					"-type", "refund", "-test",
					"-amount", "25.00", "-token",
				},
			},
			assert: []interface{}{
				blockchyp.AuthorizationResponse{
					Success:         true,
					Approved:        true,
					Test:            true,
					TransactionType: "charge",
					Token:           notEmpty,
				},
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "charge",
					EntryMethod:      "TOKEN",
					RequestedAmount:  "25.91",
					AuthorizedAmount: "25.91",
				},
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "preauth",
					EntryMethod:      "TOKEN",
					RequestedAmount:  "25.92",
					AuthorizedAmount: "25.92",
				},
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "refund",
					EntryMethod:      "TOKEN",
					RequestedAmount:  "25.00",
					AuthorizedAmount: "25.00",
				},
			},
		},
		"Preauth": {
			instructions: "Insert an EMV test card when prompted.",
			args: [][]string{
				{
					"-type", "preauth", "-terminal", "Test Terminal", "-test",
					"-amount", "26.90", "-enroll",
				},
				{
					"-type", "charge", "-test",
					"-amount", "26.91", "-token",
				},
				{
					"-type", "preauth", "-test",
					"-amount", "26.92", "-token",
				},
				{
					"-type", "refund", "-test",
					"-amount", "26.00", "-token",
				},
			},
			assert: []interface{}{
				blockchyp.AuthorizationResponse{
					Success:         true,
					Approved:        true,
					Test:            true,
					TransactionType: "preauth",
					Token:           notEmpty,
				},
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "charge",
					EntryMethod:      "TOKEN",
					RequestedAmount:  "26.91",
					AuthorizedAmount: "26.91",
				},
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "preauth",
					EntryMethod:      "TOKEN",
					RequestedAmount:  "26.92",
					AuthorizedAmount: "26.92",
				},
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "refund",
					EntryMethod:      "TOKEN",
					RequestedAmount:  "26.00",
					AuthorizedAmount: "26.00",
				},
			},
		},
	}

	cli := newCLI(t)

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			setup(t, test.instructions, true)

			for i := range test.args {
				if i > 0 && test.token != "" {
					test.args[i] = append(test.args[i], test.token)
				}

				res := cli.run(test.args[i], test.assert[i]).(*blockchyp.AuthorizationResponse)

				if test.token == "" {
					test.token = res.Token
				}
			}
		})
	}
}
