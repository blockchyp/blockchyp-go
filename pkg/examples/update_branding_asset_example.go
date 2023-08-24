package main

import (
	"fmt"
	"log"

	blockchyp "github.com/blockchyp/blockchyp-go/v2"
)

func updateBrandingAssetExample() {
	// sample credentials
	creds := blockchyp.APICredentials{
		APIKey:      "ZDSMMZLGRPBPRTJUBTAFBYZ33Q",
		BearerToken: "ZLBW5NR4U5PKD5PNP3ZP3OZS5U",
		SigningKey:  "9c6a5e8e763df1c9256e3d72bd7f53dfbd07312938131c75b3bfd254da787947",
	}

	// instantiate the client
	client := blockchyp.NewClient(creds)

	// setup request object
	request := blockchyp.BrandingAsset{
		MediaID:   "<MEDIA ID>",
		Padded:    true,
		Ordinal:   10,
		StartDate: "01/06/2021",
		StartTime: "14:00",
		EndDate:   "11/05/2024",
		EndTime:   "16:00",
		Notes:     "Test Branding Asset",
		Preview:   false,
		Enabled:   true,
	}

	response, err := client.UpdateBrandingAsset(request)

	if err != nil {
		log.Fatal(err)
	}

	//process the result
	if response.Success {
		fmt.Println("Success")
	}

	fmt.Printf("Response: %+v\n", response)
}
