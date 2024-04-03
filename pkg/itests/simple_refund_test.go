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

func TestSimpleRefund(t *testing.T) {
	assert := assert.New(t)

	config := loadTestConfiguration(t)
	client := config.newTestClient(t, "")

	processTestDelay(t, config, "SimpleRefund")

	// setup request object
	setupRequest := blockchyp.AuthorizationRequest{
		PAN:            "4111111111111111",
		ExpMonth:       "12",
		ExpYear:        "2025",
		Amount:         "25.55",
		Test:           true,
		TransactionRef: randomID(),
	}

	logObj(t, "Request:", setupRequest)

	setupResponse, err := client.Charge(setupRequest)

	assert.NoError(err)

	logObj(t, "Response:", setupResponse)

	// setup request object
	request := blockchyp.RefundRequest{
		TransactionID: setupResponse.TransactionID,
		Test:          true,
	}

	logObj(t, "Request:", request)

	response, err := client.Refund(request)

	assert.NoError(err)

	logObj(t, "Response:", response)

	// response assertions
	assert.True(response.Success)
	assert.True(response.Approved)
}
