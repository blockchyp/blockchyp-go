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
		"DirectEMV": {
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
					"-amount", amount(2), "-token",
				},
			},
			assert: []interface{}{
				blockchyp.AuthorizationResponse{
					Success:         true,
					Approved:        true,
					Test:            true,
					TransactionType: "enroll",
					EntryMethod:     "CHIP",
					Token:           notEmpty,
					MaskedPAN:       notEmpty,
					ReceiptSuggestions: blockchyp.ReceiptSuggestions{
						MerchantName:    notEmpty,
						MaskedPAN:       notEmpty,
						TransactionType: "enroll",
						EntryMethod:     "CHIP",
					},
				},
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "charge",
					EntryMethod:      "TOKEN",
					RequestedAmount:  amount(0),
					AuthorizedAmount: amount(0),
					MaskedPAN:        notEmpty,
					ReceiptSuggestions: blockchyp.ReceiptSuggestions{
						AuthorizedAmount: amount(0),
						MerchantName:     notEmpty,
						MaskedPAN:        notEmpty,
						TransactionType:  "charge",
						EntryMethod:      "TOKEN",
					},
				},
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "preauth",
					EntryMethod:      "TOKEN",
					RequestedAmount:  amount(1),
					AuthorizedAmount: amount(1),
					MaskedPAN:        notEmpty,
					ReceiptSuggestions: blockchyp.ReceiptSuggestions{
						AuthorizedAmount: amount(1),
						MerchantName:     notEmpty,
						MaskedPAN:        notEmpty,
						TransactionType:  "preauth",
						EntryMethod:      "TOKEN",
					},
				},
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "refund",
					EntryMethod:      "TOKEN",
					RequestedAmount:  amount(2),
					AuthorizedAmount: amount(2),
					MaskedPAN:        notEmpty,
					ReceiptSuggestions: blockchyp.ReceiptSuggestions{
						AuthorizedAmount: amount(2),
						MerchantName:     notEmpty,
						MaskedPAN:        notEmpty,
						TransactionType:  "refund",
						EntryMethod:      "TOKEN",
					},
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
					"-amount", amount(3), "-token",
				},
			},
			assert: []interface{}{
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "charge",
					EntryMethod:      "CHIP",
					RequestedAmount:  amount(0),
					AuthorizedAmount: amount(0),
					Token:            notEmpty,
					MaskedPAN:        notEmpty,
					ReceiptSuggestions: blockchyp.ReceiptSuggestions{
						AuthorizedAmount: amount(0),
						MerchantName:     notEmpty,
						MaskedPAN:        notEmpty,
						TransactionType:  "charge",
						EntryMethod:      "CHIP",
					},
				},
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "charge",
					EntryMethod:      "TOKEN",
					RequestedAmount:  amount(1),
					AuthorizedAmount: amount(1),
					MaskedPAN:        notEmpty,
					ReceiptSuggestions: blockchyp.ReceiptSuggestions{
						AuthorizedAmount: amount(1),
						MerchantName:     notEmpty,
						MaskedPAN:        notEmpty,
						TransactionType:  "charge",
						EntryMethod:      "TOKEN",
					},
				},
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "preauth",
					EntryMethod:      "TOKEN",
					RequestedAmount:  amount(2),
					AuthorizedAmount: amount(2),
					MaskedPAN:        notEmpty,
					ReceiptSuggestions: blockchyp.ReceiptSuggestions{
						AuthorizedAmount: amount(2),
						MerchantName:     notEmpty,
						MaskedPAN:        notEmpty,
						TransactionType:  "preauth",
						EntryMethod:      "TOKEN",
					},
				},
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "refund",
					EntryMethod:      "TOKEN",
					RequestedAmount:  amount(3),
					AuthorizedAmount: amount(3),
					MaskedPAN:        notEmpty,
					ReceiptSuggestions: blockchyp.ReceiptSuggestions{
						AuthorizedAmount: amount(3),
						MerchantName:     notEmpty,
						MaskedPAN:        notEmpty,
						TransactionType:  "refund",
						EntryMethod:      "TOKEN",
					},
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
					"-amount", amount(3), "-token",
				},
			},
			assert: []interface{}{
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "preauth",
					EntryMethod:      "CHIP",
					RequestedAmount:  amount(0),
					AuthorizedAmount: amount(0),
					Token:            notEmpty,
					MaskedPAN:        notEmpty,
					ReceiptSuggestions: blockchyp.ReceiptSuggestions{
						AuthorizedAmount: amount(0),
						MerchantName:     notEmpty,
						MaskedPAN:        notEmpty,
						TransactionType:  "preauth",
						EntryMethod:      "CHIP",
					},
				},
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "charge",
					EntryMethod:      "TOKEN",
					RequestedAmount:  amount(1),
					AuthorizedAmount: amount(1),
					MaskedPAN:        notEmpty,
					ReceiptSuggestions: blockchyp.ReceiptSuggestions{
						AuthorizedAmount: amount(1),
						MerchantName:     notEmpty,
						MaskedPAN:        notEmpty,
						TransactionType:  "charge",
						EntryMethod:      "TOKEN",
					},
				},
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "preauth",
					EntryMethod:      "TOKEN",
					RequestedAmount:  amount(2),
					AuthorizedAmount: amount(2),
					MaskedPAN:        notEmpty,
					ReceiptSuggestions: blockchyp.ReceiptSuggestions{
						AuthorizedAmount: amount(2),
						MerchantName:     notEmpty,
						MaskedPAN:        notEmpty,
						TransactionType:  "preauth",
						EntryMethod:      "TOKEN",
					},
				},
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "refund",
					EntryMethod:      "TOKEN",
					RequestedAmount:  amount(3),
					AuthorizedAmount: amount(3),
					MaskedPAN:        notEmpty,
					ReceiptSuggestions: blockchyp.ReceiptSuggestions{
						AuthorizedAmount: amount(3),
						MerchantName:     notEmpty,
						MaskedPAN:        notEmpty,
						TransactionType:  "refund",
						EntryMethod:      "TOKEN",
					},
				},
			},
		},
		"DirectMSR": {
			instructions: "Swipe an MSR test card when prompted.",
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
					"-amount", amount(2), "-token",
				},
			},
			assert: []interface{}{
				blockchyp.AuthorizationResponse{
					Success:         true,
					Approved:        true,
					Test:            true,
					TransactionType: "enroll",
					EntryMethod:     "SWIPE",
					Token:           notEmpty,
					MaskedPAN:       notEmpty,
					ReceiptSuggestions: blockchyp.ReceiptSuggestions{
						MerchantName:     notEmpty,
						MaskedPAN:        notEmpty,
						TransactionType:  "enroll",
						EntryMethod:      "SWIPE",
						RequestSignature: false,
					},
				},
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "charge",
					EntryMethod:      "TOKEN",
					RequestedAmount:  amount(0),
					AuthorizedAmount: amount(0),
					MaskedPAN:        notEmpty,
					ReceiptSuggestions: blockchyp.ReceiptSuggestions{
						AuthorizedAmount: amount(0),
						MerchantName:     notEmpty,
						MaskedPAN:        notEmpty,
						TransactionType:  "charge",
						EntryMethod:      "TOKEN",
					},
				},
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "preauth",
					EntryMethod:      "TOKEN",
					RequestedAmount:  amount(1),
					AuthorizedAmount: amount(1),
					MaskedPAN:        notEmpty,
					ReceiptSuggestions: blockchyp.ReceiptSuggestions{
						AuthorizedAmount: amount(1),
						MerchantName:     notEmpty,
						MaskedPAN:        notEmpty,
						TransactionType:  "preauth",
						EntryMethod:      "TOKEN",
					},
				},
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "refund",
					EntryMethod:      "TOKEN",
					RequestedAmount:  amount(2),
					AuthorizedAmount: amount(2),
					MaskedPAN:        notEmpty,
					ReceiptSuggestions: blockchyp.ReceiptSuggestions{
						AuthorizedAmount: amount(2),
						MerchantName:     notEmpty,
						MaskedPAN:        notEmpty,
						TransactionType:  "refund",
						EntryMethod:      "TOKEN",
					},
				},
			},
		},
		"DirectManual": {
			instructions: "Enter PAN '4111 1111 1111 1111' and CVV2 '1234' when prompted",
			args: [][]string{
				{
					"-type", "enroll", "-terminal", "Test Terminal", "-test",
					"-manual",
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
					"-amount", amount(2), "-token",
				},
			},
			assert: []interface{}{
				blockchyp.AuthorizationResponse{
					Success:         true,
					Approved:        true,
					Test:            true,
					TransactionType: "enroll",
					EntryMethod:     "KEYED",
					Token:           notEmpty,
					MaskedPAN:       notEmpty,
					ReceiptSuggestions: blockchyp.ReceiptSuggestions{
						MerchantName:    notEmpty,
						MaskedPAN:       notEmpty,
						TransactionType: "enroll",
						EntryMethod:     "KEYED",
					},
				},
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "charge",
					EntryMethod:      "TOKEN",
					RequestedAmount:  amount(0),
					AuthorizedAmount: amount(0),
					MaskedPAN:        notEmpty,
					ReceiptSuggestions: blockchyp.ReceiptSuggestions{
						AuthorizedAmount: amount(0),
						MerchantName:     notEmpty,
						MaskedPAN:        notEmpty,
						TransactionType:  "charge",
						EntryMethod:      "TOKEN",
					},
				},
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "preauth",
					EntryMethod:      "TOKEN",
					RequestedAmount:  amount(1),
					AuthorizedAmount: amount(1),
					MaskedPAN:        notEmpty,
					ReceiptSuggestions: blockchyp.ReceiptSuggestions{
						AuthorizedAmount: amount(1),
						MerchantName:     notEmpty,
						MaskedPAN:        notEmpty,
						TransactionType:  "preauth",
						EntryMethod:      "TOKEN",
					},
				},
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "refund",
					EntryMethod:      "TOKEN",
					RequestedAmount:  amount(2),
					AuthorizedAmount: amount(2),
					MaskedPAN:        notEmpty,
					ReceiptSuggestions: blockchyp.ReceiptSuggestions{
						AuthorizedAmount: amount(2),
						MerchantName:     notEmpty,
						MaskedPAN:        notEmpty,
						TransactionType:  "refund",
						EntryMethod:      "TOKEN",
					},
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
