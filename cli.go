package blockchyp

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
)

// ConfigSettings contains configuration options for the CLI.
type ConfigSettings struct {
	APIKey          string `json:"apiKey"`
	BearerToken     string `json:"bearerToken"`
	SigningKey      string `json:"signingKey"`
	GatewayHost     string `json:"gatewayHost"`
	DashboardHost   string `json:"dashboardHost"`
	TestGatewayHost string `json:"testGatewayHost"`
	Secure          bool   `json:"https"`
	RouteCacheTTL   int    `json:"routeCacheTTL"`
	GatewayTimeout  int    `json:"gatewayTimeout"`
	TerminalTimeout int    `json:"terminalTimeout"`
}

// CommandLineArguments contains arguments which are passed in at runtime.
type CommandLineArguments struct {
	Type                        string `args:"type"` //deprecated - use cmd instead
	Command                     string `args:"cmd"`
	ManualEntry                 bool   `arg:"manual"`
	ConfigFile                  string `arg:"f"`
	GatewayHost                 string `arg:"gateway"`
	DashboardHost               string `arg:"dashboard"`
	TestGatewayHost             string `arg:"testGateway"`
	Dashboard                   string `arg:"dashboard"`
	Test                        bool   `arg:"test"`
	APIKey                      string `arg:"apiKey"`
	BearerToken                 string `arg:"bearerToken"`
	SigningKey                  string `arg:"signingKey"`
	TransactionRef              string `arg:"txRef"`
	Description                 string `arg:"desc"`
	TerminalName                string `arg:"terminal"`
	Token                       string `arg:"token"`
	Amount                      string `arg:"amount"`
	PromptForTip                bool   `arg:"promptForTip"`
	Message                     string `arg:"message"`
	TipAmount                   string `arg:"tip"`
	TaxAmount                   string `arg:"tax"`
	TaxExempt                   bool   `arg:"taxExempt"`
	CurrencyCode                string `arg:"currency"`
	TransactionID               string `arg:"txId"`
	RouteCache                  string `arg:"routeCache"`
	OutputFile                  string `arg:"out"`
	SigFormat                   string `arg:"sigFormat"`
	SigWidth                    int    `arg:"sigWidth"`
	SigFile                     string `arg:"sigFile"`
	HTTPS                       bool   `arg:"secure"`
	Version                     bool   `arg:"version"`
	DisplayTotal                string `arg:"displayTotal"`
	DisplayTax                  string `arg:"displayTax"`
	DisplaySubtotal             string `arg:"displaySubtotal"`
	LineItemID                  string `arg:"lineItemId"`
	LineItemDescription         string `arg:"lineItemDescription"`
	LineItemPrice               string `arg:"lineItemPrice"`
	LineItemQty                 string `arg:"lineItemQty"`
	LineItemExtended            string `arg:"lineItemExtended"`
	LineItemDiscountDescription string `arg:"lineItemDiscountDescription"`
	LineItemDiscountAmount      string `arg:"lineItemDiscountAmount"`
	Prompt                      string `arg:"prompt"`
	PromptType                  string `arg:"promptType"`
	YesCaption                  string `arg:"yesCaption"`
	NoCaption                   string `arg:"noCaption"`
	EBT                         bool   `arg:"ebt"`
	Debit                       bool   `arg:"debit"`
	TCAlias                     string `arg:"tcAlias"`
	TCName                      string `arg:"tcName"`
	TCContent                   string `arg:"tcContent"`
	Timeout                     int    `arg:"timeout"`
	WaitForRemovedCard          bool   `arg:"waitForRemovedCard"`
	Force                       bool   `arg:"force"`
	SigRequired                 bool   `arg:"sigRequired"`
	CashBackEnabled             bool   `arg:"cashback"`
	Enroll                      bool   `arg:"enroll"`
	EnrollOnly                  bool   `arg:"enrollOnly"`
	DisableSignature            bool   `arg:"disableSignature"`
	CustomerID                  string `arg:"customerId"`
	CustomerRef                 string `arg:"customerRef"`
	FirstName                   string `arg:"firstName"`
	LastName                    string `arg:"lastName"`
	CompanyName                 string `arg:"companyName"`
	EMailAddress                string `arg:"email"`
	SMSNumber                   string `arg:"sms"`
	PAN                         string `arg:"pan"`
	ExpiryMonth                 string `arg:"expMonth"`
	ExpiryYear                  string `arg:"expYear"`
	Subject                     string `arg:"subject"`
	AutoSend                    bool   `arg:"autoSend"`
	OrderRef                    string `arg:"orderRef"`
	Query                       string `arg:"query"`
	CallbackURL                 string `arg:"callbackUrl"`
	Surcharge                   bool   `arg:"surcharge"`
	CashDiscount                bool   `arg:"cashDiscount"`
	PostalCode                  string `arg:"postalCode"`
	Address                     string `arg:"address"`
	Cashier                     bool   `arg:"cashier"`
	StartDate                   string `arg:"startDate"`
	EndDate                     string `arg:"endDate"`
	BatchID                     string `arg:"batchId"`
	MaxResults                  int    `arg:"maxResults"`
	StartIndex                  int    `arg:"startIndex"`
	Queue                       bool   `arg:"queue"`
	Async                       bool   `arg:"async"`
	LogRequests                 bool   `arg:"logRequests"`
	LinkCode                    string `arg:"linkCode"`
	Cryptocurrency              string `arg:"crypto"`
	CryptoNetwork               string `arg:"cryptoNetwork"`
	CryptoReceiveAddress        string `arg:"receiveAddress"`
	Label                       string `arg:"label"`
	DBAName                     string `arg:"dbaName"`
	PolicyID                    string `arg:"policyId"`
	MerchantID                  string `arg:"merchantId"`
	TerminalID                  string `arg:"terminalId"`
	Code                        string `arg:"code"`
	TemplateID                  string `arg:"templateId"`
	LogEntryID                  string `arg:"logEntryId"`
	QuestionID                  string `arg:"questionId"`
	IncludeResponseData         bool   `arg:"includeResponseData"`
	QuestionType                string `arg:"questionType"`
	QuestionText                string `arg:"questionText"`
	Enabled                     bool   `arg:"enabled"`
	Ordinal                     int    `arg:"ordinal"`
	File                        string `arg:"file"`
	UploadID                    string `arg:"uploadId"`
	MediaID                     string `arg:"mediaId"`
	Name                        string `arg:"name"`
	Delay                       int    `arg:"delay"`
	SlideShowID                 string `arg:"slideShowId"`
	AssetID                     string `arg:"assetId"`
	JSON                        string `args:"json"`
	JSONFile                    string `args:"jsonFile"`
	Profile                     string `args:"profile"`
	QRCodeBinary                bool   `args:"qrcodeBinary"`
	QRCodeSize                  int    `args:"qrcodeSize"`
	DaysToExpiration            int    `args:"daysToExpiration"`
	ResetConnection             bool   `args:"resetConnection"`
	RoundingMode                string `args:"roundingMode"`
	Channel                     string `args:"channel"`
	Full                        bool   `args:"full"`
	Archive                     string `args:"archive"`
	Dist                        string `args:"dist"`
	Incremental                 bool   `args:"incremental"`
	ChipRejection               bool   `args:"chipRejection"`
	OutOfOrderReversal          bool   `args:"outOfOrderReversal"`
	AsyncReversals              bool   `args:"asyncReversals"`
	CardOnFile                  bool   `args:"cardOnFile"`
	Recurring                   bool   `args:"recurring"`
	TestCase                    string `args:"testCase"`
	CIT                         bool   `args:"cit"`
	MIT                         bool   `args:"mit"`
	Subscription                bool   `args:"subscription"`
	PONumber                    string `args:"po"`
	SupplierReferenceNumber     string `args:"srn"`
	StatementID                 string `json:"statementId"`
	InvoiceID                   string `json:"invoiceId"`
	ShipmentNumber              int    `json:"shipmentNumber"`
	ShipmentCount               int    `json:"shipmentCount"`
	EntryMethod                 string `json:"entryMethod"`
	DeleteProtected             bool   `json:"deteleProtected"`
	Roles                       string `json:"roles"`
	Notes                       string `json:"notes"`
	Healthcare                  bool   `json:"healthcare"`
	HealthcareTotal             string `json:"healthcareTotal"`
	EBTTotal                    string `json:"ebtTotal"`
	CardMetadataLookup          bool   `json:"cardMetadataLookup"`
	CredType                    string `json:"credType"`
}

var defaultSettings = &ConfigSettings{
	GatewayHost:     "https://api.blockchyp.com",
	TestGatewayHost: "https://test.blockchyp.com",
	DashboardHost:   "https://dashboard.blockchyp.com",
	Secure:          true,
}

// LoadConfigSettings loads settings from the command line and/or the
// configuration file.
func LoadConfigSettings(args CommandLineArguments) (*ConfigSettings, error) {
	fileName := args.ConfigFile
	if fileName == "" {
		var configHome string

		if runtime.GOOS == "windows" {
			configHome = os.Getenv("userprofile")
		} else {
			configHome = os.Getenv("XDG_CONFIG_HOME")
			if configHome == "" {
				user, err := user.Current()
				if err != nil {
					return nil, err
				}
				configHome = user.HomeDir + "/.config"
			}
		}

		fileName = filepath.Join(configHome, ConfigDir, ConfigFile)
	}

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		if args.ConfigFile != "" {
			return nil, errors.New(fileName + " not found")
		}
		return defaultSettings, nil
	}

	b, err := ioutil.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	var profiles map[string]ConfigSettings
	if err = json.Unmarshal(b, &profiles); err != nil {
		if args.Profile != "" {
			return nil, fmt.Errorf("profile `%s` does not exist", args.Profile)
		}

		config := &ConfigSettings{}
		err = json.Unmarshal(b, config)
		return config, err
	}

	if args.Profile == "" {
		args.Profile = "default"
	}

	v, ok := profiles[args.Profile]
	if !ok {
		return nil, errors.New("default profile must be set")
	}
	return &v, nil
}
