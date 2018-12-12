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
$ ./blockchyp -type=charge -terminal="Test Terminal" -amount="25.00"
{
  "responseDescription":"Approved",
  "transactionId":"DD62YSX6G4I6RM3XNSLM7WZLHE",
  "transactionType":"charge",
  "timestamp":"2018-12-12T20:23:57Z",
  "tickBlock":"009991a8ac7b6a4420760e1e14e1689c88be2a610a033d6908c1b04b5c00f9da",
  "approved":true,
  "authCode":"554738",
  "entryMethod":"CHIP",
  "paymentType":"VISA",
  "maskedPan":"************0119",
  "partialAuth":false,
  "altCurrency":false,
  "currencyCode":"USD",
  "requestedAmount":"25.00",
  "authorizedAmount":"25.00",
  "tipAmount":"0.00",
  "taxAmount":"0.00",
  "receiptSuggestions":{
    "AID":"A0000000031010",
    "ARQC":"8A3054E0EA328A2A",
    "IAD":"06010A03A0A800",
    "TVR":"8000008000",
    "TSI":"6800",
    "requestSignature":true
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

On Windows, use the following command to build the Windows CLI:

```
> make cli-windows
```

This will create a Windows exe file at `/build/blockchyp.exe` that you can then deploy
as part of your solution.

## Configuration

Key settings like API credentials can be passsed in on the command line with
every request, but the best option is to configure the command line by dropping a
`blockchyp.json` file on your file system.

You can specify the location of this file via the `-f` command line argument, but
BlockChyp looks for this file a few default locations.

For **Linux**, BlockChyp looks for the file in the directory specified by the
`XDG_CONFIG_HOME` environment variable.  If the environment variable isn't defined,
the file location is assumed to be `~/.config/blockchyp.json`.

For **Windows**, BlockChyp looks for the file in the user's home directory.

These default locations are fine for development, but we recommend production systems
explicitly specify a file location via the `-f` argument.

## Command Line Options

| Option         | Description                                    | Example                                  |
|----------------|------------------------------------------------|------------------------------------------|
| -f             | Specifies config file location.                | -f="/path/to/blockchyp.json"             |
| -t             | Flags the transaction as a test transaction.   | -t (no value required)                   |
| -type          | Transaction type (charge, preauth, etc)        | -type=charge                             |
| -gateway       | Used to override gateway host name.            | -gateway=https://api.blockchyp.com       |
| -testGateway   | Used to override the test gateway host name.   | -testGateway=https://test.blockchyp.com  |
| -apiKey        | Used to override the API Key.                  | -apiKey=ZDSMMZLGRPBPRTJUBTAFBYZ33Q       |
| -bearerToken   | Used to override the bearer token.             | -bearerToken=ZLBW5NR4U5PKD5PNP3ZP3OZS5U  |
| -signingKey    | Used to override the signing key.              | -signingKey=9c6a5e8e763df1c9256e3d72..   |
| -terminal      | Name of the terminal for terminal transactions.| -terminal="Cashier #1"                   |
| -token         | Token for token based transactions.            | -token= ZLBW5NR4U5PKD5PNP3ZP3OZS5U       |
| -amount        | Amount to authorize for the transaction.       | -amount=50.00                            |
| -tip           | Tip amount, if needed.                         | -tip=5.00                                |
| -tax           | Tax amount, if needed.                         | -tax=23.45                               |
| -currency      | Currency code, defaults to USD.                | -currency=USD                            |
| -tx            | Transaction ID.  Required for voids and captures.   | -tx=DD62YSX6G4I6RM3XNSLM7WZLHE      |
| -txRef         | Transaction reference.  Typically your application's internal ID. Required for reversable transactions  |  -txRef=MYID |
| -desc          | Narrative description of the transaction.   | -desc="Adventures Underground #1"  |
| -secure   | Can disable https for terminal transactions. Defaults to true.  | -secure=false   |
