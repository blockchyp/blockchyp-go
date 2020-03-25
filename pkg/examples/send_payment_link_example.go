package main

import (
	"fmt"
	"log"

	blockchyp "github.com/blockchyp/blockchyp-go"
)

func sendPaymentLinkExample() {
	// sample credentials
	creds := blockchyp.APICredentials{
		APIKey:      "ZDSMMZLGRPBPRTJUBTAFBYZ33Q",
		BearerToken: "ZLBW5NR4U5PKD5PNP3ZP3OZS5U",
		SigningKey:  "9c6a5e8e763df1c9256e3d72bd7f53dfbd07312938131c75b3bfd254da787947",
	}

	// instantiate the client
	client := blockchyp.NewClient(creds)

	// setup request object
	request := blockchyp.PaymentLinkRequest{
		Amount:      "199.99",
		Description: "Widget",
		Subject:     "Widget invoice",
		Transaction: &blockchyp.TransactionDisplayTransaction{
			Subtotal: "195.00",
			Tax:      "4.99",
			Total:    "199.99",
			Items: []*blockchyp.TransactionDisplayItem{
				&blockchyp.TransactionDisplayItem{
					Description: "Widget",
					Price:       "195.00",
					Quantity:    1,
				},
			},
		},
		AutoSend: true,
		Customer: blockchyp.Customer{
			CustomerRef:  "Customer reference string",
			FirstName:    "FirstName",
			LastName:     "LastName",
			CompanyName:  "Company Name",
			EmailAddress: "support@blockchyp.com",
			SmsNumber:    "(123) 123-1231",
		},
	}

	response, err := client.SendPaymentLink(request)

	if err != nil {
		log.Fatal(err)
	}

	//process the result
	if response.Success {
		fmt.Println("Success")
	}

	fmt.Printf("Response: %+v\n", response)
}
