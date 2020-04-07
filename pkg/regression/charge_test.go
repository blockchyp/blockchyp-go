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
	}{
		"ContactEMVNoCVMApproved": {
			instructions: "Insert a No-CVM EMV test card when prompted",
			args: []string{
				"-type", "charge", "-terminal", "Test Terminal",
				"-test", "-amount", "59.00",
			},
			assert: blockchyp.AuthorizationResponse{
				Success:          true,
				Approved:         true,
				Test:             true,
				TransactionType:  "charge",
				RequestedAmount:  "59.00",
				AuthorizedAmount: "59.00",
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
					AuthorizedAmount: "59.00",
					TransactionType:  "charge",
					EntryMethod:      "CHIP",
				},
			},
		},
		"ContactlessEMVApproved": {
			instructions: "Tap a valid contactless EMV test card when prompted",
			args: []string{
				"-type", "charge", "-terminal", "Test Terminal",
				"-test", "-amount", "69.00",
			},
			assert: blockchyp.AuthorizationResponse{
				Success:          true,
				Approved:         true,
				Test:             true,
				TransactionType:  "charge",
				RequestedAmount:  "69.00",
				AuthorizedAmount: "69.00",
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
					AuthorizedAmount: "69.00",
					TransactionType:  "charge",
					EntryMethod:      "CONTACTLESS EMV",
				},
			},
		},
		"MSRVisa": {
			instructions: "Swipe a Visa MSR test card when prompted",
			args: []string{
				"-type", "charge", "-terminal", "Test Terminal",
				"-test", "-amount", "45.77",
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
			instructions: "Swipe a MasterCard MSR test card when prompted",
			args: []string{
				"-type", "charge", "-terminal", "Test Terminal",
				"-test", "-amount", "45.77",
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
			instructions: "Swipe a Discover MSR test card when prompted",
			args: []string{
				"-type", "charge", "-terminal", "Test Terminal",
				"-test", "-amount", "45.77",
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
			instructions: "Swipe an Amex MSR test card when prompted",
			args: []string{
				"-type", "charge", "-terminal", "Test Terminal",
				"-test", "-amount", "45.77",
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
			instructions: "Swipe a Diner's Club MSR test card when prompted",
			args: []string{
				"-type", "charge", "-terminal", "Test Terminal",
				"-test", "-amount", "45.77",
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
			instructions: "Swipe a JCB MSR test card when prompted",
			args: []string{
				"-type", "charge", "-terminal", "Test Terminal",
				"-test", "-amount", "45.77",
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
			instructions: "Swipe a UnionPay MSR test card when prompted",
			args: []string{
				"-type", "charge", "-terminal", "Test Terminal",
				"-test", "-amount", "45.77",
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
			instructions: "Insert a signature CVM test card when prompted",
			args: []string{
				"-type", "charge", "-terminal", "Test Terminal",
				"-test", "-amount", "58.00",
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
			instructions: "Insert a signature CVM test card when prompted",
			args: []string{
				"-type", "charge", "-terminal", "Test Terminal",
				"-test", "-amount", "58.00",
				"-sigWidth", "100", "-sigFile", "/tmp/sig.jpg",
			},
			assert: blockchyp.AuthorizationResponse{
				Success:  true,
				Approved: true,
				Test:     true,
			},
			validation: validation{
				prompt: "Does '/tmp/sig.jpg' contain the signature you entered on the terminal?",
				expect: true,
			},
		},
		"SignatureRefused": {
			instructions: `Insert a signature CVM test card when prompted.

When prompted for a signature, hit 'Done' without signing.`,
			args: []string{
				"-type", "charge", "-terminal", "Test Terminal",
				"-test", "-amount", "57.00",
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
				"-type", "charge", "-terminal", "Test Terminal",
				"-test", "-amount", "55.01",
			},
			assert: blockchyp.AuthorizationResponse{
				Success:             false,
				Approved:            false,
				Test:                true,
				ResponseDescription: "context canceled",
			},
		},
		"SignatureDisabled": {
			instructions: "Insert a signature CVM test card when prompted.",
			args: []string{
				"-type", "charge", "-terminal", "Test Terminal",
				"-test", "-amount", "61.00",
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
				"-type", "charge", "-terminal", "Test Terminal",
				"-test", "-amount", "56.00",
			},
			assert: blockchyp.AuthorizationResponse{
				Success:             false,
				Approved:            false,
				Test:                true,
				ResponseDescription: "user canceled",
			},
		},
		"ManualApproval": {
			instructions: "Enter PAN '4111 1111 1111 1111' and CVV2 '1234' when prompted",
			args: []string{
				"-type", "charge", "-terminal", "Test Terminal",
				"-test", "-amount", "54.00",
				"-manual",
			},
			assert: blockchyp.AuthorizationResponse{
				Success:          true,
				Approved:         true,
				Test:             true,
				TransactionType:  "charge",
				RequestedAmount:  "54.00",
				AuthorizedAmount: "54.00",
				EntryMethod:      "KEYED",
				MaskedPAN:        "************1111",
			},
		},
		"ManualDecline": {
			instructions: "Enter PAN '4111 1111 1111 1129' and CVV2 '1234' when prompted",
			args: []string{
				"-type", "charge", "-terminal", "Test Terminal",
				"-test", "-amount", "59.00",
				"-manual",
			},
			assert: blockchyp.AuthorizationResponse{
				Success:          true,
				Approved:         false,
				Test:             true,
				TransactionType:  "charge",
				RequestedAmount:  "59.00",
				AuthorizedAmount: "0.00",
				EntryMethod:      "KEYED",
				MaskedPAN:        "************1129",
			},
		},
		"EMVDecline": {
			instructions: "Insert any EMV test card.",
			args: []string{
				"-type", "charge", "-terminal", "Test Terminal",
				"-test", "-amount", "201.00",
			},
			assert: blockchyp.AuthorizationResponse{
				Success:          true,
				Approved:         false,
				Test:             true,
				TransactionType:  "charge",
				RequestedAmount:  "201.00",
				AuthorizedAmount: "0.00",
			},
		},
		"EMVTimeout": {
			instructions: "Insert any EMV test card.",
			args: []string{
				"-type", "charge", "-terminal", "Test Terminal",
				"-test", "-amount", "68.00",
			},
			assert: blockchyp.AuthorizationResponse{
				Success:             false,
				Approved:            false,
				Test:                true,
				ResponseDescription: "Transaction was reversed because there was a problem during authorization",
				TransactionType:     "charge",
				RequestedAmount:     "68.00",
				AuthorizedAmount:    "0.00",
			},
		},
		"EMVPartialAuth": {
			instructions: "Insert any EMV test card.",
			args: []string{
				"-type", "charge", "-terminal", "Test Terminal",
				"-test", "-amount", "55.00",
			},
			assert: blockchyp.AuthorizationResponse{
				Success:          true,
				Approved:         true,
				Test:             true,
				PartialAuth:      true,
				TransactionType:  "charge",
				RequestedAmount:  "55.00",
				AuthorizedAmount: "25.00",
			},
		},
		"EMVError": {
			instructions: "Insert any EMV test card.",
			args: []string{
				"-type", "charge", "-terminal", "Test Terminal",
				"-test", "-amount", "0.11",
			},
			assert: blockchyp.AuthorizationResponse{
				Success:             false,
				Approved:            false,
				Test:                true,
				TransactionType:     "charge",
				ResponseDescription: notEmpty,
				RequestedAmount:     "0.11",
				AuthorizedAmount:    "0.00",
			},
		},
		"EMVNoResponse": {
			instructions: "Insert any EMV test card.",
			args: []string{
				"-type", "charge", "-terminal", "Test Terminal",
				"-test", "-amount", "72.00",
			},
			assert: blockchyp.AuthorizationResponse{
				Success:             false,
				Approved:            false,
				Test:                true,
				TransactionType:     "charge",
				ResponseDescription: "Transaction was reversed because there was a problem during authorization",
				RequestedAmount:     "72.00",
				AuthorizedAmount:    "0.00",
			},
		},
		"SFUnderLimit": {
			instructions: "Tap a contactless EMV test card.",
			args: []string{
				"-type", "charge", "-terminal", "Test Terminal",
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
			instructions: "Tap a contactless EMV test card.",
			args: []string{
				"-type", "charge", "-terminal", "Test Terminal",
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
	}

	cli := newCLI(t)

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			setup(t, test.instructions, true)

			cli.run(test.args, test.assert)

			validate(t, test.validation)
		})
	}
}
