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

	"github.com/stretchr/testify/assert"

	blockchyp "github.com/blockchyp/blockchyp-go/v2"
)

func TestSurveyResults(t *testing.T) {
	assert := assert.New(t)

	config := loadTestConfiguration(t)
	client := config.newTestClient(t, "")

	// setup request object
	setupRequest := blockchyp.SurveyQuestionRequest{}

	logObj(t, "Request:", setupRequest)

	setupResponse, err := client.SurveyQuestions(setupRequest)

	assert.NoError(err)

	logObj(t, "Response:", setupResponse)

	// setup request object
	request := blockchyp.SurveyResultsRequest{
		QuestionID: setupResponse.Results[0].ID,
	}

	logObj(t, "Request:", request)

	response, err := client.SurveyResults(request)

	assert.NoError(err)

	logObj(t, "Response:", response)

	// response assertions
	assert.True(response.Success)
}
