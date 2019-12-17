// +build manual

package itests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	blockchyp "github.com/blockchyp/blockchyp-go"
)

func TestSimpleReversal(t *testing.T) {

	assert := assert.New(t)

	client := newTestClient(t)

	// setup request object
	request0 := blockchyp.AuthorizationRequest{}

	request0.PAN = "4111111111111111"

	request0.Amount = "25.55"

	request0.Test = true

	logRequest(request0)

	response0, err := client.Charge(request0)

	assert.NoError(err)

	logResponse(response0)

	// setup request object
	request := blockchyp.AuthorizationRequest{}
	request.TransactionRef = lastTransactionRef
	request.Test = true
	logRequest(request)

	response, err := client.Reverse(request)

	assert.NoError(err)

	logResponse(response)

	// response assertions
	assert.True(response.Approved)
}