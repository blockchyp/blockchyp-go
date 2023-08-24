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

func TestLinkToken(t *testing.T) {
	assert := assert.New(t)

	config := loadTestConfiguration(t)
	client := config.newTestClient(t, "")

	// setup request object
	setupRequest := blockchyp.EnrollRequest{
		PAN:  "4111111111111111",
		Test: true,
		Customer: &blockchyp.Customer{
			CustomerRef: "TESTCUSTOMER",
			FirstName:   "Test",
			LastName:    "Customer",
		},
	}

	logObj(t, "Request:", setupRequest)

	setupResponse, err := client.Enroll(setupRequest)

	assert.NoError(err)

	logObj(t, "Response:", setupResponse)

	// setup request object
	request := blockchyp.LinkTokenRequest{
		Token:      setupResponse.Token,
		CustomerID: setupResponse.Customer.ID,
	}

	logObj(t, "Request:", request)

	response, err := client.LinkToken(request)

	assert.NoError(err)

	logObj(t, "Response:", response)

	// response assertions
	assert.True(response.Success)
}
