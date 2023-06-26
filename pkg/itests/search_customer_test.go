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

func TestSearchCustomer(t *testing.T) {
	assert := assert.New(t)

	config := loadTestConfiguration(t)
	client := config.newTestClient(t, "")

	// setup request object
	setupRequest := blockchyp.UpdateCustomerRequest{
		Customer: blockchyp.Customer{
			FirstName:    "Test",
			LastName:     "Customer",
			CompanyName:  "Test Company",
			EmailAddress: "support@blockchyp.com",
			SmsNumber:    "(123) 123-1234",
		},
	}

	logObj(t, "Request:", setupRequest)

	setupResponse, err := client.UpdateCustomer(setupRequest)

	assert.NoError(err)

	logObj(t, "Response:", setupResponse)

	// setup request object
	request := blockchyp.CustomerSearchRequest{
		Query: "123123",
	}

	logObj(t, "Request:", request)

	response, err := client.CustomerSearch(request)

	assert.NoError(err)

	logObj(t, "Response:", response)

	// response assertions
	assert.True(response.Success)
}
