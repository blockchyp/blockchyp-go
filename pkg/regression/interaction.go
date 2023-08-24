package regression

import (
	"github.com/blockchyp/blockchyp-go/v2"
)

var interactionTests = testCases{
	{
		name:  "Interaction/Message",
		group: testGroupInteractive,
		sim:   true,
		operations: []operation{
			{
				args: []string{
					"-type", "message", "-terminal", terminalName, "-test",
					"-message", "Your father was a hamster.",
				},
				expect: blockchyp.Acknowledgement{
					Success: true,
				},
			},
			{
				validation: &validation{
					prompt: "Was 'Your father was a hamster.' displayed on the terminal?",
					expect: true,
				},
			},
		},
	},
	{
		name:  "Interaction/BooleanPrompt",
		group: testGroupInteractive,
		sim:   true,
		operations: []operation{
			{
				msg: "When prompted, select 'Machines'. You can never be too careful.",
				args: []string{
					"-type", "boolean-prompt", "-terminal", terminalName, "-test",
					"-prompt", "Which side will you take in the machine uprising?",
					"-yesCaption", "Machines",
					"-noCaption", "Humans",
				},
				expect: blockchyp.BooleanPromptResponse{
					Success:  true,
					Response: true,
				},
			},
		},
	},
	{
		name:  "Interaction/EmailPrompt",
		group: testGroupInteractive,
		sim:   true,
		operations: []operation{
			{
				msg: "When prompted, enter any text and hit 'Next' or the green 'O' button.",
				args: []string{
					"-type", "text-prompt", "-terminal", terminalName, "-test",
					"-promptType", "email",
				},
				expect: blockchyp.TextPromptResponse{
					Success:  true,
					Response: notEmpty,
				},
			},
		},
	},
	{
		name:  "Interaction/PhonePrompt",
		group: testGroupInteractive,
		sim:   true,
		operations: []operation{
			{
				msg: "When prompted, enter any number and hit the green 'O' button.",
				args: []string{
					"-type", "text-prompt", "-terminal", terminalName, "-test",
					"-promptType", "phone",
				},
				expect: blockchyp.TextPromptResponse{
					Success:  true,
					Response: notEmpty,
				},
			},
		},
	},
	{
		name:  "Interaction/CustomerNumberPrompt",
		group: testGroupInteractive,
		sim:   true,
		operations: []operation{
			{
				msg: "When prompted, enter any number and hit the green 'O' button.",
				args: []string{
					"-type", "text-prompt", "-terminal", terminalName, "-test",
					"-promptType", "customer-number",
				},
				expect: blockchyp.TextPromptResponse{
					Success:  true,
					Response: notEmpty,
				},
			},
		},
	},
	{
		name:  "Interaction/RewardsNumberPrompt",
		group: testGroupInteractive,
		sim:   true,
		operations: []operation{
			{
				msg: "When prompted, enter any number and hit the green 'O' button.",
				args: []string{
					"-type", "text-prompt", "-terminal", terminalName, "-test",
					"-promptType", "rewards-number",
				},
				expect: blockchyp.TextPromptResponse{
					Success:  true,
					Response: notEmpty,
				},
			},
		},
	},
	{
		name:  "Interaction/FirstNamePrompt",
		group: testGroupInteractive,
		sim:   true,
		operations: []operation{
			{
				msg: "When prompted, enter any text and hit 'Next' or the green 'O' button.",
				args: []string{
					"-type", "text-prompt", "-terminal", terminalName, "-test",
					"-promptType", "first-name",
				},
				expect: blockchyp.TextPromptResponse{
					Success:  true,
					Response: notEmpty,
				},
			},
		},
	},
	{
		name:  "Interaction/LastNamePrompt",
		group: testGroupInteractive,
		sim:   true,
		operations: []operation{
			{
				msg: "When prompted, enter any text and hit 'Next' or the green 'O' button.",
				args: []string{
					"-type", "text-prompt", "-terminal", terminalName, "-test",
					"-promptType", "last-name",
				},
				expect: blockchyp.TextPromptResponse{
					Success:  true,
					Response: notEmpty,
				},
			},
		},
	},
	{
		name:  "Interaction/SignatureCapture",
		group: testGroupInteractive,
		sim:   true,
		operations: []operation{
			{
				msg: "When prompted, sign and hit 'Done'.",
				args: []string{
					"-type", "capture-signature", "-terminal", terminalName, "-test",
					"-sigFormat", blockchyp.SignatureFormatJPG,
					"-sigWidth", "50",
				},
				expect: blockchyp.CaptureSignatureResponse{
					Success: true,
					SigFile: notEmpty,
				},
			},
		},
	},
	{
		name:  "Interaction/LineItemDisplay",
		group: testGroupInteractive,
		sim:   true,
		operations: []operation{
			{
				args: []string{
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
				expect: blockchyp.Acknowledgement{
					Success: true,
				},
			},
			{
				validation: &validation{
					prompt: "Was a transaction with a discount displayed on the terminal?",
					expect: true,
				},
			},
			{
				args: []string{
					"-type", "clear", "-terminal", terminalName, "-test",
				},
				expect: blockchyp.Acknowledgement{
					Success: true,
				},
			},
			{
				validation: &validation{
					prompt: "Was the transaction cleared?",
					expect: true,
				},
			},
		},
	},
	{
		name:  "Interaction/WhitelistedBIN",
		group: testGroupMSR,
		sim:   true,
		operations: []operation{
			{
				msg: `From the admin console, set the whitelisted BIN range to the first few numbers of a test card.

When prompted for a card, use the test card you whitelisted.`,
				args: []string{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", "1.23",
				},
				expect: blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					ResponseDescription: "WHITELISTED",
				},
			},
		},
	},
}
