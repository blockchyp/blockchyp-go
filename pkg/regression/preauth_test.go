// +build regression

package regression

import (
	"testing"

	"github.com/blockchyp/blockchyp-go"
)

func TestPreauth(t *testing.T) {
	tests := map[string]struct {
		instructions string

		authArgs   []string
		authAssert blockchyp.AuthorizationResponse

		voidArgs   []string
		voidAssert blockchyp.VoidResponse

		closeBatchArgs   []string
		closeBatchAssert blockchyp.CloseBatchResponse

		captureArgs   []string
		captureAssert blockchyp.CaptureResponse

		reCaptureArgs   []string
		reCaptureAssert blockchyp.CaptureResponse

		txID       string
		validation validation
	}{
		"EMVApproved": {
			instructions: "Insert an EMV test card when prompted.",
			authArgs: []string{
				"-type", "preauth", "-terminal", "Test Terminal",
				"-test", "-amount", amount(0),
			},
			authAssert: blockchyp.AuthorizationResponse{
				Success:          true,
				Approved:         true,
				Test:             true,
				TransactionType:  "preauth",
				RequestedAmount:  amount(0),
				AuthorizedAmount: amount(0),
			},
			captureArgs: []string{
				"-type", "capture", "-test",
				"-tx",
			},
			captureAssert: blockchyp.CaptureResponse{
				Success:          true,
				Approved:         true,
				Test:             true,
				TransactionType:  "capture",
				AuthorizedAmount: amount(0),
			},
			reCaptureArgs: []string{
				"-type", "capture", "-test",
				"-tx",
			},
			reCaptureAssert: blockchyp.CaptureResponse{
				Success:             true,
				Approved:            false,
				Test:                true,
				TransactionType:     "capture",
				ResponseDescription: "Already Captured",
			},
		},
		"SignatureInResponse": {
			instructions: "Insert a signature CVM test card when prompted.",
			authArgs: []string{
				"-type", "preauth", "-terminal", "Test Terminal",
				"-test", "-amount", amount(0),
				"-sigFormat", blockchyp.SignatureFormatPNG,
				"-sigWidth", "50",
			},
			authAssert: blockchyp.AuthorizationResponse{
				Success:         true,
				Approved:        true,
				Test:            true,
				TransactionType: "preauth",
				SigFile:         notEmpty,
			},
		},
		"SignatureInFile": {
			instructions: "Insert a signature CVM test card when prompted",
			authArgs: []string{
				"-type", "preauth", "-terminal", "Test Terminal",
				"-test", "-amount", amount(0),
				"-sigWidth", "400", "-sigFile", "/tmp/blockchyp-regression-test/sig.png",
			},
			authAssert: blockchyp.AuthorizationResponse{
				Success:         true,
				Approved:        true,
				Test:            true,
				TransactionType: "preauth",
			},
			validation: validation{
				prompt: "Does the signature appear valid in the browser?",
				expect: true,
			},
		},
		"SignatureRefused": {
			instructions: `Insert a signature CVM test card when prompted.

When prompted for a signature, hit 'Done' without signing.`,
			authArgs: []string{
				"-type", "preauth", "-terminal", "Test Terminal",
				"-test", "-amount", amount(0),
			},
			authAssert: blockchyp.AuthorizationResponse{
				Success:             false,
				Approved:            false,
				Test:                true,
				ResponseDescription: "Transaction was reversed because the customer did not sign",
			},
		},
		"UserCanceled": {
			instructions: "Hit the red 'X' button when prompted for a card.",
			authArgs: []string{
				"-type", "preauth", "-terminal", "Test Terminal",
				"-test", "-amount", amount(0),
			},
			authAssert: blockchyp.AuthorizationResponse{
				Success:             true,
				Approved:            false,
				Test:                true,
				ResponseDescription: "user canceled",
			},
		},
		"SignatureTimeout": {
			instructions: `Insert a signature CVM test card when prompted.

Let the transaction time out when prompted for a signature. It should take 90 seconds.`,
			authArgs: []string{
				"-type", "preauth", "-terminal", "Test Terminal",
				"-test", "-amount", amount(0),
			},
			authAssert: blockchyp.AuthorizationResponse{
				Success:  true,
				Approved: false,
				Test:     true,
			},
		},
		"ManualApproval": {
			instructions: "Enter PAN '4111 1111 1111 1111' and CVV2 '1234' when prompted",
			authArgs: []string{
				"-type", "preauth", "-terminal", "Test Terminal",
				"-test", "-amount", amount(0),
				"-manual",
			},
			authAssert: blockchyp.AuthorizationResponse{
				Success:          true,
				Approved:         true,
				Test:             true,
				TransactionType:  "preauth",
				RequestedAmount:  amount(0),
				AuthorizedAmount: amount(0),
				EntryMethod:      "KEYED",
				MaskedPAN:        "************1111",
			},
			captureArgs: []string{
				"-type", "capture", "-test",
				"-tx",
			},
			captureAssert: blockchyp.CaptureResponse{
				Success:          true,
				Approved:         true,
				Test:             true,
				TransactionType:  "capture",
				AuthorizedAmount: amount(0),
			},
		},
		"EMVDecline": {
			instructions: "Insert any EMV test card.",
			authArgs: []string{
				"-type", "preauth", "-terminal", "Test Terminal",
				"-test", "-amount", declineTriggerAmount,
			},
			authAssert: blockchyp.AuthorizationResponse{
				Success:          true,
				Approved:         false,
				Test:             true,
				TransactionType:  "preauth",
				RequestedAmount:  declineTriggerAmount,
				AuthorizedAmount: "0.00",
			},
			captureArgs: []string{
				"-type", "capture", "-test",
				"-tx",
			},
			captureAssert: blockchyp.CaptureResponse{
				Success:         false,
				Approved:        false,
				Test:            true,
				TransactionType: "capture",
			},
		},
		"EMVTimeout": {
			instructions: "Insert any EMV test card.",
			authArgs: []string{
				"-type", "preauth", "-terminal", "Test Terminal",
				"-test", "-amount", timeOutTriggerAmount,
			},
			authAssert: blockchyp.AuthorizationResponse{
				Success:             false,
				Approved:            false,
				Test:                true,
				ResponseDescription: "Transaction was reversed because there was a problem during authorization",
				TransactionType:     "preauth",
				RequestedAmount:     timeOutTriggerAmount,
				AuthorizedAmount:    "0.00",
			},
		},
		"EMVPartialAuth": {
			instructions: "Insert any EMV test card.",
			authArgs: []string{
				"-type", "preauth", "-terminal", "Test Terminal",
				"-test", "-amount", partialAuthTriggerAmount,
			},
			authAssert: blockchyp.AuthorizationResponse{
				Success:          true,
				Approved:         true,
				Test:             true,
				PartialAuth:      true,
				TransactionType:  "preauth",
				RequestedAmount:  partialAuthTriggerAmount,
				AuthorizedAmount: partialAuthAuthorizedAmount,
			},
		},
		"EMVError": {
			instructions: "Insert any EMV test card.",
			authArgs: []string{
				"-type", "preauth", "-terminal", "Test Terminal",
				"-test", "-amount", errorTriggerAmount,
			},
			authAssert: blockchyp.AuthorizationResponse{
				Success:             false,
				Approved:            false,
				Test:                true,
				TransactionType:     "preauth",
				ResponseDescription: notEmpty,
				RequestedAmount:     errorTriggerAmount,
				AuthorizedAmount:    "0.00",
			},
		},
		"EMVNoResponse": {
			instructions: "Insert any EMV test card.",
			authArgs: []string{
				"-type", "preauth", "-terminal", "Test Terminal",
				"-test", "-amount", noResponseTriggerAmount,
			},
			authAssert: blockchyp.AuthorizationResponse{
				Success:             false,
				Approved:            false,
				Test:                true,
				TransactionType:     "preauth",
				ResponseDescription: "Transaction was reversed because there was a problem during authorization",
				RequestedAmount:     noResponseTriggerAmount,
				AuthorizedAmount:    "0.00",
			},
		},
		"TipAdjust": {
			instructions: "Insert an EMV test card when prompted.",
			authArgs: []string{
				"-type", "preauth", "-terminal", "Test Terminal",
				"-test", "-amount", amount(0),
			},
			authAssert: blockchyp.AuthorizationResponse{
				Success:          true,
				Approved:         true,
				Test:             true,
				TransactionType:  "preauth",
				RequestedAmount:  amount(0),
				AuthorizedAmount: amount(0),
			},
			captureArgs: []string{
				"-type", "capture", "-test",
				"-tip", amount(1), "-amount", amount(2),
				"-tx",
			},
			captureAssert: blockchyp.CaptureResponse{
				Success:          true,
				Approved:         true,
				Test:             true,
				TransactionType:  "capture",
				AuthorizedAmount: amount(2),
				TipAmount:        amount(1),
			},
		},
		"OrphanCapture": {
			captureArgs: []string{
				"-type", "capture", "-test",
				"-tx",
			},
			captureAssert: blockchyp.CaptureResponse{
				Success:             false,
				Approved:            false,
				Test:                true,
				TransactionType:     "capture",
				ResponseDescription: "Invalid Transaction",
			},
			txID: "FAKE",
		},
		"VoidCapture": {
			instructions: "Insert an EMV test card when prompted.",
			authArgs: []string{
				"-type", "preauth", "-terminal", "Test Terminal",
				"-test", "-amount", amount(0),
			},
			authAssert: blockchyp.AuthorizationResponse{
				Success:          true,
				Approved:         true,
				Test:             true,
				TransactionType:  "preauth",
				RequestedAmount:  amount(0),
				AuthorizedAmount: amount(0),
			},
			voidArgs: []string{
				"-type", "void", "-test",
				"-tx",
			},
			voidAssert: blockchyp.VoidResponse{
				Success:  true,
				Approved: true,
				Test:     true,
			},
			captureArgs: []string{
				"-type", "capture", "-test",
				"-tx",
			},
			captureAssert: blockchyp.CaptureResponse{
				Success:             false,
				Approved:            false,
				Test:                true,
				TransactionType:     "capture",
				ResponseDescription: "Voided Transaction",
			},
		},
		"ClosedBatchCapture": {
			instructions: "Insert an EMV test card when prompted.",
			authArgs: []string{
				"-type", "preauth", "-terminal", "Test Terminal",
				"-test", "-amount", amount(0),
			},
			authAssert: blockchyp.AuthorizationResponse{
				Success:          true,
				Approved:         true,
				Test:             true,
				TransactionType:  "preauth",
				RequestedAmount:  amount(0),
				AuthorizedAmount: amount(0),
			},
			closeBatchArgs: []string{
				"-type", "close-batch", "-test",
			},
			closeBatchAssert: blockchyp.CloseBatchResponse{
				Success:       true,
				Test:          true,
				CapturedTotal: notEmpty,
				OpenPreauths:  notEmpty,
			},
			captureArgs: []string{
				"-type", "capture", "-test",
				"-tx",
			},
			captureAssert: blockchyp.CaptureResponse{
				Success:          true,
				Approved:         true,
				Test:             true,
				TransactionType:  "capture",
				AuthorizedAmount: amount(0),
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			cli := newCLI(t)

			setup(t, test.instructions, true)

			txID := test.txID

			if len(test.authArgs) > 0 {
				res := cli.run(test.authArgs, test.authAssert).(*blockchyp.AuthorizationResponse)
				txID = res.TransactionID
			}

			if len(test.voidArgs) > 0 {
				test.voidArgs = append(test.voidArgs, txID)

				cli.run(test.voidArgs, test.voidAssert)
			}

			if len(test.closeBatchArgs) > 0 {
				cli.run(test.closeBatchArgs, test.closeBatchAssert)
			}

			if len(test.captureArgs) > 0 {
				test.captureArgs = append(test.captureArgs, txID)

				cli.run(test.captureArgs, test.captureAssert)
			}

			if len(test.reCaptureArgs) > 0 {
				test.reCaptureArgs = append(test.reCaptureArgs, txID)

				cli.run(test.reCaptureArgs, test.reCaptureAssert)
			}

			validate(t, test.validation)
		})
	}
}
