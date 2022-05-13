package blockchyp

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"os"
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

func (client *Client) assembleDashboardURL(path string) string {

	buffer := bytes.Buffer{}

	if len(client.GatewayHost) > 0 {
		buffer.WriteString(client.DashboardHost)
	} else {
		buffer.WriteString(DefaultDashboardHost)
	}

	buffer.WriteString(path)
	return buffer.String()

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
	buffer.WriteString(path)
	return buffer.String()

}

func consumeResponse(resp *http.Response, responseEntity interface{}) error {

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		var ack Acknowledgement
		err := json.Unmarshal(b, &ack)
		if err != nil || ack.Error == "" {
			return errors.New(resp.Status)
		}
		return errors.New(ack.Error)
	}

	err = json.Unmarshal(b, responseEntity)

	if err != nil {
		return err
	}

	return nil
}

// DashboardRequest sends a gateway request with the default timeout.
func (client *Client) DashboardRequest(path, method string, request, response interface{}, requestTimeout interface{}) error {
	content, err := json.Marshal(request)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(method, client.assembleDashboardURL(path), bytes.NewBuffer(content))
	if err != nil {
		return err
	}

	if err := addAPIRequestHeaders(req, client.Credentials); err != nil {
		return err
	}

	timeout := getTimeout(requestTimeout, client.GatewayTimeout)
	ctx, cancel := context.WithTimeout(req.Context(), timeout)
	defer cancel()

	req = req.WithContext(ctx)

	if client.LogRequests {
		b, err := httputil.DumpRequestOut(req, true)
		if err != nil {
			return err
		}
		fmt.Fprintln(os.Stderr, "DASHBOARD REQUEST:")
		fmt.Fprintln(os.Stderr, string(b))
	}

	res, err := client.gatewayHTTPClient.Do(req)
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

	return consumeResponse(res, response)
}

// GatewayRequest sends a gateway request with the default timeout.
func (client *Client) GatewayRequest(path, method string, request, response interface{}, testTx bool, requestTimeout interface{}) error {
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

	timeout := getTimeout(requestTimeout, client.GatewayTimeout)
	ctx, cancel := context.WithTimeout(req.Context(), timeout)
	defer cancel()

	req = req.WithContext(ctx)

	if client.LogRequests {
		b, err := httputil.DumpRequestOut(req, true)
		if err != nil {
			return err
		}
		fmt.Fprintln(os.Stderr, "GATEWAY REQUEST:")
		fmt.Fprintln(os.Stderr, string(b))
	}

	res, err := client.gatewayHTTPClient.Do(req)
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

	return consumeResponse(res, response)
}

// RelayRequest sends a request to the gateway to be relayed to a terminal.
func (client *Client) RelayRequest(path, method string, request, response interface{}, testTx bool, requestTimeout interface{}) error {
	timeout := getTimeout(requestTimeout, client.TerminalTimeout)
	return client.GatewayRequest(path, method, request, response, testTx, timeout)
}

// GatewayPost posts a request to the api gateway.
func (client *Client) GatewayPost(path string, requestEntity interface{}, responseEntity interface{}, testTx bool) error {
	return client.GatewayRequest(path, http.MethodPost, requestEntity, responseEntity, testTx, nil)
}

// GatewayGet retrieves a get request from the api gateway.
func (client *Client) GatewayGet(path string, responseEntity interface{}) error {
	return client.GatewayRequest(path, http.MethodGet, nil, responseEntity, false, nil)
}

func (client *Client) highClockDiff() bool {

	response := HeartbeatResponse{}
	err := client.GatewayRequest("/api/heartbeat", http.MethodGet, nil, &response, false, nil)
	if err != nil {
		return false
	}

	dur := time.Since(response.Timestamp)

	if math.Abs(dur.Minutes()) > 10 {
		return true
	}

	return false

}

// PopulateHeaders takes header values and adds them to the given http request.
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

// generateAPIRequestHeaders returns the standard API requests headers given a set of
// credentials.
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

// ComputeHmac computes an hmac for the the given headers and secret key.
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
