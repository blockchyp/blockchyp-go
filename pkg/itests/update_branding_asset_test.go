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

	"os"

	"github.com/stretchr/testify/assert"

	blockchyp "github.com/blockchyp/blockchyp-go/v2"
)

func TestUpdateBrandingAsset(t *testing.T) {
	assert := assert.New(t)

	config := loadTestConfiguration(t)
	client := config.newTestClient(t, "")

	// setup request object
	setupRequest := blockchyp.UploadMetadata{
		FileName: "aviato.png",
		FileSize: 18843,
		UploadID: randomID(),
	}

	logObj(t, "Request:", setupRequest)

	file, err := os.Open("testdata/aviato.png")
	assert.NoError(err)
	setupResponse, err := client.UploadMedia(setupRequest, file)

	assert.NoError(err)

	logObj(t, "Response:", setupResponse)

	// setup request object
	request := blockchyp.BrandingAsset{
		MediaID:   setupResponse.ID,
		Padded:    true,
		Ordinal:   10,
		StartDate: "01/06/2021",
		StartTime: "14:00",
		EndDate:   "11/05/2024",
		EndTime:   "16:00",
		Notes:     "Test Branding Asset",
		Preview:   false,
		Enabled:   true,
	}

	logObj(t, "Request:", request)

	response, err := client.UpdateBrandingAsset(request)

	assert.NoError(err)

	logObj(t, "Response:", response)

	// response assertions
	assert.True(response.Success)
}
