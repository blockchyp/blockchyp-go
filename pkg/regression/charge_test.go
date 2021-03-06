// +build regression

package regression

import (
	"testing"

	"github.com/blockchyp/blockchyp-go"
)

func TestCharge(t *testing.T) {
	tests := map[string]struct {
		instructions string
		args         []string
		assert       blockchyp.AuthorizationResponse
		validation   validation

		// localOnly causes tests to be skipped when running in cloud relay
		// mode.
		localOnly bool

		// simOnly causes tests to be skipped when running in acquirer mode.
		simOnly bool
	}{
		"ContactEMVNoCVMApproved": {
			instructions: "Insert a No-CVM EMV test card when prompted.",
			args: []string{
				"-type", "charge", "-terminal", terminalName,
				"-test", "-amount", amount(0),
			},
			assert: blockchyp.AuthorizationResponse{
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
		"ContactlessEMVApproved": {
			instructions: "Tap a contactless EMV test card when prompted.",
			args: []string{
				"-type", "charge", "-terminal", terminalName,
				"-test", "-amount", amountRange(0, 100, 1000),
			},
			assert: blockchyp.AuthorizationResponse{
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
		"MSRVisa": {
			instructions: "Swipe a Visa MSR test card when prompted.",
			args: []string{
				"-type", "charge", "-terminal", terminalName,
				"-test", "-amount", amount(0),
			},
			assert: blockchyp.AuthorizationResponse{
				Success:     true,
				Approved:    true,
				Test:        true,
				EntryMethod: "SWIPE",
				PaymentType: "VISA",
				MaskedPAN:   notEmpty,
				CardHolder:  notEmpty,
			},
		},
		"MSRMasterCard": {
			instructions: "Swipe a MasterCard MSR test card when prompted.",
			args: []string{
				"-type", "charge", "-terminal", terminalName,
				"-test", "-amount", amount(0),
			},
			assert: blockchyp.AuthorizationResponse{
				Success:     true,
				Approved:    true,
				Test:        true,
				EntryMethod: "SWIPE",
				PaymentType: "MC",
				MaskedPAN:   notEmpty,
				CardHolder:  notEmpty,
			},
		},
		"MSRDiscover": {
			instructions: "Swipe a Discover MSR test card when prompted.",
			args: []string{
				"-type", "charge", "-terminal", terminalName,
				"-test", "-amount", amount(0),
			},
			assert: blockchyp.AuthorizationResponse{
				Success:     true,
				Approved:    true,
				Test:        true,
				EntryMethod: "SWIPE",
				PaymentType: "DISC",
				MaskedPAN:   notEmpty,
				CardHolder:  notEmpty,
			},
		},
		"MSRAmex": {
			instructions: "Swipe an Amex MSR test card when prompted.",
			args: []string{
				"-type", "charge", "-terminal", terminalName,
				"-test", "-amount", amount(0),
			},
			assert: blockchyp.AuthorizationResponse{
				Success:     true,
				Approved:    true,
				Test:        true,
				EntryMethod: "SWIPE",
				PaymentType: "AMEX",
				MaskedPAN:   notEmpty,
				CardHolder:  notEmpty,
			},
		},
		"MSRDiners": {
			instructions: "Swipe a Diner's Club MSR test card when prompted.",
			args: []string{
				"-type", "charge", "-terminal", terminalName,
				"-test", "-amount", amount(0),
			},
			assert: blockchyp.AuthorizationResponse{
				Success:     true,
				Approved:    true,
				Test:        true,
				EntryMethod: "SWIPE",
				PaymentType: "DINERS",
				MaskedPAN:   notEmpty,
				CardHolder:  notEmpty,
			},
		},
		"MSRJCB": {
			instructions: "Swipe a JCB MSR test card when prompted.",
			args: []string{
				"-type", "charge", "-terminal", terminalName,
				"-test", "-amount", amount(0),
			},
			assert: blockchyp.AuthorizationResponse{
				Success:     true,
				Approved:    true,
				Test:        true,
				EntryMethod: "SWIPE",
				PaymentType: "JCB",
				MaskedPAN:   notEmpty,
				CardHolder:  notEmpty,
			},
		},
		"MSRUnionPay": {
			instructions: "Swipe a UnionPay MSR test card when prompted.",
			args: []string{
				"-type", "charge", "-terminal", terminalName,
				"-test", "-amount", amount(0),
			},
			assert: blockchyp.AuthorizationResponse{
				Success:     true,
				Approved:    true,
				Test:        true,
				EntryMethod: "SWIPE",
				PaymentType: "CUP",
				MaskedPAN:   notEmpty,
				CardHolder:  notEmpty,
			},
		},
		"SignatureInResponse": {
			simOnly:      true,
			instructions: "Insert a signature CVM test card when prompted.",
			args: []string{
				"-type", "charge", "-terminal", terminalName,
				"-test", "-amount", amount(0),
				"-sigFormat", blockchyp.SignatureFormatJPG,
				"-sigWidth", "50",
			},
			assert: blockchyp.AuthorizationResponse{
				Success:  true,
				Approved: true,
				Test:     true,
				SigFile:  notEmpty,
			},
		},
		"SignatureInFile": {
			simOnly:      true,
			instructions: "Insert a signature CVM test card when prompted.",
			args: []string{
				"-type", "charge", "-terminal", terminalName,
				"-test", "-amount", amount(0),
				"-sigWidth", "400", "-sigFile", "/tmp/blockchyp-regression-test/sig.jpg",
			},
			assert: blockchyp.AuthorizationResponse{
				Success:  true,
				Approved: true,
				Test:     true,
			},
			validation: validation{
				prompt: "Does the signature appear valid in the browser?",
				serve:  "/tmp/blockchyp-regression-test/sig.jpg",
				expect: true,
			},
		},
		"SignatureRefused": {
			simOnly: true,
			instructions: `Insert a signature CVM test card when prompted.

When prompted for a signature, hit 'Done' without signing.`,
			args: []string{
				"-type", "charge", "-terminal", terminalName,
				"-test", "-amount", amount(0),
			},
			assert: blockchyp.AuthorizationResponse{
				Success:             false,
				Approved:            false,
				Test:                true,
				ResponseDescription: "Reversed: Customer did not sign",
			},
		},
		"SignatureTimeout": {
			simOnly: true,
			instructions: `Insert a signature CVM test card when prompted.

Let the transaction time out when prompted for a signature. It should take 90 seconds.`,
			args: []string{
				"-type", "charge", "-terminal", terminalName,
				"-test", "-amount", amount(0),
			},
			assert: blockchyp.AuthorizationResponse{
				Success:  false,
				Approved: false,
				Test:     true,
			},
		},
		"SignatureDisabled": {
			instructions: "Insert a signature CVM test card when prompted.",
			args: []string{
				"-type", "charge", "-terminal", terminalName,
				"-test", "-amount", amount(0),
				"-disableSignature",
			},
			assert: blockchyp.AuthorizationResponse{
				Success:  true,
				Approved: true,
				Test:     true,
				ReceiptSuggestions: blockchyp.ReceiptSuggestions{
					RequestSignature: true,
				},
			},
		},
		"UserCanceled": {
			simOnly:      true,
			instructions: "Hit the red 'X' button when prompted for a card.",
			args: []string{
				"-type", "charge", "-terminal", terminalName,
				"-test", "-amount", amount(0),
			},
			assert: blockchyp.AuthorizationResponse{
				Success:             false,
				Approved:            false,
				Test:                true,
				ResponseDescription: "User canceled",
			},
		},
		"ManualApproval": {
			instructions: "Enter PAN '4111 1111 1111 1111' and CVV2 '123' when prompted.",
			args: []string{
				"-type", "charge", "-terminal", terminalName,
				"-test", "-amount", amount(0),
				"-manual",
			},
			assert: blockchyp.AuthorizationResponse{
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
		"ManualDecline": {
			simOnly:      true,
			instructions: "Enter PAN '4111 1111 1111 1129' and CVV2 '123' when prompted.",
			args: []string{
				"-type", "charge", "-terminal", terminalName,
				"-test", "-amount", amount(0),
				"-manual",
			},
			assert: blockchyp.AuthorizationResponse{
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
		"EMVDecline": {
			simOnly:      true,
			instructions: "Insert an EMV test card when prompted.",
			args: []string{
				"-type", "charge", "-terminal", terminalName,
				"-test", "-amount", declineTriggerAmount,
			},
			assert: blockchyp.AuthorizationResponse{
				Success:          true,
				Approved:         false,
				Test:             true,
				TransactionType:  "charge",
				RequestedAmount:  declineTriggerAmount,
				AuthorizedAmount: "0.00",
			},
		},
		"EMVTimeout": {
			simOnly:      true,
			instructions: "Insert an EMV test card when prompted.",
			args: []string{
				"-type", "charge", "-terminal", terminalName,
				"-test", "-amount", timeOutTriggerAmount,
			},
			assert: blockchyp.AuthorizationResponse{
				Success:             false,
				Approved:            false,
				Test:                true,
				ResponseDescription: "Reversed: Network problem",
				TransactionType:     "charge",
				RequestedAmount:     timeOutTriggerAmount,
				AuthorizedAmount:    "0.00",
			},
		},
		"EMVPartialAuth": {
			simOnly:      true,
			instructions: "Insert an EMV test card when prompted.",
			args: []string{
				"-type", "charge", "-terminal", terminalName,
				"-test", "-amount", partialAuthTriggerAmount,
			},
			assert: blockchyp.AuthorizationResponse{
				Success:          true,
				Approved:         true,
				Test:             true,
				PartialAuth:      true,
				TransactionType:  "charge",
				RequestedAmount:  partialAuthTriggerAmount,
				AuthorizedAmount: partialAuthAuthorizedAmount,
			},
		},
		"EMVError": {
			simOnly:      true,
			instructions: "Insert an EMV test card when prompted.",
			args: []string{
				"-type", "charge", "-terminal", terminalName,
				"-test", "-amount", errorTriggerAmount,
			},
			assert: blockchyp.AuthorizationResponse{
				Success:             false,
				Approved:            false,
				Test:                true,
				TransactionType:     "charge",
				ResponseDescription: notEmpty,
				RequestedAmount:     errorTriggerAmount,
				AuthorizedAmount:    "0.00",
			},
		},
		"EMVNoResponse": {
			simOnly:      true,
			instructions: "Insert an EMV test card when prompted.",
			args: []string{
				"-type", "charge", "-terminal", terminalName,
				"-test", "-amount", noResponseTriggerAmount,
			},
			assert: blockchyp.AuthorizationResponse{
				Success:             false,
				Approved:            false,
				Test:                true,
				TransactionType:     "charge",
				ResponseDescription: "Reversed: Network problem",
				RequestedAmount:     noResponseTriggerAmount,
				AuthorizedAmount:    "0.00",
			},
		},
		"SFUnderLimit": {
			simOnly:      true,
			localOnly:    true,
			instructions: "Tap a contactless EMV test card when prompted.",
			args: []string{
				"-type", "charge", "-terminal", terminalName,
				"-test", "-amount", "7.77",
			},
			assert: blockchyp.AuthorizationResponse{
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
		"SFOverLimit": {
			simOnly:      true,
			localOnly:    true,
			instructions: "Tap a contactless EMV test card when prompted.",
			args: []string{
				"-type", "charge", "-terminal", terminalName,
				"-test", "-amount", "77.77",
			},
			assert: blockchyp.AuthorizationResponse{
				Success:          false,
				Approved:         false,
				Test:             true,
				TransactionType:  "charge",
				RequestedAmount:  "77.77",
				AuthorizedAmount: "0.00",
			},
		},
		"EMVCashBack": {
			instructions: `Insert an EMV debit card when prompted.

Select $1 when prompted for cash back.`,
			args: []string{
				"-type", "charge", "-terminal", terminalName,
				"-test", "-amount", amount(0),
				"-cashback",
			},
			assert: blockchyp.AuthorizationResponse{
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
		"MSRCashBackSupported": {
			instructions: `Swipe the Visa MSR test card when prompted. Enter PIN '1234'.

Select $10 when prompted for cash back.`,
			args: []string{
				"-type", "charge", "-terminal", terminalName,
				"-test", "-amount", amount(0),
				"-debit", "-cashback",
			},
			assert: blockchyp.AuthorizationResponse{
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
		"MSRCashBackNotSupported": {
			instructions: "Swipe the AMEX MSR test card when prompted. Enter PIN '1234'.",
			args: []string{
				"-type", "charge", "-terminal", terminalName,
				"-test", "-amount", amount(0),
				"-debit", "-cashback",
			},
			assert: blockchyp.AuthorizationResponse{
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
			validation: validation{
				prompt: "Did the PIN entry screen show cash back options?",
				expect: false,
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if test.simOnly && acquirerMode {
				t.Skip("skipped for acquirer test run")
			}

			cli := newCLI(t)

			if test.localOnly {
				cli.skipCloudRelay()
			}

			setup(t, test.instructions, true)

			cli.run(test.args, test.assert)

			validate(t, test.validation)
		})
	}
}
