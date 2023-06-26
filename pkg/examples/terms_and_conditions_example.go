package main

import (
	"fmt"
	"log"

	blockchyp "github.com/blockchyp/blockchyp-go/v2"
)

func termsAndConditionsExample() {
	// sample credentials
	creds := blockchyp.APICredentials{
		APIKey:      "ZDSMMZLGRPBPRTJUBTAFBYZ33Q",
		BearerToken: "ZLBW5NR4U5PKD5PNP3ZP3OZS5U",
		SigningKey:  "9c6a5e8e763df1c9256e3d72bd7f53dfbd07312938131c75b3bfd254da787947",
	}

	// instantiate the client
	client := blockchyp.NewClient(creds)

	// setup request object
	request := blockchyp.TermsAndConditionsRequest{
		Test:         true,
		TerminalName: "Test Terminal",

		// Alias for a Terms and Conditions template configured in the BlockChyp dashboard.
		TCAlias: "hippa",

		// Name of the contract or document if not using an alias.
		TCName: "HIPPA Disclosure",

		// Full text of the contract or disclosure if not using an alias.
		TCContent: "Full contract text",

		// file format for the signature image.
		SigFormat: blockchyp.SignatureFormatPNG,

		// width of the signature image in pixels.
		SigWidth: 200,

		// Whether or not a signature is required. Defaults to true.
		SigRequired: true,
	}

	response, err := client.TermsAndConditions(request)

	if err != nil {
		log.Fatal(err)
	}

	//process the result
	if response.Success {
		fmt.Println("Success")
	}

	fmt.Printf("Response: %+v\n", response)
}
