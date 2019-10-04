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

/*
Default client configuration constants.
*/
const (
	DefaultGatewayHost     = "https://api.blockchyp.com"
	DefaultTestGatewayHost = "https://test.blockchyp.com"
	DefaultHTTPS           = true
	DefaultRouteCacheTTL   = 60 * time.Minute
	DefaultGatewayTimeout  = 20 * time.Second
	DefaultTerminalTimeout = 2 * time.Minute
)

/*
Default filesystem configuration.
*/
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

/*
Client is the main interface used by application developers.
*/
type Client struct {
	Credentials        APICredentials
	GatewayHost        string
	TestGatewayHost    string
	HTTPS              bool
	RouteCache         string
	routeCacheTTL      time.Duration
	gatewayHTTPClient  *http.Client
	terminalHTTPClient *http.Client
}

/*
NewClient returns a default Client configured with the given credentials.
*/
func NewClient(creds APICredentials) Client {
	return Client{
		Credentials:   creds,
		GatewayHost:   DefaultGatewayHost,
		HTTPS:         DefaultHTTPS,
		routeCacheTTL: DefaultRouteCacheTTL,
		RouteCache:    filepath.Join(os.TempDir(), ".blockchyp_routes"),
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

/*
AsyncCharge executes an asynchronous auth and capture.
*/
func (client *Client) AsyncCharge(request AuthorizationRequest, responseChan chan<- AuthorizationResponse) error {

	if !isValidAsyncMethod(request.PaymentMethod) {
		return ErrInvalidAsyncRequest
	}

	return nil
}

/*
TextPrompt asks the consumer text based question.
*/
func (client *Client) TextPrompt(request TextPromptRequest) (*TextPromptResponse, error) {

	var response TextPromptResponse

	route, err := client.resolveTerminalRoute(request.TerminalName)
	if err != nil {
		if err == ErrUnknownTerminal {
			response.Error = ResponseUnknownTerminal
			return &response, err
		}

		return nil, err
	}

	if route.CloudRelayEnabled {
		err = client.RelayRequest("/text-prompt", http.MethodPost, request, &response, request.Test)
	} else {
		terminalRequest := TerminalTextPromptRequest{
			APICredentials: route.TransientCredentials,
			Request:        request,
		}
		err = client.terminalPost(route, "/text-prompt", terminalRequest, &response)
	}

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.Error = ResponseTimedOut
	} else if err != nil {
		response.Error = err.Error()
	}

	return &response, err

}

/*
TC prompts the user to accept terms and conditions.
*/
func (client *Client) TC(request TermsAndConditionsRequest) (*TermsAndConditionsResponse, error) {

	var response TermsAndConditionsResponse

	route, err := client.resolveTerminalRoute(request.TerminalName)
	if err != nil {
		if err == ErrUnknownTerminal {
			response.Error = ResponseUnknownTerminal
			return &response, err
		}

		return nil, err
	}

	if route.CloudRelayEnabled {
		err = client.RelayRequest("/tc", http.MethodPost, request, &response, request.Test)
	} else {
		terminalRequest := TerminalTermsAndConditionsRequest{
			APICredentials: route.TransientCredentials,
			Request:        request,
		}
		err = client.terminalPost(route, "/tc", terminalRequest, &response)
	}

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.Error = ResponseTimedOut
	} else if err != nil {
		response.Error = err.Error()
	}

	return &response, err

}

/*
BooleanPrompt asks the consumer a yes/no question.
*/
func (client *Client) BooleanPrompt(request BooleanPromptRequest) (*BooleanPromptResponse, error) {

	var response BooleanPromptResponse

	route, err := client.resolveTerminalRoute(request.TerminalName)
	if err != nil {
		if err == ErrUnknownTerminal {
			response.Error = ResponseUnknownTerminal
			return &response, err
		}

		return nil, err
	}

	if route.CloudRelayEnabled {
		err = client.RelayRequest("/boolean-prompt", http.MethodPost, request, &response, request.Test)
	} else {
		terminalRequest := TerminalBooleanPromptRequest{
			APICredentials: route.TransientCredentials,
			Request:        request,
		}
		err = client.terminalPost(route, "/boolean-prompt", terminalRequest, &response)
	}

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.Error = ResponseTimedOut
	} else if err != nil {
		response.Error = err.Error()
	}

	return &response, err

}

/*
Message displays a short message on the terminal.
*/
func (client *Client) Message(request MessageRequest) (*Acknowledgement, error) {

	var response Acknowledgement

	route, err := client.resolveTerminalRoute(request.TerminalName)
	if err != nil {
		if err == ErrUnknownTerminal {
			response.Error = ResponseUnknownTerminal
			return &response, err
		}

		return nil, err
	}

	if route.CloudRelayEnabled {
		err = client.RelayRequest("/message", http.MethodPost, request, &response, request.Test)
	} else {
		terminalRequest := TerminalMessageRequest{
			APICredentials: route.TransientCredentials,
			Request:        request,
		}
		err = client.terminalPost(route, "/message", terminalRequest, &response)
	}

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.Error = ResponseTimedOut
	} else if err != nil {
		response.Error = err.Error()
	}

	return &response, err

}

/*
Charge executes a standard direct preauth and capture.
*/
func (client *Client) Charge(request AuthorizationRequest) (*AuthorizationResponse, error) {

	var response AuthorizationResponse
	var err error

	if request.IsTerminalRouted() {
		var route TerminalRoute
		route, err = client.resolveTerminalRoute(request.TerminalName)
		if err != nil {
			if err == ErrUnknownTerminal {
				response.ResponseDescription = ResponseUnknownTerminal
				return &response, err
			}

			return nil, err
		}

		if route.CloudRelayEnabled {
			err = client.RelayRequest("/charge", http.MethodPost, request, &response, request.Test)
		} else {
			authRequest := TerminalAuthorizationRequest{
				APICredentials: route.TransientCredentials,
				Request:        request,
			}
			err = client.terminalPost(route, "/charge", authRequest, &response)
		}
	} else {
		err = client.GatewayRequest("/charge", http.MethodPost, request, &response, request.Test)
	}

	if timeout, ok := err.(net.Error); ok && timeout.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err

}

/*
AsyncPreauth executes an asynchronous preauthorization.
*/
func (client *Client) AsyncPreauth(request AuthorizationRequest, responseChan chan<- AuthorizationResponse) error {

	if !isValidAsyncMethod(request.PaymentMethod) {
		return ErrInvalidAsyncRequest
	}

	return nil
}

/*
Preauth executes a preauthorization intended to be captured later.
*/
func (client *Client) Preauth(request AuthorizationRequest) (*AuthorizationResponse, error) {

	var response AuthorizationResponse
	var err error

	if request.IsTerminalRouted() {
		var route TerminalRoute
		route, err = client.resolveTerminalRoute(request.TerminalName)
		if err != nil {
			if err == ErrUnknownTerminal {
				response.ResponseDescription = ResponseUnknownTerminal
				return &response, err
			}

			return nil, err
		}

		if route.CloudRelayEnabled {
			err = client.RelayRequest("/preauth", http.MethodPost, request, &response, request.Test)
		} else {
			terminalRequest := TerminalAuthorizationRequest{
				APICredentials: route.TransientCredentials,
				Request:        request,
			}
			err = client.terminalPost(route, "/preauth", terminalRequest, &response)
		}
	} else {
		err = client.GatewayRequest("/preauth", http.MethodPost, request, &response, request.Test)
	}

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err

}

/*
AsyncRefund executes an asynchronous refund
*/
func (client *Client) AsyncRefund(request AuthorizationRequest, responseChan chan<- AuthorizationResponse) error {

	if !isValidAsyncMethod(request.PaymentMethod) {
		return ErrInvalidAsyncRequest
	}

	return nil
}

/*
Refund executes a refund.
*/
func (client *Client) Refund(request RefundRequest) (*AuthorizationResponse, error) {

	if request.TransactionID != "" {
		request.TerminalName = ""
	}

	var response AuthorizationResponse
	var err error

	if request.IsTerminalRouted() {
		var route TerminalRoute
		route, err = client.resolveTerminalRoute(request.TerminalName)
		if err != nil {
			if err == ErrUnknownTerminal {
				response.ResponseDescription = ResponseUnknownTerminal
				return &response, err
			}

			return nil, err
		}

		if route.CloudRelayEnabled {
			err = client.RelayRequest("/refund", http.MethodPost, request, &response, request.Test)
		} else {
			terminalRequest := TerminalRefundAuthorizationRequest{
				APICredentials: route.TransientCredentials,
				Request:        request,
			}
			err = client.terminalPost(route, "/refund", terminalRequest, &response)
		}
	} else {
		err := client.GatewayRequest("/refund", http.MethodPost, request, &response, request.Test)
		if err, ok := err.(net.Error); ok && err.Timeout() {
			response.ResponseDescription = ResponseTimedOut
		} else if err != nil {
			response.ResponseDescription = err.Error()
		}
		return &response, err
	}

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err

}

/*
Reverse executes a manual time out reversal.
*/
func (client *Client) Reverse(request AuthorizationRequest) (*AuthorizationResponse, error) {

	var response AuthorizationResponse

	err := client.GatewayRequest("/reverse", http.MethodPost, request, &response, request.Test)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err

}

/*
Capture captures a preauthorization.
*/
func (client *Client) Capture(request CaptureRequest) (*CaptureResponse, error) {

	var captureResponse CaptureResponse

	err := client.GatewayRequest("/capture", http.MethodPost, request, &captureResponse, request.Test)
	if err, ok := err.(net.Error); ok && err.Timeout() {
		captureResponse.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		captureResponse.ResponseDescription = err.Error()
	}

	return &captureResponse, err

}

/*
Void discards a previous preauth transaction.
*/
func (client *Client) Void(request VoidRequest) (*VoidResponse, error) {

	var response VoidResponse

	err := client.GatewayRequest("/void", http.MethodPost, request, &response, request.Test)
	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

/*
AsyncEnroll executes an asynchronous vault enrollment.
*/
func (client *Client) AsyncEnroll(request EnrollRequest, responseChan chan<- EnrollResponse) error {

	if !isValidAsyncMethod(request.PaymentMethod) {
		return ErrInvalidAsyncRequest
	}

	return nil
}

/*
Enroll adds a new payment method to the token vault.
*/
func (client *Client) Enroll(request EnrollRequest) (*EnrollResponse, error) {

	var response EnrollResponse
	var err error

	if request.IsTerminalRouted() {
		_, err = client.resolveTerminalRoute(request.TerminalName)
		if err != nil {
			if err == ErrUnknownTerminal {
				response.ResponseDescription = ResponseUnknownTerminal
				return &response, err
			}

			return nil, err
		}
	} else {
		err = client.GatewayRequest("/enroll", http.MethodPost, request, &response, request.Test)
	}

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

/*
Ping tests connectivity with a payment terminal.
*/
func (client *Client) Ping(request PingRequest) (*PingResponse, error) {
	var response PingResponse

	route, err := client.resolveTerminalRoute(request.TerminalName)
	if err != nil {
		if err == ErrUnknownTerminal {
			response.ResponseDescription = ResponseUnknownTerminal
			return &response, err
		}

		return nil, err
	}

	if route.CloudRelayEnabled {
		err = client.RelayRequest("/terminal-test", http.MethodPost, request, &response, request.Test)
	} else {
		terminalRequest := TerminalPingRequest{
			APICredentials: route.TransientCredentials,
			Request:        request,
		}
		err = client.terminalPost(route, "/test", terminalRequest, &response)
	}

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

/*
GiftActivate activates or recharges a gift card.
*/
func (client *Client) GiftActivate(request GiftActivateRequest) (*GiftActivateResponse, error) {
	var response GiftActivateResponse

	route, err := client.resolveTerminalRoute(request.TerminalName)
	if err != nil {
		if err == ErrUnknownTerminal {
			response.ResponseDescription = ResponseUnknownTerminal
			return &response, err
		}

		return nil, err
	}

	if route.CloudRelayEnabled {
		err = client.RelayRequest("/gift-activate", http.MethodPost, request, &response, false)
	} else {
		terminalRequest := TerminalGiftActivateRequest{
			APICredentials: route.TransientCredentials,
			Request:        request,
		}
		err = client.terminalPost(route, "/gift-activate", terminalRequest, &response)
	}

	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err
}

/*
CloseBatch closes the current credit card batch.
*/
func (client *Client) CloseBatch(request CloseBatchRequest) (*CloseBatchResponse, error) {

	var response CloseBatchResponse

	err := client.GatewayRequest("/close-batch", http.MethodPost, request, &response, request.Test)
	if err, ok := err.(net.Error); ok && err.Timeout() {
		response.ResponseDescription = ResponseTimedOut
	} else if err != nil {
		response.ResponseDescription = err.Error()
	}

	return &response, err

}

// NewTransactionDisplay displays a new transaction on the terminal.
func (client *Client) NewTransactionDisplay(request TransactionDisplayRequest) error {
	return client.sendTransactionDisplay(request, http.MethodPost)
}

// UpdateTransactionDisplay appends items to an existing transaction display.
// Subtotal, Tax, and Total are overwritten by the request. Items with the same
// description are combined into groups.
func (client *Client) UpdateTransactionDisplay(request TransactionDisplayRequest) error {
	return client.sendTransactionDisplay(request, http.MethodPut)
}

// Clear clears the line item display and any in progress transaction
func (client *Client) Clear(request ClearTerminalRequest) (*Acknowledgement, error) {
	var ack Acknowledgement

	route, err := client.resolveTerminalRoute(request.TerminalName)
	if err != nil {
		if err == ErrUnknownTerminal {
			ack.Error = ResponseUnknownTerminal
			return &ack, err
		}

		return nil, err
	}

	if route.CloudRelayEnabled {
		err = client.RelayRequest("/terminal-clear", http.MethodPost, request, &ack, request.Test)
	} else {
		terminalRequest := TerminalClearTerminalRequest{
			APICredentials: route.TransientCredentials,
			Request:        request,
		}
		err = client.terminalPost(route, "/clear", terminalRequest, &ack)
	}

	if err, ok := err.(net.Error); ok && err.Timeout() {
		ack.Error = ResponseTimedOut
	} else if err != nil {
		ack.Error = err.Error()
	}

	return &ack, err

}

// sendTransactionDisplay sends a transaction display request to a terminal.
func (client *Client) sendTransactionDisplay(request TransactionDisplayRequest, method string) error {
	var response Acknowledgement

	route, err := client.resolveTerminalRoute(request.TerminalName)
	if err != nil {
		return err
	}

	if route.CloudRelayEnabled {
		err = client.RelayRequest("/terminal-txdisplay", method, request, &response, false)
	} else {
		terminalRequest := TerminalTransactionDisplayRequest{
			APICredentials: route.TransientCredentials,
			Request:        request,
		}
		err = client.terminalRequest(route, "/txdisplay", method, terminalRequest, &response)
	}
	if err != nil {
		return err
	}

	if !response.Success {
		return errors.New(response.Error)
	}

	return nil
}

func isValidAsyncMethod(method PaymentMethod) bool {

	if method.TerminalName == "" {
		return false
	} else if method.Token != "" {
		return false
	} else if method.Track1 != "" {
		return false
	} else if method.Track2 != "" {
		return false
	} else if method.PAN != "" {
		return false
	}

	return true

}
