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

func TestDeactivateTerminal(t *testing.T) {
	assert := assert.New(t)

	config := loadTestConfiguration(t)
	client := config.newTestClient(t, "")

	// setup request object
	request := blockchyp.TerminalDeactivationRequest{
		TerminalID: randomID(),
	}

	logObj(t, "Request:", request)

	response, err := client.DeactivateTerminal(request)

	assert.Error(err)

	logObj(t, "Response:", response)

	// response assertions
	assert.False(response.Success)
}
