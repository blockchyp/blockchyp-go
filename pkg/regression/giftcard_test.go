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
					"-type", "charge", "-terminal", "Test Terminal", "-test",
					"-amount", "100.00",
				},
			},
			assert: []interface{}{
				blockchyp.AuthorizationResponse{
					Success:             false,
					Approved:            false,
					Test:                true,
					ResponseDescription: "Unknown Card",
				},
			},
		},
		"Lifecycle": {
			txRef: randomStr(),
			args: [][]string{
				{
					"-type", "gift-activate", "-terminal", "Test Terminal", "-test",
					"-amount", "100.00",
				},
				{
					"-type", "gift-activate", "-terminal", "Test Terminal", "-test",
					"-amount", "50.00",
				},
				{
					"-type", "charge", "-terminal", "Test Terminal", "-test",
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
					"-type", "charge", "-terminal", "Test Terminal", "-test",
					"-amount", "500.00",
				},
				{
					"-type", "void", "-test",
					"-tx",
				},
				{
					"-type", "charge", "-terminal", "Test Terminal", "-test",
					"-amount", "10.00",
					"-txRef",
				},
				{
					"-type", "reverse", "-test",
					"-txRef",
				},
				{
					"-type", "charge", "-terminal", "Test Terminal", "-test",
					"-amount", "500.00",
				},
				{
					"-type", "charge", "-terminal", "Test Terminal", "-test",
					"-amount", "500.00",
				},
			},
			assert: []interface{}{
				blockchyp.GiftActivateResponse{
					Success:        true,
					Approved:       true,
					Test:           true,
					CurrentBalance: "100.00",
				},
				blockchyp.GiftActivateResponse{
					Success:        true,
					Approved:       true,
					Test:           true,
					CurrentBalance: "150.00",
				},
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					PaymentType:      "BC_GIFT",
					RequestedAmount:  "100.00",
					AuthorizedAmount: "100.00",
					RemainingBalance: "50.00",
				},
				blockchyp.AuthorizationResponse{
					Success:     true,
					Approved:    true,
					Test:        true,
					PaymentType: "BC_GIFT",
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
				},
				blockchyp.VoidResponse{
					Success:  true,
					Approved: true,
					Test:     true,
				},
				blockchyp.AuthorizationResponse{
					Success:     true,
					Approved:    true,
					Test:        true,
					PaymentType: "BC_GIFT",
				},
				blockchyp.AuthorizationResponse{
					Success:         true,
					Approved:        true,
					Test:            true,
					TransactionType: "reverse",
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
				},
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         false,
					Test:             true,
					PaymentType:      "BC_GIFT",
					RequestedAmount:  "500.00",
					AuthorizedAmount: "0.00",
					RemainingBalance: "0.00",
				},
			},
		},
	}

	cli := newCLI(t)

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
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
