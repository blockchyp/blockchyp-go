// +build regression

package regression

import (
	"testing"

	"github.com/blockchyp/blockchyp-go"
)

func TestEBT(t *testing.T) {
	tests := map[string]struct {
		instructions string
		args         []string
		assert       interface{}
	}{
		"Charge": {
			instructions: "Swipe an EBT test card when prompted. Enter PIN '1234'.",
			args: []string{
				"-type", "charge", "-terminal", "Test Terminal",
				"-test", "-amount", partialAuthAuthorizedAmount,
				"-ebt",
			},
			assert: blockchyp.AuthorizationResponse{
				Success:          true,
				Approved:         true,
				Test:             true,
				TransactionType:  "charge",
				RequestedAmount:  partialAuthAuthorizedAmount,
				AuthorizedAmount: partialAuthAuthorizedAmount,
				PaymentType:      "EBT",
				ReceiptSuggestions: blockchyp.ReceiptSuggestions{
					PINVerified: true,
				},
			},
		},
		"MSRBalance": {
			instructions: "Swipe an EBT test card when prompted. Enter PIN '1234'.",
			args: []string{
				"-type", "balance", "-terminal", "Test Terminal",
				"-test",
				"-ebt",
			},
			assert: blockchyp.BalanceResponse{
				Success:          true,
				Test:             true,
				TransactionType:  "balance",
				RemainingBalance: "100.00",
				PaymentType:      "EBT",
				ReceiptSuggestions: blockchyp.ReceiptSuggestions{
					PINVerified: true,
				},
			},
		},
		"ManualBalance": {
			instructions: "Key in the number '4111 1111 1111 1111' when prompted. Enter PIN '1234'.",
			args: []string{
				"-type", "balance", "-terminal", "Test Terminal",
				"-test",
				"-ebt", "-manual",
			},
			assert: blockchyp.BalanceResponse{
				Success:          true,
				Test:             true,
				TransactionType:  "balance",
				RemainingBalance: "100.00",
				PaymentType:      "EBT",
				ReceiptSuggestions: blockchyp.ReceiptSuggestions{
					PINVerified: true,
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			cli := newCLI(t)

			setup(t, test.instructions, true)

			cli.run(test.args, test.assert)
		})
	}
}
