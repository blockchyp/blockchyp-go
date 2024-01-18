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

func TestTCEntry(t *testing.T) {
	assert := assert.New(t)

	config := loadTestConfiguration(t)
	client := config.newTestClient(t, "")

	// setup request object
	setupRequest := blockchyp.TermsAndConditionsLogRequest{}

	logObj(t, "Request:", setupRequest)

	setupResponse, err := client.TCLog(setupRequest)

	assert.NoError(err)

	logObj(t, "Response:", setupResponse)

	// setup request object
	request := blockchyp.TermsAndConditionsLogRequest{
		LogEntryID: setupResponse.Results[0].ID,
	}

	logObj(t, "Request:", request)

	response, err := client.TCEntry(request)

	assert.NoError(err)

	logObj(t, "Response:", response)

	// response assertions
	assert.True(response.Success)
	assert.NotEmpty(response.ID)
	assert.NotEmpty(response.TerminalID)
	assert.NotEmpty(response.TerminalName)
	assert.NotEmpty(response.Timestamp)
	assert.NotEmpty(response.Name)
	assert.NotEmpty(response.Content)
	assert.True(response.HasSignature)
	assert.NotEmpty(response.Signature)
}
