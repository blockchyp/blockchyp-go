package main

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	blockchyp "github.com/blockchyp/blockchyp-go"
)

var validSignatureFormats = []string{
	"gif",
	"jpeg",
	"jpg",
	"png",
}

// compileTimeVersion is populated at build time. It contains the version
// string of the current build.
var compileTimeVersion string

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
	flag.StringVar(&args.CurrencyCode, "currency", "USD", "currency code")
	flag.StringVar(&args.TransactionID, "tx", "", "transaction id")
	flag.StringVar(&args.Description, "desc", "", "transaction description")
	flag.BoolVar(&args.Test, "test", false, "sets test mode")
	flag.BoolVar(&args.PromptForTip, "promptForTip", false, "prompt for tip flag")
	flag.BoolVar(&args.ManualEntry, "manual", false, "key in card data manually")
	flag.BoolVar(&args.HTTPS, "secure", true, "enables or disables https with terminal")
	flag.BoolVar(&args.Version, "version", false, "print version and exit")
	flag.StringVar(&args.Message, "message", "", "short message to be displayed on the terminal")
	flag.StringVar(&args.RouteCache, "routeCache", "", "specifies local file location for route cache")
	flag.StringVar(&args.OutputFile, "out", "", "directs output to a file instead of stdout")
	flag.StringVar(&args.SigFormat, "sigFormat", "", "format for signature file (jpeg, png, gif)")
	flag.IntVar(&args.SigWidth, "sigWidth", -1, "optional width in pixels the signature file should be scaled to")
	flag.StringVar(&args.SigFile, "sigFile", "", "optional location to output sig file")

	flag.Parse()

	if args.Version {
		fmt.Println(compileTimeVersion)
		os.Exit(0)
	}

	validateArgs(&args)

	return args

}

func validateArgs(args *blockchyp.CommandLineArguments) {
	if args.Type == "" {
		fatalError("-type is required")
	}

	if args.SigFile != "" {
		if args.SigFormat == "" {
			args.SigFormat = strings.ToLower(strings.TrimPrefix(filepath.Ext(args.SigFile), "."))
		}

		if !validSigFormat(args.SigFormat) {
			fatalErrorf("Invalid signature format: %s", args.SigFormat)
		}
	}
}

func validSigFormat(format string) bool {
	for _, valid := range validSignatureFormats {
		if format == valid {
			return true
		}
	}

	return false
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

	if !args.HTTPS {
		client.HTTPS = args.HTTPS
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
	case "prompt":
		processPrompt(client, args)
	default:
		fatalErrorf("%s is unknown transaction type", args.Type)
	}

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
		if res == nil {
			handleError(&args, err)
		} else if len(res.ResponseDescription) == 0 {
			handleError(&args, err)
		}
	}
	dumpResponse(&args, res)

}

func processPrompt(client *blockchyp.Client, args blockchyp.CommandLineArguments) {
	fmt.Println("not supported yet")
}

func processRefund(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := blockchyp.RefundRequest{}
	req.TransactionRef = args.TransactionRef
	req.TransactionID = args.TransactionID
	req.Amount = args.Amount
	if req.TransactionID == "" {
		req.TerminalName = args.TerminalName
	}
	req.Test = args.Test
	req.ManualEntry = args.ManualEntry
	req.SigWidth = args.SigWidth
	req.SigFormat = args.SigFormat

	res, err := client.Refund(req)

	if err != nil {
		if res == nil {
			handleError(&args, err)
		} else if len(res.ResponseDescription) == 0 {
			handleError(&args, err)
		}
	}
	if args.SigFile != "" && res.SigFile != "" {
		content, err := hex.DecodeString(res.SigFile)
		if err != nil {
			fmt.Println(err)
		} else {
			ioutil.WriteFile(args.SigFile, content, 0644)
			res.SigFile = ""
		}
	}
	dumpResponse(&args, res)
}

func processReverse(client *blockchyp.Client, args blockchyp.CommandLineArguments) {
	validateRequired(args.TransactionRef, "txRef")
	req := blockchyp.AuthorizationRequest{}
	req.TransactionRef = args.TransactionRef
	req.Test = args.Test
	req.ManualEntry = args.ManualEntry

	res, err := client.Reverse(req)

	if err != nil {
		if res == nil {
			handleError(&args, err)
		} else if len(res.ResponseDescription) == 0 {
			handleError(&args, err)
		}
	}
	dumpResponse(&args, res)
}

func processCloseBatch(client *blockchyp.Client, args blockchyp.CommandLineArguments) {

	req := blockchyp.CloseBatchRequest{}
	req.TransactionRef = args.TransactionRef
	req.Test = args.Test

	res, err := client.CloseBatch(req)

	if err != nil {
		if res == nil {
			handleError(&args, err)
		} else if len(res.ResponseDescription) == 0 {
			handleError(&args, err)
		}
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
		if res == nil {
			handleError(&args, err)
		} else if len(res.ResponseDescription) == 0 {
			handleError(&args, err)
		}
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
		if res == nil {
			handleError(&args, err)
		} else if len(res.ResponseDescription) == 0 {
			handleError(&args, err)
		}
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

func processAuth(client *blockchyp.Client, args blockchyp.CommandLineArguments) {
	validateRequired(args.Amount, "amount")
	if (args.TerminalName == "") && (args.Token == "") {
		fatalError("-terminal or -token requred")
	}
	req := blockchyp.AuthorizationRequest{}
	req.TerminalName = args.TerminalName
	req.TransactionRef = args.TransactionRef
	req.Token = args.Token
	req.Description = args.Description
	req.Amount = args.Amount
	req.PromptForTip = args.PromptForTip
	req.TaxAmount = args.TaxAmount
	req.TipAmount = args.TipAmount
	req.Test = args.Test
	req.ManualEntry = args.ManualEntry
	req.SigWidth = args.SigWidth
	req.SigFormat = args.SigFormat

	res := &blockchyp.AuthorizationResponse{}
	var err error
	switch args.Type {
	case "charge":
		res, err = client.Charge(req)
	case "preauth":
		res, err = client.Preauth(req)
	}

	if err != nil {
		if res == nil {
			handleError(&args, err)
		} else if len(res.ResponseDescription) == 0 {
			handleError(&args, err)
		}
	}
	if args.SigFile != "" && res.SigFile != "" {
		content, err := hex.DecodeString(res.SigFile)
		if err != nil {
			fmt.Println(err)
		} else {
			ioutil.WriteFile(args.SigFile, content, 0644)
			res.SigFile = ""
		}
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
		if res == nil {
			handleError(&args, err)
		} else if len(res.ResponseDescription) == 0 {
			handleError(&args, err)
		}
	}
	dumpResponse(&args, res)
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
