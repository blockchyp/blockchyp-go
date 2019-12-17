// +build manual

package itests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	blockchyp "github.com/blockchyp/blockchyp-go"
)

func TestTextPrompt(t *testing.T) {

	assert := assert.New(t)

	client := newTestClient(t)

	// setup request object
	request := blockchyp.TextPromptRequest{}
	request.Test = true
	request.TerminalName = "Test Terminal"
	request.PromptType = "email"
	logRequest(request)

	response, err := client.TextPrompt(request)

	assert.NoError(err)

	logResponse(response)

	// response assertions
	assert.True(response.Success)
	assert.NotEmpty(response.Response)
}