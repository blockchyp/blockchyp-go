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

func TestBooleanPrompt(t *testing.T) {

	assert := assert.New(t)

	client := newTestClient(t)

	testDelay := os.Getenv(TestDelay)
	if testDelay != "" {
		testDelayInt, err := strconv.Atoi(testDelay)
		assert.NoError(err)
		messageRequest := blockchyp.MessageRequest{
			TerminalName: "Test Terminal",
			Test:         true,
			Message:      fmt.Sprintf("Running TestBooleanPrompt in %v seconds...", testDelay),
		}
		messageResponse, err := client.Message(messageRequest)
		assert.NoError(err)
		assert.True(true, messageResponse.Success)
		time.Sleep(time.Duration(testDelayInt) * time.Second)
	}

	// setup request object
	request := blockchyp.BooleanPromptRequest{}
	request.Test = true
	request.TerminalName = "Test Terminal"
	request.Prompt = "Would you like to become a member?"
	request.YesCaption = "Yes"
	request.NoCaption = "No"
	logRequest(request)

	response, err := client.BooleanPrompt(request)

	assert.NoError(err)

	logResponse(response)

	// response assertions
	assert.True(response.Success)
	assert.True(response.Response)
}