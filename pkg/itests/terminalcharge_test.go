// +build manual

package itests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	blockchyp "github.com/blockchyp/sdk-generator/platforms/go/project"
)

func TestTerminalCharge(t *testing.T) {

	assert := assert.New(t)

	client := newTestClient(t)

	// setup request object
	request := blockchyp.AuthorizationRequest{}
	request.TerminalName = "Test Terminal"
	request.Amount = "25.15"
	request.Test = true
	logRequest(request)

	response, err := client.Charge(request)

	assert.NoError(err)

	logResponse(response)

	// response assertions
	assert.True(response.Approved)
	assert.True(response.Test)
	assert.Len(response.AuthCode, 6)
	assert.NotEmpty(response.TransactionID)
	assert.NotEmpty(response.Timestamp)
	assert.NotEmpty(response.TickBlock)
	assert.Equal("Approved", response.ResponseDescription)
	assert.NotEmpty(response.PaymentType)
	assert.NotEmpty(response.MaskedPAN)
	assert.NotEmpty(response.EntryMethod)
	assert.Equal("25.15", response.AuthorizedAmount)
}