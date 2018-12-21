package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"runtime"

	blockchyp "github.com/blockchyp/blockchyp-go"
)

type configSettings struct {
	APIKey          string `json:"apiKey"`
	BearerToken     string `json:"bearerToken"`
	SigningKey      string `json:"signingKey"`
	GatewayHost     string `json:"gateway"`
	TestGatewayHost string `json:"testGateway"`
	Secure          bool   `json:"https"`
	RouteCacheTTL   int    `json:"routeCacheTTL"`
	GatewayTimeout  int    `json:"gatewayTimeout"`
	TerminalTimeout int    `json:"terminalTimeout"`
}

type commandLineArguments struct {
	Type            string `arg:"type"`
	ConfigFile      string `arg:"f"`
	GatewayHost     string `arg:"gateway"`
	TestGatewayHost string `arg:"testGateway"`
	Test            bool   `arg:"test"`
	APIKey          string `arg:"apiKey"`
	BearerToken     string `arg:"bearerToken"`
	SigningKey      string `arg:"signingKey"`
	TransactionRef  string `arg:"txRef"`
	Description     string `arg:"desc"`
	TerminalName    string `arg:"terminal"`
	Token           string `arg:"token"`
	Amount          string `arg:"amount"`
	PromptForTip    bool   `arg:"promptForTip"`
	TipAmount       string `arg:"tip"`
	TaxAmount       string `arg:"tax"`
	CurrencyCode    string `arg:"currency"`
	TransactionID   string `arg:"txId"`
	HTTPS           bool   `arg:"secure"`
}

var currentConfig *configSettings

func main() {

	commandLineArgs := parseArgs()

	processCommand(commandLineArgs)

}

func parseArgs() commandLineArguments {

	args := commandLineArguments{}

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
	flag.BoolVar(&args.HTTPS, "secure", true, "enables or disables https with terminal")

	flag.Parse()

	if args.Type == "" {
		fmt.Println("-type is required")
		handleFatal()
	}

	return args

}

func resolveCredentials(args commandLineArguments) (*blockchyp.APICredentials, error) {

	creds := &blockchyp.APICredentials{}

	if args.APIKey != "" {
		creds.APIKey = args.APIKey
		creds.BearerToken = args.BearerToken
		creds.SigningKey = args.SigningKey
	} else {
		settings, err := loadConfigSettings(args)
		if err != nil {
			return nil, err
		}
		if settings != nil {
			creds.APIKey = settings.APIKey
			creds.BearerToken = settings.BearerToken
			creds.SigningKey = settings.SigningKey
		}
	}

	if creds.APIKey == "" {
		fmt.Println("-apiKey or blockchyp.json file required")
		handleFatal()
	}
	if creds.BearerToken == "" {
		fmt.Println("-bearerToken or blockchyp.json file required")
		handleFatal()
	}
	if creds.SigningKey == "" {
		fmt.Println("-signingKey or blockchyp.json file required")
		handleFatal()
	}

	return creds, nil

}

func loadConfigSettings(args commandLineArguments) (*configSettings, error) {

	if currentConfig != nil {
		return currentConfig, nil
	}

	fileName := args.ConfigFile
	if fileName == "" {
		if runtime.GOOS == "windows" {
			configHome := os.Getenv("userprofile")
			fileName = configHome + "\\blockchyp.json"
		} else {
			configHome := os.Getenv("XDG_CONFIG_HOME")
			if configHome == "" {
				user, err := user.Current()
				if err != nil {
					return nil, err
				}
				configHome = user.HomeDir + "/.config"
			}
			fileName = configHome + "/blockchyp.json"
		}
	}

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		if args.ConfigFile != "" {
			return nil, errors.New(fileName + " not found")
		}
		return nil, nil
	}

	b, err := ioutil.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	currentConfig = &configSettings{}
	err = json.Unmarshal(b, currentConfig)
	return currentConfig, err

}

func resolveClient(args commandLineArguments) (*blockchyp.Client, error) {

	creds, err := resolveCredentials(args)
	if err != nil {
		return nil, err
	}
	client := blockchyp.NewClient(*creds)

	settings, err := loadConfigSettings(args)

	if err != nil {
		return nil, err
	}

	if !args.HTTPS {
		client.HTTPS = args.HTTPS
	} else if settings != nil {
		client.HTTPS = settings.Secure
	}

	if args.GatewayHost != "" {
		client.GatewayHost = args.GatewayHost
	} else if settings.GatewayHost != "" {
		client.GatewayHost = settings.GatewayHost
	} else {
		client.GatewayHost = "https://api.blockchyp.com"
	}

	if args.TestGatewayHost != "" {
		client.TestGatewayHost = args.TestGatewayHost
	} else if settings != nil && settings.TestGatewayHost != "" {
		client.TestGatewayHost = settings.TestGatewayHost
	} else {
		client.TestGatewayHost = "https://test.blockchyp.com"
	}

	return &client, nil
}

func processCommand(args commandLineArguments) {

	client, err := resolveClient(args)

	if err != nil {
		fmt.Println(err)
		handleFatal()
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
	default:
		fmt.Println(args.Type, "is unknown transaction type")
		handleFatal()
	}

}

func processRefund(client *blockchyp.Client, args commandLineArguments) {
	validateRequired(args.TransactionID, "tx")
	req := blockchyp.RefundRequest{}
	req.TransactionRef = args.TransactionRef
	req.TransactionID = args.TransactionID
	req.Test = args.Test

	res, err := client.Refund(req)

	if err != nil {
		if res == nil {
			handleError(err)
		} else if len(res.ResponseDescription) == 0 {
			handleError(err)
		}
	}
	dumpResponse(res)
}

func processReverse(client *blockchyp.Client, args commandLineArguments) {
	validateRequired(args.TransactionRef, "txRef")
	req := blockchyp.AuthorizationRequest{}
	req.TransactionRef = args.TransactionRef
	req.Test = args.Test

	res, err := client.Reverse(req)

	if err != nil {
		if res == nil {
			handleError(err)
		} else if len(res.ResponseDescription) == 0 {
			handleError(err)
		}
	}
	dumpResponse(res)
}

func processCloseBatch(client *blockchyp.Client, args commandLineArguments) {

	req := blockchyp.CloseBatchRequest{}
	req.TransactionRef = args.TransactionRef
	req.Test = args.Test

	res, err := client.CloseBatch(req)

	if err != nil {
		if res == nil {
			handleError(err)
		} else if len(res.ResponseDescription) == 0 {
			handleError(err)
		}
	}
}

func processVoid(client *blockchyp.Client, args commandLineArguments) {
	validateRequired(args.TransactionID, "tx")
	req := blockchyp.VoidRequest{}
	req.TransactionRef = args.TransactionRef
	req.TransactionID = args.TransactionID
	req.Test = args.Test

	res, err := client.Void(req)

	if err != nil {
		if res == nil {
			handleError(err)
		} else if len(res.ResponseDescription) == 0 {
			handleError(err)
		}
	}
	dumpResponse(res)
}

func processCapture(client *blockchyp.Client, args commandLineArguments) {
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
			handleError(err)
		} else if len(res.ResponseDescription) == 0 {
			handleError(err)
		}
	}
	dumpResponse(res)
}

func processGiftActivate(client *blockchyp.Client, args commandLineArguments) {
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
	dumpResponse(res)
}

func processAuth(client *blockchyp.Client, args commandLineArguments) {
	validateRequired(args.Amount, "amount")
	if (args.TerminalName == "") && (args.Token == "") {
		fmt.Println("-terminal or -token requred")
		handleFatal()
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
			handleError(err)
		} else if len(res.ResponseDescription) == 0 {
			handleError(err)
		}
	}
	dumpResponse(res)
}

func processPing(client *blockchyp.Client, args commandLineArguments) {
	validateRequired(args.TerminalName, "terminal")
	req := blockchyp.PingRequest{
		TerminalName: args.TerminalName,
	}
	res, err := client.Ping(req)
	if err != nil {
		if res == nil {
			handleError(err)
		} else if len(res.ResponseDescription) == 0 {
			handleError(err)
		}
	}
	dumpResponse(res)
}

func validateRequired(value string, arg string) {
	if value == "" {
		fmt.Println("-" + arg + " is required")
		handleFatal()
	}
}

func dumpResponse(res interface{}) {

	content, err := json.Marshal(res)
	if err != nil {
		handleFatalError(err)
	}
	fmt.Println(string(content))

}

func handleError(err error) {

	ack := blockchyp.Acknowledgement{}
	ack.Error = err.Error()
	dumpResponse(ack)
	handleFatal()

}

func handleFatalError(err error) {

	fmt.Println(err)
	handleFatal()

}

func handleFatal() {

	os.Exit(1)

}
