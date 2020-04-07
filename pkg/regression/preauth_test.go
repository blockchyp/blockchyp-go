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
				"-test", "-amount", "79.00",
			},
			authAssert: blockchyp.AuthorizationResponse{
				Success:          true,
				Approved:         true,
				Test:             true,
				TransactionType:  "preauth",
				RequestedAmount:  "79.00",
				AuthorizedAmount: "79.00",
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
				AuthorizedAmount: "79.00",
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
				"-test", "-amount", "79.01",
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
				"-test", "-amount", "58.00",
				"-sigWidth", "100", "-sigFile", "/tmp/sig.png",
			},
			authAssert: blockchyp.AuthorizationResponse{
				Success:         true,
				Approved:        true,
				Test:            true,
				TransactionType: "preauth",
			},
			validation: validation{
				prompt: "Does '/tmp/sig.png' contain the signature you entered on the terminal?",
				expect: true,
			},
		},
		"SignatureRefused": {
			instructions: `Insert a signature CVM test card when prompted.

When prompted for a signature, hit 'Done' without signing.`,
			authArgs: []string{
				"-type", "preauth", "-terminal", "Test Terminal",
				"-test", "-amount", "79.02",
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
				"-test", "-amount", "56.00",
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
				"-test", "-amount", "80.00",
			},
			authAssert: blockchyp.AuthorizationResponse{
				Success:             true,
				Approved:            false,
				Test:                true,
				ResponseDescription: "context canceled",
			},
		},
		"ManualApproval": {
			instructions: "Enter PAN '4111 1111 1111 1111' and CVV2 '1234' when prompted",
			authArgs: []string{
				"-type", "preauth", "-terminal", "Test Terminal",
				"-test", "-amount", "80.01",
				"-manual",
			},
			authAssert: blockchyp.AuthorizationResponse{
				Success:          true,
				Approved:         true,
				Test:             true,
				TransactionType:  "preauth",
				RequestedAmount:  "80.01",
				AuthorizedAmount: "80.01",
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
				AuthorizedAmount: "80.01",
			},
		},
		"EMVDecline": {
			instructions: "Insert any EMV test card.",
			authArgs: []string{
				"-type", "preauth", "-terminal", "Test Terminal",
				"-test", "-amount", "201.00",
			},
			authAssert: blockchyp.AuthorizationResponse{
				Success:          true,
				Approved:         false,
				Test:             true,
				TransactionType:  "preauth",
				RequestedAmount:  "201.00",
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
				"-test", "-amount", "68.00",
			},
			authAssert: blockchyp.AuthorizationResponse{
				Success:             false,
				Approved:            false,
				Test:                true,
				ResponseDescription: "Transaction was reversed because there was a problem during authorization",
				TransactionType:     "preauth",
				RequestedAmount:     "68.00",
				AuthorizedAmount:    "0.00",
			},
		},
		"EMVPartialAuth": {
			instructions: "Insert any EMV test card.",
			authArgs: []string{
				"-type", "preauth", "-terminal", "Test Terminal",
				"-test", "-amount", "55.00",
			},
			authAssert: blockchyp.AuthorizationResponse{
				Success:          true,
				Approved:         true,
				Test:             true,
				PartialAuth:      true,
				TransactionType:  "preauth",
				RequestedAmount:  "55.00",
				AuthorizedAmount: "25.00",
			},
		},
		"EMVError": {
			instructions: "Insert any EMV test card.",
			authArgs: []string{
				"-type", "preauth", "-terminal", "Test Terminal",
				"-test", "-amount", "0.11",
			},
			authAssert: blockchyp.AuthorizationResponse{
				Success:             false,
				Approved:            false,
				Test:                true,
				TransactionType:     "preauth",
				ResponseDescription: notEmpty,
				RequestedAmount:     "0.11",
				AuthorizedAmount:    "0.00",
			},
		},
		"EMVNoResponse": {
			instructions: "Insert any EMV test card.",
			authArgs: []string{
				"-type", "preauth", "-terminal", "Test Terminal",
				"-test", "-amount", "72.00",
			},
			authAssert: blockchyp.AuthorizationResponse{
				Success:             false,
				Approved:            false,
				Test:                true,
				TransactionType:     "preauth",
				ResponseDescription: "Transaction was reversed because there was a problem during authorization",
				RequestedAmount:     "72.00",
				AuthorizedAmount:    "0.00",
			},
		},
		"TipAdjust": {
			instructions: "Insert an EMV test card when prompted.",
			authArgs: []string{
				"-type", "preauth", "-terminal", "Test Terminal",
				"-test", "-amount", "61.00",
			},
			authAssert: blockchyp.AuthorizationResponse{
				Success:          true,
				Approved:         true,
				Test:             true,
				TransactionType:  "preauth",
				RequestedAmount:  "61.00",
				AuthorizedAmount: "61.00",
			},
			captureArgs: []string{
				"-type", "capture", "-test",
				"-tip", "5.00", "-amount", "66.00",
				"-tx",
			},
			captureAssert: blockchyp.CaptureResponse{
				Success:          true,
				Approved:         true,
				Test:             true,
				TransactionType:  "capture",
				AuthorizedAmount: "66.00",
				TipAmount:        "5.00",
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
				"-test", "-amount", "82.00",
			},
			authAssert: blockchyp.AuthorizationResponse{
				Success:          true,
				Approved:         true,
				Test:             true,
				TransactionType:  "preauth",
				RequestedAmount:  "82.00",
				AuthorizedAmount: "82.00",
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
				"-test", "-amount", "83.00",
			},
			authAssert: blockchyp.AuthorizationResponse{
				Success:          true,
				Approved:         true,
				Test:             true,
				TransactionType:  "preauth",
				RequestedAmount:  "83.00",
				AuthorizedAmount: "83.00",
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
				AuthorizedAmount: "83.00",
			},
		},
	}

	cli := newCLI(t)

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
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
