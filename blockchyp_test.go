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
	testPAN = "4111111111111111"
	testTrack1 = "B4111111111111111^SATOSHI/NAKAMOTO^2512101000003280"
	testTrack2 = "4111111111111111=2512101000003280"
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

func TestMSRCharge(t *testing.T) {

	request := AuthorizationRequest{}
	request.Amount = "43.55"
	request.Track1 = testTrack1
	request.Track2 = testTrack2

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

	if !response.Approved {
		t.Error("transaction was not approved")
	}

	if response.EntryMethod != "SWIPE" {
		t.Error("entry method not swipe")
	}

	if response.PaymentType != "VISA" {
		t.Error("payment type not VISA")
	}

	if !response.ScopeAlert {
		t.Error("transaction failed to trigger scope alert")
	}

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
