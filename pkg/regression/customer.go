package regression

import (
	"github.com/blockchyp/blockchyp-go"
)

var customerTests = testCases{
	{
		sim:   true,
		name:  "Customer/Lifecycle",
		group: testGroupNonInteractive,
		operations: []operation{
			{
				args: []string{
					"-type", "update-customer",
					"-firstName", customerN(0).FirstName,
					"-lastName", customerN(0).LastName,
					"-companyName", customerN(0).CompanyName,
					"-email", customerN(0).EmailAddress,
					"-sms", customerN(0).SmsNumber,
				},
				expect: blockchyp.CustomerResponse{
					Success:  true,
					Customer: customerN(0),
				},
			},
			{
				args: []string{
					"-type", "get-customer",
					"-customerId", customerIDN(0),
				},
				expect: blockchyp.CustomerResponse{
					Success:  true,
					Customer: customerN(0),
				},
			},
			{
				args: []string{
					"-type", "update-customer",
					"-email", customerEndStateN(0).EmailAddress,
					"-customerId", customerIDN(0),
				},
				expect: blockchyp.CustomerResponse{
					Success:  true,
					Customer: customerEndStateN(0),
				},
			},
			{
				args: []string{
					"-type", "get-customer",
					"-customerId", customerIDN(0),
				},
				expect: blockchyp.CustomerResponse{
					Success:  true,
					Customer: customerEndStateN(0),
				},
			},
			{
				args: []string{
					"-type", "search-customer",
					"-query", customerEndStateN(0).FirstName[:10],
				},
				expect: blockchyp.CustomerSearchResponse{
					Success: true,
					Customers: []blockchyp.Customer{
						*customerEndStateN(0),
					},
				},
			},
			{
				args: []string{
					"-type", "search-customer",
					"-query", customerEndStateN(0).LastName,
				},
				expect: blockchyp.CustomerSearchResponse{
					Success: true,
					Customers: []blockchyp.Customer{
						*customerEndStateN(0),
					},
				},
			},
			{
				args: []string{
					"-type", "search-customer",
					"-query", customerEndStateN(0).SmsNumber,
				},
				expect: blockchyp.CustomerSearchResponse{
					Success: true,
					Customers: []blockchyp.Customer{
						*customerEndStateN(0),
					},
				},
			},
			{
				args: []string{
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
					"-customerId", customerIDN(0),
				},
				expect: blockchyp.PaymentLinkResponse{
					Success: true,
					URL:     notEmpty,
				},
			},
		},
	},
}

var initialState = []*blockchyp.Customer{}
var endState = []*blockchyp.Customer{}

func customerN(n int) *blockchyp.Customer {
	if n >= len(initialState) {
		cust := blockchyp.Customer{
			FirstName:    randomStr(),
			LastName:     randomStr(),
			CompanyName:  randomStr(),
			EmailAddress: randomStr(),
			SmsNumber:    randomSMSNum(),
		}
		end := cust
		end.EmailAddress = randomStr()

		initialState = append(initialState, &cust)
		endState = append(endState, &end)
	}

	return initialState[n]
}

func customerEndStateN(n int) *blockchyp.Customer {
	return endState[n]
}
