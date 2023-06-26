package regression

import (
	"github.com/blockchyp/blockchyp-go/v2"
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
					"-queue", "-desc", "Dummy #1", "-txRef", "placeholder1",
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
					"-queue", "-desc", "Dummy #2", "-txRef", "placeholder2",
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
					"-queue", "-desc", "Dummy #3", "-txRef", "placeholder3",
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
					"-queue", "-desc", "Dummy #4", "-txRef", "placeholder4",
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
					"-queue", "-desc", "Dummy #5", "-txRef", "placeholder5",
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
			{
				args: []string{
					"-type", "delete-queue", "-terminal", terminalName, "-test",
					"-txRef", "*",
				},
				expect: blockchyp.DeleteQueuedTransactionResponse{
					Success: true,
				},
			},
			{
				args: []string{
					"-type", "clear", "-terminal", terminalName, "-test",
					"-txRef", "*",
				},
				expect: blockchyp.Acknowledgement{
					Success: true,
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
					"-queue", "-desc", "Dummy #1", "-txRef", "placeholder1",
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
					"-queue", "-desc", "Dummy #2", "-txRef", "placeholder2",
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
					"-queue", "-desc", "Dummy #3", "-txRef", "placeholder3",
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
					"-queue", "-desc", "Dummy #4", "-txRef", "placeholder4",
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
					"-queue", "-desc", "Dummy #5", "-txRef", "placeholder5",
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
			{
				args: []string{
					"-type", "delete-queue", "-terminal", terminalName, "-test",
					"-txRef", "*",
				},
				expect: blockchyp.DeleteQueuedTransactionResponse{
					Success: true,
				},
			},
			{
				args: []string{
					"-type", "clear", "-terminal", terminalName, "-test",
					"-txRef", "*",
				},
				expect: blockchyp.Acknowledgement{
					Success: true,
				},
			},
		},
	},
	{
		name:  "Async/QueueManagement",
		group: testGroupInteractive,
		operations: []operation{
			{
				args: []string{
					"-type", "delete-queue", "-terminal", terminalName, "-test",
					"-txRef", "*",
				},
				expect: blockchyp.DeleteQueuedTransactionResponse{
					Success: true,
				},
			},
			{
				args: []string{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", amount(0),
					"-queue", "-desc", "Test Ticket #1", "-txRef", newTxRef,
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
					"-amount", amount(0),
					"-queue", "-desc", "Test Ticket #2", "-txRef", newTxRef,
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
					"-txRef", txRefN(1),
				},
				expect: blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					ResponseDescription: "PENDING",
				},
			},
			{
				args: []string{
					"-type", "list-queue", "-terminal", terminalName, "-test",
				},
				expect: blockchyp.ListQueuedTransactionsResponse{
					Success: true,
					TransactionRefs: []string{
						txRefN(1),
						txRefN(2),
					},
				},
			},
			{
				args: []string{
					"-type", "delete-queue", "-terminal", terminalName, "-test",
					"-txRef", txRefN(1),
				},
				expect: blockchyp.DeleteQueuedTransactionResponse{
					Success: true,
				},
			},
			{
				args: []string{
					"-type", "list-queue", "-terminal", terminalName, "-test",
				},
				expect: blockchyp.ListQueuedTransactionsResponse{
					Success: true,
					TransactionRefs: []string{
						txRefN(2),
					},
				},
			},
			{
				args: []string{
					"-type", "tx-status",
					"-txRef", txRefN(1),
				},
				expect: blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					ResponseDescription: "CANCELED",
				},
			},
			{
				args: []string{
					"-type", "delete-queue", "-terminal", terminalName, "-test",
					"-txRef", txRefN(1),
				},
				expect: blockchyp.DeleteQueuedTransactionResponse{
					Success: false,
				},
			},
			{
				args: []string{
					"-type", "delete-queue", "-terminal", terminalName, "-test",
					"-txRef", "*",
				},
				expect: blockchyp.DeleteQueuedTransactionResponse{
					Success: true,
				},
			},
			{
				args: []string{
					"-type", "tx-status",
					"-txRef", txRefN(2),
				},
				expect: blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					ResponseDescription: "CANCELED",
				},
			},
			{
				args: []string{
					"-type", "list-queue", "-terminal", terminalName, "-test",
				},
				expect: blockchyp.ListQueuedTransactionsResponse{
					Success:         true,
					TransactionRefs: []string{},
				},
			},
		},
	},
}
