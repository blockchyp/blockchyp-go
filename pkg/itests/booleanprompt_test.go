// +build manual

package itests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	blockchyp "github.com/blockchyp/sdk-generator/platforms/go/project"
)

func TestBooleanPrompt(t *testing.T) {

	assert := assert.New(t)

	client := newTestClient(t)

	// setup request object
	request := blockchyp.BooleanPromptRequest{}
	request.Test = true
	request.TerminalName = "Test Terminal"
	request.Prompt = "Would you like to become a member?"
	request.YesCaption = "Yes"
	request.NoCaption = "No"
	logRequest(request)

	response, err := client.BooleanPrompt(request)

	assert.NoError(err)

	logResponse(response)

	// response assertions
	assert.True(response.Success)
	assert.True(response.Response)
}