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
					"-amount", amount(0), "-token",
				},
				{
					"-type", "preauth", "-test",
					"-amount", amount(1), "-token",
				},
				{
					"-type", "refund", "-test",
					"-amount", amount(0), "-token",
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
					RequestedAmount:  amount(0),
					AuthorizedAmount: amount(0),
				},
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "preauth",
					EntryMethod:      "TOKEN",
					RequestedAmount:  amount(1),
					AuthorizedAmount: amount(1),
				},
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "refund",
					EntryMethod:      "TOKEN",
					RequestedAmount:  amount(0),
					AuthorizedAmount: amount(0),
				},
			},
		},
		"Charge": {
			instructions: "Insert an EMV test card when prompted.",
			args: [][]string{
				{
					"-type", "charge", "-terminal", "Test Terminal", "-test",
					"-amount", amount(0), "-enroll",
				},
				{
					"-type", "charge", "-test",
					"-amount", amount(1), "-token",
				},
				{
					"-type", "preauth", "-test",
					"-amount", amount(2), "-token",
				},
				{
					"-type", "refund", "-test",
					"-amount", partialAuthAuthorizedAmount, "-token",
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
					RequestedAmount:  amount(1),
					AuthorizedAmount: amount(1),
				},
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "preauth",
					EntryMethod:      "TOKEN",
					RequestedAmount:  amount(2),
					AuthorizedAmount: amount(2),
				},
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "refund",
					EntryMethod:      "TOKEN",
					RequestedAmount:  partialAuthAuthorizedAmount,
					AuthorizedAmount: partialAuthAuthorizedAmount,
				},
			},
		},
		"Preauth": {
			instructions: "Insert an EMV test card when prompted.",
			args: [][]string{
				{
					"-type", "preauth", "-terminal", "Test Terminal", "-test",
					"-amount", amount(0), "-enroll",
				},
				{
					"-type", "charge", "-test",
					"-amount", amount(1), "-token",
				},
				{
					"-type", "preauth", "-test",
					"-amount", amount(2), "-token",
				},
				{
					"-type", "refund", "-test",
					"-amount", amount(0), "-token",
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
					RequestedAmount:  amount(1),
					AuthorizedAmount: amount(1),
				},
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "preauth",
					EntryMethod:      "TOKEN",
					RequestedAmount:  amount(2),
					AuthorizedAmount: amount(2),
				},
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "refund",
					EntryMethod:      "TOKEN",
					RequestedAmount:  amount(0),
					AuthorizedAmount: amount(0),
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			cli := newCLI(t)

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
