// +build manual

package itests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	blockchyp "github.com/blockchyp/sdk-generator/platforms/go/project"
)

func TestSimpleBatchClose(t *testing.T) {

	assert := assert.New(t)

	client := newTestClient(t)

	// setup request object
	request := blockchyp.CloseBatchRequest{}
	request.Test = true
	logRequest(request)

	response, err := client.CloseBatch(request)

	assert.NoError(err)

	logResponse(response)

	// response assertions
	assert.True(response.Success)
	assert.NotEmpty(response.CapturedTotal)
	assert.NotEmpty(response.OpenPreauths)
}