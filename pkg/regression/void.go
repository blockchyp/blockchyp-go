package regression

import (
	"github.com/blockchyp/blockchyp-go/v2"
)

var voidTests = testCases{
	{
		name:  "Void/Charge",
		group: testGroupNoCVM,
		operations: []operation{
			{
				msg: "Insert an EMV test card when prompted.",
				args: []string{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", amount(0),
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
					"-type", "void", "-test",
					"-tx", txIDN(0),
				},
				expect: blockchyp.VoidResponse{
					Success:         true,
					Approved:        true,
					Test:            true,
					TransactionType: "void",
				},
			},
		},
	},
	{
		name:  "Void/Preauth",
		group: testGroupNoCVM,
		operations: []operation{
			{
				msg: "Insert an EMV test card when prompted.",
				args: []string{
					"-type", "preauth", "-terminal", terminalName, "-test",
					"-amount", amount(0),
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
					"-type", "void", "-test",
					"-tx", txIDN(0),
				},
				expect: blockchyp.VoidResponse{
					Success:         true,
					Approved:        true,
					Test:            true,
					TransactionType: "void",
				},
			},
		},
	},
	{
		name:  "Void/Capture",
		group: testGroupNoCVM,
		operations: []operation{
			{
				msg: "Insert an EMV test card when prompted.",
				args: []string{
					"-type", "preauth", "-terminal", terminalName, "-test",
					"-amount", "63.00",
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
					"-type", "void", "-test",
					"-tx", txIDN(0),
				},
				expect: blockchyp.VoidResponse{
					Success:         true,
					Approved:        true,
					Test:            true,
					TransactionType: "void",
				},
			},
		},
	},
	{
		name:  "Void/Unknown",
		group: testGroupNonInteractive,
		operations: []operation{
			{
				args: []string{
					"-type", "void", "-test",
					"-tx", "NOT A REAL TRANSACTION",
				},
				expect: blockchyp.VoidResponse{
					Success:             false,
					Approved:            false,
					Test:                true,
					TransactionType:     "void",
					ResponseDescription: "Invalid Transaction",
				},
			},
		},
	},
	{
		name:  "Void/Double",
		group: testGroupNoCVM,
		operations: []operation{
			{
				msg: "Insert an EMV test card when prompted.",
				args: []string{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", amount(0),
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
					"-type", "void", "-test",
					"-tx", txIDN(0),
				},
				expect: blockchyp.VoidResponse{
					Success:         true,
					Approved:        true,
					Test:            true,
					TransactionType: "void",
				},
			},
			{
				args: []string{
					"-type", "void", "-test",
					"-tx", txIDN(0),
				},
				expect: blockchyp.VoidResponse{
					Success:             false,
					Approved:            false,
					Test:                true,
					TransactionType:     "void",
					ResponseDescription: "Already Voided",
				},
			},
		},
	},
}
