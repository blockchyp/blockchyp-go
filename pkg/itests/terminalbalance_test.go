// +build manual

package itests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	blockchyp "github.com/blockchyp/blockchyp-go"
)

func TestTerminalBalance(t *testing.T) {

	assert := assert.New(t)

	client := newTestClient(t)

	// setup request object
	request := blockchyp.BalanceRequest{}
	request.Test = true
	request.TerminalName = "Test Terminal"
	request.CardType = 3
	logRequest(request)

	response, err := client.Balance(request)

	assert.NoError(err)

	logResponse(response)

	// response assertions
	assert.True(response.Success)
	assert.NotEmpty(response.RemainingBalance)
}