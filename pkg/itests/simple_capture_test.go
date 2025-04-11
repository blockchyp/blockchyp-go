//go:build integration
// +build integration

// Copyright 2019-2025 BlockChyp, Inc. All rights reserved. Use of this code
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

func TestSimpleCapture(t *testing.T) {
	assert := assert.New(t)

	config := loadTestConfiguration(t)
	client := config.newTestClient(t, "")

	// setup request object
	setupRequest := blockchyp.AuthorizationRequest{
		PAN:              "4111111111111111",
		ExpMonth:         "12",
		ExpYear:          "2025",
		Amount:           "42.45",
		Test:             true,
		BypassDupeFilter: true,
	}

	logObj(t, "Request:", setupRequest)

	setupResponse, err := client.Preauth(setupRequest)

	assert.NoError(err)

	logObj(t, "Response:", setupResponse)

	// setup request object
	request := blockchyp.CaptureRequest{
		TransactionID: setupResponse.TransactionID,
		Test:          true,
	}

	logObj(t, "Request:", request)

	response, err := client.Capture(request)

	assert.NoError(err)

	logObj(t, "Response:", response)

	// response assertions
	assert.True(response.Success)
	assert.True(response.Approved)
}
