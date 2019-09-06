package blockchyp

import (
	"bytes"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
)

var (
	routeCache map[string]routeCacheEntry
)

const (
	offlineFixedKey = "cb22789c9d5c344a10e0474f134db39e25eb3bbf5a1b1a5e89b507f15ea9519c"
)

type routeCacheEntry struct {
	TTL   time.Time
	Route TerminalRoute
}

/*
TerminalTermsAndConditionsRequest adds API credentials to auth requests for use in
direct terminal transactions.
*/
type TerminalTermsAndConditionsRequest struct {
	APICredentials
	Request TermsAndConditionsRequest `json:"request"`
}

/*
TerminalAuthorizationRequest adds API credentials to auth requests for use in
direct terminal transactions.
*/
type TerminalAuthorizationRequest struct {
	APICredentials
	Request AuthorizationRequest `json:"request"`
}

/*
TerminalMessageRequest adds API credentials to auth requests for use in
direct terminal message display requests.
*/
type TerminalMessageRequest struct {
	APICredentials
	Request MessageRequest `json:"request"`
}

/*
TerminalClearTerminalRequest adds API credentials to a clear terminal request.
*/
type TerminalClearTerminalRequest struct {
	APICredentials
	Request ClearTerminalRequest `json:"request"`
}

/*
TerminalBooleanPromptRequest adds API credentials to boolean prompt requests.
*/
type TerminalBooleanPromptRequest struct {
	APICredentials
	Request BooleanPromptRequest `json:"request"`
}

/*
TerminalTextPromptRequest adds API credentials to boolean prompt requests.
*/
type TerminalTextPromptRequest struct {
	APICredentials
	Request TextPromptRequest `json:"request"`
}

/*
TerminalRefundAuthorizationRequest adds API credentials to refund requests for use in
free range refund transactions.
*/
type TerminalRefundAuthorizationRequest struct {
	APICredentials
	Request RefundRequest `json:"request"`
}

/*
TerminalGiftActivateRequest adds API credentials to gift activation requests
for use in direct terminal transactions.
*/
type TerminalGiftActivateRequest struct {
	APICredentials
	Request GiftActivateRequest `json:"request"`
}

/*
TerminalPingRequest adds API credentials to a terminal ping request.
*/
type TerminalPingRequest struct {
	APICredentials
	Request PingRequest `json:"request"`
}

// TerminalTransactionDisplayRequest adds API credentials to a terminal line
// item display request.
type TerminalTransactionDisplayRequest struct {
	APICredentials
	Request TransactionDisplayRequest `json:"request"`
}

/*
TerminalRoute models route information for a payment terminal.
*/
type TerminalRoute struct {
	Exists               bool           `json:"exists"`
	TerminalName         string         `json:"terminalName"`
	IPAddress            string         `json:"ipAddress"`
	CloudRelayEnabled    bool           `json:"cloudRelayEnabled"`
	TransientCredentials APICredentials `json:"transientCredentials,omitempty"`
	PublicKey            string         `json:"publicKey"`
	RawKey               RawPublicKey   `json:"rawKey"`
	Timestamp            time.Time      `json:"timestamp"`
}

/*
RouteCache models offline route cache information.
*/
type RouteCache struct {
	Routes map[string]routeCacheEntry `json:"routes"`
}

/*
RawPublicKey models the primitive form of an ECC public key.  A little
simpler than X509, ASN and the usual nonsense.
*/
type RawPublicKey struct {
	Curve string `json:"curve"`
	X     string `json:"x"`
	Y     string `json:"Y"`
}

func (client *Client) readFromOfflineCache(terminalName string) *routeCacheEntry {

	cache := client.readOfflineCache()

	if cache == nil {
		return nil
	}

	route, ok := cache.Routes[client.Credentials.APIKey+terminalName]
	if ok {
		return &route
	}

	return nil

}

func (client *Client) readOfflineCache() *RouteCache {

	if _, err := os.Stat(client.RouteCache); os.IsNotExist(err) {
		return nil
	}

	content, err := ioutil.ReadFile(client.RouteCache)

	if err != nil {
		fmt.Print(err)
		return nil
	}

	cache := RouteCache{}

	err = json.Unmarshal(content, &cache)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &cache
}

/*
resolveTerminalRoute returns the route to the given terminal along with
transient credentials mapped to the given API credentials.
*/
func (client *Client) resolveTerminalRoute(terminalName string) (TerminalRoute, error) {

	route := client.routeCacheGet(terminalName)

	if route == nil {
		path := "/terminal-route?terminal=" + url.QueryEscape(terminalName)
		routeResponse := TerminalRouteResponse{}
		err := client.GatewayRequest(path, http.MethodGet, nil, &routeResponse, false)
		if err != nil {
			return routeResponse.TerminalRoute, err
		}
		if routeResponse.Success {
			route = &routeResponse.TerminalRoute
			route.Exists = true
			if len(route.IPAddress) > 0 {
				client.routeCachePut(*route)
			}
		} else {
			route = &TerminalRoute{}
			route.Exists = false
		}
	}

	return *route, nil

}

func (client *Client) routeCachePut(terminalRoute TerminalRoute) {

	if routeCache == nil {
		routeCache = make(map[string]routeCacheEntry)
	}

	cacheEntry := routeCacheEntry{
		Route: terminalRoute,
		TTL:   time.Now().Add(client.routeCacheTTL),
	}

	routeCache[client.Credentials.APIKey+terminalRoute.TerminalName] = cacheEntry

	go client.updateOfflineCache(&cacheEntry)

}

func (client *Client) deriveOfflineKey() []byte {

	hash := sha256.New()
	fixedKey, err := hex.DecodeString(offlineFixedKey)
	if err != nil {
		fmt.Println(err)
		return []byte{}
	}
	hash.Write(fixedKey)
	dynamicKey, err := hex.DecodeString(client.Credentials.SigningKey)
	if err != nil {
		fmt.Println(err)
		return []byte{}
	}
	hash.Write(dynamicKey)
	return hash.Sum(nil)

}

func (client *Client) encrypt(value string) string {

	key := client.deriveOfflineKey()
	return Encrypt(key, value)

}

func (client *Client) decrypt(value string) string {

	key := client.deriveOfflineKey()
	return Decrypt(key, value)

}

func (client *Client) updateOfflineCache(cacheEntry *routeCacheEntry) {

	cache := client.readOfflineCache()

	if cache == nil {
		cache = &RouteCache{}
	}

	if cache.Routes == nil {
		cache.Routes = make(map[string]routeCacheEntry)
	}

	cacheEntry.Route.TransientCredentials.APIKey = client.encrypt(cacheEntry.Route.TransientCredentials.APIKey)
	cacheEntry.Route.TransientCredentials.BearerToken = client.encrypt(cacheEntry.Route.TransientCredentials.BearerToken)
	cacheEntry.Route.TransientCredentials.SigningKey = client.encrypt(cacheEntry.Route.TransientCredentials.SigningKey)

	cache.Routes[client.Credentials.APIKey+cacheEntry.Route.TerminalName] = *cacheEntry

	content, err := json.Marshal(cache)

	if err != nil {
		fmt.Print(err)
		return
	}

	err = ioutil.WriteFile(client.RouteCache, content, 0644)

	if err != nil {
		fmt.Print(err)
	}

}

func (client *Client) routeCacheGet(terminalName string) *TerminalRoute {

	if routeCache != nil {
		route, ok := routeCache[terminalName]
		if ok {
			if time.Now().After(route.TTL) {
				return nil
			}
			return &route.Route
		}
	}

	cacheEntry := client.readFromOfflineCache(terminalName)

	//check expiry
	if cacheEntry != nil {
		if time.Now().After(cacheEntry.TTL) {
			return nil
		}
		cacheEntry.Route.TransientCredentials.APIKey = client.decrypt(cacheEntry.Route.TransientCredentials.APIKey)
		cacheEntry.Route.TransientCredentials.BearerToken = client.decrypt(cacheEntry.Route.TransientCredentials.BearerToken)
		cacheEntry.Route.TransientCredentials.SigningKey = client.decrypt(cacheEntry.Route.TransientCredentials.SigningKey)
		return &cacheEntry.Route
	}

	return nil

}

func isTerminalRouted(auth PaymentMethod) bool {
	return auth.TerminalName != ""
}

func (client *Client) assembleTerminalURL(route TerminalRoute, path string) string {

	buffer := bytes.Buffer{}
	if client.HTTPS {
		buffer.WriteString("https://")
	} else {
		buffer.WriteString("http://")
	}
	buffer.WriteString(route.IPAddress)
	if client.HTTPS {
		buffer.WriteString(":8443")
	} else {
		buffer.WriteString(":8080")
	}
	buffer.WriteString("/api")
	buffer.WriteString(path)
	return buffer.String()

}

/*
terminalPost posts a request to the api gateway.
*/
func (client *Client) terminalPost(route TerminalRoute, path string, requestEntity interface{}, responseEntity interface{}) error {
	return client.terminalRequest(route, path, http.MethodPost, requestEntity, responseEntity)
}

// terminalRequest sends an HTTP request to a terminal.
func (client *Client) terminalRequest(route TerminalRoute, path, method string, requestEntity, responseEntity interface{}) error {
	content, err := json.Marshal(requestEntity)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(method, client.assembleTerminalURL(route, path), bytes.NewBuffer(content))
	if err != nil {
		return err
	}

	err = addAPIRequestHeaders(req, client.Credentials)
	if err != nil {
		return err
	}

	res, err := client.terminalHTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return errors.New(res.Status)
	}

	return consumeResponse(res, responseEntity)
}

func terminalCertPool() *x509.CertPool {
	pool := x509.NewCertPool()

	pool.AppendCertsFromPEM([]byte(terminalRootCA))

	return pool
}

const terminalRootCA = `
-----BEGIN CERTIFICATE-----
MIIFAjCCAuqgAwIBAgIBATANBgkqhkiG9w0BAQsFADAgMR4wHAYDVQQDDBVCbG9j
a0NoeXAgSW50ZXJuYWwgQ0EwIBcNMTgwMTAxMDAwMDA4WhgPNDc1NTExMjkwMDAw
MDhaMCAxHjAcBgNVBAMMFUJsb2NrQ2h5cCBJbnRlcm5hbCBDQTCCAiIwDQYJKoZI
hvcNAQEBBQADggIPADCCAgoCggIBANyWuVhDiqeCrHMxbTv5PN5UOZdR8n4PPwUV
z0dALnLS7Lkl9nnuBxUK5XFGsZHBQ3GqSsWgA0HBUAAkKY/hzDIY+mrKOTMFMhoF
SKmcNwmdt+NXuUtYwL5STsr1U/XnxcizsSEHcGP5LhIH16AY0XYMVzNTBXrylH7O
Hf/pPJaVbuywAkiyrEV+lTo1mVTOCucGoNRPogluuyfbBCUH9bWBajbjHWdyiX58
IV786JWkw5ogLXgDekrrzdVxQH1t2kN2PvXNHGOBlB0NL/QwKHxfbvgIu6EkyEXv
vSuFclgaM3x38zcEaIS8id/wZYkwZXAqquR5Hi5fqPILC1xmRF+zC1GH1uJ+gsQu
wqwaiwmD9Rcbm2ZOSVntQy5bCF7IzPlMHzMlt33dF9mZo9bJwFO1APdpeWy+Ooga
n1k/yS2EPnkAv+DiRpNf2it6n86+X7Z4C6QGgP5+rfc53uxeaF8gPLgXViaHHTZD
NflxaNjgKD0xAwB3Yhca8RQSjRPwKYk1FrbhTSAIidnwmA4jrV7juZ2RSWA99VzR
O68OmE/7NygxGgo995pPc+s6DO6IOnZvT2tSs0b2UmEKT51/cf93lv+phX/69hTC
ctMEYoIGNRAvcISA0lfTWHAbiRzMyagtuiRMttS7C+IshsgBrjHSHMsEYj8RhRnR
0FvmChUNAgMBAAGjRTBDMB0GA1UdDgQWBBSBl1rnpf7Omve8fXPl9EltnlcqGTAS
BgNVHRMBAf8ECDAGAQH/AgEBMA4GA1UdDwEB/wQEAwIBBjANBgkqhkiG9w0BAQsF
AAOCAgEAkt9ywLJvM0TjEUjlC32niE8mNIPX5azHJ0++PlZ2Fc7ZKy4nntt2YErl
l4qEOB8ED2VaLQuxx0O9H2oh1QsMuxT3rQ4SDNmQVH9vUYJWgIkYjY1zKubEyktv
oZyi8xK5e0/ME//vU0ru6y0dmcFtDvpwm/JZPjoVKHK58JpCKH8xhVxQo7NxAIf8
Ow+fr58plDQP1CbfjO1gJpFg7lQ282rz9n0Ju2mXm3guclcx74mDJGlzGLGCJCnu
Qxta8Dv/Cg8+kNM36boORMChaoAgIerXL17EhyUh3ZsSaxEchqvCWtLv1+ekhGpF
A08xS33r1GgQV/cyunuz3czQ0Y/7UjKluo6sbS0RmVtAWJA/DhwXgQlHlFyROmhG
pcKXeLc7+LrBZxITVuQk8Mg9aceAnzBqjeTjQNPQJkOwqIFgDUXNNqvA5mhn/j25
u8CcDY/0p5C4tFQc1npgQwJZAwRGEvFmXVDgEZ8FFkzhn74oxI99Xs1HGc9zO/uP
GV0cahaj9xspMPMBe5Q2mNhVca6+RIZPSIcVbsgYy+2QDBep7NpraQgG7V0f2XTu
uLBaPXbY9PZLFklSSZOLXAuuOk0G57lfyVFRNAZ2R3uQdkDpx90Ti6PDWj9M6x1p
jD1XNpXvgH2k91jjsK67khN+4bWoFBsfrMYt6vgjtXyv0kf12y0=
-----END CERTIFICATE-----
`
