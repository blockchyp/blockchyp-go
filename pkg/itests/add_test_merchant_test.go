//go:build integration
// +build integration

// Copyright 2019-2022 BlockChyp, Inc. All rights reserved. Use of this code
// is governed by a license that can be found in the LICENSE file.
//
// This file was generated automatically by the BlockChyp SDK Generator.
// Changes to this file will be lost every time the code is regenerated.

package itests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	blockchyp "github.com/blockchyp/blockchyp-go"
)

func TestAddTestMerchant(t *testing.T) {
	assert := assert.New(t)

	config := loadTestConfiguration(t)
	client := config.newTestClient(t, "partner")

	// setup request object
	request := blockchyp.AddTestMerchantRequest{
		DbaName:     "Test Merchant",
		CompanyName: "Test Merchant",
	}

	logObj(t, "Request:", request)

	response, err := client.AddTestMerchant(request)

	assert.NoError(err)

	logObj(t, "Response:", response)

	// response assertions
	assert.True(response.Success)
	assert.Equal("Test Merchant", response.DbaName)
	assert.Equal("Test Merchant", response.CompanyName)
	assert.True(response.Visa)
}