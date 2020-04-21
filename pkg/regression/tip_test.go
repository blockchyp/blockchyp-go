// +build regression

package regression

import (
	"testing"

	"github.com/blockchyp/blockchyp-go"
)

func TestTip(t *testing.T) {
	tests := map[string]struct {
		instructions string
		args         []string
		assert       blockchyp.AuthorizationResponse
	}{
		"Percentage": {
			instructions: `Insert an EMV test card when prompted.

Select 15% when prompted for a tip.`,
			args: []string{
				"-type", "charge", "-terminal", terminalName,
				"-test", "-amount", amount(0),
				"-promptForTip",
			},
			assert: blockchyp.AuthorizationResponse{
				Success:          true,
				Approved:         true,
				Test:             true,
				TransactionType:  "charge",
				RequestedAmount:  mult(amount(0), 1.15),
				AuthorizedAmount: mult(amount(0), 1.15),
				TipAmount:        mult(amount(0), 0.15),
			},
		},
		"Custom": {
			instructions: `Insert an EMV test card when prompted.

Select 'Custom Amount' and enter '1.00' when prompted for a tip.`,
			args: []string{
				"-type", "charge", "-terminal", terminalName,
				"-test", "-amount", amount(0),
				"-promptForTip",
			},
			assert: blockchyp.AuthorizationResponse{
				Success:          true,
				Approved:         true,
				Test:             true,
				TransactionType:  "charge",
				RequestedAmount:  add(amount(0), 100),
				AuthorizedAmount: add(amount(0), 100),
				TipAmount:        "1.00",
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
