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

		// localMode causes tests to be skipped when running in cloud relay
		// mode.
		localMode bool
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
				"-test", "-amount", amountRange(0, 1, 10),
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
				expect: true,
			},
		},
		"SignatureRefused": {
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
				ResponseDescription: "Transaction was reversed because the customer did not sign",
			},
		},
		"SignatureTimeout": {
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
			instructions: "Hit the red 'X' button when prompted for a card.",
			args: []string{
				"-type", "charge", "-terminal", terminalName,
				"-test", "-amount", amount(0),
			},
			assert: blockchyp.AuthorizationResponse{
				Success:             false,
				Approved:            false,
				Test:                true,
				ResponseDescription: "user canceled",
			},
		},
		"ManualApproval": {
			instructions: "Enter PAN '4111 1111 1111 1111' and CVV2 '1234' when prompted.",
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
			instructions: "Enter PAN '4111 1111 1111 1129' and CVV2 '1234' when prompted.",
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
			instructions: "Insert an EMV test card when prompted.",
			args: []string{
				"-type", "charge", "-terminal", terminalName,
				"-test", "-amount", timeOutTriggerAmount,
			},
			assert: blockchyp.AuthorizationResponse{
				Success:             false,
				Approved:            false,
				Test:                true,
				ResponseDescription: "Transaction was reversed because there was a problem during authorization",
				TransactionType:     "charge",
				RequestedAmount:     timeOutTriggerAmount,
				AuthorizedAmount:    "0.00",
			},
		},
		"EMVPartialAuth": {
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
				ResponseDescription: "Transaction was reversed because there was a problem during authorization",
				RequestedAmount:     noResponseTriggerAmount,
				AuthorizedAmount:    "0.00",
			},
		},
		"SFUnderLimit": {
			localMode:    true,
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
			localMode:    true,
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
		"CashBack": {
			instructions: `Insert a debit card when prompted.

Select $1 when prompted for cash back.`,
			args: []string{
				"-type", "charge", "-terminal", terminalName,
				"-test", "-amount", "5.00",
				"-cashback",
			},
			assert: blockchyp.AuthorizationResponse{
				Success:                  true,
				Approved:                 true,
				Test:                     true,
				TransactionType:          "charge",
				RequestedAmount:          "5.00",
				RequestedCashBackAmount:  "1.00",
				AuthorizedAmount:         "6.00",
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
					AuthorizedAmount: "6.00",
					CashBackAmount:   "1.00",
					PINVerified:      true,
					TransactionType:  "charge",
					EntryMethod:      "CHIP",
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			cli := newCLI(t)
			if test.localMode {
				cli.skipCloudRelay()
			}

			setup(t, test.instructions, true)

			cli.run(test.args, test.assert)

			validate(t, test.validation)
		})
	}
}
