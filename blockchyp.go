package blockchyp

import (
	"crypto/tls"
	"errors"
	"net/http"
	"time"
)

/*
Default client configuration constants.
*/
const (
	DefaultGatewayHost     = "api.blockchyp.com"
	DefaultTestGatewayHost = "test.blockchyp.com"
	DefaultHTTPS           = true
	DefaultRouteCacheTTL   = 60 * time.Minute
	DefaultGatewayTimeout  = 20 * time.Second
	DefaultTerminalTimeout = 2 * time.Minute
)

/*
Client is the main interface used by application developers.
*/
type Client struct {
	Credentials     APICredentials
	GatewayHost     string
	TestGatewayHost string
	HTTPS           bool

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
		gatewayHTTPClient: &http.Client{
			Timeout: DefaultGatewayTimeout,
		},
		terminalHTTPClient: &http.Client{
			Timeout: DefaultTerminalTimeout,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					RootCAs: terminalCertPool(),
				},
			},
		},
	}
}

/*
AsyncCharge executes an asynchronous auth and capture.
*/
func (client *Client) AsyncCharge(request AuthorizationRequest, responseChan chan<- AuthorizationResponse) error {

	if !isValidAsyncMethod(request.PaymentMethod) {
		return newInvalidAsyncRequestError()
	}

	return nil
}

/*
Charge executes a standard direct preauth and capture.
*/
func (client *Client) Charge(request AuthorizationRequest) (*AuthorizationResponse, error) {

	if isTerminalRouted(request.PaymentMethod) {
		route, err := client.resolveTerminalRoute(request.TerminalName)
		if err != nil {
			return nil, err
		}
		authRequest := TerminalAuthorizationRequest{
			APICredentials: route.TransientCredentials,
			Request:        request,
		}
		authResponse := AuthorizationResponse{}
		err = client.terminalPost(route, "/charge", authRequest, &authResponse)
		return &authResponse, err
	}
	authResponse := AuthorizationResponse{}
	err := client.GatewayPost("/charge", request, &authResponse)
	return &authResponse, err

}

/*
AsyncPreauth executes an asynchronous preauthorization.
*/
func (client *Client) AsyncPreauth(request AuthorizationRequest, responseChan chan<- AuthorizationResponse) error {

	if !isValidAsyncMethod(request.PaymentMethod) {
		return newInvalidAsyncRequestError()
	}

	return nil
}

/*
Preauth executes a preauthorization intended to be captured later.
*/
func (client *Client) Preauth(request AuthorizationRequest) (*AuthorizationResponse, error) {

	if isTerminalRouted(request.PaymentMethod) {
		route, err := client.resolveTerminalRoute(request.TerminalName)
		if err != nil {
			return nil, err
		}
		authRequest := TerminalAuthorizationRequest{
			APICredentials: route.TransientCredentials,
			Request:        request,
		}
		authResponse := AuthorizationResponse{}
		err = client.terminalPost(route, "/preauth", authRequest, &authResponse)
		return &authResponse, err

	}

	authResponse := AuthorizationResponse{}
	err := client.GatewayPost("/preauth", request, &authResponse)
	return &authResponse, err

}

/*
AsyncRefund executes an asynchronous refund
*/
func (client *Client) AsyncRefund(request AuthorizationRequest, responseChan chan<- AuthorizationResponse) error {

	if !isValidAsyncMethod(request.PaymentMethod) {
		return newInvalidAsyncRequestError()
	}

	return nil
}

/*
Refund executes a refund.
*/
func (client *Client) Refund(request RefundRequest) (*AuthorizationResponse, error) {

	if isTerminalRouted(request.PaymentMethod) {
		route, err := client.resolveTerminalRoute(request.TerminalName)
		if err != nil {
			return nil, err
		}
		authRequest := TerminalAuthorizationRequest{
			APICredentials: route.TransientCredentials,
			Request: AuthorizationRequest{
				CoreRequest:   request.CoreRequest,
				PaymentMethod: request.PaymentMethod,
				RequestAmount: request.RequestAmount,
				Subtotals:     request.Subtotals,
			},
		}
		authResponse := AuthorizationResponse{}
		err = client.terminalPost(route, "/refund", authRequest, &authResponse)
		return &authResponse, err

	}
	authResponse := AuthorizationResponse{}
	err := client.GatewayPost("/refund", request, &authResponse)
	return &authResponse, err

}

/*
Reverse executes a manual time out reversal.
*/
func (client *Client) Reverse(request AuthorizationRequest) (*AuthorizationResponse, error) {

	authResponse := AuthorizationResponse{}
	err := client.GatewayPost("/reverse", request, &authResponse)
	return &authResponse, err

}

/*
Capture captures a preauthorization.
*/
func (client *Client) Capture(request CaptureRequest) (*CaptureResponse, error) {

	captureResponse := CaptureResponse{}
	err := client.GatewayPost("/capture", request, &captureResponse)
	return &captureResponse, err

}

/*
Void discards a previous preauth transaction.
*/
func (client *Client) Void(request VoidRequest) (*VoidResponse, error) {

	voidResponse := VoidResponse{}
	err := client.GatewayPost("/void", request, &voidResponse)
	return &voidResponse, err
}

/*
AsyncEnroll executes an asynchronous vault enrollment.
*/
func (client *Client) AsyncEnroll(request EnrollRequest, responseChan chan<- EnrollResponse) error {

	if !isValidAsyncMethod(request.PaymentMethod) {
		return newInvalidAsyncRequestError()
	}

	return nil
}

/*
Enroll adds a new payment method to the token vault.
*/
func (client *Client) Enroll(request EnrollRequest) (*EnrollResponse, error) {

	if isTerminalRouted(request.PaymentMethod) {
		_, err := client.resolveTerminalRoute(request.TerminalName)
		if err != nil {
			return nil, err
		}

	} else {
		enrollResponse := EnrollResponse{}
		err := client.GatewayPost("/enroll", request, &enrollResponse)
		return &enrollResponse, err
	}

	return &EnrollResponse{}, nil
}

/*
Ping tests connectivity with a payment terminal.
*/
func (client *Client) Ping(request PingRequest) (*PingResponse, error) {
	route, err := client.resolveTerminalRoute(request.TerminalName)
	if err != nil {
		return nil, err
	}
	pingResponse := PingResponse{}
	err = client.terminalPost(route, "/test", request, &pingResponse)
	return &pingResponse, err
}

/*
GiftActivate activates or recharges a gift card.
*/
func (client *Client) GiftActivate(request GiftActivateRequest) (*GiftActivateResponse, error) {

	return &GiftActivateResponse{}, nil
}

/*
CloseBatch closes the current credit card batch.
*/
func (client *Client) CloseBatch(request CloseBatchRequest) (*CloseBatchResponse, error) {

	return &CloseBatchResponse{}, nil

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

func newInvalidAsyncRequestError() error {
	return errors.New("async requests must be terminal requests")
}
