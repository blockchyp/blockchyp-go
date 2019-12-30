package main

import (
	"fmt"
	"log"

	blockchyp "github.com/blockchyp/blockchyp-go"
)

func booleanPromptExample() {

	// sample credentials
	creds := blockchyp.APICredentials{
		APIKey:      "ZDSMMZLGRPBPRTJUBTAFBYZ33Q",
		BearerToken: "ZLBW5NR4U5PKD5PNP3ZP3OZS5U",
		SigningKey:  "9c6a5e8e763df1c9256e3d72bd7f53dfbd07312938131c75b3bfd254da787947",
	}

	// instantiate the client
	client := blockchyp.NewClient(creds)

	// setup request object
	request := blockchyp.BooleanPromptRequest{}
	request.Test = true
	request.TerminalName = "Test Terminal"
	request.Prompt = "Would you like to become a member?"
	request.YesCaption = "Yes"
	request.NoCaption = "No"

	response, err := client.BooleanPrompt(request)

	if err != nil {
		log.Fatal(err)
	}

	//process the result
	if response.Success {
		fmt.Println("Success")
	}

	fmt.Println(response.Response)
}
