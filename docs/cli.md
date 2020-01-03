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
| `-taxExempt`     | Flags a transaction as tax exempt for Level 2 processing.  | `-taxExempt`                        |
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
| `-sigWidth`    | If provided, signature images will be scaled to this max width.      | `-sigWidth="300"`           |
| `-lineItemId`    | Optional ID for a line item.   | `-lineItemId="1234"`           |
| `-lineItemDescription`    | Description of a line item for line item display.     | `-lineItemDescription="Black Diamond Trekking Poles"`           |
| `-lineItemQty`    | Quantity of the associated line item.  Decimals are supported.     | `-lineItemQty="2"`           |
| `-lineItemPrice`    | Price of the line item.    | `-lineItemPrice="129.99"`           |
| `-lineItemExtended`    | Price times quantity less discounts.  Will auto-calculate if you don't provide it.   | `-lineItemExtended="259.98"`           |
| `-lineItemDiscountDescription`    | A line item specific discount.  | `-lineItemDiscountDescription="Member Discount"`           |
| `-lineItemDiscountAmount`    | Amount of the discount. | `-lineItemDiscountDiscountAmount="20.00"`           |
| `-displaySubtotal`    | Subtotal for all line items on the display. | `-displaySubtotal="239.98"`           |
| `-displayTax`    | Tax to be displayed on the terminal | `-displayTax="11.02"`           |
| `-displayTotal`    | Grand total for the line item display | `-displayTotal="250.00"`           |
| `-prompt`    | Text to be display on a boolean prompt screen. | `-prompt="Would you like to supersize that?"`           |
| `-promptType`    | Type of prompt for text-prompts. Could be 'email', 'phone', 'customer-number', or 'rewards-number' | `-prompt="email"`           |
| `-yesCaption`    | Overrides the label for the 'Yes' button on boolean-prompt screens. | `-yesCaption="Definitely"`           |
| `-noCaption`    | Overrides the label for the 'No' button on boolean-prompt screens. | `-noCaption="I Think Not"`           |

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

### Refunds

BlockChyp supports refunds that reference previous transactions and refunds that
do not reference a previous transaction, which we call free range refunds.

We strongly recommend that refunds always reference a previous transaction, but
we know we don't live in an ideal world.

The example below shows how to refund the entire amound of a previous transaction
without needing to lookup or pass in the amount.

```
$ ./blockchyp -type=refund -tx=IU245GBFV4I6THQ3AJBKYEIAAY
  {
    "responseDescription":"Approved",
    "transactionId":"IU245GBFV4I6THQ3AJBKYEIAAY",
    "transactionType":"refund",
    "timestamp":"2019-01-31T23:37:56Z",
    "tickBlock":"00014bd97158fdbc02a0923a5a339db853c074aa3a460b9f54aafd003e630296",
    "test":false,
    "approved":true,
    "authCode":"313097",
    "currencyCode":"USD",
    "authorizedAmount":"55.55"
  }
```

#### Partial refunds

To refund some, but not all, of a previous transaction, just add an amount argument
as shown below...

```
$ ./blockchyp -type=refund -tx=IU245HBFV4I6THQ3AJBKYEIAAY -amount="25.00"
  {
    "responseDescription":"Approved",
    "transactionId":"IU245HBFV4I6THQ3AJBKYEIAAY",
    "transactionType":"refund",
    "timestamp":"2019-01-31T23:41:25Z",
    "tickBlock":"00014bd97158fdbc02a0923a5a339db853c074aa3a460b9f54aafd003e630296",
    "test":false,
    "approved":true,
    "authCode":"139089",
    "entryMethod":"CHIP",
    "paymentType":"MC",
    "currencyCode":"USD",
    "requestedAmount":"25.00",
    "authorizedAmount":"25.00"
  }
```

#### Free Range Refunds

If you have to refund a card directly without referencing a previous transaction,
the syntax is similar to charge and preauth transactions.

```
$ ./blockchyp -type=refund -terminal="Test Terminal" -amount="25.00"
  {
    "responseDescription":"Approved",
    "transactionId":"IU245HRFV4I6THQ3AJBKYEIAAY",
    "transactionRef":"2073d065659065e1ab0a565765992743993bc75e8875bdaad5e00947dddc029c",
    "transactionType":"refund",
    "timestamp":"2019-01-31T23:43:55Z",
    "tickBlock":"00014bd97158fdbc02a0923a5a339db853c074aa3a460b9f54aafd003e630296",
    "test":false,
    "approved":true,
    "authCode":"557464",
    "sigFile":"",
    "entryMethod":"CHIP",
    "paymentType":"MC",
    "maskedPan":"************0434",
    "cardHolder":"Test/Card 10            ",
    "currencyCode":"USD",
    "requestedAmount":"25.00",
    "authorizedAmount":"25.00",
    "receiptSuggestions":{
      "AID":"A0000000041010",
      "ARQC":"47197067E0E722D2",
      "IAD":"0110A0000F220000000000000000000000FF",
      "TVR":"0840008000",
      "TSI":"E800",
      "merchantName":"Test Merchant",
      "applicationLabel":"MasterCard",
      "maskedPan":"************0434",
      "authorizedAmount":"25.00",
      "transactionType":"refund",
      "entryMethod":"CHIP"
    }
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


### Balance Checks

This transaction type is used to check the remaining balance for payment
types for which a remaining balance is relevant, like gift cards and ebt.

This first example shows the process for a gift card balance check:

```
./blockchyp -type=balance -terminal="Test Terminal"
{
  "success": true,
  "responseDescription": "Approved",
  "transactionId": "EJJL6PBOJUI6VIGYNSLM7WZLHE",
  "transactionType": "balance",
  "tickBlock": "009991a8ac7b6a4420760e1e14e1689c88be2a610a033d6908c1b04b5c00f9da",
  "test": false,
  "remainingBalance": "100.00"
}
```

For EBT balance checks, append the `-ebt` argument to the command as shown below:

```
./blockchyp -type=balance -terminal="Test Terminal" -ebt
{
  "success": true,
  "responseDescription": "Approved",
  "transactionId": "EJJL6RROJUI6VIGYNSLM7WZLHE",
  "transactionType": "balance",
  "tickBlock": "009991a8ac7b6a4420760e1e14e1689c88be2a610a033d6908c1b04b5c00f9da",
  "test": false,
  "remainingBalance": "100.00"
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

## Manual/Keyed Transactions

If enabled in the Dashboard (under Merchant Settings), you can bypass
the usual terminal behavior and enter keyed transactions manually.

We don't recommend this for conventional retail transactions, but it might
be necessary for telephone orders.

To explicitly put the terminal in manual mode, just add `-manual`
to the arguments for a `charge`, `preauth`, or `refund` transaction.

The example below shows a typical manual transaction.

```
./blockchyp -type=charge -terminal="Test Terminal" -amount="25.00" -manual
  {
    "responseDescription":"Approved",
    "transactionId":"IU245IBFV4I6THQ3AJBKYEIAAY",
    "transactionRef":"e1e3620daa63e849733434993acf41be64815e9074988efb67701316f07e8eac",
    "transactionType":"charge",
    "timestamp":"2019-01-31T23:52:51Z",
    "tickBlock":"000e5faa310afb68fb78792d6728b55a07fb5476cd37a33a86154765b4905201",
    "test":false,
    "approved":true,
    "authCode":"044879",
    "entryMethod":"MANUAL",
    "paymentType":"VISA",
    "maskedPan":"************1111",
    "currencyCode":"USD",
    "requestedAmount":"25.00",
    "authorizedAmount":"25.00",
    "receiptSuggestions":{
      "merchantName":"Test Merchant",
      "maskedPan":"************1111",
      "authorizedAmount":"25.00",
      "transactionType":"charge",
      "entryMethod":"MANUAL"
    }
  }
```

### Line Item Display

This command add items to the line item display.

If you run a charge or preauth transaction immediately after populating the line item display, the line item data will be used for Level 3 processing and the display data will be cleared after the transaction.

-displaySubtotal, -displayTax, and -displayTotal are required.  -lineItemSubtotal will be autocalculated if you don't provide it.

```
$ ./blockchyp -type=display -terminal="Test Terminal" -displaySubtotal="120.05" -displayTax="5.00" -displayTotal="125.05" -lineItemDescription="Leki Trekking Poles" -lineItemQty=1 -lineItemPrice="135.05" -lineItemDiscountDescription="Member Discount" -lineItemDiscountAmount="10.00" -lineItemSubtotal="120.05"
{
  "success":true,
  "error":""
}
```

### Clear Terminal

This command clears the terminal if a transaction is in progress. It also clears the line item display buffer.

```
$ ./blockchyp -type=clear -terminal="Test Terminal"
{
  "success":true,
  "error":""
}
```

### Display Message

This command displays a free form message on the terminal.

```
$ ./blockchyp -type=message -terminal="Test Terminal" -message="Thank you for your business."
{
  "success":true,
  "error":""
}
```

### Boolean Prompt

This command asks the user a yes or no question.

-yesCaption and -noCaption are optional.

```
$ ./blockchyp -type="boolean-prompt" -terminal="Test Terminal" -prompt="Would you like to become a member?" -yesCaption="Yes" -noCaption="No"
{
  "success":true,
  "error":"",
  "response": true
}
```

### Text Prompt

This command captures text input from the user.  Due to PCI restrictions, free
form prompts are not allowed. You must pick between email, phone numbers, customer
numbers, and rewards numbers.  We'll add more types over time.

```
$ ./blockchyp -type="text-prompt" -terminal="Test Terminal" -promptType="phone"
{
  "success":true,
  "error":"",
  "response": "5095901945"
}
```

## Signature Images

If digital signature capture is enabled in the dashboard, BlockChyp will upload the
signature image to the gateway for archival after each transaction.  You can easily
bring up the transaction in the dashboard if you need to and inspect the signature,
so by default signatures are not returned with API calls.

This means you don't need to archive signatures in your system, but if you absolutely
must have the signatures, you have options for returning them with the transaction.

You can retrieve the signature file in png, jpeg, or gif format and you can also
scale the signature image to whatever maximum width you specify.  If you leave
out a width argument, the image will be returned in its native resolution.  We recommend
using images in their native resolution because terminal response time will be
all that much faster.

Images can be returned in the JSON request as hex or you can specify a file to
write the image to if that works better you.  We like to give you options.

### In Hex

This example scales an image to a max width of 200 pixels and returns the image
in hex.

```
./blockchyp -type=charge -terminal="Test Terminal" -amount="25.00" -sigFormat=png -sigWidth=200
  {
    "responseDescription":"Approved",
    "transactionId":"ITAK4GBFWUI6TONDAJBKYEIAAM",
    "transactionRef":"715b31bc958b4de6f0afab60c39823bc39f792f2c151420d52b0bc3330776db9",
    "transactionType":"charge",
    "timestamp":"2019-02-01T00:07:55Z",
    "tickBlock":"000e5faa310afb68fb78792d6728b55a07fb5476cd37a33a86154765b4905201",
    "test":false,
    "approved":true,
    "authCode":"523979",
    "entryMethod":"SWIPE",
    "paymentType":"DISC",
    "maskedPan":"************1117",
    "currencyCode":"USD",
    "requestedAmount":"25.00",
    "authorizedAmount":"25.00",
    "receiptSuggestions":{
      "merchantName":"Test Merchant",
      "maskedPan":"************1117",
      "authorizedAmount":"25.00",
      "transactionType":"charge",
      "entryMethod":"SWIPE"
    },
    "sigFile":"89504e470d0a1a0a0000000d49484452000000c8000000370802000000c6fe92050000150e49444154789cec7c7b549357d6779e272181818624dc6f96695150b90d8c33163ea752b04ea59ddea8ab33305d382d438719c401aa4871ba28b2049652d17670891d14e948c70b282092414590292a7211546e0a09888110024848489ecbb75ef7fb3e6f564242a05c9437bf3f58e4c9b9ecb3cfefecbdcf3ee7098324499a1146cc37d0a516c088e50923b18c5810188965c482c048ac6714e4532cb51473879158cf100882c09f8246a3214f41a3d1e0e3730723b1961e2449027b5014a53f054110c3c3c3bdbdbd2449d2e9f4a516702e409e6b7bbb0c4010048afed7f2c630eceeddbbf7efdf1f1818989898303333a3d3e9bdbdbd2b57ae8c8d8da58a3d2f602cb500342a9e204912411075f5e97abe6c4092248aa21289e4c2850b9d9d9df6f6f61e1e1eebd7af777474643299341a2d2525c5cbcb0b3ce3520b3b3b2c99c5a248a36dea0982807f3448f63cea570f6044b76fdf2e2c2c7cedb5d7828383cdcdcdd5bf55a954341a0d18f6dc61092c164110c02760094992fdfdfd9d9d9d5d5d5d9393934aa512c3b0ff168ec16032993c1ecfdbdbdbcfcf6ff1455d5080b93a71e2446262e28b2fbe48c5e9288a42e40e947aee9c2060512d168ee3a0351a8d363636d6d0d0d0dcdc2c97cb592c96b3b3b3878787adadada5a5e54f7ef213282f93c9a452a94020a8abab532a959f7efaa9b5b535b55d7ade81e3389d4ebf74e99248248a8c8cc4308cc17826229379c122114b7dd9ddb973e7f2e5cb128964d5aa557e7e7e6e6e6ea6a6a633b6505959d9dcdcbc6bd72e988f851779314010048661a9a9a9212121414141cfa9719a168bb144800a1886151717b7b5b53199ccd75f7f7dddba755401708ec8ff40bd2e158a050505b5b4b468045ecf3bc0df7df1c5178989891e1e1ef6f6f6a087a5966b1eb0b0c40235d1e9f4ab57af565454ac59b3262626c6cece0ebea53ca31eae00d54892c4306c6a6a6acef615c330f429e63a940501822004413099cca8a8a813274e2425251104b13c88b5808a061da954aaafbefaeafaf5eb090909919191767676388ec3be8f8adf0d69a7acaccccece0e6662b662d068b463c78e0d0d0d517bb16707288ae238eee5e5a552a9040201aca2a5166a1eb050c4026b343434b46bd7ae152b56ecd9b3072805fb41c32d076c9dbefbeebbb6b6b6f7df7f1f3ece4a121445a7a6a6868686acadad173961612045804c9b366d3a7ffefc1c56ceb3890571851054ddb973273737372626c6cbcb0b783687a01b78b065cb162e973b0749c017f7f5f5b1d96c0683b1c881bf812446519424c95ffce217e5e5e522916879445af36fb160f24a4b4bf3f3f35353538155067a3d5de072b910e0cfb62254191e1eb6b2b2a21e2e8249807e070707e572b921e5613f181818585555b538122e34e69958c0a1e3c78f373737676565d9dada1204f1e38d04ccd39c89f5e8d123cae081335d9c38262f2f0f88356377e0df7ffef39f0b85c26761e7ab2130b53737bc85f91c00b0eaca952b42a170cf9e3d262626730889a605ec1ce7d01498c9aeae2e48dcc326f4debd7b0b1a2383237bf4e8119d4ee7f178866cf4a080b5b535499272b97c69fda0ba2356bfc333ab6b3cf3462ce0d0e0e060717171525212c4e9b35510a906ea21f8855bb76ef1f9fc59b9091069606000c3304747470cc31004110a853b76ec282c2c44106481ae3a81f09d9d9d7050632083615ca6a6a6a3a3a38bb07bd5d53e2c03a150a8502860e74ea3d12412494f4f8f4422a19ecc28debc05ef308bb9b9b99f7cf20993c99c6d3e06d4aa6e93a82c1758171717979a9a9ac6c64660ad21ee15bcf0cd9b37d7ae5d4b7571ead4a9a3478f9695959597978786861ade14a8923a92d20328f9e0c183595d4c805a743a5d269319521e34acab7110585d5a8df2d356046db4b5b57df7dd775f7cf105499267ce9c696e6ee672b9743a1dc771a954eae3e313161636b3f720e703609f6a6a6ab2b3b3e1587e56d5410b24494a2492aea7181e1e566f99c2e1c3878b8b8b215f6aa0547bf7ee7dfcf8313c91c964a9a9a9f0d5f6eddb452211f114068a37ed475de5d3d3d36114339607c088727272bababa0cafa50d4813aa7fd4d0214992535353df7ffffdb402080482e8e86891484492e4810307fefef7bf8f8f8f5365c6c6c6f2f3f3e3e3e37b7a7a343ad2c0fc582c302a151515f1f1f1b3bdf4081ba22b57ae5cbf7e9d4ea75b5858d06834b1586c6b6bfbd1471f71381c5846a0a34f3ffdf4b3cf3e5bbd7ab5bbbbbbfe933558afc3c3c3341acddede1e1ae9edede572b9f0d51ffef0876fbef9e6cb2fbfd47ffa0bbd343535d5d6d6dadbdb878484f0783c3d5e1ebe92482424495a5959cd361e50a954339607911a1b1b5f7ae9250e87437541d9541a8dd6d4d4d4d8d808fc0043c862b1c2c3c36d6d6d21e81408049489a2fee9e9e9d9bf7fffeeddbbedececeaebeba7a6a6e2e3e3a165e894cd66474646b6b5b589c5625757573d7e691e88055d363434383838585b5bcf2a570485333232984c664444c44b2fbd447d555e5e9e9a9aba65cb964d9b360177216791989898959575e0c001fd13007ef0f6eddbaeaeae544712890402648542e1e3e3535757575050f0d1471fe9e216cc9940203871e2c4ef7ffffbbebebeb4b4b4f4f47438359f5600a872f7ee5d2727274a0c435481a2a85c2e1f1f1f7776768646f4af9cf2f2f2a8a828f5f142e1aaaaaaab57af72389cc0c0c00f3ef8c0dcdc1c4190a9a929894462696909f2fcf0c30febd7afa7b808ca110804fbf7ef873b3c2449b2582cc8d1c0b7d428701cf7f4f4a464d6399eb9d95b6d139a9d9ddddada0aaf03ccaa626969e9e1c387e1099cf663180636766464243d3d3d232343a55241b350a5acac2c272747bf4384f269696942a190f2cefffce73ff97cbe7ab1f4f4f42b57aee86a0a1e666565353535c1133e9f9f9999a9a76b78fecd37df3436366abb72fdaae0f3f9c78e1d838ffa5da152a9fcf2cb2fa9c6e11fb158bc73e7ce83070f0a04025d150982989898d8bd7bb752a9842ee02f7840a808bdab54aa9d3b77ca6432ed68615af7aa8179d815026d711c777171d10827a9d70474559c9c9cbc76edda279f7c020e1bcc1258261cc7b95c6e7272f24f7ffad3cccc4c1445310c8310323434747474f4c68d1bf0de8176cbd0544747079d4e777171817d008d46130a85dedede0d0d0dc78f1ffffaebaff3f3f38383832f5cb8f0f0e143eda6c0a70f0d0dc9e5725f5f5f0cc3542ad5a64d9b0882a8ababd3d53574343a3a0a9672562992eaeaead0d050eaa2476565a576200f9d76747470381c386704cf0e299ead5bb7c6c5c5ad58b1823a90d5980b04412e5ebce8ebeb0bc9a0fafa7a82203a3a3af6eddbb77bf76ea808810783c178f5d5574f9d3aa57dca6448ea677ed20d24494aa552068341519bba0943a7d3150a8576021a06595555e5efef0f9e454356181e86615bb76eb5b1b1c9cbcb833319486f42ca00f4aebdf585ae2b2b2b376fde4ca5c1262626140a059c64af5ab56af3e6cdab56adb2b0b0484848c8cece9e9898d0680afe1708042e2e2ed00883c12049f2e38f3fbe78f1a22e5520082293c948928460ce10ed5167152b56acb0b7b7878b18fff8c73fbabbbb4d4d4db5f90177dadcdddda9271289242b2b6bd7ae5dfefefe60fcb40f6441094aa5b2b1b1f137bff90dac3d9224fff297bf9c3a752a2525e5c5175fa41c378aa204416cd9b2a5b3b3532412cd21a53c0fc482590c0c0c2c2828002ec318100479f8f0e1c99327d3d2d2209855170e0c5b474747404080ae950d738961d81ffff8479148c4e7f3816d2449b2d9ecf7de7b2f3f3f1f56ad863c288a8ac5e2c1c1c19ffdec6754d6beadadcdd1d19124497f7fff808080952b57060606ae5dbbd6d9d9393c3c7cefdebda04df5762085c3e17028790882b0b1b1717272aaaeaed6284f55e9ebeb834bb086a4dc8055030303555555dbb66dc371dcc4c4442c16777777fff9cf7fd6b60da0b7fefe7e087460b0df7fff7d585898abab2b048bd3067fc0a4d6d656575757535353c8eab5b6b6deba752b2525c5c9c969da90ee77bffbddd1a347e770343e3fae10d82d168b8f1f3fdeddddddd5d5555555959292525454646161919292e2ececac11ed8209999c9c747171d1bf7502a7b373e7cef2f2f2c1c1416807c7f1a0a020b1585c5f5f0f968c2a0f1a3c7dfaf4af7ffd6b5894d0fecd9b37fdfdfd1104a12218f88be3f82bafbce2e5e575e4c811b890a8debb52a934313151179b24c9f7df7f1f0ef53400da6f6d6d5dbd7ab521192c60954c26dbb76f5f4c4c0c93c9046ab6b4b4787b7b8384dada96c9640a85c2c1c1011c964c261b1818d8b8712341107ab6b7d0727d7dbd9f9f1f49924c26f3d0a1432449464545414241435a98561f1f1f369b7dfefc795dae5f17e6c71582bafff6b7bf31188cf2f2f28b172ff6f4f4fce94f7f4a4e4e7ef7dd77cdcccc34e486196d6d6db5b3b38338914a8a68370e15592cd6b66ddbf2f2f260c0f0373e3efee4c9938f1f3f86d88b5ac122914828146ed8b0818adbc6c7c7c562b1afaf2f4c86ba5905328587870f0c0c3437378343571f9afa4758bb7676768e8e8e7c3e5fc35e82a83d3d3dfa8945659be8747a5757575252d2871f7e080914a8323838e8ececac9dff0449babbbbe1ee3fac81969696975f7e79c66c380c442291ac59b306419023478e1004111d1d3d3939299148a61515aac4c5c555575743dadd7087386f473a2059444444dc534445458175a50e9bd40b03cf1a1a1a828383e9743a93c98498803230da23c471dcd7d71745d1f6f6766aeb6b6969191b1b9b9999a952a9604941cb478f1efded6f7f0b6560321a1a1adcdcdcb49d1700eac6c6c61e3b764c2a95aa87140a85827ab983128620888888888a8a0a9812aa30c8a952a91c1c1cf4100bd88ca2684949c9d75f7f1d171717181848ed3020f687adbeb6dee08d31202ee0c993276c365bffec0065efddbb676262626e6e5e5a5a3a3636b663c70e8220381c8e542ad5431a0441c2c2c2befdf65b5dda9b16f37fbb8102180fed2c0e3c97c9647d7d7db5b5b5393939870f1f3e77ee5c4b4b8b542ad5754c01f3f7ce3bef14151551e7151886797878040404bcf9e69b4d4d4d40cd8282020e87e3e3e3a31e34343535050606ea09e668349a8d8d4d7474747a7aba4aa5a2cca7542ad59863f887cd664744441c3c7890523794178bc5b054a69d0358360f1f3e2c2c2c4c4e4eeeebebcbc8c8707373d348fe6118a6ee7f29a028aa52a97a7a7a200b05556c6c6ce07871461417177ff8e18742a190cfe7272626aa542a1445391c8e1e6b044b2530309020887ffffbdf94679811f37cd1cff023bc8a8a0a7f7fff575e7945a150e0382e1008fef39fff5cba7409c3b0d0d0505f5f5f8d5812acc89aa7181f1f878c33b8adb7df7ebba2a2e2ead5ab2525250882585a5afef5af7fa5aad3e9f4d1a778f9e597f50473d415e10d1b36e4e6e66edfbe5da954d268b4898909b872a31120e238eeefef7ff3e6cdb2b2b237df7c1336b9e00721bbadc700b4b7b75b5858c4c7c7c3a556edb38a69eb82d5696868b0b7b7373333a3b8e8e2e2525252a2dff3a2282a1008582c1687c3f9fcf3cf93929228abcce5724522919ef9820d535c5cdc9e3d7b02020234ecb72e2cc18b6c9008be71e3465a5a1af5e21765db474646a81daf76459224b76edd4a7d04adb158acf5ebd7bffefaeb161616a6a6a62fbcf0827a768d4ea7d7d6d67a7a7a42c8af87fa106cbdfdf6db7bf7eeada9a9f9d5af7e45108452a9e4f178d316260862dbb66dc9c9c9afbdf69aa9a9296ce8eedebdbb66cd1a5dd34c5d88858fea2f5aaa83c5624d4d4d693c047e545656464646aa0fd0c6c6c6ceceeedab56bafbefaeab43b3b58c9b5b5b5dedede070e1c8044179559b0b2b2ba7fffbe2e9d502d989999f9f9f9d5d4d4bcf1c61b861cae2cf6853248d29c3973e697bffc256c7a3592b93c1ecfd2d2525775ed7d2fcccaead5ab2f5fbe6c6363f3c20b2f6844d3535353d5d5d5c1c1c1865c0e03ba6cdfbefddcb97313131390b207a64e4b145353d3b7de7aebabafbe4251d4c4c444a9543e78f0c0c7c7477f5fd4f984ae8bb59696961a9767602e2f5fbecc66b3d50fe9c0f0848585555454e84a7040cc303030c0e7f3376fdeeceeeeae1ecf71b9dc19afb982da3d3c3ce0851443b0a8c4c2719cc160b4b7b7373535418e0e922eea0977032f5b5200fdae5dbb163211ea6773a0be8282828d1b37c211a12117ee2049f6ce3befe4e5e58111d295c50187181414646d6d9d9999393939f9edb7dfae5bb7cedcdc7cc60d9afebbda6c367b686888dac7505bdd92929298981875b344ed5257ae5c595858a89179a1eee5b5b5b51515150507076fd8b0010e3028d5f1783cf5ab57da80658fa2a842a1181b1bd3afc0ff1da381e57e3c60818e8f8f1f3a7428212181c964ea71168603a860656505770aa8d00126e3c99327bdbdbd6fbdf596e12f198343dcb8712383c1282c2c84fd9d2ea280858b8e8e7677774f484878f2e4c97befbda76e0f660ba8e8e9e9d9dddd4df97ab04f393939d1d1d1da6f8d4340fdf1c71fb7b7b79f3d7b563be1842048565696bfbf7f6868a8c6713b499210334d4e4eaa57814ea9db9a262626c3c3c34545451f7cf081a1b782f51f25ce176093259148626363e140d7f0b3ea1901fe343737b7a5a585ea0b1e1616169e3b77cec0fb5bead2120431323212121272faf4e919abc358e472f99c2f5169f7fef9e79f4ba552ea617676f6c9932775490255701cdfb76f5f5555152512fc6d6e6e5eb76eddd8d898f6d93614d8bf7f7f7f7fbfae4b602291e8ecd9b33b76ecb875eb96e117c5162f7897c964870e1d0a0f0ff7f5f55d88d7b0783c1e4400609c1004191919b973e74e5a5ada6c5fe880dc2397cb7573739b76dbaf01f08986fc0285210093f0c61b6f646464a4a5a5c9e5f223478e40764397dec0b6210892949404e7a76054a883c59c9c1c369bad7d7d0a0a5858588c8c8c38393951d775babbbb1f3f7edcdfdf0f2f77383939a5a6a6420b061ae3c57bef16c3301cc7592cd6bcfff405a8636c6c8c244948434030979191f1ff9e620e3d922439313171f0e0c18484043333b3457ebb0104aea8a8e0f3f96c36dbd7d7f7dd77df35641433de40d478084cfdd7bffec5e3f1424242e0a352a92c2a2a6230180e0e0e9e9e9e363636ea52193884c57ea17b117e5005ba282d2d6d6f6fffecb3cfe66c1de572f9e8e8281cc92dfe6b33d46ab1b0b080b069ce7a03df346d75ed35a93d52eafae8ac94b04c7e2940fddd430441eaeaeacacaca525353a963c11fd3ec52bd8ca5716f7891bbd678eb6bb6583ec402000f7a7b7be1673c7f242d96fc55f7c51160217e8673b9114b1d4b4e8bffcb583ebf4da88ee5f74bb8cf1d9627b18c945a723c5bbf7067c4b2819158462c088cc4326241602496110b0223b18c5810188965c482c0482c231604ff3f0000ffffa9cab5d5b0bdd6540000000049454e44ae426082"
  }
```

### As A File

As you can see, that clutters up your JSON and you might prefer to convert the image directly to a file.

Just add the `-sigFile` argument as shown below.  The hex will go away and the image will be written to the
file you specify.  This example is shown without the -sigWidth option to show
how a signature image can be captured in full resolution.

```
./blockchyp -type=charge -terminal="Test Terminal" -amount="25.00" -sigFormat=jpeg -sigFile="sig.jpeg"
  {
    "responseDescription":"Approved",
    "transactionId":"ITAK4GZFWUI6TONDAJBKYEIAAM",
    "transactionRef":"d940770e8b1aacc7ea831f53a5311558afc40da0f50f0517ab762a4c820d84b1",
    "transactionType":"charge",
    "timestamp":"2019-02-01T00:20:16Z",
    "tickBlock":"000e5faa310afb68fb78792d6728b55a07fb5476cd37a33a86154765b4905201",
    "test":false,
    "approved":true,
    "authCode":"086848",
    "entryMethod":"SWIPE",
    "paymentType":"MC",
    "maskedPan":"************5118",
    "currencyCode":"USD",
    "requestedAmount":"25.00",
    "authorizedAmount":"25.00",
    "receiptSuggestions":{
      "merchantName":"Test Merchant",
      "maskedPan":"************5118",
      "authorizedAmount":"25.00",
      "transactionType":"charge",
      "entryMethod":"SWIPE"
    },
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
