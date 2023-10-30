package main

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	blockchyp "github.com/blockchyp/blockchyp-go/v2"
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
	flag.StringVar(&args.MediaID, "mediaId", "", "media id to be used for media related commands")
	flag.StringVar(&args.Name, "name", "", "specifies the name field for a data update")
	flag.IntVar(&args.Delay, "delay", 5, "specifies the delay between slides in seconds")
	flag.StringVar(&args.SlideShowID, "slideShowId", "", "id of a slide show for slide show related operations")
	flag.StringVar(&args.AssetID, "assetId", "", "id of a branding asset")
	flag.StringVar(&args.JSON, "json", "", "raw json request, will override any other command line parameters if used")
	flag.StringVar(&args.JSONFile, "jsonFile", "", "path to a json file to be used for raw json input, will override any other command line parameters if used")
	flag.StringVar(&args.Profile, "profile", "", "profile to source configuration from in blockchyp.json")
	flag.IntVar(&args.QRCodeSize, "qrcodeSize", 256, "default size of the qrcode in pixels if binary for the qr code is requested")
	flag.BoolVar(&args.QRCodeBinary, "qrcodeBinary", false, "if true, a payment link response should also return the image binary")
	flag.IntVar(&args.DaysToExpiration, "daysToExpiration", 0, "days until the payment link should expire")
	flag.BoolVar(&args.ResetConnection, "resetConnection", false, "resets the terminal websocket connection")
	flag.StringVar(&args.RoundingMode, "roundingMode", "", "optional rounding mode for use in surcharge calculation")
	flag.StringVar(&args.Channel, "channel", "stable", "firmware release channel")
	flag.BoolVar(&args.Full, "full", false, "perform full firmware install with transitive dependencies")
	flag.BoolVar(&args.HTTPS, "https", true, "use https for all communication")
	flag.StringVar(&args.Archive, "archive", "", "firmware archive for manual package installation")
	flag.StringVar(&args.Dist, "dist", "", "terminal model distribution")
	flag.StringVar(&args.TestCase, "testCase", "", "test case code for testing and certification")
	flag.BoolVar(&args.Incremental, "incremental", false, "force incremental firmware downloads")
	flag.BoolVar(&args.ChipRejection, "chipRejection", false, "simulates a chip rejection")
	flag.BoolVar(&args.OutOfOrderReversal, "outOfOrderReversal", false, "simulates an out of order auto reversal")
	flag.BoolVar(&args.AsyncReversals, "asyncReversals", false, "causes auto-reversals to run asynchronously")
	flag.BoolVar(&args.CardOnFile, "cardOnFile", false, "flags a transaction as MOTO / card on file.")
	flag.BoolVar(&args.Recurring, "recurring", false, "flags a transaction as recurring.")
	flag.BoolVar(&args.MIT, "mit", false, "manually sets the MIT flag.")
	flag.BoolVar(&args.CIT, "cit", false, "manually sets the CIT flag.")
	flag.StringVar(&args.PONumber, "po", "", "purchase order for L2 transactions")
	flag.StringVar(&args.SupplierReferenceNumber, "srn", "", "supplier reference number for L2 transactions")
	flag.StringVar(&args.PolicyID, "policy", "", "policy id for pricing policy related operations")
	flag.StringVar(&args.StatementID, "statementId", "", "statement id for partner or merchant statement operations")
	flag.StringVar(&args.InvoiceID, "invoiceId", "", "invoice id for partner or merchant statement/invoice operations")
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
	case "merchant-invoices":
		processMerchantInvoices(client, args)
	case "merchant-invoice-detail":
		processMerchantInvoiceDetail(client, args)
	case "partner-statement-detail":
		processPartnerStatementDetail(client, args)
	case "partner-commission-breakdown":
		processPartnerCommissionBreakdown(client, args)
	case "partner-statements":
		processPartnerStatements(client, args)
	case "pricing":
		processPricing(client, args)
	case "sideload":
		processSideLoad(client, args)
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
	case "terminals", "list-terminals":
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
	case "upload-status":
		processUploadStatus(client, args)
	case "media":
		processMedia(client, args)
	case "delete-media":
		processDeleteMedia(client, args)
	case "update-slide-show":
		processUpdateSlideShow(client, args)
	case "slide-shows":
		processSlideShows(client, args)
	case "slide-show":
		processSlideShow(client, args)
	case "delete-slide-show":
		processDeleteSlideShow(client, args)
	case "terminal-branding":
		processTerminalBranding(client, args)
	case "update-branding-asset":
		processUpdateBrandingAsset(client, args)
	case "delete-branding-asset":
		processDeleteBrandingAsset(client, args)
	case "ping":
		processPing(client, args)
	case "reboot":
		processReboot(client, args)
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
	case "payment-link-status":
		processLinkStatus(client, args)
	case "resend-link":
		processResendLink(client, args)
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
	case "update-merchant":
		processUpdateMerchant(client, args)
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
	case "drop-terminal-socket":
		processDropSocket(client, args)
	default:
		fatalErrorf("unknown command: %s", cmd)
	}

}

func processUnlinkToken(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	request := &blockchyp.UnlinkTokenRequest{}

	if !parseJSONInput(args, request) {
		request = &blockchyp.UnlinkTokenRequest{
			Token:      args.Token,
			CustomerID: args.CustomerID,
		}
	}

	ack, err := client.UnlinkToken(*request)
	if err != nil {
		handleError(&args, err)
	}

	dumpResponse(&args, ack)

}

func processPartnerCommissionBreakdown(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	validateRequired(args.StatementID, "statementId")

	request := blockchyp.PartnerCommissionBreakdownRequest{
		StatementID: args.StatementID,
	}

	res, err := client.PartnerCommissionBreakdown(request)

	if nErr, ok := err.(net.Error); ok && nErr.Timeout() {
		res.ResponseDescription = blockchyp.ResponseTimedOut
	} else if err != nil {
		handleError(&args, err)
		return
	}

	dumpResponse(&args, res)

}

func processMerchantInvoiceDetail(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	validateRequired(args.InvoiceID, "invoiceId")

	request := blockchyp.MerchantInvoiceDetailRequest{
		ID: args.InvoiceID,
	}

	res, err := client.MerchantInvoiceDetail(request)

	if nErr, ok := err.(net.Error); ok && nErr.Timeout() {
		res.ResponseDescription = blockchyp.ResponseTimedOut
	} else if err != nil {
		handleError(&args, err)
		return
	}

	dumpResponse(&args, res)

}

func processPartnerStatementDetail(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	validateRequired(args.StatementID, "statementId")

	request := blockchyp.PartnerStatementDetailRequest{
		ID: args.StatementID,
	}

	res, err := client.PartnerStatementDetail(request)

	if nErr, ok := err.(net.Error); ok && nErr.Timeout() {
		res.ResponseDescription = blockchyp.ResponseTimedOut
	} else if err != nil {
		handleError(&args, err)
		return
	}

	dumpResponse(&args, res)

}

func processMerchantInvoices(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	request := blockchyp.MerchantInvoiceListRequest{}

	if args.MerchantID != "" {
		request.MerchantID = &args.MerchantID
		request.InvoiceType = &args.Type
	}

	if args.StartDate != "" {
		ts, err := parseTimestamp(args.StartDate)
		if err != nil {
			handleError(&args, err)
			return
		}
		request.StartDate = &ts
	}
	if args.EndDate != "" {
		ts, err := parseTimestamp(args.EndDate)
		if err != nil {
			handleError(&args, err)
			return
		}
		request.EndDate = &ts
	}

	res, err := client.MerchantInvoices(request)

	if nErr, ok := err.(net.Error); ok && nErr.Timeout() {
		res.ResponseDescription = blockchyp.ResponseTimedOut
	} else if err != nil {
		handleError(&args, err)
		return
	}

	dumpResponse(&args, res)

}

func processPartnerStatements(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	request := blockchyp.PartnerStatementListRequest{}

	if args.StartDate != "" {
		ts, err := parseTimestamp(args.StartDate)
		if err != nil {
			handleError(&args, err)
			return
		}
		request.StartDate = &ts
	}
	if args.EndDate != "" {
		ts, err := parseTimestamp(args.EndDate)
		if err != nil {
			handleError(&args, err)
			return
		}
		request.EndDate = &ts
	}

	res, err := client.PartnerStatements(request)

	if nErr, ok := err.(net.Error); ok && nErr.Timeout() {
		res.ResponseDescription = blockchyp.ResponseTimedOut
	} else if err != nil {
		handleError(&args, err)
		return
	}

	dumpResponse(&args, res)

}

func processPricing(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	validateRequired(args.MerchantID, "merchantId")

	request := blockchyp.PricingPolicyRequest{
		MerchantID: args.MerchantID,
		ID:         args.PolicyID,
	}

	res, err := client.PricingPolicy(request)

	if nErr, ok := err.(net.Error); ok && nErr.Timeout() {
		res.ResponseDescription = blockchyp.ResponseTimedOut
	} else if err != nil {
		handleError(&args, err)
		return
	}

	dumpResponse(&args, res)

}

func processSideLoad(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	validateRequired(args.TerminalName, "terminal")

	request := blockchyp.SideLoadRequest{
		Terminal:        args.TerminalName,
		Channel:         args.Channel,
		Dist:            args.Dist,
		Archive:         args.Archive,
		Full:            args.Full,
		HTTPS:           args.HTTPS,
		Incremental:     args.Incremental,
		TempDir:         os.TempDir(),
		BlockChypClient: client,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Minute,
			Transport: &http.Transport{
				Dial: (&net.Dialer{
					Timeout: 5 * time.Second,
				}).Dial,
				TLSHandshakeTimeout: 5 * time.Second,
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		},
	}

	baseLogger := logrus.New()

	err := blockchyp.SideLoad(request, baseLogger)

	if err != nil {
		handleError(&args, err)
		return
	}

}

func processDropSocket(client *blockchyp.Client, args blockchyp.CommandLineArguments) {
	validateRequired(args.TerminalName, "terminal")

	var req blockchyp.TerminalDeactivationRequest
	if !parseJSONInput(args, &req) {
		req = blockchyp.TerminalDeactivationRequest{
			TerminalName: args.TerminalName,
			TerminalID:   args.TerminalID,
			Test:         args.Test,
			Timeout:      args.Timeout,
		}
	}

	var res blockchyp.Acknowledgement
	err := client.GatewayRequest("/api/drop-terminal-socket", http.MethodPost, req, &res, req.Test, req.Timeout)

	if nErr, ok := err.(net.Error); ok && nErr.Timeout() {
		res.ResponseDescription = blockchyp.ResponseTimedOut
	} else if err != nil {
		handleError(&args, err)
		return
	}

	dumpResponse(&args, res)
}

func processLinkToken(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	request := &blockchyp.LinkTokenRequest{}

	if !parseJSONInput(args, request) {

		request = &blockchyp.LinkTokenRequest{
			Token:      args.Token,
			CustomerID: args.CustomerID,
		}

	}

	ack, err := client.LinkToken(*request)
	if err != nil {
		handleError(&args, err)
	}

	dumpResponse(&args, ack)

}

func processTokenMetadata(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	request := &blockchyp.TokenMetadataRequest{}

	if !parseJSONInput(args, request) {

		request = &blockchyp.TokenMetadataRequest{
			Token: args.Token,
		}

	}

	ack, err := client.TokenMetadata(*request)
	if err != nil {
		handleError(&args, err)
	}

	dumpResponse(&args, ack)

}

func processMerchantProfile(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	request := &blockchyp.MerchantProfileRequest{}

	if !parseJSONInput(args, request) {

		request = &blockchyp.MerchantProfileRequest{
			Test:       args.Test,
			MerchantID: args.MerchantID,
		}

	}

	ack, err := client.MerchantProfile(*request)
	if err != nil {
		handleError(&args, err)
	}

	dumpResponse(&args, ack)

}

func processBatchHistory(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	request := &blockchyp.BatchHistoryRequest{}

	if !parseJSONInput(args, request) {
		request = &blockchyp.BatchHistoryRequest{
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

	}

	ack, err := client.BatchHistory(*request)
	if err != nil {
		handleError(&args, err)
	}

	dumpResponse(&args, ack)

}

func processUpdateMerchant(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	request := &blockchyp.MerchantProfile{}

	if !parseJSONInput(args, request) {

		handleError(&args, errors.New("json input required"))

	}

	ack, err := client.UpdateMerchant(*request)
	if err != nil {
		handleError(&args, err)
	}

	dumpResponse(&args, ack)

}

func processBatchDetails(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	request := &blockchyp.BatchDetailsRequest{}

	if !parseJSONInput(args, request) {

		if args.BatchID == "" {
			fatalErrorf("-batchId is required")
		}

		request = &blockchyp.BatchDetailsRequest{
			BatchID: args.BatchID,
			Test:    args.Test,
		}

	}

	ack, err := client.BatchDetails(*request)
	if err != nil {
		handleError(&args, err)
	}

	dumpResponse(&args, ack)

}

func processTransactionHistory(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	request := &blockchyp.TransactionHistoryRequest{}

	if !parseJSONInput(args, request) {

		request = &blockchyp.TransactionHistoryRequest{
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

	}

	ack, err := client.TransactionHistory(*request)
	if err != nil {
		handleError(&args, err)
	}

	dumpResponse(&args, ack)

}

func processTransactionStatus(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	request := &blockchyp.TransactionStatusRequest{}

	if !parseJSONInput(args, request) {

		if args.TransactionID == "" && args.TransactionRef == "" {
			fatalErrorf("-tx or -txRef are required")
		}

		request = &blockchyp.TransactionStatusRequest{
			TransactionID:  args.TransactionID,
			TransactionRef: args.TransactionRef,
			Test:           args.Test,
		}

	}

	ack, err := client.TransactionStatus(*request)
	if err != nil {
		handleError(&args, err)
	}

	dumpResponse(&args, ack)

}

func processSendLink(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	request := &blockchyp.PaymentLinkRequest{}

	if !parseJSONInput(args, request) {

		if !args.EnrollOnly {
			validateRequired(args.OrderRef, "orderRef")
			validateRequired(args.Amount, "amount")
		}

		if !args.Cashier && !hasCustomerFields(args) {
			fatalErrorf("customer fields (-customerId, -email, etc ) are required")
		}

		request = &blockchyp.PaymentLinkRequest{
			TransactionRef:   args.TransactionRef,
			Description:      args.Description,
			Subject:          args.Subject,
			Amount:           args.Amount,
			OrderRef:         args.OrderRef,
			Test:             args.Test,
			Timeout:          args.Timeout,
			TaxExempt:        args.TaxExempt,
			Transaction:      assembleDisplayTransaction(args),
			Customer:         *populateCustomer(args),
			AutoSend:         args.AutoSend,
			CallbackURL:      args.CallbackURL,
			TCAlias:          args.TCAlias,
			TCName:           args.TCName,
			TCContent:        args.TCContent,
			Cashier:          args.Cashier,
			Enroll:           args.Enroll,
			EnrollOnly:       args.EnrollOnly,
			QrcodeBinary:     args.QRCodeBinary,
			QrcodeSize:       args.QRCodeSize,
			DaysToExpiration: args.DaysToExpiration,
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

	}

	ack, err := client.SendPaymentLink(*request)
	if err != nil {
		handleError(&args, err)
	}

	dumpResponse(&args, ack)
}

func processCancelLink(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	request := &blockchyp.CancelPaymentLinkRequest{}

	if !parseJSONInput(args, request) {

		validateRequired(args.LinkCode, "linkCode")

		request = &blockchyp.CancelPaymentLinkRequest{
			Test:     args.Test,
			Timeout:  args.Timeout,
			LinkCode: args.LinkCode,
		}

	}

	res, err := client.CancelPaymentLink(*request)
	if err != nil {
		handleError(&args, err)
	}

	dumpResponse(&args, res)
}

func processLinkStatus(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	request := &blockchyp.PaymentLinkStatusRequest{}

	if !parseJSONInput(args, request) {

		validateRequired(args.LinkCode, "linkCode")

		request = &blockchyp.PaymentLinkStatusRequest{
			Test:     args.Test,
			Timeout:  args.Timeout,
			LinkCode: args.LinkCode,
		}

	}

	res, err := client.PaymentLinkStatus(*request)
	if err != nil {
		handleError(&args, err)
	}

	dumpResponse(&args, res)
}

func processResendLink(client *blockchyp.Client, args blockchyp.CommandLineArguments) {
	request := &blockchyp.ResendPaymentLinkRequest{}

	if !parseJSONInput(args, request) {

		validateRequired(args.LinkCode, "linkCode")

		request = &blockchyp.ResendPaymentLinkRequest{
			Test:     args.Test,
			Timeout:  args.Timeout,
			LinkCode: args.LinkCode,
		}

	}

	res, err := client.ResendPaymentLink(*request)
	if err != nil {
		handleError(&args, err)
	}

	dumpResponse(&args, res)
}

func getCustomer(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	request := &blockchyp.CustomerRequest{}

	if !parseJSONInput(args, request) {

		if args.CustomerID == "" && args.CustomerRef == "" {
			fatalError("-customerId or -customerRef are required")
		}

		request = &blockchyp.CustomerRequest{
			CustomerID:  args.CustomerID,
			CustomerRef: args.CustomerRef,
			Timeout:     args.Timeout,
		}

	}

	res, err := client.Customer(*request)

	if err != nil {
		handleError(&args, err)
	}

	dumpResponse(&args, res)
}

func searchCustomer(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	request := &blockchyp.CustomerSearchRequest{}

	if !parseJSONInput(args, request) {

		validateRequired(args.Query, "query")

		request = &blockchyp.CustomerSearchRequest{
			Query:   args.Query,
			Timeout: args.Timeout,
		}

	}
	res, err := client.CustomerSearch(*request)
	if err != nil {
		handleError(&args, err)
	}

	dumpResponse(&args, res)
}

func updateCustomer(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	request := &blockchyp.UpdateCustomerRequest{}

	if !parseJSONInput(args, request) {

		request = &blockchyp.UpdateCustomerRequest{
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
		}

	}

	res, err := client.UpdateCustomer(*request)
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

	request := &blockchyp.TerminalStatusRequest{}

	if !parseJSONInput(args, request) {
		validateRequired(args.TerminalName, "terminal")

		request = &blockchyp.TerminalStatusRequest{
			TerminalName:       args.TerminalName,
			Timeout:            args.Timeout,
			WaitForRemovedCard: args.WaitForRemovedCard,
			Force:              args.Force,
			Test:               args.Test,
		}

	}

	response, err := client.TerminalStatus(*request)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, response)
}

func processCaptureSignature(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	request := &blockchyp.CaptureSignatureRequest{}

	if !parseJSONInput(args, request) {
		validateRequired(args.TerminalName, "terminal")
		if args.SigFile == "" && args.SigFormat == blockchyp.SignatureFormatNone {
			fatalErrorf("-%s or -%s are required", "sigFile", "sigFormat")
		}

		request = &blockchyp.CaptureSignatureRequest{
			TerminalName:       args.TerminalName,
			SigFile:            args.SigFile,
			SigFormat:          blockchyp.SignatureFormat(args.SigFormat),
			SigWidth:           args.SigWidth,
			Test:               args.Test,
			Timeout:            args.Timeout,
			WaitForRemovedCard: args.WaitForRemovedCard,
			Force:              args.Force,
		}

	}

	response, err := client.CaptureSignature(*request)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, response)
}

func processBalance(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	request := &blockchyp.BalanceRequest{}

	if !parseJSONInput(args, request) {

		validateRequired(args.TerminalName, "terminal")

		request = &blockchyp.BalanceRequest{
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

	}

	ack, err := client.Balance(*request)
	if err != nil {
		handleError(&args, err)
	}

	dumpResponse(&args, ack)

}

func processCashDiscount(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	request := &blockchyp.CashDiscountRequest{}

	if !parseJSONInput(args, request) {
		validateRequired(args.Amount, "amount")

		request = &blockchyp.CashDiscountRequest{
			Amount:       args.Amount,
			CashDiscount: args.CashDiscount,
			CurrencyCode: args.CurrencyCode,
			Surcharge:    args.Surcharge,
			Test:         args.Test,
			Timeout:      args.Timeout,
		}
	}

	response, err := client.CashDiscount(*request)
	if err != nil {
		handleError(&args, err)
	}

	dumpResponse(&args, response)

}

func processClear(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	request := &blockchyp.ClearTerminalRequest{}

	if !parseJSONInput(args, request) {

		validateRequired(args.TerminalName, "terminal")

		request = &blockchyp.ClearTerminalRequest{
			TerminalName:       args.TerminalName,
			Timeout:            args.Timeout,
			WaitForRemovedCard: args.WaitForRemovedCard,
			Force:              args.Force,
			Test:               args.Test,
		}

	}

	ack, err := client.Clear(*request)
	if err != nil {
		handleError(&args, err)
	}

	dumpResponse(&args, ack)

}

func processTermsAndConditions(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	request := &blockchyp.TermsAndConditionsRequest{}

	if !parseJSONInput(args, request) {

		validateRequired(args.TerminalName, "terminal")

		request = &blockchyp.TermsAndConditionsRequest{
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

	}

	ack, err := client.TermsAndConditions(*request)
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

func parseJSONInput(args blockchyp.CommandLineArguments, req interface{}) bool {

	rawJSON := args.JSON

	if args.JSONFile != "" {
		b, err := ioutil.ReadFile(args.JSONFile)
		if err != nil {
			handleError(&args, err)
			return false
		}
		rawJSON = string(b)
	}

	if len(rawJSON) == 0 {
		return false
	}

	err := json.Unmarshal([]byte(rawJSON), req)
	if err != nil {
		handleError(&args, err)
		return false
	}

	return true

}

func processDisplay(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	request := &blockchyp.TransactionDisplayRequest{}

	if !parseJSONInput(args, request) {

		validateRequired(args.TerminalName, "terminal")

		request = &blockchyp.TransactionDisplayRequest{
			TerminalName:       args.TerminalName,
			Transaction:        assembleDisplayTransaction(args),
			Timeout:            args.Timeout,
			WaitForRemovedCard: args.WaitForRemovedCard,
			Force:              args.Force,
			Test:               args.Test,
		}

	}

	ack, err := client.UpdateTransactionDisplay(*request)
	if err != nil {
		handleError(&args, err)
	}

	dumpResponse(&args, ack)

}

func processMessage(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := &blockchyp.MessageRequest{}

	if !parseJSONInput(args, req) {

		validateRequired(args.Message, "message")
		validateRequired(args.TerminalName, "terminal")

		req = &blockchyp.MessageRequest{
			Message:            args.Message,
			TerminalName:       args.TerminalName,
			Timeout:            args.Timeout,
			WaitForRemovedCard: args.WaitForRemovedCard,
			Force:              args.Force,
			Test:               args.Test,
		}

	}

	res, err := client.Message(*req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)

}

func processBooleanPrompt(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := &blockchyp.BooleanPromptRequest{}

	if !parseJSONInput(args, req) {

		validateRequired(args.Prompt, "prompt")
		validateRequired(args.TerminalName, "terminal")

		req = &blockchyp.BooleanPromptRequest{
			NoCaption:          args.NoCaption,
			Prompt:             args.Prompt,
			TerminalName:       args.TerminalName,
			Test:               args.Test,
			Timeout:            args.Timeout,
			WaitForRemovedCard: args.WaitForRemovedCard,
			Force:              args.Force,
			YesCaption:         args.YesCaption,
		}

	}

	res, err := client.BooleanPrompt(*req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)

}

func processTextPrompt(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := &blockchyp.TextPromptRequest{}

	if !parseJSONInput(args, req) {

		validateRequired(args.PromptType, "promptType")
		validateRequired(args.TerminalName, "terminal")

		req = &blockchyp.TextPromptRequest{
			PromptType:         blockchyp.PromptType(args.PromptType),
			TerminalName:       args.TerminalName,
			Timeout:            args.Timeout,
			WaitForRemovedCard: args.WaitForRemovedCard,
			Force:              args.Force,
			Test:               args.Test,
		}

	}

	res, err := client.TextPrompt(*req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)

}

func processRefund(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := &blockchyp.RefundRequest{}

	if !parseJSONInput(args, req) {
		req = &blockchyp.RefundRequest{
			Address:                    args.Address,
			Amount:                     args.Amount,
			DisableSignature:           args.DisableSignature,
			ManualEntry:                args.ManualEntry,
			PostalCode:                 args.PostalCode,
			SigFile:                    args.SigFile,
			SigFormat:                  blockchyp.SignatureFormat(args.SigFormat),
			SigWidth:                   args.SigWidth,
			TerminalName:               args.TerminalName,
			Test:                       args.Test,
			Timeout:                    args.Timeout,
			WaitForRemovedCard:         args.WaitForRemovedCard,
			Force:                      args.Force,
			Token:                      args.Token,
			TransactionID:              args.TransactionID,
			TransactionRef:             args.TransactionRef,
			ResetConnection:            args.ResetConnection,
			SimulateChipRejection:      args.ChipRejection,
			SimulateOutOfOrderReversal: args.OutOfOrderReversal,
			AsyncReversals:             args.AsyncReversals,
			TestCase:                   args.TestCase,
			Mit:                        args.MIT,
			Cit:                        args.CIT,
			PAN:                        args.PAN,
			ExpMonth:                   args.ExpiryMonth,
			ExpYear:                    args.ExpiryYear,
		}

		if args.Debit {
			req.CardType = blockchyp.CardTypeDebit
		} else if args.EBT {
			req.CardType = blockchyp.CardTypeEBT
			// EBT free range refunds are not permitted.
			req.TerminalName = args.TerminalName
		}
	}

	res, err := client.Refund(*req)

	if err != nil {
		handleError(&args, err)
	}

	dumpResponse(&args, res)
}

func processReverse(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := &blockchyp.AuthorizationRequest{}

	if !parseJSONInput(args, req) {
		validateRequired(args.TransactionRef, "txRef")
		req = &blockchyp.AuthorizationRequest{
			ManualEntry:        args.ManualEntry,
			Test:               args.Test,
			Timeout:            args.Timeout,
			WaitForRemovedCard: args.WaitForRemovedCard,
			Force:              args.Force,
			TransactionRef:     args.TransactionRef,
			TestCase:           args.TestCase,
		}

		if args.Debit {
			req.CardType = blockchyp.CardTypeDebit
		} else if args.EBT {
			req.CardType = blockchyp.CardTypeEBT
		}
	}

	res, err := client.Reverse(*req)

	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)
}

func processCloseBatch(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := &blockchyp.CloseBatchRequest{}

	if !parseJSONInput(args, req) {
		req = &blockchyp.CloseBatchRequest{
			Test:               args.Test,
			Timeout:            args.Timeout,
			WaitForRemovedCard: args.WaitForRemovedCard,
			Force:              args.Force,
			TransactionRef:     args.TransactionRef,
		}
	}

	res, err := client.CloseBatch(*req)

	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)
}

func processVoid(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := &blockchyp.VoidRequest{}

	if !parseJSONInput(args, req) {
		validateRequired(args.TransactionID, "tx")
		req = &blockchyp.VoidRequest{
			Test:               args.Test,
			Timeout:            args.Timeout,
			WaitForRemovedCard: args.WaitForRemovedCard,
			Force:              args.Force,
			TransactionID:      args.TransactionID,
			TransactionRef:     args.TransactionRef,
			TestCase:           args.TestCase,
		}
	}

	res, err := client.Void(*req)

	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)

}

func processCapture(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := &blockchyp.CaptureRequest{}

	if !parseJSONInput(args, req) {
		validateRequired(args.TransactionID, "tx")
		req = &blockchyp.CaptureRequest{
			Amount:             args.Amount,
			TaxAmount:          args.TaxAmount,
			Test:               args.Test,
			Timeout:            args.Timeout,
			WaitForRemovedCard: args.WaitForRemovedCard,
			Force:              args.Force,
			TipAmount:          args.TipAmount,
			TransactionID:      args.TransactionID,
			TransactionRef:     args.TransactionRef,
			TestCase:           args.TestCase,
		}
	}

	res, err := client.Capture(*req)

	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)
}

func processGiftActivate(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := &blockchyp.GiftActivateRequest{}

	if !parseJSONInput(args, req) {
		validateRequired(args.Amount, "amount")
		validateRequired(args.TerminalName, "terminal")
		req = &blockchyp.GiftActivateRequest{
			Amount:             args.Amount,
			TerminalName:       args.TerminalName,
			Timeout:            args.Timeout,
			WaitForRemovedCard: args.WaitForRemovedCard,
			Force:              args.Force,
			TransactionRef:     args.TransactionRef,
			Test:               args.Test,
		}
	}

	res, err := client.GiftActivate(*req)

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

	req := &blockchyp.EnrollRequest{}

	if !parseJSONInput(args, req) {

		if (args.TerminalName == "") && (args.Token == "") && (args.PAN == "") {
			fatalError("-terminal or -token requred")
		}
		req = &blockchyp.EnrollRequest{
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
			ResetConnection:    args.ResetConnection,
		}
		if hasCustomerFields(args) {
			req.Customer = populateCustomer(args)
		}

		if args.Debit {
			req.CardType = blockchyp.CardTypeDebit
		} else if args.EBT {
			req.CardType = blockchyp.CardTypeEBT
		}
	}

	res, err := client.Enroll(*req)

	if err != nil {
		handleError(&args, err)
	}

	dumpResponse(&args, res)
}

func processAuth(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := &blockchyp.AuthorizationRequest{}

	if !parseJSONInput(args, req) {

		validateRequired(args.Amount, "amount")
		if (args.TerminalName == "") && (args.Token == "") && (args.PAN == "") && (args.TransactionID == "") {
			fatalError("-terminal, -token, or -txId required")
		}

		req = &blockchyp.AuthorizationRequest{
			Address:                    args.Address,
			Amount:                     args.Amount,
			Async:                      args.Async,
			CashBackEnabled:            args.CashBackEnabled,
			CashDiscount:               args.CashDiscount,
			Description:                args.Description,
			DisableSignature:           args.DisableSignature,
			Enroll:                     args.Enroll,
			ExpMonth:                   args.ExpiryMonth,
			ExpYear:                    args.ExpiryYear,
			ManualEntry:                args.ManualEntry,
			OrderRef:                   args.OrderRef,
			PAN:                        args.PAN,
			PostalCode:                 args.PostalCode,
			PromptForTip:               args.PromptForTip,
			Queue:                      args.Queue,
			SigFile:                    args.SigFile,
			SigFormat:                  blockchyp.SignatureFormat(args.SigFormat),
			SigWidth:                   args.SigWidth,
			Surcharge:                  args.Surcharge,
			TaxAmount:                  args.TaxAmount,
			TerminalName:               args.TerminalName,
			Test:                       args.Test,
			Timeout:                    args.Timeout,
			WaitForRemovedCard:         args.WaitForRemovedCard,
			Force:                      args.Force,
			TipAmount:                  args.TipAmount,
			Token:                      args.Token,
			TransactionRef:             args.TransactionRef,
			ResetConnection:            args.ResetConnection,
			SimulateChipRejection:      args.ChipRejection,
			SimulateOutOfOrderReversal: args.OutOfOrderReversal,
			AsyncReversals:             args.AsyncReversals,
			CardOnFile:                 args.CardOnFile,
			Recurring:                  args.Recurring,
			TestCase:                   args.TestCase,
			Mit:                        args.MIT,
			Cit:                        args.CIT,
			TransactionID:              args.TransactionID,
			PurchaseOrderNumber:        args.PONumber,
			SupplierReferenceNumber:    args.SupplierReferenceNumber,
		}

		displayTx := assembleDisplayTransaction(args)

		if displayTx != nil && displayTx.Items != nil {
			req.LineItems = displayTx.Items
		}

		if args.TransactionID != "" {
			req.TransactionID = args.TransactionID
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

		if args.RoundingMode != "" {
			mode := blockchyp.RoundingMode(args.RoundingMode)
			req.RoundingMode = &mode
		}

	}

	cmd := args.Command
	if cmd == "" {
		cmd = args.Type
	}

	res := &blockchyp.AuthorizationResponse{}
	var err error
	switch cmd {
	case "charge":
		res, err = client.Charge(*req)
	case "preauth":
		res, err = client.Preauth(*req)
	}

	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)
}

func processLocate(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := &blockchyp.LocateRequest{}

	if !parseJSONInput(args, req) {

		validateRequired(args.TerminalName, "terminal")
		req = &blockchyp.LocateRequest{
			TerminalName: args.TerminalName,
			Timeout:      args.Timeout,
		}

	}
	res, err := client.Locate(*req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)
}

func processDeleteSurveyQuestion(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := &blockchyp.SurveyQuestionRequest{}

	if !parseJSONInput(args, req) {
		req = &blockchyp.SurveyQuestionRequest{
			Timeout:    args.Timeout,
			QuestionID: args.QuestionID,
		}
	}
	res, err := client.DeleteSurveyQuestion(*req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)

}

func processSurveyQuestions(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := &blockchyp.SurveyQuestionRequest{}

	if !parseJSONInput(args, req) {
		req = &blockchyp.SurveyQuestionRequest{
			Timeout: args.Timeout,
		}
	}
	res, err := client.SurveyQuestions(*req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)

}

func processSurveyQuestion(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := &blockchyp.SurveyQuestionRequest{}

	if !parseJSONInput(args, req) {
		req = &blockchyp.SurveyQuestionRequest{
			Timeout:    args.Timeout,
			QuestionID: args.QuestionID,
		}
	}
	res, err := client.SurveyQuestion(*req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)

}

func processUpdateSurveyQuestion(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := &blockchyp.SurveyQuestion{}

	if !parseJSONInput(args, req) {
		req = &blockchyp.SurveyQuestion{
			Timeout:      args.Timeout,
			ID:           args.QuestionID,
			QuestionText: args.QuestionText,
			QuestionType: args.QuestionType,
			Enabled:      args.Enabled,
			Ordinal:      args.Ordinal,
		}
	}
	res, err := client.UpdateSurveyQuestion(*req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)

}

func processSurveyResults(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := &blockchyp.SurveyResultsRequest{}

	if !parseJSONInput(args, req) {
		req = &blockchyp.SurveyResultsRequest{
			Timeout:    args.Timeout,
			QuestionID: args.QuestionID,
			StartDate:  args.StartDate,
			EndDate:    args.EndDate,
		}
	}
	res, err := client.SurveyResults(*req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)

}

func processTCEntry(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := &blockchyp.TermsAndConditionsLogRequest{}

	if !parseJSONInput(args, req) {
		req = &blockchyp.TermsAndConditionsLogRequest{
			Timeout:    args.Timeout,
			LogEntryID: args.LogEntryID,
		}
	}
	res, err := client.TCEntry(*req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)

}

func processTCLog(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := &blockchyp.TermsAndConditionsLogRequest{}

	if !parseJSONInput(args, req) {
		req = &blockchyp.TermsAndConditionsLogRequest{
			Timeout:       args.Timeout,
			TransactionID: args.TransactionID,
			StartIndex:    args.StartIndex,
			MaxResults:    args.MaxResults,
		}
	}
	res, err := client.TCLog(*req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)

}

func processDeleteTCTemplate(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := &blockchyp.TermsAndConditionsTemplateRequest{}

	if !parseJSONInput(args, req) {
		req = &blockchyp.TermsAndConditionsTemplateRequest{
			Timeout:    args.Timeout,
			TemplateID: args.TemplateID,
		}
	}
	res, err := client.TCDeleteTemplate(*req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)

}

func processTCTemplate(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := &blockchyp.TermsAndConditionsTemplateRequest{}

	if !parseJSONInput(args, req) {
		req = &blockchyp.TermsAndConditionsTemplateRequest{
			Timeout:    args.Timeout,
			TemplateID: args.TemplateID,
		}
	}
	res, err := client.TCTemplate(*req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)

}

func processTCTemplates(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := &blockchyp.TermsAndConditionsTemplateRequest{}

	if !parseJSONInput(args, req) {
		req = &blockchyp.TermsAndConditionsTemplateRequest{
			Timeout: args.Timeout,
		}
	}
	res, err := client.TCTemplates(*req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)

}

func processUpdateTCTemplate(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := &blockchyp.TermsAndConditionsTemplate{}

	if !parseJSONInput(args, req) {
		req = &blockchyp.TermsAndConditionsTemplate{
			Timeout: args.Timeout,
			Alias:   args.TCAlias,
			Name:    args.TCName,
			Content: args.TCContent,
			ID:      args.TemplateID,
		}
	}
	res, err := client.TCUpdateTemplate(*req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)

}

func processActivateTerminal(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := &blockchyp.TerminalActivationRequest{}

	if !parseJSONInput(args, req) {
		req = &blockchyp.TerminalActivationRequest{
			Timeout:        args.Timeout,
			TerminalName:   args.TerminalName,
			ActivationCode: args.Code,
		}
	}
	res, err := client.ActivateTerminal(*req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)

}

func processDeactivateTerminal(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := &blockchyp.TerminalDeactivationRequest{}

	if !parseJSONInput(args, req) {
		req = &blockchyp.TerminalDeactivationRequest{
			Timeout:    args.Timeout,
			TerminalID: args.TerminalID,
		}
	}
	res, err := client.DeactivateTerminal(*req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)

}

func processTerminals(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := &blockchyp.TerminalProfileRequest{}

	if !parseJSONInput(args, req) {
		req = &blockchyp.TerminalProfileRequest{
			Timeout: args.Timeout,
		}
	}
	res, err := client.Terminals(*req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)

}

func processMerchantUsers(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := &blockchyp.MerchantProfileRequest{}

	if !parseJSONInput(args, req) {
		req = &blockchyp.MerchantProfileRequest{
			MerchantID: args.MerchantID,
		}
	}
	res, err := client.MerchantUsers(*req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)

}

func processGetMerchants(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := &blockchyp.GetMerchantsRequest{}

	if !parseJSONInput(args, req) {
		req = &blockchyp.GetMerchantsRequest{
			Test:       args.Test,
			StartIndex: args.StartIndex,
			MaxResults: args.MaxResults,
		}
	}
	res, err := client.GetMerchants(*req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)
}

func processInviteMerchantUser(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := &blockchyp.InviteMerchantUserRequest{}

	if !parseJSONInput(args, req) {
		validateRequired(args.EMailAddress, "email")
		validateRequired(args.FirstName, "firstName")
		validateRequired(args.LastName, "lastName")
		req = &blockchyp.InviteMerchantUserRequest{
			MerchantID: args.MerchantID,
			Email:      args.EMailAddress,
			FirstName:  args.FirstName,
			LastName:   args.LastName,
		}
	}
	res, err := client.InviteMerchantUser(*req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)
}

func processDeleteTestMerchant(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := &blockchyp.MerchantProfileRequest{}

	if !parseJSONInput(args, req) {
		validateRequired(args.MerchantID, "merchantId")
		req = &blockchyp.MerchantProfileRequest{
			MerchantID: args.MerchantID,
		}
	}
	res, err := client.DeleteTestMerchant(*req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)
}

func processAddTestMerchant(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := &blockchyp.AddTestMerchantRequest{}

	if !parseJSONInput(args, req) {
		validateRequired(args.CompanyName, "companyName")
		req = &blockchyp.AddTestMerchantRequest{
			DBAName:     args.DBAName,
			CompanyName: args.CompanyName,
		}
	}
	res, err := client.AddTestMerchant(*req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)
}

func processSlideShows(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := &blockchyp.SlideShowRequest{}

	if !parseJSONInput(args, req) {
		req = &blockchyp.SlideShowRequest{
			Timeout: args.Timeout,
		}
	}

	res, err := client.SlideShows(*req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)
}

func processSlideShow(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := &blockchyp.SlideShowRequest{
		Timeout:     args.Timeout,
		SlideShowID: args.SlideShowID,
	}

	if !parseJSONInput(args, req) {
		req = &blockchyp.SlideShowRequest{
			Timeout:     args.Timeout,
			SlideShowID: args.SlideShowID,
		}
	}

	res, err := client.SlideShow(*req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)
}

func processTerminalBranding(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := &blockchyp.BrandingAssetRequest{}

	if !parseJSONInput(args, req) {
		req = &blockchyp.BrandingAssetRequest{
			Timeout: args.Timeout,
		}
	}
	res, err := client.TerminalBranding(*req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)
}

func processDeleteBrandingAsset(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := &blockchyp.BrandingAssetRequest{}

	if !parseJSONInput(args, req) {
		req = &blockchyp.BrandingAssetRequest{
			AssetID: args.AssetID,
			Timeout: args.Timeout,
		}
	}
	res, err := client.DeleteBrandingAsset(*req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)
}

func processUpdateBrandingAsset(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := &blockchyp.BrandingAsset{}

	if !parseJSONInput(args, req) {
		req = &blockchyp.BrandingAsset{
			ID:          args.AssetID,
			Timeout:     args.Timeout,
			MediaID:     args.MediaID,
			SlideShowID: args.SlideShowID,
			Enabled:     args.Enabled,
			Ordinal:     args.Ordinal,
			Notes:       args.Message,
		}
	}
	res, err := client.UpdateBrandingAsset(*req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)
}

func processDeleteSlideShow(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := &blockchyp.SlideShowRequest{}

	if !parseJSONInput(args, req) {
		req = &blockchyp.SlideShowRequest{
			Timeout:     args.Timeout,
			SlideShowID: args.SlideShowID,
		}
	}
	res, err := client.DeleteSlideShow(*req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)
}

func processUpdateSlideShow(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := &blockchyp.SlideShow{}

	if !parseJSONInput(args, req) {
		slideIds := strings.Split(args.MediaID, ",")

		req = &blockchyp.SlideShow{
			ID:      args.SlideShowID,
			Timeout: args.Timeout,
			Delay:   args.Delay,
			Name:    args.Name,
		}

		slides := make([]*blockchyp.Slide, len(slideIds))

		for idx, id := range slideIds {
			slides[idx] = &blockchyp.Slide{
				MediaID: id,
				Ordinal: idx,
			}
		}

		req.Slides = slides
	}

	res, err := client.UpdateSlideShow(*req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)
}

func processDeleteMedia(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := &blockchyp.MediaRequest{}

	if !parseJSONInput(args, req) {
		req = &blockchyp.MediaRequest{
			Timeout: args.Timeout,
			MediaID: args.MediaID,
		}
	}

	res, err := client.DeleteMediaAsset(*req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)
}

func processMedia(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := &blockchyp.MediaRequest{}

	if !parseJSONInput(args, req) {
		req = &blockchyp.MediaRequest{
			Timeout: args.Timeout,
		}

		if args.MediaID != "" {
			req.MediaID = args.MediaID

		}
	}

	if req.MediaID != "" {
		res, err := client.MediaAsset(*req)
		if err != nil {
			handleError(&args, err)
		}
		dumpResponse(&args, res)
		return
	}

	res, err := client.Media(*req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)
}

func processUploadStatus(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := &blockchyp.UploadStatusRequest{}

	if !parseJSONInput(args, req) {
		validateRequired(args.UploadID, "uploadId")
		req = &blockchyp.UploadStatusRequest{
			Timeout:  args.Timeout,
			UploadID: args.UploadID,
		}
	}
	res, err := client.UploadStatus(*req)
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

	req := &blockchyp.UploadMetadata{}

	if !parseJSONInput(args, req) {
		req = &blockchyp.UploadMetadata{
			Timeout:  args.Timeout,
			UploadID: args.UploadID,
			FileSize: info.Size(),
			FileName: fileName,
		}
	}

	file, err := os.Open(args.File)
	if err != nil {
		handleError(&args, err)
		return
	}

	defer file.Close()

	res, err := client.UploadMedia(*req, file)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)
}

func processPing(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := &blockchyp.PingRequest{}

	if !parseJSONInput(args, req) {
		validateRequired(args.TerminalName, "terminal")
		req = &blockchyp.PingRequest{
			TerminalName: args.TerminalName,
			Timeout:      args.Timeout,
		}
	}
	res, err := client.Ping(*req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)
}

func processReboot(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := &blockchyp.PingRequest{}

	if !parseJSONInput(args, req) {
		validateRequired(args.TerminalName, "terminal")
		req = &blockchyp.PingRequest{
			TerminalName: args.TerminalName,
			Timeout:      args.Timeout,
		}
	}
	res, err := client.Reboot(*req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)
}

func processQueueList(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := &blockchyp.ListQueuedTransactionsRequest{}

	if !parseJSONInput(args, req) {
		validateRequired(args.TerminalName, "terminal")
		req = &blockchyp.ListQueuedTransactionsRequest{
			TerminalName: args.TerminalName,
			Timeout:      args.Timeout,
		}
	}
	res, err := client.ListQueuedTransactions(*req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)
}

func processQueueDelete(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := &blockchyp.DeleteQueuedTransactionRequest{}

	if !parseJSONInput(args, req) {
		validateRequired(args.TerminalName, "terminal")
		validateRequired(args.TransactionRef, "txRef")
		req = &blockchyp.DeleteQueuedTransactionRequest{
			TerminalName:   args.TerminalName,
			Timeout:        args.Timeout,
			TransactionRef: args.TransactionRef,
		}
	}
	res, err := client.DeleteQueuedTransaction(*req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)
}

func processCustomerDelete(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := &blockchyp.DeleteCustomerRequest{}

	if !parseJSONInput(args, req) {
		validateRequired(args.CustomerID, "customerID")
		req = &blockchyp.DeleteCustomerRequest{
			Timeout:    args.Timeout,
			CustomerID: args.CustomerID,
		}
	}
	res, err := client.DeleteCustomer(*req)
	if err != nil {
		handleError(&args, err)
	}
	dumpResponse(&args, res)
}

func processTokenDelete(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := &blockchyp.DeleteTokenRequest{}

	if !parseJSONInput(args, req) {
		validateRequired(args.Token, "token")
		req = &blockchyp.DeleteTokenRequest{
			Timeout: args.Timeout,
			Token:   args.Token,
		}
	}
	res, err := client.DeleteToken(*req)
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
