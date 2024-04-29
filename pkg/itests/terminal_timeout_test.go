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

func TestTerminalTimeout(t *testing.T) {
	assert := assert.New(t)

	config := loadTestConfiguration(t)
	client := config.newTestClient(t, "")

	processTestDelay(t, config, "TerminalTimeout")

	// setup request object
	request := blockchyp.AuthorizationRequest{
		Timeout:      1,
		TerminalName: config.DefaultTerminalName,
		Amount:       "25.15",
		Test:         true,
	}

	logObj(t, "Request:", request)

	response, err := client.Charge(request)

	logObj(t, "Response:", response)
	t.Logf("Response Error: %+v", err)

	assert.Error(err)
	assert.Equal(blockchyp.ResponseTimedOut, response.ResponseDescription)

	return
}
