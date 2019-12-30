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

func TestTermsAndConditions(t *testing.T) {

	assert := assert.New(t)

	client := newTestClient(t)

	// setup request object
	request := blockchyp.TermsAndConditionsRequest{}
	request.Test = true
	request.TerminalName = "Test Terminal"
	request.TCName = "HIPPA Disclosure"
	request.TCContent = "Full contract text"
	request.SigFormat = "png"
	request.SigWidth = 200
	request.SigRequired = true
	logRequest(request)

	response, err := client.TermsAndConditions(request)

	assert.NoError(err)

	logResponse(response)

	// response assertions
	assert.True(response.Success)
}
