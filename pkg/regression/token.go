package regression

import (
	"github.com/blockchyp/blockchyp-go/v2"
)

var tokenTests = testCases{
	{
		name:  "Token/DirectEMV",
		group: testGroupNoCVM,

		operations: []operation{
			{
				msg: "Insert an EMV test card when prompted.",
				args: []string{
					"-type", "enroll", "-terminal", terminalName, "-test",
				},
				expect: blockchyp.AuthorizationResponse{
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
			},
			{
				args: []string{
					"-type", "charge", "-test",
					"-amount", amount(0), "-token", tokenN(0),
				},
				expect: blockchyp.AuthorizationResponse{
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
			},
			{
				args: []string{
					"-type", "preauth", "-test",
					"-amount", amount(1), "-token", tokenN(0),
				},
				expect: blockchyp.AuthorizationResponse{
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
			},
			{
				args: []string{
					"-type", "refund", "-test",
					"-amount", amount(2), "-token", tokenN(0),
				},
				expect: blockchyp.AuthorizationResponse{
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
	},
	{
		name:  "Token/Charge",
		group: testGroupNoCVM,
		operations: []operation{
			{
				msg: "Insert an EMV test card when prompted.",
				args: []string{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", amount(0), "-enroll",
				},
				expect: blockchyp.AuthorizationResponse{
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
			},
			{
				args: []string{
					"-type", "charge", "-test",
					"-amount", amount(1), "-token", tokenN(0),
				},
				expect: blockchyp.AuthorizationResponse{
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
			},
			{
				args: []string{
					"-type", "preauth", "-test",
					"-amount", amount(2), "-token", tokenN(0),
				},
				expect: blockchyp.AuthorizationResponse{
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
			},
			{
				args: []string{
					"-type", "refund", "-test",
					"-amount", amount(3), "-token", tokenN(0),
				},
				expect: blockchyp.AuthorizationResponse{
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
	},
	{
		name:  "Token/Preauth",
		group: testGroupNoCVM,
		operations: []operation{
			{
				msg: "Insert an EMV test card when prompted.",
				args: []string{
					"-type", "preauth", "-terminal", terminalName, "-test",
					"-amount", amount(0), "-enroll",
				},
				expect: blockchyp.AuthorizationResponse{
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
			},
			{
				args: []string{
					"-type", "charge", "-test",
					"-amount", amount(1), "-token", tokenN(0),
				},
				expect: blockchyp.AuthorizationResponse{
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
			},
			{
				args: []string{
					"-type", "preauth", "-test",
					"-amount", amount(2), "-token", tokenN(0),
				},
				expect: blockchyp.AuthorizationResponse{
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
			},
			{
				args: []string{
					"-type", "refund", "-test",
					"-amount", amount(3), "-token", tokenN(0),
				},
				expect: blockchyp.AuthorizationResponse{
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
	},
	{
		name:  "Token/DirectMSR",
		group: testGroupMSR,
		operations: []operation{
			{
				msg: "Swipe an MSR test card when prompted.",
				args: []string{
					"-type", "enroll", "-terminal", terminalName, "-test",
				},
				expect: blockchyp.AuthorizationResponse{
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
			},
			{
				args: []string{
					"-type", "charge", "-test",
					"-amount", amount(0), "-token", tokenN(0),
				},
				expect: blockchyp.AuthorizationResponse{
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
			},
			{
				args: []string{
					"-type", "preauth", "-test",
					"-amount", amount(1), "-token", tokenN(0),
				},
				expect: blockchyp.AuthorizationResponse{
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
			},
			{
				args: []string{
					"-type", "refund", "-test",
					"-amount", amount(2), "-token", tokenN(0),
				},
				expect: blockchyp.AuthorizationResponse{
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
	},
	{
		name:  "Token/DirectManual",
		group: testGroupInteractive,
		operations: []operation{
			{
				msg: "Enter PAN '4111 1111 1111 1111' and CVV2 '123' when prompted",
				args: []string{
					"-type", "enroll", "-terminal", terminalName, "-test",
					"-manual",
				},
				expect: blockchyp.AuthorizationResponse{
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
			},
			{
				args: []string{
					"-type", "charge", "-test",
					"-amount", amount(0), "-token", tokenN(0),
				},
				expect: blockchyp.AuthorizationResponse{
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
			},
			{
				args: []string{
					"-type", "preauth", "-test",
					"-amount", amount(1), "-token", tokenN(0),
				},
				expect: blockchyp.AuthorizationResponse{
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
			},
			{
				args: []string{
					"-type", "refund", "-test",
					"-amount", amount(2), "-token", tokenN(0),
				},
				expect: blockchyp.AuthorizationResponse{
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
	},
}
