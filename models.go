// Copyright 2019 BlockChyp, Inc. All rights reserved. Use of this code is
// governed by a license that can be found in the LICENSE file.
//
// This file was generated automatically. Changes to this file will be lost
// every time the code is regenerated.

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

// SignatureFormat is used to specify the output format for customer
// signature images.
type SignatureFormat string

// SignatureFormats.
const (
	SignatureFormatNone = ""
	SignatureFormatPNG  = "png"
	SignatureFormatJPG  = "jpg"
	SignatureFormatGIF  = "gif"
)

// PromptType is used to specify the type of text input data being requested
// from a customer.
type PromptType string

// PromptTypes.
const (
	PromptTypeAmount         = "amount"
	PromptTypeEmail          = "email"
	PromptTypePhone          = "phone"
	PromptTypeCustomerNumber = "customer-number"
	PromptTypeRewardsNumber  = "rewards-number"
)

// ReceiptSuggestions contains EMV fields we recommend developers put on their
// receipts.
type ReceiptSuggestions struct {
	// AID is the EMV Application Identifier.
	AID string `json:"aid,omitempty"`

	// ARQC is the EMV Application Request Cryptogram.
	ARQC string `json:"arqc,omitempty"`

	// IAD is the EMV Issuer Application Data.
	IAD string `json:"iad,omitempty"`

	// ARC is the EMV Authorization Response Code.
	ARC string `json:"arc,omitempty"`

	// TC is the EMV Transaction Certificate.
	TC string `json:"tc,omitempty"`

	// TVR is the EMV Terminal Verification Response.
	TVR string `json:"tvr,omitempty"`

	// TSI is the EMV Transaction Status Indicator.
	TSI string `json:"tsi,omitempty"`

	// TerminalID is the ID of the payment terminal.
	TerminalID string `json:"terminalId,omitempty"`

	// MerchantName is the name of the merchant's business.
	MerchantName string `json:"merchantName,omitempty"`

	// MerchantID is the ID of the merchant.
	MerchantID string `json:"merchantId,omitempty"`

	// MerchantKey is the partially masked merchant key required on EMV receipts.
	MerchantKey string `json:"merchantKey,omitempty"`

	// ApplicationLabel is a description of the selected AID.
	ApplicationLabel string `json:"applicationLabel,omitempty"`

	// RequestSignature indicates that the receipt should contain a signature
	// line.
	RequestSignature bool `json:"requestSignature"`

	// MaskedPAN is the masked primary account number of the payment card, as
	// required.
	MaskedPAN string `json:"maskedPan,omitempty"`

	// AuthorizedAmount is the amount authorized by the payment network. Could be
	// less than the requested amount for partial auth.
	AuthorizedAmount string `json:"authorizedAmount"`

	// TransactionType is the type of transaction performed (CHARGE, PREAUTH,
	// REFUND, etc).
	TransactionType string `json:"transactionType"`

	// EntryMethod is the method by which the payment card was entered (MSR,
	// CHIP, KEYED, etc.).
	EntryMethod string `json:"entryMethod,omitempty"`

	// PINVerified indicates that PIN verification was performed.
	PINVerified bool `json:"pinVerified,omitempty"`

	// CashBackAmount is the amount of cash back that was approved.
	CashBackAmount string `json:"cashBackAmount,omitempty"`
}

// Acknowledgement contains a basic api acknowledgement.
type Acknowledgement struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`
}

// PingRequest contains information needed to test connectivity with a
// terminal.
type PingRequest struct {
	// TransactionRef is the transaction reference string assigned to the
	// transaction request. If no transaction ref was assiged on the request,
	// then the gateway will randomly generate one.
	TransactionRef string `json:"transactionRef,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// Timeout is the request timeout in milliseconds.
	Timeout int `json:"timeout"`

	// TerminalName is the name of the target payment terminal.
	TerminalName string `json:"terminalName,omitempty"`
}

// PingResponse contains the response to a ping request.
type PingResponse struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// TransactionID is the ID assigned to the transaction.
	TransactionID string `json:"transactionId"`

	// BatchID is the ID assigned to the batch.
	BatchID string `json:"batchId,omitempty"`

	// TransactionRef is the transaction reference string assigned to the
	// transaction request. If no transaction ref was assiged on the request,
	// then the gateway will randomly generate one.
	TransactionRef string `json:"transactionRef,omitempty"`

	// TransactionType is the type of transaction.
	TransactionType string `json:"transactionType"`

	// Timestamp is the timestamp of the transaction.
	Timestamp string `json:"timestamp"`

	// TickBlock is the hash of the last tick block.
	TickBlock string `json:"tickBlock"`

	// Test indicates that the transaction was processed on the test gateway.
	Test bool `json:"test"`

	// Sig is the ECC signature of the response. Can be used to ensure that it
	// was signed by the terminal and detect man-in-the middle attacks.
	Sig string `json:"sig,omitempty"`
}

// MessageRequest contains a message to be displayed on the terminal screen.
type MessageRequest struct {
	// TransactionRef is the transaction reference string assigned to the
	// transaction request. If no transaction ref was assiged on the request,
	// then the gateway will randomly generate one.
	TransactionRef string `json:"transactionRef,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// Timeout is the request timeout in milliseconds.
	Timeout int `json:"timeout"`

	// TerminalName is the name of the target payment terminal.
	TerminalName string `json:"terminalName,omitempty"`

	// Message is the message to be displayed on the terminal.
	Message string `json:"message"`
}

// BooleanPromptRequest contains a simple yes no prompt request.
type BooleanPromptRequest struct {
	// TransactionRef is the transaction reference string assigned to the
	// transaction request. If no transaction ref was assiged on the request,
	// then the gateway will randomly generate one.
	TransactionRef string `json:"transactionRef,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// Timeout is the request timeout in milliseconds.
	Timeout int `json:"timeout"`

	// TerminalName is the name of the target payment terminal.
	TerminalName string `json:"terminalName,omitempty"`

	// YesCaption is the preferred caption for the 'yes' button.
	YesCaption string `json:"yesCaption"`

	// NoCaption is the preferred caption for the 'no' button.
	NoCaption string `json:"noCaption"`

	// Prompt is the text to be displayed on the terminal.
	Prompt string `json:"prompt"`
}

// TextPromptRequest contains a text prompt request.
type TextPromptRequest struct {
	// TransactionRef is the transaction reference string assigned to the
	// transaction request. If no transaction ref was assiged on the request,
	// then the gateway will randomly generate one.
	TransactionRef string `json:"transactionRef,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// Timeout is the request timeout in milliseconds.
	Timeout int `json:"timeout"`

	// TerminalName is the name of the target payment terminal.
	TerminalName string `json:"terminalName,omitempty"`

	// PromptType is the prompt type (email, phone, etc).
	PromptType PromptType `json:"promptType"`
}

// TextPromptResponse contains the response to a text prompt request.
type TextPromptResponse struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// Response is the text prompt response.
	Response string `json:"response"`
}

// BooleanPromptResponse contains the response to a boolean prompt request.
type BooleanPromptResponse struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// Response is the boolean prompt response.
	Response bool `json:"response"`
}

// WhiteListedCard shows details about a white listed card.
type WhiteListedCard struct {
	// Bin is the card BIN.
	Bin string `json:"bin"`

	// Track1 is the track 1 data from the card.
	Track1 string `json:"track1"`

	// Track2 is the track 2 data from the card.
	Track2 string `json:"track2"`

	// PAN is the card primary account number.
	PAN string `json:"pan"`
}

// AuthorizationRequest contains auth requests for charge, preauth, and
// reverse transaction types.
type AuthorizationRequest struct {
	// TransactionRef is the transaction reference string assigned to the
	// transaction request. If no transaction ref was assiged on the request,
	// then the gateway will randomly generate one.
	TransactionRef string `json:"transactionRef,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// Timeout is the request timeout in milliseconds.
	Timeout int `json:"timeout"`

	// TerminalName is the name of the target payment terminal.
	TerminalName string `json:"terminalName,omitempty"`

	// Token is the payment token to be used for this transaction. This should be
	// used for recurring transactions.
	Token string `json:"token,omitempty"`

	// Track1 contains track 1 magnetic stripe data.
	Track1 string `json:"track1,omitempty"`

	// Track2 contains track 2 magnetic stripe data.
	Track2 string `json:"track2,omitempty"`

	// PAN contains the primary account number. We recommend using the terminal
	// or e-commerce tokenization libraries instead of passing account numbers in
	// directly, as this would put your application in PCI scope.
	PAN string `json:"pan,omitempty"`

	// RoutingNumber is the ACH routing number for ACH transactions.
	RoutingNumber string `json:"routingNumber,omitempty"`

	// CardholderName is the cardholder name. Only required if the request
	// includes a primary account number or track data.
	CardholderName string `json:"cardholderName,omitempty"`

	// ExpMonth is the card expiration month for use with PAN based transactions.
	ExpMonth string `json:"expMonth,omitempty"`

	// ExpYear is the card expiration year for use with PAN based transactions.
	ExpYear string `json:"expYear,omitempty"`

	// CVV is the card CVV for use with PAN based transactions.
	CVV string `json:"cvv,omitempty"`

	// Address is the cardholder address for use with address verification.
	Address string `json:"address,omitempty"`

	// PostalCode is the cardholder postal code for use with address
	// verification.
	PostalCode string `json:"postalCode,omitempty"`

	// ManualEntry specifies that the payment entry method is a manual keyed
	// transaction. If this is true, no other payment method will be accepted.
	ManualEntry bool `json:"manualEntry,omitempty"`

	// KSN is the key serial number used for DUKPT encryption.
	KSN string `json:"ksn,omitempty"`

	// PINBlock is the encrypted pin block.
	PINBlock string `json:"pinBlock,omitempty"`

	// CardType designates categories of cards: credit, debit, EBT.
	CardType CardType `json:"cardType,omitempty"`

	// PaymentType designates brands of payment methods: Visa, Discover, etc.
	PaymentType string `json:"paymentType,omitempty"`

	// CurrencyCode indicates the transaction currency code.
	CurrencyCode string `json:"currencyCode"`

	// Amount is the requested amount.
	Amount string `json:"amount"`

	// TaxExempt indicates that the request is tax exempt. Only required for tax
	// exempt level 2 processing.
	TaxExempt bool `json:"taxExempt"`

	// TipAmount is the tip amount.
	TipAmount string `json:"tipAmount,omitempty"`

	// TaxAmount is the tax amount.
	TaxAmount string `json:"taxAmount,omitempty"`

	// CashBackAmount is the amount of cash back requested.
	CashBackAmount string `json:"cashBackAmount,omitempty"`

	// FSAEligibleAmount is the amount of the transaction that should be charged
	// to an FSA card. This amount may be less than the transaction total, in
	// which case only this amount will be charged if an FSA card is presented.
	// If the FSA amount is paid on an FSA card, then the FSA amount authorized
	// will be indicated on the response.
	FSAEligibleAmount string `json:"fsaEligibleAmount,omitempty"`

	// HSAEligibleAmount is the amount of the transaction that should be charged
	// to an HSA card.
	HSAEligibleAmount string `json:"hsaEligibleAmount,omitempty"`

	// EBTEligibleAmount is the amount of the transaction that should be charged
	// to an EBT card.
	EBTEligibleAmount string `json:"ebtEligibleAmount,omitempty"`

	// SigFile is a location on the filesystem which a customer signature should
	// be written to.
	SigFile string `json:"sigFile,omitempty"`

	// SigFormat specifies the image format to be used for returning signatures.
	SigFormat SignatureFormat `json:"sigFormat,omitempty"`

	// SigWidth is the width that the signature image should be scaled to,
	// preserving the aspect ratio. If not provided, the signature is returned in
	// the terminal's max resolution.
	SigWidth int `json:"sigWidth,omitempty"`

	// TransactionID is the ID of the previous transaction being referenced.
	TransactionID string `json:"transactionId"`

	// OnlineAuthCode is used to validate online gift card authorizations.
	OnlineAuthCode string `json:"onlineAuthCode,omitempty"`

	// Enroll indicates that the payment method should be added to the token
	// vault alongside the authorization.
	Enroll bool `json:"enroll,omitempty"`

	// Description contains a narrative description of the transaction.
	Description string `json:"description,omitempty"`

	// PromptForTip indicates that the terminal should request a tip from the
	// user before starting the transaction.
	PromptForTip bool `json:"promptForTip,omitempty"`

	// CashBackEnabled indicates that cash back should be enabled for supported
	// cards.
	CashBackEnabled bool `json:"cashBackEnabled,omitempty"`

	// AltPrices is a map of alternate currencies and the price in each currency.
	AltPrices map[string]string `json:"altPrices,omitempty"`
}

// BalanceRequest contains balance requests.
type BalanceRequest struct {
	// TransactionRef is the transaction reference string assigned to the
	// transaction request. If no transaction ref was assiged on the request,
	// then the gateway will randomly generate one.
	TransactionRef string `json:"transactionRef,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// Timeout is the request timeout in milliseconds.
	Timeout int `json:"timeout"`

	// Token is the payment token to be used for this transaction. This should be
	// used for recurring transactions.
	Token string `json:"token,omitempty"`

	// Track1 contains track 1 magnetic stripe data.
	Track1 string `json:"track1,omitempty"`

	// Track2 contains track 2 magnetic stripe data.
	Track2 string `json:"track2,omitempty"`

	// PAN contains the primary account number. We recommend using the terminal
	// or e-commerce tokenization libraries instead of passing account numbers in
	// directly, as this would put your application in PCI scope.
	PAN string `json:"pan,omitempty"`

	// RoutingNumber is the ACH routing number for ACH transactions.
	RoutingNumber string `json:"routingNumber,omitempty"`

	// CardholderName is the cardholder name. Only required if the request
	// includes a primary account number or track data.
	CardholderName string `json:"cardholderName,omitempty"`

	// ExpMonth is the card expiration month for use with PAN based transactions.
	ExpMonth string `json:"expMonth,omitempty"`

	// ExpYear is the card expiration year for use with PAN based transactions.
	ExpYear string `json:"expYear,omitempty"`

	// CVV is the card CVV for use with PAN based transactions.
	CVV string `json:"cvv,omitempty"`

	// Address is the cardholder address for use with address verification.
	Address string `json:"address,omitempty"`

	// PostalCode is the cardholder postal code for use with address
	// verification.
	PostalCode string `json:"postalCode,omitempty"`

	// ManualEntry specifies that the payment entry method is a manual keyed
	// transaction. If this is true, no other payment method will be accepted.
	ManualEntry bool `json:"manualEntry,omitempty"`

	// KSN is the key serial number used for DUKPT encryption.
	KSN string `json:"ksn,omitempty"`

	// PINBlock is the encrypted pin block.
	PINBlock string `json:"pinBlock,omitempty"`

	// CardType designates categories of cards: credit, debit, EBT.
	CardType CardType `json:"cardType,omitempty"`

	// PaymentType designates brands of payment methods: Visa, Discover, etc.
	PaymentType string `json:"paymentType,omitempty"`

	// TerminalName is the name of the target payment terminal.
	TerminalName string `json:"terminalName,omitempty"`
}

// BalanceResponse contains the response to a balance request.
type BalanceResponse struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// TransactionID is the ID assigned to the transaction.
	TransactionID string `json:"transactionId"`

	// BatchID is the ID assigned to the batch.
	BatchID string `json:"batchId,omitempty"`

	// TransactionRef is the transaction reference string assigned to the
	// transaction request. If no transaction ref was assiged on the request,
	// then the gateway will randomly generate one.
	TransactionRef string `json:"transactionRef,omitempty"`

	// TransactionType is the type of transaction.
	TransactionType string `json:"transactionType"`

	// Timestamp is the timestamp of the transaction.
	Timestamp string `json:"timestamp"`

	// TickBlock is the hash of the last tick block.
	TickBlock string `json:"tickBlock"`

	// Test indicates that the transaction was processed on the test gateway.
	Test bool `json:"test"`

	// Sig is the ECC signature of the response. Can be used to ensure that it
	// was signed by the terminal and detect man-in-the middle attacks.
	Sig string `json:"sig,omitempty"`

	// Token is the payment token, if the payment was enrolled in the vault.
	Token string `json:"token,omitempty"`

	// EntryMethod is the entry method for the transaction (CHIP, MSR, KEYED,
	// etc).
	EntryMethod string `json:"entryMethod,omitempty"`

	// PaymentType is the card brand (VISA, MC, AMEX, etc).
	PaymentType string `json:"paymentType,omitempty"`

	// MaskedPAN is the masked primary account number.
	MaskedPAN string `json:"maskedPan,omitempty"`

	// PublicKey is the BlockChyp public key if the user presented a BlockChyp
	// payment card.
	PublicKey string `json:"publicKey,omitempty"`

	// ScopeAlert indicates that the transaction did something that would put the
	// system in PCI scope.
	ScopeAlert bool `json:"ScopeAlert,omitempty"`

	// CardHolder is the cardholder name.
	CardHolder string `json:"cardHolder,omitempty"`

	// ReceiptSuggestions contains suggested receipt fields.
	ReceiptSuggestions ReceiptSuggestions `json:"receiptSuggestions"`

	// RemainingBalance remaining balance on the payment method.
	RemainingBalance string `json:"remainingBalance,omitempty"`
}

// RefundRequest contains refund requests.
type RefundRequest struct {
	// TransactionRef is the transaction reference string assigned to the
	// transaction request. If no transaction ref was assiged on the request,
	// then the gateway will randomly generate one.
	TransactionRef string `json:"transactionRef,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// Timeout is the request timeout in milliseconds.
	Timeout int `json:"timeout"`

	// TerminalName is the name of the target payment terminal.
	TerminalName string `json:"terminalName,omitempty"`

	// Token is the payment token to be used for this transaction. This should be
	// used for recurring transactions.
	Token string `json:"token,omitempty"`

	// Track1 contains track 1 magnetic stripe data.
	Track1 string `json:"track1,omitempty"`

	// Track2 contains track 2 magnetic stripe data.
	Track2 string `json:"track2,omitempty"`

	// PAN contains the primary account number. We recommend using the terminal
	// or e-commerce tokenization libraries instead of passing account numbers in
	// directly, as this would put your application in PCI scope.
	PAN string `json:"pan,omitempty"`

	// RoutingNumber is the ACH routing number for ACH transactions.
	RoutingNumber string `json:"routingNumber,omitempty"`

	// CardholderName is the cardholder name. Only required if the request
	// includes a primary account number or track data.
	CardholderName string `json:"cardholderName,omitempty"`

	// ExpMonth is the card expiration month for use with PAN based transactions.
	ExpMonth string `json:"expMonth,omitempty"`

	// ExpYear is the card expiration year for use with PAN based transactions.
	ExpYear string `json:"expYear,omitempty"`

	// CVV is the card CVV for use with PAN based transactions.
	CVV string `json:"cvv,omitempty"`

	// Address is the cardholder address for use with address verification.
	Address string `json:"address,omitempty"`

	// PostalCode is the cardholder postal code for use with address
	// verification.
	PostalCode string `json:"postalCode,omitempty"`

	// ManualEntry specifies that the payment entry method is a manual keyed
	// transaction. If this is true, no other payment method will be accepted.
	ManualEntry bool `json:"manualEntry,omitempty"`

	// KSN is the key serial number used for DUKPT encryption.
	KSN string `json:"ksn,omitempty"`

	// PINBlock is the encrypted pin block.
	PINBlock string `json:"pinBlock,omitempty"`

	// CardType designates categories of cards: credit, debit, EBT.
	CardType CardType `json:"cardType,omitempty"`

	// PaymentType designates brands of payment methods: Visa, Discover, etc.
	PaymentType string `json:"paymentType,omitempty"`

	// CurrencyCode indicates the transaction currency code.
	CurrencyCode string `json:"currencyCode"`

	// Amount is the requested amount.
	Amount string `json:"amount"`

	// TaxExempt indicates that the request is tax exempt. Only required for tax
	// exempt level 2 processing.
	TaxExempt bool `json:"taxExempt"`

	// TipAmount is the tip amount.
	TipAmount string `json:"tipAmount,omitempty"`

	// TaxAmount is the tax amount.
	TaxAmount string `json:"taxAmount,omitempty"`

	// CashBackAmount is the amount of cash back requested.
	CashBackAmount string `json:"cashBackAmount,omitempty"`

	// FSAEligibleAmount is the amount of the transaction that should be charged
	// to an FSA card. This amount may be less than the transaction total, in
	// which case only this amount will be charged if an FSA card is presented.
	// If the FSA amount is paid on an FSA card, then the FSA amount authorized
	// will be indicated on the response.
	FSAEligibleAmount string `json:"fsaEligibleAmount,omitempty"`

	// HSAEligibleAmount is the amount of the transaction that should be charged
	// to an HSA card.
	HSAEligibleAmount string `json:"hsaEligibleAmount,omitempty"`

	// EBTEligibleAmount is the amount of the transaction that should be charged
	// to an EBT card.
	EBTEligibleAmount string `json:"ebtEligibleAmount,omitempty"`

	// SigFile is a location on the filesystem which a customer signature should
	// be written to.
	SigFile string `json:"sigFile,omitempty"`

	// SigFormat specifies the image format to be used for returning signatures.
	SigFormat SignatureFormat `json:"sigFormat,omitempty"`

	// SigWidth is the width that the signature image should be scaled to,
	// preserving the aspect ratio. If not provided, the signature is returned in
	// the terminal's max resolution.
	SigWidth int `json:"sigWidth,omitempty"`

	// TransactionID is the ID of the previous transaction being referenced.
	TransactionID string `json:"transactionId"`
}

// CaptureRequest contains the information needed to capture a preauth.
type CaptureRequest struct {
	// TransactionRef is the transaction reference string assigned to the
	// transaction request. If no transaction ref was assiged on the request,
	// then the gateway will randomly generate one.
	TransactionRef string `json:"transactionRef,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// Timeout is the request timeout in milliseconds.
	Timeout int `json:"timeout"`

	// CurrencyCode indicates the transaction currency code.
	CurrencyCode string `json:"currencyCode"`

	// Amount is the requested amount.
	Amount string `json:"amount"`

	// TaxExempt indicates that the request is tax exempt. Only required for tax
	// exempt level 2 processing.
	TaxExempt bool `json:"taxExempt"`

	// TipAmount is the tip amount.
	TipAmount string `json:"tipAmount,omitempty"`

	// TaxAmount is the tax amount.
	TaxAmount string `json:"taxAmount,omitempty"`

	// CashBackAmount is the amount of cash back requested.
	CashBackAmount string `json:"cashBackAmount,omitempty"`

	// FSAEligibleAmount is the amount of the transaction that should be charged
	// to an FSA card. This amount may be less than the transaction total, in
	// which case only this amount will be charged if an FSA card is presented.
	// If the FSA amount is paid on an FSA card, then the FSA amount authorized
	// will be indicated on the response.
	FSAEligibleAmount string `json:"fsaEligibleAmount,omitempty"`

	// HSAEligibleAmount is the amount of the transaction that should be charged
	// to an HSA card.
	HSAEligibleAmount string `json:"hsaEligibleAmount,omitempty"`

	// EBTEligibleAmount is the amount of the transaction that should be charged
	// to an EBT card.
	EBTEligibleAmount string `json:"ebtEligibleAmount,omitempty"`

	// TransactionID is the ID of the previous transaction being referenced.
	TransactionID string `json:"transactionId"`
}

// CaptureResponse contains the response to a capture request.
type CaptureResponse struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// TransactionID is the ID assigned to the transaction.
	TransactionID string `json:"transactionId"`

	// BatchID is the ID assigned to the batch.
	BatchID string `json:"batchId,omitempty"`

	// TransactionRef is the transaction reference string assigned to the
	// transaction request. If no transaction ref was assiged on the request,
	// then the gateway will randomly generate one.
	TransactionRef string `json:"transactionRef,omitempty"`

	// TransactionType is the type of transaction.
	TransactionType string `json:"transactionType"`

	// Timestamp is the timestamp of the transaction.
	Timestamp string `json:"timestamp"`

	// TickBlock is the hash of the last tick block.
	TickBlock string `json:"tickBlock"`

	// Test indicates that the transaction was processed on the test gateway.
	Test bool `json:"test"`

	// Sig is the ECC signature of the response. Can be used to ensure that it
	// was signed by the terminal and detect man-in-the middle attacks.
	Sig string `json:"sig,omitempty"`

	// Approved indicates that the transaction was approved.
	Approved bool `json:"approved"`

	// AuthCode is the auth code from the payment network.
	AuthCode string `json:"authCode,omitempty"`

	// SigFile contains the hex encoded signature data.
	SigFile string `json:"sigFile"`

	// Token is the payment token, if the payment was enrolled in the vault.
	Token string `json:"token,omitempty"`

	// EntryMethod is the entry method for the transaction (CHIP, MSR, KEYED,
	// etc).
	EntryMethod string `json:"entryMethod,omitempty"`

	// PaymentType is the card brand (VISA, MC, AMEX, etc).
	PaymentType string `json:"paymentType,omitempty"`

	// MaskedPAN is the masked primary account number.
	MaskedPAN string `json:"maskedPan,omitempty"`

	// PublicKey is the BlockChyp public key if the user presented a BlockChyp
	// payment card.
	PublicKey string `json:"publicKey,omitempty"`

	// ScopeAlert indicates that the transaction did something that would put the
	// system in PCI scope.
	ScopeAlert bool `json:"ScopeAlert,omitempty"`

	// CardHolder is the cardholder name.
	CardHolder string `json:"cardHolder,omitempty"`

	// ReceiptSuggestions contains suggested receipt fields.
	ReceiptSuggestions ReceiptSuggestions `json:"receiptSuggestions"`

	// PartialAuth indicates whether or not the transaction was approved for a
	// partial amount.
	PartialAuth bool `json:"partialAuth"`

	// AltCurrency indicates whether or not an alternate currency was used.
	AltCurrency bool `json:"altCurrency"`

	// FSAAuth indicates whether or not a request was settled on an FSA card.
	FSAAuth bool `json:"fsaAuth"`

	// CurrencyCode is the currency code used for the transaction.
	CurrencyCode string `json:"currencyCode"`

	// RequestedAmount is the requested amount.
	RequestedAmount string `json:"requestedAmount"`

	// AuthorizedAmount is the authorized amount. May not match the requested
	// amount in the event of a partial auth.
	AuthorizedAmount string `json:"authorizedAmount"`

	// RemainingBalance is the remaining balance on the payment method.
	RemainingBalance string `json:"remainingBalance"`

	// TipAmount is the tip amount.
	TipAmount string `json:"tipAmount"`

	// TaxAmount is the tax amount.
	TaxAmount string `json:"taxAmount"`

	// RequestedCashBackAmount is the cash back amount the customer requested
	// during the transaction.
	RequestedCashBackAmount string `json:"requestedCashBackAmount"`

	// AuthorizedCashBackAmount is the amount of cash back authorized by the
	// gateway. This amount will be the entire amount requested, or zero.
	AuthorizedCashBackAmount string `json:"authorizedCashBackAmount"`
}

// VoidRequest contains a void request.
type VoidRequest struct {
	// TransactionRef is the transaction reference string assigned to the
	// transaction request. If no transaction ref was assiged on the request,
	// then the gateway will randomly generate one.
	TransactionRef string `json:"transactionRef,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// Timeout is the request timeout in milliseconds.
	Timeout int `json:"timeout"`

	// TransactionID is the ID of the previous transaction being referenced.
	TransactionID string `json:"transactionId"`
}

// VoidResponse contains the response to a void request.
type VoidResponse struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// TransactionID is the ID assigned to the transaction.
	TransactionID string `json:"transactionId"`

	// BatchID is the ID assigned to the batch.
	BatchID string `json:"batchId,omitempty"`

	// TransactionRef is the transaction reference string assigned to the
	// transaction request. If no transaction ref was assiged on the request,
	// then the gateway will randomly generate one.
	TransactionRef string `json:"transactionRef,omitempty"`

	// TransactionType is the type of transaction.
	TransactionType string `json:"transactionType"`

	// Timestamp is the timestamp of the transaction.
	Timestamp string `json:"timestamp"`

	// TickBlock is the hash of the last tick block.
	TickBlock string `json:"tickBlock"`

	// Test indicates that the transaction was processed on the test gateway.
	Test bool `json:"test"`

	// Sig is the ECC signature of the response. Can be used to ensure that it
	// was signed by the terminal and detect man-in-the middle attacks.
	Sig string `json:"sig,omitempty"`

	// Approved indicates that the transaction was approved.
	Approved bool `json:"approved"`

	// AuthCode is the auth code from the payment network.
	AuthCode string `json:"authCode,omitempty"`

	// SigFile contains the hex encoded signature data.
	SigFile string `json:"sigFile"`

	// Token is the payment token, if the payment was enrolled in the vault.
	Token string `json:"token,omitempty"`

	// EntryMethod is the entry method for the transaction (CHIP, MSR, KEYED,
	// etc).
	EntryMethod string `json:"entryMethod,omitempty"`

	// PaymentType is the card brand (VISA, MC, AMEX, etc).
	PaymentType string `json:"paymentType,omitempty"`

	// MaskedPAN is the masked primary account number.
	MaskedPAN string `json:"maskedPan,omitempty"`

	// PublicKey is the BlockChyp public key if the user presented a BlockChyp
	// payment card.
	PublicKey string `json:"publicKey,omitempty"`

	// ScopeAlert indicates that the transaction did something that would put the
	// system in PCI scope.
	ScopeAlert bool `json:"ScopeAlert,omitempty"`

	// CardHolder is the cardholder name.
	CardHolder string `json:"cardHolder,omitempty"`

	// ReceiptSuggestions contains suggested receipt fields.
	ReceiptSuggestions ReceiptSuggestions `json:"receiptSuggestions"`
}

// EnrollRequest contains the information needed to enroll a new payment
// method in the token vault.
type EnrollRequest struct {
	// TransactionRef is the transaction reference string assigned to the
	// transaction request. If no transaction ref was assiged on the request,
	// then the gateway will randomly generate one.
	TransactionRef string `json:"transactionRef,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// Timeout is the request timeout in milliseconds.
	Timeout int `json:"timeout"`

	// TerminalName is the name of the target payment terminal.
	TerminalName string `json:"terminalName,omitempty"`

	// Token is the payment token to be used for this transaction. This should be
	// used for recurring transactions.
	Token string `json:"token,omitempty"`

	// Track1 contains track 1 magnetic stripe data.
	Track1 string `json:"track1,omitempty"`

	// Track2 contains track 2 magnetic stripe data.
	Track2 string `json:"track2,omitempty"`

	// PAN contains the primary account number. We recommend using the terminal
	// or e-commerce tokenization libraries instead of passing account numbers in
	// directly, as this would put your application in PCI scope.
	PAN string `json:"pan,omitempty"`

	// RoutingNumber is the ACH routing number for ACH transactions.
	RoutingNumber string `json:"routingNumber,omitempty"`

	// CardholderName is the cardholder name. Only required if the request
	// includes a primary account number or track data.
	CardholderName string `json:"cardholderName,omitempty"`

	// ExpMonth is the card expiration month for use with PAN based transactions.
	ExpMonth string `json:"expMonth,omitempty"`

	// ExpYear is the card expiration year for use with PAN based transactions.
	ExpYear string `json:"expYear,omitempty"`

	// CVV is the card CVV for use with PAN based transactions.
	CVV string `json:"cvv,omitempty"`

	// Address is the cardholder address for use with address verification.
	Address string `json:"address,omitempty"`

	// PostalCode is the cardholder postal code for use with address
	// verification.
	PostalCode string `json:"postalCode,omitempty"`

	// ManualEntry specifies that the payment entry method is a manual keyed
	// transaction. If this is true, no other payment method will be accepted.
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

// EnrollResponse contains the response to an enroll request.
type EnrollResponse struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// TransactionID is the ID assigned to the transaction.
	TransactionID string `json:"transactionId"`

	// BatchID is the ID assigned to the batch.
	BatchID string `json:"batchId,omitempty"`

	// TransactionRef is the transaction reference string assigned to the
	// transaction request. If no transaction ref was assiged on the request,
	// then the gateway will randomly generate one.
	TransactionRef string `json:"transactionRef,omitempty"`

	// TransactionType is the type of transaction.
	TransactionType string `json:"transactionType"`

	// Timestamp is the timestamp of the transaction.
	Timestamp string `json:"timestamp"`

	// TickBlock is the hash of the last tick block.
	TickBlock string `json:"tickBlock"`

	// Test indicates that the transaction was processed on the test gateway.
	Test bool `json:"test"`

	// Sig is the ECC signature of the response. Can be used to ensure that it
	// was signed by the terminal and detect man-in-the middle attacks.
	Sig string `json:"sig,omitempty"`

	// Token is the payment token, if the payment was enrolled in the vault.
	Token string `json:"token,omitempty"`

	// EntryMethod is the entry method for the transaction (CHIP, MSR, KEYED,
	// etc).
	EntryMethod string `json:"entryMethod,omitempty"`

	// PaymentType is the card brand (VISA, MC, AMEX, etc).
	PaymentType string `json:"paymentType,omitempty"`

	// MaskedPAN is the masked primary account number.
	MaskedPAN string `json:"maskedPan,omitempty"`

	// PublicKey is the BlockChyp public key if the user presented a BlockChyp
	// payment card.
	PublicKey string `json:"publicKey,omitempty"`

	// ScopeAlert indicates that the transaction did something that would put the
	// system in PCI scope.
	ScopeAlert bool `json:"ScopeAlert,omitempty"`

	// CardHolder is the cardholder name.
	CardHolder string `json:"cardHolder,omitempty"`

	// ReceiptSuggestions contains suggested receipt fields.
	ReceiptSuggestions ReceiptSuggestions `json:"receiptSuggestions"`

	// Approved indicates that the transaction was approved.
	Approved bool `json:"approved"`

	// AuthCode is the auth code from the payment network.
	AuthCode string `json:"authCode,omitempty"`

	// SigFile contains the hex encoded signature data.
	SigFile string `json:"sigFile"`
}

// ClearTerminalRequest contains the information needed to enroll a new
// payment method in the token vault.
type ClearTerminalRequest struct {
	// TransactionRef is the transaction reference string assigned to the
	// transaction request. If no transaction ref was assiged on the request,
	// then the gateway will randomly generate one.
	TransactionRef string `json:"transactionRef,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// Timeout is the request timeout in milliseconds.
	Timeout int `json:"timeout"`

	// TerminalName is the name of the target payment terminal.
	TerminalName string `json:"terminalName,omitempty"`
}

// GiftActivateRequest contains the information needed to activate or recharge
// a gift card.
type GiftActivateRequest struct {
	// TransactionRef is the transaction reference string assigned to the
	// transaction request. If no transaction ref was assiged on the request,
	// then the gateway will randomly generate one.
	TransactionRef string `json:"transactionRef,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// Timeout is the request timeout in milliseconds.
	Timeout int `json:"timeout"`

	// CurrencyCode indicates the transaction currency code.
	CurrencyCode string `json:"currencyCode"`

	// Amount is the requested amount.
	Amount string `json:"amount"`

	// TaxExempt indicates that the request is tax exempt. Only required for tax
	// exempt level 2 processing.
	TaxExempt bool `json:"taxExempt"`

	// TerminalName is the name of the target payment terminal.
	TerminalName string `json:"terminalName,omitempty"`
}

// GiftActivateResponse contains the response to a gift activate request.
type GiftActivateResponse struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// TransactionID is the ID assigned to the transaction.
	TransactionID string `json:"transactionId"`

	// BatchID is the ID assigned to the batch.
	BatchID string `json:"batchId,omitempty"`

	// TransactionRef is the transaction reference string assigned to the
	// transaction request. If no transaction ref was assiged on the request,
	// then the gateway will randomly generate one.
	TransactionRef string `json:"transactionRef,omitempty"`

	// TransactionType is the type of transaction.
	TransactionType string `json:"transactionType"`

	// Timestamp is the timestamp of the transaction.
	Timestamp string `json:"timestamp"`

	// TickBlock is the hash of the last tick block.
	TickBlock string `json:"tickBlock"`

	// Test indicates that the transaction was processed on the test gateway.
	Test bool `json:"test"`

	// Sig is the ECC signature of the response. Can be used to ensure that it
	// was signed by the terminal and detect man-in-the middle attacks.
	Sig string `json:"sig,omitempty"`

	// Approved indicates that the card was activated.
	Approved bool `json:"approved"`

	// Amount is the amount of the transaction.
	Amount string `json:"amount"`

	// CurrentBalance is the current balance of the gift card.
	CurrentBalance string `json:"currentBalance"`

	// CurrencyCode is the currency code used for the transaction.
	CurrencyCode string `json:"currencyCode"`

	// PublicKey is the public key of the activated card.
	PublicKey string `json:"publicKey"`
}

// CloseBatchRequest contains the information needed to manually close a
// credit card batch.
type CloseBatchRequest struct {
	// TransactionRef is the transaction reference string assigned to the
	// transaction request. If no transaction ref was assiged on the request,
	// then the gateway will randomly generate one.
	TransactionRef string `json:"transactionRef,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// Timeout is the request timeout in milliseconds.
	Timeout int `json:"timeout"`
}

// CloseBatchResponse contains the response to a close batch request.
type CloseBatchResponse struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// TransactionID is the ID assigned to the transaction.
	TransactionID string `json:"transactionId"`

	// BatchID is the ID assigned to the batch.
	BatchID string `json:"batchId,omitempty"`

	// TransactionRef is the transaction reference string assigned to the
	// transaction request. If no transaction ref was assiged on the request,
	// then the gateway will randomly generate one.
	TransactionRef string `json:"transactionRef,omitempty"`

	// TransactionType is the type of transaction.
	TransactionType string `json:"transactionType"`

	// Timestamp is the timestamp of the transaction.
	Timestamp string `json:"timestamp"`

	// TickBlock is the hash of the last tick block.
	TickBlock string `json:"tickBlock"`

	// Test indicates that the transaction was processed on the test gateway.
	Test bool `json:"test"`

	// Sig is the ECC signature of the response. Can be used to ensure that it
	// was signed by the terminal and detect man-in-the middle attacks.
	Sig string `json:"sig,omitempty"`

	// CurrencyCode is the currency code of amounts indicated.
	CurrencyCode string `json:"currencyCode"`

	// CapturedTotal is the total captured amount for this batch. Should be the
	// expected deposit amount.
	CapturedTotal string `json:"capturedTotal"`

	// OpenPreauths contains the total amount of preauths opened during the batch
	// that weren't captured.
	OpenPreauths string `json:"openPreauths"`

	// CardBrands contains the captured totals by card brand.
	CardBrands map[string]string `json:"cardBrands"`
}

// TermsAndConditionsRequest contains the fields needed for custom T&C
// prompts.
type TermsAndConditionsRequest struct {
	// TransactionRef is the transaction reference string assigned to the
	// transaction request. If no transaction ref was assiged on the request,
	// then the gateway will randomly generate one.
	TransactionRef string `json:"transactionRef,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// Timeout is the request timeout in milliseconds.
	Timeout int `json:"timeout"`

	// TerminalName is the name of the target payment terminal.
	TerminalName string `json:"terminalName,omitempty"`

	// TransactionID is the ID of the previous transaction being referenced.
	TransactionID string `json:"transactionId"`

	// SigFile is a location on the filesystem which a customer signature should
	// be written to.
	SigFile string `json:"sigFile,omitempty"`

	// SigFormat specifies the image format to be used for returning signatures.
	SigFormat SignatureFormat `json:"sigFormat,omitempty"`

	// SigWidth is the width that the signature image should be scaled to,
	// preserving the aspect ratio. If not provided, the signature is returned in
	// the terminal's max resolution.
	SigWidth int `json:"sigWidth,omitempty"`

	// TCAlias is a reference to a T&C assembled in the dashboard.
	TCAlias string `json:"tcAlias"`

	// TCName contains the name of the T&Cs the user is accepting.
	TCName string `json:"tcName"`

	// TCContent is the content of the terms and conditions that will be
	// presented to the user.
	TCContent string `json:"tcContent"`

	// SigRequired indicates that a signature should be requested.
	SigRequired bool `json:"sigRequired"`
}

// TermsAndConditionsResponse contains a T&C signature capture response.
type TermsAndConditionsResponse struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// TransactionID is the ID assigned to the transaction.
	TransactionID string `json:"transactionId"`

	// BatchID is the ID assigned to the batch.
	BatchID string `json:"batchId,omitempty"`

	// TransactionRef is the transaction reference string assigned to the
	// transaction request. If no transaction ref was assiged on the request,
	// then the gateway will randomly generate one.
	TransactionRef string `json:"transactionRef,omitempty"`

	// TransactionType is the type of transaction.
	TransactionType string `json:"transactionType"`

	// Timestamp is the timestamp of the transaction.
	Timestamp string `json:"timestamp"`

	// TickBlock is the hash of the last tick block.
	TickBlock string `json:"tickBlock"`

	// Test indicates that the transaction was processed on the test gateway.
	Test bool `json:"test"`

	// Sig is the ECC signature of the response. Can be used to ensure that it
	// was signed by the terminal and detect man-in-the middle attacks.
	Sig string `json:"sig,omitempty"`

	// SigFile contains the hex encoded signature data.
	SigFile string `json:"sigFile,omitempty"`
}

// AuthorizationResponse contains the response to authorization requests.
type AuthorizationResponse struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// TransactionID is the ID assigned to the transaction.
	TransactionID string `json:"transactionId"`

	// BatchID is the ID assigned to the batch.
	BatchID string `json:"batchId,omitempty"`

	// TransactionRef is the transaction reference string assigned to the
	// transaction request. If no transaction ref was assiged on the request,
	// then the gateway will randomly generate one.
	TransactionRef string `json:"transactionRef,omitempty"`

	// TransactionType is the type of transaction.
	TransactionType string `json:"transactionType"`

	// Timestamp is the timestamp of the transaction.
	Timestamp string `json:"timestamp"`

	// TickBlock is the hash of the last tick block.
	TickBlock string `json:"tickBlock"`

	// Test indicates that the transaction was processed on the test gateway.
	Test bool `json:"test"`

	// Sig is the ECC signature of the response. Can be used to ensure that it
	// was signed by the terminal and detect man-in-the middle attacks.
	Sig string `json:"sig,omitempty"`

	// Approved indicates that the transaction was approved.
	Approved bool `json:"approved"`

	// AuthCode is the auth code from the payment network.
	AuthCode string `json:"authCode,omitempty"`

	// SigFile contains the hex encoded signature data.
	SigFile string `json:"sigFile"`

	// Token is the payment token, if the payment was enrolled in the vault.
	Token string `json:"token,omitempty"`

	// EntryMethod is the entry method for the transaction (CHIP, MSR, KEYED,
	// etc).
	EntryMethod string `json:"entryMethod,omitempty"`

	// PaymentType is the card brand (VISA, MC, AMEX, etc).
	PaymentType string `json:"paymentType,omitempty"`

	// MaskedPAN is the masked primary account number.
	MaskedPAN string `json:"maskedPan,omitempty"`

	// PublicKey is the BlockChyp public key if the user presented a BlockChyp
	// payment card.
	PublicKey string `json:"publicKey,omitempty"`

	// ScopeAlert indicates that the transaction did something that would put the
	// system in PCI scope.
	ScopeAlert bool `json:"ScopeAlert,omitempty"`

	// CardHolder is the cardholder name.
	CardHolder string `json:"cardHolder,omitempty"`

	// ReceiptSuggestions contains suggested receipt fields.
	ReceiptSuggestions ReceiptSuggestions `json:"receiptSuggestions"`

	// PartialAuth indicates whether or not the transaction was approved for a
	// partial amount.
	PartialAuth bool `json:"partialAuth"`

	// AltCurrency indicates whether or not an alternate currency was used.
	AltCurrency bool `json:"altCurrency"`

	// FSAAuth indicates whether or not a request was settled on an FSA card.
	FSAAuth bool `json:"fsaAuth"`

	// CurrencyCode is the currency code used for the transaction.
	CurrencyCode string `json:"currencyCode"`

	// RequestedAmount is the requested amount.
	RequestedAmount string `json:"requestedAmount"`

	// AuthorizedAmount is the authorized amount. May not match the requested
	// amount in the event of a partial auth.
	AuthorizedAmount string `json:"authorizedAmount"`

	// RemainingBalance is the remaining balance on the payment method.
	RemainingBalance string `json:"remainingBalance"`

	// TipAmount is the tip amount.
	TipAmount string `json:"tipAmount"`

	// TaxAmount is the tax amount.
	TaxAmount string `json:"taxAmount"`

	// RequestedCashBackAmount is the cash back amount the customer requested
	// during the transaction.
	RequestedCashBackAmount string `json:"requestedCashBackAmount"`

	// AuthorizedCashBackAmount is the amount of cash back authorized by the
	// gateway. This amount will be the entire amount requested, or zero.
	AuthorizedCashBackAmount string `json:"authorizedCashBackAmount"`

	// WhiteListedCard contains card BIN ranges can be whitelisted so that they
	// are read instead of being processed directly. This is useful for
	// integration with legacy gift card systems.
	WhiteListedCard *WhiteListedCard `json:"whiteListedCard"`

	// StoreAndForward indicates that the transaction was flagged for store and
	// forward due to network problems.
	StoreAndForward bool `json:"storeAndForward"`
}

// TransactionDisplayDiscount is an item level discount for transaction
// display. Discounts never combine.
type TransactionDisplayDiscount struct {
	// Description is the discount description.
	Description string `json:"description"`

	// Amount is the amount of the discount.
	Amount string `json:"amount"`
}

// TransactionDisplayItem is an item category in a transaction display. Groups
// combine if their descriptions match. Calculated subtotal amounts are
// rounded to two decimal places of precision. Quantity is a floating point
// number that is not rounded at all.
type TransactionDisplayItem struct {
	// ID is a unique value identifying the item. This is not required, but
	// recommended since it is required to update or delete line items.
	ID string `json:"id"`

	// Description is a description of the line item.
	Description string `json:"description"`

	// Price is the price of the line item.
	Price string `json:"price"`

	// Quantity is the quantity of the line item.
	Quantity float64 `json:"quantity"`

	// Extended is an item category in a transaction display. Groups combine if
	// their descriptions match. Calculated subtotal amounts are rounded to two
	// decimal places of precision. Quantity is a floating point number that is
	// not rounded at all.
	Extended string `json:"extended"`

	// Discounts are displayed under their corresponding item.
	Discounts []*TransactionDisplayDiscount `json:"discounts"`
}

// TransactionDisplayTransaction contains the items to display on a terminal.
type TransactionDisplayTransaction struct {
	// Subtotal is the subtotal to display.
	Subtotal string `json:"subtotal"`

	// Tax is the tax to display.
	Tax string `json:"tax"`

	// Total is the total to display.
	Total string `json:"total"`

	// Items is an item to display. Can be overwritten or appended, based on the
	// request type.
	Items []*TransactionDisplayItem `json:"items"`
}

// TransactionDisplayRequest is used to start or update a transaction line
// item display on a terminal.
type TransactionDisplayRequest struct {
	// TransactionRef is the transaction reference string assigned to the
	// transaction request. If no transaction ref was assiged on the request,
	// then the gateway will randomly generate one.
	TransactionRef string `json:"transactionRef,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// Timeout is the request timeout in milliseconds.
	Timeout int `json:"timeout"`

	// TerminalName is the name of the target payment terminal.
	TerminalName string `json:"terminalName,omitempty"`

	// Transaction transaction to display on the terminal.
	Transaction *TransactionDisplayTransaction `json:"transaction"`
}

// HeartbeatResponse contains the response to a basic API health check. If the
// security context permits it, the response may also include the public key
// of the current merchant.
type HeartbeatResponse struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// Timestamp is the timestamp of the heartbeat.
	Timestamp time.Time `json:"timestamp"`

	// Clockchain is the public key of the clockchain. This is blockchain stuff
	// that you don't really need to worry about. It is a base 58 encoded and
	// compressed eliptic curve public key. For the production clockchain, this
	// will always be: '3cuhsckVUd9HzMjbdUSW17aY5kCcm1d6YAphJMUwmtXRj7WLyU'.
	Clockchain string `json:"clockchain"`

	// LatestTick is the hash of the last tick block.
	LatestTick string `json:"latestTick"`

	// MerchantPublicKey is the public key for the merchant's blockchain.
	MerchantPublicKey string `json:"merchantPk"`
}

// TerminalPingRequest contains information needed to test connectivity with a
// terminal.
type TerminalPingRequest struct {
	APICredentials
	Request PingRequest `json:"request"`
}

// TerminalMessageRequest contains a message to be displayed on the terminal
// screen.
type TerminalMessageRequest struct {
	APICredentials
	Request MessageRequest `json:"request"`
}

// TerminalBooleanPromptRequest contains a simple yes no prompt request.
type TerminalBooleanPromptRequest struct {
	APICredentials
	Request BooleanPromptRequest `json:"request"`
}

// TerminalTextPromptRequest contains a text prompt request.
type TerminalTextPromptRequest struct {
	APICredentials
	Request TextPromptRequest `json:"request"`
}

// TerminalAuthorizationRequest contains auth requests for charge, preauth,
// and reverse transaction types.
type TerminalAuthorizationRequest struct {
	APICredentials
	Request AuthorizationRequest `json:"request"`
}

// TerminalBalanceRequest contains balance requests.
type TerminalBalanceRequest struct {
	APICredentials
	Request BalanceRequest `json:"request"`
}

// TerminalRefundRequest contains refund requests.
type TerminalRefundRequest struct {
	APICredentials
	Request RefundRequest `json:"request"`
}

// TerminalEnrollRequest contains the information needed to enroll a new
// payment method in the token vault.
type TerminalEnrollRequest struct {
	APICredentials
	Request EnrollRequest `json:"request"`
}

// TerminalClearTerminalRequest contains the information needed to enroll a
// new payment method in the token vault.
type TerminalClearTerminalRequest struct {
	APICredentials
	Request ClearTerminalRequest `json:"request"`
}

// TerminalGiftActivateRequest contains the information needed to activate or
// recharge a gift card.
type TerminalGiftActivateRequest struct {
	APICredentials
	Request GiftActivateRequest `json:"request"`
}

// TerminalTermsAndConditionsRequest contains the fields needed for custom T&C
// prompts.
type TerminalTermsAndConditionsRequest struct {
	APICredentials
	Request TermsAndConditionsRequest `json:"request"`
}

// TerminalTermsAndConditionsResponse contains a T&C signature capture
// response.
type TerminalTermsAndConditionsResponse struct {
	APICredentials
	Request TermsAndConditionsResponse `json:"request"`
}

// TerminalTransactionDisplayRequest is used to start or update a transaction
// line item display on a terminal.
type TerminalTransactionDisplayRequest struct {
	APICredentials
	Request TransactionDisplayRequest `json:"request"`
}
