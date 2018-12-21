# BlockChyp Go SDK

[![Travis CI](https://img.shields.io/travis/blockchyp/blockchyp-go.svg?style=shield)](https://travis-ci.org/blockchyp/blockchyp-go)
[![Release](https://img.shields.io/github/release/blockchyp/blockchyp-go/all.svg?style=shield)](https://github.com/blockchyp/blockchyp-go/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/blockchyp/blockchyp-go)](https://goreportcard.com/report/github.com/blockchyp/blockchyp-go)
[![GoDoc](https://godoc.org/github.com/blockchyp/blockchyp-go?status.svg)](https://godoc.org/github.com/blockchyp/blockchyp-go)

This is the reference SDK implementation for BlockChyp maintained by the BlockChyp engineering team.

It's based on the [BlockChyp SDK Developers Guide](https://docs.blockchyp.com/sdk-guide/index.html).

BlockChyp is still pre-release and developer access is by invitation only.  Godocs are coming soon.

This project contains a full native Go client for BlockChyp along with a CLI for Windows,
Linux, and Mac OS developers.

## Command Line Interface

In addition to the standard Go SDK, the Makefile includes special targets for
Windows and Linux command line binaries.

These binaries are intended for unique situations where using an SDK or doing
a direct REST integration aren't practical.

Check out the [CLI Reference](docs/cli.md) for more information.

## Go Installation

For Go developers, you can install BlockChyp in the usual way with `go get`.

```
go get github.com/blockchyp/blockchyp-go
```

## A Simple Example

The following code snippet shows a minimal example of a charge transaction.

```
package main

import (
	"encoding/json"
	"fmt"
	"log"

	blockchyp "github.com/blockchyp/blockchyp-go"
)

func main() {

	creds := blockchyp.APICredentials{
		APIKey:      "ZDSMMZLGRPBPRTJUBTAFBYZ33Q",
		BearerToken: "ZLBW5NR4U5PKD5PNP3ZP3OZS5U",
		SigningKey:  "9c6a5e8e763df1c9256e3d72bd7f53dfbd07312938131c75b3bfd254da787947",
	}

	client := blockchyp.NewClient(creds)

	req := blockchyp.AuthorizationRequest{}
	req.TerminalName = "Test Terminal"
	req.Amount = "55.00"

	response, err := client.Charge(req)

	if err != nil {
		log.Fatal(err)
	}

	if response.Approved {
		fmt.Println("Approved")
	} else {
		fmt.Println(response.ResponseDescription)
	}

	b, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(b))

}
```
