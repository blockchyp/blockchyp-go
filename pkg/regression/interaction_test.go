// +build regression

package regression

import (
	"testing"

	"github.com/blockchyp/blockchyp-go"
)

func TestInteraction(t *testing.T) {
	tests := map[string]struct {
		instructions string
		wait         bool
		args         [][]string
		assert       []interface{}
		validation   []validation
	}{
		"Message": {
			args: [][]string{
				{
					"-type", "message", "-terminal", terminalName, "-test",
					"-message", "Your father was a hamster.",
				},
			},
			assert: []interface{}{
				blockchyp.Acknowledgement{
					Success: true,
				},
			},
			validation: []validation{
				{
					prompt: "Was 'Your father was a hamster.' displayed on the terminal?",
					expect: true,
				},
			},
		},
		"BooleanPrompt": {
			instructions: "When prompted, select 'Machines'. You can never be too careful.",
			args: [][]string{
				{
					"-type", "boolean-prompt", "-terminal", terminalName, "-test",
					"-prompt", "Which side will you take in the machine uprising?",
					"-yesCaption", "Machines",
					"-noCaption", "Humans",
				},
			},
			assert: []interface{}{
				blockchyp.BooleanPromptResponse{
					Success:  true,
					Response: true,
				},
			},
		},
		"EmailPrompt": {
			instructions: "When prompted, enter any text and hit 'Next' or the green 'O' button.",
			args: [][]string{
				{
					"-type", "text-prompt", "-terminal", terminalName, "-test",
					"-promptType", "email",
				},
			},
			assert: []interface{}{
				blockchyp.TextPromptResponse{
					Success:  true,
					Response: notEmpty,
				},
			},
		},
		"PhonePrompt": {
			instructions: "When prompted, enter any number and hit the green 'O' button.",
			args: [][]string{
				{
					"-type", "text-prompt", "-terminal", terminalName, "-test",
					"-promptType", "phone",
				},
			},
			assert: []interface{}{
				blockchyp.TextPromptResponse{
					Success:  true,
					Response: notEmpty,
				},
			},
		},
		"CustomerNumberPrompt": {
			instructions: "When prompted, enter any number and hit the green 'O' button.",
			args: [][]string{
				{
					"-type", "text-prompt", "-terminal", terminalName, "-test",
					"-promptType", "customer-number",
				},
			},
			assert: []interface{}{
				blockchyp.TextPromptResponse{
					Success:  true,
					Response: notEmpty,
				},
			},
		},
		"RewardsNumberPrompt": {
			instructions: "When prompted, enter any number and hit the green 'O' button.",
			args: [][]string{
				{
					"-type", "text-prompt", "-terminal", terminalName, "-test",
					"-promptType", "rewards-number",
				},
			},
			assert: []interface{}{
				blockchyp.TextPromptResponse{
					Success:  true,
					Response: notEmpty,
				},
			},
		},
		"FirstNamePrompt": {
			instructions: "When prompted, enter any text and hit 'Next' or the green 'O' button.",
			args: [][]string{
				{
					"-type", "text-prompt", "-terminal", terminalName, "-test",
					"-promptType", "first-name",
				},
			},
			assert: []interface{}{
				blockchyp.TextPromptResponse{
					Success:  true,
					Response: notEmpty,
				},
			},
		},
		"LastNamePrompt": {
			instructions: "When prompted, enter any text and hit 'Next' or the green 'O' button.",
			args: [][]string{
				{
					"-type", "text-prompt", "-terminal", terminalName, "-test",
					"-promptType", "last-name",
				},
			},
			assert: []interface{}{
				blockchyp.TextPromptResponse{
					Success:  true,
					Response: notEmpty,
				},
			},
		},
		"SignatureCapture": {
			instructions: "When prompted, sign and hit 'Done'.",
			args: [][]string{
				{
					"-type", "capture-signature", "-terminal", terminalName, "-test",
					"-sigFormat", blockchyp.SignatureFormatJPG,
					"-sigWidth", "50",
				},
			},
			assert: []interface{}{
				blockchyp.CaptureSignatureResponse{
					Success: true,
					SigFile: notEmpty,
				},
			},
		},
		"LineItemDisplay": {
			args: [][]string{
				{
					"-type", "display", "-terminal", terminalName, "-test",
					"-displaySubtotal", "120.05",
					"-displayTax", "5.00",
					"-displayTotal", "125.05",
					"-lineItemDescription", "Leki Trekking Poles",
					"-lineItemQty", "1",
					"-lineItemPrice", "135.05",
					"-lineItemDiscountDescription", "Member Discount",
					"-lineItemDiscountAmount", "10.00",
					"-lineItemExtended", "120.05",
				},
				{
					"-type", "clear", "-terminal", terminalName, "-test",
				},
			},
			assert: []interface{}{
				blockchyp.Acknowledgement{
					Success: true,
				},
				blockchyp.Acknowledgement{
					Success: true,
				},
			},
			validation: []validation{
				{
					prompt: "Was a transaction with a discount displayed on the terminal?",
					expect: true,
				},
				{
					prompt: "Was the transaction cleared?",
					expect: true,
				},
			},
		},
		"WhitelistedBIN": {
			instructions: `From the admin console, set the whitelisted BIN range to the first few numbers of a test card.

When prompted for a card, use the test card you whitelisted.`,
			args: [][]string{
				{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", "1.23",
				},
			},
			assert: []interface{}{
				blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					ResponseDescription: "WHITELISTED",
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			cli := newCLI(t)

			setup(t, test.instructions, true)

			for i := range test.args {
				cli.run(test.args[i], test.assert[i])

				if len(test.validation) > i {
					validate(t, test.validation[i])
				}
			}
		})
	}
}
