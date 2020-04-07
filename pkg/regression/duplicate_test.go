// +build regression

package regression

import (
	"testing"

	"github.com/blockchyp/blockchyp-go"
)

func TestDuplicate(t *testing.T) {
	cli := newCLI(t)
	setup(t, `Insert a valid test card when prompted.

Insert the same test card when prompted again.`, true)

	args := []string{
		"-type", "charge", "-terminal", "Test Terminal",
		"-test",
		"-amount", "51.00",
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
