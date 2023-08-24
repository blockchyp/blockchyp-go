package main

import (
	"fmt"
	"log"

	blockchyp "github.com/blockchyp/blockchyp-go/v2"
)

func updateMerchantExample() {
	// sample credentials
	creds := blockchyp.APICredentials{
		APIKey:      "ZDSMMZLGRPBPRTJUBTAFBYZ33Q",
		BearerToken: "ZLBW5NR4U5PKD5PNP3ZP3OZS5U",
		SigningKey:  "9c6a5e8e763df1c9256e3d72bd7f53dfbd07312938131c75b3bfd254da787947",
	}

	// instantiate the client
	client := blockchyp.NewClient(creds)

	// setup request object
	request := blockchyp.MerchantProfile{
		MerchantID:  "<MERCHANT ID>",
		Test:        true,
		DBAName:     "Test Merchant",
		CompanyName: "Test Merchant",
		BillingAddress: blockchyp.Address{
			Address1:        "1060 West Addison",
			City:            "Chicago",
			StateOrProvince: "IL",
			PostalCode:      "60613",
		},
	}

	response, err := client.UpdateMerchant(request)

	if err != nil {
		log.Fatal(err)
	}

	//process the result
	if response.Success {
		fmt.Println("Success")
	}

	fmt.Printf("Response: %+v\n", response)
}
