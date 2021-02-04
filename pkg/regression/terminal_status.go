package regression

import (
	"time"

	"github.com/blockchyp/blockchyp-go"
)

var terminalStatusTests = testCases{
	{
		name: "TerminalStatus/Idle",
		sim:  true,
		operations: []operation{
			{
				args: []string{
					"-type", "terminal-status", "-terminal", terminalName,
				},
				expect: blockchyp.TerminalStatusResponse{
					Success: true,
					Idle:    true,
				},
			},
		},
	},
	{
		name: "TerminalStatus/Charge",
		sim:  true,
		operations: []operation{
			{
				background: true,
				args: []string{
					"-type", "charge", "-terminal", terminalName, "-test", "-amount", "1.00",
				},
			},
			{
				wait: 2 * time.Second,
				args: []string{
					"-type", "terminal-status", "-terminal", terminalName,
				},
				expect: blockchyp.TerminalStatusResponse{
					Success: true,
					Idle:    false,
					Status:  "charge",
				},
			},
			{
				args: []string{
					"-type", "clear", "-terminal", terminalName,
				},
				expect: blockchyp.Acknowledgement{
					Success: true,
				},
			},
		},
	},
	{
		name: "TerminalStatus/TC",
		sim:  true,
		operations: []operation{
			{
				background: true,
				args: []string{
					"-type", "tc", "-terminal", terminalName, "-test",
					"-tcName", "Contract Title", "-tcContent", "Blah Blah Blah",
				},
			},
			{
				wait: 2 * time.Second,
				args: []string{
					"-type", "terminal-status", "-terminal", terminalName,
				},
				expect: blockchyp.TerminalStatusResponse{
					Success: true,
					Idle:    false,
					Status:  "terms-and-conditions",
				},
			},
			{
				args: []string{
					"-type", "clear", "-terminal", terminalName,
				},
				expect: blockchyp.Acknowledgement{
					Success: true,
				},
			},
		},
	},
}
