package regression

import (
	"github.com/blockchyp/blockchyp-go"
)

var asyncTests = testCases{
	{
		name:  "Async/AsyncCharge",
		group: testGroupNoCVM,
		operations: []operation{
			{
				msg: "Insert an EMV test card when prompted",
				args: []string{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", amount(0),
					"-async", "-txRef", newTxRef,
				},
				expect: blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					ResponseDescription: "Accepted",
				},
			},
			{
				args: []string{
					"-type", "tx-status",
					"-txRef", txRefN(0),
				},
				expect: blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					ResponseDescription: "PENDING",
				},
			},
			{
				msg: "Complete the transaction before continuing",
				args: []string{
					"-type", "tx-status",
					"-txRef", txRefN(0),
				},
				expect: blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "charge",
					RequestedAmount:  amount(0),
					AuthorizedAmount: amount(0),
					EntryMethod:      "CHIP",
					PaymentType:      notEmpty,
					MaskedPAN:        notEmpty,
					CardHolder:       notEmpty,
					ReceiptSuggestions: blockchyp.ReceiptSuggestions{
						AID:              notEmpty,
						ARQC:             notEmpty,
						IAD:              notEmpty,
						TVR:              notEmpty,
						TSI:              notEmpty,
						MerchantName:     notEmpty,
						ApplicationLabel: notEmpty,
						RequestSignature: false,
						MaskedPAN:        notEmpty,
						AuthorizedAmount: amount(0),
						TransactionType:  "charge",
						EntryMethod:      "CHIP",
					},
				},
			},
		},
	},
	{
		name:  "Async/AsyncPreauth",
		group: testGroupNoCVM,
		operations: []operation{
			{
				msg: "Insert an EMV test card when prompted",
				args: []string{
					"-type", "preauth", "-terminal", terminalName, "-test",
					"-amount", amount(0),
					"-async", "-txRef", newTxRef,
				},
				expect: blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					ResponseDescription: "Accepted",
				},
			},
			{
				args: []string{
					"-type", "tx-status",
					"-txRef", txRefN(0),
				},
				expect: blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					ResponseDescription: "PENDING",
				},
			},
			{
				msg: "Complete the transaction before continuing",
				args: []string{
					"-type", "tx-status",
					"-txRef", txRefN(0),
				},
				expect: blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "preauth",
					RequestedAmount:  amount(0),
					AuthorizedAmount: amount(0),
					EntryMethod:      "CHIP",
					PaymentType:      notEmpty,
					MaskedPAN:        notEmpty,
					CardHolder:       notEmpty,
					ReceiptSuggestions: blockchyp.ReceiptSuggestions{
						AID:              notEmpty,
						ARQC:             notEmpty,
						IAD:              notEmpty,
						TVR:              notEmpty,
						TSI:              notEmpty,
						MerchantName:     notEmpty,
						ApplicationLabel: notEmpty,
						RequestSignature: false,
						MaskedPAN:        notEmpty,
						AuthorizedAmount: amount(0),
						TransactionType:  "preauth",
						EntryMethod:      "CHIP",
					},
				},
			},
		},
	},
	{
		name:  "Async/QueueCharge",
		group: testGroupInteractive,
		operations: []operation{
			{
				args: []string{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", amount(0),
					"-queue", "-desc", "Test Ticket", "-txRef", newTxRef,
				},
				expect: blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					ResponseDescription: "Queued",
				},
			},
			{
				args: []string{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", amount(1),
					"-queue", "-desc", "Dummy #1", "-txRef", "placeholder",
				},
				expect: blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					ResponseDescription: "Queued",
				},
			},
			{
				args: []string{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", amount(2),
					"-queue", "-desc", "Dummy #2", "-txRef", "placeholder",
				},
				expect: blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					ResponseDescription: "Queued",
				},
			},
			{
				args: []string{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", amount(3),
					"-queue", "-desc", "Dummy #3", "-txRef", "placeholder",
				},
				expect: blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					ResponseDescription: "Queued",
				},
			},
			{
				args: []string{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", amount(4),
					"-queue", "-desc", "Dummy #4", "-txRef", "placeholder",
				},
				expect: blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					ResponseDescription: "Queued",
				},
			},
			{
				args: []string{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", amount(5),
					"-queue", "-desc", "Dummy #5", "-txRef", "placeholder",
				},
				expect: blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					ResponseDescription: "Queued",
				},
			},
			{
				args: []string{
					"-type", "tx-status",
					"-txRef", txRefN(0),
				},
				expect: blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					ResponseDescription: "PENDING",
				},
			},
			{
				msg: "Select the ticket labeled 'Test Ticket` and insert an EMV test card when prompted.",
				args: []string{
					"-type", "tx-status",
					"-txRef", txRefN(0),
				},
				expect: blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "charge",
					RequestedAmount:  amount(0),
					AuthorizedAmount: amount(0),
					EntryMethod:      "CHIP",
					PaymentType:      notEmpty,
					MaskedPAN:        notEmpty,
					CardHolder:       notEmpty,
					ReceiptSuggestions: blockchyp.ReceiptSuggestions{
						AID:              notEmpty,
						ARQC:             notEmpty,
						IAD:              notEmpty,
						TVR:              notEmpty,
						TSI:              notEmpty,
						MerchantName:     notEmpty,
						ApplicationLabel: notEmpty,
						RequestSignature: false,
						MaskedPAN:        notEmpty,
						AuthorizedAmount: amount(0),
						TransactionType:  "charge",
						EntryMethod:      "CHIP",
					},
				},
			},
		},
	},
	{
		name:  "Async/QueuePreauth",
		group: testGroupInteractive,
		operations: []operation{
			{
				args: []string{
					"-type", "preauth", "-terminal", terminalName, "-test",
					"-amount", amount(0),
					"-queue", "-desc", "Test Ticket", "-txRef", newTxRef,
				},
				expect: blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					ResponseDescription: "Queued",
				},
			},
			{
				args: []string{
					"-type", "preauth", "-terminal", terminalName, "-test",
					"-amount", amount(1),
					"-queue", "-desc", "Dummy #1", "-txRef", "placeholder",
				},
				expect: blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					ResponseDescription: "Queued",
				},
			},
			{
				args: []string{
					"-type", "preauth", "-terminal", terminalName, "-test",
					"-amount", amount(2),
					"-queue", "-desc", "Dummy #2", "-txRef", "placeholder",
				},
				expect: blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					ResponseDescription: "Queued",
				},
			},
			{
				args: []string{
					"-type", "preauth", "-terminal", terminalName, "-test",
					"-amount", amount(3),
					"-queue", "-desc", "Dummy #3", "-txRef", "placeholder",
				},
				expect: blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					ResponseDescription: "Queued",
				},
			},
			{
				args: []string{
					"-type", "preauth", "-terminal", terminalName, "-test",
					"-amount", amount(4),
					"-queue", "-desc", "Dummy #4", "-txRef", "placeholder",
				},
				expect: blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					ResponseDescription: "Queued",
				},
			},
			{
				args: []string{
					"-type", "preauth", "-terminal", terminalName, "-test",
					"-amount", amount(5),
					"-queue", "-desc", "Dummy #5", "-txRef", "placeholder",
				},
				expect: blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					ResponseDescription: "Queued",
				},
			},
			{
				args: []string{
					"-type", "tx-status",
					"-txRef", txRefN(0),
				},
				expect: blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					ResponseDescription: "PENDING",
				},
			},
			{
				msg: "Select the ticket labeled 'Test Ticket` and insert an EMV test card when prompted.",
				args: []string{
					"-type", "tx-status",
					"-txRef", txRefN(0),
				},
				expect: blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "preauth",
					RequestedAmount:  amount(0),
					AuthorizedAmount: amount(0),
					EntryMethod:      "CHIP",
					PaymentType:      notEmpty,
					MaskedPAN:        notEmpty,
					CardHolder:       notEmpty,
					ReceiptSuggestions: blockchyp.ReceiptSuggestions{
						AID:              notEmpty,
						ARQC:             notEmpty,
						IAD:              notEmpty,
						TVR:              notEmpty,
						TSI:              notEmpty,
						MerchantName:     notEmpty,
						ApplicationLabel: notEmpty,
						RequestSignature: false,
						MaskedPAN:        notEmpty,
						AuthorizedAmount: amount(0),
						TransactionType:  "preauth",
						EntryMethod:      "CHIP",
					},
				},
			},
		},
	},
}
