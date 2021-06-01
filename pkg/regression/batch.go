package regression

import (
	"time"

	"github.com/blockchyp/blockchyp-go"
)

var batchTests = testCases{
	{
		name:  "Batch/Charge",
		group: testGroupNoCVM,
		operations: []operation{
			{
				// Close the batch and ignore the result to make sure we're
				// starting from a clean slate.
				args: []string{
					"-type", "close-batch", "-test",
				},
			},
			{
				msg: "Insert an EMV test card when prompted.",
				args: []string{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", amount(0),
				},
				expect: blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "charge",
					AuthorizedAmount: amount(0),
				},
			},
			{
				args: []string{
					"-type", "batch-history", "-test",
					"-maxResults", "1",
					"-startDate", time.Now().Add(-5 * time.Second).Format(time.RFC3339),
				},
				expect: blockchyp.BatchHistoryResponse{
					Success: true,
					Test:    true,
					Batches: []blockchyp.BatchSummary{
						{
							Open: true,
						},
					},
				},
			},
			{
				args: []string{
					"-type", "close-batch", "-test",
				},
				expect: blockchyp.CloseBatchResponse{
					Success: true,
					Test:    true,
					Batches: []blockchyp.BatchSummary{
						{
							CapturedAmount: amount(0),
							OpenPreauths:   "0.00",
						},
					},
				},
			},
			{
				args: []string{
					"-type", "close-batch", "-test",
				},
				expect: blockchyp.CloseBatchResponse{
					Success:             false,
					Test:                true,
					ResponseDescription: "no open batches",
				},
			},
			{
				args: []string{
					"-type", "batch-history", "-test",
					"-maxResults", "1",
					"-startDate", time.Now().Add(-5 * time.Second).Format(time.RFC3339),
				},
				expect: blockchyp.BatchHistoryResponse{
					Success: true,
					Test:    true,
					Batches: []blockchyp.BatchSummary{
						{
							Open: false,
						},
					},
				},
			},
			{
				args: []string{
					"-type", "batch-details", "-test", "-batchId", batchIDN(1),
				},
				expect: blockchyp.BatchDetailsResponse{
					Success:          true,
					Test:             true,
					CapturedAmount:   amount(0),
					TotalVolume:      amount(0),
					TransactionCount: 1,
					Open:             false,
				},
			},
			{
				args: []string{
					"-type", "tx-history", "-test", "-batchId", batchIDN(1),
				},
				expect: blockchyp.TransactionHistoryResponse{
					Success:          true,
					Test:             true,
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
	},
	{
		name:  "Batch/PreauthRollover",
		group: testGroupNoCVM,
		operations: []operation{
			{
				// Close the batch and ignore the result to make sure we're
				// starting from a clean slate.
				args: []string{
					"-type", "close-batch", "-test",
				},
			},
			{
				msg: "Insert an EMV test card when prompted.",
				args: []string{
					"-type", "preauth", "-terminal", terminalName, "-test",
					"-amount", amount(0),
				},
				expect: blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "preauth",
					AuthorizedAmount: amount(0),
				},
			},
			{
				args: []string{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", amount(1),
				},
				expect: blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "charge",
					AuthorizedAmount: amount(1),
				},
			},
			{
				args: []string{
					"-type", "close-batch", "-test",
				},
				expect: blockchyp.CloseBatchResponse{
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
	},
	{
		name:  "Batch/BadCredentials",
		group: testGroupNonInteractive,
		operations: []operation{
			{
				args: []string{
					"-type", "close-batch", "-test",
					"-apiKey", "X6N2KIQEWYI6TCADNSLM7WZLHE",
					"-bearerToken", "RIKLAPSMSMG2YII27N2NPAMCS5",
					"-signingKey", "4b556bc4e73ffc86fc5f8bfbba1598e7a8cd91f44fd7072d070c92fae7f48cd9",
				},
				expect: blockchyp.CloseBatchResponse{
					Success: false,
				},
			},
		},
	},
}
