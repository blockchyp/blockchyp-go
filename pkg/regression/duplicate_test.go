// +build regression

package regression

import (
	"testing"

	"github.com/blockchyp/blockchyp-go"
)

func TestDuplicate(t *testing.T) {
	cli := newCLI(t)
	setup(t, `Insert an EMV test card when prompted.

Leave it in the terminal until the test completes.`, true)

	args := []string{
		"-type", "charge", "-terminal", "Test Terminal",
		"-test",
		"-amount", randomAmount(),
	}

	expect0 := blockchyp.AuthorizationResponse{
		Success:  true,
		Approved: true,
		Test:     true,
	}

	cli.run(args, expect0)

	expect1 := blockchyp.AuthorizationResponse{
		Success:             true,
		Approved:            false,
		Test:                true,
		ResponseDescription: "Duplicate Transaction",
	}

	cli.run(args, expect1)
}
