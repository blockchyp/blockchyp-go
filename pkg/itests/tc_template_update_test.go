//go:build integration
// +build integration

// Copyright 2019-2025 BlockChyp, Inc. All rights reserved. Use of this code
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

func TestTCTemplateUpdate(t *testing.T) {
	assert := assert.New(t)

	config := loadTestConfiguration(t)
	client := config.newTestClient(t, "")

	// setup request object
	request := blockchyp.TermsAndConditionsTemplate{
		Alias:   randomID(),
		Name:    "HIPPA Disclosure",
		Content: "Lorem ipsum dolor sit amet.",
	}

	logObj(t, "Request:", request)

	response, err := client.TCUpdateTemplate(request)

	assert.NoError(err)

	logObj(t, "Response:", response)

	// response assertions
	assert.True(response.Success)
	assert.NotEmpty(response.Alias)
	assert.Equal("HIPPA Disclosure", response.Name)
	assert.Equal("Lorem ipsum dolor sit amet.", response.Content)
}
