package main

import (
	"fmt"
	"log"

	blockchyp "github.com/blockchyp/blockchyp-go"
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
	request := blockchyp.TermsAndConditionsRequest{}
	request.Test = true
	request.TerminalName = "Test Terminal"
	request.TCAlias = "hippa"                        // Alias for a T&C template configured in blockchyp.
	request.TCName = "HIPPA Disclosure"              // Name of the contract or document if not using an alias.
	request.TCContent = "Full contract text"         // Full text of the contract or disclosure if not using an alias.
	request.SigFormat = blockchyp.SignatureFormatPNG // file format for the signature image.
	request.SigWidth = 200                           // width of the signature image in pixels.
	request.SigRequired = true                       // Whether or not a signature is required. Defaults to true.

	response, err := client.TermsAndConditions(request)

	if err != nil {
		log.Fatal(err)
	}

	//process the result
	if response.Success {
		fmt.Println("Success")
	}

	fmt.Println(response.Sig)
	fmt.Println(response.SigFile)
}
