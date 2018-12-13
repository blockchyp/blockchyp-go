package blockchyp

import (
	"bytes"
	"crypto/x509"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"time"
)

var (
	routeCache map[string]routeCacheEntry
)

type routeCacheEntry struct {
	TTL   time.Time
	Route TerminalRoute
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

/*
TerminalRoute models route information for a payment terminal.
*/
type TerminalRoute struct {
	TerminalName         string         `json:"terminalName"`
	IPAddress            string         `json:"ipAddress"`
	CloudRelayEnabled    bool           `json:"cloudRelayEnabled"`
	TransientCredentials APICredentials `json:"transientCredentials,omitempty"`
	PublicKey            string         `json:"publicKey"`
	RawKey               RawPublicKey   `json:"rawKey"`
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

/*
resolveTerminalRoute returns the route to the given terminal along with
transient credentials mapped to the given API credentials.
*/
func (client *Client) resolveTerminalRoute(terminalName string) (TerminalRoute, error) {

	route := client.routeCacheGet(terminalName)

	if route == nil {
		path := "/terminal-route?terminal=" + url.QueryEscape(terminalName)
		routeResponse := TerminalRouteResponse{}
		err := client.GatewayGet(path, &routeResponse)
		if err != nil {
			log.Fatal(err)
			return routeResponse.TerminalRoute, err
		}
		if routeResponse.Success {
			route = &routeResponse.TerminalRoute
			client.routeCachePut(*route)
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

	routeCache[terminalRoute.TerminalName] = cacheEntry

}

func (client *Client) routeCacheGet(terminalName string) *TerminalRoute {

	if routeCache == nil {
		return nil
	}
	route, ok := routeCache[terminalName]
	if ok {
		if time.Now().After(route.TTL) {
			return nil
		}
		return &route.Route
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

	content, err := json.Marshal(requestEntity)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", client.assembleTerminalURL(route, path), bytes.NewBuffer(content))
	if err != nil {
		return err
	}

	err = addAPIRequestHeaders(req, client.Credentials)
	if err != nil {
		return err
	}
	resp, err := client.terminalHTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}

	err = consumeResponse(resp, responseEntity)

	return err
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
