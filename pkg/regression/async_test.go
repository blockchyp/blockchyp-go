// +build regression

package regression

import (
	"testing"

	"github.com/blockchyp/blockchyp-go"
)

func TestAsync(t *testing.T) {
	tests := map[string]struct {
		args       []interface{}
		assert     []interface{}
		validation validation
		txRef      string
	}{
		"AsyncCharge": {
			args: []interface{}{
				"Insert an EMV test card when prompted",
				[]string{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", amount(0),
					"-async", "-txRef",
				},
				[]string{
					"-type", "tx-status",
					"-txRef",
				},
				"Complete the transaction before continuing",
				[]string{
					"-type", "tx-status",
					"-txRef",
				},
			},
			assert: []interface{}{
				nil,
				blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					ResponseDescription: "Accepted",
				},
				blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					ResponseDescription: "PENDING",
				},
				nil,
				blockchyp.AuthorizationResponse{
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
		},
		"AsyncPreauth": {
			args: []interface{}{
				"Insert an EMV test card when prompted",
				[]string{
					"-type", "preauth", "-terminal", terminalName, "-test",
					"-amount", amount(0),
					"-async", "-txRef",
				},
				[]string{
					"-type", "tx-status",
					"-txRef",
				},
				"Complete the transaction before continuing",
				[]string{
					"-type", "tx-status",
					"-txRef",
				},
			},
			assert: []interface{}{
				nil,
				blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					ResponseDescription: "Accepted",
				},
				blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					ResponseDescription: "PENDING",
				},
				nil,
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "preauth",
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
						TransactionType:  "preauth",
						EntryMethod:      "CHIP",
					},
				},
			},
		},
		"QueueCharge": {
			args: []interface{}{
				[]string{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", amount(0),
					"-queue", "-desc", "Test Ticket", "-txRef",
				},
				[]string{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", amount(1),
					"-queue", "-desc", "Dummy #1", "-txRef", "placeholder",
				},
				[]string{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", amount(2),
					"-queue", "-desc", "Dummy #2", "-txRef", "placeholder",
				},
				[]string{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", amount(3),
					"-queue", "-desc", "Dummy #3", "-txRef", "placeholder",
				},
				[]string{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", amount(4),
					"-queue", "-desc", "Dummy #4", "-txRef", "placeholder",
				},
				[]string{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", amount(5),
					"-queue", "-desc", "Dummy #5", "-txRef", "placeholder",
				},
				[]string{
					"-type", "tx-status",
					"-txRef",
				},
				"Select the ticket labeled 'Test Ticket` and insert an EMV test card when prompted.",
				[]string{
					"-type", "tx-status",
					"-txRef",
				},
			},
			assert: []interface{}{
				blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					ResponseDescription: "Queued",
				},
				blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					ResponseDescription: "Queued",
				},
				blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					ResponseDescription: "Queued",
				},
				blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					ResponseDescription: "Queued",
				},
				blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					ResponseDescription: "Queued",
				},
				blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					ResponseDescription: "Queued",
				},

				blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					ResponseDescription: "PENDING",
				},
				nil,
				blockchyp.AuthorizationResponse{
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
		},
		"QueuePreauth": {
			args: []interface{}{
				[]string{
					"-type", "preauth", "-terminal", terminalName, "-test",
					"-amount", amount(0),
					"-queue", "-desc", "Test Ticket", "-txRef",
				},
				[]string{
					"-type", "preauth", "-terminal", terminalName, "-test",
					"-amount", amount(1),
					"-queue", "-desc", "Dummy #1", "-txRef", "placeholder",
				},
				[]string{
					"-type", "preauth", "-terminal", terminalName, "-test",
					"-amount", amount(2),
					"-queue", "-desc", "Dummy #2", "-txRef", "placeholder",
				},
				[]string{
					"-type", "preauth", "-terminal", terminalName, "-test",
					"-amount", amount(3),
					"-queue", "-desc", "Dummy #3", "-txRef", "placeholder",
				},
				[]string{
					"-type", "preauth", "-terminal", terminalName, "-test",
					"-amount", amount(4),
					"-queue", "-desc", "Dummy #4", "-txRef", "placeholder",
				},
				[]string{
					"-type", "preauth", "-terminal", terminalName, "-test",
					"-amount", amount(5),
					"-queue", "-desc", "Dummy #5", "-txRef", "placeholder",
				},
				[]string{
					"-type", "tx-status",
					"-txRef",
				},
				"Select the ticket labeled 'Test Ticket` and insert an EMV test card when prompted.",
				[]string{
					"-type", "tx-status",
					"-txRef",
				},
			},
			assert: []interface{}{
				blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					ResponseDescription: "Queued",
				},
				blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					ResponseDescription: "Queued",
				},
				blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					ResponseDescription: "Queued",
				},
				blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					ResponseDescription: "Queued",
				},
				blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					ResponseDescription: "Queued",
				},
				blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					ResponseDescription: "Queued",
				},

				blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					ResponseDescription: "PENDING",
				},
				nil,
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "preauth",
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
						TransactionType:  "preauth",
						EntryMethod:      "CHIP",
					},
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			cli := newCLI(t)

			if test.txRef == "" {
				test.txRef = randomStr()
			}

			for i := range test.args {
				args, ok := test.args[i].([]string)
				if !ok {
					continue
				}
				if args[len(args)-1] == "-txRef" {
					test.args[i] = append(args, test.txRef)
				}
			}

			for i := range test.args {
				switch v := test.args[i].(type) {
				case string:
					setup(t, v, true)
				case func(*testing.T):
					v(t)
				case []string:
					cli.run(v, test.assert[i])
				}
			}

			validate(t, test.validation)
		})
	}
}
