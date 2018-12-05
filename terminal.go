package blockchyp

import (
	"bytes"
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

	log.Println("Resolving terminal route...")

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

	content, _ := json.Marshal(route)

	log.Println(string(content))

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

	httpClient := &http.Client{}

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
	resp, err := httpClient.Do(req)
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

const terminalRootCA = `
-----BEGIN CERTIFICATE-----
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
lx3y1Bdb/A3T0axxAwax4jhNQbDQ2dqyXN0=
-----END CERTIFICATE-----
`
