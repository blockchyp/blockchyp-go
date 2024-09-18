//go:build integration
// +build integration

// Copyright 2019-2024 BlockChyp, Inc. All rights reserved. Use of this code
// is governed by a license that can be found in the LICENSE file.
//
// This file was generated automatically by the BlockChyp SDK Generator.
// Changes to this file will be lost every time the code is regenerated.

package itests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	blockchyp "github.com/blockchyp/blockchyp-go/v2"
)

func TestSubmitApplication(t *testing.T) {
	assert := assert.New(t)

	config := loadTestConfiguration(t)
	client := config.newTestClient(t, "partner")

	// setup request object
	request := blockchyp.SubmitApplicationRequest{
		Test:                 true,
		InviteCode:           "asdf",
		DBAName:              "BlockChyp",
		CorporateName:        "BlockChyp Inc.",
		WebSite:              "https://www.blockchyp.com",
		TaxIDNumber:          "123456789",
		EntityType:           "CORPORATION",
		StateOfIncorporation: "UT",
		MerchantType:         "RETAIL",
		BusinessDescription:  "Payment processing solutions",
		YearsInBusiness:      "5",
		BusinessPhoneNumber:  "5555551234",
		PhysicalAddress: blockchyp.Address{
			Address1:        "355 S 520 W",
			City:            "Lindon",
			StateOrProvince: "UT",
			PostalCode:      "84042",
			CountryCode:     "US",
		},
		MailingAddress: blockchyp.Address{
			Address1:        "355 S 520 W",
			City:            "Lindon",
			StateOrProvince: "UT",
			PostalCode:      "84042",
			CountryCode:     "US",
		},
		ContactFirstName:         "John",
		ContactLastName:          "Doe",
		ContactPhoneNumber:       "5555555678",
		ContactEmail:             "john.doe@example.com",
		ContactTitle:             "CEO",
		ContactTaxIDNumber:       "987654321",
		ContactDob:               "1980-01-01",
		ContactDlNumber:          "D1234567",
		ContactDlStateOrProvince: "NY",
		ContactDlExpiration:      "2025-12-31",
		ContactHomeAddress: blockchyp.Address{
			Address1:        "355 S 520 W",
			City:            "Lindon",
			StateOrProvince: "UT",
			PostalCode:      "84042",
			CountryCode:     "US",
		},
		ContactRole: "OWNER",
		Owners: []*blockchyp.Owner{
			&blockchyp.Owner{
				FirstName:         "John",
				LastName:          "Doe",
				JobTitle:          "CEO",
				TaxIDNumber:       "876543210",
				PhoneNumber:       "5555559876",
				Dob:               "1981-02-02",
				Ownership:         "50",
				Email:             "john.doe@example.com",
				DlNumber:          "D7654321",
				DlStateOrProvince: "UT",
				DlExpiration:      "2024-12-31",
				Address: blockchyp.Address{
					Address1:        "355 S 520 W",
					City:            "Lindon",
					StateOrProvince: "UT",
					PostalCode:      "84042",
					CountryCode:     "US",
				},
			},
		},
		ManualAccount: blockchyp.ApplicationAccount{
			Name:              "Business Checking",
			Bank:              "Test Bank",
			AccountHolderName: "BlockChyp Inc.",
			RoutingNumber:     "124001545",
			AccountNumber:     "987654321",
		},
		AverageTransaction:    "100.00",
		HighTransaction:       "1000.00",
		AverageMonth:          "10000.00",
		HighMonth:             "20000.00",
		RefundPolicy:          "30_DAYS",
		RefundDays:            "30",
		TimeZone:              "America/Denver",
		BatchCloseTime:        "23:59",
		MultipleLocations:     "false",
		EBTRequested:          "false",
		Ecommerce:             "true",
		CardPresentPercentage: "70",
		PhoneOrderPercentage:  "10",
		EcomPercentage:        "20",
		SignerName:            "John Doe",
	}

	logObj(t, "Request:", request)

	response, err := client.SubmitApplication(request)

	assert.NoError(err)

	logObj(t, "Response:", response)

	// response assertions
	assert.True(response.Success)
}
