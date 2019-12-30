// +build manual
// Copyright 2019 BlockChyp, Inc. All rights reserved. Use of this code is
// governed by a license that can be found in the LICENSE file.
//
// This file was generated automatically. Changes to this file will be lost
// every time the code is regenerated.

package itests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	blockchyp "github.com/blockchyp/blockchyp-go"
)

func TestSimpleRefund(t *testing.T) {

	assert := assert.New(t)

	client := newTestClient(t)

	// setup request object
	request0 := blockchyp.AuthorizationRequest{}

	request0.PAN = "4111111111111111"

	request0.Amount = "25.55"

	request0.Test = true

	request0.TransactionRef = lastTransactionRef

	logRequest(request0)

	response0, err := client.Charge(request0)

	assert.NoError(err)

	logResponse(response0)

	// setup request object
	request := blockchyp.RefundRequest{}
	request.TransactionID = lastTransactionID
	request.Test = true
	logRequest(request)

	response, err := client.Refund(request)

	assert.NoError(err)

	logResponse(response)

	// response assertions
	assert.True(response.Approved)
}
