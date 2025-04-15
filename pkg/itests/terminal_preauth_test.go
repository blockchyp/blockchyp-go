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

func TestTerminalPreauth(t *testing.T) {
	assert := assert.New(t)

	config := loadTestConfiguration(t)
	client := config.newTestClient(t, "")

	processTestDelay(t, config, "TerminalPreauth")

	// setup request object
	request := blockchyp.AuthorizationRequest{
		TerminalName: config.DefaultTerminalName,
		Amount:       "15.15",
		Test:         true,
	}

	logObj(t, "Request:", request)

	response, err := client.Preauth(request)

	assert.NoError(err)

	logObj(t, "Response:", response)

	// response assertions
	assert.True(response.Success)
	assert.True(response.Approved)
	assert.True(response.Test)
	assert.Len(response.AuthCode, 6)
	assert.NotEmpty(response.TransactionID)
	assert.NotEmpty(response.Timestamp)
	assert.NotEmpty(response.TickBlock)
	assert.Equal("approved", response.ResponseDescription)
	assert.NotEmpty(response.PaymentType)
	assert.NotEmpty(response.MaskedPAN)
	assert.NotEmpty(response.EntryMethod)
	assert.Equal("15.15", response.AuthorizedAmount)
}
