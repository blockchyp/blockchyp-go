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

func TestNewTransactionDisplay(t *testing.T) {
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
			Message:      fmt.Sprintf("Running TestNewTransactionDisplay in %v seconds...", testDelay),
		}
		if _, err := client.Message(messageRequest); err != nil {
			t.Fatal(err)
		}
		time.Sleep(time.Duration(testDelayInt) * time.Second)
	}

	// setup request object
	request := blockchyp.TransactionDisplayRequest{
		Test:         true,
		TerminalName: config.DefaultTerminalName,
		Transaction: &blockchyp.TransactionDisplayTransaction{
			Subtotal: "35.00",
			Tax:      "5.00",
			Total:    "70.00",
			Items: []*blockchyp.TransactionDisplayItem{
				&blockchyp.TransactionDisplayItem{
					Description: "Leki Trekking Poles",
					Price:       "35.00",
					Quantity:    2,
					Extended:    "70.00",
					Discounts: []*blockchyp.TransactionDisplayDiscount{
						&blockchyp.TransactionDisplayDiscount{
							Description: "memberDiscount",
							Amount:      "10.00",
						},
					},
				},
			},
		},
	}

	logObj(t, "Request:", request)

	response, err := client.NewTransactionDisplay(request)

	assert.NoError(err)

	logObj(t, "Response:", response)

	// response assertions
	assert.True(response.Success)
}
