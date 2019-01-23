package blockchyp

/*
APICredentials models gateway credentials.
*/
type APICredentials struct {
	APIKey      string `json:"apiKey"`
	BearerToken string `json:"bearerToken"`
	SigningKey  string `json:"signingKey"`
}

/*
AuthorizationRequest models auth requests for charge, preauth,
and reverse transaction types.
*/
type AuthorizationRequest struct {
	CoreRequest
	PaymentMethod
	RequestAmount
	Subtotals
	Enroll       bool              `json:"enroll"`
	Description  string            `json:"description,omitempty"`
	PromptForTip bool              `json:"promptForTip,omitempty"`
	AltPrices    map[string]string `json:"altPrices,omitempty"`
}

/*
RefundRequest models refund requests.
*/
type RefundRequest struct {
	CoreRequest
	PaymentMethod
	RequestAmount
	Subtotals
	PreviousTransaction
}

/*
PaymentMethod models fields for transactions that work with payment method data.
*/
type PaymentMethod struct {
	TerminalName   string `json:"terminalName,omitempty"`
	Token          string `json:"token,omitempty"`
	Track1         string `json:"track1,omitempty"`
	Track2         string `json:"track2,omitempty"`
	PAN            string `json:"pan,omitempty"`
	CardholderName string `json:"cardholderName,omitempty"`
	ExpMonth       string `json:"expMonth,omitempty"`
	ExpYear        string `json:"expYear,omitempty"`
	CVV            string `json:"cvv,omitempty"`
	Address        string `json:"address,omitempty"`
	PostalCode     string `json:"postalCode,omitempty"`
}

/*
RequestAmount models currency amounts in transaction requests.
*/
type RequestAmount struct {
	CurrencyCode string `json:"currencyCode"`
	Amount       string `json:"amount"`
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
	Test           bool   `json:"test"`
}

/*
CoreResponse models elements common to all API responses.
*/
type CoreResponse struct {
	ResponseDescription string `json:"responseDescription"`
	TransactionID       string `json:"transactionId"`
	BatchID             string `json:"batchId,omitempty"`
	TransactionRef      string `json:"transactionRef,omitempty"`
	TransactionType     string `json:"transactionType"`
	Timestamp           string `json:"timestamp"`
	TickBlock           string `json:"tickBlock"`
}

/*
ApprovalResponse models data related to approval or failure of a transaction.
*/
type ApprovalResponse struct {
	Approved   bool   `json:"approved"`
	AuthCode   string `json:"authCode"`
	SigCapture string `json:"sigCapture"`
}

/*
AuthorizationResponse models the response to authorization requests.
*/
type AuthorizationResponse struct {
	CoreResponse
	ApprovalResponse
	PaymentMethodResponse
	PaymentAmounts
	Sig string `json:"sig,omitempty"`
}

/*
PaymentAmounts models the amounts and currency data in responses.
*/
type PaymentAmounts struct {
	PartialAuth      bool   `json:"partialAuth"`
	AltCurrency      bool   `json:"altCurrency"`
	CurrencyCode     string `json:"currencyCode"`
	RequestedAmount  string `json:"requestedAmount"`
	AuthorizedAmount string `json:"authorizedAmount"`
	TipAmount        string `json:"tipAmount,omitempty"`
	TaxAmount        string `json:"taxAmount,omitempty"`
}

/*
PaymentMethodResponse models response data about payment methods.  Could be
used for non-authorization transactions that still work with payment methods.
*/
type PaymentMethodResponse struct {
	Token              string             `json:"token,omitempty"`
	EntryMethod        string             `json:"entryMethod,omitempty"`
	PaymentType        string             `json:"paymentType,omitempty"`
	MaskedPAN          string             `json:"maskedPan,omitempty"`
	PublicKey          string             `json:"publicKey,omitempty"`
	ScopeAlert         bool               `json:"scopeAlert,omitempty"`
	CardHolder         string             `json:"cardHolder,omitempty"`
	ReceiptSuggestions ReceiptSuggestions `json:"receiptSuggestions,omitempty"`
}

/*
ReceiptSuggestions models EMV fields we recommend developers put on their receipts.
*/
type ReceiptSuggestions struct {
	AID              string `json:"AID,omitempty"`
	ARQC             string `json:"ARQC,omitempty"`
	IAD              string `json:"IAD,omitempty"`
	ARC              string `json:"ARC,omitempty"`
	TC               string `json:"TC,omitempty"`
	TVR              string `json:"TVR,omitempty"`
	TSI              string `json:"TSI,omitempty"`
	TerminalID       string `json:"terminalId,omitempty"`
	MerchantName     string `json:"merchantName,omitempty"`
	MerchantID       string `json:"merchantId,omitempty"`
	MerchantKey      string `json:"merchantKey,omitempty"`
	ApplicationLabel string `json:"applicationLabel,omitempty"`
	RequestSignature bool   `json:"requestSignature,omitempty"`
	MaskedPAN        string `json:"maskedPan,omitempty"`
	AuthorizedAmount string `json:"authorizedAmount"`
	TransactionType  string `json:"transactionType"`
	EntryMethod      string `json:"entryMethod,omitempty"`
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
Acknowledgement models a basic api acknowledgement.
*/
type Acknowledgement struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

/*
CaptureResponse models the response to a capture request.
*/
type CaptureResponse struct {
	CoreResponse
	ApprovalResponse
	PaymentMethodResponse
	PaymentAmounts
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
	CoreResponse
	ApprovalResponse
	PaymentMethodResponse
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
	CoreResponse
	PaymentMethodResponse
	ApprovalResponse
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
	Success bool `json:"success"`
	CoreResponse
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
	CoreResponse
	Approved       bool   `json:"approved"`
	Amount         string `json:"amount"`
	CurrentBalance string `json:"currentBalance"`
	CurrencyCode   string `json:"currencyCode"`
	PublicKey      string `json:"publicKey"`
	Sig            string `json:"sig,omitempty"`
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
	CoreResponse
	Success       bool              `json:"success"`
	CurrencyCode  string            `json:"currencyCode"`
	CapturedTotal string            `json:"capturedTotal"`
	OpenPreauths  string            `json:"openPreauths"`
	CardBrands    map[string]string `json:"cardBrands"`
}
