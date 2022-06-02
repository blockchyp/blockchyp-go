//go:build integration
// +build integration

// Copyright 2019-2022 BlockChyp, Inc. All rights reserved. Use of this code
// is governed by a license that can be found in the LICENSE file.
//
// This file was generated automatically by the BlockChyp SDK Generator.
// Changes to this file will be lost every time the code is regenerated.

package itests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	blockchyp "github.com/blockchyp/blockchyp-go"
)

func TestDeleteSurveyQuestion(t *testing.T) {
	assert := assert.New(t)

	config := loadTestConfiguration(t)
	client := config.newTestClient(t, "")

	// setup request object
	setupRequest := blockchyp.SurveyQuestion{
		Ordinal:      1,
		QuestionText: "Would you shop here again?",
		QuestionType: "yes_no",
	}

	logObj(t, "Request:", setupRequest)

	setupResponse, err := client.UpdateSurveyQuestion(setupRequest)

	assert.NoError(err)

	logObj(t, "Response:", setupResponse)

	// setup request object
	request := blockchyp.SurveyQuestionRequest{
		QuestionID: setupResponse.ID,
	}

	logObj(t, "Request:", request)

	response, err := client.DeleteSurveyQuestion(request)

	assert.NoError(err)

	logObj(t, "Response:", response)

	// response assertions
	assert.True(response.Success)
}
