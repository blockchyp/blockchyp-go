package blockchyp


/*
APICredentials models gateway credentials.
*/
type APICredentials struct {
  APIKey string `json:"apiKey"`
  BearerToken string `json:"bearerToken"`
  SigningKey string `json:"signingKey"`
}

/*
AuthorizationRequest models auth requests for charge, preauth, refund,
and reverse transaction types.
*/
type AuthorizationRequest struct {
  CoreRequest
  PaymentMethod
  RequestAmount
  Subtotals
  Enroll bool `json:"enroll"`
  Description string `json:"description,omitempty"`
  PromptForTip bool `json:"promptForTip,omitempty"`
  AltPrices map[string]string `json:"altPrices,omitempty"`
}

/*
PaymentMethod models fields for transactions that work with payment method data.
*/
type PaymentMethod struct {
  TerminalName string `json:"terminalName,omitempty"`
  Token string `json:"token,omitempty"`
  Track1 string `json:"track1,omitempty"`
  Track2 string `json:"track2,omitempty"`
  PAN string `json:"pan,omitempty"`
  CardholderName string `json:"cardholderName,omitempty"`
  ExpMonth string `json:"expMonth,omitempty"`
  ExpYear string `json:"expYear,omitempty"`
  CVV string `json:"cvv,omitempty"`
  Address string `json:"address,omitempty"`
  PostalCode string `json:"postalCode,omitempty"`
}

/*
RequestAmount models currency amounts in transaction requests.
*/
type RequestAmount struct {
  CurrencyCode string `json:"currencyCode"`
  Amount string `json:"amount"`
}

/*
Subtotals models subtotals like tip and tax amounts.
*/
type Subtotals struct {
  TipAmount string `json:"tipAmount,omitempty"`
  TaxAmount string `json:"taxAmount,omitempty"`
}

/*
PreviousTransaction models reference to a previous transaction.
*/
type PreviousTransaction struct {
  TransactionID string `json:"transactionId"`
}

/*
CoreRequest models fields that are common to all API requests.
*/
type CoreRequest struct {
  TransactionRef string `json:"transactionRef,omitempty"`
  Test bool `json:"test"`
}

/*
AuthorizationResponse models the response to authorization requests.
*/
type AuthorizationResponse struct {

}


/*
CaptureRequest models the information needed to capture a preauth.
*/
type CaptureRequest struct {
  CoreRequest
  PreviousTransaction
  RequestAmount
  Subtotals
}

/*
CaptureResponse models the response to a capture request.
*/
type CaptureResponse struct {

}

/*
VoidRequest models the information needed to nuke a preauth.
*/
type VoidRequest struct {
  CoreRequest
  PreviousTransaction
}

/*
VoidResponse models the response to a void request.
*/
type VoidResponse struct {

}

/*
EnrollRequest models the information needed to enroll a new payment method
in the token vault.
*/
type EnrollRequest struct {
  CoreRequest
  PaymentMethod
}

/*
EnrollResponse models the response to an enroll request.
*/
type EnrollResponse struct {

}

/*
PingRequest models the information needed to ping a terminal.
*/
type PingRequest struct {
  CoreRequest
  TerminalName string `json:"terminalName"`
}

/*
PingResponse models the response to a ping request.
*/
type PingResponse struct {

}

/*
GiftActivateRequest models the information needed to activate or recharge a
gift card.
*/
type GiftActivateRequest struct {
  CoreRequest
  RequestAmount
  TerminalName string `json:"terminalName"`
}

/*
GiftActivateResponse models the response to a gift activate request.
*/
type GiftActivateResponse struct {

}

/*
CloseBatchRequest models the information needed to manually close a credit
card batch.
*/
type CloseBatchRequest struct {
  CoreRequest
}

/*
CloseBatchResponse models the response to a close batch request.
*/
type CloseBatchResponse struct {

}
