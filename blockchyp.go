package blockchyp

/*
Default client configuration constants.
*/
const (
  DefaultGatewayHost = "api.blockchyp.com"
  DefaultTestGatewayHost = "test.blockchyp.com"
  DefaultHTTPS = true
  DefaultRouteCacheTTL = 60 //in minutes
)

/*
Client is the main interface used by application developers.
*/
type Client struct {
  Credentials APICredentials
  GatewayHost string
  HTTPS bool
  RouteCacheTTL uint
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
Charge executes a standard direct preauth and capture.
*/
func (client *Client) Charge(request AuthorizationRequest) (*AuthorizationResponse, error) {

  if isTerminalRouted(request.PaymentMethod) {
    _, err := resolveTerminalRoute(client.Credentials, request.TerminalName)
    if err != nil {
      return nil, err
    }

  } else {

  }

  return &AuthorizationResponse{}, nil
}

/*
Preauth executes a preauthorization intended to be captured later.
*/
func (client *Client) Preauth(request AuthorizationRequest) (*AuthorizationResponse, error) {

  return &AuthorizationResponse{}, nil
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

  return &CaptureResponse{}, nil
}

/*
Void discards a previous preauth transaction.
*/
func (client *Client) Void(request VoidRequest) (*VoidResponse, error) {

  return &VoidResponse{}, nil
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
