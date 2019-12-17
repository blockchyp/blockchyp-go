package blockchyp

import (
	"time"
)

// APICredentials models gateway credentials.
type APICredentials struct {
	APIKey      string `json:"apiKey"`
	BearerToken string `json:"bearerToken"`
	SigningKey  string `json:"signingKey"`
}

// CardType is used to differentiate credit, debit, and EBT.
type CardType int

// CardTypes.
const (
	CardTypeCredit CardType = iota
	CardTypeDebit
	CardTypeEBT
	CardTypeBlockchainGift
)

// Acknowledgement models a basic api acknowledgement.
type Acknowledgement struct {
	Success bool `json:"success"`

	Error string `json:"error"`

	ResponseDescription string `json:"responseDescription"`
}

// PingRequest models information needed to test connectivity with a terminal.
type PingRequest struct {
	TransactionRef string `json:"transactionRef,omitempty"`

	OrderRef string `json:"orderRef,omitempty"`

	DestinationAccount string `json:"destinationAccount,omitempty"`

	Test bool `json:"test"`

	Timeout int `json:"timeout"`

	TerminalName string `json:"terminalName,omitempty"`
}

// PingResponse models the response to a ping request.
type PingResponse struct {
	Success bool `json:"success"`

	Error string `json:"error"`

	ResponseDescription string `json:"responseDescription"`

	TransactionID string `json:"transactionId"`

	BatchID string `json:"batchId,omitempty"`

	TransactionRef string `json:"transactionRef,omitempty"`

	TransactionType string `json:"transactionType"`

	Timestamp string `json:"timestamp"`

	TickBlock string `json:"tickBlock"`

	Test bool `json:"test"`

	Sig string `json:"sig,omitempty"`
}

// MessageRequest models a message to be displayed on the terminal screen.
type MessageRequest struct {
	TransactionRef string `json:"transactionRef,omitempty"`

	OrderRef string `json:"orderRef,omitempty"`

	DestinationAccount string `json:"destinationAccount,omitempty"`

	Test bool `json:"test"`

	Timeout int `json:"timeout"`

	TerminalName string `json:"terminalName,omitempty"`

	Message string `json:"message"`
}

// BooleanPromptRequest models a simple yes no prompt request.
type BooleanPromptRequest struct {
	TransactionRef string `json:"transactionRef,omitempty"`

	OrderRef string `json:"orderRef,omitempty"`

	DestinationAccount string `json:"destinationAccount,omitempty"`

	Test bool `json:"test"`

	Timeout int `json:"timeout"`

	TerminalName string `json:"terminalName,omitempty"`

	YesCaption string `json:"yesCaption"`

	NoCaption string `json:"noCaption"`

	Prompt string `json:"prompt"`
}

// TextPromptRequest models a text prompt request.
type TextPromptRequest struct {
	TransactionRef string `json:"transactionRef,omitempty"`

	OrderRef string `json:"orderRef,omitempty"`

	DestinationAccount string `json:"destinationAccount,omitempty"`

	Test bool `json:"test"`

	Timeout int `json:"timeout"`

	TerminalName string `json:"terminalName,omitempty"`

	PromptType string `json:"promptType"`
}

// TextPromptResponse models the response to a text prompt request.
type TextPromptResponse struct {
	Success bool `json:"success"`

	Error string `json:"error"`

	ResponseDescription string `json:"responseDescription"`

	Response string `json:"response"`
}

// BooleanPromptResponse models the response to a boolean prompt request.
type BooleanPromptResponse struct {
	Success bool `json:"success"`

	Error string `json:"error"`

	ResponseDescription string `json:"responseDescription"`

	Response bool `json:"response"`
}

// WhiteListedCard shows details about a white listed card.
type WhiteListedCard struct {
	Bin string `json:"bin"`

	Track1 string `json:"track1"`

	Track2 string `json:"track2"`

	Pan string `json:"pan"`
}

// AuthorizationRequest models auth requests for charge, preauth, and reverse transaction types.
type AuthorizationRequest struct {
	TransactionRef string `json:"transactionRef,omitempty"`

	OrderRef string `json:"orderRef,omitempty"`

	DestinationAccount string `json:"destinationAccount,omitempty"`

	Test bool `json:"test"`

	Timeout int `json:"timeout"`

	TerminalName string `json:"terminalName,omitempty"`

	Token string `json:"token,omitempty"`

	Track1 string `json:"track1,omitempty"`

	Track2 string `json:"track2,omitempty"`

	PAN string `json:"pan,omitempty"`

	RoutingNumber string `json:"routingNumber,omitempty"`

	CardholderName string `json:"cardholderName,omitempty"`

	ExpMonth string `json:"expMonth,omitempty"`

	ExpYear string `json:"expYear,omitempty"`

	CVV string `json:"cvv,omitempty"`

	Address string `json:"address,omitempty"`

	PostalCode string `json:"postalCode,omitempty"`

	// ManualEntry specifies that the payment entry method is a manual keyed transaction. If this is true, no other payment method will be accepted.
	ManualEntry bool `json:"manualEntry,omitempty"`

	// KSN is the key serial number used for DUKPT encryption.
	KSN string `json:"ksn,omitempty"`

	// PINBlock is the encrypted pin block.
	PINBlock string `json:"pinBlock,omitempty"`

	// CardType designates categories of cards: credit, debit, EBT.
	CardType CardType `json:"cardType,omitempty"`

	// PaymentType designates brands of payment methods: Visa, Discover, etc.
	PaymentType string `json:"paymentType,omitempty"`

	CurrencyCode string `json:"currencyCode"`

	Amount string `json:"amount"`

	TaxExempt bool `json:"taxExempt"`

	TipAmount string `json:"tipAmount,omitempty"`

	TaxAmount string `json:"taxAmount,omitempty"`

	CashBackAmount string `json:"cashBackAmount,omitempty"`

	FSAEligibleAmount string `json:"fsaEligibleAmount,omitempty"`

	HSAEligibleAmount string `json:"hsaEligibleAmount,omitempty"`

	EBTEligibleAmount string `json:"ebtEligibleAmount,omitempty"`

	SigFile string `json:"sigFile,omitempty"`

	// SigFormat specifies the image format to be used for returning signatures. (PNG or JPEG).
	SigFormat string `json:"sigFormat,omitempty"`

	// SigWidth scales the signature image to the given width, preserving the aspect ratio.  If not provided, the signature is returned in the terminal's max resolution.
	SigWidth int `json:"sigWidth,omitempty"`

	TransactionID string `json:"transactionId"`

	// OnlineAuthCode is used to validate online gift card authorizations.
	OnlineAuthCode string `json:"onlineAuthCode,omitempty"`

	// Enroll adds the payment method to the token vault alongside the authorization.
	Enroll bool `json:"enroll,omitempty"`

	Description string `json:"description,omitempty"`

	PromptForTip bool `json:"promptForTip,omitempty"`

	CashBackEnabled bool `json:"cashBackEnabled,omitempty"`

	AltPrices map[string]string `json:"altPrices,omitempty"`
}

// BalanceRequest models balance requests.
type BalanceRequest struct {
	TransactionRef string `json:"transactionRef,omitempty"`

	OrderRef string `json:"orderRef,omitempty"`

	DestinationAccount string `json:"destinationAccount,omitempty"`

	Test bool `json:"test"`

	Timeout int `json:"timeout"`

	Token string `json:"token,omitempty"`

	Track1 string `json:"track1,omitempty"`

	Track2 string `json:"track2,omitempty"`

	PAN string `json:"pan,omitempty"`

	RoutingNumber string `json:"routingNumber,omitempty"`

	CardholderName string `json:"cardholderName,omitempty"`

	ExpMonth string `json:"expMonth,omitempty"`

	ExpYear string `json:"expYear,omitempty"`

	CVV string `json:"cvv,omitempty"`

	Address string `json:"address,omitempty"`

	PostalCode string `json:"postalCode,omitempty"`

	// ManualEntry specifies that the payment entry method is a manual keyed transaction. If this is true, no other payment method will be accepted.
	ManualEntry bool `json:"manualEntry,omitempty"`

	// KSN is the key serial number used for DUKPT encryption.
	KSN string `json:"ksn,omitempty"`

	// PINBlock is the encrypted pin block.
	PINBlock string `json:"pinBlock,omitempty"`

	// CardType designates categories of cards: credit, debit, EBT.
	CardType CardType `json:"cardType,omitempty"`

	// PaymentType designates brands of payment methods: Visa, Discover, etc.
	PaymentType string `json:"paymentType,omitempty"`

	TerminalName string `json:"terminalName,omitempty"`
}

// BalanceResponse models the response to a balance request.
type BalanceResponse struct {
	Success bool `json:"success"`

	Error string `json:"error"`

	ResponseDescription string `json:"responseDescription"`

	TransactionID string `json:"transactionId"`

	BatchID string `json:"batchId,omitempty"`

	TransactionRef string `json:"transactionRef,omitempty"`

	TransactionType string `json:"transactionType"`

	Timestamp string `json:"timestamp"`

	TickBlock string `json:"tickBlock"`

	Test bool `json:"test"`

	Sig string `json:"sig,omitempty"`

	Token string `json:"token,omitempty"`

	EntryMethod string `json:"entryMethod,omitempty"`

	PaymentType string `json:"paymentType,omitempty"`

	MaskedPAN string `json:"maskedPan,omitempty"`

	PublicKey string `json:"publicKey,omitempty"`

	ScopeAlert bool `json:"ScopeAlert,omitempty"`

	CardHolder string `json:"cardHolder,omitempty"`

	ReceiptSuggestions ReceiptSuggestions `json:"receiptSuggestions"`

	// RemainingBalance remaining balance on the payment method.
	RemainingBalance string `json:"remainingBalance,omitempty"`
}

// RefundRequest models refund requests.
type RefundRequest struct {
	TransactionRef string `json:"transactionRef,omitempty"`

	OrderRef string `json:"orderRef,omitempty"`

	DestinationAccount string `json:"destinationAccount,omitempty"`

	Test bool `json:"test"`

	Timeout int `json:"timeout"`

	TerminalName string `json:"terminalName,omitempty"`

	Token string `json:"token,omitempty"`

	Track1 string `json:"track1,omitempty"`

	Track2 string `json:"track2,omitempty"`

	PAN string `json:"pan,omitempty"`

	RoutingNumber string `json:"routingNumber,omitempty"`

	CardholderName string `json:"cardholderName,omitempty"`

	ExpMonth string `json:"expMonth,omitempty"`

	ExpYear string `json:"expYear,omitempty"`

	CVV string `json:"cvv,omitempty"`

	Address string `json:"address,omitempty"`

	PostalCode string `json:"postalCode,omitempty"`

	// ManualEntry specifies that the payment entry method is a manual keyed transaction. If this is true, no other payment method will be accepted.
	ManualEntry bool `json:"manualEntry,omitempty"`

	// KSN is the key serial number used for DUKPT encryption.
	KSN string `json:"ksn,omitempty"`

	// PINBlock is the encrypted pin block.
	PINBlock string `json:"pinBlock,omitempty"`

	// CardType designates categories of cards: credit, debit, EBT.
	CardType CardType `json:"cardType,omitempty"`

	// PaymentType designates brands of payment methods: Visa, Discover, etc.
	PaymentType string `json:"paymentType,omitempty"`

	CurrencyCode string `json:"currencyCode"`

	Amount string `json:"amount"`

	TaxExempt bool `json:"taxExempt"`

	TipAmount string `json:"tipAmount,omitempty"`

	TaxAmount string `json:"taxAmount,omitempty"`

	CashBackAmount string `json:"cashBackAmount,omitempty"`

	FSAEligibleAmount string `json:"fsaEligibleAmount,omitempty"`

	HSAEligibleAmount string `json:"hsaEligibleAmount,omitempty"`

	EBTEligibleAmount string `json:"ebtEligibleAmount,omitempty"`

	SigFile string `json:"sigFile,omitempty"`

	// SigFormat specifies the image format to be used for returning signatures. (PNG or JPEG).
	SigFormat string `json:"sigFormat,omitempty"`

	// SigWidth scales the signature image to the given width, preserving the aspect ratio.  If not provided, the signature is returned in the terminal's max resolution.
	SigWidth int `json:"sigWidth,omitempty"`

	TransactionID string `json:"transactionId"`
}

// CaptureRequest models the information needed to capture a preauth.
type CaptureRequest struct {
	TransactionRef string `json:"transactionRef,omitempty"`

	OrderRef string `json:"orderRef,omitempty"`

	DestinationAccount string `json:"destinationAccount,omitempty"`

	Test bool `json:"test"`

	Timeout int `json:"timeout"`

	CurrencyCode string `json:"currencyCode"`

	Amount string `json:"amount"`

	TaxExempt bool `json:"taxExempt"`

	TipAmount string `json:"tipAmount,omitempty"`

	TaxAmount string `json:"taxAmount,omitempty"`

	CashBackAmount string `json:"cashBackAmount,omitempty"`

	FSAEligibleAmount string `json:"fsaEligibleAmount,omitempty"`

	HSAEligibleAmount string `json:"hsaEligibleAmount,omitempty"`

	EBTEligibleAmount string `json:"ebtEligibleAmount,omitempty"`

	TransactionID string `json:"transactionId"`
}

// CaptureResponse models the response to a capture request.
type CaptureResponse struct {
	Success bool `json:"success"`

	Error string `json:"error"`

	ResponseDescription string `json:"responseDescription"`

	TransactionID string `json:"transactionId"`

	BatchID string `json:"batchId,omitempty"`

	TransactionRef string `json:"transactionRef,omitempty"`

	TransactionType string `json:"transactionType"`

	Timestamp string `json:"timestamp"`

	TickBlock string `json:"tickBlock"`

	Test bool `json:"test"`

	Sig string `json:"sig,omitempty"`

	Approved bool `json:"approved"`

	AuthCode string `json:"authCode,omitempty"`

	SigFile string `json:"sigFile"`

	Token string `json:"token,omitempty"`

	EntryMethod string `json:"entryMethod,omitempty"`

	PaymentType string `json:"paymentType,omitempty"`

	MaskedPAN string `json:"maskedPan,omitempty"`

	PublicKey string `json:"publicKey,omitempty"`

	ScopeAlert bool `json:"ScopeAlert,omitempty"`

	CardHolder string `json:"cardHolder,omitempty"`

	ReceiptSuggestions ReceiptSuggestions `json:"receiptSuggestions"`

	PartialAuth bool `json:"partialAuth"`

	AltCurrency bool `json:"altCurrency"`

	FSAAuth bool `json:"fasAuth"`

	CurrencyCode string `json:"currencyCode"`

	RequestedAmount string `json:"requestedAmount"`

	AuthorizedAmount string `json:"authorizedAmount"`

	RemainingBalance string `json:"remainingBalance"`

	TipAmount string `json:"tipAmount"`

	TaxAmount string `json:"taxAmount"`

	RequestedCashBackAmount string `json:"requestedCashBackAmount"`

	AuthorizedCashBackAmount string `json:"authorizedCashBackAmount"`
}

// VoidRequest models a void request.
type VoidRequest struct {
	TransactionRef string `json:"transactionRef,omitempty"`

	OrderRef string `json:"orderRef,omitempty"`

	DestinationAccount string `json:"destinationAccount,omitempty"`

	Test bool `json:"test"`

	Timeout int `json:"timeout"`

	TransactionID string `json:"transactionId"`
}

// VoidResponse models the response to a void request.
type VoidResponse struct {
	Success bool `json:"success"`

	Error string `json:"error"`

	ResponseDescription string `json:"responseDescription"`

	TransactionID string `json:"transactionId"`

	BatchID string `json:"batchId,omitempty"`

	TransactionRef string `json:"transactionRef,omitempty"`

	TransactionType string `json:"transactionType"`

	Timestamp string `json:"timestamp"`

	TickBlock string `json:"tickBlock"`

	Test bool `json:"test"`

	Sig string `json:"sig,omitempty"`

	Approved bool `json:"approved"`

	AuthCode string `json:"authCode,omitempty"`

	SigFile string `json:"sigFile"`

	Token string `json:"token,omitempty"`

	EntryMethod string `json:"entryMethod,omitempty"`

	PaymentType string `json:"paymentType,omitempty"`

	MaskedPAN string `json:"maskedPan,omitempty"`

	PublicKey string `json:"publicKey,omitempty"`

	ScopeAlert bool `json:"ScopeAlert,omitempty"`

	CardHolder string `json:"cardHolder,omitempty"`

	ReceiptSuggestions ReceiptSuggestions `json:"receiptSuggestions"`
}

// EnrollRequest models the information needed to enroll a new payment method in the token vault.
type EnrollRequest struct {
	TransactionRef string `json:"transactionRef,omitempty"`

	OrderRef string `json:"orderRef,omitempty"`

	DestinationAccount string `json:"destinationAccount,omitempty"`

	Test bool `json:"test"`

	Timeout int `json:"timeout"`

	TerminalName string `json:"terminalName,omitempty"`

	Token string `json:"token,omitempty"`

	Track1 string `json:"track1,omitempty"`

	Track2 string `json:"track2,omitempty"`

	PAN string `json:"pan,omitempty"`

	RoutingNumber string `json:"routingNumber,omitempty"`

	CardholderName string `json:"cardholderName,omitempty"`

	ExpMonth string `json:"expMonth,omitempty"`

	ExpYear string `json:"expYear,omitempty"`

	CVV string `json:"cvv,omitempty"`

	Address string `json:"address,omitempty"`

	PostalCode string `json:"postalCode,omitempty"`

	// ManualEntry specifies that the payment entry method is a manual keyed transaction. If this is true, no other payment method will be accepted.
	ManualEntry bool `json:"manualEntry,omitempty"`

	// KSN is the key serial number used for DUKPT encryption.
	KSN string `json:"ksn,omitempty"`

	// PINBlock is the encrypted pin block.
	PINBlock string `json:"pinBlock,omitempty"`

	// CardType designates categories of cards: credit, debit, EBT.
	CardType CardType `json:"cardType,omitempty"`

	// PaymentType designates brands of payment methods: Visa, Discover, etc.
	PaymentType string `json:"paymentType,omitempty"`
}

// EnrollResponse models the response to an enroll request.
type EnrollResponse struct {
	Success bool `json:"success"`

	Error string `json:"error"`

	ResponseDescription string `json:"responseDescription"`

	TransactionID string `json:"transactionId"`

	BatchID string `json:"batchId,omitempty"`

	TransactionRef string `json:"transactionRef,omitempty"`

	TransactionType string `json:"transactionType"`

	Timestamp string `json:"timestamp"`

	TickBlock string `json:"tickBlock"`

	Test bool `json:"test"`

	Sig string `json:"sig,omitempty"`

	Token string `json:"token,omitempty"`

	EntryMethod string `json:"entryMethod,omitempty"`

	PaymentType string `json:"paymentType,omitempty"`

	MaskedPAN string `json:"maskedPan,omitempty"`

	PublicKey string `json:"publicKey,omitempty"`

	ScopeAlert bool `json:"ScopeAlert,omitempty"`

	CardHolder string `json:"cardHolder,omitempty"`

	ReceiptSuggestions ReceiptSuggestions `json:"receiptSuggestions"`

	Approved bool `json:"approved"`

	AuthCode string `json:"authCode,omitempty"`

	SigFile string `json:"sigFile"`
}

// ClearTerminalRequest models the information needed to enroll a new payment method in the token vault.
type ClearTerminalRequest struct {
	TransactionRef string `json:"transactionRef,omitempty"`

	OrderRef string `json:"orderRef,omitempty"`

	DestinationAccount string `json:"destinationAccount,omitempty"`

	Test bool `json:"test"`

	Timeout int `json:"timeout"`

	TerminalName string `json:"terminalName,omitempty"`
}

// GiftActivateRequest models the information needed to activate or recharge a gift card.
type GiftActivateRequest struct {
	TransactionRef string `json:"transactionRef,omitempty"`

	OrderRef string `json:"orderRef,omitempty"`

	DestinationAccount string `json:"destinationAccount,omitempty"`

	Test bool `json:"test"`

	Timeout int `json:"timeout"`

	CurrencyCode string `json:"currencyCode"`

	Amount string `json:"amount"`

	TaxExempt bool `json:"taxExempt"`

	TerminalName string `json:"terminalName,omitempty"`
}

// GiftActivateResponse models the response to a gift activate request.
type GiftActivateResponse struct {
	Success bool `json:"success"`

	Error string `json:"error"`

	ResponseDescription string `json:"responseDescription"`

	TransactionID string `json:"transactionId"`

	BatchID string `json:"batchId,omitempty"`

	TransactionRef string `json:"transactionRef,omitempty"`

	TransactionType string `json:"transactionType"`

	Timestamp string `json:"timestamp"`

	TickBlock string `json:"tickBlock"`

	Test bool `json:"test"`

	Sig string `json:"sig,omitempty"`

	Approved bool `json:"approved"`

	Amount string `json:"amount"`

	CurrentBalance string `json:"currentBalance"`

	CurrencyCode string `json:"currencyCode"`

	PublicKey string `json:"publicKey"`
}

// CloseBatchRequest models the information needed to manually close a credit card batch.
type CloseBatchRequest struct {
	TransactionRef string `json:"transactionRef,omitempty"`

	OrderRef string `json:"orderRef,omitempty"`

	DestinationAccount string `json:"destinationAccount,omitempty"`

	Test bool `json:"test"`

	Timeout int `json:"timeout"`
}

// CloseBatchResponse models the response to a close batch request.
type CloseBatchResponse struct {
	Success bool `json:"success"`

	Error string `json:"error"`

	ResponseDescription string `json:"responseDescription"`

	TransactionID string `json:"transactionId"`

	BatchID string `json:"batchId,omitempty"`

	TransactionRef string `json:"transactionRef,omitempty"`

	TransactionType string `json:"transactionType"`

	Timestamp string `json:"timestamp"`

	TickBlock string `json:"tickBlock"`

	Test bool `json:"test"`

	Sig string `json:"sig,omitempty"`

	CurrencyCode string `json:"currencyCode"`

	CapturedTotal string `json:"capturedTotal"`

	OpenPreauths string `json:"openPreauths"`

	CardBrands map[string]string `json:"cardBrands"`
}

// TermsAndConditionsRequest models the fields needed for custom T&C prompts.
type TermsAndConditionsRequest struct {
	TransactionRef string `json:"transactionRef,omitempty"`

	OrderRef string `json:"orderRef,omitempty"`

	DestinationAccount string `json:"destinationAccount,omitempty"`

	Test bool `json:"test"`

	Timeout int `json:"timeout"`

	TerminalName string `json:"terminalName,omitempty"`

	TransactionID string `json:"transactionId"`

	SigFile string `json:"sigFile,omitempty"`

	// SigFormat specifies the image format to be used for returning signatures. (PNG or JPEG).
	SigFormat string `json:"sigFormat,omitempty"`

	// SigWidth scales the signature image to the given width, preserving the aspect ratio.  If not provided, the signature is returned in the terminal's max resolution.
	SigWidth int `json:"sigWidth,omitempty"`

	TCAlias string `json:"tcAlias"`

	TCName string `json:"tcName"`

	TCContent string `json:"tcContent"`

	SigRequired bool `json:"sigRequired"`
}

// TermsAndConditionsResponse models a T&C signature capture response.
type TermsAndConditionsResponse struct {
	Success bool `json:"success"`

	Error string `json:"error"`

	ResponseDescription string `json:"responseDescription"`

	TransactionID string `json:"transactionId"`

	BatchID string `json:"batchId,omitempty"`

	TransactionRef string `json:"transactionRef,omitempty"`

	TransactionType string `json:"transactionType"`

	Timestamp string `json:"timestamp"`

	TickBlock string `json:"tickBlock"`

	Test bool `json:"test"`

	Sig string `json:"sig,omitempty"`

	SigFile string `json:"sigFile,omitempty"`
}

// ReceiptSuggestions models EMV fields we recommend developers put on their receipts.
type ReceiptSuggestions struct {

	// AID is the EMV Application Identifier.
	AID string `json:"AID,omitempty"`

	// ARQC is the EMV Application Request Cryptogram.
	ARQC string `json:"ARQC,omitempty"`

	// IAD is EMV Issuer Application Data.
	IAD string `json:"IAD,omitempty"`

	// ARC is the EMV Authorization Response Code.
	ARC string `json:"ARC,omitempty"`

	// TC is the EMV Transaction Certificate.
	TC string `json:"TC,omitempty"`

	// TVR is the EMV Terminal Verification Response.
	TVR string `json:"TVR,omitempty"`

	// TSI is the EMV Transaction Status Indicator.
	TSI string `json:"TSI,omitempty"`

	TerminalID string `json:"terminalId,omitempty"`

	MerchantName string `json:"merchantName,omitempty"`

	MerchantID string `json:"merchantId,omitempty"`

	MerchantKey string `json:"merchantKey,omitempty"`

	ApplicationLabel string `json:"applicationLabel,omitempty"`

	RequestSignature bool `json:"requestSignature"`

	MaskedPAN string `json:"maskedPan,omitempty"`

	AuthorizedAmount string `json:"authorizedAmount"`

	TransactionType string `json:"transactionType"`

	EntryMethod string `json:"entryMethod,omitempty"`

	PINVerified bool `json:"pinVerified,omitempty"`

	CashBackAmount string `json:"cashBackAmount,omitempty"`
}

// AuthorizationResponse models the response to authorization requests.
type AuthorizationResponse struct {
	Success bool `json:"success"`

	Error string `json:"error"`

	ResponseDescription string `json:"responseDescription"`

	TransactionID string `json:"transactionId"`

	BatchID string `json:"batchId,omitempty"`

	TransactionRef string `json:"transactionRef,omitempty"`

	TransactionType string `json:"transactionType"`

	Timestamp string `json:"timestamp"`

	TickBlock string `json:"tickBlock"`

	Test bool `json:"test"`

	Sig string `json:"sig,omitempty"`

	Approved bool `json:"approved"`

	AuthCode string `json:"authCode,omitempty"`

	SigFile string `json:"sigFile"`

	Token string `json:"token,omitempty"`

	EntryMethod string `json:"entryMethod,omitempty"`

	PaymentType string `json:"paymentType,omitempty"`

	MaskedPAN string `json:"maskedPan,omitempty"`

	PublicKey string `json:"publicKey,omitempty"`

	ScopeAlert bool `json:"ScopeAlert,omitempty"`

	CardHolder string `json:"cardHolder,omitempty"`

	ReceiptSuggestions ReceiptSuggestions `json:"receiptSuggestions"`

	PartialAuth bool `json:"partialAuth"`

	AltCurrency bool `json:"altCurrency"`

	FSAAuth bool `json:"fasAuth"`

	CurrencyCode string `json:"currencyCode"`

	RequestedAmount string `json:"requestedAmount"`

	AuthorizedAmount string `json:"authorizedAmount"`

	RemainingBalance string `json:"remainingBalance"`

	TipAmount string `json:"tipAmount"`

	TaxAmount string `json:"taxAmount"`

	RequestedCashBackAmount string `json:"requestedCashBackAmount"`

	AuthorizedCashBackAmount string `json:"authorizedCashBackAmount"`

	WhiteListedCard *WhiteListedCard `json:"whiteListedCard"`

	StoreAndForward bool `json:"storeAndForward"`
}

// TransactionDisplayDiscount is an item level discount for transaction display. Discounts never combine.
type TransactionDisplayDiscount struct {
	Description string `json:"description"`

	Amount string `json:"amount"`
}

// TransactionDisplayItem is an item category in a transaction display. Groups combine if their descriptions match. Calculated subtotal amounts are rounded to two decimal places of precision. Quantity is a floating point number that is not rounded at all.
type TransactionDisplayItem struct {

	// ID is not required, but recommended since it is required to update or delete line items.
	ID string `json:"id"`

	Description string `json:"description"`

	Price string `json:"price"`

	Quantity float64 `json:"quantity"`

	// Extended is an item category in a transaction display. Groups combine if their descriptions match. Calculated subtotal amounts are rounded to two decimal places of precision. Quantity is a floating point number that is not rounded at all.
	Extended string `json:"extended"`

	// Discounts are displayed under their corresponding item.
	Discounts []*TransactionDisplayDiscount `json:"discounts"`
}

// TransactionDisplayTransaction contains the items to display on a terminal.
type TransactionDisplayTransaction struct {
	Subtotal string `json:"subtotal"`

	Tax string `json:"tax"`

	Total string `json:"total"`

	Quantity float64 `json:"quantity"`

	// Items can be overwritten or appended, based on the request type.
	Items []*TransactionDisplayItem `json:"items"`
}

// TransactionDisplayRequest is used to start or update a transaction line item display on a terminal.
type TransactionDisplayRequest struct {
	TransactionRef string `json:"transactionRef,omitempty"`

	OrderRef string `json:"orderRef,omitempty"`

	DestinationAccount string `json:"destinationAccount,omitempty"`

	Test bool `json:"test"`

	Timeout int `json:"timeout"`

	TerminalName string `json:"terminalName,omitempty"`

	// Transaction is the transaction to display on the terminal.
	Transaction *TransactionDisplayTransaction `json:"transaction"`
}

// HeartbeatResponse models the response to a basic API health check. If the security context permits it, the response may also include the public key of the current merchant.
type HeartbeatResponse struct {
	Success bool `json:"success"`

	Error string `json:"error"`

	ResponseDescription string `json:"responseDescription"`

	Timestamp time.Time `json:"timestamp"`

	Clockchain string `json:"clockchain"`

	LatestTick string `json:"latestTick"`

	MerchantPublicKey string `json:"merchantPk"`
}

// TerminalPingRequest models information needed to test connectivity with a terminal.
type TerminalPingRequest struct {
	APICredentials
	Request PingRequest `json:"request"`
}

// TerminalMessageRequest models a message to be displayed on the terminal screen.
type TerminalMessageRequest struct {
	APICredentials
	Request MessageRequest `json:"request"`
}

// TerminalBooleanPromptRequest models a simple yes no prompt request.
type TerminalBooleanPromptRequest struct {
	APICredentials
	Request BooleanPromptRequest `json:"request"`
}

// TerminalTextPromptRequest models a text prompt request.
type TerminalTextPromptRequest struct {
	APICredentials
	Request TextPromptRequest `json:"request"`
}

// TerminalAuthorizationRequest models auth requests for charge, preauth, and reverse transaction types.
type TerminalAuthorizationRequest struct {
	APICredentials
	Request AuthorizationRequest `json:"request"`
}

// TerminalBalanceRequest models balance requests.
type TerminalBalanceRequest struct {
	APICredentials
	Request BalanceRequest `json:"request"`
}

// TerminalRefundRequest models refund requests.
type TerminalRefundRequest struct {
	APICredentials
	Request RefundRequest `json:"request"`
}

// TerminalEnrollRequest models the information needed to enroll a new payment method in the token vault.
type TerminalEnrollRequest struct {
	APICredentials
	Request EnrollRequest `json:"request"`
}

// TerminalClearTerminalRequest models the information needed to enroll a new payment method in the token vault.
type TerminalClearTerminalRequest struct {
	APICredentials
	Request ClearTerminalRequest `json:"request"`
}

// TerminalGiftActivateRequest models the information needed to activate or recharge a gift card.
type TerminalGiftActivateRequest struct {
	APICredentials
	Request GiftActivateRequest `json:"request"`
}

// TerminalTermsAndConditionsRequest models the fields needed for custom T&C prompts.
type TerminalTermsAndConditionsRequest struct {
	APICredentials
	Request TermsAndConditionsRequest `json:"request"`
}

// TerminalTermsAndConditionsResponse models a T&C signature capture response.
type TerminalTermsAndConditionsResponse struct {
	APICredentials
	Request TermsAndConditionsResponse `json:"request"`
}

// TerminalTransactionDisplayRequest is used to start or update a transaction line item display on a terminal.
type TerminalTransactionDisplayRequest struct {
	APICredentials
	Request TransactionDisplayRequest `json:"request"`
}