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
		assert       blockchyp.AuthorizationResponse
	}{
		"Charge": {
			instructions: "Swipe an EBT test card when prompted. Enter PIN '1234'.",
			args: []string{
				"-type", "charge", "-terminal", "Test Terminal",
				"-test", "-amount", "25.00",
				"-ebt",
			},
			assert: blockchyp.AuthorizationResponse{
				Success:          true,
				Approved:         true,
				Test:             true,
				TransactionType:  "charge",
				RequestedAmount:  "25.00",
				AuthorizedAmount: "25.00",
				PaymentType:      "EBT",
				ReceiptSuggestions: blockchyp.ReceiptSuggestions{
					PINVerified: true,
				},
			},
		},
		"Balance": {
			instructions: "Swipe an EBT test card when prompted. Enter PIN '1234'.",
			args: []string{
				"-type", "balance", "-terminal", "Test Terminal",
				"-test",
				"-ebt",
			},
			assert: blockchyp.AuthorizationResponse{
				Success:          true,
				Approved:         true,
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
			assert: blockchyp.AuthorizationResponse{
				Success:          true,
				Approved:         true,
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

	cli := newCLI(t)

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			setup(t, test.instructions, true)

			cli.run(test.args, test.assert)
		})
	}
}
