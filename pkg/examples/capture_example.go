package main

import (
	"fmt"
	"log"

	blockchyp "github.com/blockchyp/blockchyp-go/v2"
)

func captureExample() {
	// sample credentials
	creds := blockchyp.APICredentials{
		APIKey:      "ZDSMMZLGRPBPRTJUBTAFBYZ33Q",
		BearerToken: "ZLBW5NR4U5PKD5PNP3ZP3OZS5U",
		SigningKey:  "9c6a5e8e763df1c9256e3d72bd7f53dfbd07312938131c75b3bfd254da787947",
	}

	// instantiate the client
	client := blockchyp.NewClient(creds)

	// setup request object
	request := blockchyp.CaptureRequest{
		Test:          true,
		TransactionID: "<ORIGINAL TRANSACTION ID>",
		Amount:        "32.00",
	}

	response, err := client.Capture(request)

	if err != nil {
		log.Fatal(err)
	}

	//process the result
	if response.Approved {
		fmt.Println("approved")
	}

	fmt.Printf("Response: %+v\n", response)
}
