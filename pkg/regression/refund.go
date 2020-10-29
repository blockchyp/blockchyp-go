package regression

import (
	"github.com/blockchyp/blockchyp-go"
)

var refundTests = testCases{
	{
		name:  "Refund/Charge",
		group: testGroupNoCVM,
		operations: []operation{
			{
				msg: "Insert an EMV test card when prompted.",
				args: []string{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", amount(0),
				},
				expect: blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "charge",
					AuthorizedAmount: amount(0),
				},
			},
			{
				args: []string{
					"-type", "refund", "-test",
					"-tx", txIDN(0),
				},
				expect: blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "void", // Same-day refund processed as void
					AuthorizedAmount: amount(0),
				},
			},
		},
	},
	{
		name:  "Refund/Partial",
		group: testGroupNoCVM,
		operations: []operation{
			{
				msg: "Insert an EMV test card when prompted.",
				args: []string{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", amountRange(0, 501, 999),
				},
				expect: blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "charge",
					AuthorizedAmount: amount(0),
				},
			},
			{
				msg: "Insert an EMV test card when prompted.",
				args: []string{
					"-type", "refund", "-test",
					"-amount", "5.00",
					"-tx", txIDN(0),
				},
				expect: blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "refund",
					AuthorizedAmount: "5.00",
				},
			},
			{
				msg: "Insert an EMV test card when prompted.",
				args: []string{
					"-type", "refund", "-test",
					"-amount", "10.00",
					"-tx", txIDN(0),
				},
				expect: blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					TransactionType:     "refund",
					AuthorizedAmount:    "0.00",
					ResponseDescription: "Refund would exceed the original transaction amount",
				},
			},
		},
	},
	{
		name:  "Refund/Excess",
		group: testGroupNoCVM,
		operations: []operation{
			{
				msg: "Insert an EMV test card when prompted.",
				args: []string{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", amountRange(0, 500, 1000),
				},
				expect: blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "charge",
					AuthorizedAmount: amount(0),
				},
			},
			{
				args: []string{
					"-type", "refund", "-test",
					"-amount", amountRange(1, 1001, 2000),
					"-tx", txIDN(0),
				},
				expect: blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					TransactionType:     "refund",
					AuthorizedAmount:    "0.00",
					ResponseDescription: "Refund would exceed the original transaction amount",
				},
			},
		},
	},
	{
		name:  "Refund/SF",
		group: testGroupNoCVM,
		sim:   true,
		local: true,
		operations: []operation{
			{
				msg: "Insert an EMV test card when prompted.",
				args: []string{
					"-type", "refund", "-terminal", terminalName, "-test",
					"-amount", "7.77",
				},
				expect: blockchyp.AuthorizationResponse{
					Success:          false,
					Approved:         false,
					Test:             true,
					TransactionType:  "refund",
					RequestedAmount:  "7.77",
					AuthorizedAmount: "0.00",
				},
			},
		},
	},
	{
		name:  "Refund/BadCredentials",
		group: testGroupNonInteractive,
		operations: []operation{
			{
				args: []string{
					"-type", "refund", "-test",
					"-tx", "OFE3TTQFJ4I6TNTUNSLM7WZLHE",
					"-apiKey", "X6N2KIQEWYI6TCADNSLM7WZLHE",
					"-bearerToken", "RIKLAPSMSMG2YII27N2NPAMCS5",
					"-signingKey", "4b556bc4e73ffc86fc5f8bfbba1598e7a8cd91f44fd7072d070c92fae7f48cd9",
				},
				expect: blockchyp.AuthorizationResponse{
					Success:  false,
					Approved: false,
				},
			},
		},
	},
	{
		name:  "Refund/EMVFreeRange",
		group: testGroupNoCVM,
		operations: []operation{
			{
				msg: `Insert an EMV test card when prompted.

Leave the card in the terminal until the test completes.`,
				args: []string{
					"-type", "refund", "-terminal", terminalName, "-test",
					"-amount", amount(0),
				},
				expect: blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "refund",
					RequestedAmount:  amount(0),
					AuthorizedAmount: amount(0),
					EntryMethod:      "CHIP",
				},
			},
			{
				args: []string{
					"-type", "refund", "-terminal", terminalName, "-test",
					"-amount", amount(0),
				},
				expect: blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					TransactionType:     "refund",
					RequestedAmount:     amount(0),
					AuthorizedAmount:    "0.00",
					EntryMethod:         "CHIP",
					ResponseDescription: "Duplicate Transaction",
				},
			},
		},
	},
	{
		name:  "Refund/SignatureInResponse",
		group: testGroupSignature,
		sim:   true,
		operations: []operation{
			{
				msg: "Insert a signature CVM test card when prompted.",
				args: []string{
					"-type", "refund", "-terminal", terminalName,
					"-test", "-amount", amount(0),
					"-sigFormat", blockchyp.SignatureFormatJPG,
					"-sigWidth", "50",
				},
				expect: blockchyp.AuthorizationResponse{
					Success:  true,
					Approved: true,
					Test:     true,
					SigFile:  notEmpty,
				},
			},
		},
	},
	{
		name:  "Refund/SignatureInFile",
		group: testGroupSignature,
		sim:   true,
		operations: []operation{
			{
				msg: "Insert a signature CVM test card when prompted.",
				args: []string{
					"-type", "refund", "-terminal", terminalName,
					"-test", "-amount", amount(0),
					"-sigWidth", "400", "-sigFile", "/tmp/blockchyp-regression-test/sig.jpg",
				},
				expect: blockchyp.AuthorizationResponse{
					Success:  true,
					Approved: true,
					Test:     true,
				},
			},
			{
				validation: &validation{
					prompt: "Does the signature appear valid in the browser?",
					serve:  "/tmp/blockchyp-regression-test/sig.jpg",
					expect: true,
				},
			},
		},
	},
	{
		name:  "Refund/SignatureRefused",
		group: testGroupSignature,
		sim:   true,
		operations: []operation{
			{
				msg: `Insert a signature CVM test card when prompted.

When prompted for a signature, hit 'Done' without signing.`,
				args: []string{
					"-type", "charge", "-terminal", terminalName,
					"-test", "-amount", amount(0),
				},
				expect: blockchyp.AuthorizationResponse{
					Success:             false,
					Approved:            false,
					Test:                true,
					ResponseDescription: "Reversed: Customer did not sign",
				},
			},
		},
	},
	{
		name:  "Refund/SignatureTimeout",
		group: testGroupSignature,
		sim:   true,
		operations: []operation{
			{
				msg: `Insert a signature CVM test card when prompted.

Let the transaction time out when prompted for a signature. It should take 20 seconds.`,
				args: []string{
					"-type", "charge", "-terminal", terminalName,
					"-test", "-timeout", "20", "-amount", amount(0),
				},
				expect: blockchyp.Acknowledgement{
					Success: false,
				},
			},
		},
	},
	{
		name:  "Refund/SignatureDisabled",
		group: testGroupSignature,
		operations: []operation{
			{
				msg: "Insert a signature CVM test card when prompted.",
				args: []string{
					"-type", "charge", "-terminal", terminalName,
					"-test", "-amount", amount(0),
					"-disableSignature",
				},
				expect: blockchyp.AuthorizationResponse{
					Success:  true,
					Approved: true,
					Test:     true,
					ReceiptSuggestions: blockchyp.ReceiptSuggestions{
						RequestSignature: true,
					},
				},
			},
		},
	},
	{
		name:  "Refund/UserCanceled",
		group: testGroupInteractive,
		sim:   true,
		operations: []operation{
			{
				msg: "Hit the red 'X' button when prompted for a card.",
				args: []string{
					"-type", "charge", "-terminal", terminalName,
					"-test", "-amount", amount(0),
				},
				expect: blockchyp.AuthorizationResponse{
					Success:             false,
					Approved:            false,
					Test:                true,
					ResponseDescription: "User canceled",
				},
			},
		},
	},
	{
		name:  "Refund/MSRFreeRange",
		group: testGroupMSR,
		operations: []operation{
			{
				msg: "Swipe an MSR test card when prompted.",
				args: []string{
					"-type", "refund", "-terminal", terminalName, "-test",
					"-amount", amount(0),
				},
				expect: blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "refund",
					EntryMethod:      "SWIPE",
					AuthorizedAmount: amount(0),
				},
			},
		},
	},
	{
		name:  "Refund/ManualFreeRange",
		group: testGroupInteractive,
		operations: []operation{
			{
				msg: "Enter PAN '4111 1111 1111 1111' and CVV2 '123' when prompted.",
				args: []string{
					"-type", "refund", "-terminal", terminalName, "-test",
					"-amount", amount(0), "-manual",
				},
				expect: blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "refund",
					EntryMethod:      "KEYED",
					AuthorizedAmount: amount(0),
				},
			},
		},
	},
	{
		name:  "Refund/DeclineFreeRange",
		group: testGroupMSR,
		sim:   true,
		operations: []operation{
			{
				msg: "Swipe the 'Decline' MSR test card when prompted.",
				args: []string{
					"-type", "refund", "-terminal", terminalName, "-test",
					"-amount", amount(0),
				},
				expect: blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         false,
					Test:             true,
					TransactionType:  "refund",
					EntryMethod:      "SWIPE",
					RequestedAmount:  amount(0),
					AuthorizedAmount: "0.00",
				},
			},
		},
	},
}
