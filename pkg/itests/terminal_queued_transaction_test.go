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

func TestTerminalQueuedTransaction(t *testing.T) {
	assert := assert.New(t)

	config := loadTestConfiguration(t)
	client := config.newTestClient(t, "")

	processTestDelay(t, config, "TerminalQueuedTransaction")

	// setup request object
	request := blockchyp.AuthorizationRequest{
		TerminalName:   config.DefaultTerminalName,
		TransactionRef: randomID(),
		Description:    "1060 West Addison",
		Amount:         "25.15",
		Test:           true,
		Queue:          true,
	}

	logObj(t, "Request:", request)

	response, err := client.Charge(request)

	assert.NoError(err)

	logObj(t, "Response:", response)

	// response assertions
	assert.True(response.Success)
	assert.False(response.Approved)
	assert.Equal("Queued", response.ResponseDescription)
}
