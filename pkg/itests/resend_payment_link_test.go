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

func TestResendPaymentLink(t *testing.T) {
	assert := assert.New(t)

	config := loadTestConfiguration(t)
	client := config.newTestClient(t, "")

	// setup request object
	setupRequest := blockchyp.PaymentLinkRequest{
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
			EmailAddress: "notifications@blockchypteam.m8r.co",
			SmsNumber:    "(123) 123-1231",
		},
	}

	logObj(t, "Request:", setupRequest)

	setupResponse, err := client.SendPaymentLink(setupRequest)

	assert.NoError(err)

	logObj(t, "Response:", setupResponse)

	// setup request object
	request := blockchyp.ResendPaymentLinkRequest{
		Test:     true,
		LinkCode: setupResponse.LinkCode,
	}

	logObj(t, "Request:", request)

	response, err := client.ResendPaymentLink(request)

	assert.NoError(err)

	logObj(t, "Response:", response)

	// response assertions
	assert.True(response.Success)
}
