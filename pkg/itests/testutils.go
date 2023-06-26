package itests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"testing"
	"time"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"

	blockchyp "github.com/blockchyp/blockchyp-go/v2"
)

// TestDelay is an environment variable constant for integration test delays
const TestDelay = "BC_TEST_DELAY"

const (
	defaultConfigFile = "sdk-itest-config.json"
	defaultConfigDir  = "blockchyp"
)

var (
	testConfig         *TestConfiguration
	lastTransactionID  string
	lastTransactionRef string
)

// TestConfiguration models test configuration
type TestConfiguration struct {
	GatewayHost            string                        `json:"gatewayHost"`
	TestGatewayHost        string                        `json:"testGatewayHost"`
	DashboardHost          string                        `json:"dashboardHost"`
	DefaultTerminalName    string                        `json:"defaultTerminalName"`
	DefaultTerminalAddress string                        `json:"defaultTerminalAddress"`
	APIKey                 string                        `json:"apiKey"`
	BearerToken            string                        `json:"bearerToken"`
	SigningKey             string                        `json:"signingKey"`
	Profiles               map[string]ProfileCredentials `json:"profiles"`
}

// ProfileCredentials model alternate test credentials
type ProfileCredentials struct {
	APIKey      string `json:"apiKey"`
	BearerToken string `json:"bearerToken"`
	SigningKey  string `json:"signingKey"`
}

func processTestDelay(t *testing.T, config *TestConfiguration, testName string) {

	testDelay := os.Getenv(TestDelay)
	if testDelay != "" {
		testDelayInt, err := strconv.Atoi(testDelay)
		if err != nil {
			t.Fatal(err)
		}
		messageRequest := blockchyp.MessageRequest{
			TerminalName: config.DefaultTerminalName,
			Test:         true,
			Message:      fmt.Sprintf("Running Test %v in %v seconds...", testName, testDelay),
		}
		delayClient := config.newTestClient(t, "")
		if _, err := delayClient.Message(messageRequest); err != nil {
			t.Fatal(err)
		}
		time.Sleep(time.Duration(testDelayInt) * time.Second)
	}

}

func loadTestConfiguration(t *testing.T) *TestConfiguration {

	assert := assert.New(t)

	var configHome string

	if runtime.GOOS == "windows" {
		configHome = os.Getenv("userprofile")
	} else {
		configHome = os.Getenv("XDG_CONFIG_HOME")
		if configHome == "" {
			configHome = filepath.Join(os.Getenv("HOME"), ".config")
		}
	}

	fileName := filepath.Join(configHome, defaultConfigDir, defaultConfigFile)

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		assert.NoError(err)
	}

	b, err := ioutil.ReadFile(fileName)

	assert.NoError(err)

	config := TestConfiguration{}

	err = json.Unmarshal(b, &config)
	if err != nil {
		t.Error(err)
	}

	return &config
}

func updateLastTransaction(response interface{}) string {

	el := reflect.ValueOf(response).Elem()
	val := el.FieldByName("TransactionID")
	lastTransactionID = val.String()
	val = el.FieldByName("TransactionRef")
	lastTransactionRef = val.String()

	return ""
}

func (c *TestConfiguration) newTestClient(t *testing.T, profile string) blockchyp.Client {
	creds := blockchyp.APICredentials{
		APIKey:      c.APIKey,
		BearerToken: c.BearerToken,
		SigningKey:  c.SigningKey,
	}

	altCreds, ok := c.Profiles[profile]
	if ok {
		creds.APIKey = altCreds.APIKey
		creds.BearerToken = altCreds.BearerToken
		creds.SigningKey = altCreds.SigningKey
	}

	client := blockchyp.NewClient(creds)
	client.HTTPS = false
	client.GatewayHost = c.GatewayHost
	client.DashboardHost = c.DashboardHost
	client.TestGatewayHost = c.TestGatewayHost

	logObj(t, "Client:", client)

	return client
}

func randomID() string {

	u2, err := uuid.NewV1()
	if err != nil {
		log.Fatal(err)
	}
	return u2.String()

}

func logObj(t *testing.T, args ...interface{}) {
	var fmtStr string
	var content []byte
	for _, arg := range args {
		switch v := arg.(type) {
		case string:
			fmtStr += v + " "
		default:
			content, _ = json.MarshalIndent(arg, "", " ")
		}
	}

	fmtStr += "%s"

	t.Logf(fmtStr, string(content))
}
