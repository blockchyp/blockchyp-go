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
			instructions: `Select 15% when prompted for a tip.

Insert a valid test card when prompted.`,
			args: []string{
				"-type", "charge", "-terminal", "Test Terminal",
				"-test", "-amount", "59.00",
				"-promptForTip",
			},
			assert: blockchyp.AuthorizationResponse{
				Success:          true,
				Approved:         true,
				Test:             true,
				TransactionType:  "charge",
				RequestedAmount:  "67.85",
				AuthorizedAmount: "67.85",
				TipAmount:        "8.85",
			},
		},
		"Custom": {
			instructions: `Select 'Custom Amount' and enter '1.00' when prompted for a tip.

Insert a valid test card when prompted.`,
			args: []string{
				"-type", "charge", "-terminal", "Test Terminal",
				"-test", "-amount", "59.00",
				"-promptForTip",
			},
			assert: blockchyp.AuthorizationResponse{
				Success:          true,
				Approved:         true,
				Test:             true,
				TransactionType:  "charge",
				RequestedAmount:  "60.00",
				AuthorizedAmount: "60.00",
				TipAmount:        "1.00",
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
