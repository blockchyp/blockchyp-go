// +build regression

package regression

import (
	"testing"

	"github.com/blockchyp/blockchyp-go"
)

func TestGiftCard(t *testing.T) {
	tests := map[string]struct {
		args   [][]string
		assert []interface{}
		txID   string
		txRef  string
	}{
		"Inactive": {
			args: [][]string{
				{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", "100.00",
				},
			},
			assert: []interface{}{
				blockchyp.AuthorizationResponse{
					Success:             false,
					Approved:            false,
					Test:                true,
					ResponseDescription: "Unknown Card",
					MaskedPAN:           notEmpty,
					PublicKey:           notEmpty,
					ReceiptSuggestions: blockchyp.ReceiptSuggestions{
						MerchantName:     notEmpty,
						MaskedPAN:        notEmpty,
						AuthorizedAmount: "0.00",
						TransactionType:  "charge",
						EntryMethod:      "SWIPE",
					},
				},
			},
		},
		"Lifecycle": {
			txRef: randomStr(),
			args: [][]string{
				{
					"-type", "gift-activate", "-terminal", terminalName, "-test",
					"-amount", "100.00",
				},
				{
					"-type", "gift-activate", "-terminal", terminalName, "-test",
					"-amount", "50.00",
				},
				{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", "100.00",
				},
				{
					"-type", "refund", "-test",
					"-tx",
				},
				{
					"-type", "refund", "-test",
					"-tx",
				},
				{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", "500.00",
				},
				{
					"-type", "void", "-test",
					"-tx",
				},
				{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", "10.00",
					"-txRef",
				},
				{
					"-type", "reverse", "-test",
					"-txRef",
				},
				{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", "500.00",
				},
				{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", "500.00",
				},
			},
			assert: []interface{}{
				blockchyp.GiftActivateResponse{
					Success:        true,
					Approved:       true,
					Test:           true,
					CurrentBalance: "100.00",
					MaskedPAN:      notEmpty,
					PublicKey:      notEmpty,
					CurrencyCode:   notEmpty,
					TickBlock:      notEmpty,
				},
				blockchyp.GiftActivateResponse{
					Success:        true,
					Approved:       true,
					Test:           true,
					CurrentBalance: "150.00",
					MaskedPAN:      notEmpty,
					PublicKey:      notEmpty,
					CurrencyCode:   notEmpty,
					TickBlock:      notEmpty,
				},
				blockchyp.AuthorizationResponse{
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
				blockchyp.AuthorizationResponse{
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
				// Second refund for the same transaction fails
				blockchyp.AuthorizationResponse{
					Success:  true,
					Approved: false,
					Test:     true,
				},
				blockchyp.AuthorizationResponse{
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
				blockchyp.VoidResponse{
					Success:         true,
					Approved:        true,
					Test:            true,
					TickBlock:       notEmpty,
					TransactionType: "void",
					PaymentType:     "BC_GIFT",
					EntryMethod:     "SWIPE",
				},
				blockchyp.AuthorizationResponse{
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
				blockchyp.AuthorizationResponse{
					Success:         true,
					Approved:        true,
					Test:            true,
					TransactionType: "reverse",
					MaskedPAN:       notEmpty,
					PublicKey:       notEmpty,
					TickBlock:       notEmpty,
				},
				blockchyp.AuthorizationResponse{
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
				blockchyp.AuthorizationResponse{
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
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			cli := newCLI(t)

			setup(t, "Select an unused gift card and swipe it when prompted.", true)

			for i := range test.args {
				if test.txID != "" && test.args[i][len(test.args[i])-1] == "-tx" {
					test.args[i] = append(test.args[i], test.txID)
				}
				if test.txRef != "" && test.args[i][len(test.args[i])-1] == "-txRef" {
					test.args[i] = append(test.args[i], test.txRef)
				}

				if res, ok := cli.run(test.args[i], test.assert[i]).(*blockchyp.AuthorizationResponse); ok {
					test.txID = res.TransactionID
				}
			}
		})
	}
}
