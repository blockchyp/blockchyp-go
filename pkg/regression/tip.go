package regression

import (
	"github.com/blockchyp/blockchyp-go/v2"
)

var tipTests = testCases{
	{
		name:  "Tip/Percentage",
		group: testGroupNonInteractive,
		operations: []operation{
			{
				msg: `Insert an EMV test card when prompted.

Select 15% when prompted for a tip.`,
				args: []string{
					"-type", "charge", "-terminal", terminalName,
					"-test", "-amount", amount(0),
					"-promptForTip",
				},
				expect: blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "charge",
					RequestedAmount:  mult(amount(0), 1.15),
					AuthorizedAmount: mult(amount(0), 1.15),
					TipAmount:        mult(amount(0), 0.15),
				},
			},
		},
	},
	{
		name:  "Tip/Custom",
		group: testGroupNonInteractive,
		operations: []operation{
			{
				msg: `Insert an EMV test card when prompted.

Select 'Custom Amount' and enter '1.00' when prompted for a tip.`,
				args: []string{
					"-type", "charge", "-terminal", terminalName,
					"-test", "-amount", amount(0),
					"-promptForTip",
				},
				expect: blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "charge",
					RequestedAmount:  add(amount(0), 100),
					AuthorizedAmount: add(amount(0), 100),
					TipAmount:        "1.00",
				},
			},
		},
	},
}
