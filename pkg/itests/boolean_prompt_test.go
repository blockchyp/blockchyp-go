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

func TestBooleanPrompt(t *testing.T) {
	assert := assert.New(t)

	config := loadTestConfiguration(t)
	client := config.newTestClient(t, "")

	processTestDelay(t, config, "BooleanPrompt")

	// setup request object
	request := blockchyp.BooleanPromptRequest{
		Test:         true,
		TerminalName: config.DefaultTerminalName,
		Prompt:       "Would you like to become a member?",
		YesCaption:   "Yes",
		NoCaption:    "No",
	}

	logObj(t, "Request:", request)

	response, err := client.BooleanPrompt(request)

	assert.NoError(err)

	logObj(t, "Response:", response)

	// response assertions
	assert.True(response.Success)
	assert.True(response.Response)
}
