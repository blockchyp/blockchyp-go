package blockchyp

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"time"
)

/*
TerminalRouteResponse models a terminal route response from the gateway.
*/
type TerminalRouteResponse struct {
	TerminalRoute
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

/*
APIRequestHeaders models the http headers required for BlockChyp API requests.
*/
type APIRequestHeaders struct {
	Timestamp   string
	Nonce       string
	BearerToken string
	APIKey      string
	Signature   string
}

func (client *Client) assembleGatewayURL(path string, testTx bool) string {

	buffer := bytes.Buffer{}

	if testTx {
		if len(client.TestGatewayHost) > 0 {
			buffer.WriteString(client.TestGatewayHost)
		} else {
			buffer.WriteString(DefaultTestGatewayHost)
		}
	} else {
		if len(client.GatewayHost) > 0 {
			buffer.WriteString(client.GatewayHost)
		} else {
			buffer.WriteString(DefaultGatewayHost)
		}
	}
	buffer.WriteString("/api")
	buffer.WriteString(path)
	return buffer.String()

}

func consumeResponse(resp *http.Response, responseEntity interface{}) error {

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	err = json.Unmarshal(b, responseEntity)

	if err != nil {
		return err
	}

	return nil
}

// GatewayRequest sends a gateway request with the default timeout.
func (client *Client) GatewayRequest(path, method string, request, response interface{}, testTx bool) error {
	return client.GatewayRequestWithTimeout(path, method, request, response, testTx, DefaultGatewayTimeout)
}

// RelayRequest sends a gateway request with the cloud relay timeout.
func (client *Client) RelayRequest(path, method string, request, response interface{}, testTx bool) error {
	return client.GatewayRequestWithTimeout(path, method, request, response, testTx, DefaultTerminalTimeout)
}

// GatewayRequestWithTimeout sends an HTTP request to the gateway using a
// context to set timeout per request.
func (client *Client) GatewayRequestWithTimeout(path, method string, request, response interface{}, testTx bool, timeout time.Duration) error {
	content, err := json.Marshal(request)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(method, client.assembleGatewayURL(path, testTx), bytes.NewBuffer(content))
	if err != nil {
		return err
	}

	if err := addAPIRequestHeaders(req, client.Credentials); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(req.Context(), timeout)
	defer cancel()

	res, err := client.gatewayHTTPClient.Do(req.WithContext(ctx))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusForbidden {
		//on 403's we check time diffs in order to help troubleshoot time issues
		if client.highClockDiff() {
			return errors.New("high clock drift, reset time on client")
		}
	}

	if res.StatusCode != http.StatusOK {
		return errors.New(res.Status)
	}

	return consumeResponse(res, response)
}

/*
GatewayPost posts a request to the api gateway.
*/
func (client *Client) GatewayPost(path string, requestEntity interface{}, responseEntity interface{}, testTx bool) error {
	return client.GatewayRequest(path, http.MethodPost, requestEntity, responseEntity, testTx)
}

/*
GatewayGet retrieves a get request from the api gateway.
*/
func (client *Client) GatewayGet(path string, responseEntity interface{}) error {
	return client.GatewayRequest(path, http.MethodGet, nil, responseEntity, false)
}

func (client *Client) highClockDiff() bool {

	response := HeartbeatResponse{}
	err := client.GatewayRequest("/heartbeat", http.MethodGet, nil, &response, false)
	if err != nil {
		return false
	}

	dur := time.Since(response.Timestamp)

	if math.Abs(dur.Minutes()) > 10 {
		return true
	}

	return false

}

/*
PopulateHeaders takes header values and adds them to the given http request.
*/
func populateHeaders(headers APIRequestHeaders, req *http.Request) {

	req.Header.Add("Nonce", headers.Nonce)
	req.Header.Add("Timestamp", headers.Timestamp)
	req.Header.Add("Authorization", "Dual "+headers.BearerToken+":"+headers.APIKey+":"+headers.Signature)

}

func addAPIRequestHeaders(req *http.Request, creds APICredentials) error {

	headers, err := generateAPIRequestHeaders(creds)
	if err != nil {
		return err
	}
	populateHeaders(headers, req)
	return nil

}

/*
generateAPIRequestHeaders returns the standard API requests headers given a set of
credentials.
*/
func generateAPIRequestHeaders(creds APICredentials) (APIRequestHeaders, error) {

	headers := APIRequestHeaders{
		APIKey:      creds.APIKey,
		BearerToken: creds.BearerToken,
	}
	headers.Nonce = generateNonce()
	headers.Timestamp = time.Now().UTC().Format(time.RFC3339)

	sig, err := computeHmac(headers, creds.SigningKey)

	if err != nil {
		return headers, err
	}
	headers.Signature = sig

	return headers, nil

}

/*
ComputeHmac computes an hmac for the the given headers and secret key.
*/
func computeHmac(headers APIRequestHeaders, signingKey string) (string, error) {

	buf := bytes.Buffer{}

	buf.WriteString(headers.APIKey)
	buf.WriteString(headers.BearerToken)
	buf.WriteString(headers.Timestamp)
	buf.WriteString(headers.Nonce)

	key, err := hex.DecodeString(signingKey)

	if err != nil {
		return "", errors.New("Malformed Signing Key")
	}

	mac := hmac.New(sha256.New, key)
	mac.Write(buf.Bytes())
	hash := mac.Sum(nil)

	return hex.EncodeToString(hash), nil

}

func generateNonce() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]byte, 32)
	r.Read(result)
	return hex.EncodeToString(result)
}
