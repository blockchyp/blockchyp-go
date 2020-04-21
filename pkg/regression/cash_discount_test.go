// +build regression

package regression

import (
	"testing"

	"github.com/blockchyp/blockchyp-go"
)

func TestCashDiscount(t *testing.T) {
	tests := map[string]struct {
		instructions string
		args         [][]string
		assert       []interface{}
	}{
		"CreditSurchargeCashDiscount": {
			instructions: "Insert a test EMV " + format(bold, yellow) + "credit" + format(normal, magenta) + " card when prompted.",
			args: [][]string{
				{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", amount(0),
					"-surcharge", "-cashDiscount",
				},
			},
			assert: []interface{}{
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "charge",
					EntryMethod:      "CHIP",
					RequestedAmount:  addFees(amount(0)),
					AuthorizedAmount: addFees(amount(0)),
					MaskedPAN:        notEmpty,
					ReceiptSuggestions: blockchyp.ReceiptSuggestions{
						MerchantName:    notEmpty,
						MaskedPAN:       notEmpty,
						TransactionType: "charge",
						EntryMethod:     "CHIP",
						// Surcharge is applied and no discount is applied
						AuthorizedAmount: addFees(amount(0)),
						Surcharge:        fees(amount(0)),
						CashDiscount:     "0.00",
					},
				},
			},
		},
		"CreditSurchargeNoCashDiscount": {
			instructions: "Insert a test EMV " + format(bold, yellow) + "credit" + format(normal, magenta) + " card when prompted.",
			args: [][]string{
				{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", amount(0),
					"-surcharge",
				},
			},
			assert: []interface{}{
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "charge",
					EntryMethod:      "CHIP",
					RequestedAmount:  addFees(amount(0)),
					AuthorizedAmount: addFees(amount(0)),
					MaskedPAN:        notEmpty,
					ReceiptSuggestions: blockchyp.ReceiptSuggestions{
						MerchantName:    notEmpty,
						MaskedPAN:       notEmpty,
						TransactionType: "charge",
						EntryMethod:     "CHIP",
						// Surcharge is applied and no discount is applied
						AuthorizedAmount: addFees(amount(0)),
						Surcharge:        fees(amount(0)),
						CashDiscount:     "0.00",
					},
				},
			},
		},
		"CreditNoSurchargeCashDiscount": {
			instructions: "Insert a test EMV " + format(bold, yellow) + "credit" + format(normal, magenta) + " card when prompted.",
			args: [][]string{
				{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", amount(0),
					"-cashDiscount",
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
					MaskedPAN:        notEmpty,
					ReceiptSuggestions: blockchyp.ReceiptSuggestions{
						MerchantName:    notEmpty,
						MaskedPAN:       notEmpty,
						TransactionType: "charge",
						EntryMethod:     "CHIP",
						// Surcharge is not applied
						AuthorizedAmount: amount(0),
						Surcharge:        "0.00",
						CashDiscount:     "0.00",
					},
				},
			},
		},
		"DebitSurchargeCashDiscount": {
			instructions: "Insert a test EMV " + format(bold, yellow) + "debit" + format(normal, magenta) + " card when prompted.",
			args: [][]string{
				{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", amount(0),
					"-surcharge", "-cashDiscount",
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
					MaskedPAN:        notEmpty,
					ReceiptSuggestions: blockchyp.ReceiptSuggestions{
						MerchantName:    notEmpty,
						MaskedPAN:       notEmpty,
						TransactionType: "charge",
						EntryMethod:     "CHIP",
						// Surcharge and discount offset
						AuthorizedAmount: amount(0),
						Surcharge:        fees(amount(0)),
						CashDiscount:     fees(amount(0)),
					},
				},
			},
		},
		"DebitSurchargeNoCashDiscount": {
			instructions: "Insert a test EMV " + format(bold, yellow) + "debit" + format(normal, magenta) + " card when prompted.",
			args: [][]string{
				{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", amount(0),
					"-surcharge",
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
					MaskedPAN:        notEmpty,
					ReceiptSuggestions: blockchyp.ReceiptSuggestions{
						MerchantName:    notEmpty,
						MaskedPAN:       notEmpty,
						TransactionType: "charge",
						EntryMethod:     "CHIP",
						// Surcharge is not applied
						AuthorizedAmount: amount(0),
						Surcharge:        "0.00",
						CashDiscount:     "0.00",
					},
				},
			},
		},
		"DebitNoSurchargeCashDiscount": {
			instructions: "Insert a test EMV " + format(bold, yellow) + "debit" + format(normal, magenta) + " card when prompted.",
			args: [][]string{
				{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", amount(0),
					"-cashDiscount",
				},
			},
			assert: []interface{}{
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "charge",
					EntryMethod:      "CHIP",
					RequestedAmount:  cashDiscount(amount(0)),
					AuthorizedAmount: cashDiscount(amount(0)),
					MaskedPAN:        notEmpty,
					ReceiptSuggestions: blockchyp.ReceiptSuggestions{
						MerchantName:    notEmpty,
						MaskedPAN:       notEmpty,
						TransactionType: "charge",
						EntryMethod:     "CHIP",
						// Cash discount is applied, fees are not applied
						AuthorizedAmount: cashDiscount(amount(0)),
						Surcharge:        "0.00",
						CashDiscount:     fees(amount(0)),
					},
				},
			},
		},
		"APISurchargeCashDiscount": {
			args: [][]string{
				{
					"-type", "cash-discount", "-test",
					"-amount", amount(0),
					"-surcharge", "-cashDiscount",
				},
			},
			assert: []interface{}{
				blockchyp.CashDiscountResponse{
					Success:      true,
					Amount:       amount(0),
					Surcharge:    fees(amount(0)),
					CashDiscount: fees(amount(0)),
				},
			},
		},
		"APISurchargeNoCashDiscount": {
			args: [][]string{
				{
					"-type", "cash-discount", "-test",
					"-amount", amount(0),
					"-surcharge",
				},
			},
			assert: []interface{}{
				blockchyp.CashDiscountResponse{
					Success: true,
					Amount:  amount(0),
				},
			},
		},
		"APINoSurchargeCashDiscount": {
			args: [][]string{
				{
					"-type", "cash-discount", "-test",
					"-amount", amount(0),
					"-cashDiscount",
				},
			},
			assert: []interface{}{
				blockchyp.CashDiscountResponse{
					Success:      true,
					Amount:       cashDiscount(amount(0)),
					CashDiscount: fees(amount(0)),
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			cli := newCLI(t)

			setup(t, test.instructions, true)

			for i := range test.args {
				cli.run(test.args[i], test.assert[i])
			}
		})
	}
}
