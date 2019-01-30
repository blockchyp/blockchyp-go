# BlockChyp Command Line Interface

BlockChyp has a good set of SDK's for application developers and over the coming
months and years we'll be adding more.

But some platforms - for a variety of reasons ranging from legacy to architectural
complexity - would be better suited with a simple executable client.

The BlockChyp CLI allows developers to invoke the client via the standard shell
for their operating system.  Requests are passed in as command line arguments and
the results are returned to stdout as JSON.

## Sample Transaction

The following example shows a basic CLI charge transaction.

```
$ ./blockchyp -type=charge -terminal="Test Terminal" -amount="25.55"
{
  "approved":true,
  "responseDescription":"Approved",
  "transactionId":"NZ6FGYAYLYI6TLVWNSLM7WZLHE",
  "transactionRef":"cfd068099a4280f1f09a965d9cc522f25ef4e06a95c9a9461d59fa5deed62180",
  "transactionType":"charge","timestamp":"2019-01-15T00:42:36Z",
  "tickBlock":"000e61f8204a2a372cac288f833a8e0949dd50d0074d5133432dce4e78d97913",
  "authCode":"612797",
  "entryMethod":"CHIP",
  "paymentType":"VISA",
  "maskedPan":"************0010",
  "cardHolder":"Test/Card 01              ",
  "partialAuth":false,
  "altCurrency":false,
  "currencyCode":"USD",
  "requestedAmount":"20.55",
  "authorizedAmount":"20.55",
  "receiptSuggestions":{
    "AID":"A0000000031010",
    "ARQC":"E0A09074268A87F4",
    "IAD":"06010A03A0B800",
    "TVR":"0240008000",
    "TSI":"E800",
    "merchantName":"Test Merchant",
    "applicationLabel":"VISA CREDIT",
    "requestSignature":true,
    "maskedPan":"************0010",
    "authorizedAmount":"20.55",
    "transactionType":"charge",
    "entryMethod":"CHIP"
  }
}
```

## Building and Installing

For security reasons, BlockChyp does not distribute binaries for the CLI (at
least not yet).  But, binaries can be easily built from source.

Start by installing Go 1.11 or later on your workstation or CI environment.
If you're not already using `make`, install make as well.

In both cases, the build process is very fast and should add very little time
to your build process.



### For Linux and Mac OS

On Linux systems, use the following command to build the Linux CLI:

```
$ make cli-linux
```

This will create a Linux binary at `/build/blockchyp` that you can then deploy as part of your solution.

### For Windows

---
**NOTE**

These instructions assume you have a git client installed and available on your
path.  If you don't have a Windows git client, you can get the standard git
Windows client here...

https://git-scm.com/download/

---

If you have `make` in your Windows development environment, use the following command to build the Windows CLI:

```
> make cli-windows
```

If you don't have `make`, you can use `go build` directly...

```
go build -o builds\blockchyp.exe cmd\blockchyp\main.go
```

This will create a Windows exe file at `/build/blockchyp.exe` that you can then deploy
as part of your solution.

## Configuration

Key settings like API credentials can be passed in on the command line with
every request, but the best option is to configure the command line by dropping a
`blockchyp.json` file on your file system.

A typical `blockchyp.json` file looks like this:

```
{
  "apiKey":"ZDSMMZLGRPBPRTJUBTAFBYZ33Q",
  "bearerToken":"ZLBW5NR4U5PKD5PNP3ZP3OZS5U",
  "signingKey":"9c6a5e8e763df1c9256e3d72bd7f53dfbd07312938131c75b3bfd254da787947",
  "routeCacheTTL": 60,
  "gatewayTimeout": 20,
  "terminalTimeout": 2
}
```

The TTL and timeout settings are optional and so are the credentials, but the
credentials are highly recommended.  Otherwise malicious users might be able
to see your API credentials by looking at your shell command history.

You can specify the location of this file via the `-f` command line argument, but
BlockChyp does look for this file in a few default locations depending on your operating system.

For **Unix Like Systems**, BlockChyp looks for the file at
`$XDG_CONFIG_HOME/blockchyp/blockchyp.json`.  If the environment variable
isn't defined, the file location is assumed to be
`~/.config/blockchyp/blockchyp.json`.

For **Windows**, BlockChyp looks for the file in
`%HomeDrive%%HomePath%\blockchyp\blockchyp.json`.

These default locations are fine for development, but we recommend production systems
explicitly specify a file location via the `-f` argument.

## Command Line Options

| Option           | Description                                         | Example                                    |
|------------------|-----------------------------------------------------|--------------------------------------------|
| `-f`             | Specifies config file location.                     | `-f="/path/to/blockchyp.json"`             |
| `-t`             | Flags the transaction as a test transaction.        | `-t` (no value required)                   |
| `-type`          | Transaction type (charge, preauth, etc)             | `-type=charge`                             |
| `-gateway`       | Used to override gateway host name.                 | `-gateway=https://api.blockchyp.com`       |
| `-testGateway`   | Used to override the test gateway host name.        | `-testGateway=https://test.blockchyp.com`  |
| `-manual`        | Enter card information manually.                    | `-manual`                                  |
| `-apiKey`        | Used to override the API Key.                       | `-apiKey=ZDSMMZLGRPBPRTJUBTAFBYZ33Q`       |
| `-bearerToken`   | Used to override the bearer token.                  | `-bearerToken=ZLBW5NR4U5PKD5PNP3ZP3OZS5U`  |
| `-signingKey`    | Used to override the signing key.                   | `-signingKey=9c6a5e8e763df1c9256e3d72..`   |
| `-terminal`      | Name of the terminal for terminal transactions.     | `-terminal="Cashier #1"`                   |
| `-token`         | Token for token based transactions.                 | `-token=ZLBW5NR4U5PKD5PNP3ZP3OZS5U`        |
| `-amount`        | Amount to authorize for the transaction.            | `-amount=50.00`                            |
| `-tip`           | Tip amount, if needed.                              | `-tip=5.00`                                |
| `-tax`           | Tax amount, if needed.                              | `-tax=23.45`                               |
| `-currency`      | Currency code, defaults to USD.                     | `-currency=USD`                            |
| `-tx`            | Transaction ID.  Required for voids and captures.   | `-tx=DD62YSX6G4I6RM3XNSLM7WZLHE`           |
| `-txRef`         | Transaction reference.  Typically your application's internal ID. Required for reversable transactions  |  `-txRef=MYID` |
| `-desc`          | Narrative description of the transaction.           | `-desc="Adventures Underground #1"`        |
| `-secure`        | Can disable https for terminal transactions. Defaults to true.  | `-secure=false`                |
| `-version`       | Print the CLI version and exit.                     | `-version`                                 |
| `-out`           | Direct output to a file instead of stdout.          | `-out="output.json"`                       |
| `-routeCache`    | Specify a custom offline route cache location.      | `-routeCache="route_cache.json"`           |
| `-sigFormat`    | File format for signatures, if you'd like it returned with the transaction.  gif, jpeg, and png formats are supported.      | `-sigFormat="png"`           |
| `-sigWidth`    | If provided, signature images will be scaled to this max width.      | `-sigWidth="300"`           |
| `-sigFile`    | By default, signatures are returned in the response as hex.  If you'd rather have a file, use this option.     | `-sigFile="signature.png"`           |

## Test Transactions

BlockChyp's test system is designed to prevent mixups about whether a configuration
is pointed at live or test environments.  When using test credentials, all transactions
must have the `-test` flag set.

To run the sample transactions below against a test merchant account, just append
`-test` to the list of arguments.  For example,

```
> blockchyp.exe -type=charge -terminal="Test Terminal" -amount="20.55" -test
```


## Sample Transactions

The section below gives a few sample transactions for most common scenarios.

Note  that responses are shown below with standard JSON pretty printing white space.
Real CLI responses are more compact.

### Terminal Ping

This transaction tests connectivity with a payment terminal.

```
$ ./blockchyp -type=ping -terminal="Test Terminal"
{
  "success":true
}
```

### Charge

This transaction executes a direct auth and capture transaction against a BlockChyp
payment terminal.

```
> blockchyp.exe -type=charge -terminal="Test Terminal" -amount="20.55"
{
  "approved":true,
  "responseDescription":"Approved",
  "transactionId":"NZ6FGYAYLYI6TLVWNSLM7WZLHE",
  "transactionRef":"cfd068099a4280f1f09a965d9cc522f25ef4e06a95c9a9461d59fa5deed62180",
  "transactionType":"charge","timestamp":"2019-01-15T00:42:36Z",
  "tickBlock":"000e61f8204a2a372cac288f833a8e0949dd50d0074d5133432dce4e78d97913",
  "authCode":"612797",
  "entryMethod":"CHIP",
  "paymentType":"VISA",
  "maskedPan":"************0010",
  "cardHolder":"Test/Card 01              ",
  "partialAuth":false,
  "altCurrency":false,
  "currencyCode":"USD",
  "requestedAmount":"20.55",
  "authorizedAmount":"20.55",
  "receiptSuggestions":{
    "AID":"A0000000031010",
    "ARQC":"E0A09074268A87F4",
    "IAD":"06010A03A0B800",
    "TVR":"0240008000",
    "TSI":"E800",
    "merchantName":"Test Merchant",
    "applicationLabel":"VISA CREDIT",
    "requestSignature":true,
    "maskedPan":"************0010",
    "authorizedAmount":"20.55",
    "transactionType":"charge",
    "entryMethod":"CHIP"
  }
}
```

### Preauth

This transaction executes a preauthorization against a BlockChyp
payment terminal.

```
./blockchyp -type=preauth -terminal="Test Terminal" -amount="20.55"
{
  "approved":true,
  "responseDescription":"Approved",
  "transactionId":"NZ6FGYAYLYI6TLVWNSLM7WZLHE",
  "transactionRef":"cfd068099a4280f1f09a965d9cc522f25ef4e06a95c9a9461d59fa5deed62180",
  "transactionType":"preauth",
  "timestamp":"2019-01-15T00:42:36Z",
  "tickBlock":"000e61f8204a2a372cac288f833a8e0949dd50d0074d5133432dce4e78d97913",
  "authCode":"612797",
  "entryMethod":"CHIP",
  "paymentType":"VISA",
  "maskedPan":"************0010",
  "cardHolder":"Test/Card 01              ",
  "partialAuth":false,
  "altCurrency":false,
  "currencyCode":"USD",
  "requestedAmount":"20.55",
  "authorizedAmount":"20.55",
  "receiptSuggestions":{
    "AID":"A0000000031010",
    "ARQC":"E0A09074268A87F4",
    "IAD":"06010A03A0B800",
    "TVR":"0240008000",
    "TSI":"E800",
    "merchantName":"Test Merchant",
    "applicationLabel":"VISA CREDIT",
    "requestSignature":true,
    "maskedPan":"************0010",
    "authorizedAmount":"20.55",
    "transactionType":"charge",
    "entryMethod":"CHIP"
  }
}
```

### Capture

Captures an existing preauthorization.  `-tx` is required and developers have
the option of adding tip adjustments or changing the amount.

```
> blockchyp.exe -type=capture -tx=DD62YXX6G4INSLM7WZLHE -tip=5.00 -amount=55.00
{
  "responseDescription":"Approved",
  "transactionId":"DD62YXX6G4I6RM35NSLM7WZLHE",
  "batchId":"OGMJ72X5MUI6RD7MNSLM7WZLHE",
  "transactionType":"capture",
  "timestamp":"2018-12-12T21:20:11Z",
  "tickBlock":"009991a8ac7b6a4420760e1e14e1689c88be2a610a033d6908c1b04b5c00f9da",
  "approved":true,
  "authCode":"349143",
  "entryMethod":"CHIP",
  "paymentType":"VISA",
  "partialAuth":false,
  "altCurrency":false,
  "currencyCode":"USD",
  "requestedAmount":"55.00",
  "authorizedAmount":"55.00",
  "tipAmount":"5.00",
  "taxAmount":"0.00"
}
```

### Void

Voids an existing transaction in the current batch.  `-tx` is required.

```
$ ./blockchyp -type=void -tx=DD62YVH6G4I6RM33NSLM7WZLHE
{
  "responseDescription":"Approved",
  "transactionId":"DD62YXX6G4I6RM36NSLM7WZLHE",
  "batchId":"OGMJ72X5MUI6RD7MNSLM7WZLHE",
  "transactionType":"void",
  "timestamp":"2018-12-12T21:24:19Z",
  "tickBlock":"009991a8ac7b6a4420760e1e14e1689c88be2a610a033d6908c1b04b5c00f9da",
  "approved":true,
  "authCode":"686941",
  "entryMethod":"CHIP",
  "paymentType":"VISA"
}
```

### Time Out Reversals

Time out reversals are used to cancel transactions that may or may not have
gone through.  In order to use reversals, always provide a value for the `-txRef`
option as shown in the sample charge transaction below.

```
$ ./blockchyp -type=charge -terminal="Test Terminal" -amount=25.00 -txRef=4373223444
Request Timed Out
```

If the request times out, you have 2 minutes to submit a reversal as shown in the
next sample transaction.

```
./blockchyp -type=reverse -txRef=4373223444
{
  "responseDescription":"No Action Taken",
  "transactionId":"DD62Y2H6G4I6RM4ANSLM7WZLHE",
  "batchId":"OGMJ72X5MUI6RD7MNSLM7WZLHE",
  "transactionType":"reverse",
  "timestamp":"2018-12-12T21:32:19Z",
  "tickBlock":"009991a8ac7b6a4420760e1e14e1689c88be2a610a033d6908c1b04b5c00f9da",
  "approved":true
}
```

### Gift Card Activation

This transaction can be used to activate or add value to a BlockChyp gift card.

Note that BlockChyp gift cards do not have numbers.  They're identified by
public key.

```
./blockchyp -type=gift-activate -terminal="Test Terminal" -amount=25.00
{
  "responseDescription":"Approved",
  "transactionId":"DD62Y2H6G4I6RM4ANSLM7WZLHE",
  "transactionType":"gift-activate",
  "timestamp":"2018-12-12T21:32:19Z",
  "tickBlock":"009991a8ac7b6a4420760e1e14e1689c88be2a610a033d6908c1b04b5c00f9da",
  "approved":true,
  "currencyCode":"USD",
  "amount": "25.00",
  "currentBalance": "25.00",
  "publicKey": "342a40ada947bd35886f19c8908cd84e521f713cc2637c0bf70b3b2ea63ffe7d"
}
```

### Close Batch

This transaction will close and submit the current batch for settlement.

```
$ ./blockchyp -type=close-batch
{
  "responseDescription": "Closed",
  "batchId": "UEOHSRX2MYI6RA2WSSDM7WZLHE",
  "transactionRef": "b944f032e997d944cdabb03cf1aa260ba3cde3d3b572b138eceb27bb41e54332",
  "test": false,
  "transactionId":"UEOHSRX2MYI6RA2LNSLM7WZLHE",
  "transactionType":"close-batch",
  "timestamp":"2018-12-07T21:25:37Z",
  "tickBlock":"000a40ada947bd35886f19c8908cd84e521f713cc2637c0bf70b3b2ea63ffe7d",
  "currencyCode":"USD",
  "capturedTotal": "1712.04",
  "openPreauths": "120.00",
  "cardBrands": {
    "VISA": "500.00",
    "MC": "120.00",
    "AMEX": "800.00",
    "DISC": "292.04"
  }

}
```

## The Route Cache

BlockChyp automatically locates payment terminals on your network, even if you
stick with DHCP (which you shouldn't do).  Every time a payment terminal comes
online, it reports its internal IP address to the BlockChyp gateway.  It also
updates its network status periodically.

When the CLI is asked to run a transaction, it starts by looking up the terminal
on the gateway.  The gateway returns the IP Address, Public Key, and a set of
transient credentials.  For most BlockChyp SDK's, the SDK maintains an in memory
cache of recent terminal routes.  For the CLI, this is problematic because every
invocation of the CLI is a new process, making in memory caches unfeasible.  We
address this by maintaining an offline cache file.  This file is stored in your
temp directory by default. You can use the `-routeCache`
parameter to override this location if you'd like.




## What Are Tick Blocks?

You may have noticed that almost every API response in BlockChyp returns something
called a **tick block**.

This is essentially an internal blockchain timestamp generated by BlockChyp's
proof-of-work mining system.  BlockChyp uses a blockchain data
model under the hood and this system uses block hashes alongside timestamps to
record when transactions actually happened.

BlockChyp's internals are mostly hidden from developers (for now) and you don't
really need to store or worry about tick blocks in your application.
