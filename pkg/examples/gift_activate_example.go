package main

import (
	"fmt"
	"log"

	blockchyp "github.com/blockchyp/blockchyp-go"
)

func giftActivateExample() {

	// sample credentials
	creds := blockchyp.APICredentials{
		APIKey:      "ZDSMMZLGRPBPRTJUBTAFBYZ33Q",
		BearerToken: "ZLBW5NR4U5PKD5PNP3ZP3OZS5U",
		SigningKey:  "9c6a5e8e763df1c9256e3d72bd7f53dfbd07312938131c75b3bfd254da787947",
	}

	// instantiate the client
	client := blockchyp.NewClient(creds)

	// setup request object
	request := blockchyp.GiftActivateRequest{}
	request.Test = true
	request.TerminalName = "Test Terminal"
	request.Amount = "50.00"

	response, err := client.GiftActivate(request)

	if err != nil {
		log.Fatal(err)
	}

	//process the result
	if response.Approved {
		fmt.Println("Approved")
	}

	fmt.Println(response.Amount)
	fmt.Println(response.CurrentBalance)
	fmt.Println(response.PublicKey)
}