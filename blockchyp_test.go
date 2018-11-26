// +build manual

package blockchyp

import (
	"testing"
  "log"
  "encoding/json"
	"io/ioutil"
)

const (
	defaultConfigLocation = "/etc/blockchyp/sdk-itest-config.json"
)

var (
	testConfig *TestConfiguration
)

type TestConfiguration struct {
	GatewayHost string `json:"gatewayHost"`
	DefaultTerminalName string `json:"defaultTerminalName"`
	DefaultTerminalAddress string `json:"defaultTerminalAddress"`
	APIKey string `json:"apiKey"`
	BearerToken string `json:"bearerToken"`
	SigningKey string `json:"signingKey"`
}

func loadTestConfiguration(t *testing.T) TestConfiguration {

	if testConfig == nil {


		content, err := ioutil.ReadFile(defaultConfigLocation)
		if err != nil {
			t.Error(err)
		}

		testConfig = &TestConfiguration{}

		err = json.Unmarshal(content, testConfig)
		if err != nil {
			t.Error(err)
		}

	}

	return *testConfig

}

func newTestClient(t *testing.T) Client {

	config := loadTestConfiguration(t)

	creds := APICredentials{
		APIKey: config.APIKey,
		BearerToken: config.BearerToken,
		SigningKey: config.SigningKey,
	}

	client := NewClient(creds)
	client.HTTPS = false
	client.GatewayHost = config.GatewayHost

	return client

}

func TestMinimalCharge(t *testing.T) {

	config := loadTestConfiguration(t)

  request := AuthorizationRequest{}
  request.Amount = "20.55"
  request.TerminalName = config.DefaultTerminalName


  content, err := json.Marshal(request)

	if err != nil {
		t.Error(err)
	}


  log.Println("SDK Request:", string(content))

	client := newTestClient(t)

	response, err := client.Charge(request)

	if err != nil {
		t.Error(err)
	}

	content, err = json.Marshal(response)

	if err != nil {
		t.Error(err)
	}

	log.Println("SDK Response:", string(content))

	if response.TransactionID == "" {
		t.Error("transaction id not returned")
	}

	if response.Approved == false {
		t.Error("transaction was not approved")
	}

}
