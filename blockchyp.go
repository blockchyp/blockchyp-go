// Copyright 2019 BlockChyp, Inc. All rights reserved. Use of this code is
// governed by a license that can be found in the LICENSE file.
//
// This file was generated automatically. Changes to this file will be lost
// every time the code is regenerated.

package blockchyp

import (
	"crypto/tls"
	"errors"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// Default client configuration constants.
const (
	DefaultGatewayHost     = "https://api.blockchyp.com"
	DefaultTestGatewayHost = "https://test.blockchyp.com"
	DefaultHTTPS           = true
	DefaultRouteCacheTTL   = 60 * time.Minute
	DefaultGatewayTimeout  = 20 * time.Second
	DefaultTerminalTimeout = 2 * time.Minute
)

// Default filesystem configuration.
const (
	ConfigDir  = "blockchyp"
	ConfigFile = "blockchyp.json"
)

// terminalCN is the common name on a terminal certificate.
const terminalCN = "blockchyp-terminal"

// Clientside response constants.
const (
	ResponseUnknownTerminal = "Unknown Terminal"
	ResponseTimedOut        = "Request Timed Out"
)

// ErrInvalidAsyncRequest is returned when a request cannot be called
// asynchronously.
var ErrInvalidAsyncRequest = errors.New("async requests must be terminal requests")

// Version contains the version at build time
var Version string

// Client is the main interface used by application developers.
type Client struct {
	Credentials     APICredentials
	GatewayHost     string
	TestGatewayHost string
	HTTPS           bool
	RouteCache      string

	GatewayTimeout  time.Duration
	TerminalTimeout time.Duration

	routeCacheTTL      time.Duration
	gatewayHTTPClient  *http.Client
	terminalHTTPClient *http.Client
}

// NewClient returns a default Client configured with the given credentials.
func NewClient(creds APICredentials) Client {
	return Client{
		Credentials: creds,
		GatewayHost: DefaultGatewayHost,
		HTTPS:       DefaultHTTPS,
		RouteCache:  filepath.Join(os.TempDir(), ".blockchyp_routes"),

		GatewayTimeout:  DefaultGatewayTimeout,
		TerminalTimeout: DefaultTerminalTimeout,

		routeCacheTTL: DefaultRouteCacheTTL,
		gatewayHTTPClient: &http.Client{
			Transport: AddUserAgent(
				&http.Transport{},
				BuildUserAgent(),
			),
		}, // Timeout is set per request
		terminalHTTPClient: &http.Client{
			Timeout: DefaultTerminalTimeout,
			Transport: AddUserAgent(
				&http.Transport{
					TLSClientConfig: &tls.Config{
						RootCAs:    terminalCertPool(),
						ServerName: terminalCN,
					},
				},
				BuildUserAgent(),
			),
		},
	}
}

// ExpireRouteCache invalidates the route cache to for testing.
func (client *Client) ExpireRouteCache() {
	for key, value := range routeCache {
		value.TTL = time.Now()
		routeCache[key] = value
	}

	offlineCache := client.readOfflineCache()

	if offlineCache != nil {
		for _, route := range offlineCache.Routes {
			route.TTL = time.Now()
			client.updateOfflineCache(&route)
		}
	}
}

// Charge executes a standard direct preauth and capture.
func (client *Client) Charge(request AuthorizationRequest) (*AuthorizationResponse, error) {
	var response AuthorizationResponse
	var err error

	if request.TerminalName != "" {
		var route TerminalRoute
		route, err = client.resolveTerminalRoute(request.TerminalName)
		if err != nil {
			if errors.Is(err, ErrUnknownTerminal) {
				response.ResponseDescription = ResponseUnknownTerminal
				return &response, err
			}

			return nil, err
		}

		if route.CloudRelayEnabled {
			err = client.RelayRequest("/charge", "POST", request, &response, request.Test, request.Timeout)
		} else {
			authRequest := TerminalAuthorizationRequest{
				APICredentials: route.TransientCredentials,
				Request:        request,
			}
			err = client.terminalRequest(route, "/charge", "POST", authRequest, &response, request.Timeout)
		}
	} else {
		err = client.GatewayRequest("/charge", "POST", request, &response, request.Test, request.Timeout)
	}

	if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// Preauth executes a preauthorization intended to be captured later.
func (client *Client) Preauth(request AuthorizationRequest) (*AuthorizationResponse, error) {
	var response AuthorizationResponse
	var err error

	if request.TerminalName != "" {
		var route TerminalRoute
		route, err = client.resolveTerminalRoute(request.TerminalName)
		if err != nil {
			if errors.Is(err, ErrUnknownTerminal) {
				response.ResponseDescription = ResponseUnknownTerminal
				return &response, err
			}

			return nil, err
		}

		if route.CloudRelayEnabled {
			err = client.RelayRequest("/preauth", "POST", request, &response, request.Test, request.Timeout)
		} else {
			authRequest := TerminalAuthorizationRequest{
				APICredentials: route.TransientCredentials,
				Request:        request,
			}
			err = client.terminalRequest(route, "/preauth", "POST", authRequest, &response, request.Timeout)
		}
	} else {
		err = client.GatewayRequest("/preauth", "POST", request, &response, request.Test, request.Timeout)
	}

	if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// Ping tests connectivity with a payment terminal.
func (client *Client) Ping(request PingRequest) (*PingResponse, error) {
	var response PingResponse
	var err error

	if request.TerminalName != "" {
		var route TerminalRoute
		route, err = client.resolveTerminalRoute(request.TerminalName)
		if err != nil {
			if errors.Is(err, ErrUnknownTerminal) {
				response.ResponseDescription = ResponseUnknownTerminal
				return &response, err
			}

			return nil, err
		}

		if route.CloudRelayEnabled {
			err = client.RelayRequest("/terminal-test", "POST", request, &response, request.Test, request.Timeout)
		} else {
			authRequest := TerminalPingRequest{
				APICredentials: route.TransientCredentials,
				Request:        request,
			}
			err = client.terminalRequest(route, "/test", "POST", authRequest, &response, request.Timeout)
		}
	} else {
		err = client.GatewayRequest("/terminal-test", "POST", request, &response, request.Test, request.Timeout)
	}

	if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// Balance checks the remaining balance on a payment method.
func (client *Client) Balance(request BalanceRequest) (*BalanceResponse, error) {
	var response BalanceResponse
	var err error

	if request.TerminalName != "" {
		var route TerminalRoute
		route, err = client.resolveTerminalRoute(request.TerminalName)
		if err != nil {
			if errors.Is(err, ErrUnknownTerminal) {
				response.ResponseDescription = ResponseUnknownTerminal
				return &response, err
			}

			return nil, err
		}

		if route.CloudRelayEnabled {
			err = client.RelayRequest("/balance", "POST", request, &response, request.Test, request.Timeout)
		} else {
			authRequest := TerminalBalanceRequest{
				APICredentials: route.TransientCredentials,
				Request:        request,
			}
			err = client.terminalRequest(route, "/balance", "POST", authRequest, &response, request.Timeout)
		}
	} else {
		err = client.GatewayRequest("/balance", "POST", request, &response, request.Test, request.Timeout)
	}

	if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// Clear clears the line item display and any in progress transaction.
func (client *Client) Clear(request ClearTerminalRequest) (*Acknowledgement, error) {
	var response Acknowledgement
	var err error

	if request.TerminalName != "" {
		var route TerminalRoute
		route, err = client.resolveTerminalRoute(request.TerminalName)
		if err != nil {
			if errors.Is(err, ErrUnknownTerminal) {
				response.ResponseDescription = ResponseUnknownTerminal
				return &response, err
			}

			return nil, err
		}

		if route.CloudRelayEnabled {
			err = client.RelayRequest("/terminal-clear", "POST", request, &response, request.Test, request.Timeout)
		} else {
			authRequest := TerminalClearTerminalRequest{
				APICredentials: route.TransientCredentials,
				Request:        request,
			}
			err = client.terminalRequest(route, "/clear", "POST", authRequest, &response, request.Timeout)
		}
	} else {
		err = client.GatewayRequest("/terminal-clear", "POST", request, &response, request.Test, request.Timeout)
	}

	if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// TermsAndConditions prompts the user to accept terms and conditions.
func (client *Client) TermsAndConditions(request TermsAndConditionsRequest) (*TermsAndConditionsResponse, error) {
	var response TermsAndConditionsResponse
	var err error

	if request.TerminalName != "" {
		var route TerminalRoute
		route, err = client.resolveTerminalRoute(request.TerminalName)
		if err != nil {
			if errors.Is(err, ErrUnknownTerminal) {
				response.ResponseDescription = ResponseUnknownTerminal
				return &response, err
			}

			return nil, err
		}

		if route.CloudRelayEnabled {
			err = client.RelayRequest("/terminal-tc", "POST", request, &response, request.Test, request.Timeout)
		} else {
			authRequest := TerminalTermsAndConditionsRequest{
				APICredentials: route.TransientCredentials,
				Request:        request,
			}
			err = client.terminalRequest(route, "/tc", "POST", authRequest, &response, request.Timeout)
		}
	} else {
		err = client.GatewayRequest("/terminal-tc", "POST", request, &response, request.Test, request.Timeout)
	}

	if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// UpdateTransactionDisplay appends items to an existing transaction display
// Subtotal, Tax, and Total are overwritten by the request. Items with the
// same description are combined into groups.
func (client *Client) UpdateTransactionDisplay(request TransactionDisplayRequest) (*Acknowledgement, error) {
	var response Acknowledgement
	var err error

	if request.TerminalName != "" {
		var route TerminalRoute
		route, err = client.resolveTerminalRoute(request.TerminalName)
		if err != nil {
			if errors.Is(err, ErrUnknownTerminal) {
				response.ResponseDescription = ResponseUnknownTerminal
				return &response, err
			}

			return nil, err
		}

		if route.CloudRelayEnabled {
			err = client.RelayRequest("/terminal-txdisplay", "PUT", request, &response, request.Test, request.Timeout)
		} else {
			authRequest := TerminalTransactionDisplayRequest{
				APICredentials: route.TransientCredentials,
				Request:        request,
			}
			err = client.terminalRequest(route, "/txdisplay", "PUT", authRequest, &response, request.Timeout)
		}
	} else {
		err = client.GatewayRequest("/terminal-txdisplay", "PUT", request, &response, request.Test, request.Timeout)
	}

	if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// NewTransactionDisplay displays a new transaction on the terminal.
func (client *Client) NewTransactionDisplay(request TransactionDisplayRequest) (*Acknowledgement, error) {
	var response Acknowledgement
	var err error

	if request.TerminalName != "" {
		var route TerminalRoute
		route, err = client.resolveTerminalRoute(request.TerminalName)
		if err != nil {
			if errors.Is(err, ErrUnknownTerminal) {
				response.ResponseDescription = ResponseUnknownTerminal
				return &response, err
			}

			return nil, err
		}

		if route.CloudRelayEnabled {
			err = client.RelayRequest("/terminal-txdisplay", "POST", request, &response, request.Test, request.Timeout)
		} else {
			authRequest := TerminalTransactionDisplayRequest{
				APICredentials: route.TransientCredentials,
				Request:        request,
			}
			err = client.terminalRequest(route, "/txdisplay", "POST", authRequest, &response, request.Timeout)
		}
	} else {
		err = client.GatewayRequest("/terminal-txdisplay", "POST", request, &response, request.Test, request.Timeout)
	}

	if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// TextPrompt asks the consumer text based question.
func (client *Client) TextPrompt(request TextPromptRequest) (*TextPromptResponse, error) {
	var response TextPromptResponse
	var err error

	if request.TerminalName != "" {
		var route TerminalRoute
		route, err = client.resolveTerminalRoute(request.TerminalName)
		if err != nil {
			if errors.Is(err, ErrUnknownTerminal) {
				response.ResponseDescription = ResponseUnknownTerminal
				return &response, err
			}

			return nil, err
		}

		if route.CloudRelayEnabled {
			err = client.RelayRequest("/text-prompt", "POST", request, &response, request.Test, request.Timeout)
		} else {
			authRequest := TerminalTextPromptRequest{
				APICredentials: route.TransientCredentials,
				Request:        request,
			}
			err = client.terminalRequest(route, "/text-prompt", "POST", authRequest, &response, request.Timeout)
		}
	} else {
		err = client.GatewayRequest("/text-prompt", "POST", request, &response, request.Test, request.Timeout)
	}

	if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// BooleanPrompt asks the consumer a yes/no question.
func (client *Client) BooleanPrompt(request BooleanPromptRequest) (*BooleanPromptResponse, error) {
	var response BooleanPromptResponse
	var err error

	if request.TerminalName != "" {
		var route TerminalRoute
		route, err = client.resolveTerminalRoute(request.TerminalName)
		if err != nil {
			if errors.Is(err, ErrUnknownTerminal) {
				response.ResponseDescription = ResponseUnknownTerminal
				return &response, err
			}

			return nil, err
		}

		if route.CloudRelayEnabled {
			err = client.RelayRequest("/boolean-prompt", "POST", request, &response, request.Test, request.Timeout)
		} else {
			authRequest := TerminalBooleanPromptRequest{
				APICredentials: route.TransientCredentials,
				Request:        request,
			}
			err = client.terminalRequest(route, "/boolean-prompt", "POST", authRequest, &response, request.Timeout)
		}
	} else {
		err = client.GatewayRequest("/boolean-prompt", "POST", request, &response, request.Test, request.Timeout)
	}

	if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// Message displays a short message on the terminal.
func (client *Client) Message(request MessageRequest) (*Acknowledgement, error) {
	var response Acknowledgement
	var err error

	if request.TerminalName != "" {
		var route TerminalRoute
		route, err = client.resolveTerminalRoute(request.TerminalName)
		if err != nil {
			if errors.Is(err, ErrUnknownTerminal) {
				response.ResponseDescription = ResponseUnknownTerminal
				return &response, err
			}

			return nil, err
		}

		if route.CloudRelayEnabled {
			err = client.RelayRequest("/message", "POST", request, &response, request.Test, request.Timeout)
		} else {
			authRequest := TerminalMessageRequest{
				APICredentials: route.TransientCredentials,
				Request:        request,
			}
			err = client.terminalRequest(route, "/message", "POST", authRequest, &response, request.Timeout)
		}
	} else {
		err = client.GatewayRequest("/message", "POST", request, &response, request.Test, request.Timeout)
	}

	if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// Refund executes a refund.
func (client *Client) Refund(request RefundRequest) (*AuthorizationResponse, error) {
	var response AuthorizationResponse
	var err error

	if request.TerminalName != "" {
		var route TerminalRoute
		route, err = client.resolveTerminalRoute(request.TerminalName)
		if err != nil {
			if errors.Is(err, ErrUnknownTerminal) {
				response.ResponseDescription = ResponseUnknownTerminal
				return &response, err
			}

			return nil, err
		}

		if route.CloudRelayEnabled {
			err = client.RelayRequest("/refund", "POST", request, &response, request.Test, request.Timeout)
		} else {
			authRequest := TerminalRefundRequest{
				APICredentials: route.TransientCredentials,
				Request:        request,
			}
			err = client.terminalRequest(route, "/refund", "POST", authRequest, &response, request.Timeout)
		}
	} else {
		err = client.GatewayRequest("/refund", "POST", request, &response, request.Test, request.Timeout)
	}

	if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// Enroll adds a new payment method to the token vault.
func (client *Client) Enroll(request EnrollRequest) (*EnrollResponse, error) {
	var response EnrollResponse
	var err error

	if request.TerminalName != "" {
		var route TerminalRoute
		route, err = client.resolveTerminalRoute(request.TerminalName)
		if err != nil {
			if errors.Is(err, ErrUnknownTerminal) {
				response.ResponseDescription = ResponseUnknownTerminal
				return &response, err
			}

			return nil, err
		}

		if route.CloudRelayEnabled {
			err = client.RelayRequest("/enroll", "POST", request, &response, request.Test, request.Timeout)
		} else {
			authRequest := TerminalEnrollRequest{
				APICredentials: route.TransientCredentials,
				Request:        request,
			}
			err = client.terminalRequest(route, "/enroll", "POST", authRequest, &response, request.Timeout)
		}
	} else {
		err = client.GatewayRequest("/enroll", "POST", request, &response, request.Test, request.Timeout)
	}

	if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// GiftActivate activates or recharges a gift card.
func (client *Client) GiftActivate(request GiftActivateRequest) (*GiftActivateResponse, error) {
	var response GiftActivateResponse
	var err error

	if request.TerminalName != "" {
		var route TerminalRoute
		route, err = client.resolveTerminalRoute(request.TerminalName)
		if err != nil {
			if errors.Is(err, ErrUnknownTerminal) {
				response.ResponseDescription = ResponseUnknownTerminal
				return &response, err
			}

			return nil, err
		}

		if route.CloudRelayEnabled {
			err = client.RelayRequest("/gift-activate", "POST", request, &response, request.Test, request.Timeout)
		} else {
			authRequest := TerminalGiftActivateRequest{
				APICredentials: route.TransientCredentials,
				Request:        request,
			}
			err = client.terminalRequest(route, "/gift-activate", "POST", authRequest, &response, request.Timeout)
		}
	} else {
		err = client.GatewayRequest("/gift-activate", "POST", request, &response, request.Test, request.Timeout)
	}

	if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// TerminalStatus returns the current status of a terminal.
func (client *Client) TerminalStatus(request TerminalStatusRequest) (*TerminalStatusResponse, error) {
	var response TerminalStatusResponse
	var err error

	if request.TerminalName != "" {
		var route TerminalRoute
		route, err = client.resolveTerminalRoute(request.TerminalName)
		if err != nil {
			if errors.Is(err, ErrUnknownTerminal) {
				response.ResponseDescription = ResponseUnknownTerminal
				return &response, err
			}

			return nil, err
		}

		if route.CloudRelayEnabled {
			err = client.RelayRequest("/terminal-status", "POST", request, &response, request.Test, request.Timeout)
		} else {
			authRequest := TerminalTerminalStatusRequest{
				APICredentials: route.TransientCredentials,
				Request:        request,
			}
			err = client.terminalRequest(route, "/terminal-status", "POST", authRequest, &response, request.Timeout)
		}
	} else {
		err = client.GatewayRequest("/terminal-status", "POST", request, &response, request.Test, request.Timeout)
	}

	if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// Reverse executes a manual time out reversal.
//
// We love time out reversals. Don't be afraid to use them whenever a request
// to a BlockChyp terminal times out. You have up to two minutes to reverse
// any transaction. The only caveat is that you must assign transactionRef
// values when you build the original request. Otherwise, we have no real way
// of knowing which transaction you're trying to reverse because we may not
// have assigned it an id yet. And if we did assign it an id, you wouldn't
// know what it is because your request to the terminal timed out before you
// got a response.
func (client *Client) Reverse(request AuthorizationRequest) (*AuthorizationResponse, error) {
	var response AuthorizationResponse

	err := client.GatewayRequest("/reverse", "POST", request, &response, request.Test, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// Capture captures a preauthorization.
func (client *Client) Capture(request CaptureRequest) (*CaptureResponse, error) {
	var response CaptureResponse

	err := client.GatewayRequest("/capture", "POST", request, &response, request.Test, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// CloseBatch closes the current credit card batch.
func (client *Client) CloseBatch(request CloseBatchRequest) (*CloseBatchResponse, error) {
	var response CloseBatchResponse

	err := client.GatewayRequest("/close-batch", "POST", request, &response, request.Test, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// Void discards a previous preauth transaction.
func (client *Client) Void(request VoidRequest) (*VoidResponse, error) {
	var response VoidResponse

	err := client.GatewayRequest("/void", "POST", request, &response, request.Test, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

func getTimeout(requestTimeout interface{}, defaultTimeout time.Duration) time.Duration {
	var requestTimeoutDuration time.Duration
	switch v := requestTimeout.(type) {
	case int:
		requestTimeoutDuration = time.Duration(v) * time.Second
	case time.Duration:
		requestTimeoutDuration = v
	case nil:
	default:
		panic("must be int or time.Duration")
	}

	if requestTimeoutDuration <= 0 {
		return defaultTimeout
	}
	return requestTimeoutDuration
}
