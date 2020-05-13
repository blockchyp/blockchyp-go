package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	blockchyp "github.com/blockchyp/blockchyp-go"
)

var validSignatureFormats = []string{
	"gif",
	"jpeg",
	"jpg",
	"png",
}

/*
Constants for various time/date formats.
*/
const (
	ShortDateFormat = "01/02/2006"
	HTMLDateFormat  = "2006-01-02"
	ISOFormat       = "2006-01-02T15:04:05Z0700"
)

var config *blockchyp.ConfigSettings

func main() {

	commandLineArgs := parseArgs()

	loadConfig(commandLineArgs)

	processCommand(commandLineArgs)

}

// loadConfig loads configuration from disk.
func loadConfig(args blockchyp.CommandLineArguments) {
	c, err := blockchyp.LoadConfigSettings(args)
	if err != nil {
		fatalErrorf("Failed to load configuration: %+v", err)
	}

	config = c
}

func parseArgs() blockchyp.CommandLineArguments {

	args := blockchyp.CommandLineArguments{}

	flag.StringVar(&args.Type, "type", "", "transaction type")
	flag.StringVar(&args.ConfigFile, "f", "", "config location")
	flag.StringVar(&args.GatewayHost, "gateway", "", "gateway host address")
	flag.StringVar(&args.TestGatewayHost, "testGateway", "", "test gateway host address")
	flag.StringVar(&args.APIKey, "apiKey", "", "api key")
	flag.StringVar(&args.BearerToken, "bearerToken", "", "bearer token")
	flag.StringVar(&args.SigningKey, "signingKey", "", "signing key")
	flag.StringVar(&args.TransactionRef, "txRef", "", "transaction reference")
	flag.StringVar(&args.TerminalName, "terminal", "", "terminal name")
	flag.StringVar(&args.Token, "token", "", "reusable token")
	flag.StringVar(&args.Amount, "amount", "", "requested tx amount")
	flag.StringVar(&args.TipAmount, "tip", "0.00", "tip amount")
	flag.StringVar(&args.TaxAmount, "tax", "0.00", "tax amount")
	flag.BoolVar(&args.TaxExempt, "taxExempt", false, "tax exempt flag")
	flag.StringVar(&args.CurrencyCode, "currency", "USD", "currency code")
	flag.StringVar(&args.TransactionID, "tx", "", "transaction id")
	flag.StringVar(&args.Description, "desc", "", "transaction description")
	flag.BoolVar(&args.Test, "test", false, "sets test mode")
	flag.BoolVar(&args.PromptForTip, "promptForTip", false, "prompt for tip flag")
	flag.BoolVar(&args.ManualEntry, "manual", false, "key in card data manually")
	flag.BoolVar(&args.HTTPS, "secure", true, "enables or disables https with terminal")
	flag.BoolVar(&args.Version, "version", false, "print version and exit")
	flag.StringVar(&args.Message, "message", "", "short message to be displayed on the terminal")
	flag.BoolVar(&args.EBT, "ebt", false, "EBT transaction")
	flag.BoolVar(&args.Debit, "debit", false, "process card as a debit card")
	flag.StringVar(&args.RouteCache, "routeCache", "", "specifies local file location for route cache")
	flag.StringVar(&args.OutputFile, "out", "", "directs output to a file instead of stdout")
	flag.StringVar(&args.SigFormat, "sigFormat", "", "format for signature file (jpeg, png, gif)")
	flag.IntVar(&args.SigWidth, "sigWidth", -1, "optional width in pixels the signature file should be scaled to")
	flag.StringVar(&args.SigFile, "sigFile", "", "optional location to output sig file")
	flag.StringVar(&args.DisplayTotal, "displayTotal", "", "grand total for line item display")
	flag.StringVar(&args.DisplayTax, "displayTax", "", "tax for line item display")
	flag.StringVar(&args.DisplaySubtotal, "displaySubtotal", "", "subtotal for line item display")
	flag.StringVar(&args.LineItemID, "lineItemId", "", "line item id")
	flag.StringVar(&args.LineItemDescription, "lineItemDescription", "", "line item description")
	flag.StringVar(&args.LineItemPrice, "lineItemPrice", "", "line item price")
	flag.StringVar(&args.LineItemQty, "lineItemQty", "", "line item qty")
	flag.StringVar(&args.LineItemExtended, "lineItemExtended", "", "line item extended total")
	flag.StringVar(&args.LineItemDiscountDescription, "lineItemDiscountDescription", "", "line item discount description")
	flag.StringVar(&args.LineItemDiscountAmount, "lineItemDiscountAmount", "", "line item discount description")
	flag.StringVar(&args.Prompt, "prompt", "", "prompt for boolean or text prompts")
	flag.StringVar(&args.YesCaption, "yesCaption", "Yes", "caption for the 'yes' button")
	flag.StringVar(&args.NoCaption, "noCaption", "No", "caption for the 'no' button")
	flag.StringVar(&args.PromptType, "promptType", "", "type of prompt: email, phone, customer-number, rewards-number")
	flag.StringVar(&args.TCAlias, "tcAlias", "", "alias for a terms and conditions template")
	flag.StringVar(&args.TCName, "tcName", "", "optional name for a terms and conditions template")
	flag.StringVar(&args.TCContent, "tcContent", "", "raw content for the terms and conditions, plain text")
	flag.BoolVar(&args.SigRequired, "sigRequired", true, "optional flag that indicates whether signatures are required, defaults to true")
	flag.IntVar(&args.Timeout, "timeout", 90, "overrides default timeouts for terminal interaction")
	flag.BoolVar(&args.CashBackEnabled, "cashback", false, "enables cash back transactions")
	flag.BoolVar(&args.Enroll, "enroll", false, "enroll the payment in the token vault")
	flag.BoolVar(&args.DisableSignature, "disableSignature", false, "prevent terminal from prompting for signatures")
	flag.StringVar(&args.CustomerID, "customerId", "", "customer id for existing customer record")
	flag.StringVar(&args.CustomerRef, "customerRef", "", "customer reference")
	flag.StringVar(&args.FirstName, "firstName", "", "customer first name")
	flag.StringVar(&args.LastName, "lastName", "", "customer last name")
	flag.StringVar(&args.CompanyName, "companyName", "", "customer company name")
	flag.StringVar(&args.EMailAddress, "email", "", "customer email address")
	flag.StringVar(&args.SMSNumber, "sms", "", "customer sms or mobile number")
	flag.StringVar(&args.PAN, "pan", "", "primary account number for direct process (use not recommended)")
	flag.StringVar(&args.ExpiryMonth, "expMonth", "", "expiration month for pan, if provided")
	flag.StringVar(&args.ExpiryYear, "expYear", "", "expiration year for pan, if provided")
	flag.StringVar(&args.Subject, "subject", "", "subject for email links if auto sending")
	flag.BoolVar(&args.AutoSend, "autoSend", false, "if true, BlockChyp will send email payment links for you")
	flag.StringVar(&args.OrderRef, "orderRef", "", "your system's order or invoice number")
	flag.StringVar(&args.Query, "query", "", "search string")
	flag.StringVar(&args.CallbackURL, "callbackUrl", "", "optional callback url to which a response is posted for payment links")
	flag.BoolVar(&args.Surcharge, "surcharge", false, "adds fee surcharges to transactions, if eligible.")
	flag.BoolVar(&args.CashDiscount, "cashDiscount", false, "adds a cash discount to transactions, if eligible")
	flag.StringVar(&args.PostalCode, "postalCode", "", "postal code to use for address verification")
	flag.StringVar(&args.Address, "address", "", "street address to use for address verification")
	flag.BoolVar(&args.Cashier, "cashier", false, "indicates that a payment link should be displayed in cashier facing mode")
	flag.StringVar(&args.StartDate, "startDate", "", "start date for filtering history results")
	flag.StringVar(&args.EndDate, "endDate", "", "end date for filtering history results")
	flag.StringVar(&args.BatchID, "batchId", "", "batch id for filtering history results")
	flag.IntVar(&args.MaxResults, "maxResults", 250, "max results for query and history functions")
	flag.IntVar(&args.StartIndex, "startIndex", 0, "start index for paged queries")

	flag.Parse()

	if args.Version {
		fmt.Println(blockchyp.Version)
		os.Exit(0)
	}

	validateArgs(&args)

	return args

}

func validateArgs(args *blockchyp.CommandLineArguments) {
	if args.Type == "" {
		fatalError("-type is required")
	}
}

func resolveCredentials(args blockchyp.CommandLineArguments) (*blockchyp.APICredentials, error) {

	creds := &blockchyp.APICredentials{}

	if args.APIKey != "" {
		creds.APIKey = args.APIKey
		creds.BearerToken = args.BearerToken
		creds.SigningKey = args.SigningKey
	} else {
		creds.APIKey = config.APIKey
		creds.BearerToken = config.BearerToken
		creds.SigningKey = config.SigningKey
	}

	if creds.APIKey == "" {
		fatalError("-apiKey or blockchyp.json file required")
	}
	if creds.BearerToken == "" {
		fatalError("-bearerToken or blockchyp.json file required")
	}
	if creds.SigningKey == "" {
		fatalError("-signingKey or blockchyp.json file required")
	}

	return creds, nil

}

func resolveClient(args blockchyp.CommandLineArguments) (*blockchyp.Client, error) {

	creds, err := resolveCredentials(args)
	if err != nil {
		return nil, err
	}
	client := blockchyp.NewClient(*creds)

	if args.HTTPS {
		client.HTTPS = true
	} else {
		client.HTTPS = config.Secure
	}

	if args.GatewayHost != "" {
		client.GatewayHost = args.GatewayHost
	} else {
		client.GatewayHost = config.GatewayHost
	}

	if args.TestGatewayHost != "" {
		client.TestGatewayHost = args.TestGatewayHost
	} else {
		client.TestGatewayHost = config.TestGatewayHost
	}

	if args.RouteCache != "" {
		client.RouteCache = args.RouteCache
	}

	return &client, nil
}

func processCommand(args blockchyp.CommandLineArguments) {

	client, err := resolveClient(args)

	if err != nil {
		handleFatalError(err)
	}

	switch args.Type {
	case "ping":
		processPing(client, args)
	case "enroll":
		processEnroll(client, args)
	case "charge", "preauth":
		processAuth(client, args)
	case "gift-activate":
		processGiftActivate(client, args)
	case "capture":
		processCapture(client, args)
	case "void":
		processVoid(client, args)
	case "refund":
		processRefund(client, args)
	case "reverse":
		processReverse(client, args)
	case "close-batch":
		processCloseBatch(client, args)
	case "message":
		processMessage(client, args)
	case "boolean-prompt":
		processBooleanPrompt(client, args)
	case "text-prompt":
		processTextPrompt(client, args)
	case "clear":
		processClear(client, args)
	case "balance":
		processBalance(client, args)
	case "display":
		processDisplay(client, args)
	case "tc":
		processTermsAndConditions(client, args)
	case "cache":
		processCache(client, args)
	case "cache-expire":
		processCacheExpire(client, args)
	case "terminal-status":
		processTerminalStatus(client, args)
	case "capture-signature":
		processCaptureSignature(client, args)
	case "send-link":
		processSendLink(client, args)
	case "get-customer":
		getCustomer(client, args)
	case "search-customer":
		searchCustomer(client, args)
	case "update-customer":
		updateCustomer(client, args)
	case "tx-status":
		processTransactionStatus(client, args)
	case "cash-discount":
		processCashDiscount(client, args)
	case "batch-history":
		processBatchHistory(client, args)
	case "batch-details":
		processBatchDetails(client, args)
	case "tx-history":
		processTransactionHistory(client, args)
	default:
		fatalErrorf("unknown transaction type: %s", args.Type)
	}

}

func processBatchHistory(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	request := blockchyp.BatchHistoryRequest{
		MaxResults: args.MaxResults,
		StartIndex: args.StartIndex,
	}

	if args.StartDate != "" {
		parsedDate, err := parseTimestamp(args.StartDate)
		if err != nil {
			handleError(&args, err)
		}
		request.StartDate = parsedDate
	}
	if args.EndDate != "" {
		parsedDate, err := parseTimestamp(args.EndDate)
		if err != nil {
			handleError(&args, err)
		}
		request.EndDate = parsedDate
	}

	ack, err := client.BatchHistory(request)
	if err != nil {
		handleError(&args, err)
	}

	dumpResponse(&args, ack)

}

func processBatchDetails(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	if args.BatchID == "" {
		fatalErrorf("-batchId is required")
	}

	request := blockchyp.BatchDetailsRequest{
		BatchID: args.BatchID,
	}

	ack, err := client.BatchDetails(request)
	if err != nil {
		handleError(&args, err)
	}

	dumpResponse(&args, ack)

}

func processTransactionHistory(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	request := blockchyp.TransactionHistoryRequest{
		MaxResults:   args.MaxResults,
		StartIndex:   args.StartIndex,
		BatchID:      args.BatchID,
		TerminalName: args.TerminalName,
	}

	if args.StartDate != "" {
		parsedDate, err := parseTimestamp(args.StartDate)
		if err != nil {
			handleError(&args, err)
		}
		request.StartDate = parsedDate
	}
	if args.EndDate != "" {
		parsedDate, err := parseTimestamp(args.EndDate)
		if err != nil {
			handleError(&args, err)
		}
		request.EndDate = parsedDate
	}

	ack, err := client.TransactionHistory(request)
	if err != nil {
		handleError(&args, err)
	}

	dumpResponse(&args, ack)

}

func processTransactionStatus(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	if args.TransactionID == "" && args.TransactionRef == "" {
		fatalErrorf("-tx or -txRef are required")
	}

	request := blockchyp.TransactionStatusRequest{
		TransactionID:  args.TransactionID,
		TransactionRef: args.TransactionRef,
		Test:           args.Test,
	}

	ack, err := client.TransactionStatus(request)
	if err != nil {
		handleError(&args, err)
	}

	dumpResponse(&args, ack)

}

func processSendLink(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	validateRequired(args.OrderRef, "orderRef")
	validateRequired(args.Amount, "amount")

	if !args.Cashier && !hasCustomerFields(args) {
		fatalErrorf("customer fields (-customerId, -email, etc ) are required")
	}

	request := blockchyp.PaymentLinkRequest{
		TransactionRef: args.TransactionRef,
		Description:    args.Description,
		Subject:        args.Subject,
		Amount:         args.Amount,
		OrderRef:       args.OrderRef,
		Test:           args.Test,
		Timeout:        args.Timeout,
		TaxExempt:      args.TaxExempt,
		Transaction:    assembleDisplayTransaction(args),
		Customer:       *populateCustomer(args),
		AutoSend:       args.AutoSend,
		CallbackURL:    args.CallbackURL,
		TCAlias:        args.TCAlias,
		TCName:         args.TCName,
		TCContent:      args.TCContent,
		Cashier:        args.Cashier,
	}

	ack, err := client.SendPaymentLink(request)
	if err != nil {
		handleError(&args, err)
	}

	dumpResponse(&args, ack)
}

func getCustomer(client *blockchyp.Client, args blockchyp.CommandLineArguments) {
	if args.CustomerID == "" && args.CustomerRef == "" {
		fatalError("-customerId or -customerRef are required")
	}

	res, err := client.Customer(blockchyp.CustomerRequest{
		CustomerID:  args.CustomerID,
		CustomerRef: args.CustomerRef,
		Timeout:     args.Timeout,
	})
	if err != nil {
		handleError(&args, err)
	}

	dumpResponse(&args, res)
}

func searchCustomer(client *blockchyp.Client, args blockchyp.CommandLineArguments) {
	validateRequired(args.Query, "query")

	res, err := client.CustomerSearch(blockchyp.CustomerSearchRequest{
		Query:   args.Query,
		Timeout: args.Timeout,
	})
	if err != nil {
		handleError(&args, err)
	}

	dumpResponse(&args, res)
}

func updateCustomer(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	res, err := client.UpdateCustomer(blockchyp.UpdateCustomerRequest{
		Customer: blockchyp.Customer{
			ID:           args.CustomerID,
			CustomerRef:  args.CustomerRef,
			FirstName:    args.FirstName,
			LastName:     args.LastName,
			CompanyName:  args.CompanyName,
			EmailAddress: args.EMailAddress,
			SmsNumber:    args.SMSNumber,
		},
		Timeout: args.Timeout,
	})
	if err != nil {
		handleError(&args, err)
	}

	dumpResponse(&args, res)
}

func processCacheExpire(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	client.ExpireRouteCache()

}

func processCache(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	fmt.Println("Cache Location:", client.RouteCache)

}

func processTerminalStatus(client *blockchyp.Client, args blockchyp.CommandLineArguments) {
	validateRequired(args.TerminalName, "terminal")

	response, err := client.TerminalStatus(blockchyp.TerminalStatusRequest{
		TerminalName: args.TerminalName,
	})
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, response)
}

func processCaptureSignature(client *blockchyp.Client, args blockchyp.CommandLineArguments) {
	validateRequired(args.TerminalName, "terminal")
	if args.SigFile == "" && args.SigFormat == blockchyp.SignatureFormatNone {
		fatalErrorf("-%s or -%s are required", "sigFile", "sigFormat")
	}

	response, err := client.CaptureSignature(blockchyp.CaptureSignatureRequest{
		TerminalName: args.TerminalName,
		SigFile:      args.SigFile,
		SigFormat:    blockchyp.SignatureFormat(args.SigFormat),
		SigWidth:     args.SigWidth,
	})
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, response)
}

func processBalance(client *blockchyp.Client, args blockchyp.CommandLineArguments) {
	validateRequired(args.TerminalName, "terminal")

	request := blockchyp.BalanceRequest{}
	request.TerminalName = args.TerminalName
	request.ManualEntry = args.ManualEntry
	request.Test = args.Test

	if args.Debit {
		request.CardType = blockchyp.CardTypeDebit
	} else if args.EBT {
		request.CardType = blockchyp.CardTypeEBT
	}

	ack, err := client.Balance(request)
	if err != nil {
		handleError(&args, err)
	}

	dumpResponse(&args, ack)

}

func processCashDiscount(client *blockchyp.Client, args blockchyp.CommandLineArguments) {
	validateRequired(args.Amount, "amount")

	request := blockchyp.CashDiscountRequest{}
	request.Amount = args.Amount
	request.CurrencyCode = args.CurrencyCode
	request.Test = args.Test
	request.Surcharge = args.Surcharge
	request.CashDiscount = args.CashDiscount

	response, err := client.CashDiscount(request)
	if err != nil {
		handleError(&args, err)
	}

	dumpResponse(&args, response)

}

func processClear(client *blockchyp.Client, args blockchyp.CommandLineArguments) {
	validateRequired(args.TerminalName, "terminal")

	request := blockchyp.ClearTerminalRequest{}
	request.TerminalName = args.TerminalName

	ack, err := client.Clear(request)
	if err != nil {
		handleError(&args, err)
	}

	dumpResponse(&args, ack)

}

func processTermsAndConditions(client *blockchyp.Client, args blockchyp.CommandLineArguments) {
	validateRequired(args.TerminalName, "terminal")

	request := blockchyp.TermsAndConditionsRequest{}
	request.TerminalName = args.TerminalName
	request.Timeout = args.Timeout
	request.TCAlias = args.TCAlias
	request.TCName = args.TCName
	request.TCContent = args.TCContent
	request.TransactionID = args.TransactionID
	request.TransactionRef = args.TransactionRef
	request.SigRequired = args.SigRequired
	request.SigWidth = args.SigWidth
	request.SigFile = args.SigFile
	request.SigFormat = blockchyp.SignatureFormat(args.SigFormat)
	request.DisableSignature = args.DisableSignature

	ack, err := client.TermsAndConditions(request)
	if err != nil {
		handleError(&args, err)
	}

	dumpResponse(&args, ack)

}

func assembleDisplayTransaction(args blockchyp.CommandLineArguments) *blockchyp.TransactionDisplayTransaction {

	tx := &blockchyp.TransactionDisplayTransaction{
		Subtotal: args.DisplaySubtotal,
		Tax:      args.DisplayTax,
		Total:    args.DisplayTotal,
	}

	ids := strings.Split(args.LineItemID, "|")
	descs := strings.Split(args.LineItemDescription, "|")
	prices := strings.Split(args.LineItemPrice, "|")
	qtys := strings.Split(args.LineItemQty, "|")
	extendeds := strings.Split(args.LineItemExtended, "|")
	discounts := strings.Split(args.LineItemDiscountDescription, "|")
	discountAmounts := strings.Split(args.LineItemDiscountAmount, "|")

	lines := make([]*blockchyp.TransactionDisplayItem, 0)
	for idx, desc := range descs {
		if desc == "" {
			continue
		}

		line := &blockchyp.TransactionDisplayItem{}
		line.Description = desc

		if len(ids) >= (idx - 1) {
			line.ID = ids[idx]
		}

		if len(qtys) >= (idx - 1) {
			line.Quantity, _ = strconv.ParseFloat(qtys[idx], 64)
		}

		if len(prices) >= (idx - 1) {
			line.Price = prices[idx]
		}

		if len(extendeds) >= (idx - 1) {
			line.Extended = extendeds[idx]
		}

		if len(discounts) >= (idx - 1) {

			discountLines := make([]*blockchyp.TransactionDisplayDiscount, 0)
			discountLine := blockchyp.TransactionDisplayDiscount{}
			discountLine.Description = discounts[idx]
			if len(discountAmounts) >= (idx - 1) {
				discountLine.Amount = discountAmounts[idx]
			}

			discountLines = append(discountLines, &discountLine)
			line.Discounts = discountLines
		}

		lines = append(lines, line)
	}

	tx.Items = lines

	return tx

}

func processDisplay(client *blockchyp.Client, args blockchyp.CommandLineArguments) {
	validateRequired(args.TerminalName, "terminal")

	request := blockchyp.TransactionDisplayRequest{}
	request.TerminalName = args.TerminalName
	request.Transaction = assembleDisplayTransaction(args)

	ack, err := client.UpdateTransactionDisplay(request)
	if err != nil {
		handleError(&args, err)
	}

	dumpResponse(&args, ack)

}

func processMessage(client *blockchyp.Client, args blockchyp.CommandLineArguments) {
	validateRequired(args.Message, "message")
	validateRequired(args.TerminalName, "terminal")

	req := blockchyp.MessageRequest{}
	req.Message = args.Message
	req.TerminalName = args.TerminalName
	req.Test = args.Test

	res, err := client.Message(req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)

}

func processBooleanPrompt(client *blockchyp.Client, args blockchyp.CommandLineArguments) {
	validateRequired(args.Prompt, "prompt")
	validateRequired(args.TerminalName, "terminal")

	req := blockchyp.BooleanPromptRequest{}
	req.Prompt = args.Prompt
	req.TerminalName = args.TerminalName
	req.YesCaption = args.YesCaption
	req.NoCaption = args.NoCaption
	req.Test = args.Test

	res, err := client.BooleanPrompt(req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)

}

func processTextPrompt(client *blockchyp.Client, args blockchyp.CommandLineArguments) {
	validateRequired(args.PromptType, "promptType")
	validateRequired(args.TerminalName, "terminal")

	req := blockchyp.TextPromptRequest{}
	req.PromptType = blockchyp.PromptType(args.PromptType)
	req.TerminalName = args.TerminalName
	req.Test = args.Test

	res, err := client.TextPrompt(req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)

}

func processRefund(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := blockchyp.RefundRequest{}
	req.TransactionRef = args.TransactionRef
	req.TransactionID = args.TransactionID
	req.Amount = args.Amount
	req.TerminalName = args.TerminalName
	req.Token = args.Token
	req.PostalCode = args.PostalCode
	req.Address = args.Address

	if args.Debit {
		req.CardType = blockchyp.CardTypeDebit
	} else if args.EBT {
		req.CardType = blockchyp.CardTypeEBT
		// EBT free range refunds are not permitted.
		req.TerminalName = args.TerminalName
	}

	req.Test = args.Test
	req.ManualEntry = args.ManualEntry
	req.SigWidth = args.SigWidth
	req.SigFile = args.SigFile
	req.SigFormat = blockchyp.SignatureFormat(args.SigFormat)
	req.DisableSignature = args.DisableSignature

	res, err := client.Refund(req)

	if err != nil {
		handleError(&args, err)
	}

	dumpResponse(&args, res)
}

func processReverse(client *blockchyp.Client, args blockchyp.CommandLineArguments) {
	validateRequired(args.TransactionRef, "txRef")
	req := blockchyp.AuthorizationRequest{}
	req.TransactionRef = args.TransactionRef
	req.Test = args.Test
	req.ManualEntry = args.ManualEntry

	if args.Debit {
		req.CardType = blockchyp.CardTypeDebit
	} else if args.EBT {
		req.CardType = blockchyp.CardTypeEBT
	}

	res, err := client.Reverse(req)

	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)
}

func processCloseBatch(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := blockchyp.CloseBatchRequest{}
	req.TransactionRef = args.TransactionRef
	req.Test = args.Test

	res, err := client.CloseBatch(req)

	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)
}

func processVoid(client *blockchyp.Client, args blockchyp.CommandLineArguments) {
	validateRequired(args.TransactionID, "tx")
	req := blockchyp.VoidRequest{}
	req.TransactionRef = args.TransactionRef
	req.TransactionID = args.TransactionID
	req.Test = args.Test

	res, err := client.Void(req)

	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)
}

func processCapture(client *blockchyp.Client, args blockchyp.CommandLineArguments) {
	validateRequired(args.TransactionID, "tx")
	req := blockchyp.CaptureRequest{}
	req.TransactionRef = args.TransactionRef
	req.Amount = args.Amount
	req.TransactionID = args.TransactionID
	req.TipAmount = args.TipAmount
	req.TaxAmount = args.TaxAmount
	req.Test = args.Test

	res, err := client.Capture(req)

	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)
}

func processGiftActivate(client *blockchyp.Client, args blockchyp.CommandLineArguments) {
	validateRequired(args.Amount, "amount")
	validateRequired(args.TerminalName, "terminal")
	req := blockchyp.GiftActivateRequest{}
	req.TerminalName = args.TerminalName
	req.TransactionRef = args.TransactionRef
	req.Amount = args.Amount
	req.Test = args.Test

	res, err := client.GiftActivate(req)

	if err != nil {
		handleFatalError(err)
	}
	dumpResponse(&args, res)
}

func hasCustomerFields(args blockchyp.CommandLineArguments) bool {

	if args.CustomerID != "" {
		return true
	} else if args.CustomerRef != "" {
		return true
	} else if args.FirstName != "" {
		return true
	} else if args.LastName != "" {
		return true
	} else if args.CompanyName != "" {
		return true
	} else if args.EMailAddress != "" {
		return true
	} else if args.SMSNumber != "" {
		return true
	}

	return false

}

func populateCustomer(args blockchyp.CommandLineArguments) *blockchyp.Customer {

	return &blockchyp.Customer{
		ID:           args.CustomerID,
		CustomerRef:  args.CustomerRef,
		FirstName:    args.FirstName,
		LastName:     args.LastName,
		CompanyName:  args.CompanyName,
		EmailAddress: args.EMailAddress,
		SmsNumber:    args.SMSNumber,
	}

}

func processEnroll(client *blockchyp.Client, args blockchyp.CommandLineArguments) {
	if (args.TerminalName == "") && (args.Token == "") && (args.PAN == "") {
		fatalError("-terminal or -token requred")
	}
	req := blockchyp.EnrollRequest{}
	req.TerminalName = args.TerminalName
	req.TransactionRef = args.TransactionRef
	req.Test = args.Test
	req.ManualEntry = args.ManualEntry
	req.PAN = args.PAN
	req.ExpMonth = args.ExpiryMonth
	req.ExpYear = args.ExpiryYear
	req.PostalCode = args.PostalCode
	req.Address = args.Address
	if hasCustomerFields(args) {
		req.Customer = populateCustomer(args)
	}

	res, err := client.Enroll(req)

	if err != nil {
		handleError(&args, err)
	}

	dumpResponse(&args, res)
}

func processAuth(client *blockchyp.Client, args blockchyp.CommandLineArguments) {
	validateRequired(args.Amount, "amount")
	if (args.TerminalName == "") && (args.Token == "") && (args.PAN == "") {
		fatalError("-terminal or -token requred")
	}
	req := blockchyp.AuthorizationRequest{}
	req.TerminalName = args.TerminalName
	req.TransactionRef = args.TransactionRef
	req.Token = args.Token
	req.PAN = args.PAN
	req.ExpMonth = args.ExpiryMonth
	req.ExpYear = args.ExpiryYear
	req.Description = args.Description
	req.Amount = args.Amount
	req.PromptForTip = args.PromptForTip
	req.TaxAmount = args.TaxAmount
	req.TipAmount = args.TipAmount
	req.Test = args.Test
	req.Enroll = args.Enroll
	req.ManualEntry = args.ManualEntry
	req.SigWidth = args.SigWidth
	req.SigFile = args.SigFile
	req.SigFormat = blockchyp.SignatureFormat(args.SigFormat)
	req.DisableSignature = args.DisableSignature
	req.CashBackEnabled = args.CashBackEnabled
	req.Surcharge = args.Surcharge
	req.CashDiscount = args.CashDiscount
	req.PostalCode = args.PostalCode
	req.Address = args.Address

	if args.Debit {
		req.CardType = blockchyp.CardTypeDebit
	} else if args.EBT {
		req.CardType = blockchyp.CardTypeEBT
	}
	if hasCustomerFields(args) {
		req.Customer = populateCustomer(args)
	}

	res := &blockchyp.AuthorizationResponse{}
	var err error
	switch args.Type {
	case "charge":
		res, err = client.Charge(req)
	case "preauth":
		res, err = client.Preauth(req)
	}

	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)
}

func processPing(client *blockchyp.Client, args blockchyp.CommandLineArguments) {
	validateRequired(args.TerminalName, "terminal")
	req := blockchyp.PingRequest{
		TerminalName: args.TerminalName,
	}
	res, err := client.Ping(req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)
}

func parseTimestamp(ts string) (time.Time, error) {

	parsedResult, err := parseTimestampWithFormat(ts, ISOFormat)
	if err == nil {
		return parsedResult, nil
	}
	parsedResult, err = parseTimestampWithFormat(ts, ShortDateFormat)
	if err == nil {
		return parsedResult, nil
	}
	return parseTimestampWithFormat(ts, HTMLDateFormat)

}

func parseTimestampWithFormat(ts string, format string) (time.Time, error) {

	return time.Parse(format, ts)

}

func validateRequired(value string, arg string) {
	if value == "" {
		fatalErrorf("-%s is required", arg)
	}
}

func dumpResponse(args *blockchyp.CommandLineArguments, res interface{}) {

	content, err := json.Marshal(res)
	if err != nil {
		handleFatalError(err)
	}
	if args.OutputFile != "" {
		err := ioutil.WriteFile(args.OutputFile, content, 0644)
		if err != nil {
			fmt.Print(err)
		}
	} else {
		fmt.Println(string(content))
	}

}

func handleError(args *blockchyp.CommandLineArguments, err error) {

	ack := blockchyp.Acknowledgement{}
	ack.Error = err.Error()
	dumpResponse(args, ack)
	handleFatal()

}

func handleFatalError(err error) {

	fmt.Println(err)
	handleFatal()

}

func fatalError(msg string) {

	handleFatalError(errors.New(msg))

}

func fatalErrorf(format string, args ...interface{}) {

	handleFatalError(fmt.Errorf(format, args...))

}

func handleFatal() {

	os.Exit(1)

}
