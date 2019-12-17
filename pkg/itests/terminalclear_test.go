// +build manual

package itests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	blockchyp "github.com/blockchyp/blockchyp-go"
)

func TestTerminalClear(t *testing.T) {

	assert := assert.New(t)

	client := newTestClient(t)

	// setup request object
	request := blockchyp.ClearTerminalRequest{}
	request.Test = true
	request.TerminalName = "Test Terminal"
	logRequest(request)

	response, err := client.Clear(request)

	assert.NoError(err)

	logResponse(response)

	// response assertions
	assert.True(response.Success)
}