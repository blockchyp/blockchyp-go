// +build manual

package itests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	blockchyp "github.com/blockchyp/sdk-generator/platforms/go/project"
)

func TestUpdateTransactionDisplay(t *testing.T) {

	assert := assert.New(t)

	client := newTestClient(t)

	// setup request object
	request := blockchyp.TransactionDisplayRequest{}
	request.Test = true
	request.TerminalName = "Test Terminal"
	request.Transaction = &blockchyp.TransactionDisplayTransaction{
		Subtotal: "35.00",
		Tax: "5.00",
		Total: "70.00",
		Items: []*blockchyp.TransactionDisplayItem{
			&blockchyp.TransactionDisplayItem{
				Description: "Leki Trekking Poles",
				Price: "35.00",
				Quantity: 2,
				Extended: "70.00",
				Discounts: []*blockchyp.TransactionDisplayDiscount{
					&blockchyp.TransactionDisplayDiscount{
						Description: "memberDiscount",
						Amount: "10.00",
					},
				},
			},
		},
	}
	logRequest(request)

	response, err := client.UpdateTransactionDisplay(request)

	assert.NoError(err)

	logResponse(response)

	// response assertions
	assert.True(response.Success)
}