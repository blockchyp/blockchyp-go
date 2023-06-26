package main

import (
	"fmt"
	"log"

	blockchyp "github.com/blockchyp/blockchyp-go/v2"
)

func updateSlideShowExample() {
	// sample credentials
	creds := blockchyp.APICredentials{
		APIKey:      "ZDSMMZLGRPBPRTJUBTAFBYZ33Q",
		BearerToken: "ZLBW5NR4U5PKD5PNP3ZP3OZS5U",
		SigningKey:  "9c6a5e8e763df1c9256e3d72bd7f53dfbd07312938131c75b3bfd254da787947",
	}

	// instantiate the client
	client := blockchyp.NewClient(creds)

	// setup request object
	request := blockchyp.SlideShow{
		Name:  "Test Slide Show",
		Delay: 5,
		Slides: []*blockchyp.Slide{
			&blockchyp.Slide{
				MediaID: "<MEDIA ID>",
			},
		},
	}

	response, err := client.UpdateSlideShow(request)

	if err != nil {
		log.Fatal(err)
	}

	//process the result
	if response.Success {
		fmt.Println("Success")
	}

	fmt.Printf("Response: %+v\n", response)
}
