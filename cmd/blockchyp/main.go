package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
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

	flag.StringVar(&args.Type, "type", "", "transaction type (deprecated, use cmd instead)")
	flag.StringVar(&args.Command, "cmd", "", "command")
	flag.StringVar(&args.ConfigFile, "f", "", "config location")
	flag.StringVar(&args.GatewayHost, "gateway", "", "gateway host address")
	flag.StringVar(&args.DashboardHost, "dashboard", "", "dashboard host address")
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
	flag.IntVar(&args.Timeout, "timeout", 120, "overrides default timeouts for terminal interaction")
	flag.BoolVar(&args.WaitForRemovedCard, "waitForRemovedCard", false, "causes the request to block until all cards have been removed from the card reader.")
	flag.BoolVar(&args.Force, "force", false, "overrides any in-progress transactions.")
	flag.BoolVar(&args.CashBackEnabled, "cashback", false, "enables cash back transactions")
	flag.BoolVar(&args.Enroll, "enroll", false, "enroll the payment in the token vault")
	flag.BoolVar(&args.EnrollOnly, "enrollOnly", false, "use to make a cashier facing payment link enroll only")
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
	flag.IntVar(&args.MaxResults, "maxResults", 50, "max results for query and history functions")
	flag.IntVar(&args.StartIndex, "startIndex", 0, "start index for paged queries")
	flag.BoolVar(&args.Queue, "queue", false, "queue transaction without running it")
	flag.BoolVar(&args.Async, "async", false, "run transaction asynchronously and don't wait for the response")
	flag.BoolVar(&args.LogRequests, "logRequests", false, "log full http request for API calls")
	flag.StringVar(&args.LinkCode, "linkCode", "", "payment link code")
	flag.StringVar(&args.Cryptocurrency, "crypto", "", "crypto currency code for crypto transaction")
	flag.StringVar(&args.CryptoNetwork, "cryptoNetwork", "L1", "optional network code for crypto currency (L1 or L2)")
	flag.StringVar(&args.CryptoReceiveAddress, "receiveAddress", "", "destination address for cryptocurrency transactions")
	flag.StringVar(&args.Label, "label", "", "optional label for cryptocurrency transactions")
	flag.StringVar(&args.DBAName, "dbaName", "", "dba name for merchant account commands")
	flag.StringVar(&args.MerchantID, "merchantId", "", "merchant id for partner and org related apis")
	flag.StringVar(&args.TerminalID, "terminalId", "", "terminal id for terminal related operations")
	flag.StringVar(&args.Code, "code", "", "code for use with commands like terminal activation")
	flag.StringVar(&args.TemplateID, "templateId", "", "template id for terms and conditions template operations")
	flag.StringVar(&args.LogEntryID, "logEntryId", "", "log entry id for terms and conditions operations")
	flag.StringVar(&args.QuestionID, "questionId", "", "question id for survey question related operations")
	flag.StringVar(&args.QuestionType, "questionType", "", "question type for survey question related operations")
	flag.StringVar(&args.QuestionText, "questionText", "", "question text for survey question related operations")
	flag.BoolVar(&args.Enabled, "enabled", false, "enabled flag for various update operations")
	flag.IntVar(&args.Ordinal, "ordinal", 0, "ordinal value used to specify sort order for certain update operations.")
	flag.StringVar(&args.File, "file", "", "is a file name for file upload operations")
	flag.StringVar(&args.UploadID, "uploadId", "", "upload id to be used for tracking upload progress")

	flag.Parse()

	if args.Version {
		fmt.Println(blockchyp.Version)
		os.Exit(0)
	}

	validateArgs(&args)

	return args

}

func validateArgs(args *blockchyp.CommandLineArguments) {
	if args.Type == "" && args.Command == "" {
		fatalError("-cmd is required")
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
	client.LogRequests = args.LogRequests

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

	if args.DashboardHost != "" {
		client.DashboardHost = args.DashboardHost
	} else {
		client.DashboardHost = config.DashboardHost
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

	cmd := args.Command

	if cmd == "" {
		cmd = args.Type
	}

	switch cmd {
	case "add-test-merchant":
		processAddTestMerchant(client, args)
	case "delete-test-merchant":
		processDeleteTestMerchant(client, args)
	case "invite-merchant-user":
		processInviteMerchantUser(client, args)
	case "get-merchants":
		processGetMerchants(client, args)
	case "merchant-users":
		processMerchantUsers(client, args)
	case "terminals":
		processTerminals(client, args)
	case "deactivate-terminal":
		processDeactivateTerminal(client, args)
	case "activate-terminal":
		processActivateTerminal(client, args)
	case "update-tc-template":
		processUpdateTCTemplate(client, args)
	case "tc-templates":
		processTCTemplates(client, args)
	case "tc-template":
		processTCTemplate(client, args)
	case "delete-tc-template":
		processDeleteTCTemplate(client, args)
	case "tc-log":
		processTCLog(client, args)
	case "tc-entry":
		processTCEntry(client, args)
	case "survey-questions":
		processSurveyQuestions(client, args)
	case "survey-question":
		processSurveyQuestion(client, args)
	case "survey-results":
		processSurveyResults(client, args)
	case "update-survey-question":
		processUpdateSurveyQuestion(client, args)
	case "delete-survey-question":
		processDeleteSurveyQuestion(client, args)
	case "upload-media":
		processUploadMedia(client, args)
	case "ping":
		processPing(client, args)
	case "locate":
		processLocate(client, args)
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
	case "cancel-link":
		processCancelLink(client, args)
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
	case "merchant-profile":
		processMerchantProfile(client, args)
	case "list-queue":
		processQueueList(client, args)
	case "delete-queue":
		processQueueDelete(client, args)
	case "delete-customer":
		processCustomerDelete(client, args)
	case "delete-token":
		processTokenDelete(client, args)
	case "token-metadata":
		processTokenMetadata(client, args)
	case "link-token":
		processLinkToken(client, args)
	case "unlink-token":
		processUnlinkToken(client, args)
	default:
		fatalErrorf("unknown transaction type: %s", args.Type)
	}

}

func processUnlinkToken(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	request := blockchyp.UnlinkTokenRequest{
		Token:      args.Token,
		CustomerID: args.CustomerID,
	}

	ack, err := client.UnlinkToken(request)
	if err != nil {
		handleError(&args, err)
	}

	dumpResponse(&args, ack)

}

func processLinkToken(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	request := blockchyp.LinkTokenRequest{
		Token:      args.Token,
		CustomerID: args.CustomerID,
	}

	ack, err := client.LinkToken(request)
	if err != nil {
		handleError(&args, err)
	}

	dumpResponse(&args, ack)

}

func processTokenMetadata(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	request := blockchyp.TokenMetadataRequest{
		Token: args.Token,
	}

	ack, err := client.TokenMetadata(request)
	if err != nil {
		handleError(&args, err)
	}

	dumpResponse(&args, ack)

}

func processMerchantProfile(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	request := blockchyp.MerchantProfileRequest{
		Test:       args.Test,
		MerchantID: args.MerchantID,
	}

	ack, err := client.MerchantProfile(request)
	if err != nil {
		handleError(&args, err)
	}

	dumpResponse(&args, ack)

}

func processBatchHistory(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	request := blockchyp.BatchHistoryRequest{
		MaxResults: args.MaxResults,
		StartIndex: args.StartIndex,
		Test:       args.Test,
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
		Test:    args.Test,
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
		Test:         args.Test,
		Query:        args.Query,
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

	if !args.EnrollOnly {
		validateRequired(args.OrderRef, "orderRef")
		validateRequired(args.Amount, "amount")
	}

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
		Enroll:         args.Enroll,
		EnrollOnly:     args.EnrollOnly,
	}

	if args.Cryptocurrency != "" {
		request.Cryptocurrency = &args.Cryptocurrency
		if args.CryptoNetwork != "" {
			request.CryptoNetwork = &args.CryptoNetwork
		}
		if args.CryptoReceiveAddress != "" {
			request.CryptoReceiveAddress = &args.CryptoReceiveAddress
		}
		if args.Label != "" {
			request.PaymentRequestLabel = &args.Label
		}
		if args.Message != "" {
			request.PaymentRequestMessage = &args.Message
		}
	}

	if args.EnrollOnly {
		request.Enroll = true
	}

	ack, err := client.SendPaymentLink(request)
	if err != nil {
		handleError(&args, err)
	}

	dumpResponse(&args, ack)
}

func processCancelLink(client *blockchyp.Client, args blockchyp.CommandLineArguments) {
	validateRequired(args.LinkCode, "linkCode")

	request := blockchyp.CancelPaymentLinkRequest{
		Test:     args.Test,
		Timeout:  args.Timeout,
		LinkCode: args.LinkCode,
	}

	res, err := client.CancelPaymentLink(request)
	if err != nil {
		handleError(&args, err)
	}

	dumpResponse(&args, res)
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
		TerminalName:       args.TerminalName,
		Timeout:            args.Timeout,
		WaitForRemovedCard: args.WaitForRemovedCard,
		Force:              args.Force,
		Test:               args.Test,
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
		TerminalName:       args.TerminalName,
		SigFile:            args.SigFile,
		SigFormat:          blockchyp.SignatureFormat(args.SigFormat),
		SigWidth:           args.SigWidth,
		Test:               args.Test,
		Timeout:            args.Timeout,
		WaitForRemovedCard: args.WaitForRemovedCard,
		Force:              args.Force,
	})
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, response)
}

func processBalance(client *blockchyp.Client, args blockchyp.CommandLineArguments) {
	validateRequired(args.TerminalName, "terminal")

	request := blockchyp.BalanceRequest{
		ManualEntry:        args.ManualEntry,
		TerminalName:       args.TerminalName,
		Timeout:            args.Timeout,
		WaitForRemovedCard: args.WaitForRemovedCard,
		Force:              args.Force,
		Test:               args.Test,
	}

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

	request := blockchyp.CashDiscountRequest{
		Amount:       args.Amount,
		CashDiscount: args.CashDiscount,
		CurrencyCode: args.CurrencyCode,
		Surcharge:    args.Surcharge,
		Test:         args.Test,
		Timeout:      args.Timeout,
	}

	response, err := client.CashDiscount(request)
	if err != nil {
		handleError(&args, err)
	}

	dumpResponse(&args, response)

}

func processClear(client *blockchyp.Client, args blockchyp.CommandLineArguments) {
	validateRequired(args.TerminalName, "terminal")

	request := blockchyp.ClearTerminalRequest{
		TerminalName:       args.TerminalName,
		Timeout:            args.Timeout,
		WaitForRemovedCard: args.WaitForRemovedCard,
		Force:              args.Force,
		Test:               args.Test,
	}

	ack, err := client.Clear(request)
	if err != nil {
		handleError(&args, err)
	}

	dumpResponse(&args, ack)

}

func processTermsAndConditions(client *blockchyp.Client, args blockchyp.CommandLineArguments) {
	validateRequired(args.TerminalName, "terminal")

	request := blockchyp.TermsAndConditionsRequest{
		DisableSignature:   args.DisableSignature,
		SigFile:            args.SigFile,
		SigFormat:          blockchyp.SignatureFormat(args.SigFormat),
		SigRequired:        args.SigRequired,
		SigWidth:           args.SigWidth,
		TCAlias:            args.TCAlias,
		TCContent:          args.TCContent,
		TCName:             args.TCName,
		TerminalName:       args.TerminalName,
		Timeout:            args.Timeout,
		WaitForRemovedCard: args.WaitForRemovedCard,
		Force:              args.Force,
		TransactionID:      args.TransactionID,
		TransactionRef:     args.TransactionRef,
		Test:               args.Test,
	}

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

	request := blockchyp.TransactionDisplayRequest{
		TerminalName:       args.TerminalName,
		Transaction:        assembleDisplayTransaction(args),
		Timeout:            args.Timeout,
		WaitForRemovedCard: args.WaitForRemovedCard,
		Force:              args.Force,
		Test:               args.Test,
	}

	ack, err := client.UpdateTransactionDisplay(request)
	if err != nil {
		handleError(&args, err)
	}

	dumpResponse(&args, ack)

}

func processMessage(client *blockchyp.Client, args blockchyp.CommandLineArguments) {
	validateRequired(args.Message, "message")
	validateRequired(args.TerminalName, "terminal")

	req := blockchyp.MessageRequest{
		Message:            args.Message,
		TerminalName:       args.TerminalName,
		Timeout:            args.Timeout,
		WaitForRemovedCard: args.WaitForRemovedCard,
		Force:              args.Force,
		Test:               args.Test,
	}

	res, err := client.Message(req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)

}

func processBooleanPrompt(client *blockchyp.Client, args blockchyp.CommandLineArguments) {
	validateRequired(args.Prompt, "prompt")
	validateRequired(args.TerminalName, "terminal")

	req := blockchyp.BooleanPromptRequest{
		NoCaption:          args.NoCaption,
		Prompt:             args.Prompt,
		TerminalName:       args.TerminalName,
		Test:               args.Test,
		Timeout:            args.Timeout,
		WaitForRemovedCard: args.WaitForRemovedCard,
		Force:              args.Force,
		YesCaption:         args.YesCaption,
	}

	res, err := client.BooleanPrompt(req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)

}

func processTextPrompt(client *blockchyp.Client, args blockchyp.CommandLineArguments) {
	validateRequired(args.PromptType, "promptType")
	validateRequired(args.TerminalName, "terminal")

	req := blockchyp.TextPromptRequest{
		PromptType:         blockchyp.PromptType(args.PromptType),
		TerminalName:       args.TerminalName,
		Timeout:            args.Timeout,
		WaitForRemovedCard: args.WaitForRemovedCard,
		Force:              args.Force,
		Test:               args.Test,
	}

	res, err := client.TextPrompt(req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)

}

func processRefund(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := blockchyp.RefundRequest{
		Address:            args.Address,
		Amount:             args.Amount,
		DisableSignature:   args.DisableSignature,
		ManualEntry:        args.ManualEntry,
		PostalCode:         args.PostalCode,
		SigFile:            args.SigFile,
		SigFormat:          blockchyp.SignatureFormat(args.SigFormat),
		SigWidth:           args.SigWidth,
		TerminalName:       args.TerminalName,
		Test:               args.Test,
		Timeout:            args.Timeout,
		WaitForRemovedCard: args.WaitForRemovedCard,
		Force:              args.Force,
		Token:              args.Token,
		TransactionID:      args.TransactionID,
		TransactionRef:     args.TransactionRef,
	}

	if args.Debit {
		req.CardType = blockchyp.CardTypeDebit
	} else if args.EBT {
		req.CardType = blockchyp.CardTypeEBT
		// EBT free range refunds are not permitted.
		req.TerminalName = args.TerminalName
	}

	res, err := client.Refund(req)

	if err != nil {
		handleError(&args, err)
	}

	dumpResponse(&args, res)
}

func processReverse(client *blockchyp.Client, args blockchyp.CommandLineArguments) {
	validateRequired(args.TransactionRef, "txRef")
	req := blockchyp.AuthorizationRequest{
		ManualEntry:        args.ManualEntry,
		Test:               args.Test,
		Timeout:            args.Timeout,
		WaitForRemovedCard: args.WaitForRemovedCard,
		Force:              args.Force,
		TransactionRef:     args.TransactionRef,
	}

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

	req := blockchyp.CloseBatchRequest{
		Test:               args.Test,
		Timeout:            args.Timeout,
		WaitForRemovedCard: args.WaitForRemovedCard,
		Force:              args.Force,
		TransactionRef:     args.TransactionRef,
	}

	res, err := client.CloseBatch(req)

	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)
}

func processVoid(client *blockchyp.Client, args blockchyp.CommandLineArguments) {
	validateRequired(args.TransactionID, "tx")
	req := blockchyp.VoidRequest{
		Test:               args.Test,
		Timeout:            args.Timeout,
		WaitForRemovedCard: args.WaitForRemovedCard,
		Force:              args.Force,
		TransactionID:      args.TransactionID,
		TransactionRef:     args.TransactionRef,
	}

	res, err := client.Void(req)

	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)
}

func processCapture(client *blockchyp.Client, args blockchyp.CommandLineArguments) {
	validateRequired(args.TransactionID, "tx")
	req := blockchyp.CaptureRequest{
		Amount:             args.Amount,
		TaxAmount:          args.TaxAmount,
		Test:               args.Test,
		Timeout:            args.Timeout,
		WaitForRemovedCard: args.WaitForRemovedCard,
		Force:              args.Force,
		TipAmount:          args.TipAmount,
		TransactionID:      args.TransactionID,
		TransactionRef:     args.TransactionRef,
	}

	res, err := client.Capture(req)

	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)
}

func processGiftActivate(client *blockchyp.Client, args blockchyp.CommandLineArguments) {
	validateRequired(args.Amount, "amount")
	validateRequired(args.TerminalName, "terminal")
	req := blockchyp.GiftActivateRequest{
		Amount:             args.Amount,
		TerminalName:       args.TerminalName,
		Timeout:            args.Timeout,
		WaitForRemovedCard: args.WaitForRemovedCard,
		Force:              args.Force,
		TransactionRef:     args.TransactionRef,
		Test:               args.Test,
	}

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
	req := blockchyp.EnrollRequest{
		Address:            args.Address,
		ExpMonth:           args.ExpiryMonth,
		ExpYear:            args.ExpiryYear,
		ManualEntry:        args.ManualEntry,
		PAN:                args.PAN,
		PostalCode:         args.PostalCode,
		TerminalName:       args.TerminalName,
		Test:               args.Test,
		Timeout:            args.Timeout,
		WaitForRemovedCard: args.WaitForRemovedCard,
		Force:              args.Force,
		TransactionRef:     args.TransactionRef,
	}
	if hasCustomerFields(args) {
		req.Customer = populateCustomer(args)
	}

	if args.Debit {
		req.CardType = blockchyp.CardTypeDebit
	} else if args.EBT {
		req.CardType = blockchyp.CardTypeEBT
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

	req := blockchyp.AuthorizationRequest{
		Address:            args.Address,
		Amount:             args.Amount,
		Async:              args.Async,
		CashBackEnabled:    args.CashBackEnabled,
		CashDiscount:       args.CashDiscount,
		Description:        args.Description,
		DisableSignature:   args.DisableSignature,
		Enroll:             args.Enroll,
		ExpMonth:           args.ExpiryMonth,
		ExpYear:            args.ExpiryYear,
		ManualEntry:        args.ManualEntry,
		OrderRef:           args.OrderRef,
		PAN:                args.PAN,
		PostalCode:         args.PostalCode,
		PromptForTip:       args.PromptForTip,
		Queue:              args.Queue,
		SigFile:            args.SigFile,
		SigFormat:          blockchyp.SignatureFormat(args.SigFormat),
		SigWidth:           args.SigWidth,
		Surcharge:          args.Surcharge,
		TaxAmount:          args.TaxAmount,
		TerminalName:       args.TerminalName,
		Test:               args.Test,
		Timeout:            args.Timeout,
		WaitForRemovedCard: args.WaitForRemovedCard,
		Force:              args.Force,
		TipAmount:          args.TipAmount,
		Token:              args.Token,
		TransactionRef:     args.TransactionRef,
	}

	if args.Cryptocurrency != "" {
		req.Cryptocurrency = &args.Cryptocurrency
		if args.CryptoNetwork != "" {
			req.CryptoNetwork = &args.CryptoNetwork
		}
		if args.CryptoReceiveAddress != "" {
			req.CryptoReceiveAddress = &args.CryptoReceiveAddress
		}
		if args.Label != "" {
			req.PaymentRequestLabel = &args.Label
		}
		if args.Message != "" {
			req.PaymentRequestMessage = &args.Message
		}
	}

	if args.Debit {
		req.CardType = blockchyp.CardTypeDebit
	} else if args.EBT {
		req.CardType = blockchyp.CardTypeEBT
	}
	if hasCustomerFields(args) {
		req.Customer = populateCustomer(args)
	}

	cmd := args.Command
	if cmd == "" {
		cmd = args.Type
	}

	res := &blockchyp.AuthorizationResponse{}
	var err error
	switch cmd {
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

func processLocate(client *blockchyp.Client, args blockchyp.CommandLineArguments) {
	validateRequired(args.TerminalName, "terminal")
	req := blockchyp.LocateRequest{
		TerminalName: args.TerminalName,
		Timeout:      args.Timeout,
	}
	res, err := client.Locate(req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)
}

func processDeleteSurveyQuestion(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := blockchyp.SurveyQuestionRequest{
		Timeout:    args.Timeout,
		QuestionID: args.QuestionID,
	}
	res, err := client.DeleteSurveyQuestion(req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)

}

func processSurveyQuestions(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := blockchyp.SurveyQuestionRequest{
		Timeout: args.Timeout,
	}
	res, err := client.SurveyQuestions(req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)

}

func processSurveyQuestion(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := blockchyp.SurveyQuestionRequest{
		Timeout:    args.Timeout,
		QuestionID: args.QuestionID,
	}
	res, err := client.SurveyQuestion(req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)

}

func processUpdateSurveyQuestion(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := blockchyp.SurveyQuestion{
		Timeout:      args.Timeout,
		ID:           args.QuestionID,
		QuestionText: args.QuestionText,
		QuestionType: args.QuestionType,
		Enabled:      args.Enabled,
		Ordinal:      args.Ordinal,
	}
	res, err := client.UpdateSurveyQuestion(req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)

}

func processSurveyResults(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := blockchyp.SurveyResultsRequest{
		Timeout:    args.Timeout,
		QuestionID: args.QuestionID,
		StartDate:  args.StartDate,
		EndDate:    args.EndDate,
	}
	res, err := client.SurveyResults(req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)

}

func processTCEntry(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := blockchyp.TermsAndConditionsLogRequest{
		Timeout:    args.Timeout,
		LogEntryID: args.LogEntryID,
	}
	res, err := client.TCEntry(req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)

}

func processTCLog(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := blockchyp.TermsAndConditionsLogRequest{
		Timeout:       args.Timeout,
		TransactionID: args.TransactionID,
		StartIndex:    args.StartIndex,
		MaxResults:    args.MaxResults,
	}
	res, err := client.TCLog(req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)

}

func processDeleteTCTemplate(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := blockchyp.TermsAndConditionsTemplateRequest{
		Timeout:    args.Timeout,
		TemplateID: args.TemplateID,
	}
	res, err := client.TCDeleteTemplate(req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)

}

func processTCTemplate(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := blockchyp.TermsAndConditionsTemplateRequest{
		Timeout:    args.Timeout,
		TemplateID: args.TemplateID,
	}
	res, err := client.TCTemplate(req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)

}

func processTCTemplates(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := blockchyp.TermsAndConditionsTemplateRequest{
		Timeout: args.Timeout,
	}
	res, err := client.TCTemplates(req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)

}

func processUpdateTCTemplate(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := blockchyp.TermsAndConditionsTemplate{
		Timeout: args.Timeout,
		Alias:   args.TCAlias,
		Name:    args.TCName,
		Content: args.TCContent,
		ID:      args.TemplateID,
	}
	res, err := client.TCUpdateTemplate(req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)

}

func processActivateTerminal(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := blockchyp.TerminalActivationRequest{
		Timeout:        args.Timeout,
		TerminalName:   args.TerminalName,
		ActivationCode: args.Code,
	}
	res, err := client.ActivateTerminal(req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)

}

func processDeactivateTerminal(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := blockchyp.TerminalDeactivationRequest{
		Timeout:    args.Timeout,
		TerminalID: args.TerminalID,
	}
	res, err := client.DeactivateTerminal(req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)

}

func processTerminals(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := blockchyp.TerminalProfileRequest{
		Timeout: args.Timeout,
	}
	res, err := client.Terminals(req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)

}

func processMerchantUsers(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := blockchyp.MerchantProfileRequest{
		MerchantID: args.MerchantID,
	}
	res, err := client.MerchantUsers(req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)

}

func processGetMerchants(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := blockchyp.GetMerchantsRequest{
		Test:       args.Test,
		StartIndex: args.StartIndex,
		MaxResults: args.MaxResults,
	}
	res, err := client.GetMerchants(req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)
}

func processInviteMerchantUser(client *blockchyp.Client, args blockchyp.CommandLineArguments) {
	validateRequired(args.EMailAddress, "email")
	validateRequired(args.FirstName, "firstName")
	validateRequired(args.LastName, "lastName")
	req := blockchyp.InviteMerchantUserRequest{
		MerchantID: args.MerchantID,
		Email:      args.EMailAddress,
		FirstName:  args.FirstName,
		LastName:   args.LastName,
	}
	res, err := client.InviteMerchantUser(req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)
}

func processDeleteTestMerchant(client *blockchyp.Client, args blockchyp.CommandLineArguments) {
	validateRequired(args.MerchantID, "merchantId")
	req := blockchyp.MerchantProfileRequest{
		MerchantID: args.MerchantID,
	}
	res, err := client.DeleteTestMerchant(req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)
}

func processAddTestMerchant(client *blockchyp.Client, args blockchyp.CommandLineArguments) {
	validateRequired(args.CompanyName, "companyName")
	req := blockchyp.AddTestMerchantRequest{
		DbaName:     args.DBAName,
		CompanyName: args.CompanyName,
	}
	res, err := client.AddTestMerchant(req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)
}

func processUploadMedia(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	validateRequired(args.File, "file")

	info, err := os.Stat(args.File)
	if err != nil {
		handleError(&args, err)
		return
	}

	_, fileName := filepath.Split(info.Name())

	req := blockchyp.UploadMetadata{
		Timeout:  args.Timeout,
		UploadID: args.UploadID,
		FileSize: info.Size(),
		FileName: fileName,
	}

	file, err := os.Open(args.File)
	if err != nil {
		handleError(&args, err)
		return
	}

	defer file.Close()

	res, err := client.UploadMedia(req, file)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)
}

func processPing(client *blockchyp.Client, args blockchyp.CommandLineArguments) {
	validateRequired(args.TerminalName, "terminal")
	req := blockchyp.PingRequest{
		TerminalName: args.TerminalName,
		Timeout:      args.Timeout,
	}
	res, err := client.Ping(req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)
}

func processQueueList(client *blockchyp.Client, args blockchyp.CommandLineArguments) {
	validateRequired(args.TerminalName, "terminal")
	req := blockchyp.ListQueuedTransactionsRequest{
		TerminalName: args.TerminalName,
		Timeout:      args.Timeout,
	}
	res, err := client.ListQueuedTransactions(req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)
}

func processQueueDelete(client *blockchyp.Client, args blockchyp.CommandLineArguments) {
	validateRequired(args.TerminalName, "terminal")
	validateRequired(args.TransactionRef, "txRef")
	req := blockchyp.DeleteQueuedTransactionRequest{
		TerminalName:   args.TerminalName,
		Timeout:        args.Timeout,
		TransactionRef: args.TransactionRef,
	}
	res, err := client.DeleteQueuedTransaction(req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)
}

func processCustomerDelete(client *blockchyp.Client, args blockchyp.CommandLineArguments) {
	validateRequired(args.CustomerID, "customerID")
	req := blockchyp.DeleteCustomerRequest{
		Timeout:    args.Timeout,
		CustomerID: args.CustomerID,
	}
	res, err := client.DeleteCustomer(req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)
}

func processTokenDelete(client *blockchyp.Client, args blockchyp.CommandLineArguments) {
	validateRequired(args.Token, "token")
	req := blockchyp.DeleteTokenRequest{
		Timeout: args.Timeout,
		Token:   args.Token,
	}
	res, err := client.DeleteToken(req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)
}

func parseTimestamp(ts string) (time.Time, error) {

	parsedResult, err := parseTimestampWithFormat(ts, time.RFC3339)
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

	ack := blockchyp.Acknowledgement{
		Error: err.Error(),
	}
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
