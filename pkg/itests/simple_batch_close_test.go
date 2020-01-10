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

func TestSimpleBatchClose(t *testing.T) {
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
			Message:      fmt.Sprintf("Running TestSimpleBatchClose in %v seconds...", testDelay),
		}
		if _, err := client.Message(messageRequest); err != nil {
			t.Fatal(err)
		}
		time.Sleep(time.Duration(testDelayInt) * time.Second)
	}

	// setup request object
	request0 := blockchyp.AuthorizationRequest{
		PAN:            "4111111111111111",
		Amount:         "25.55",
		Test:           true,
		TransactionRef: randomID(),
	}

	logRequest(request0)

	response0, err := client.Charge(request0)

	assert.NoError(err)

	logResponse(response0)

	// setup request object
	request := blockchyp.CloseBatchRequest{
		Test: true,
	}

	logRequest(request)

	response, err := client.CloseBatch(request)

	assert.NoError(err)

	logResponse(response)

	// response assertions
	assert.True(response.Success)
	assert.NotEmpty(response.CapturedTotal)
	assert.NotEmpty(response.OpenPreauths)
}
