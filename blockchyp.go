package blockchyp

/*
Charge executes a standard direct preauth and capture.
*/
func Charge(creds APICredentials, request AuthorizationRequest) (AuthorizationResponse, error) {

  return AuthorizationResponse{}, nil
}

/*
Preauth executes a preauthorization intended to be captured later.
*/
func Preauth(creds APICredentials, request AuthorizationRequest) (AuthorizationResponse, error) {

  return AuthorizationResponse{}, nil
}

/*
Refund executes a refund.
*/
func Refund(creds APICredentials, request AuthorizationRequest) (AuthorizationResponse, error) {

  return AuthorizationResponse{}, nil
}

/*
Reverse executes a manual time out reversal.
*/
func Reverse(creds APICredentials, request AuthorizationRequest) (AuthorizationResponse, error) {

  return AuthorizationResponse{}, nil
}

/*
Capture captures a preauthorization.
*/
func Capture(creds APICredentials, request CaptureRequest) (CaptureResponse, error) {

  return CaptureResponse{}, nil
}

/*
Void discards a previous preauth transaction.
*/
func Void(creds APICredentials, request VoidRequest) (VoidResponse, error) {

  return VoidResponse{}, nil
}

/*
Enroll adds a new payment method to the token vault.
*/
func Enroll(creds APICredentials, request EnrollRequest) (EnrollResponse, error) {

  return EnrollResponse{}, nil
}

/*
Ping tests connectivity with a payment terminal.
*/
func Ping(creds APICredentials, request PingRequest) (PingResponse, error) {

  return PingResponse{}, nil
}

/*
GiftActivate activates or recharges a gift card.
*/
func GiftActivate(creds APICredentials, request GiftActivateRequest) (GiftActivateResponse, error) {

  return GiftActivateResponse{}, nil
}

/*
CloseBatch closes the current credit card batch.
*/
func CloseBatch(creds APICredentials, request CloseBatchRequest) (CloseBatchResponse, error) {

  return CloseBatchResponse{}, nil
  
}
