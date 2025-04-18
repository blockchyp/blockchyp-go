//go:build integration
// +build integration

// Copyright 2019-2025 BlockChyp, Inc. All rights reserved. Use of this code
// is governed by a license that can be found in the LICENSE file.
//
// This file was generated automatically by the BlockChyp SDK Generator.
// Changes to this file will be lost every time the code is regenerated.

package itests

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	blockchyp "github.com/blockchyp/blockchyp-go/v2"
)

func TestMediaUpload(t *testing.T) {
	assert := assert.New(t)

	config := loadTestConfiguration(t)
	client := config.newTestClient(t, "")

	// setup request object
	request := blockchyp.UploadMetadata{
		FileName: "aviato.png",
		FileSize: 18843,
		UploadID: randomID(),
	}

	logObj(t, "Request:", request)

	file, err := os.Open("testdata/aviato.png")
	assert.NoError(err)
	response, err := client.UploadMedia(request, file)

	assert.NoError(err)

	logObj(t, "Response:", response)

	// response assertions
	assert.True(response.Success)
	assert.NotEmpty(response.ID)
	assert.Equal("aviato.png", response.OriginalFile)
	assert.NotEmpty(response.FileURL)
	assert.NotEmpty(response.ThumbnailURL)
}
