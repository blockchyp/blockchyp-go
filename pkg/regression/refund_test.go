// +build regression

package regression

import (
	"testing"

	"github.com/blockchyp/blockchyp-go"
)

func TestRefund(t *testing.T) {
	tests := map[string]struct {
		instructions string
		args         [][]string
		assert       []interface{}
		txID         string
		validation   validation

		// localMode causes tests to be skipped when running in cloud relay
		// mode.
		localMode bool
	}{
		"Charge": {
			instructions: "Insert an EMV test card when prompted.",
			args: [][]string{
				{
					"-type", "charge", "-terminal", "Test Terminal", "-test",
					"-amount", amount(0),
				},
				{
					"-type", "refund", "-test",
					"-tx",
				},
			},
			assert: []interface{}{
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "charge",
					AuthorizedAmount: amount(0),
				},
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "refund",
					AuthorizedAmount: amount(0),
				},
			},
		},
		"Partial": {
			instructions: "Insert an EMV test card when prompted.",
			args: [][]string{
				{
					"-type", "charge", "-terminal", "Test Terminal", "-test",
					"-amount", amountRange(0, 5.01, 9.99),
				},
				{
					"-type", "refund", "-test",
					"-amount", "5.00",
					"-tx",
				},
				{
					"-type", "refund", "-test",
					"-amount", "10.00",
					"-tx",
				},
			},
			assert: []interface{}{
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "charge",
					AuthorizedAmount: amount(0),
				},
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "refund",
					AuthorizedAmount: "5.00",
				},
				blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					TransactionType:     "refund",
					AuthorizedAmount:    "0.00",
					ResponseDescription: "Refund would exceed the original transaction amount",
				},
			},
		},
		"Excess": {
			instructions: "Insert an EMV test card when prompted.",
			args: [][]string{
				{
					"-type", "charge", "-terminal", "Test Terminal", "-test",
					"-amount", amountRange(0, 5, 10),
				},
				{
					"-type", "refund", "-test",
					"-amount", amountRange(1, 10.01, 20),
					"-tx",
				},
			},
			assert: []interface{}{
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "charge",
					AuthorizedAmount: amount(0),
				},
				blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					TransactionType:     "refund",
					AuthorizedAmount:    "0.00",
					ResponseDescription: "Refund would exceed the original transaction amount",
				},
			},
		},
		"SF": {
			localMode:    true,
			instructions: "Insert an EMV test card when prompted.",
			args: [][]string{
				{
					"-type", "refund", "-terminal", "Test Terminal", "-test",
					"-amount", "7.77",
				},
			},
			assert: []interface{}{
				blockchyp.AuthorizationResponse{
					Success:          false,
					Approved:         false,
					Test:             true,
					TransactionType:  "refund",
					RequestedAmount:  "7.77",
					AuthorizedAmount: "0.00",
				},
			},
		},
		"BadCredentials": {
			args: [][]string{
				{
					"-type", "refund", "-test",
					"-tx", "OFE3TTQFJ4I6TNTUNSLM7WZLHE",
					"-apiKey", "X6N2KIQEWYI6TCADNSLM7WZLHE",
					"-bearerToken", "RIKLAPSMSMG2YII27N2NPAMCS5",
					"-signingKey", "4b556bc4e73ffc86fc5f8bfbba1598e7a8cd91f44fd7072d070c92fae7f48cd9",
				},
			},
			assert: []interface{}{
				blockchyp.AuthorizationResponse{
					Success:  false,
					Approved: false,
				},
			},
		},
		"EMVFreeRange": {
			instructions: `Insert an EMV test card when prompted.

Leave the card in the terminal until the test completes.`,
			args: [][]string{
				{
					"-type", "refund", "-terminal", "Test Terminal", "-test",
					"-amount", amount(0),
				},
				{
					"-type", "refund", "-terminal", "Test Terminal", "-test",
					"-amount", amount(0),
				},
			},
			assert: []interface{}{
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "refund",
					RequestedAmount:  amount(0),
					AuthorizedAmount: amount(0),
					EntryMethod:      "CHIP",
				},
				blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					TransactionType:     "refund",
					RequestedAmount:     amount(0),
					AuthorizedAmount:    "0.00",
					EntryMethod:         "CHIP",
					ResponseDescription: "Duplicate Transaction",
				},
			},
		},
		"SignatureInResponse": {
			instructions: "Insert a signature CVM test card when prompted.",
			args: [][]string{
				{
					"-type", "refund", "-terminal", "Test Terminal",
					"-test", "-amount", partialAuthTriggerAmount,
					"-sigFormat", blockchyp.SignatureFormatJPG,
					"-sigWidth", "50",
				},
			},
			assert: []interface{}{
				blockchyp.AuthorizationResponse{
					Success:  true,
					Approved: true,
					Test:     true,
					SigFile:  notEmpty,
				},
			},
		},
		"SignatureInFile": {
			instructions: "Insert a signature CVM test card when prompted.",
			args: [][]string{
				{
					"-type", "refund", "-terminal", "Test Terminal",
					"-test", "-amount", amount(0),
					"-sigWidth", "400", "-sigFile", "/tmp/blockchyp-regression-test/sig.jpg",
				},
			},
			assert: []interface{}{
				blockchyp.AuthorizationResponse{
					Success:  true,
					Approved: true,
					Test:     true,
				},
			},
			validation: validation{
				prompt: "Does the signature appear valid in the browser?",
				expect: true,
			},
		},
		"SignatureRefused": {
			instructions: `Insert a signature CVM test card when prompted.

When prompted for a signature, hit 'Done' without signing.`,
			args: [][]string{
				{
					"-type", "charge", "-terminal", "Test Terminal",
					"-test", "-amount", amount(0),
				},
			},
			assert: []interface{}{
				blockchyp.AuthorizationResponse{
					Success:             false,
					Approved:            false,
					Test:                true,
					ResponseDescription: "Transaction was reversed because the customer did not sign",
				},
			},
		},
		"SignatureTimeout": {
			instructions: `Insert a signature CVM test card when prompted.

Let the transaction time out when prompted for a signature. It should take 90 seconds.`,
			args: [][]string{
				{
					"-type", "charge", "-terminal", "Test Terminal",
					"-test", "-amount", amount(0),
				},
			},
			assert: []interface{}{
				blockchyp.AuthorizationResponse{
					Success:  true,
					Approved: false,
					Test:     true,
				},
			},
		},
		"SignatureDisabled": {
			instructions: "Insert a signature CVM test card when prompted.",
			args: [][]string{
				{
					"-type", "charge", "-terminal", "Test Terminal",
					"-test", "-amount", amount(0),
					"-disableSignature",
				},
			},
			assert: []interface{}{
				blockchyp.AuthorizationResponse{
					Success:  true,
					Approved: true,
					Test:     true,
					ReceiptSuggestions: blockchyp.ReceiptSuggestions{
						RequestSignature: true,
					},
				},
			},
		},
		"UserCanceled": {
			instructions: "Hit the red 'X' button when prompted for a card.",
			args: [][]string{
				{
					"-type", "charge", "-terminal", "Test Terminal",
					"-test", "-amount", amount(0),
				},
			},
			assert: []interface{}{
				blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					ResponseDescription: "user canceled",
				},
			},
		},
		"MSRFreeRange": {
			instructions: "Swipe an MSR test card when prompted.",
			args: [][]string{
				{
					"-type", "refund", "-terminal", "Test Terminal", "-test",
					"-amount", amount(0),
				},
			},
			assert: []interface{}{
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "refund",
					EntryMethod:      "SWIPE",
					AuthorizedAmount: amount(0),
				},
			},
		},
		"ManualFreeRange": {
			instructions: "Enter PAN '4111 1111 1111 1111' and CVV2 '1234' when prompted.",
			args: [][]string{
				{
					"-type", "refund", "-terminal", "Test Terminal", "-test",
					"-amount", amount(0), "-manual",
				},
			},
			assert: []interface{}{
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "refund",
					EntryMethod:      "KEYED",
					AuthorizedAmount: amount(0),
				},
			},
		},
		"DeclineFreeRange": {
			instructions: "Swipe the 'Decline' MSR test card when prompted.",
			args: [][]string{
				{
					"-type", "refund", "-terminal", "Test Terminal", "-test",
					"-amount", amount(0),
				},
			},
			assert: []interface{}{
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         false,
					Test:             true,
					TransactionType:  "refund",
					EntryMethod:      "SWIPE",
					RequestedAmount:  amount(0),
					AuthorizedAmount: "0.00",
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

			for i := range test.args {
				if test.txID != "" && test.args[i][len(test.args[i])-1] == "-tx" {
					test.args[i] = append(test.args[i], test.txID)
				}

				if res, ok := cli.run(test.args[i], test.assert[i]).(*blockchyp.AuthorizationResponse); ok && test.txID == "" {
					test.txID = res.TransactionID
				}
			}

			validate(t, test.validation)
		})
	}
}
