# BlockChyp Go SDK

[![Build Status](https://circleci.com/gh/blockchyp/blockchyp-go/tree/master.svg?style=shield)](https://circleci.com/gh/blockchyp/blockchyp-go/tree/master)
[![Release](https://img.shields.io/github/release/blockchyp/blockchyp-go/all.svg?style=shield)](https://github.com/blockchyp/blockchyp-go/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/blockchyp/blockchyp-go)](https://goreportcard.com/report/github.com/blockchyp/blockchyp-go)
[![GoDoc](https://godoc.org/github.com/blockchyp/blockchyp-go?status.svg)](https://godoc.org/github.com/blockchyp/blockchyp-go)
[![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/blockchyp/blockchyp-go/blob/master/LICENSE)

This is the Go SDK for BlockChyp. Like all BlockChyp SDKs, it provides a full
Go client for the BlockChyp gateway and BlockChyp payment terminals.

This project also contains a command line interface for Windows, Linux, and
Mac OS developers working in languages or on platforms for which BlockChyp doesn't
currently provide a supported SDK.

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

Running your first terminal transaction is easy. Make sure you have a BlockChyp
terminal, activate it, and generate a set of API keys.

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
  req.Test = true
	req.TerminalName = "Test Terminal"
	req.Amount = "55.00"

	response, err := client.Charge(req)

	if err != nil {
		log.Fatal(err)
	}

	if response.Approved {
		fmt.Println("Approved")
    fmt.Println(response.AuthCode)
    fmt.Println(response.AuthorizedAmount)
    fmt.Println(response.ReceiptSuggestions.AID)
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

The response contains all the information you'll need to complete processing
a transaction. Of particular importance is the ReceiptSuggestions struct, which
contains all the fields that are required or recommended for PCI or EMV compliance.



## The Rest APIs

All BlockChyp SDKs provide a convenient way of accessing the BlockChyp REST APIs.
You can checkout the REST API documentation via the links below.

[Terminal REST API Docs](https://docs.blockchyp.com/rest-api/terminal/index.html)

[Gateway REST API Docs](https://docs.blockchyp.com/rest-api/gateway/index.html)

## Other SDKs

BlockChyp has officially supported SDKs for eight different development platforms and counting.
Here's the full list with links to their GitHub repositories.

[Go SDK](https://github.com/blockchyp/blockchyp-go)

[Node.js/JavaScript SDK](https://github.com/blockchyp/blockchyp-js)

[Java SDK](https://github.com/blockchyp/blockchyp-java)

[.net/C# SDK](https://github.com/blockchyp/blockchyp-csharp)

[Ruby SDK](https://github.com/blockchyp/blockchyp-ruby)

[PHP SDK](https://github.com/blockchyp/blockchyp-php)

[Python SDK](https://github.com/blockchyp/blockchyp-python)

[iOS (Objective-C/Swift) SDK](https://github.com/blockchyp/blockchyp-ios)

## Getting a Developer Kit

In order to test your integration with real terminals, you'll need a BlockChyp
Developer Kit. Our kits include a fully functioning payment terminal with
test pin encryption keys. Every kit includes a comprehensive set of test
cards with test cards for every major card brand and entry method, including
Contactless and Contact EMV and mag stripe cards. Each kit also includes
test gift cards for our blockchain gift card system.

Access to BlockChyp's developer program is currently invite only, but you
can request an invitation by contacting our engineering team at **nerds@blockchyp.com**.

You can also view a number of long form demos and learn more about us on our [YouTube Channel](https://www.youtube.com/channel/UCE-iIVlJic_XArs_U65ZcJg).

## Transaction Code Examples

You don't want to read words. You want examples. Here's a quick rundown of the
stuff you can do with the BlockChyp Go SDK and a few basic examples.

#### Charge

Executes a standard direct preauth and capture.

```go
package main

import (
    "fmt"
    "log"

    blockchyp "github.com/blockchyp/blockchyp-go"
)

func chargeExample() {
    // sample credentials
    creds := blockchyp.APICredentials{
        APIKey:      "ZDSMMZLGRPBPRTJUBTAFBYZ33Q",
        BearerToken: "ZLBW5NR4U5PKD5PNP3ZP3OZS5U",
        SigningKey:  "9c6a5e8e763df1c9256e3d72bd7f53dfbd07312938131c75b3bfd254da787947",
    }

    // instantiate the client
    client := blockchyp.NewClient(creds)

    // setup request object
    request := blockchyp.AuthorizationRequest{
        Test:         true,
        TerminalName: "Test Terminal",
        Amount:       "55.00",
    }

    response, err := client.Charge(request)

    if err != nil {
        log.Fatal(err)
    }

    //process the result
    if response.Approved {
        fmt.Println("Approved")
    }

    fmt.Printf("Response: %+v\n", response)
}

```

#### Preauthorization

Executes a preauthorization intended to be captured later.

```go
package main

import (
    "fmt"
    "log"

    blockchyp "github.com/blockchyp/blockchyp-go"
)

func preauthExample() {
    // sample credentials
    creds := blockchyp.APICredentials{
        APIKey:      "ZDSMMZLGRPBPRTJUBTAFBYZ33Q",
        BearerToken: "ZLBW5NR4U5PKD5PNP3ZP3OZS5U",
        SigningKey:  "9c6a5e8e763df1c9256e3d72bd7f53dfbd07312938131c75b3bfd254da787947",
    }

    // instantiate the client
    client := blockchyp.NewClient(creds)

    // setup request object
    request := blockchyp.AuthorizationRequest{
        Test:         true,
        TerminalName: "Test Terminal",
        Amount:       "27.00",
    }

    response, err := client.Preauth(request)

    if err != nil {
        log.Fatal(err)
    }

    //process the result
    if response.Approved {
        fmt.Println("Approved")
    }

    fmt.Printf("Response: %+v\n", response)
}

```

#### Terminal Ping

Tests connectivity with a payment terminal.

```go
package main

import (
    "fmt"
    "log"

    blockchyp "github.com/blockchyp/blockchyp-go"
)

func pingExample() {
    // sample credentials
    creds := blockchyp.APICredentials{
        APIKey:      "ZDSMMZLGRPBPRTJUBTAFBYZ33Q",
        BearerToken: "ZLBW5NR4U5PKD5PNP3ZP3OZS5U",
        SigningKey:  "9c6a5e8e763df1c9256e3d72bd7f53dfbd07312938131c75b3bfd254da787947",
    }

    // instantiate the client
    client := blockchyp.NewClient(creds)

    // setup request object
    request := blockchyp.PingRequest{
        TerminalName: "Test Terminal",
    }

    response, err := client.Ping(request)

    if err != nil {
        log.Fatal(err)
    }

    //process the result
    if response.Success {
        fmt.Println("Success")
    }

    fmt.Printf("Response: %+v\n", response)
}

```

#### Balance

Checks the remaining balance on a payment method.

```go
package main

import (
    "fmt"
    "log"

    blockchyp "github.com/blockchyp/blockchyp-go"
)

func balanceExample() {
    // sample credentials
    creds := blockchyp.APICredentials{
        APIKey:      "ZDSMMZLGRPBPRTJUBTAFBYZ33Q",
        BearerToken: "ZLBW5NR4U5PKD5PNP3ZP3OZS5U",
        SigningKey:  "9c6a5e8e763df1c9256e3d72bd7f53dfbd07312938131c75b3bfd254da787947",
    }

    // instantiate the client
    client := blockchyp.NewClient(creds)

    // setup request object
    request := blockchyp.BalanceRequest{
        Test:         true,
        TerminalName: "Test Terminal",
        CardType:     blockchyp.CardTypeEBT,
    }

    response, err := client.Balance(request)

    if err != nil {
        log.Fatal(err)
    }

    //process the result
    if response.Success {
        fmt.Println("Success")
    }

    fmt.Printf("Response: %+v\n", response)
}

```

#### Terminal Clear

Clears the line item display and any in progress transaction.

```go
package main

import (
    "fmt"
    "log"

    blockchyp "github.com/blockchyp/blockchyp-go"
)

func clearExample() {
    // sample credentials
    creds := blockchyp.APICredentials{
        APIKey:      "ZDSMMZLGRPBPRTJUBTAFBYZ33Q",
        BearerToken: "ZLBW5NR4U5PKD5PNP3ZP3OZS5U",
        SigningKey:  "9c6a5e8e763df1c9256e3d72bd7f53dfbd07312938131c75b3bfd254da787947",
    }

    // instantiate the client
    client := blockchyp.NewClient(creds)

    // setup request object
    request := blockchyp.ClearTerminalRequest{
        Test:         true,
        TerminalName: "Test Terminal",
    }

    response, err := client.Clear(request)

    if err != nil {
        log.Fatal(err)
    }

    //process the result
    if response.Success {
        fmt.Println("Success")
    }

    fmt.Printf("Response: %+v\n", response)
}

```

#### Terms & Conditions Capture

Prompts the user to accept terms and conditions.

```go
package main

import (
    "fmt"
    "log"

    blockchyp "github.com/blockchyp/blockchyp-go"
)

func termsAndConditionsExample() {
    // sample credentials
    creds := blockchyp.APICredentials{
        APIKey:      "ZDSMMZLGRPBPRTJUBTAFBYZ33Q",
        BearerToken: "ZLBW5NR4U5PKD5PNP3ZP3OZS5U",
        SigningKey:  "9c6a5e8e763df1c9256e3d72bd7f53dfbd07312938131c75b3bfd254da787947",
    }

    // instantiate the client
    client := blockchyp.NewClient(creds)

    // setup request object
    request := blockchyp.TermsAndConditionsRequest{
        Test:         true,
        TerminalName: "Test Terminal",

        // Alias for a Terms and Conditions template configured in the BlockChyp dashboard.
        TCAlias: "hippa",

        // Name of the contract or document if not using an alias.
        TCName: "HIPPA Disclosure",

        // Full text of the contract or disclosure if not using an alias.
        TCContent: "Full contract text",

        // file format for the signature image.
        SigFormat: blockchyp.SignatureFormatPNG,

        // width of the signature image in pixels.
        SigWidth: 200,

        // Whether or not a signature is required. Defaults to true.
        SigRequired: true,
    }

    response, err := client.TermsAndConditions(request)

    if err != nil {
        log.Fatal(err)
    }

    //process the result
    if response.Success {
        fmt.Println("Success")
    }

    fmt.Printf("Response: %+v\n", response)
}

```

#### Update Transaction Display

Appends items to an existing transaction display Subtotal, Tax, and Total are
overwritten by the request. Items with the same description are combined into
groups.

```go
package main

import (
    "fmt"
    "log"

    blockchyp "github.com/blockchyp/blockchyp-go"
)

func updateTransactionDisplayExample() {
    // sample credentials
    creds := blockchyp.APICredentials{
        APIKey:      "ZDSMMZLGRPBPRTJUBTAFBYZ33Q",
        BearerToken: "ZLBW5NR4U5PKD5PNP3ZP3OZS5U",
        SigningKey:  "9c6a5e8e763df1c9256e3d72bd7f53dfbd07312938131c75b3bfd254da787947",
    }

    // instantiate the client
    client := blockchyp.NewClient(creds)

    // setup request object
    request := blockchyp.TransactionDisplayRequest{
        Test:         true,
        TerminalName: "Test Terminal",
        Transaction: &blockchyp.TransactionDisplayTransaction{
            Subtotal: "60.00",
            Tax:      "5.00",
            Total:    "65.00",
            Items: []*blockchyp.TransactionDisplayItem{
                &blockchyp.TransactionDisplayItem{
                    Description: "Leki Trekking Poles",
                    Price:       "35.00",
                    Quantity:    2,
                    Extended:    "70.00",
                    Discounts: []*blockchyp.TransactionDisplayDiscount{
                        &blockchyp.TransactionDisplayDiscount{
                            Description: "memberDiscount",
                            Amount:      "10.00",
                        },
                    },
                },
            },
        },
    }

    response, err := client.UpdateTransactionDisplay(request)

    if err != nil {
        log.Fatal(err)
    }

    //process the result
    if response.Success {
        fmt.Println("Success")
    }

    fmt.Printf("Response: %+v\n", response)
}

```

#### New Transaction Display

Displays a new transaction on the terminal.

```go
package main

import (
    "fmt"
    "log"

    blockchyp "github.com/blockchyp/blockchyp-go"
)

func newTransactionDisplayExample() {
    // sample credentials
    creds := blockchyp.APICredentials{
        APIKey:      "ZDSMMZLGRPBPRTJUBTAFBYZ33Q",
        BearerToken: "ZLBW5NR4U5PKD5PNP3ZP3OZS5U",
        SigningKey:  "9c6a5e8e763df1c9256e3d72bd7f53dfbd07312938131c75b3bfd254da787947",
    }

    // instantiate the client
    client := blockchyp.NewClient(creds)

    // setup request object
    request := blockchyp.TransactionDisplayRequest{
        Test:         true,
        TerminalName: "Test Terminal",
        Transaction: &blockchyp.TransactionDisplayTransaction{
            Subtotal: "60.00",
            Tax:      "5.00",
            Total:    "65.00",
            Items: []*blockchyp.TransactionDisplayItem{
                &blockchyp.TransactionDisplayItem{
                    Description: "Leki Trekking Poles",
                    Price:       "35.00",
                    Quantity:    2,
                    Extended:    "70.00",
                    Discounts: []*blockchyp.TransactionDisplayDiscount{
                        &blockchyp.TransactionDisplayDiscount{
                            Description: "memberDiscount",
                            Amount:      "10.00",
                        },
                    },
                },
            },
        },
    }

    response, err := client.NewTransactionDisplay(request)

    if err != nil {
        log.Fatal(err)
    }

    //process the result
    if response.Success {
        fmt.Println("Success")
    }

    fmt.Printf("Response: %+v\n", response)
}

```

#### Text Prompt

Asks the consumer text based question.

```go
package main

import (
    "fmt"
    "log"

    blockchyp "github.com/blockchyp/blockchyp-go"
)

func textPromptExample() {
    // sample credentials
    creds := blockchyp.APICredentials{
        APIKey:      "ZDSMMZLGRPBPRTJUBTAFBYZ33Q",
        BearerToken: "ZLBW5NR4U5PKD5PNP3ZP3OZS5U",
        SigningKey:  "9c6a5e8e763df1c9256e3d72bd7f53dfbd07312938131c75b3bfd254da787947",
    }

    // instantiate the client
    client := blockchyp.NewClient(creds)

    // setup request object
    request := blockchyp.TextPromptRequest{
        Test:         true,
        TerminalName: "Test Terminal",

        // Type of prompt. Can be 'email', 'phone', 'customer-number', or 'rewards-number'.
        PromptType: blockchyp.PromptTypeEmail,
    }

    response, err := client.TextPrompt(request)

    if err != nil {
        log.Fatal(err)
    }

    //process the result
    if response.Success {
        fmt.Println("Success")
    }

    fmt.Printf("Response: %+v\n", response)
}

```

#### Boolean Prompt

Asks the consumer a yes/no question.

```go
package main

import (
    "fmt"
    "log"

    blockchyp "github.com/blockchyp/blockchyp-go"
)

func booleanPromptExample() {
    // sample credentials
    creds := blockchyp.APICredentials{
        APIKey:      "ZDSMMZLGRPBPRTJUBTAFBYZ33Q",
        BearerToken: "ZLBW5NR4U5PKD5PNP3ZP3OZS5U",
        SigningKey:  "9c6a5e8e763df1c9256e3d72bd7f53dfbd07312938131c75b3bfd254da787947",
    }

    // instantiate the client
    client := blockchyp.NewClient(creds)

    // setup request object
    request := blockchyp.BooleanPromptRequest{
        Test:         true,
        TerminalName: "Test Terminal",
        Prompt:       "Would you like to become a member?",
        YesCaption:   "Yes",
        NoCaption:    "No",
    }

    response, err := client.BooleanPrompt(request)

    if err != nil {
        log.Fatal(err)
    }

    //process the result
    if response.Success {
        fmt.Println("Success")
    }

    fmt.Printf("Response: %+v\n", response)
}

```

#### Display Message

Displays a short message on the terminal.

```go
package main

import (
    "fmt"
    "log"

    blockchyp "github.com/blockchyp/blockchyp-go"
)

func messageExample() {
    // sample credentials
    creds := blockchyp.APICredentials{
        APIKey:      "ZDSMMZLGRPBPRTJUBTAFBYZ33Q",
        BearerToken: "ZLBW5NR4U5PKD5PNP3ZP3OZS5U",
        SigningKey:  "9c6a5e8e763df1c9256e3d72bd7f53dfbd07312938131c75b3bfd254da787947",
    }

    // instantiate the client
    client := blockchyp.NewClient(creds)

    // setup request object
    request := blockchyp.MessageRequest{
        Test:         true,
        TerminalName: "Test Terminal",
        Message:      "Thank you for your business.",
    }

    response, err := client.Message(request)

    if err != nil {
        log.Fatal(err)
    }

    //process the result
    if response.Success {
        fmt.Println("Success")
    }

    fmt.Printf("Response: %+v\n", response)
}

```

#### Refund

Executes a refund.

```go
package main

import (
    "fmt"
    "log"

    blockchyp "github.com/blockchyp/blockchyp-go"
)

func refundExample() {
    // sample credentials
    creds := blockchyp.APICredentials{
        APIKey:      "ZDSMMZLGRPBPRTJUBTAFBYZ33Q",
        BearerToken: "ZLBW5NR4U5PKD5PNP3ZP3OZS5U",
        SigningKey:  "9c6a5e8e763df1c9256e3d72bd7f53dfbd07312938131c75b3bfd254da787947",
    }

    // instantiate the client
    client := blockchyp.NewClient(creds)

    // setup request object
    request := blockchyp.RefundRequest{
        TerminalName:  "Test Terminal",
        TransactionID: "<PREVIOUS TRANSACTION ID>",

        // Optional amount for partial refunds.
        Amount: "5.00",
    }

    response, err := client.Refund(request)

    if err != nil {
        log.Fatal(err)
    }

    //process the result
    if response.Approved {
        fmt.Println("Approved")
    }

    fmt.Printf("Response: %+v\n", response)
}

```

#### Enroll

Adds a new payment method to the token vault.

```go
package main

import (
    "fmt"
    "log"

    blockchyp "github.com/blockchyp/blockchyp-go"
)

func enrollExample() {
    // sample credentials
    creds := blockchyp.APICredentials{
        APIKey:      "ZDSMMZLGRPBPRTJUBTAFBYZ33Q",
        BearerToken: "ZLBW5NR4U5PKD5PNP3ZP3OZS5U",
        SigningKey:  "9c6a5e8e763df1c9256e3d72bd7f53dfbd07312938131c75b3bfd254da787947",
    }

    // instantiate the client
    client := blockchyp.NewClient(creds)

    // setup request object
    request := blockchyp.EnrollRequest{
        Test:         true,
        TerminalName: "Test Terminal",
    }

    response, err := client.Enroll(request)

    if err != nil {
        log.Fatal(err)
    }

    //process the result
    if response.Approved {
        fmt.Println("Approved")
    }

    fmt.Printf("Response: %+v\n", response)
}

```

#### Gift Card Activation

Activates or recharges a gift card.

```go
package main

import (
    "fmt"
    "log"

    blockchyp "github.com/blockchyp/blockchyp-go"
)

func giftActivateExample() {
    // sample credentials
    creds := blockchyp.APICredentials{
        APIKey:      "ZDSMMZLGRPBPRTJUBTAFBYZ33Q",
        BearerToken: "ZLBW5NR4U5PKD5PNP3ZP3OZS5U",
        SigningKey:  "9c6a5e8e763df1c9256e3d72bd7f53dfbd07312938131c75b3bfd254da787947",
    }

    // instantiate the client
    client := blockchyp.NewClient(creds)

    // setup request object
    request := blockchyp.GiftActivateRequest{
        Test:         true,
        TerminalName: "Test Terminal",
        Amount:       "50.00",
    }

    response, err := client.GiftActivate(request)

    if err != nil {
        log.Fatal(err)
    }

    //process the result
    if response.Approved {
        fmt.Println("Approved")
    }

    fmt.Printf("Response: %+v\n", response)
}

```

#### Time Out Reversal

Executes a manual time out reversal.

We love time out reversals. Don't be afraid to use them whenever a request to a
BlockChyp terminal times out. You have up to two minutes to reverse any
transaction. The only caveat is that you must assign transactionRef values when
you build the original request. Otherwise, we have no real way of knowing which
transaction you're trying to reverse because we may not have assigned it an id
yet. And if we did assign it an id, you wouldn't know what it is because your
request to the terminal timed out before you got a response.

```go
package main

import (
    "fmt"
    "log"

    blockchyp "github.com/blockchyp/blockchyp-go"
)

func reverseExample() {
    // sample credentials
    creds := blockchyp.APICredentials{
        APIKey:      "ZDSMMZLGRPBPRTJUBTAFBYZ33Q",
        BearerToken: "ZLBW5NR4U5PKD5PNP3ZP3OZS5U",
        SigningKey:  "9c6a5e8e763df1c9256e3d72bd7f53dfbd07312938131c75b3bfd254da787947",
    }

    // instantiate the client
    client := blockchyp.NewClient(creds)

    // setup request object
    request := blockchyp.AuthorizationRequest{
        TerminalName:   "Test Terminal",
        TransactionRef: "<LAST TRANSACTION REF>",
    }

    response, err := client.Reverse(request)

    if err != nil {
        log.Fatal(err)
    }

    //process the result
    if response.Approved {
        fmt.Println("Approved")
    }

    fmt.Printf("Response: %+v\n", response)
}

```

#### Capture Preauthorization

Captures a preauthorization.

```go
package main

import (
    "fmt"
    "log"

    blockchyp "github.com/blockchyp/blockchyp-go"
)

func captureExample() {
    // sample credentials
    creds := blockchyp.APICredentials{
        APIKey:      "ZDSMMZLGRPBPRTJUBTAFBYZ33Q",
        BearerToken: "ZLBW5NR4U5PKD5PNP3ZP3OZS5U",
        SigningKey:  "9c6a5e8e763df1c9256e3d72bd7f53dfbd07312938131c75b3bfd254da787947",
    }

    // instantiate the client
    client := blockchyp.NewClient(creds)

    // setup request object
    request := blockchyp.CaptureRequest{
        Test:          true,
        TransactionID: "<PREAUTH TRANSACTION ID>",
    }

    response, err := client.Capture(request)

    if err != nil {
        log.Fatal(err)
    }

    //process the result
    if response.Approved {
        fmt.Println("Approved")
    }

    fmt.Printf("Response: %+v\n", response)
}

```

#### Close Batch

Closes the current credit card batch.

```go
package main

import (
    "fmt"
    "log"

    blockchyp "github.com/blockchyp/blockchyp-go"
)

func closeBatchExample() {
    // sample credentials
    creds := blockchyp.APICredentials{
        APIKey:      "ZDSMMZLGRPBPRTJUBTAFBYZ33Q",
        BearerToken: "ZLBW5NR4U5PKD5PNP3ZP3OZS5U",
        SigningKey:  "9c6a5e8e763df1c9256e3d72bd7f53dfbd07312938131c75b3bfd254da787947",
    }

    // instantiate the client
    client := blockchyp.NewClient(creds)

    // setup request object
    request := blockchyp.CloseBatchRequest{
        Test: true,
    }

    response, err := client.CloseBatch(request)

    if err != nil {
        log.Fatal(err)
    }

    //process the result
    if response.Success {
        fmt.Println("Success")
    }

    fmt.Printf("Response: %+v\n", response)
}

```

#### Void Transaction

Discards a previous preauth transaction.

```go
package main

import (
    "fmt"
    "log"

    blockchyp "github.com/blockchyp/blockchyp-go"
)

func voidExample() {
    // sample credentials
    creds := blockchyp.APICredentials{
        APIKey:      "ZDSMMZLGRPBPRTJUBTAFBYZ33Q",
        BearerToken: "ZLBW5NR4U5PKD5PNP3ZP3OZS5U",
        SigningKey:  "9c6a5e8e763df1c9256e3d72bd7f53dfbd07312938131c75b3bfd254da787947",
    }

    // instantiate the client
    client := blockchyp.NewClient(creds)

    // setup request object
    request := blockchyp.VoidRequest{
        Test:          true,
        TransactionID: "<PREVIOUS TRANSACTION ID>",
    }

    response, err := client.Void(request)

    if err != nil {
        log.Fatal(err)
    }

    //process the result
    if response.Approved {
        fmt.Println("Approved")
    }

    fmt.Printf("Response: %+v\n", response)
}

```

#### Terminal Status

Returns the current status of a terminal.

```go
package main

import (
    "fmt"
    "log"

    blockchyp "github.com/blockchyp/blockchyp-go"
)

func terminalStatusExample() {
    // sample credentials
    creds := blockchyp.APICredentials{
        APIKey:      "ZDSMMZLGRPBPRTJUBTAFBYZ33Q",
        BearerToken: "ZLBW5NR4U5PKD5PNP3ZP3OZS5U",
        SigningKey:  "9c6a5e8e763df1c9256e3d72bd7f53dfbd07312938131c75b3bfd254da787947",
    }

    // instantiate the client
    client := blockchyp.NewClient(creds)

    // setup request object
    request := blockchyp.TerminalStatusRequest{
        TerminalName: "Test Terminal",
    }

    response, err := client.TerminalStatus(request)

    if err != nil {
        log.Fatal(err)
    }

    //process the result
    if response.Success {
        fmt.Println("Success")
    }

    fmt.Printf("Response: %+v\n", response)
}

```

## Running Integration Tests

If you'd like to run the integration tests, create a new file on your system
called `sdk-itest-config.json` with the API credentials you'll be using as
shown in the example below.

```
{
 "gatewayHost": "https://api.blockchyp.com",
 "testGatewayHost": "https://test.blockchyp.com",
 "apiKey": "PZZNEFK7HFULCB3HTLA7HRQDJU",
 "bearerToken": "QUJCHIKNXOMSPGQ4QLT2UJX5DI",
 "signingKey": "f88a72d8bc0965f193abc7006bbffa240663c10e4d1dc3ba2f81e0ca10d359f5"
}
```

This file can be located in a few different places, but is usually located
at `<USER_HOME>/.config/blockchyp/sdk-itest-config.json`. All BlockChyp SDKs
use the same configuration file.

To run the integration test suite via `make`, type the following command:

`make integration`


## Contributions

BlockChyp welcomes contributions from the open source community, but bear in mind
that this repository has been generated by our internal SDK Generator tool. If
we choose to accept a PR or contribution, your code will be moved into our SDK
Generator project, which is a private repository.

## License

Copyright BlockChyp, Inc., 2019

Distributed under the terms of the [MIT] license, blockchyp-go is free and open source software.

[MIT]: https://github.com/blockchyp/blockchyp-go/blob/master/LICENSE
