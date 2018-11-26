package blockchyp

import (
  "bytes"
  "net/http"
  "time"
  "encoding/hex"
  "encoding/json"
  "math/rand"
  "crypto/hmac"
  "crypto/sha256"
  "io/ioutil"
)

/*
TerminalRouteResponse models a terminal route response from the gateway.
*/
type TerminalRouteResponse struct {
  TerminalRoute
  Success bool `json:"success"`
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


func (client *Client) assembleFullURL(path string) string {

  buffer := bytes.Buffer{}

  buffer.WriteString(client.GatewayHost)
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

func (client *Client) GatewayPost(path string, requestEntity interface{}, responseEntity interface{}) error {

  httpClient := &http.Client{}

  content, err := json.Marshal(requestEntity)
  if err != nil {
    return err
  }

  req, err := http.NewRequest("POST", client.assembleFullURL(path), bytes.NewBuffer(content))
  if err != nil {
    return err
  }

  err = addAPIRequestHeaders(req, client.Credentials)
  if err != nil {
    return err
  }
  resp, err := httpClient.Do(req)
  if err != nil {
    return err
  }
  defer resp.Body.Close()

	err = consumeResponse(resp, responseEntity)

  return err
}

func (client *Client) GatewayGet(path string, responseEntity interface{}) error {

  httpClient := &http.Client{}

  req, err := http.NewRequest("GET", client.assembleFullURL(path), nil)
  if err != nil {
    return err
  }

  err = addAPIRequestHeaders(req, client.Credentials)
  if err != nil {
    return err
  }
  resp, err := httpClient.Do(req)
  if err != nil {
    return err
  }
  defer resp.Body.Close()

	err = consumeResponse(resp, responseEntity)

  return err
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
    APIKey: creds.APIKey,
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
		return "", err
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
