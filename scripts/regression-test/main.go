package main

import (
	"os"

	"github.com/blockchyp/blockchyp-go/v2/pkg/regression"
)

func main() {
	runner := regression.NewTestRunner()

	if err := runner.Run(); err != nil {
		os.Exit(1)
	}
}
