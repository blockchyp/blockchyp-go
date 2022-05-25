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

func TestUpdateBrandingAsset(t *testing.T) {
	assert := assert.New(t)

	config := loadTestConfiguration(t)
	client := config.newTestClient(t)

	testDelay := os.Getenv(TestDelay)
	if testDelay != "" {
		testDelayInt, err := strconv.Atoi(testDelay)
		if err != nil {
			t.Fatal(err)
		}
		messageRequest := blockchyp.MessageRequest{
			TerminalName: config.DefaultTerminalName,
			Test:         true,
			Message:      fmt.Sprintf("Running TestUpdateBrandingAsset in %v seconds...", testDelay),
		}
		if _, err := client.Message(messageRequest); err != nil {
			t.Fatal(err)
		}
		time.Sleep(time.Duration(testDelayInt) * time.Second)
	}

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
