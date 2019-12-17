// +build manual

package itests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	blockchyp "github.com/blockchyp/sdk-generator/platforms/go/project"
)

func TestSimplePing(t *testing.T) {

	assert := assert.New(t)

	client := newTestClient(t)

	// setup request object
	request := blockchyp.PingRequest{}
	request.Test = true
	request.TerminalName = "Test Terminal"
	logRequest(request)

	response, err := client.Ping(request)

	assert.NoError(err)

	logResponse(response)

	// response assertions
	assert.True(response.Success)
}