package regression

import (
	"github.com/blockchyp/blockchyp-go"
)

var preauthTests = testCases{
	{
		name:  "Preauth/EMVApproved",
		group: testGroupNoCVM,
		operations: []operation{
			{
				msg: "Insert an EMV test card when prompted.",
				args: []string{
					"-type", "preauth", "-terminal", terminalName,
					"-test", "-amount", amount(0),
				},
				expect: blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "preauth",
					RequestedAmount:  amount(0),
					AuthorizedAmount: amount(0),
				},
			},
			{
				args: []string{
					"-type", "capture", "-test",
					"-tx", txIDN(0),
				},
				expect: blockchyp.CaptureResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "capture",
					AuthorizedAmount: amount(0),
				},
			},
			{
				args: []string{
					"-type", "capture", "-test",
					"-tx", txIDN(0),
				},
				expect: blockchyp.CaptureResponse{
					Success:             false,
					Approved:            false,
					Test:                true,
					TransactionType:     "capture",
					ResponseDescription: "Already Captured",
				},
			},
		},
	},
	{
		name:  "Preauth/SignatureInResponse",
		group: testGroupSignature,
		sim:   true,
		operations: []operation{
			{
				msg: "Insert a signature CVM test card when prompted.",
				args: []string{
					"-type", "preauth", "-terminal", terminalName,
					"-test", "-amount", amount(0),
					"-sigFormat", blockchyp.SignatureFormatPNG,
					"-sigWidth", "50",
				},
				expect: blockchyp.AuthorizationResponse{
					Success:         true,
					Approved:        true,
					Test:            true,
					TransactionType: "preauth",
					SigFile:         notEmpty,
				},
			},
		},
	},
	{
		name:  "Preauth/SignatureInFile",
		group: testGroupSignature,
		sim:   true,
		operations: []operation{
			{
				msg: "Insert a signature CVM test card when prompted",
				args: []string{
					"-type", "preauth", "-terminal", terminalName,
					"-test", "-amount", amount(0),
					"-sigWidth", "400", "-sigFile", "/tmp/blockchyp-regression-test/sig.png",
				},
				expect: blockchyp.AuthorizationResponse{
					Success:         true,
					Approved:        true,
					Test:            true,
					TransactionType: "preauth",
				},
			},
			{
				validation: &validation{
					prompt: "Does the signature appear valid in the browser?",
					serve:  "/tmp/blockchyp-regression-test/sig.png",
					expect: true,
				},
			},
		},
	},
	{
		name:  "Preauth/SignatureRefused",
		group: testGroupSignature,
		sim:   true,
		operations: []operation{
			{
				msg: `Insert a signature CVM test card when prompted.

When prompted for a signature, hit 'Done' without signing.`,
				args: []string{
					"-type", "preauth", "-terminal", terminalName,
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
		name:  "Preauth/UserCanceled",
		group: testGroupInteractive,
		sim:   true,
		operations: []operation{
			{
				msg: "Hit the red 'X' button when prompted for a card.",
				args: []string{
					"-type", "preauth", "-terminal", terminalName,
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
		name:  "Preauth/SignatureTimeout",
		group: testGroupSignature,
		sim:   true,
		operations: []operation{
			{
				msg: `Insert a signature CVM test card when prompted.

Let the transaction time out when prompted for a signature. It should take 20 seconds.`,
				args: []string{
					"-type", "preauth", "-terminal", terminalName,
					"-test", "-timeout", "20", "-amount", amount(0),
				},
				expect: blockchyp.Acknowledgement{
					Success: false,
				},
			},
		},
	},
	{
		name:  "Preauth/ManualApproval",
		group: testGroupInteractive,
		operations: []operation{
			{
				msg: "Enter PAN '4111 1111 1111 1111' and CVV2 '123' when prompted",
				args: []string{
					"-type", "preauth", "-terminal", terminalName,
					"-test", "-amount", amount(0),
					"-manual",
				},
				expect: blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "preauth",
					RequestedAmount:  amount(0),
					AuthorizedAmount: amount(0),
					EntryMethod:      "KEYED",
					MaskedPAN:        "************1111",
				},
			},
			{
				args: []string{
					"-type", "capture", "-test",
					"-tx", txIDN(0),
				},
				expect: blockchyp.CaptureResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "capture",
					AuthorizedAmount: amount(0),
				},
			},
		},
	},
	{
		name:  "Preauth/EMVDecline",
		group: testGroupNoCVM,
		sim:   true,
		operations: []operation{
			{
				msg: "Insert any EMV test card.",
				args: []string{
					"-type", "preauth", "-terminal", terminalName,
					"-test", "-amount", declineTriggerAmount,
				},
				expect: blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         false,
					Test:             true,
					TransactionType:  "preauth",
					RequestedAmount:  declineTriggerAmount,
					AuthorizedAmount: "0.00",
				},
			},
			{
				args: []string{
					"-type", "capture", "-test",
					"-tx", txIDN(0),
				},
				expect: blockchyp.CaptureResponse{
					Success:         false,
					Approved:        false,
					Test:            true,
					TransactionType: "capture",
				},
			},
		},
	},
	{
		name:  "Preauth/EMVTimeout",
		group: testGroupNoCVM,
		sim:   true,
		operations: []operation{
			{
				msg: "Insert any EMV test card.",
				args: []string{
					"-type", "preauth", "-terminal", terminalName,
					"-test", "-amount", timeOutTriggerAmount,
				},
				expect: blockchyp.AuthorizationResponse{
					Success:             false,
					Approved:            false,
					Test:                true,
					ResponseDescription: "Reversed: Network problem",
					TransactionType:     "preauth",
					RequestedAmount:     timeOutTriggerAmount,
					AuthorizedAmount:    "0.00",
				},
			},
		},
	},
	{
		name:  "Preauth/EMVPartialAuth",
		group: testGroupNoCVM,
		sim:   true,
		operations: []operation{
			{
				msg: "Insert any EMV test card.",
				args: []string{
					"-type", "preauth", "-terminal", terminalName,
					"-test", "-amount", partialAuthTriggerAmount,
				},
				expect: blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					PartialAuth:      true,
					TransactionType:  "preauth",
					RequestedAmount:  partialAuthTriggerAmount,
					AuthorizedAmount: partialAuthAuthorizedAmount,
				},
			},
		},
	},
	{
		name:  "Preauth/EMVError",
		group: testGroupNoCVM,
		sim:   true,
		operations: []operation{
			{
				msg: "Insert any EMV test card.",
				args: []string{
					"-type", "preauth", "-terminal", terminalName,
					"-test", "-amount", errorTriggerAmount,
				},
				expect: blockchyp.AuthorizationResponse{
					Success:             false,
					Approved:            false,
					Test:                true,
					TransactionType:     "preauth",
					ResponseDescription: notEmpty,
					RequestedAmount:     errorTriggerAmount,
					AuthorizedAmount:    "0.00",
				},
			},
		},
	},
	{
		name:  "Preauth/EMVNoResponse",
		group: testGroupNoCVM,
		sim:   true,
		operations: []operation{
			{
				msg: "Insert any EMV test card.",
				args: []string{
					"-type", "preauth", "-terminal", terminalName,
					"-test", "-amount", noResponseTriggerAmount,
				},
				expect: blockchyp.AuthorizationResponse{
					Success:             false,
					Approved:            false,
					Test:                true,
					TransactionType:     "preauth",
					ResponseDescription: "Reversed: Network problem",
					RequestedAmount:     noResponseTriggerAmount,
					AuthorizedAmount:    "0.00",
				},
			},
		},
	},
	{
		name:  "Preauth/TipAdjust",
		group: testGroupNoCVM,
		operations: []operation{
			{
				msg: "Insert an EMV test card when prompted.",
				args: []string{
					"-type", "preauth", "-terminal", terminalName,
					"-test", "-amount", amount(0),
				},
				expect: blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "preauth",
					RequestedAmount:  amount(0),
					AuthorizedAmount: amount(0),
				},
			},
			{
				args: []string{
					"-type", "capture", "-test",
					"-tip", "1.00", "-amount", add(amount(0), 100),
					"-tx", txIDN(0),
				},
				expect: blockchyp.CaptureResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "capture",
					AuthorizedAmount: add(amount(0), 100),
					TipAmount:        "1.00",
				},
			},
		},
	},
	{
		name:  "Preauth/OrphanCapture",
		group: testGroupNonInteractive,
		operations: []operation{
			{
				args: []string{
					"-type", "capture", "-test",
					"-tx", newTxRef,
				},
				expect: blockchyp.CaptureResponse{
					Success:             false,
					Approved:            false,
					Test:                true,
					TransactionType:     "capture",
					ResponseDescription: "Invalid Transaction",
				},
			},
		},
	},
	{
		name:  "Preauth/VoidCapture",
		group: testGroupNoCVM,
		operations: []operation{
			{
				msg: "Insert an EMV test card when prompted.",
				args: []string{
					"-type", "preauth", "-terminal", terminalName,
					"-test", "-amount", amount(0),
				},
				expect: blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "preauth",
					RequestedAmount:  amount(0),
					AuthorizedAmount: amount(0),
				},
			},
			{
				args: []string{
					"-type", "void", "-test",
					"-tx", txIDN(0),
				},
				expect: blockchyp.VoidResponse{
					Success:  true,
					Approved: true,
					Test:     true,
				},
			},
			{
				args: []string{
					"-type", "capture", "-test",
					"-tx", txIDN(0),
				},
				expect: blockchyp.CaptureResponse{
					Success:             false,
					Approved:            false,
					Test:                true,
					TransactionType:     "capture",
					ResponseDescription: "Voided Transaction",
				},
			},
		},
	},
	{
		name:  "Preauth/ClosedBatchCapture",
		group: testGroupNoCVM,
		operations: []operation{
			{
				msg: "Insert an EMV test card when prompted.",
				args: []string{
					"-type", "preauth", "-terminal", terminalName,
					"-test", "-amount", amount(0),
				},
				expect: blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "preauth",
					RequestedAmount:  amount(0),
					AuthorizedAmount: amount(0),
				},
			},
			{
				args: []string{
					"-type", "close-batch", "-test",
				},
				expect: blockchyp.CloseBatchResponse{
					Success: true,
					Test:    true,
					Batches: []blockchyp.BatchSummary{
						{
							CapturedAmount: notEmpty,
							OpenPreauths:   notEmpty,
						},
					},
				},
			},
			{
				args: []string{
					"-type", "capture", "-test",
					"-tx", txIDN(0),
				},
				expect: blockchyp.CaptureResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "capture",
					AuthorizedAmount: amount(0),
				},
			},
		},
	},
}
