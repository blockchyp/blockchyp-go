package regression

import (
	"time"

	"github.com/blockchyp/blockchyp-go"
)

var reversalTests = testCases{
	{
		name:  "Reversal/Charge",
		group: testGroupNoCVM,
		operations: []operation{
			{
				msg: "Insert an EMV test card when prompted.",
				args: []string{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", amount(0),
					"-txRef", newTxRef,
				},
				expect: blockchyp.AuthorizationResponse{
					Success:         true,
					Approved:        true,
					Test:            true,
					TransactionType: "charge",
				},
			},
			{
				args: []string{
					"-type", "reverse", "-test",
					"-txRef", txRefN(0),
				},
				expect: blockchyp.AuthorizationResponse{
					Success:         true,
					Approved:        true,
					Test:            true,
					TransactionType: "reverse",
				},
			},
			{
				args: []string{
					"-type", "reverse", "-test",
					"-txRef", txRefN(0),
				},
				expect: blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					TransactionType:     "reverse",
					ResponseDescription: "Already Reversed",
				},
			},
		},
	},
	{
		name:  "Reversal/Preauth",
		group: testGroupNoCVM,
		operations: []operation{
			{
				msg: "Insert an EMV test card when prompted.",
				args: []string{
					"-type", "preauth", "-terminal", terminalName, "-test",
					"-amount", amount(0),
					"-txRef", newTxRef,
				},
				expect: blockchyp.AuthorizationResponse{
					Success:         true,
					Approved:        true,
					Test:            true,
					TransactionType: "preauth",
				},
			},
			{
				args: []string{
					"-type", "reverse", "-test",
					"-txRef", txRefN(0),
				},
				expect: blockchyp.AuthorizationResponse{
					Success:         true,
					Approved:        true,
					Test:            true,
					TransactionType: "reverse",
				},
			},
		},
	},
	{
		name:  "Reversal/Capture",
		group: testGroupNoCVM,
		operations: []operation{
			{
				msg: "Insert an EMV test card when prompted.",
				args: []string{
					"-type", "preauth", "-terminal", terminalName, "-test",
					"-amount", amount(0),
					"-txRef", newTxRef,
				},
				expect: blockchyp.AuthorizationResponse{
					Success:         true,
					Approved:        true,
					Test:            true,
					TransactionType: "preauth",
				},
			},
			{
				args: []string{
					"-type", "capture", "-test",
					"-tx", txIDN(0),
				},
				expect: blockchyp.CaptureResponse{
					Success:         true,
					Approved:        true,
					Test:            true,
					TransactionType: "capture",
				},
			},
			{
				args: []string{
					"-type", "reverse", "-test",
					"-txRef", txRefN(0),
				},
				expect: blockchyp.AuthorizationResponse{
					Success:         true,
					Approved:        true,
					Test:            true,
					TransactionType: "reverse",
				},
			},
		},
	},
	{
		name:  "Reversal/TimeLimit",
		group: testGroupNoCVM,
		operations: []operation{
			{
				msg: "Insert an EMV test card when prompted.",
				args: []string{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", amount(0),
					"-txRef", newTxRef,
				},
				expect: blockchyp.AuthorizationResponse{
					Success:         true,
					Approved:        true,
					Test:            true,
					TransactionType: "charge",
				},
			},
			{
				wait: 125 * time.Second,
				args: []string{
					"-type", "reverse", "-test",
					"-txRef", txRefN(0),
				},
				expect: blockchyp.AuthorizationResponse{
					Success:             false,
					Approved:            false,
					Test:                true,
					TransactionType:     "reverse",
					ResponseDescription: "Reverse Time Limit Exceeded. Use Void Instead.",
				},
			},
		},
	},
}
