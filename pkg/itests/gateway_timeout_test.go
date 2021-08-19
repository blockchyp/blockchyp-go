//go:build integration
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

func TestGatewayTimeout(t *testing.T) {
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
			Message:      fmt.Sprintf("Running TestGatewayTimeout in %v seconds...", testDelay),
		}
		if _, err := client.Message(messageRequest); err != nil {
			t.Fatal(err)
		}
		time.Sleep(time.Duration(testDelayInt) * time.Second)
	}

	// setup request object
	request := blockchyp.AuthorizationRequest{
		Timeout:        1,
		PAN:            "5555555555554444",
		ExpMonth:       "12",
		ExpYear:        "2025",
		Amount:         "25.55",
		Test:           true,
		TransactionRef: randomID(),
	}

	logRequest(request)

	response, err := client.Charge(request)

	logResponse(response)
	t.Logf("Response Error: %+v", err)

	assert.Error(err)
	assert.Equal(blockchyp.ResponseTimedOut, response.ResponseDescription)

	return
}
