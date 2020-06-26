// +build regression

package regression

import (
	"fmt"
	"testing"
	"time"

	"github.com/blockchyp/blockchyp-go"
)

func TestBatch(t *testing.T) {
	tests := map[string]struct {
		instructions string
		args         [][]string
		assert       []interface{}
		batchID      string
	}{
		"Charge": {
			instructions: "Insert an EMV test card when prompted.",
			args: [][]string{
				{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", amount(0),
				},
				{
					"-type", "batch-history", "-test",
					"-maxResults", "1",
					"-startDate", time.Now().Add(-5 * time.Second).Format(time.RFC3339),
				},
				{
					"-type", "close-batch", "-test",
				},
				{
					"-type", "close-batch", "-test",
				},
				{
					"-type", "batch-history", "-test",
					"-maxResults", "1",
					"-startDate", time.Now().Add(-5 * time.Second).Format(time.RFC3339),
				},
				{
					"-type", "batch-details", "-test",
					"-batchId",
				},
				{
					"-type", "tx-history", "-test",
					"-batchId",
				},
			},
			assert: []interface{}{
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "charge",
					AuthorizedAmount: amount(0),
				},
				blockchyp.BatchHistoryResponse{
					Success: true,
					Batches: []blockchyp.BatchSummary{
						{
							Open: true,
						},
					},
				},
				blockchyp.CloseBatchResponse{
					Success: true,
					Test:    true,
					Batches: []blockchyp.BatchSummary{
						{
							CapturedAmount: amount(0),
							OpenPreauths:   "0.00",
						},
					},
				},
				blockchyp.CloseBatchResponse{
					Success:             false,
					Test:                true,
					ResponseDescription: "No batch",
				},
				blockchyp.BatchHistoryResponse{
					Success: true,
					Batches: []blockchyp.BatchSummary{
						{
							Open: false,
						},
					},
				},
				blockchyp.BatchDetailsResponse{
					Success:          true,
					CapturedAmount:   amount(0),
					TotalVolume:      amount(0),
					TransactionCount: 1,
					Open:             false,
				},
				blockchyp.TransactionHistoryResponse{
					Success:          true,
					TotalResultCount: 1,
					Transactions: []blockchyp.AuthorizationResponse{
						{
							Success:          true,
							Approved:         true,
							Test:             true,
							TransactionType:  "charge",
							AuthorizedAmount: amount(0),
						},
					},
				},
			},
		},
		"PreauthRollover": {
			instructions: `Insert an EMV test card when prompted.

Leave it in the terminal until the test completes.`,
			args: [][]string{
				{
					"-type", "preauth", "-terminal", terminalName, "-test",
					"-amount", amount(0),
				},
				{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", amount(1),
				},
				{
					"-type", "close-batch", "-test",
				},
			},
			assert: []interface{}{
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "preauth",
					AuthorizedAmount: amount(0),
				},
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "charge",
					AuthorizedAmount: amount(1),
				},
				blockchyp.CloseBatchResponse{
					Success: true,
					Test:    true,
					Batches: []blockchyp.BatchSummary{
						{
							CapturedAmount: amount(1),
							OpenPreauths:   amount(0),
						},
					},
				},
			},
		},
		"BadCredentials": {
			args: [][]string{
				{
					"-type", "close-batch", "-test",
					"-apiKey", "X6N2KIQEWYI6TCADNSLM7WZLHE",
					"-bearerToken", "RIKLAPSMSMG2YII27N2NPAMCS5",
					"-signingKey", "4b556bc4e73ffc86fc5f8bfbba1598e7a8cd91f44fd7072d070c92fae7f48cd9",
				},
			},
			assert: []interface{}{
				blockchyp.CloseBatchResponse{
					Success: false,
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			cli := newCLI(t)

			setup(t, test.instructions, true)

			// Close the batch and ignore the result to make sure we're
			// starting from a clean slate.
			cli.run([]string{"-type", "close-batch", "-test"}, struct{}{})

			for i := range test.args {
				if test.batchID != "" && test.args[i][len(test.args[i])-1] == "-batchId" {
					fmt.Println("foo")
					test.args[i] = append(test.args[i], test.batchID)
				}

				res := cli.run(test.args[i], test.assert[i])
				switch v := res.(type) {
				case *blockchyp.AuthorizationResponse:
					if test.batchID == "" {
						test.batchID = v.BatchID
					}
					fmt.Println(test.batchID)
				}
			}
		})
	}
}
