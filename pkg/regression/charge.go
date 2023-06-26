package regression

import (
	"github.com/blockchyp/blockchyp-go/v2"
)

var chargeTests = testCases{
	{
		name:  "Charge/ContactEMVNoCVMApproved",
		group: testGroupNoCVM,
		operations: []operation{
			{
				msg: "Insert a No-CVM EMV test card when prompted.",
				args: []string{
					"-type", "charge", "-terminal", terminalName,
					"-test", "-amount", amount(0),
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
		name:  "Charge/ContactlessEMVApproved",
		group: testGroupInteractive,
		operations: []operation{
			{
				msg: "Tap a contactless EMV test card when prompted.",
				args: []string{
					"-type", "charge", "-terminal", terminalName,
					"-test", "-amount", amountRange(0, 100, 1000),
				},
				expect: blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "charge",
					RequestedAmount:  amount(0),
					AuthorizedAmount: amount(0),
					EntryMethod:      "CONTACTLESS EMV",
					PaymentType:      notEmpty,
					MaskedPAN:        notEmpty,
					CardHolder:       notEmpty,
					ReceiptSuggestions: blockchyp.ReceiptSuggestions{
						AID:              notEmpty,
						ARQC:             notEmpty,
						IAD:              notEmpty,
						TVR:              notEmpty,
						MerchantName:     notEmpty,
						ApplicationLabel: notEmpty,
						RequestSignature: false,
						MaskedPAN:        notEmpty,
						AuthorizedAmount: amount(0),
						TransactionType:  "charge",
						EntryMethod:      "CONTACTLESS EMV",
					},
				},
			},
		},
	},
	{
		name:  "Charge/MSRVisa",
		group: testGroupMSR,
		operations: []operation{
			{
				msg: "Swipe a Visa MSR test card when prompted.",
				args: []string{
					"-type", "charge", "-terminal", terminalName,
					"-test", "-amount", amount(0),
				},
				expect: blockchyp.AuthorizationResponse{
					Success:     true,
					Approved:    true,
					Test:        true,
					EntryMethod: "SWIPE",
					PaymentType: "VISA",
					MaskedPAN:   notEmpty,
					CardHolder:  notEmpty,
				},
			},
		},
	},
	{
		name:  "Charge/MSRMasterCard",
		group: testGroupMSR,
		operations: []operation{
			{
				msg: "Swipe a MasterCard MSR test card when prompted.",
				args: []string{
					"-type", "charge", "-terminal", terminalName,
					"-test", "-amount", amount(0),
				},
				expect: blockchyp.AuthorizationResponse{
					Success:     true,
					Approved:    true,
					Test:        true,
					EntryMethod: "SWIPE",
					PaymentType: "MC",
					MaskedPAN:   notEmpty,
					CardHolder:  notEmpty,
				},
			},
		},
	},
	{
		name:  "Charge/MSRDiscover",
		group: testGroupMSR,
		operations: []operation{
			{
				msg: "Swipe a Discover MSR test card when prompted.",
				args: []string{
					"-type", "charge", "-terminal", terminalName,
					"-test", "-amount", amount(0),
				},
				expect: blockchyp.AuthorizationResponse{
					Success:     true,
					Approved:    true,
					Test:        true,
					EntryMethod: "SWIPE",
					PaymentType: "DISC",
					MaskedPAN:   notEmpty,
					CardHolder:  notEmpty,
				},
			},
		},
	},
	{
		name:  "Charge/MSRAmex",
		group: testGroupMSR,
		operations: []operation{
			{
				msg: "Swipe an Amex MSR test card when prompted.",
				args: []string{
					"-type", "charge", "-terminal", terminalName,
					"-test", "-amount", amount(0),
				},
				expect: blockchyp.AuthorizationResponse{
					Success:     true,
					Approved:    true,
					Test:        true,
					EntryMethod: "SWIPE",
					PaymentType: "AMEX",
					MaskedPAN:   notEmpty,
					CardHolder:  notEmpty,
				},
			},
		},
	},
	{
		name:  "Charge/MSRDiners",
		group: testGroupMSR,
		operations: []operation{
			{
				msg: "Swipe a Diner's Club MSR test card when prompted.",
				args: []string{
					"-type", "charge", "-terminal", terminalName,
					"-test", "-amount", amount(0),
				},
				expect: blockchyp.AuthorizationResponse{
					Success:     true,
					Approved:    true,
					Test:        true,
					EntryMethod: "SWIPE",
					PaymentType: "DINERS",
					MaskedPAN:   notEmpty,
					CardHolder:  notEmpty,
				},
			},
		},
	},
	{
		name:  "Charge/MSRJCB",
		group: testGroupMSR,
		operations: []operation{
			{
				msg: "Swipe a JCB MSR test card when prompted.",
				args: []string{
					"-type", "charge", "-terminal", terminalName,
					"-test", "-amount", amount(0),
				},
				expect: blockchyp.AuthorizationResponse{
					Success:     true,
					Approved:    true,
					Test:        true,
					EntryMethod: "SWIPE",
					PaymentType: "JCB",
					MaskedPAN:   notEmpty,
					CardHolder:  notEmpty,
				},
			},
		},
	},
	{
		name:  "Charge/MSRUnionPay",
		group: testGroupMSR,
		operations: []operation{
			{
				msg: "Swipe a UnionPay MSR test card when prompted.",
				args: []string{
					"-type", "charge", "-terminal", terminalName,
					"-test", "-amount", amount(0),
				},
				expect: blockchyp.AuthorizationResponse{
					Success:     true,
					Approved:    true,
					Test:        true,
					EntryMethod: "SWIPE",
					PaymentType: "CUP",
					MaskedPAN:   notEmpty,
					CardHolder:  notEmpty,
				},
			},
		},
	},
	{
		name:  "Charge/SignatureInResponse",
		group: testGroupSignature,
		sim:   true,
		operations: []operation{
			{
				msg: "Insert a signature CVM test card when prompted.",
				args: []string{
					"-type", "charge", "-terminal", terminalName,
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
		name:  "Charge/SignatureInFile",
		group: testGroupSignature,
		sim:   true,
		operations: []operation{
			{
				msg: "Insert a signature CVM test card when prompted.",
				args: []string{
					"-type", "charge", "-terminal", terminalName,
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
		name:  "Charge/SignatureRefused",
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
		name:  "Charge/SignatureTimeout",
		group: testGroupSignature,
		sim:   true,
		operations: []operation{
			{
				msg: `Insert a signature CVM test card when prompted.

Let the transaction time out when prompted for a signature. It should take 20 seconds.`,
				args: []string{
					"-type", "charge", "-terminal", terminalName,
					"-timeout", "20",
					"-test", "-amount", amount(0),
				},
				expect: blockchyp.Acknowledgement{
					Success: false,
				},
			},
		},
	},
	{
		name:  "Charge/SignatureDisabled",
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
		name:  "Charge/UserCanceled",
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
		name:  "Charge/ManualApproval",
		group: testGroupInteractive,
		operations: []operation{
			{
				msg: "Enter PAN '4111 1111 1111 1111' and CVV2 '123' when prompted.",
				args: []string{
					"-type", "charge", "-terminal", terminalName,
					"-test", "-amount", amount(0),
					"-manual",
				},
				expect: blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "charge",
					RequestedAmount:  amount(0),
					AuthorizedAmount: amount(0),
					EntryMethod:      "KEYED",
					MaskedPAN:        "************1111",
				},
			},
		},
	},
	{
		name:  "Charge/ManualDecline",
		group: testGroupInteractive,
		sim:   true,
		operations: []operation{
			{
				msg: "Enter PAN '4111 1111 1111 1129' and CVV2 '123' when prompted.",
				args: []string{
					"-type", "charge", "-terminal", terminalName,
					"-test", "-amount", amount(0),
					"-manual",
				},
				expect: blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         false,
					Test:             true,
					TransactionType:  "charge",
					RequestedAmount:  amount(0),
					AuthorizedAmount: "0.00",
					EntryMethod:      "KEYED",
					MaskedPAN:        "************1129",
				},
			},
		},
	},
	{
		name:  "Charge/EMVDecline",
		group: testGroupNoCVM,
		sim:   true,
		operations: []operation{
			{
				msg: "Insert an EMV test card when prompted.",
				args: []string{
					"-type", "charge", "-terminal", terminalName,
					"-test", "-amount", declineTriggerAmount,
				},
				expect: blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         false,
					Test:             true,
					TransactionType:  "charge",
					RequestedAmount:  declineTriggerAmount,
					AuthorizedAmount: "0.00",
				},
			},
		},
	},
	{
		name:  "Charge/EMVTimeout",
		group: testGroupNoCVM,
		sim:   true,
		operations: []operation{
			{
				msg: "Insert an EMV test card when prompted.",
				args: []string{
					"-type", "charge", "-terminal", terminalName,
					"-test", "-amount", timeOutTriggerAmount,
				},
				expect: blockchyp.AuthorizationResponse{
					Success:             false,
					Approved:            false,
					Test:                true,
					ResponseDescription: "Reversed: Network problem",
					TransactionType:     "charge",
					RequestedAmount:     timeOutTriggerAmount,
					AuthorizedAmount:    "0.00",
				},
			},
		},
	},
	{
		name:  "Charge/EMVPartialAuth",
		group: testGroupNoCVM,
		sim:   true,
		operations: []operation{
			{
				msg: "Insert an EMV test card when prompted.",
				args: []string{
					"-type", "charge", "-terminal", terminalName,
					"-test", "-amount", partialAuthTriggerAmount,
				},
				expect: blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					PartialAuth:      true,
					TransactionType:  "charge",
					RequestedAmount:  partialAuthTriggerAmount,
					AuthorizedAmount: partialAuthAuthorizedAmount,
				},
			},
		},
	},
	{
		name:  "Charge/EMVError",
		group: testGroupNoCVM,
		sim:   true,
		operations: []operation{
			{
				msg: "Insert an EMV test card when prompted.",
				args: []string{
					"-type", "charge", "-terminal", terminalName,
					"-test", "-amount", errorTriggerAmount,
				},
				expect: blockchyp.AuthorizationResponse{
					Success:             false,
					Approved:            false,
					Test:                true,
					TransactionType:     "charge",
					ResponseDescription: notEmpty,
					RequestedAmount:     errorTriggerAmount,
					AuthorizedAmount:    "0.00",
				},
			},
		},
	},
	{
		name:  "Charge/EMVNoResponse",
		group: testGroupNoCVM,
		sim:   true,
		operations: []operation{
			{
				msg: "Insert an EMV test card when prompted.",
				args: []string{
					"-type", "charge", "-terminal", terminalName,
					"-test", "-amount", noResponseTriggerAmount,
				},
				expect: blockchyp.AuthorizationResponse{
					Success:             false,
					Approved:            false,
					Test:                true,
					TransactionType:     "charge",
					ResponseDescription: "Reversed: Network problem",
					RequestedAmount:     noResponseTriggerAmount,
					AuthorizedAmount:    "0.00",
				},
			},
		},
	},
	{
		name:  "Charge/SFUnderLimit",
		group: testGroupInteractive,
		sim:   true,
		local: true,
		operations: []operation{
			{
				msg: "Tap a contactless EMV test card when prompted.",
				args: []string{
					"-type", "charge", "-terminal", terminalName,
					"-test", "-amount", "7.77",
				},
				expect: blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "charge",
					RequestedAmount:  "7.77",
					AuthorizedAmount: "7.77",
					StoreAndForward:  true,
					MaskedPAN:        notEmpty,
				},
			},
		},
	},
	{
		name:  "Charge/SFOverLimit",
		group: testGroupInteractive,
		sim:   true,
		local: true,
		operations: []operation{
			{
				msg: "Tap a contactless EMV test card when prompted.",
				args: []string{
					"-type", "charge", "-terminal", terminalName,
					"-test", "-amount", "77.77",
				},
				expect: blockchyp.AuthorizationResponse{
					Success:          false,
					Approved:         false,
					Test:             true,
					TransactionType:  "charge",
					RequestedAmount:  "77.77",
					AuthorizedAmount: "0.00",
				},
			},
		},
	},
	{
		name:  "Charge/EMVCashBack",
		group: testGroupInteractive,
		operations: []operation{
			{
				msg: `Insert an EMV debit card when prompted.

Select $1 when prompted for cash back.`,
				args: []string{
					"-type", "charge", "-terminal", terminalName,
					"-test", "-amount", amount(0),
					"-cashback",
				},
				expect: blockchyp.AuthorizationResponse{
					Success:                  true,
					Approved:                 true,
					Test:                     true,
					TransactionType:          "charge",
					RequestedAmount:          amount(0),
					RequestedCashBackAmount:  "1.00",
					AuthorizedAmount:         add(amount(0), 100),
					AuthorizedCashBackAmount: "1.00",
					EntryMethod:              "CHIP",
					PaymentType:              notEmpty,
					MaskedPAN:                notEmpty,
					CardHolder:               notEmpty,
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
						AuthorizedAmount: add(amount(0), 100),
						CashBackAmount:   "1.00",
						PINVerified:      true,
						TransactionType:  "charge",
						EntryMethod:      "CHIP",
					},
				},
			},
		},
	},
	{
		name:  "Charge/MSRCashBackSupported",
		group: testGroupMSR,
		operations: []operation{
			{
				msg: `Swipe the Visa MSR test card when prompted. Enter PIN '1234'.

Select $10 when prompted for cash back.`,
				args: []string{
					"-type", "charge", "-terminal", terminalName,
					"-test", "-amount", amount(0),
					"-debit", "-cashback",
				},
				expect: blockchyp.AuthorizationResponse{
					Success:                  true,
					Approved:                 true,
					Test:                     true,
					TransactionType:          "charge",
					RequestedAmount:          amount(0),
					RequestedCashBackAmount:  "10.00",
					AuthorizedAmount:         add(amount(0), 1000),
					AuthorizedCashBackAmount: "10.00",
					EntryMethod:              "SWIPE",
					PaymentType:              notEmpty,
					MaskedPAN:                notEmpty,
					CardHolder:               notEmpty,
					ReceiptSuggestions: blockchyp.ReceiptSuggestions{
						MerchantName:     notEmpty,
						RequestSignature: false,
						MaskedPAN:        notEmpty,
						AuthorizedAmount: add(amount(0), 1000),
						CashBackAmount:   "10.00",
						PINVerified:      true,
						TransactionType:  "charge",
						EntryMethod:      "SWIPE",
					},
				},
			},
		},
	},
	{
		name:  "Charge/MSRCashBackNotSupported",
		group: testGroupMSR,
		operations: []operation{
			{
				msg: `Swipe the AMEX MSR test card when prompted. Enter PIN '1234'.

It should not give you the option to select cash back.`,
				args: []string{
					"-type", "charge", "-terminal", terminalName,
					"-test", "-amount", amount(0),
					"-debit", "-cashback",
				},
				expect: blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "charge",
					RequestedAmount:  amount(0),
					AuthorizedAmount: amount(0),
					EntryMethod:      "SWIPE",
					PaymentType:      notEmpty,
					MaskedPAN:        notEmpty,
					CardHolder:       notEmpty,
					ReceiptSuggestions: blockchyp.ReceiptSuggestions{
						MerchantName:     notEmpty,
						RequestSignature: false,
						MaskedPAN:        notEmpty,
						AuthorizedAmount: amount(0),
						PINVerified:      true,
						TransactionType:  "charge",
						EntryMethod:      "SWIPE",
					},
				},
			},
			{
				validation: &validation{
					prompt: "Did the PIN entry screen show cash back options?",
					expect: false,
				},
			},
		},
	},
}
