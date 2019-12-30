package main

import (
	"encoding/json"
	"fmt"
	"log"

	blockchyp "github.com/blockchyp/blockchyp-go"
)

func main() {

	creds := blockchyp.APICredentials{
		APIKey:      "ZDSMMZLGRPBPRTJUBTAFBYZ33Q",
		BearerToken: "ZLBW5NR4U5PKD5PNP3ZP3OZS5U",
		SigningKey:  "9c6a5e8e763df1c9256e3d72bd7f53dfbd07312938131c75b3bfd254da787947",
	}

	client := blockchyp.NewClient(creds)
	client.GatewayHost = "http://localhost:8000"
	client.HTTPS = false

	req := blockchyp.AuthorizationRequest{}
	req.TerminalName = "Test Terminal"
	req.Amount = "55.00"

	response, err := client.Charge(req)

	if err != nil {
		log.Fatal(err)
	}

	if response.Approved {
		fmt.Println("Approved")
		fmt.Println(response.AuthCode)
		fmt.Println(response.AuthorizedAmount)
		fmt.Println(response.ReceiptSuggestions.AID)
	} else {
		fmt.Println(response.ResponseDescription)
	}

	b, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(b))

}
