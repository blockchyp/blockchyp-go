// +build manual

package itests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	blockchyp "github.com/blockchyp/blockchyp-go"
)

func TestTermsAndConditionsTest(t *testing.T) {

	assert := assert.New(t)

	client := newTestClient(t)

	// setup request object
	request := blockchyp.TermsAndConditionsRequest{}
	request.Test = true
	request.TerminalName = "Test Terminal"
	request.TCName = "HIPPA Disclosure"
	request.TCContent = "Full contract text"
	request.SigFormat = "png"
	request.SigWidth = 200
	request.SigRequired = true
	logRequest(request)

	response, err := client.TC(request)

	assert.NoError(err)

	logResponse(response)

	// response assertions
	assert.True(response.Success)
}