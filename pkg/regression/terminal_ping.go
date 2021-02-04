package regression

import (
	"github.com/blockchyp/blockchyp-go"
)

var terminalPingTests = testCases{
	{
		name: "TerminalPing/Success",
		sim:  true,
		operations: []operation{
			{
				args: []string{
					"-type", "ping", "-terminal", terminalName,
				},
				expect: blockchyp.Acknowledgement{
					Success: true,
				},
			},
		},
	},
	{
		name: "TerminalPing/Failure",
		sim:  true,
		operations: []operation{
			{
				args: []string{
					"-type", "ping", "-terminal", "Unknown Terminal",
				},
				expect: blockchyp.Acknowledgement{
					Success: false,
					Error:   "unknown terminal",
				},
			},
		},
	},
	{
		name: "TerminalPing/BadCreds",
		sim:  true,
		operations: []operation{
			{
				args: []string{
					"-type", "ping", "-terminal", terminalName,
					"-apiKey", "RIKLAPSMSMG2YII27N2NPAMCS5",
					"-bearerToken", "RIKLAPSMSMG2YII27N2NPAMCS5",
					"-signingKey", "4b556bc4e73ffc86fc5f8bfbba1598e7a8cd91f44fd7072d070c92fae7f48cd9",
				},
				expect: blockchyp.Acknowledgement{
					Success: false,
					Error:   "Access Denied",
				},
			},
		},
	},
	{
		name: "TerminalPing/BadSigningKey",
		sim:  true,
		operations: []operation{
			{
				args: []string{
					"-type", "ping", "-terminal", terminalName,
					"-apiKey", "RIKLAPSMSMG2YII27N2NPAMCS5",
					"-bearerToken", "RIKLAPSMSMG2YII27N2NPAMCS5",
					"-signingKey", "4b556bc4e73ffc86fc5f8bfbba1598e7a8cd91f44fd7072d070c92fae7f48cd",
				},
				expect: blockchyp.Acknowledgement{
					Success: false,
					Error:   "Malformed Signing Key",
				},
			},
		},
	},
}
