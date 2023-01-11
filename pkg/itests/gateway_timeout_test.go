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

	blockchyp "github.com/blockchyp/blockchyp-go"
)

func TestGatewayTimeout(t *testing.T) {
	assert := assert.New(t)

	config := loadTestConfiguration(t)
	client := config.newTestClient(t, "")

	processTestDelay(t, config, "GatewayTimeout")

	// setup request object
	request := blockchyp.AuthorizationRequest{
		Timeout:        1,
		PAN:            "5555555555554444",
		ExpMonth:       "12",
		ExpYear:        "2025",
		Amount:         "25.55",
		Test:           true,
		TransactionRef: randomID(),
	}

	logObj(t, "Request:", request)

	response, err := client.Charge(request)

	logObj(t, "Response:", response)
	t.Logf("Response Error: %+v", err)

	assert.Error(err)
	assert.Equal(blockchyp.ResponseTimedOut, response.ResponseDescription)

	return
}
