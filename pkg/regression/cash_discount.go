package regression

import (
	"github.com/blockchyp/blockchyp-go"
)

var cashDiscountTests = testCases{
	{
		name:  "CashDiscount/CreditSurchargeCashDiscount",
		group: testGroupNoCVM,
		operations: []operation{
			{
				msg: "Insert a test EMV " + format(bold, yellow) + "credit" + format(normal, magenta) + " card when prompted.",
				args: []string{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", amount(0),
					"-surcharge", "-cashDiscount",
				},
				expect: blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "charge",
					EntryMethod:      "CHIP",
					RequestedAmount:  addFees(amount(0)),
					AuthorizedAmount: addFees(amount(0)),
					MaskedPAN:        notEmpty,
					ReceiptSuggestions: blockchyp.ReceiptSuggestions{
						MerchantName:    notEmpty,
						MaskedPAN:       notEmpty,
						TransactionType: "charge",
						EntryMethod:     "CHIP",
						// Surcharge is applied and no discount is applied
						AuthorizedAmount: addFees(amount(0)),
						Surcharge:        fees(amount(0)),
						CashDiscount:     "0.00",
					},
				},
			},
		},
	},
	{
		name:  "CashDiscount/CreditSurchargeNoCashDiscount",
		group: testGroupNoCVM,
		operations: []operation{
			{
				msg: "Insert a test EMV " + format(bold, yellow) + "credit" + format(normal, magenta) + " card when prompted.",
				args: []string{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", amount(0),
					"-surcharge",
				},
				expect: blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "charge",
					EntryMethod:      "CHIP",
					RequestedAmount:  addFees(amount(0)),
					AuthorizedAmount: addFees(amount(0)),
					MaskedPAN:        notEmpty,
					ReceiptSuggestions: blockchyp.ReceiptSuggestions{
						MerchantName:    notEmpty,
						MaskedPAN:       notEmpty,
						TransactionType: "charge",
						EntryMethod:     "CHIP",
						// Surcharge is applied and no discount is applied
						AuthorizedAmount: addFees(amount(0)),
						Surcharge:        fees(amount(0)),
						CashDiscount:     "0.00",
					},
				},
			},
		},
	},
	{
		name:  "CashDiscount/CreditNoSurchargeCashDiscount",
		group: testGroupNoCVM,
		operations: []operation{
			{
				msg: "Insert a test EMV " + format(bold, yellow) + "credit" + format(normal, magenta) + " card when prompted.",
				args: []string{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", amount(0),
					"-cashDiscount",
				},
				expect: blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "charge",
					EntryMethod:      "CHIP",
					RequestedAmount:  amount(0),
					AuthorizedAmount: amount(0),
					MaskedPAN:        notEmpty,
					ReceiptSuggestions: blockchyp.ReceiptSuggestions{
						MerchantName:    notEmpty,
						MaskedPAN:       notEmpty,
						TransactionType: "charge",
						EntryMethod:     "CHIP",
						// Surcharge is not applied
						AuthorizedAmount: amount(0),
						Surcharge:        "0.00",
						CashDiscount:     "0.00",
					},
				},
			},
		},
	},
	{
		name:  "CashDiscount/DebitSurchargeCashDiscount",
		group: testGroupInteractive,
		operations: []operation{
			{
				msg: "Insert a test EMV " + format(bold, yellow) + "debit" + format(normal, magenta) + " card when prompted.",
				args: []string{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", amount(0),
					"-surcharge", "-cashDiscount",
				},
				expect: blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "charge",
					EntryMethod:      "CHIP",
					RequestedAmount:  amount(0),
					AuthorizedAmount: amount(0),
					MaskedPAN:        notEmpty,
					ReceiptSuggestions: blockchyp.ReceiptSuggestions{
						MerchantName:    notEmpty,
						MaskedPAN:       notEmpty,
						TransactionType: "charge",
						EntryMethod:     "CHIP",
						// Surcharge and discount offset
						AuthorizedAmount: amount(0),
						Surcharge:        fees(amount(0)),
						CashDiscount:     fees(amount(0)),
						PINVerified:      true,
					},
				},
			},
		},
	},
	{
		name:  "CashDiscount/DebitSurchargeNoCashDiscount",
		group: testGroupInteractive,
		operations: []operation{
			{
				msg: "Insert a test EMV " + format(bold, yellow) + "debit" + format(normal, magenta) + " card when prompted.",
				args: []string{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", amount(0),
					"-surcharge",
				},
				expect: blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "charge",
					EntryMethod:      "CHIP",
					RequestedAmount:  amount(0),
					AuthorizedAmount: amount(0),
					MaskedPAN:        notEmpty,
					ReceiptSuggestions: blockchyp.ReceiptSuggestions{
						MerchantName:    notEmpty,
						MaskedPAN:       notEmpty,
						TransactionType: "charge",
						EntryMethod:     "CHIP",
						// Surcharge is not applied
						AuthorizedAmount: amount(0),
						Surcharge:        "0.00",
						CashDiscount:     "0.00",
						PINVerified:      true,
					},
				},
			},
		},
	},
	{
		name:  "CashDiscount/DebitNoSurchargeCashDiscount",
		group: testGroupInteractive,
		operations: []operation{
			{
				msg: "Insert a test EMV " + format(bold, yellow) + "debit" + format(normal, magenta) + " card when prompted.",
				args: []string{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", amount(0),
					"-cashDiscount",
				},
				expect: blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "charge",
					EntryMethod:      "CHIP",
					RequestedAmount:  cashDiscount(amount(0)),
					AuthorizedAmount: cashDiscount(amount(0)),
					MaskedPAN:        notEmpty,
					ReceiptSuggestions: blockchyp.ReceiptSuggestions{
						MerchantName:    notEmpty,
						MaskedPAN:       notEmpty,
						TransactionType: "charge",
						EntryMethod:     "CHIP",
						// Cash discount is applied, fees are not applied
						AuthorizedAmount: cashDiscount(amount(0)),
						Surcharge:        "0.00",
						CashDiscount:     fees(amount(0)),
						PINVerified:      true,
					},
				},
			},
		},
	},
	{
		name:  "CashDiscount/APISurchargeCashDiscount",
		group: testGroupNonInteractive,
		operations: []operation{
			{
				args: []string{
					"-type", "cash-discount", "-test",
					"-amount", amount(0),
					"-surcharge", "-cashDiscount",
				},
				expect: blockchyp.CashDiscountResponse{
					Success:      true,
					Amount:       amount(0),
					Surcharge:    fees(amount(0)),
					CashDiscount: fees(amount(0)),
				},
			},
		},
	},
	{
		name:  "CashDiscount/APISurchargeNoCashDiscount",
		group: testGroupNonInteractive,
		operations: []operation{
			{
				args: []string{
					"-type", "cash-discount", "-test",
					"-amount", amount(0),
					"-surcharge",
				},
				expect: blockchyp.CashDiscountResponse{
					Success: true,
					Amount:  amount(0),
				},
			},
		},
	},
	{
		name:  "CashDiscount/APINoSurchargeCashDiscount",
		group: testGroupNonInteractive,
		operations: []operation{
			{
				args: []string{
					"-type", "cash-discount", "-test",
					"-amount", amount(0),
					"-cashDiscount",
				},
				expect: blockchyp.CashDiscountResponse{
					Success:      true,
					Amount:       cashDiscount(amount(0)),
					CashDiscount: fees(amount(0)),
				},
			},
		},
	},
}
