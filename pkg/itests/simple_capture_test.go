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

func TestSimpleCapture(t *testing.T) {
	assert := assert.New(t)

	client := newTestClient(t)

	testDelay := os.Getenv(TestDelay)
	if testDelay != "" {
		testDelayInt, err := strconv.Atoi(testDelay)
		if err != nil {
			t.Fatal(err)
		}
		messageRequest := blockchyp.MessageRequest{
			TerminalName: "Test Terminal",
			Test:         true,
			Message:      fmt.Sprintf("Running TestSimpleCapture in %v seconds...", testDelay),
		}
		if _, err := client.Message(messageRequest); err != nil {
			t.Fatal(err)
		}
		time.Sleep(time.Duration(testDelayInt) * time.Second)
	}

	// setup request object
	setupRequest := blockchyp.AuthorizationRequest{
		PAN:    "4111111111111111",
		Amount: "25.55",
		Test:   true,
	}

	logRequest(setupRequest)

	setupResponse, err := client.Preauth(setupRequest)

	assert.NoError(err)

	logResponse(setupResponse)

	// setup request object
	request := blockchyp.CaptureRequest{
		TransactionID: setupResponse.TransactionID,
		Test:          true,
	}

	logRequest(request)

	response, err := client.Capture(request)

	assert.NoError(err)

	logResponse(response)

	// response assertions
	assert.True(response.Success)
	assert.True(response.Approved)
}
