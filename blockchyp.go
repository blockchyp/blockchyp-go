// Copyright 2019-2023 BlockChyp, Inc. All rights reserved. Use of this code
// is governed by a license that can be found in the LICENSE file.
//
// This file was generated automatically by the BlockChyp SDK Generator.
// Changes to this file will be lost every time the code is regenerated.

package blockchyp

import (
	"crypto/tls"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Default client configuration constants.
const (
	DefaultGatewayHost     = "https://api.blockchyp.com"
	DefaultDashboardHost   = "https://dashboard.blockchyp.com"
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
	DashboardHost   string
	TestGatewayHost string
	HTTPS           bool
	RouteCache      string

	GatewayTimeout  time.Duration
	TerminalTimeout time.Duration

	routeCacheTTL      time.Duration
	gatewayHTTPClient  *http.Client
	terminalHTTPClient *http.Client

	LogRequests bool
}

// NewClient returns a default Client configured with the given credentials.
func NewClient(creds APICredentials) Client {
	userAgent := BuildUserAgent()

	return Client{
		Credentials:     creds,
		GatewayHost:     DefaultGatewayHost,
		TestGatewayHost: DefaultTestGatewayHost,
		DashboardHost:   DefaultDashboardHost,
		HTTPS:           DefaultHTTPS,
		RouteCache:      filepath.Join(os.TempDir(), ".blockchyp_routes"),

		GatewayTimeout:  DefaultGatewayTimeout,
		TerminalTimeout: DefaultTerminalTimeout,

		routeCacheTTL: DefaultRouteCacheTTL,
		gatewayHTTPClient: &http.Client{
			Transport: AddUserAgent(
				&http.Transport{
					Dial: (&net.Dialer{
						Timeout: 5 * time.Second,
					}).Dial,
					TLSHandshakeTimeout: 5 * time.Second,
				},
				userAgent,
			),
		}, // Timeout is set per request
		terminalHTTPClient: &http.Client{
			Transport: AddUserAgent(
				&http.Transport{
					Dial: (&net.Dialer{
						Timeout: 5 * time.Second,
					}).Dial,
					TLSHandshakeTimeout: 5 * time.Second,
					TLSClientConfig: &tls.Config{
						RootCAs:    terminalCertPool(),
						ServerName: terminalCN,
					},
				},
				userAgent,
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

// Ping tests connectivity with a payment terminal.
func (client *Client) Ping(request PingRequest) (*PingResponse, error) {
	var response PingResponse
	var err error

	if err := populateSignatureOptions(&request); err != nil {
		return nil, err
	}

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
			err = client.RelayRequest("/api/terminal-test", "POST", request, &response, request.Test, request.Timeout)
		} else {
			authRequest := TerminalPingRequest{
				APICredentials: route.TransientCredentials,
				Request:        request,
			}
			err = client.terminalRequest(route, "/api/test", "POST", authRequest, &response, request.Timeout)
		}
	} else {
		err = client.GatewayRequest("/api/terminal-test", "POST", request, &response, request.Test, request.Timeout)
	}

	if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	if err := handleSignature(request, &response); err != nil {
		log.Printf("Failed to write signature: %+v", err)
	}

	return &response, err
}

// Charge executes a standard direct preauth and capture.
func (client *Client) Charge(request AuthorizationRequest) (*AuthorizationResponse, error) {
	var response AuthorizationResponse
	var err error

	if err := populateSignatureOptions(&request); err != nil {
		return nil, err
	}

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
			err = client.RelayRequest("/api/charge", "POST", request, &response, request.Test, request.Timeout)
		} else {
			authRequest := TerminalAuthorizationRequest{
				APICredentials: route.TransientCredentials,
				Request:        request,
			}
			err = client.terminalRequest(route, "/api/charge", "POST", authRequest, &response, request.Timeout)
		}
	} else {
		err = client.GatewayRequest("/api/charge", "POST", request, &response, request.Test, request.Timeout)
	}

	if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	if err := handleSignature(request, &response); err != nil {
		log.Printf("Failed to write signature: %+v", err)
	}

	return &response, err
}

// Preauth executes a preauthorization intended to be captured later.
func (client *Client) Preauth(request AuthorizationRequest) (*AuthorizationResponse, error) {
	var response AuthorizationResponse
	var err error

	if err := populateSignatureOptions(&request); err != nil {
		return nil, err
	}

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
			err = client.RelayRequest("/api/preauth", "POST", request, &response, request.Test, request.Timeout)
		} else {
			authRequest := TerminalAuthorizationRequest{
				APICredentials: route.TransientCredentials,
				Request:        request,
			}
			err = client.terminalRequest(route, "/api/preauth", "POST", authRequest, &response, request.Timeout)
		}
	} else {
		err = client.GatewayRequest("/api/preauth", "POST", request, &response, request.Test, request.Timeout)
	}

	if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	if err := handleSignature(request, &response); err != nil {
		log.Printf("Failed to write signature: %+v", err)
	}

	return &response, err
}

// Refund executes a refund.
func (client *Client) Refund(request RefundRequest) (*AuthorizationResponse, error) {
	var response AuthorizationResponse
	var err error

	if err := populateSignatureOptions(&request); err != nil {
		return nil, err
	}

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
			err = client.RelayRequest("/api/refund", "POST", request, &response, request.Test, request.Timeout)
		} else {
			authRequest := TerminalRefundRequest{
				APICredentials: route.TransientCredentials,
				Request:        request,
			}
			err = client.terminalRequest(route, "/api/refund", "POST", authRequest, &response, request.Timeout)
		}
	} else {
		err = client.GatewayRequest("/api/refund", "POST", request, &response, request.Test, request.Timeout)
	}

	if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	if err := handleSignature(request, &response); err != nil {
		log.Printf("Failed to write signature: %+v", err)
	}

	return &response, err
}

// Enroll adds a new payment method to the token vault.
func (client *Client) Enroll(request EnrollRequest) (*EnrollResponse, error) {
	var response EnrollResponse
	var err error

	if err := populateSignatureOptions(&request); err != nil {
		return nil, err
	}

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
			err = client.RelayRequest("/api/enroll", "POST", request, &response, request.Test, request.Timeout)
		} else {
			authRequest := TerminalEnrollRequest{
				APICredentials: route.TransientCredentials,
				Request:        request,
			}
			err = client.terminalRequest(route, "/api/enroll", "POST", authRequest, &response, request.Timeout)
		}
	} else {
		err = client.GatewayRequest("/api/enroll", "POST", request, &response, request.Test, request.Timeout)
	}

	if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	if err := handleSignature(request, &response); err != nil {
		log.Printf("Failed to write signature: %+v", err)
	}

	return &response, err
}

// GiftActivate activates or recharges a gift card.
func (client *Client) GiftActivate(request GiftActivateRequest) (*GiftActivateResponse, error) {
	var response GiftActivateResponse
	var err error

	if err := populateSignatureOptions(&request); err != nil {
		return nil, err
	}

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
			err = client.RelayRequest("/api/gift-activate", "POST", request, &response, request.Test, request.Timeout)
		} else {
			authRequest := TerminalGiftActivateRequest{
				APICredentials: route.TransientCredentials,
				Request:        request,
			}
			err = client.terminalRequest(route, "/api/gift-activate", "POST", authRequest, &response, request.Timeout)
		}
	} else {
		err = client.GatewayRequest("/api/gift-activate", "POST", request, &response, request.Test, request.Timeout)
	}

	if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	if err := handleSignature(request, &response); err != nil {
		log.Printf("Failed to write signature: %+v", err)
	}

	return &response, err
}

// Balance checks the remaining balance on a payment method.
func (client *Client) Balance(request BalanceRequest) (*BalanceResponse, error) {
	var response BalanceResponse
	var err error

	if err := populateSignatureOptions(&request); err != nil {
		return nil, err
	}

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
			err = client.RelayRequest("/api/balance", "POST", request, &response, request.Test, request.Timeout)
		} else {
			authRequest := TerminalBalanceRequest{
				APICredentials: route.TransientCredentials,
				Request:        request,
			}
			err = client.terminalRequest(route, "/api/balance", "POST", authRequest, &response, request.Timeout)
		}
	} else {
		err = client.GatewayRequest("/api/balance", "POST", request, &response, request.Test, request.Timeout)
	}

	if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	if err := handleSignature(request, &response); err != nil {
		log.Printf("Failed to write signature: %+v", err)
	}

	return &response, err
}

// Clear clears the line item display and any in progress transaction.
func (client *Client) Clear(request ClearTerminalRequest) (*Acknowledgement, error) {
	var response Acknowledgement
	var err error

	if err := populateSignatureOptions(&request); err != nil {
		return nil, err
	}

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
			err = client.RelayRequest("/api/terminal-clear", "POST", request, &response, request.Test, request.Timeout)
		} else {
			authRequest := TerminalClearTerminalRequest{
				APICredentials: route.TransientCredentials,
				Request:        request,
			}
			err = client.terminalRequest(route, "/api/clear", "POST", authRequest, &response, request.Timeout)
		}
	} else {
		err = client.GatewayRequest("/api/terminal-clear", "POST", request, &response, request.Test, request.Timeout)
	}

	if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	if err := handleSignature(request, &response); err != nil {
		log.Printf("Failed to write signature: %+v", err)
	}

	return &response, err
}

// TerminalStatus returns the current status of a terminal.
func (client *Client) TerminalStatus(request TerminalStatusRequest) (*TerminalStatusResponse, error) {
	var response TerminalStatusResponse
	var err error

	if err := populateSignatureOptions(&request); err != nil {
		return nil, err
	}

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
			err = client.RelayRequest("/api/terminal-status", "POST", request, &response, request.Test, request.Timeout)
		} else {
			authRequest := TerminalTerminalStatusRequest{
				APICredentials: route.TransientCredentials,
				Request:        request,
			}
			err = client.terminalRequest(route, "/api/terminal-status", "POST", authRequest, &response, request.Timeout)
		}
	} else {
		err = client.GatewayRequest("/api/terminal-status", "POST", request, &response, request.Test, request.Timeout)
	}

	if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	if err := handleSignature(request, &response); err != nil {
		log.Printf("Failed to write signature: %+v", err)
	}

	return &response, err
}

// TermsAndConditions prompts the user to accept terms and conditions.
func (client *Client) TermsAndConditions(request TermsAndConditionsRequest) (*TermsAndConditionsResponse, error) {
	var response TermsAndConditionsResponse
	var err error

	if err := populateSignatureOptions(&request); err != nil {
		return nil, err
	}

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
			err = client.RelayRequest("/api/terminal-tc", "POST", request, &response, request.Test, request.Timeout)
		} else {
			authRequest := TerminalTermsAndConditionsRequest{
				APICredentials: route.TransientCredentials,
				Request:        request,
			}
			err = client.terminalRequest(route, "/api/tc", "POST", authRequest, &response, request.Timeout)
		}
	} else {
		err = client.GatewayRequest("/api/terminal-tc", "POST", request, &response, request.Test, request.Timeout)
	}

	if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	if err := handleSignature(request, &response); err != nil {
		log.Printf("Failed to write signature: %+v", err)
	}

	return &response, err
}

// CaptureSignature captures and returns a signature.
func (client *Client) CaptureSignature(request CaptureSignatureRequest) (*CaptureSignatureResponse, error) {
	var response CaptureSignatureResponse
	var err error

	if err := populateSignatureOptions(&request); err != nil {
		return nil, err
	}

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
			err = client.RelayRequest("/api/capture-signature", "POST", request, &response, request.Test, request.Timeout)
		} else {
			authRequest := TerminalCaptureSignatureRequest{
				APICredentials: route.TransientCredentials,
				Request:        request,
			}
			err = client.terminalRequest(route, "/api/capture-signature", "POST", authRequest, &response, request.Timeout)
		}
	} else {
		err = client.GatewayRequest("/api/capture-signature", "POST", request, &response, request.Test, request.Timeout)
	}

	if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	if err := handleSignature(request, &response); err != nil {
		log.Printf("Failed to write signature: %+v", err)
	}

	return &response, err
}

// NewTransactionDisplay displays a new transaction on the terminal.
func (client *Client) NewTransactionDisplay(request TransactionDisplayRequest) (*Acknowledgement, error) {
	var response Acknowledgement
	var err error

	if err := populateSignatureOptions(&request); err != nil {
		return nil, err
	}

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
			err = client.RelayRequest("/api/terminal-txdisplay", "POST", request, &response, request.Test, request.Timeout)
		} else {
			authRequest := TerminalTransactionDisplayRequest{
				APICredentials: route.TransientCredentials,
				Request:        request,
			}
			err = client.terminalRequest(route, "/api/txdisplay", "POST", authRequest, &response, request.Timeout)
		}
	} else {
		err = client.GatewayRequest("/api/terminal-txdisplay", "POST", request, &response, request.Test, request.Timeout)
	}

	if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	if err := handleSignature(request, &response); err != nil {
		log.Printf("Failed to write signature: %+v", err)
	}

	return &response, err
}

// UpdateTransactionDisplay appends items to an existing transaction display.
// Subtotal, Tax, and Total are overwritten by the request. Items with the
// same description are combined into groups.
func (client *Client) UpdateTransactionDisplay(request TransactionDisplayRequest) (*Acknowledgement, error) {
	var response Acknowledgement
	var err error

	if err := populateSignatureOptions(&request); err != nil {
		return nil, err
	}

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
			err = client.RelayRequest("/api/terminal-txdisplay", "PUT", request, &response, request.Test, request.Timeout)
		} else {
			authRequest := TerminalTransactionDisplayRequest{
				APICredentials: route.TransientCredentials,
				Request:        request,
			}
			err = client.terminalRequest(route, "/api/txdisplay", "PUT", authRequest, &response, request.Timeout)
		}
	} else {
		err = client.GatewayRequest("/api/terminal-txdisplay", "PUT", request, &response, request.Test, request.Timeout)
	}

	if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	if err := handleSignature(request, &response); err != nil {
		log.Printf("Failed to write signature: %+v", err)
	}

	return &response, err
}

// Message displays a short message on the terminal.
func (client *Client) Message(request MessageRequest) (*Acknowledgement, error) {
	var response Acknowledgement
	var err error

	if err := populateSignatureOptions(&request); err != nil {
		return nil, err
	}

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
			err = client.RelayRequest("/api/message", "POST", request, &response, request.Test, request.Timeout)
		} else {
			authRequest := TerminalMessageRequest{
				APICredentials: route.TransientCredentials,
				Request:        request,
			}
			err = client.terminalRequest(route, "/api/message", "POST", authRequest, &response, request.Timeout)
		}
	} else {
		err = client.GatewayRequest("/api/message", "POST", request, &response, request.Test, request.Timeout)
	}

	if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	if err := handleSignature(request, &response); err != nil {
		log.Printf("Failed to write signature: %+v", err)
	}

	return &response, err
}

// BooleanPrompt asks the consumer a yes/no question.
func (client *Client) BooleanPrompt(request BooleanPromptRequest) (*BooleanPromptResponse, error) {
	var response BooleanPromptResponse
	var err error

	if err := populateSignatureOptions(&request); err != nil {
		return nil, err
	}

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
			err = client.RelayRequest("/api/boolean-prompt", "POST", request, &response, request.Test, request.Timeout)
		} else {
			authRequest := TerminalBooleanPromptRequest{
				APICredentials: route.TransientCredentials,
				Request:        request,
			}
			err = client.terminalRequest(route, "/api/boolean-prompt", "POST", authRequest, &response, request.Timeout)
		}
	} else {
		err = client.GatewayRequest("/api/boolean-prompt", "POST", request, &response, request.Test, request.Timeout)
	}

	if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	if err := handleSignature(request, &response); err != nil {
		log.Printf("Failed to write signature: %+v", err)
	}

	return &response, err
}

// TextPrompt asks the consumer a text based question.
func (client *Client) TextPrompt(request TextPromptRequest) (*TextPromptResponse, error) {
	var response TextPromptResponse
	var err error

	if err := populateSignatureOptions(&request); err != nil {
		return nil, err
	}

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
			err = client.RelayRequest("/api/text-prompt", "POST", request, &response, request.Test, request.Timeout)
		} else {
			authRequest := TerminalTextPromptRequest{
				APICredentials: route.TransientCredentials,
				Request:        request,
			}
			err = client.terminalRequest(route, "/api/text-prompt", "POST", authRequest, &response, request.Timeout)
		}
	} else {
		err = client.GatewayRequest("/api/text-prompt", "POST", request, &response, request.Test, request.Timeout)
	}

	if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	if err := handleSignature(request, &response); err != nil {
		log.Printf("Failed to write signature: %+v", err)
	}

	return &response, err
}

// ListQueuedTransactions returns a list of queued transactions on a terminal.
func (client *Client) ListQueuedTransactions(request ListQueuedTransactionsRequest) (*ListQueuedTransactionsResponse, error) {
	var response ListQueuedTransactionsResponse
	var err error

	if err := populateSignatureOptions(&request); err != nil {
		return nil, err
	}

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
			err = client.RelayRequest("/api/queue/list", "POST", request, &response, request.Test, request.Timeout)
		} else {
			authRequest := TerminalListQueuedTransactionsRequest{
				APICredentials: route.TransientCredentials,
				Request:        request,
			}
			err = client.terminalRequest(route, "/api/queue/list", "POST", authRequest, &response, request.Timeout)
		}
	} else {
		err = client.GatewayRequest("/api/queue/list", "POST", request, &response, request.Test, request.Timeout)
	}

	if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	if err := handleSignature(request, &response); err != nil {
		log.Printf("Failed to write signature: %+v", err)
	}

	return &response, err
}

// DeleteQueuedTransaction deletes a queued transaction from the terminal.
func (client *Client) DeleteQueuedTransaction(request DeleteQueuedTransactionRequest) (*DeleteQueuedTransactionResponse, error) {
	var response DeleteQueuedTransactionResponse
	var err error

	if err := populateSignatureOptions(&request); err != nil {
		return nil, err
	}

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
			err = client.RelayRequest("/api/queue/delete", "POST", request, &response, request.Test, request.Timeout)
		} else {
			authRequest := TerminalDeleteQueuedTransactionRequest{
				APICredentials: route.TransientCredentials,
				Request:        request,
			}
			err = client.terminalRequest(route, "/api/queue/delete", "POST", authRequest, &response, request.Timeout)
		}
	} else {
		err = client.GatewayRequest("/api/queue/delete", "POST", request, &response, request.Test, request.Timeout)
	}

	if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	if err := handleSignature(request, &response); err != nil {
		log.Printf("Failed to write signature: %+v", err)
	}

	return &response, err
}

// Reboot reboot a payment terminal.
func (client *Client) Reboot(request PingRequest) (*Acknowledgement, error) {
	var response Acknowledgement
	var err error

	if err := populateSignatureOptions(&request); err != nil {
		return nil, err
	}

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
			err = client.RelayRequest("/api/terminal-reboot", "POST", request, &response, request.Test, request.Timeout)
		} else {
			authRequest := TerminalPingRequest{
				APICredentials: route.TransientCredentials,
				Request:        request,
			}
			err = client.terminalRequest(route, "/api/reboot", "POST", authRequest, &response, request.Timeout)
		}
	} else {
		err = client.GatewayRequest("/api/terminal-reboot", "POST", request, &response, request.Test, request.Timeout)
	}

	if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	if err := handleSignature(request, &response); err != nil {
		log.Printf("Failed to write signature: %+v", err)
	}

	return &response, err
}

// Locate returns routing and location data for a payment terminal.
func (client *Client) Locate(request LocateRequest) (*LocateResponse, error) {
	var response LocateResponse

	err := client.GatewayRequest("/api/terminal-locate", "POST", request, &response, request.Test, request.Timeout)

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

	err := client.GatewayRequest("/api/capture", "POST", request, &response, request.Test, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// Void discards a previous transaction.
func (client *Client) Void(request VoidRequest) (*VoidResponse, error) {
	var response VoidResponse

	err := client.GatewayRequest("/api/void", "POST", request, &response, request.Test, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
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

	err := client.GatewayRequest("/api/reverse", "POST", request, &response, request.Test, request.Timeout)

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

	err := client.GatewayRequest("/api/close-batch", "POST", request, &response, request.Test, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// SendPaymentLink creates and send a payment link to a customer.
func (client *Client) SendPaymentLink(request PaymentLinkRequest) (*PaymentLinkResponse, error) {
	var response PaymentLinkResponse

	err := client.GatewayRequest("/api/send-payment-link", "POST", request, &response, request.Test, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// CancelPaymentLink cancels a payment link.
func (client *Client) CancelPaymentLink(request CancelPaymentLinkRequest) (*CancelPaymentLinkResponse, error) {
	var response CancelPaymentLinkResponse

	err := client.GatewayRequest("/api/cancel-payment-link", "POST", request, &response, request.Test, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// TransactionStatus retrieves the current status of a transaction.
func (client *Client) TransactionStatus(request TransactionStatusRequest) (*AuthorizationResponse, error) {
	var response AuthorizationResponse

	err := client.GatewayRequest("/api/tx-status", "POST", request, &response, request.Test, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// UpdateCustomer updates or creates a customer record.
func (client *Client) UpdateCustomer(request UpdateCustomerRequest) (*CustomerResponse, error) {
	var response CustomerResponse

	err := client.GatewayRequest("/api/update-customer", "POST", request, &response, request.Test, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// Customer retrieves a customer by id.
func (client *Client) Customer(request CustomerRequest) (*CustomerResponse, error) {
	var response CustomerResponse

	err := client.GatewayRequest("/api/customer", "POST", request, &response, request.Test, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// CustomerSearch searches the customer database.
func (client *Client) CustomerSearch(request CustomerSearchRequest) (*CustomerSearchResponse, error) {
	var response CustomerSearchResponse

	err := client.GatewayRequest("/api/customer-search", "POST", request, &response, request.Test, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// CashDiscount calculates the discount for actual cash transactions.
func (client *Client) CashDiscount(request CashDiscountRequest) (*CashDiscountResponse, error) {
	var response CashDiscountResponse

	err := client.GatewayRequest("/api/cash-discount", "POST", request, &response, request.Test, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// BatchHistory returns the batch history for a merchant.
func (client *Client) BatchHistory(request BatchHistoryRequest) (*BatchHistoryResponse, error) {
	var response BatchHistoryResponse

	err := client.GatewayRequest("/api/batch-history", "POST", request, &response, request.Test, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// BatchDetails returns the batch details for a single batch.
func (client *Client) BatchDetails(request BatchDetailsRequest) (*BatchDetailsResponse, error) {
	var response BatchDetailsResponse

	err := client.GatewayRequest("/api/batch-details", "POST", request, &response, request.Test, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// TransactionHistory returns the transaction history for a merchant.
func (client *Client) TransactionHistory(request TransactionHistoryRequest) (*TransactionHistoryResponse, error) {
	var response TransactionHistoryResponse

	err := client.GatewayRequest("/api/tx-history", "POST", request, &response, request.Test, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// MerchantProfile returns profile information for a merchant.
func (client *Client) MerchantProfile(request MerchantProfileRequest) (*MerchantProfileResponse, error) {
	var response MerchantProfileResponse

	err := client.GatewayRequest("/api/public-merchant-profile", "POST", request, &response, request.Test, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// DeleteCustomer deletes a customer record.
func (client *Client) DeleteCustomer(request DeleteCustomerRequest) (*DeleteCustomerResponse, error) {
	var response DeleteCustomerResponse

	err := client.GatewayRequest("/api/customer/"+request.CustomerID, "DELETE", request, &response, request.Test, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// TokenMetadata retrieves payment token metadata.
func (client *Client) TokenMetadata(request TokenMetadataRequest) (*TokenMetadataResponse, error) {
	var response TokenMetadataResponse

	err := client.GatewayRequest("/api/token/"+request.Token, "GET", request, &response, request.Test, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// LinkToken links a token to a customer record.
func (client *Client) LinkToken(request LinkTokenRequest) (*Acknowledgement, error) {
	var response Acknowledgement

	err := client.GatewayRequest("/api/link-token", "POST", request, &response, request.Test, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// UnlinkToken removes a link between a customer and a token.
func (client *Client) UnlinkToken(request UnlinkTokenRequest) (*Acknowledgement, error) {
	var response Acknowledgement

	err := client.GatewayRequest("/api/unlink-token", "POST", request, &response, request.Test, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// DeleteToken deletes a payment token.
func (client *Client) DeleteToken(request DeleteTokenRequest) (*DeleteTokenResponse, error) {
	var response DeleteTokenResponse

	err := client.GatewayRequest("/api/token/"+request.Token, "DELETE", request, &response, request.Test, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// GetMerchants adds a test merchant account.
func (client *Client) GetMerchants(request GetMerchantsRequest) (*GetMerchantsResponse, error) {
	var response GetMerchantsResponse

	err := client.DashboardRequest("/api/get-merchants", "POST", request, &response, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// UpdateMerchant adds or updates a merchant account. Can be used to create or
// update test merchants. Only gateway partners may create new live merchants.
func (client *Client) UpdateMerchant(request MerchantProfile) (*MerchantProfileResponse, error) {
	var response MerchantProfileResponse

	err := client.DashboardRequest("/api/update-merchant", "POST", request, &response, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// MerchantUsers list all active users and pending invites for a merchant
// account.
func (client *Client) MerchantUsers(request MerchantProfileRequest) (*MerchantUsersResponse, error) {
	var response MerchantUsersResponse

	err := client.DashboardRequest("/api/merchant-users", "POST", request, &response, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// InviteMerchantUser invites a user to join a merchant account.
func (client *Client) InviteMerchantUser(request InviteMerchantUserRequest) (*Acknowledgement, error) {
	var response Acknowledgement

	err := client.DashboardRequest("/api/invite-merchant-user", "POST", request, &response, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// AddTestMerchant adds a test merchant account.
func (client *Client) AddTestMerchant(request AddTestMerchantRequest) (*MerchantProfileResponse, error) {
	var response MerchantProfileResponse

	err := client.DashboardRequest("/api/add-test-merchant", "POST", request, &response, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// DeleteTestMerchant deletes a test merchant account. Supports partner scoped
// API credentials only. Live merchant accounts cannot be deleted.
func (client *Client) DeleteTestMerchant(request MerchantProfileRequest) (*Acknowledgement, error) {
	var response Acknowledgement

	err := client.DashboardRequest("/api/test-merchant/"+request.MerchantID, "DELETE", request, &response, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// MerchantPlatforms list all merchant platforms configured for a gateway
// merchant.
func (client *Client) MerchantPlatforms(request MerchantProfileRequest) (*MerchantPlatformsResponse, error) {
	var response MerchantPlatformsResponse

	err := client.DashboardRequest("/api/plugin-configs/"+request.MerchantID, "GET", request, &response, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// UpdateMerchantPlatforms list all merchant platforms configured for a
// gateway merchant.
func (client *Client) UpdateMerchantPlatforms(request MerchantPlatform) (*Acknowledgement, error) {
	var response Acknowledgement

	err := client.DashboardRequest("/api/plugin-configs", "POST", request, &response, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// DeleteMerchantPlatforms deletes a boarding platform configuration.
func (client *Client) DeleteMerchantPlatforms(request MerchantPlatformRequest) (*Acknowledgement, error) {
	var response Acknowledgement

	err := client.DashboardRequest("/api/plugin-config/"+request.PlatformID, "DELETE", request, &response, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// Terminals returns all terminals associated with the merchant account.
func (client *Client) Terminals(request TerminalProfileRequest) (*TerminalProfileResponse, error) {
	var response TerminalProfileResponse

	err := client.DashboardRequest("/api/terminals", "GET", request, &response, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// DeactivateTerminal deactivates a terminal.
func (client *Client) DeactivateTerminal(request TerminalDeactivationRequest) (*Acknowledgement, error) {
	var response Acknowledgement

	err := client.DashboardRequest("/api/terminal/"+request.TerminalID, "DELETE", request, &response, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// ActivateTerminal activates a terminal.
func (client *Client) ActivateTerminal(request TerminalActivationRequest) (*Acknowledgement, error) {
	var response Acknowledgement

	err := client.DashboardRequest("/api/terminal-activate", "POST", request, &response, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// TCTemplates returns a list of terms and conditions templates associated
// with a merchant account.
func (client *Client) TCTemplates(request TermsAndConditionsTemplateRequest) (*TermsAndConditionsTemplateResponse, error) {
	var response TermsAndConditionsTemplateResponse

	err := client.DashboardRequest("/api/tc-templates", "GET", request, &response, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// TCTemplate returns a single terms and conditions template.
func (client *Client) TCTemplate(request TermsAndConditionsTemplateRequest) (*TermsAndConditionsTemplate, error) {
	var response TermsAndConditionsTemplate

	err := client.DashboardRequest("/api/tc-templates/"+request.TemplateID, "GET", request, &response, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// TCUpdateTemplate updates or creates a terms and conditions template.
func (client *Client) TCUpdateTemplate(request TermsAndConditionsTemplate) (*TermsAndConditionsTemplate, error) {
	var response TermsAndConditionsTemplate

	err := client.DashboardRequest("/api/tc-templates", "POST", request, &response, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// TCDeleteTemplate deletes a single terms and conditions template.
func (client *Client) TCDeleteTemplate(request TermsAndConditionsTemplateRequest) (*Acknowledgement, error) {
	var response Acknowledgement

	err := client.DashboardRequest("/api/tc-templates/"+request.TemplateID, "DELETE", request, &response, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// TCLog returns up to 250 entries from the Terms and Conditions log.
func (client *Client) TCLog(request TermsAndConditionsLogRequest) (*TermsAndConditionsLogResponse, error) {
	var response TermsAndConditionsLogResponse

	err := client.DashboardRequest("/api/tc-log", "POST", request, &response, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// TCEntry returns a single detailed Terms and Conditions entry.
func (client *Client) TCEntry(request TermsAndConditionsLogRequest) (*TermsAndConditionsLogEntry, error) {
	var response TermsAndConditionsLogEntry

	err := client.DashboardRequest("/api/tc-entry/"+request.LogEntryID, "GET", request, &response, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// SurveyQuestions returns all survey questions for a given merchant.
func (client *Client) SurveyQuestions(request SurveyQuestionRequest) (*SurveyQuestionResponse, error) {
	var response SurveyQuestionResponse

	err := client.DashboardRequest("/api/survey-questions", "GET", request, &response, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// SurveyQuestion returns a single survey question with response data.
func (client *Client) SurveyQuestion(request SurveyQuestionRequest) (*SurveyQuestion, error) {
	var response SurveyQuestion

	err := client.DashboardRequest("/api/survey-questions/"+request.QuestionID, "GET", request, &response, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// UpdateSurveyQuestion updates or creates a survey question.
func (client *Client) UpdateSurveyQuestion(request SurveyQuestion) (*SurveyQuestion, error) {
	var response SurveyQuestion

	err := client.DashboardRequest("/api/survey-questions", "POST", request, &response, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// DeleteSurveyQuestion deletes a survey question.
func (client *Client) DeleteSurveyQuestion(request SurveyQuestionRequest) (*Acknowledgement, error) {
	var response Acknowledgement

	err := client.DashboardRequest("/api/survey-questions/"+request.QuestionID, "DELETE", request, &response, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// SurveyResults returns results for a single survey question.
func (client *Client) SurveyResults(request SurveyResultsRequest) (*SurveyQuestion, error) {
	var response SurveyQuestion

	err := client.DashboardRequest("/api/survey-results", "POST", request, &response, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// Media returns the media library for a given partner, merchant, or
// organization.
func (client *Client) Media(request MediaRequest) (*MediaLibraryResponse, error) {
	var response MediaLibraryResponse

	err := client.DashboardRequest("/api/media", "GET", request, &response, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// UploadMedia uploads a media asset to the media library.
func (client *Client) UploadMedia(request UploadMetadata, reader io.Reader) (*MediaMetadata, error) {

	var response MediaMetadata

	err := client.DashboardUpload("/api/upload-media", request, reader, &response, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err

}

// UploadStatus retrieves the current status of a file upload.
func (client *Client) UploadStatus(request UploadStatusRequest) (*UploadStatus, error) {
	var response UploadStatus

	err := client.DashboardRequest("/api/media-upload/"+request.UploadID, "GET", request, &response, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// MediaAsset returns the media details for a single media asset.
func (client *Client) MediaAsset(request MediaRequest) (*MediaMetadata, error) {
	var response MediaMetadata

	err := client.DashboardRequest("/api/media/"+request.MediaID, "GET", request, &response, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// DeleteMediaAsset deletes a media asset.
func (client *Client) DeleteMediaAsset(request MediaRequest) (*Acknowledgement, error) {
	var response Acknowledgement

	err := client.DashboardRequest("/api/media/"+request.MediaID, "DELETE", request, &response, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// SlideShows returns a collection of slide shows.
func (client *Client) SlideShows(request SlideShowRequest) (*SlideShowResponse, error) {
	var response SlideShowResponse

	err := client.DashboardRequest("/api/slide-shows", "GET", request, &response, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// SlideShow returns a single slide show with slides.
func (client *Client) SlideShow(request SlideShowRequest) (*SlideShow, error) {
	var response SlideShow

	err := client.DashboardRequest("/api/slide-shows/"+request.SlideShowID, "GET", request, &response, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// UpdateSlideShow updates or creates a slide show.
func (client *Client) UpdateSlideShow(request SlideShow) (*SlideShow, error) {
	var response SlideShow

	err := client.DashboardRequest("/api/slide-shows", "POST", request, &response, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// DeleteSlideShow deletes a single slide show.
func (client *Client) DeleteSlideShow(request SlideShowRequest) (*Acknowledgement, error) {
	var response Acknowledgement

	err := client.DashboardRequest("/api/slide-shows/"+request.SlideShowID, "DELETE", request, &response, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// TerminalBranding returns the terminal branding stack for a given set of API
// credentials.
func (client *Client) TerminalBranding(request BrandingAssetRequest) (*BrandingAssetResponse, error) {
	var response BrandingAssetResponse

	err := client.DashboardRequest("/api/terminal-branding", "GET", request, &response, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// UpdateBrandingAsset updates a branding asset.
func (client *Client) UpdateBrandingAsset(request BrandingAsset) (*BrandingAsset, error) {
	var response BrandingAsset

	err := client.DashboardRequest("/api/terminal-branding", "POST", request, &response, request.Timeout)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

// DeleteBrandingAsset deletes a branding asset.
func (client *Client) DeleteBrandingAsset(request BrandingAssetRequest) (*Acknowledgement, error) {
	var response Acknowledgement

	err := client.DashboardRequest("/api/terminal-branding/"+request.AssetID, "DELETE", request, &response, request.Timeout)

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

func populateSignatureOptions(request interface{}) error {
	sigOpts, ok := (SignatureRequest{}).From(request)
	if !ok {
		return nil
	}

	if sigOpts.SigFile != "" && sigOpts.SigFormat == "" {
		x := strings.Split(sigOpts.SigFile, ".")
		sigOpts.SigFormat = SignatureFormat(strings.ToLower(x[len(x)-1]))
	}

	switch sigOpts.SigFormat {
	case SignatureFormatNone, SignatureFormatPNG, SignatureFormatJPG, SignatureFormatGIF:
	default:
		return fmt.Errorf("invalid signature format: %s", sigOpts.SigFormat)
	}

	copyTo(sigOpts, request)

	return nil
}

func handleSignature(request, response interface{}) error {
	requestOpts, ok := (SignatureRequest{}).From(request)
	if !ok {
		return nil
	}

	responseOpts, ok := (SignatureResponse{}).From(response)
	if !ok {
		return nil
	}

	if requestOpts.SigFile == "" || responseOpts.SigFile == "" {
		return nil
	}

	clearField(response, "SigFile")

	content, err := hex.DecodeString(responseOpts.SigFile)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(requestOpts.SigFile, content, 0600)
}
