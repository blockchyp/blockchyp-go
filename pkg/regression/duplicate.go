package regression

import (
	"github.com/blockchyp/blockchyp-go/v2"
)

var duplicateTests = testCases{
	{
		name:  "Duplicate/Charge",
		group: testGroupNoCVM,
		operations: []operation{
			{
				args: []string{
					"-type", "charge", "-terminal", terminalName,
					"-test",
					"-amount", randomAmount(),
				},
				expect: blockchyp.AuthorizationResponse{
					Success:  true,
					Approved: true,
					Test:     true,
				},
			},
			{
				args: []string{
					"-type", "charge", "-terminal", terminalName,
					"-test",
					"-amount", previousAmount,
				},
				expect: blockchyp.AuthorizationResponse{
					Success:             true,
					Approved:            false,
					Test:                true,
					ResponseDescription: "Duplicate Transaction",
				},
			},
		},
	},
}
