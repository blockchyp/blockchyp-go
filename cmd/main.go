package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/blockchyp/blockchyp-core/pkg/logging"
	blockchyp "github.com/blockchyp/blockchyp-go"
)

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
	TerminalName    string `arg:"terminal"`
	Token           string `arg:"token"`
	Amount          string `arg:"amount"`
	TipAmount       string `arg:"tip"`
	TaxAmount       string `arg:"tax"`
	CurrencyCode    string `arg:"currency"`
	TransactionID   string `arg:"txId"`
	HTTPS           bool   `arg:"secure"`
}

func main() {

	commandLineArgs := parseArgs()

	logging.LogJSON(commandLineArgs)

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
	flag.BoolVar(&args.Test, "test", false, "sets test mode")
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
	}

	return creds, nil

}

func resolveClient(args commandLineArguments) (*blockchyp.Client, error) {

	creds, err := resolveCredentials(args)
	if err != nil {
		return nil, err
	}
	client := blockchyp.NewClient(*creds)

	client.HTTPS = args.HTTPS
	if args.GatewayHost != "" {
		client.GatewayHost = args.GatewayHost
	}

	if args.TestGatewayHost != "" {
		client.TestGatewayHost = args.TestGatewayHost
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
	default:
		fmt.Println(args.Type, "is unknown transaction type")
		handleFatal()
	}

}

func processPing(client *blockchyp.Client, args commandLineArguments) {
	validateRequired(args.TerminalName, "terminal")
	req := blockchyp.PingRequest{TerminalName: args.TerminalName}
	res, err := client.Ping(req)
	if err != nil {
		handleFatalError(err)
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
	fmt.Print(string(content))

}

func handleFatalError(err error) {

	fmt.Println(err)
	handleFatal()

}

func handleFatal() {

	//fmt.Println("Type 'help' for a list of commands.")
	os.Exit(1)

}
