// +build manual

package itests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	blockchyp "github.com/blockchyp/sdk-generator/platforms/go/project"
)

func TestSimpleGiftActivate(t *testing.T) {

	assert := assert.New(t)

	client := newTestClient(t)

	// setup request object
	request := blockchyp.GiftActivateRequest{}
	request.Test = true
	request.TerminalName = "Test Terminal"
	request.Amount = "50.00"
	logRequest(request)

	response, err := client.GiftActivate(request)

	assert.NoError(err)

	logResponse(response)

	// response assertions
	assert.True(response.Approved)
	assert.NotEmpty(response.PublicKey)
}