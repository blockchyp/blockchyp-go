package main

import (
	"fmt"
	"log"

	blockchyp "github.com/blockchyp/blockchyp-go"
)

func updateTransactionDisplayExample() {

	// sample credentials
	creds := blockchyp.APICredentials{
		APIKey:      "ZDSMMZLGRPBPRTJUBTAFBYZ33Q",
		BearerToken: "ZLBW5NR4U5PKD5PNP3ZP3OZS5U",
		SigningKey:  "9c6a5e8e763df1c9256e3d72bd7f53dfbd07312938131c75b3bfd254da787947",
	}

	// instantiate the client
	client := blockchyp.NewClient(creds)

	// setup request object
	request := blockchyp.TransactionDisplayRequest{}
	request.Test = true
	request.TerminalName = "Test Terminal"
	request.Transaction = &blockchyp.TransactionDisplayTransaction{
		Subtotal: "60.00",
		Tax:      "5.00",
		Total:    "65.00",
		Items: []*blockchyp.TransactionDisplayItem{
			&blockchyp.TransactionDisplayItem{
				Description: "Leki Trekking Poles",
				Price:       "35.00",
				Quantity:    2,
				Extended:    "70.00",
				Discounts: []*blockchyp.TransactionDisplayDiscount{
					&blockchyp.TransactionDisplayDiscount{
						Description: "memberDiscount",
						Amount:      "10.00",
					},
				},
			},
		},
	}

	response, err := client.UpdateTransactionDisplay(request)

	if err != nil {
		log.Fatal(err)
	}

	//process the result
	if response.Success {
		fmt.Println("Succeded")
	}

}
