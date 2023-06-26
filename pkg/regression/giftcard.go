package regression

import (
	"github.com/blockchyp/blockchyp-go/v2"
)

var giftCardTests = testCases{
	{
		name:  "Giftcard/Inactive",
		group: testGroupMSR,
		sim:   true,
		operations: []operation{
			{
				msg: "Select an unused gift card and swipe it when prompted.",
				args: []string{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", "100.00",
				},
				expect: blockchyp.AuthorizationResponse{
					Success:             false,
					Approved:            false,
					Test:                true,
					ResponseDescription: "Unknown Card",
					MaskedPAN:           notEmpty,
					PublicKey:           notEmpty,
					ReceiptSuggestions: blockchyp.ReceiptSuggestions{
						MaskedPAN:        notEmpty,
						AuthorizedAmount: "0.00",
						TransactionType:  "charge",
						EntryMethod:      "SWIPE",
					},
				},
			},
		},
	},
	{
		name:  "Giftcard/Lifecycle",
		group: testGroupMSR,
		sim:   true,
		operations: []operation{
			{
				msg: "Select an unused gift card and swipe it when prompted.",
				args: []string{
					"-type", "gift-activate", "-terminal", terminalName, "-test",
					"-amount", "100.00",
				},
				expect: blockchyp.GiftActivateResponse{
					Success:        true,
					Approved:       true,
					Test:           true,
					CurrentBalance: "100.00",
					MaskedPAN:      notEmpty,
					PublicKey:      notEmpty,
					CurrencyCode:   notEmpty,
					TickBlock:      notEmpty,
				},
			},
			{
				args: []string{
					"-type", "gift-activate", "-terminal", terminalName, "-test",
					"-amount", "50.00",
				},
				expect: blockchyp.GiftActivateResponse{
					Success:        true,
					Approved:       true,
					Test:           true,
					CurrentBalance: "150.00",
					MaskedPAN:      notEmpty,
					PublicKey:      notEmpty,
					CurrencyCode:   notEmpty,
					TickBlock:      notEmpty,
				},
			},
			{
				args: []string{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", "100.00",
				},
				expect: blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					PaymentType:      "BC_GIFT",
					RequestedAmount:  "100.00",
					AuthorizedAmount: "100.00",
					RemainingBalance: "50.00",
					MaskedPAN:        notEmpty,
					PublicKey:        notEmpty,
					CurrencyCode:     notEmpty,
					TickBlock:        notEmpty,
					ReceiptSuggestions: blockchyp.ReceiptSuggestions{
						MerchantName:     notEmpty,
						MaskedPAN:        notEmpty,
						AuthorizedAmount: "100.00",
						TransactionType:  "charge",
						EntryMethod:      "SWIPE",
					},
				},
			},
			{
				args: []string{
					"-type", "refund", "-test",
					"-tx", txIDN(-1),
				},
				expect: blockchyp.AuthorizationResponse{
					Success:      true,
					Approved:     true,
					Test:         true,
					PaymentType:  "BC_GIFT",
					MaskedPAN:    notEmpty,
					PublicKey:    notEmpty,
					CurrencyCode: notEmpty,
					TickBlock:    notEmpty,
					ReceiptSuggestions: blockchyp.ReceiptSuggestions{
						MerchantName:    notEmpty,
						MaskedPAN:       notEmpty,
						TransactionType: "refund",
						EntryMethod:     "SWIPE",
					},
				},
			},
			{
				args: []string{
					"-type", "refund", "-test",
					"-tx", txIDN(-2),
				},
				// Second refund for the same transaction fails
				expect: blockchyp.AuthorizationResponse{
					Success:  true,
					Approved: false,
					Test:     true,
				},
			},
			{
				args: []string{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", "500.00",
				},
				expect: blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					PaymentType:      "BC_GIFT",
					RequestedAmount:  "500.00",
					AuthorizedAmount: "150.00",
					RemainingBalance: "0.00",
					PartialAuth:      true,
					MaskedPAN:        notEmpty,
					PublicKey:        notEmpty,
					CurrencyCode:     notEmpty,
					TickBlock:        notEmpty,
					ReceiptSuggestions: blockchyp.ReceiptSuggestions{
						MerchantName:    notEmpty,
						MaskedPAN:       notEmpty,
						TransactionType: "charge",
						EntryMethod:     "SWIPE",
					},
				},
			},
			{
				args: []string{
					"-type", "void", "-test",
					"-tx", txIDN(-1),
				},
				expect: blockchyp.VoidResponse{
					Success:         true,
					Approved:        true,
					Test:            true,
					TickBlock:       notEmpty,
					TransactionType: "void",
					PaymentType:     "BC_GIFT",
					EntryMethod:     "SWIPE",
				},
			},
			{
				args: []string{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", "10.00",
					"-txRef", newTxRef,
				},
				expect: blockchyp.AuthorizationResponse{
					Success:      true,
					Approved:     true,
					Test:         true,
					PaymentType:  "BC_GIFT",
					MaskedPAN:    notEmpty,
					PublicKey:    notEmpty,
					CurrencyCode: notEmpty,
					TickBlock:    notEmpty,
					ReceiptSuggestions: blockchyp.ReceiptSuggestions{
						MerchantName:    notEmpty,
						MaskedPAN:       notEmpty,
						TransactionType: "charge",
						EntryMethod:     "SWIPE",
					},
				},
			},
			{
				args: []string{
					"-type", "reverse", "-test",
					"-txRef", txRefN(-1),
				},
				expect: blockchyp.AuthorizationResponse{
					Success:         true,
					Approved:        true,
					Test:            true,
					TransactionType: "reverse",
					MaskedPAN:       notEmpty,
					PublicKey:       notEmpty,
					TickBlock:       notEmpty,
				},
			},
			{
				args: []string{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", "500.00",
				},
				expect: blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					PaymentType:      "BC_GIFT",
					RequestedAmount:  "500.00",
					AuthorizedAmount: "150.00",
					RemainingBalance: "0.00",
					PartialAuth:      true,
					MaskedPAN:        notEmpty,
					PublicKey:        notEmpty,
					CurrencyCode:     notEmpty,
					TickBlock:        notEmpty,
					ReceiptSuggestions: blockchyp.ReceiptSuggestions{
						MerchantName:    notEmpty,
						MaskedPAN:       notEmpty,
						TransactionType: "charge",
						EntryMethod:     "SWIPE",
					},
				},
			},
			{
				args: []string{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", "500.00",
				},
				expect: blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         false,
					Test:             true,
					PaymentType:      "BC_GIFT",
					RequestedAmount:  "500.00",
					AuthorizedAmount: "0.00",
					RemainingBalance: "0.00",
					MaskedPAN:        notEmpty,
					PublicKey:        notEmpty,
					CurrencyCode:     notEmpty,
					TickBlock:        notEmpty,
					ReceiptSuggestions: blockchyp.ReceiptSuggestions{
						MerchantName:    notEmpty,
						MaskedPAN:       notEmpty,
						TransactionType: "charge",
						EntryMethod:     "SWIPE",
					},
				},
			},
		},
	},
}
