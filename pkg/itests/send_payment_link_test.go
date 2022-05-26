//go:build integration
// +build integration

// Copyright 2019-2022 BlockChyp, Inc. All rights reserved. Use of this code
// is governed by a license that can be found in the LICENSE file.
//
// This file was generated automatically by the BlockChyp SDK Generator.
// Changes to this file will be lost every time the code is regenerated.

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

func TestSendPaymentLink(t *testing.T) {
	assert := assert.New(t)

	config := loadTestConfiguration(t)
	client := config.newTestClient(t, "")

	testDelay := os.Getenv(TestDelay)
	if testDelay != "" {
		testDelayInt, err := strconv.Atoi(testDelay)
		if err != nil {
			t.Fatal(err)
		}
		messageRequest := blockchyp.MessageRequest{
			TerminalName: config.DefaultTerminalName,
			Test:         true,
			Message:      fmt.Sprintf("Running TestSendPaymentLink in %v seconds...", testDelay),
		}
		if _, err := client.Message(messageRequest); err != nil {
			t.Fatal(err)
		}
		time.Sleep(time.Duration(testDelayInt) * time.Second)
	}

	// setup request object
	request := blockchyp.PaymentLinkRequest{
		Amount:      "199.99",
		Description: "Widget",
		Subject:     "Widget invoice",
		Transaction: &blockchyp.TransactionDisplayTransaction{
			Subtotal: "195.00",
			Tax:      "4.99",
			Total:    "199.99",
			Items: []*blockchyp.TransactionDisplayItem{
				&blockchyp.TransactionDisplayItem{
					Description: "Widget",
					Price:       "195.00",
					Quantity:    1,
				},
			},
		},
		AutoSend: true,
		Customer: blockchyp.Customer{
			CustomerRef:  "Customer reference string",
			FirstName:    "FirstName",
			LastName:     "LastName",
			CompanyName:  "Company Name",
			EmailAddress: "support@blockchyp.com",
			SmsNumber:    "(123) 123-1231",
		},
	}

	logObj(t, "Request:", request)

	response, err := client.SendPaymentLink(request)

	assert.NoError(err)

	logObj(t, "Response:", response)

	// response assertions
	assert.True(response.Success)
	assert.NotEmpty(response.URL)
}
