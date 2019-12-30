package itests

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"

	blockchyp "github.com/blockchyp/blockchyp-go"
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

//TestConfiguration models test configuration
type TestConfiguration struct {
	GatewayHost            string `json:"gatewayHost"`
	TestGatewayHost        string `json:"testGatewayHost"`
	DefaultTerminalName    string `json:"defaultTerminalName"`
	DefaultTerminalAddress string `json:"defaultTerminalAddress"`
	APIKey                 string `json:"apiKey"`
	BearerToken            string `json:"bearerToken"`
	SigningKey             string `json:"signingKey"`
}

func loadTestConfiguration(t *testing.T) *TestConfiguration {

	assert := assert.New(t)

	var configHome string

	if runtime.GOOS == "windows" {
		configHome = os.Getenv("userprofile")
	} else {
		configHome = os.Getenv("XDG_CONFIG_HOME")
		if configHome == "" {
			user, err := user.Current()
			if err != nil {
				assert.NoError(err)
			}
			configHome = user.HomeDir + "/.config"
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

func newTestClient(t *testing.T) blockchyp.Client {

	config := loadTestConfiguration(t)

	creds := blockchyp.APICredentials{
		APIKey:      config.APIKey,
		BearerToken: config.BearerToken,
		SigningKey:  config.SigningKey,
	}

	client := blockchyp.NewClient(creds)
	client.HTTPS = false
	client.GatewayHost = config.GatewayHost
	client.TestGatewayHost = config.TestGatewayHost

	log.Printf("%+v\n", client)

	return client

}

func randomID() string {

	u2, err := uuid.NewV1()
	if err != nil {
		log.Fatal(err)
	}
	return u2.String()

}

func logRequest(request interface{}) {
	content, _ := json.Marshal(request)
	log.Println("Request:", string(content))
}

func logResponse(response interface{}) {
	updateLastTransaction(response)
	content, _ := json.Marshal(response)
	log.Println("Response:", string(content))
}
