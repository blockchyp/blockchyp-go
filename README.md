# BlockChyp Go SDK

[![CircleCI](https://circleci.com/gh/blockchyp/blockchyp-go/tree/master.svg?style=shield)](https://circleci.com/gh/blockchyp/blockchyp-go/tree/master)
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

## Transaction Code Examples

You don't want to read words. You want examples. Here's a quick rundown of the
stuff you can do with the BlockChyp Go SDK and a few basic examples.
#### Charge

Executes a standard direct preauth and capture.

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"

    blockchyp "github.com/blockchyp/blockchyp-go"
)

func main() {

    // sample credentials
    creds := blockchyp.APICredentials{
        APIKey:      "ZDSMMZLGRPBPRTJUBTAFBYZ33Q",
        BearerToken: "ZLBW5NR4U5PKD5PNP3ZP3OZS5U",
        SigningKey:  "9c6a5e8e763df1c9256e3d72bd7f53dfbd07312938131c75b3bfd254da787947",
    }

    // instantiate the client
    client := blockchyp.NewClient(creds)

    // setup request object
    request := blockchyp.AuthorizationRequest{}
    request.Test = true
    request.TerminalName = "Test Terminal"
    request.Amount = "55.00"

    response, err := client.Charge(request)

    // run the transaction
    response, err := client.Charge(req)

    if err != nil {
        log.Fatal(err)
    }

    //process the result
    if response.Approved {
        fmt.Println("Approved")
    }

    fmt.Println(response.AuthCode)
    fmt.Println(response.AuthorizedAmount)
}

```
#### Preauthorization

Executes a preauthorization intended to be captured later.

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"

    blockchyp "github.com/blockchyp/blockchyp-go"
)

func main() {

    // sample credentials
    creds := blockchyp.APICredentials{
        APIKey:      "ZDSMMZLGRPBPRTJUBTAFBYZ33Q",
        BearerToken: "ZLBW5NR4U5PKD5PNP3ZP3OZS5U",
        SigningKey:  "9c6a5e8e763df1c9256e3d72bd7f53dfbd07312938131c75b3bfd254da787947",
    }

    // instantiate the client
    client := blockchyp.NewClient(creds)

    // setup request object
    request := blockchyp.AuthorizationRequest{}
    request.Test = true
    request.TerminalName = "Test Terminal"
    request.Amount = "27.00"

    response, err := client.Preauth(request)

    // run the transaction
    response, err := client.Preauth(req)

    if err != nil {
        log.Fatal(err)
    }

    //process the result
    if response.Approved {
        fmt.Println("Approved")
    }

    fmt.Println(response.AuthCode)
    fmt.Println(response.AuthorizedAmount)
}

```
#### Terminal Ping

Tests connectivity with a payment terminal.

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"

    blockchyp "github.com/blockchyp/blockchyp-go"
)

func main() {

    // sample credentials
    creds := blockchyp.APICredentials{
        APIKey:      "ZDSMMZLGRPBPRTJUBTAFBYZ33Q",
        BearerToken: "ZLBW5NR4U5PKD5PNP3ZP3OZS5U",
        SigningKey:  "9c6a5e8e763df1c9256e3d72bd7f53dfbd07312938131c75b3bfd254da787947",
    }

    // instantiate the client
    client := blockchyp.NewClient(creds)

    // setup request object
    request := blockchyp.PingRequest{}
    request.TerminalName = "Test Terminal"

    response, err := client.Ping(request)

    // run the transaction
    response, err := client.Ping(req)

    if err != nil {
        log.Fatal(err)
    }

    //process the result
    if response.Success {
        fmt.Println("Success")
    }

}

```
#### Balance

Checks the remaining balance on a payment method.

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"

    blockchyp "github.com/blockchyp/blockchyp-go"
)

func main() {

    // sample credentials
    creds := blockchyp.APICredentials{
        APIKey:      "ZDSMMZLGRPBPRTJUBTAFBYZ33Q",
        BearerToken: "ZLBW5NR4U5PKD5PNP3ZP3OZS5U",
        SigningKey:  "9c6a5e8e763df1c9256e3d72bd7f53dfbd07312938131c75b3bfd254da787947",
    }

    // instantiate the client
    client := blockchyp.NewClient(creds)

    // setup request object
    request := blockchyp.BalanceRequest{}
    request.Test = true
    request.TerminalName = "Test Terminal"
    request.CardType = 2

    response, err := client.Balance(request)

    // run the transaction
    response, err := client.Balance(req)

    if err != nil {
        log.Fatal(err)
    }

    //process the result
    if response.Success {
        fmt.Println("Success")
    }

}

```
#### Terminal Clear

Clears the line item display and any in progress transaction.

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"

    blockchyp "github.com/blockchyp/blockchyp-go"
)

func main() {

    // sample credentials
    creds := blockchyp.APICredentials{
        APIKey:      "ZDSMMZLGRPBPRTJUBTAFBYZ33Q",
        BearerToken: "ZLBW5NR4U5PKD5PNP3ZP3OZS5U",
        SigningKey:  "9c6a5e8e763df1c9256e3d72bd7f53dfbd07312938131c75b3bfd254da787947",
    }

    // instantiate the client
    client := blockchyp.NewClient(creds)

    // setup request object
    request := blockchyp.ClearTerminalRequest{}
    request.Test = true
    request.TerminalName = "Test Terminal"

    response, err := client.Clear(request)

    // run the transaction
    response, err := client.Clear(req)

    if err != nil {
        log.Fatal(err)
    }

    //process the result
    if response.Success {
        fmt.Println("Success")
    }

}

```
#### Terms & Conditions Capture

Prompts the user to accept terms and conditions.

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"

    blockchyp "github.com/blockchyp/blockchyp-go"
)

func main() {

    // sample credentials
    creds := blockchyp.APICredentials{
        APIKey:      "ZDSMMZLGRPBPRTJUBTAFBYZ33Q",
        BearerToken: "ZLBW5NR4U5PKD5PNP3ZP3OZS5U",
        SigningKey:  "9c6a5e8e763df1c9256e3d72bd7f53dfbd07312938131c75b3bfd254da787947",
    }

    // instantiate the client
    client := blockchyp.NewClient(creds)

    // setup request object
    request := blockchyp.TermsAndConditionsRequest{}
    request.Test = true
    request.TerminalName = "Test Terminal"
    request.TCAlias = "hippa"                // Alias for a T&C template configured in blockchyp.
    request.TCName = "HIPPA Disclosure"      // Name of the contract or document if not using an alias.
    request.TCContent = "Full contract text" // Full text of the contract or disclosure if not using an alias.
    request.SigFormat = "png"                // file format for the signature image, if desired can be PNG or JPG.
    request.SigWidth = 200                   // width of the signature image in pixels.
    request.SigRequired = true               // Whether or not a signature is required. Defaults to true.

    response, err := client.TermsAndConditions(request)

    // run the transaction
    response, err := client.TermsAndConditions(req)

    if err != nil {
        log.Fatal(err)
    }

    //process the result
    if response.Success {
        fmt.Println("Success")
    }

    fmt.Println(response.Sig)
    fmt.Println(response.SigFile)
}

```
#### Update Transaction Display

Appends items to an existing transaction display Subtotal, Tax, and Total are
overwritten by the request. Items with the same description are combined into
groups.

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"

    blockchyp "github.com/blockchyp/blockchyp-go"
)

func main() {

    // sample credentials
    creds := blockchyp.APICredentials{
        APIKey:      "ZDSMMZLGRPBPRTJUBTAFBYZ33Q",
        BearerToken: "ZLBW5NR4U5PKD5PNP3ZP3OZS5U",
        SigningKey:  "9c6a5e8e763df1c9256e3d72bd7f53dfbd07312938131c75b3bfd254da787947",
    }

    // instantiate the client
    client := blockchyp.NewClient(creds)

    // setup request object
    request := blockchyp.TransactionDisplayRequest{}
    request.Test = true
    request.TerminalName = "Test Terminal"
    request.Transaction = &blockchyp.TransactionDisplayTransaction{
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
    }

    response, err := client.UpdateTransactionDisplay(request)

    // run the transaction
    response, err := client.UpdateTransactionDisplay(req)

    if err != nil {
        log.Fatal(err)
    }

    //process the result
    if response.Success {
        fmt.Println("Succeded")
    }

}

```
#### New Transaction Display

Displays a new transaction on the terminal.

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"

    blockchyp "github.com/blockchyp/blockchyp-go"
)

func main() {

    // sample credentials
    creds := blockchyp.APICredentials{
        APIKey:      "ZDSMMZLGRPBPRTJUBTAFBYZ33Q",
        BearerToken: "ZLBW5NR4U5PKD5PNP3ZP3OZS5U",
        SigningKey:  "9c6a5e8e763df1c9256e3d72bd7f53dfbd07312938131c75b3bfd254da787947",
    }

    // instantiate the client
    client := blockchyp.NewClient(creds)

    // setup request object
    request := blockchyp.TransactionDisplayRequest{}
    request.Test = true
    request.TerminalName = "Test Terminal"
    request.Transaction = &blockchyp.TransactionDisplayTransaction{
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
    }

    response, err := client.NewTransactionDisplay(request)

    // run the transaction
    response, err := client.NewTransactionDisplay(req)

    if err != nil {
        log.Fatal(err)
    }

    //process the result
    if response.Success {
        fmt.Println("Succeded")
    }

}

```
#### Text Prompt

Asks the consumer text based question.

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"

    blockchyp "github.com/blockchyp/blockchyp-go"
)

func main() {

    // sample credentials
    creds := blockchyp.APICredentials{
        APIKey:      "ZDSMMZLGRPBPRTJUBTAFBYZ33Q",
        BearerToken: "ZLBW5NR4U5PKD5PNP3ZP3OZS5U",
        SigningKey:  "9c6a5e8e763df1c9256e3d72bd7f53dfbd07312938131c75b3bfd254da787947",
    }

    // instantiate the client
    client := blockchyp.NewClient(creds)

    // setup request object
    request := blockchyp.TextPromptRequest{}
    request.Test = true
    request.TerminalName = "Test Terminal"
    request.PromptType = "email" // Type of prompt. Can be 'email', 'phone', 'customer-number', or 'rewards-number'.

    response, err := client.TextPrompt(request)

    // run the transaction
    response, err := client.TextPrompt(req)

    if err != nil {
        log.Fatal(err)
    }

    //process the result
    if response.Success {
        fmt.Println("Success")
    }

    fmt.Println(response.Response)
}

```
#### Boolean Prompt

Asks the consumer a yes/no question.

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"

    blockchyp "github.com/blockchyp/blockchyp-go"
)

func main() {

    // sample credentials
    creds := blockchyp.APICredentials{
        APIKey:      "ZDSMMZLGRPBPRTJUBTAFBYZ33Q",
        BearerToken: "ZLBW5NR4U5PKD5PNP3ZP3OZS5U",
        SigningKey:  "9c6a5e8e763df1c9256e3d72bd7f53dfbd07312938131c75b3bfd254da787947",
    }

    // instantiate the client
    client := blockchyp.NewClient(creds)

    // setup request object
    request := blockchyp.BooleanPromptRequest{}
    request.Test = true
    request.TerminalName = "Test Terminal"
    request.Prompt = "Would you like to become a member?"
    request.YesCaption = "Yes"
    request.NoCaption = "No"

    response, err := client.BooleanPrompt(request)

    // run the transaction
    response, err := client.BooleanPrompt(req)

    if err != nil {
        log.Fatal(err)
    }

    //process the result
    if response.Success {
        fmt.Println("Success")
    }

    fmt.Println(response.Response)
}

```
#### Display Message

Displays a short message on the terminal.

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"

    blockchyp "github.com/blockchyp/blockchyp-go"
)

func main() {

    // sample credentials
    creds := blockchyp.APICredentials{
        APIKey:      "ZDSMMZLGRPBPRTJUBTAFBYZ33Q",
        BearerToken: "ZLBW5NR4U5PKD5PNP3ZP3OZS5U",
        SigningKey:  "9c6a5e8e763df1c9256e3d72bd7f53dfbd07312938131c75b3bfd254da787947",
    }

    // instantiate the client
    client := blockchyp.NewClient(creds)

    // setup request object
    request := blockchyp.MessageRequest{}
    request.Test = true
    request.TerminalName = "Test Terminal"
    request.Message = "Thank you for your business."

    response, err := client.Message(request)

    // run the transaction
    response, err := client.Message(req)

    if err != nil {
        log.Fatal(err)
    }

    //process the result
    if response.Success {
        fmt.Println("Success")
    }

}

```
#### Refund

Executes a refund.

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"

    blockchyp "github.com/blockchyp/blockchyp-go"
)

func main() {

    // sample credentials
    creds := blockchyp.APICredentials{
        APIKey:      "ZDSMMZLGRPBPRTJUBTAFBYZ33Q",
        BearerToken: "ZLBW5NR4U5PKD5PNP3ZP3OZS5U",
        SigningKey:  "9c6a5e8e763df1c9256e3d72bd7f53dfbd07312938131c75b3bfd254da787947",
    }

    // instantiate the client
    client := blockchyp.NewClient(creds)

    // setup request object
    request := blockchyp.RefundRequest{}
    request.TerminalName = "Test Terminal"
    request.TransactionID = "<PREVIOUS TRANSACTION ID>"
    request.Amount = "5.00" // Optional amount for partial refunds.

    response, err := client.Refund(request)

    // run the transaction
    response, err := client.Refund(req)

    if err != nil {
        log.Fatal(err)
    }

    //process the result
    if response.Approved {
        fmt.Println("Approved")
    }

}

```
#### Enroll

Adds a new payment method to the token vault.

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"

    blockchyp "github.com/blockchyp/blockchyp-go"
)

func main() {

    // sample credentials
    creds := blockchyp.APICredentials{
        APIKey:      "ZDSMMZLGRPBPRTJUBTAFBYZ33Q",
        BearerToken: "ZLBW5NR4U5PKD5PNP3ZP3OZS5U",
        SigningKey:  "9c6a5e8e763df1c9256e3d72bd7f53dfbd07312938131c75b3bfd254da787947",
    }

    // instantiate the client
    client := blockchyp.NewClient(creds)

    // setup request object
    request := blockchyp.EnrollRequest{}
    request.Test = true
    request.TerminalName = "Test Terminal"

    response, err := client.Enroll(request)

    // run the transaction
    response, err := client.Enroll(req)

    if err != nil {
        log.Fatal(err)
    }

    //process the result
    if response.Approved {
        fmt.Println("Approved")
    }

    fmt.Println(response.Token)
}

```
#### Gift Card Activation

Activates or recharges a gift card.

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"

    blockchyp "github.com/blockchyp/blockchyp-go"
)

func main() {

    // sample credentials
    creds := blockchyp.APICredentials{
        APIKey:      "ZDSMMZLGRPBPRTJUBTAFBYZ33Q",
        BearerToken: "ZLBW5NR4U5PKD5PNP3ZP3OZS5U",
        SigningKey:  "9c6a5e8e763df1c9256e3d72bd7f53dfbd07312938131c75b3bfd254da787947",
    }

    // instantiate the client
    client := blockchyp.NewClient(creds)

    // setup request object
    request := blockchyp.GiftActivateRequest{}
    request.Test = true
    request.TerminalName = "Test Terminal"
    request.Amount = "50.00"

    response, err := client.GiftActivate(request)

    // run the transaction
    response, err := client.GiftActivate(req)

    if err != nil {
        log.Fatal(err)
    }

    //process the result
    if response.Approved {
        fmt.Println("Approved")
    }

    fmt.Println(response.Amount)
    fmt.Println(response.CurrentBalance)
    fmt.Println(response.PublicKey)
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
    "encoding/json"
    "fmt"
    "log"

    blockchyp "github.com/blockchyp/blockchyp-go"
)

func main() {

    // sample credentials
    creds := blockchyp.APICredentials{
        APIKey:      "ZDSMMZLGRPBPRTJUBTAFBYZ33Q",
        BearerToken: "ZLBW5NR4U5PKD5PNP3ZP3OZS5U",
        SigningKey:  "9c6a5e8e763df1c9256e3d72bd7f53dfbd07312938131c75b3bfd254da787947",
    }

    // instantiate the client
    client := blockchyp.NewClient(creds)

    // setup request object
    request := blockchyp.AuthorizationRequest{}
    request.TerminalName = "Test Terminal"
    request.TransactionRef = "<LAST TRANSACTION REF>"

    response, err := client.Reverse(request)

    // run the transaction
    response, err := client.Reverse(req)

    if err != nil {
        log.Fatal(err)
    }

    //process the result
    if response.Approved {
        fmt.Println("Approved")
    }

}

```
#### Capture Preauthorization

Captures a preauthorization.

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"

    blockchyp "github.com/blockchyp/blockchyp-go"
)

func main() {

    // sample credentials
    creds := blockchyp.APICredentials{
        APIKey:      "ZDSMMZLGRPBPRTJUBTAFBYZ33Q",
        BearerToken: "ZLBW5NR4U5PKD5PNP3ZP3OZS5U",
        SigningKey:  "9c6a5e8e763df1c9256e3d72bd7f53dfbd07312938131c75b3bfd254da787947",
    }

    // instantiate the client
    client := blockchyp.NewClient(creds)

    // setup request object
    request := blockchyp.CaptureRequest{}
    request.Test = true
    request.TransactionID = "<PREAUTH TRANSACTION ID>"

    response, err := client.Capture(request)

    // run the transaction
    response, err := client.Capture(req)

    if err != nil {
        log.Fatal(err)
    }

    //process the result
    if response.Approved {
        fmt.Println("Approved")
    }

}

```
#### Close Batch

Closes the current credit card batch.

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"

    blockchyp "github.com/blockchyp/blockchyp-go"
)

func main() {

    // sample credentials
    creds := blockchyp.APICredentials{
        APIKey:      "ZDSMMZLGRPBPRTJUBTAFBYZ33Q",
        BearerToken: "ZLBW5NR4U5PKD5PNP3ZP3OZS5U",
        SigningKey:  "9c6a5e8e763df1c9256e3d72bd7f53dfbd07312938131c75b3bfd254da787947",
    }

    // instantiate the client
    client := blockchyp.NewClient(creds)

    // setup request object
    request := blockchyp.CloseBatchRequest{}
    request.Test = true

    response, err := client.CloseBatch(request)

    // run the transaction
    response, err := client.CloseBatch(req)

    if err != nil {
        log.Fatal(err)
    }

    //process the result
    if response.Success {
        fmt.Println("Success")
    }

    fmt.Println(response.CapturedTotal)
    fmt.Println(response.OpenPreauths)
}

```
#### Void Transaction

Discards a previous preauth transaction.

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"

    blockchyp "github.com/blockchyp/blockchyp-go"
)

func main() {

    // sample credentials
    creds := blockchyp.APICredentials{
        APIKey:      "ZDSMMZLGRPBPRTJUBTAFBYZ33Q",
        BearerToken: "ZLBW5NR4U5PKD5PNP3ZP3OZS5U",
        SigningKey:  "9c6a5e8e763df1c9256e3d72bd7f53dfbd07312938131c75b3bfd254da787947",
    }

    // instantiate the client
    client := blockchyp.NewClient(creds)

    // setup request object
    request := blockchyp.VoidRequest{}
    request.Test = true
    request.TransactionID = "<PREVIOUS TRANSACTION ID>"

    response, err := client.Void(request)

    // run the transaction
    response, err := client.Void(req)

    if err != nil {
        log.Fatal(err)
    }

    //process the result
    if response.Approved {
        fmt.Println("Approved")
    }

}

```

## License

Copyright BlockChyp, Inc., 2019

Distributed under the terms of the [MIT] license, blockchyp-go is free and open source software.

[MIT]: https://github.com/blockchyp/blockchyp-go/blob/master/LICENSE
