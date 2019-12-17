// +build manual

package itests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	blockchyp "github.com/blockchyp/sdk-generator/platforms/go/project"
)

func TestSimpleMessage(t *testing.T) {

	assert := assert.New(t)

	client := newTestClient(t)

	// setup request object
	request := blockchyp.MessageRequest{}
	request.Test = true
	request.TerminalName = "Test Terminal"
	request.Message = "Thank You For Your Business"
	logRequest(request)

	response, err := client.Message(request)

	assert.NoError(err)

	logResponse(response)

	// response assertions
	assert.True(response.Success)
}