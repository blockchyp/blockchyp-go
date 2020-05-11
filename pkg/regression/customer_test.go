// +build regression

package regression

import (
	"testing"

	"github.com/blockchyp/blockchyp-go"
)

func TestCustomer(t *testing.T) {
	if acquirerMode {
		t.Skip("skipped for acquirer test run")
	}

	initialState := blockchyp.Customer{
		FirstName:    randomStr(),
		LastName:     randomStr(),
		CompanyName:  randomStr(),
		EmailAddress: randomStr(),
		SmsNumber:    randomSMSNum(),
	}
	endState := initialState
	endState.EmailAddress = randomStr()

	tests := map[string]struct {
		instructions string
		args         [][]string
		assert       []interface{}
		customer     *blockchyp.Customer
	}{
		"Lifecycle": {
			args: [][]string{
				{
					"-type", "update-customer",
					"-firstName", initialState.FirstName,
					"-lastName", initialState.LastName,
					"-companyName", initialState.CompanyName,
					"-email", initialState.EmailAddress,
					"-sms", initialState.SmsNumber,
				},
				{
					"-type", "get-customer",
					"-customerId",
				},
				{
					"-type", "update-customer",
					"-email", endState.EmailAddress,
					"-customerId",
				},
				{
					"-type", "get-customer",
					"-customerId",
				},
				{
					"-type", "search-customer",
					"-query", endState.FirstName[:10],
				},
				{
					"-type", "search-customer",
					"-query", endState.LastName,
				},
				{
					"-type", "search-customer",
					"-query", endState.SmsNumber,
				},
				{
					"-type", "send-link",
					"-displaySubtotal", randomAmount(),
					"-displayTax", randomAmount(),
					"-displayTotal", randomAmount(),
					"-lineItemDescription", randomStr(),
					"-lineItemQty", "1",
					"-lineItemPrice", randomAmount(),
					"-lineItemDiscountDescription", randomStr(),
					"-lineItemDiscountAmount", randomAmount(),
					"-lineItemExtended", randomAmount(),
					"-desc", "Thank you for your order. Your order will be ready in 20 minutes",
					"-amount", randomAmount(),
					"-orderRef", "12345",
					"-txRef", "12334",
					"-customerId",
				},
			},
			assert: []interface{}{
				blockchyp.CustomerResponse{
					Success:  true,
					Customer: &initialState,
				},
				blockchyp.CustomerResponse{
					Success:  true,
					Customer: &initialState,
				},
				blockchyp.CustomerResponse{
					Success:  true,
					Customer: &endState,
				},
				blockchyp.CustomerResponse{
					Success:  true,
					Customer: &endState,
				},
				blockchyp.CustomerSearchResponse{
					Success: true,
					Customers: []blockchyp.Customer{
						endState,
					},
				},
				blockchyp.CustomerSearchResponse{
					Success: true,
					Customers: []blockchyp.Customer{
						endState,
					},
				},
				blockchyp.CustomerSearchResponse{
					Success: true,
					Customers: []blockchyp.Customer{
						endState,
					},
				},
				blockchyp.PaymentLinkResponse{
					Success: true,
					URL:     notEmpty,
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			cli := newCLI(t)

			setup(t, test.instructions, true)

			for i := range test.args {
				if test.customer != nil {
					if test.args[i][len(test.args[i])-1] == "-customerId" {
						test.args[i] = append(test.args[i], test.customer.ID)
					}
				}

				if res, ok := cli.run(test.args[i], test.assert[i]).(*blockchyp.CustomerResponse); ok && test.customer == nil {
					test.customer = res.Customer
				}
			}
		})
	}
}
