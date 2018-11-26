package blockchyp

import (
  "errors"
  "time"
)

/*
Default client configuration constants.
*/
const (
  DefaultGatewayHost = "api.blockchyp.com"
  DefaultTestGatewayHost = "test.blockchyp.com"
  DefaultHTTPS = true
  DefaultRouteCacheTTL = 60 //in minutes
  DefaultTimeout = 60 // in seconds
)

/*
Client is the main interface used by application developers.
*/
type Client struct {
  Credentials APICredentials
  GatewayHost string
  HTTPS bool
  RouteCacheTTL time.Duration
  Timeout uint64 //in seconds
}

/*
NewClient returns a default Client configured with the given credentials.
*/
func NewClient(creds APICredentials) Client {
  return Client{
    Credentials: creds,
    GatewayHost: DefaultGatewayHost,
    HTTPS: DefaultHTTPS,
    RouteCacheTTL: DefaultRouteCacheTTL,
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
    _, err := client.resolveTerminalRoute(request.TerminalName)
    if err != nil {
      return nil, err
    }

  } else {
    authResponse := AuthorizationResponse{}
    err := client.gatewayPost("/charge", request, &authResponse)
    return &authResponse, err
  }

  return &AuthorizationResponse{}, nil
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
    _, err := client.resolveTerminalRoute(request.TerminalName)
    if err != nil {
      return nil, err
    }

  } else {
    authResponse := AuthorizationResponse{}
    err := client.gatewayPost("/preauth", request, &authResponse)
    return &authResponse, err
  }

  return &AuthorizationResponse{}, nil

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
func (client *Client) Refund(request AuthorizationRequest) (*AuthorizationResponse, error) {

  return &AuthorizationResponse{}, nil
}

/*
Reverse executes a manual time out reversal.
*/
func (client *Client) Reverse(request AuthorizationRequest) (*AuthorizationResponse, error) {

  return &AuthorizationResponse{}, nil
}

/*
Capture captures a preauthorization.
*/
func (client *Client) Capture(request CaptureRequest) (*CaptureResponse, error) {

  captureResponse := CaptureResponse{}
  err := client.gatewayPost("/capture", request, &captureResponse)
  return &captureResponse, err

}

/*
Void discards a previous preauth transaction.
*/
func (client *Client) Void(request VoidRequest) (*VoidResponse, error) {

  return &VoidResponse{}, nil
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

  return &EnrollResponse{}, nil
}

/*
Ping tests connectivity with a payment terminal.
*/
func (client *Client) Ping(request PingRequest) (*PingResponse, error) {

  return &PingResponse{}, nil
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
  } else if (method.Token != "") {
    return false
  } else if (method.Track1 != "") {
    return false
  } else if (method.Track2 != "") {
    return false
  } else if (method.PAN != "") {
    return false
  }


  return true

}

func newInvalidAsyncRequestError() error {
  return errors.New("async requests must be terminal requests")
}
