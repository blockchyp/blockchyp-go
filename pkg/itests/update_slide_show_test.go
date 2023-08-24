//go:build integration
// +build integration

// Copyright 2019-2023 BlockChyp, Inc. All rights reserved. Use of this code
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

func TestUpdateSlideShow(t *testing.T) {
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
	request := blockchyp.SlideShow{
		Name:  "Test Slide Show",
		Delay: 5,
		Slides: []*blockchyp.Slide{
			&blockchyp.Slide{
				MediaID: setupResponse.ID,
			},
		},
	}

	logObj(t, "Request:", request)

	response, err := client.UpdateSlideShow(request)

	assert.NoError(err)

	logObj(t, "Response:", response)

	// response assertions
	assert.True(response.Success)
	assert.Equal("Test Slide Show", response.Name)
}
