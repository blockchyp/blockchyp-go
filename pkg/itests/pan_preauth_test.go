// +build integration
// Copyright 2019 BlockChyp, Inc. All rights reserved. Use of this code is
// governed by a license that can be found in the LICENSE file.
//
// This file was generated automatically. Changes to this file will be lost
// every time the code is regenerated.

package itests

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	blockchyp "github.com/blockchyp/blockchyp-go"
)

func TestPANPreauth(t *testing.T) {

	assert := assert.New(t)

	client := newTestClient(t)

	testDelay := os.Getenv(TestDelay)
	if testDelay != "" {
		testDelayInt, err := strconv.Atoi(testDelay)
		assert.NoError(err)
		messageRequest := blockchyp.MessageRequest{
			TerminalName: "Test Terminal",
			Test:         true,
			Message:      fmt.Sprintf("Running TestPANPreauth in %v seconds...", testDelay),
		}
		messageResponse, err := client.Message(messageRequest)
		assert.NoError(err)
		assert.True(true, messageResponse.Success)
		time.Sleep(time.Duration(testDelayInt) * time.Second)
	}

	// setup request object
	request := blockchyp.AuthorizationRequest{}
	request.PAN = "4111111111111111"
	request.Amount = "25.55"
	request.Test = true
	logRequest(request)

	response, err := client.Preauth(request)

	assert.NoError(err)

	logResponse(response)

	// response assertions
	assert.True(response.Approved)
	assert.True(response.Test)
	assert.Len(response.AuthCode, 6)
	assert.NotEmpty(response.TransactionID)
	assert.NotEmpty(response.Timestamp)
	assert.NotEmpty(response.TickBlock)
	assert.Equal("Approved", response.ResponseDescription)
	assert.NotEmpty(response.PaymentType)
	assert.NotEmpty(response.MaskedPAN)
	assert.NotEmpty(response.EntryMethod)
	assert.Equal("25.55", response.AuthorizedAmount)
	assert.Equal("KEYED", response.EntryMethod)
}