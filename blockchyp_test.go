// +build manual

package blockchyp

import (
	"testing"
  "log"
  "encoding/json"
	"io/ioutil"

	"github.com/stretchr/testify/assert"
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

func TestMSRVoid(t *testing.T) {

	request := AuthorizationRequest{}
	request.Amount = "45.00"
	request.Track1 = testTrack1
	request.Track2 = testTrack2

	logRequest(request)

	client := newTestClient(t)

	response, err := client.Preauth(request)

	if err != nil {
		t.Error(err)
	}

	logResponse(response)

	assertConventionalApproval(t, *response)

	voidRequest := VoidRequest{}
	voidRequest.TransactionID = response.TransactionID


	logRequest(voidRequest)

	voidResponse, err := client.Void(voidRequest)

	if err != nil {
		t.Error(err)
	}

	logResponse(voidResponse)

	assertVoidApproval(t, *voidResponse)

}


func TestMSRPreauth(t *testing.T) {

	request := AuthorizationRequest{}
	request.Amount = "45.00"
	request.Track1 = testTrack1
	request.Track2 = testTrack2

	logRequest(request)

	client := newTestClient(t)

	response, err := client.Preauth(request)

	if err != nil {
		t.Error(err)
	}

	logResponse(response)

	assertConventionalApproval(t, *response)

	captureRequest := CaptureRequest{}
	captureRequest.TransactionID = response.TransactionID
	captureRequest.Amount = "50.00"
	captureRequest.TipAmount = "5.00"

	logRequest(captureRequest)

	captureResponse, err := client.Capture(captureRequest)

	if err != nil {
		t.Error(err)
	}

	logResponse(captureResponse)

	assertCaptureApproval(t, *captureResponse)

}

func logRequest(request interface{}) {
	content, _ := json.Marshal(request)
	log.Println("Request:", string(content))
}


func logResponse(response interface{}) {
	content, _ := json.Marshal(response)
	log.Println("Response:", string(content))
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

	assertConventionalApproval(t, *response)

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

	assertConventionalApproval(t, *response)
	assert := assert.New(t)

	assert.Equal("VISA", response.PaymentType)
	assert.Equal("SWIPE", response.EntryMethod)
	assert.Contains(response.MaskedPAN, "*1111")

}

func assertVoidApproval(t *testing.T, response VoidResponse) {

	assert := assert.New(t)

	assert.True(response.Approved)
	assert.NotEmpty(response.TransactionID)
	assert.NotEmpty(response.PaymentType)
	assert.NotEmpty(response.EntryMethod)
	assert.Equal("Approved", response.ResponseDescription)
	assert.NotEmpty(response.Timestamp)
	assert.NotEmpty(response.TickBlock)

}

func assertCaptureApproval(t *testing.T, response CaptureResponse) {

	assert := assert.New(t)

	assert.True(response.Approved)
	assert.False(response.PartialAuth)
	assert.NotEmpty(response.TransactionID)
	assert.NotEmpty(response.PaymentType)
	assert.NotEmpty(response.EntryMethod)
	assert.Equal("Approved", response.ResponseDescription)
	assert.Equal("USD", response.CurrencyCode)
	assert.NotEmpty(response.Timestamp)
	assert.NotEmpty(response.TickBlock)
	assert.NotEmpty(response.RequestedAmount)
	assert.NotEmpty(response.AuthorizedAmount)
	assert.NotEmpty(response.TipAmount)
	assert.NotEmpty(response.TaxAmount)

}

func assertConventionalApproval(t *testing.T, response AuthorizationResponse) {

	assert := assert.New(t)

	assert.True(response.Approved)
	assert.False(response.PartialAuth)
	assert.NotEmpty(response.TransactionID)
	assert.NotEmpty(response.PaymentType)
	assert.NotEmpty(response.EntryMethod)
	assert.Equal("Approved", response.ResponseDescription)
	assert.Equal("USD", response.CurrencyCode)
	assert.NotEmpty(response.Timestamp)
	assert.NotEmpty(response.TickBlock)
	assert.NotEmpty(response.RequestedAmount)
	assert.NotEmpty(response.AuthorizedAmount)
	assert.NotEmpty(response.TipAmount)
	assert.NotEmpty(response.TaxAmount)


}
