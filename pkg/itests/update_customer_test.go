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

func TestUpdateCustomer(t *testing.T) {
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
			Message:      fmt.Sprintf("Running TestUpdateCustomer in %v seconds...", testDelay),
		}
		if _, err := client.Message(messageRequest); err != nil {
			t.Fatal(err)
		}
		time.Sleep(time.Duration(testDelayInt) * time.Second)
	}

	// setup request object
	request := blockchyp.UpdateCustomerRequest{
		Customer: blockchyp.Customer{
			FirstName:    "Test",
			LastName:     "Customer",
			CompanyName:  "Test Company",
			EmailAddress: "support@blockchyp.com",
			SmsNumber:    "(123) 123-1234",
		},
	}

	logRequest(request)

	response, err := client.UpdateCustomer(request)

	assert.NoError(err)

	logResponse(response)

	// response assertions
	assert.True(response.Success)
}
