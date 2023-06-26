//go:build integration
// +build integration

// Copyright 2019-2023 BlockChyp, Inc. All rights reserved. Use of this code
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

func TestMerchantPlatforms(t *testing.T) {
	assert := assert.New(t)

	config := loadTestConfiguration(t)
	client := config.newTestClient(t, "partner")

	// setup request object
	setupRequest := blockchyp.AddTestMerchantRequest{
		DBAName:     "Test Merchant",
		CompanyName: "Test Merchant",
	}

	logObj(t, "Request:", setupRequest)

	setupResponse, err := client.AddTestMerchant(setupRequest)

	assert.NoError(err)

	logObj(t, "Response:", setupResponse)

	// setup request object
	request := blockchyp.MerchantProfileRequest{
		MerchantID: setupResponse.MerchantID,
	}

	logObj(t, "Request:", request)

	response, err := client.MerchantPlatforms(request)

	assert.NoError(err)

	logObj(t, "Response:", response)

	// response assertions
	assert.True(response.Success)
}
