// +build regression

package regression

import (
	"testing"

	"github.com/blockchyp/blockchyp-go"
)

func TestPing(t *testing.T) {
	if acquirerMode {
		t.Skip("skipped for acquirer test run")
	}

	tests := map[string]struct {
		instructions string
		args         []string
		assert       blockchyp.Acknowledgement
	}{
		"Success": {
			args: []string{
				"-type", "ping", "-terminal", terminalName,
			},
			assert: blockchyp.Acknowledgement{
				Success: true,
			},
		},
		"Failure": {
			args: []string{
				"-type", "ping", "-terminal", "Unknown Terminal",
			},
			assert: blockchyp.Acknowledgement{
				Success: false,
				Error:   "unknown terminal",
			},
		},
		"BadCreds": {
			args: []string{
				"-type", "ping", "-terminal", terminalName,
				"-apiKey", "RIKLAPSMSMG2YII27N2NPAMCS5",
				"-bearerToken", "RIKLAPSMSMG2YII27N2NPAMCS5",
				"-signingKey", "4b556bc4e73ffc86fc5f8bfbba1598e7a8cd91f44fd7072d070c92fae7f48cd9",
			},
			assert: blockchyp.Acknowledgement{
				Success: false,
				Error:   "Access Denied",
			},
		},
		"BadSigningKey": {
			args: []string{
				"-type", "ping", "-terminal", terminalName,
				"-apiKey", "RIKLAPSMSMG2YII27N2NPAMCS5",
				"-bearerToken", "RIKLAPSMSMG2YII27N2NPAMCS5",
				"-signingKey", "4b556bc4e73ffc86fc5f8bfbba1598e7a8cd91f44fd7072d070c92fae7f48cd",
			},
			assert: blockchyp.Acknowledgement{
				Success: false,
				Error:   "Malformed Signing Key",
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			cli := newCLI(t)

			setup(t, test.instructions, false)

			cli.run(test.args, test.assert)
		})
	}
}
