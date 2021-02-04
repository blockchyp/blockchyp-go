package regression

import (
	"github.com/blockchyp/blockchyp-go"
)

var ebtTests = testCases{
	{
		name:  "EBT/Charge",
		group: testGroupMSR,
		operations: []operation{
			{
				msg: "Swipe an EBT test card when prompted. Enter PIN '1234'.",
				args: []string{
					"-type", "charge", "-terminal", terminalName,
					"-test", "-amount", partialAuthAuthorizedAmount,
					"-ebt",
				},
				expect: blockchyp.AuthorizationResponse{
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
		},
	},
	{
		name:  "EBT/MSRBalance",
		group: testGroupMSR,
		operations: []operation{
			{
				msg: "Swipe an EBT test card when prompted. Enter PIN '1234'.",
				args: []string{
					"-type", "balance", "-terminal", terminalName,
					"-test",
					"-ebt",
				},
				expect: blockchyp.BalanceResponse{
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
		},
	},
	{
		name:  "EBT/ManualBalance",
		group: testGroupInteractive,
		operations: []operation{
			{
				msg: "Key in the number '4111 1111 1111 1111' when prompted. Enter PIN '1234'.",
				args: []string{
					"-type", "balance", "-terminal", terminalName,
					"-test",
					"-ebt", "-manual",
				},
				expect: blockchyp.BalanceResponse{
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
		},
	},
}
