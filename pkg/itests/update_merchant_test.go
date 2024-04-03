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

func TestUpdateMerchant(t *testing.T) {
	assert := assert.New(t)

	config := loadTestConfiguration(t)
	client := config.newTestClient(t, "partner")

	// setup request object
	request := blockchyp.MerchantProfile{
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

	logObj(t, "Request:", request)

	response, err := client.UpdateMerchant(request)

	assert.NoError(err)

	logObj(t, "Response:", response)

	// response assertions
	assert.True(response.Success)
}
