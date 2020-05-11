// Copyright 2019 BlockChyp, Inc. All rights reserved. Use of this code is
// governed by a license that can be found in the LICENSE file.
//
// This file was generated automatically. Changes to this file will be lost
// every time the code is regenerated.

package blockchyp

import (
	"reflect"
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
	PromptTypeFirstName      = "first-name"
	PromptTypeLastName       = "last-name"
)

// AVSResponse indicates the result of address verification.
type AVSResponse string

// AVSResponse types.
const (
	AVSResponseNotApplicable             AVSResponse = ""
	AVSResponseNotSupported                          = "not_supported"
	AVSResponseRetry                                 = "retry"
	AVSResponseNoMatch                               = "no_match"
	AVSResponseAddressMatch                          = "address_match"
	AVSResponsePostalCodeMatch                       = "zip_match"
	AVSResponseAddressAndPostalCodeMatch             = "match"
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

	// Surcharge is the amount added to the transaction to cover eligible credit
	// card fees.
	Surcharge string `json:"surcharge,omitempty"`

	// CashDiscount is the discount applied to the transaction for payment
	// methods ineligible for surcharges.
	CashDiscount string `json:"cashDiscount,omitempty"`
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

// CaptureSignatureRequest contains a request for customer signature data.
type CaptureSignatureRequest struct {
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

	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// SigFile is a location on the filesystem which a customer signature should
	// be written to.
	SigFile string `json:"sigFile,omitempty"`

	// SigFormat specifies the image format to be used for returning signatures.
	SigFormat SignatureFormat `json:"sigFormat,omitempty"`

	// SigWidth is the width that the signature image should be scaled to,
	// preserving the aspect ratio. If not provided, the signature is returned in
	// the terminal's max resolution.
	SigWidth int `json:"sigWidth,omitempty"`

	// DisableSignature specifies whether or not signature prompt should be
	// skipped on the terminal. The terminal will indicate whether or not a
	// signature is required by the card in the receipt suggestions response.
	DisableSignature bool `json:"disableSignature,omitempty"`

	// TerminalName is the name of the target payment terminal.
	TerminalName string `json:"terminalName,omitempty"`
}

// CaptureSignatureResponse contains customer signature data.
type CaptureSignatureResponse struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// SigFile contains the hex encoded signature data.
	SigFile string `json:"sigFile,omitempty"`
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

	// Timeout is the request timeout in seconds.
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

	// Timeout is the request timeout in seconds.
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

	// Timeout is the request timeout in seconds.
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

	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// TerminalName is the name of the target payment terminal.
	TerminalName string `json:"terminalName,omitempty"`

	// PromptType is the prompt type (email, phone, etc).
	PromptType PromptType `json:"promptType"`
}

// CustomerRequest models a customer data request.
type CustomerRequest struct {
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

	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// CustomerID BlockChyp assigned customer id.
	CustomerID string `json:"customerId"`

	// CustomerRef optional customer ref that can be used for the client's
	// system's customer id.
	CustomerRef string `json:"customerRef"`
}

// CustomerResponse models a customer data response.
type CustomerResponse struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// Customer the customer record.
	Customer *Customer `json:"customer"`
}

// CustomerSearchRequest models a customer data search request.
type CustomerSearchRequest struct {
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

	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Query search query for customer searches.
	Query string `json:"query"`
}

// UpdateCustomerRequest models a customer data search request.
type UpdateCustomerRequest struct {
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

	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Customer models a customer update request.
	Customer Customer `json:"customer"`
}

// CustomerSearchResponse models customer search results.
type CustomerSearchResponse struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// Customers the customer results matching the search query.
	Customers []Customer `json:"customers"`
}

// Customer models a customer record.
type Customer struct {
	// ID BlockChyp assigned customer id.
	ID string `json:"id"`

	// CustomerRef optional customer ref that can be used for the client's
	// system's customer id.
	CustomerRef string `json:"customerRef"`

	// FirstName customer's first name.
	FirstName string `json:"firstName"`

	// LastName customer's last name.
	LastName string `json:"lastName"`

	// CompanyName customer's company name.
	CompanyName string `json:"companyName"`

	// EmailAddress customer's email address.
	EmailAddress string `json:"emailAddress"`

	// SmsNumber customer's SMS or mobile number.
	SmsNumber string `json:"smsNumber"`

	// PaymentMethods model saved payment methods associated with a customer.
	PaymentMethods []CustomerToken `json:"paymentMethods"`
}

// CustomerToken models a customer token.
type CustomerToken struct {
	// Token BlockChyp assigned customer id.
	Token string `json:"token"`

	// MaskedPAN masked primary account number.
	MaskedPAN string `json:"maskedPan"`

	// ExpiryMonth expiration month.
	ExpiryMonth string `json:"expiryMonth"`

	// ExpiryYear expiration month.
	ExpiryYear string `json:"expiryYear"`

	// PaymentType payment type.
	PaymentType string `json:"paymentType"`
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

// AuthorizationRequest contains an authorization request for a charge,
// preauth, or reverse transaction.
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

	// Timeout is the request timeout in seconds.
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

	// TransactionID is the ID of the previous transaction being referenced.
	TransactionID string `json:"transactionId"`

	// CurrencyCode indicates the transaction currency code.
	CurrencyCode string `json:"currencyCode"`

	// Amount is the requested amount.
	Amount string `json:"amount"`

	// TaxExempt indicates that the request is tax exempt. Only required for tax
	// exempt level 2 processing.
	TaxExempt bool `json:"taxExempt"`

	// Surcharge is a flag to add a surcharge to the transaction to cover credit
	// card fees, if permitted.
	Surcharge bool `json:"surcharge"`

	// CashDiscount is a flag that applies a discount to negate the surcharge for
	// debit transactions or other surcharge ineligible payment methods.
	CashDiscount bool `json:"cashDiscount"`

	// SigFile is a location on the filesystem which a customer signature should
	// be written to.
	SigFile string `json:"sigFile,omitempty"`

	// SigFormat specifies the image format to be used for returning signatures.
	SigFormat SignatureFormat `json:"sigFormat,omitempty"`

	// SigWidth is the width that the signature image should be scaled to,
	// preserving the aspect ratio. If not provided, the signature is returned in
	// the terminal's max resolution.
	SigWidth int `json:"sigWidth,omitempty"`

	// DisableSignature specifies whether or not signature prompt should be
	// skipped on the terminal. The terminal will indicate whether or not a
	// signature is required by the card in the receipt suggestions response.
	DisableSignature bool `json:"disableSignature,omitempty"`

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

	// TerminalName is the name of the target payment terminal.
	TerminalName string `json:"terminalName,omitempty"`

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

	// Customer contains suggested receipt fields.
	Customer *Customer `json:"customer"`
}

// BalanceRequest contains a request for the remaining balance on a payment
// type.
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

	// Timeout is the request timeout in seconds.
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

	// AVSResponse contains address verification results if address information
	// was submitted.
	AVSResponse AVSResponse `json:"avsResponse"`

	// ReceiptSuggestions contains suggested receipt fields.
	ReceiptSuggestions ReceiptSuggestions `json:"receiptSuggestions"`

	// Customer contains customer data, if any.
	Customer *Customer `json:"customer"`

	// RemainingBalance remaining balance on the payment method.
	RemainingBalance string `json:"remainingBalance,omitempty"`
}

// RefundRequest contains a refund request.
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

	// Timeout is the request timeout in seconds.
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

	// TransactionID is the ID of the previous transaction being referenced.
	TransactionID string `json:"transactionId"`

	// CurrencyCode indicates the transaction currency code.
	CurrencyCode string `json:"currencyCode"`

	// Amount is the requested amount.
	Amount string `json:"amount"`

	// TaxExempt indicates that the request is tax exempt. Only required for tax
	// exempt level 2 processing.
	TaxExempt bool `json:"taxExempt"`

	// Surcharge is a flag to add a surcharge to the transaction to cover credit
	// card fees, if permitted.
	Surcharge bool `json:"surcharge"`

	// CashDiscount is a flag that applies a discount to negate the surcharge for
	// debit transactions or other surcharge ineligible payment methods.
	CashDiscount bool `json:"cashDiscount"`

	// SigFile is a location on the filesystem which a customer signature should
	// be written to.
	SigFile string `json:"sigFile,omitempty"`

	// SigFormat specifies the image format to be used for returning signatures.
	SigFormat SignatureFormat `json:"sigFormat,omitempty"`

	// SigWidth is the width that the signature image should be scaled to,
	// preserving the aspect ratio. If not provided, the signature is returned in
	// the terminal's max resolution.
	SigWidth int `json:"sigWidth,omitempty"`

	// DisableSignature specifies whether or not signature prompt should be
	// skipped on the terminal. The terminal will indicate whether or not a
	// signature is required by the card in the receipt suggestions response.
	DisableSignature bool `json:"disableSignature,omitempty"`

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

	// TerminalName is the name of the target payment terminal.
	TerminalName string `json:"terminalName,omitempty"`
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

	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// TransactionID is the ID of the previous transaction being referenced.
	TransactionID string `json:"transactionId"`

	// CurrencyCode indicates the transaction currency code.
	CurrencyCode string `json:"currencyCode"`

	// Amount is the requested amount.
	Amount string `json:"amount"`

	// TaxExempt indicates that the request is tax exempt. Only required for tax
	// exempt level 2 processing.
	TaxExempt bool `json:"taxExempt"`

	// Surcharge is a flag to add a surcharge to the transaction to cover credit
	// card fees, if permitted.
	Surcharge bool `json:"surcharge"`

	// CashDiscount is a flag that applies a discount to negate the surcharge for
	// debit transactions or other surcharge ineligible payment methods.
	CashDiscount bool `json:"cashDiscount"`

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

	// Approved indicates that the transaction was approved.
	Approved bool `json:"approved"`

	// AuthCode is the auth code from the payment network.
	AuthCode string `json:"authCode,omitempty"`

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

	// AVSResponse contains address verification results if address information
	// was submitted.
	AVSResponse AVSResponse `json:"avsResponse"`

	// ReceiptSuggestions contains suggested receipt fields.
	ReceiptSuggestions ReceiptSuggestions `json:"receiptSuggestions"`

	// Customer contains customer data, if any.
	Customer *Customer `json:"customer"`
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

	// Timeout is the request timeout in seconds.
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

	// Approved indicates that the transaction was approved.
	Approved bool `json:"approved"`

	// AuthCode is the auth code from the payment network.
	AuthCode string `json:"authCode,omitempty"`

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

	// AVSResponse contains address verification results if address information
	// was submitted.
	AVSResponse AVSResponse `json:"avsResponse"`

	// ReceiptSuggestions contains suggested receipt fields.
	ReceiptSuggestions ReceiptSuggestions `json:"receiptSuggestions"`

	// Customer contains customer data, if any.
	Customer *Customer `json:"customer"`

	// SigFile contains the hex encoded signature data.
	SigFile string `json:"sigFile,omitempty"`
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

	// Timeout is the request timeout in seconds.
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

	// Customer customer with which the new token should be associated.
	Customer *Customer `json:"customer"`
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

	// Approved indicates that the transaction was approved.
	Approved bool `json:"approved"`

	// AuthCode is the auth code from the payment network.
	AuthCode string `json:"authCode,omitempty"`

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

	// AVSResponse contains address verification results if address information
	// was submitted.
	AVSResponse AVSResponse `json:"avsResponse"`

	// ReceiptSuggestions contains suggested receipt fields.
	ReceiptSuggestions ReceiptSuggestions `json:"receiptSuggestions"`

	// Customer contains customer data, if any.
	Customer *Customer `json:"customer"`

	// SigFile contains the hex encoded signature data.
	SigFile string `json:"sigFile,omitempty"`
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

	// Timeout is the request timeout in seconds.
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

	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// CurrencyCode indicates the transaction currency code.
	CurrencyCode string `json:"currencyCode"`

	// Amount is the requested amount.
	Amount string `json:"amount"`

	// TaxExempt indicates that the request is tax exempt. Only required for tax
	// exempt level 2 processing.
	TaxExempt bool `json:"taxExempt"`

	// Surcharge is a flag to add a surcharge to the transaction to cover credit
	// card fees, if permitted.
	Surcharge bool `json:"surcharge"`

	// CashDiscount is a flag that applies a discount to negate the surcharge for
	// debit transactions or other surcharge ineligible payment methods.
	CashDiscount bool `json:"cashDiscount"`

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

	// MaskedPAN is the masked card identifier.
	MaskedPAN string `json:"maskedPan,omitempty"`
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

	// Timeout is the request timeout in seconds.
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

// TermsAndConditionsRequest contains the fields needed for custom Terms and
// Conditions prompts.
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

	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

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

	// DisableSignature specifies whether or not signature prompt should be
	// skipped on the terminal. The terminal will indicate whether or not a
	// signature is required by the card in the receipt suggestions response.
	DisableSignature bool `json:"disableSignature,omitempty"`

	// TerminalName is the name of the target payment terminal.
	TerminalName string `json:"terminalName,omitempty"`

	// TCAlias is an alias for a Terms and Conditions template configured in the
	// BlockChyp dashboard.
	TCAlias string `json:"tcAlias"`

	// TCName contains the name of the Terms and Conditions the user is
	// accepting.
	TCName string `json:"tcName"`

	// TCContent is the content of the terms and conditions that will be
	// presented to the user.
	TCContent string `json:"tcContent"`

	// SigRequired indicates that a signature should be requested.
	SigRequired bool `json:"sigRequired"`
}

// TermsAndConditionsResponse contains a signature capture response for Terms
// and Conditions.
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

// AuthorizationResponse contains the response to an authorization request.
type AuthorizationResponse struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// Approved indicates that the transaction was approved.
	Approved bool `json:"approved"`

	// AuthCode is the auth code from the payment network.
	AuthCode string `json:"authCode,omitempty"`

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

	// AVSResponse contains address verification results if address information
	// was submitted.
	AVSResponse AVSResponse `json:"avsResponse"`

	// ReceiptSuggestions contains suggested receipt fields.
	ReceiptSuggestions ReceiptSuggestions `json:"receiptSuggestions"`

	// Customer contains customer data, if any.
	Customer *Customer `json:"customer"`

	// SigFile contains the hex encoded signature data.
	SigFile string `json:"sigFile,omitempty"`

	// WhiteListedCard contains card BIN ranges can be whitelisted so that they
	// are read instead of being processed directly. This is useful for
	// integration with legacy gift card systems.
	WhiteListedCard *WhiteListedCard `json:"whiteListedCard"`

	// StoreAndForward indicates that the transaction was flagged for store and
	// forward due to network problems.
	StoreAndForward bool `json:"storeAndForward"`
}

// TransactionStatusRequest models the request for updated information about a
// transaction.
type TransactionStatusRequest struct {
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

	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// TransactionID is the BlockChyp assigned transaction id.
	TransactionID string `json:"transactionId,omitempty"`
}

// TransactionStatus models the status of a transaction.
type TransactionStatus struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// Approved indicates that the transaction was approved.
	Approved bool `json:"approved"`

	// AuthCode is the auth code from the payment network.
	AuthCode string `json:"authCode,omitempty"`

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

	// AVSResponse contains address verification results if address information
	// was submitted.
	AVSResponse AVSResponse `json:"avsResponse"`

	// ReceiptSuggestions contains suggested receipt fields.
	ReceiptSuggestions ReceiptSuggestions `json:"receiptSuggestions"`

	// Customer contains customer data, if any.
	Customer *Customer `json:"customer"`

	// SigFile contains the hex encoded signature data.
	SigFile string `json:"sigFile,omitempty"`

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

	// Timeout is the request timeout in seconds.
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

// TerminalStatusRequest contains a request for the status of a terminal.
type TerminalStatusRequest struct {
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

	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// TerminalName is the name of the target payment terminal.
	TerminalName string `json:"terminalName,omitempty"`
}

// TerminalStatusResponse contains the current status of a terminal.
type TerminalStatusResponse struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// Idle indicates that the terminal is idle.
	Idle bool `json:"idle"`

	// Status contains the operation that the terminal is performing.
	Status string `json:"status"`

	// TransactionRef contains the transaction reference for an ongoing
	// transaction, if one was specified at request time.
	TransactionRef string `json:"transactionRef"`

	// Since is the timestamp of the last status change.
	Since time.Time `json:"since"`
}

// PaymentLinkRequest creates a payment link.
type PaymentLinkRequest struct {
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

	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// CurrencyCode indicates the transaction currency code.
	CurrencyCode string `json:"currencyCode"`

	// Amount is the requested amount.
	Amount string `json:"amount"`

	// TaxExempt indicates that the request is tax exempt. Only required for tax
	// exempt level 2 processing.
	TaxExempt bool `json:"taxExempt"`

	// Surcharge is a flag to add a surcharge to the transaction to cover credit
	// card fees, if permitted.
	Surcharge bool `json:"surcharge"`

	// CashDiscount is a flag that applies a discount to negate the surcharge for
	// debit transactions or other surcharge ineligible payment methods.
	CashDiscount bool `json:"cashDiscount"`

	// TerminalName is the name of the target payment terminal.
	TerminalName string `json:"terminalName,omitempty"`

	// AutoSend automatically send the link via an email.
	AutoSend bool `json:"autoSend"`

	// Cashier flags the payment link as cashier facing.
	Cashier bool `json:"cashier"`

	// Description description explaining the transaction for display to the
	// user.
	Description string `json:"description"`

	// Subject subject of the payment email.
	Subject string `json:"subject"`

	// Transaction transaction details for display on the payment email.
	Transaction *TransactionDisplayTransaction `json:"transaction"`

	// Customer contains customer information.
	Customer Customer `json:"customer"`

	// CallbackURL optional callback url to which transaction responses for this
	// link will be posted.
	CallbackURL string `json:"callbackUrl"`

	// TCAlias is an alias for a Terms and Conditions template configured in the
	// BlockChyp dashboard.
	TCAlias string `json:"tcAlias"`

	// TCName contains the name of the Terms and Conditions the user is
	// accepting.
	TCName string `json:"tcName"`

	// TCContent is the content of the terms and conditions that will be
	// presented to the user.
	TCContent string `json:"tcContent"`
}

// PaymentLinkResponse creates a payment link.
type PaymentLinkResponse struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// LinkCode the payment link code.
	LinkCode string `json:"linkCode"`

	// URL the url for the payment link.
	URL string `json:"url"`

	// CustomerID the customer id created or used for the payment.
	CustomerID string `json:"customerId"`
}

// CashDiscountRequest computes the cash discount for a cash discount if
// enabled.
type CashDiscountRequest struct {
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

	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// CurrencyCode indicates the transaction currency code.
	CurrencyCode string `json:"currencyCode"`

	// Amount is the requested amount.
	Amount string `json:"amount"`

	// TaxExempt indicates that the request is tax exempt. Only required for tax
	// exempt level 2 processing.
	TaxExempt bool `json:"taxExempt"`

	// Surcharge is a flag to add a surcharge to the transaction to cover credit
	// card fees, if permitted.
	Surcharge bool `json:"surcharge"`

	// CashDiscount is a flag that applies a discount to negate the surcharge for
	// debit transactions or other surcharge ineligible payment methods.
	CashDiscount bool `json:"cashDiscount"`
}

// CashDiscountResponse models the results of a cash discount calculation.
type CashDiscountResponse struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// CurrencyCode indicates the transaction currency code.
	CurrencyCode string `json:"currencyCode"`

	// Amount is the new calculated total amount.
	Amount string `json:"amount"`

	// TaxExempt indicates that the request is tax exempt. Only required for tax
	// exempt level 2 processing.
	TaxExempt bool `json:"taxExempt"`

	// Surcharge is the normal surcharge for a transaction. Will only be returned
	// if an offsetting cash discount is also returned.
	Surcharge string `json:"surcharge"`

	// CashDiscount is the cash discount. Will not be returned in surcharge only
	// mode.
	CashDiscount string `json:"cashDiscount"`
}

// TransactionHistoryRequest models a batch history request.
type TransactionHistoryRequest struct {
	// Test inidicates that the request is for a test merchant account
	Test bool `json:"test"`

	// BatchID optional batch id.
	BatchID string `json:"batchId"`

	// TerminalName optional terminal name.
	TerminalName string `json:"terminalName"`

	// StartDate optional start date filter for batch history.
	StartDate time.Time `json:"startDate"`

	// EndDate optional end date filter for batch history.
	EndDate time.Time `json:"endDate"`
}

// TransactionHistoryResponse models response to a batch history request.
type TransactionHistoryResponse struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// BatchID batch identifier if filtered by batch.
	BatchID string `json:"batchId"`

	// TerminalName terminal name if filtered by terminal.
	TerminalName string `json:"terminalName"`

	// StartDate start date if filtered by start date.
	StartDate time.Time `json:"startDate"`

	// EndDate end date if filtered by end date.
	EndDate time.Time `json:"endDate"`

	// Transactions matching transaction history.
	Transactions []TransactionStatus `json:"transactions"`
}

// BatchHistoryRequest models a batch history request.
type BatchHistoryRequest struct {
	// Test inidates that the request is for a test merchant account
	Test bool `json:"test"`

	// StartDate optional start date filter for batch history.
	StartDate time.Time `json:"startDate"`

	// EndDate optional end date filter for batch history.
	EndDate time.Time `json:"endDate"`
}

// BatchHistoryResponse models response to a batch history request.
type BatchHistoryResponse struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// StartDate start date if filtered by start date.
	StartDate time.Time `json:"startDate"`

	// EndDate end date if filtered by end date.
	EndDate time.Time `json:"endDate"`

	// Batches merchant's batch history in descending order.
	Batches []BatchDetails `json:"batches"`
}

// BatchDetails models high level information about a single batch.
type BatchDetails struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// BatchID batch identifier.
	BatchID string `json:"batchId"`

	// CapturedAmount is the new captured amount.
	CapturedAmount string `json:"capturedAmount"`

	// UncapturedPreauths is the new captured amount.
	UncapturedPreauths string `json:"uncapturedPreauths"`

	// TransactionCount the number of transactions in the batch
	TransactionCount int `json:"transactionCount"`

	// Open flag indicating whether or not the batch is open
	Open bool `json:"open"`

	// OpenDate date and time of the first transaction for this batch.
	OpenDate time.Time `json:"openDate"`

	// CloseDate date and time the batch was closed.
	CloseDate time.Time `json:"closeDate"`
}

// TerminalCaptureSignatureRequest contains a request for customer signature
// data.
type TerminalCaptureSignatureRequest struct {
	APICredentials
	Request CaptureSignatureRequest `json:"request"`
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

// TerminalAuthorizationRequest contains an authorization request for a
// charge, preauth, or reverse transaction.
type TerminalAuthorizationRequest struct {
	APICredentials
	Request AuthorizationRequest `json:"request"`
}

// TerminalBalanceRequest contains a request for the remaining balance on a
// payment type.
type TerminalBalanceRequest struct {
	APICredentials
	Request BalanceRequest `json:"request"`
}

// TerminalRefundRequest contains a refund request.
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

// TerminalTermsAndConditionsRequest contains the fields needed for custom
// Terms and Conditions prompts.
type TerminalTermsAndConditionsRequest struct {
	APICredentials
	Request TermsAndConditionsRequest `json:"request"`
}

// TerminalTermsAndConditionsResponse contains a signature capture response
// for Terms and Conditions.
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

// TerminalTerminalStatusRequest contains a request for the status of a
// terminal.
type TerminalTerminalStatusRequest struct {
	APICredentials
	Request TerminalStatusRequest `json:"request"`
}

// AbstractAcknowledgement contains fields which should be returned with
// standard requests.
type AbstractAcknowledgement struct {
	// Success indicates whether or not the request succeeded.
	Success bool

	// Error is the error, if an error occurred.
	Error string

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string
}

// From creates an instance of AbstractAcknowledgement with values
// from a generic type.
func (r AbstractAcknowledgement) From(raw interface{}) (result AbstractAcknowledgement, ok bool) {
	ok = copyTo(raw, &r)
	return r, ok
}

// TerminalReference contains a reference to a terminal name.
type TerminalReference struct {
	// TerminalName is the name of the target payment terminal.
	TerminalName string
}

// From creates an instance of TerminalReference with values
// from a generic type.
func (r TerminalReference) From(raw interface{}) (result TerminalReference, ok bool) {
	ok = copyTo(raw, &r)
	return r, ok
}

// SignatureResponse contains customer signature data.
type SignatureResponse struct {
	// SigFile contains the hex encoded signature data.
	SigFile string
}

// From creates an instance of SignatureResponse with values
// from a generic type.
func (r SignatureResponse) From(raw interface{}) (result SignatureResponse, ok bool) {
	ok = copyTo(raw, &r)
	return r, ok
}

// SignatureRequest contains a request for customer signature data.
type SignatureRequest struct {
	// SigFile is a location on the filesystem which a customer signature should
	// be written to.
	SigFile string

	// SigFormat specifies the image format to be used for returning signatures.
	SigFormat SignatureFormat

	// SigWidth is the width that the signature image should be scaled to,
	// preserving the aspect ratio. If not provided, the signature is returned in
	// the terminal's max resolution.
	SigWidth int

	// DisableSignature specifies whether or not signature prompt should be
	// skipped on the terminal. The terminal will indicate whether or not a
	// signature is required by the card in the receipt suggestions response.
	DisableSignature bool
}

// From creates an instance of SignatureRequest with values
// from a generic type.
func (r SignatureRequest) From(raw interface{}) (result SignatureRequest, ok bool) {
	ok = copyTo(raw, &r)
	return r, ok
}

// ApprovalResponse contains response fields for an approved transaction.
type ApprovalResponse struct {
	// Approved indicates that the transaction was approved.
	Approved bool

	// AuthCode is the auth code from the payment network.
	AuthCode string
}

// From creates an instance of ApprovalResponse with values
// from a generic type.
func (r ApprovalResponse) From(raw interface{}) (result ApprovalResponse, ok bool) {
	ok = copyTo(raw, &r)
	return r, ok
}

// CoreRequest contains core request fields for a transaction.
type CoreRequest struct {
	// TransactionRef is the transaction reference string assigned to the
	// transaction request. If no transaction ref was assiged on the request,
	// then the gateway will randomly generate one.
	TransactionRef string

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool

	// Timeout is the request timeout in seconds.
	Timeout int
}

// From creates an instance of CoreRequest with values
// from a generic type.
func (r CoreRequest) From(raw interface{}) (result CoreRequest, ok bool) {
	ok = copyTo(raw, &r)
	return r, ok
}

// PaymentMethodResponse contains response details about a payment method.
type PaymentMethodResponse struct {
	// Token is the payment token, if the payment was enrolled in the vault.
	Token string

	// EntryMethod is the entry method for the transaction (CHIP, MSR, KEYED,
	// etc).
	EntryMethod string

	// PaymentType is the card brand (VISA, MC, AMEX, etc).
	PaymentType string

	// MaskedPAN is the masked primary account number.
	MaskedPAN string

	// PublicKey is the BlockChyp public key if the user presented a BlockChyp
	// payment card.
	PublicKey string

	// ScopeAlert indicates that the transaction did something that would put the
	// system in PCI scope.
	ScopeAlert bool

	// CardHolder is the cardholder name.
	CardHolder string

	// AVSResponse contains address verification results if address information
	// was submitted.
	AVSResponse AVSResponse

	// ReceiptSuggestions contains suggested receipt fields.
	ReceiptSuggestions ReceiptSuggestions

	// Customer contains customer data, if any.
	Customer *Customer
}

// From creates an instance of PaymentMethodResponse with values
// from a generic type.
func (r PaymentMethodResponse) From(raw interface{}) (result PaymentMethodResponse, ok bool) {
	ok = copyTo(raw, &r)
	return r, ok
}

// PaymentAmounts contains response details about tender amounts.
type PaymentAmounts struct {
	// PartialAuth indicates whether or not the transaction was approved for a
	// partial amount.
	PartialAuth bool

	// AltCurrency indicates whether or not an alternate currency was used.
	AltCurrency bool

	// FSAAuth indicates whether or not a request was settled on an FSA card.
	FSAAuth bool

	// CurrencyCode is the currency code used for the transaction.
	CurrencyCode string

	// RequestedAmount is the requested amount.
	RequestedAmount string

	// AuthorizedAmount is the authorized amount. May not match the requested
	// amount in the event of a partial auth.
	AuthorizedAmount string

	// RemainingBalance is the remaining balance on the payment method.
	RemainingBalance string

	// TipAmount is the tip amount.
	TipAmount string

	// TaxAmount is the tax amount.
	TaxAmount string

	// RequestedCashBackAmount is the cash back amount the customer requested
	// during the transaction.
	RequestedCashBackAmount string

	// AuthorizedCashBackAmount is the amount of cash back authorized by the
	// gateway. This amount will be the entire amount requested, or zero.
	AuthorizedCashBackAmount string
}

// From creates an instance of PaymentAmounts with values
// from a generic type.
func (r PaymentAmounts) From(raw interface{}) (result PaymentAmounts, ok bool) {
	ok = copyTo(raw, &r)
	return r, ok
}

// PaymentMethod contains request details about a payment method.
type PaymentMethod struct {
	// Token is the payment token to be used for this transaction. This should be
	// used for recurring transactions.
	Token string

	// Track1 contains track 1 magnetic stripe data.
	Track1 string

	// Track2 contains track 2 magnetic stripe data.
	Track2 string

	// PAN contains the primary account number. We recommend using the terminal
	// or e-commerce tokenization libraries instead of passing account numbers in
	// directly, as this would put your application in PCI scope.
	PAN string

	// RoutingNumber is the ACH routing number for ACH transactions.
	RoutingNumber string

	// CardholderName is the cardholder name. Only required if the request
	// includes a primary account number or track data.
	CardholderName string

	// ExpMonth is the card expiration month for use with PAN based transactions.
	ExpMonth string

	// ExpYear is the card expiration year for use with PAN based transactions.
	ExpYear string

	// CVV is the card CVV for use with PAN based transactions.
	CVV string

	// Address is the cardholder address for use with address verification.
	Address string

	// PostalCode is the cardholder postal code for use with address
	// verification.
	PostalCode string

	// ManualEntry specifies that the payment entry method is a manual keyed
	// transaction. If this is true, no other payment method will be accepted.
	ManualEntry bool

	// KSN is the key serial number used for DUKPT encryption.
	KSN string

	// PINBlock is the encrypted pin block.
	PINBlock string

	// CardType designates categories of cards: credit, debit, EBT.
	CardType CardType

	// PaymentType designates brands of payment methods: Visa, Discover, etc.
	PaymentType string
}

// From creates an instance of PaymentMethod with values
// from a generic type.
func (r PaymentMethod) From(raw interface{}) (result PaymentMethod, ok bool) {
	ok = copyTo(raw, &r)
	return r, ok
}

// RequestAmount contains request details about tender amounts.
type RequestAmount struct {
	// CurrencyCode indicates the transaction currency code.
	CurrencyCode string

	// Amount is the requested amount.
	Amount string

	// TaxExempt indicates that the request is tax exempt. Only required for tax
	// exempt level 2 processing.
	TaxExempt bool

	// Surcharge is a flag to add a surcharge to the transaction to cover credit
	// card fees, if permitted.
	Surcharge bool

	// CashDiscount is a flag that applies a discount to negate the surcharge for
	// debit transactions or other surcharge ineligible payment methods.
	CashDiscount bool
}

// From creates an instance of RequestAmount with values
// from a generic type.
func (r RequestAmount) From(raw interface{}) (result RequestAmount, ok bool) {
	ok = copyTo(raw, &r)
	return r, ok
}

// Subtotals contains request subtotals.
type Subtotals struct {
	// TipAmount is the tip amount.
	TipAmount string

	// TaxAmount is the tax amount.
	TaxAmount string

	// CashBackAmount is the amount of cash back requested.
	CashBackAmount string

	// FSAEligibleAmount is the amount of the transaction that should be charged
	// to an FSA card. This amount may be less than the transaction total, in
	// which case only this amount will be charged if an FSA card is presented.
	// If the FSA amount is paid on an FSA card, then the FSA amount authorized
	// will be indicated on the response.
	FSAEligibleAmount string

	// HSAEligibleAmount is the amount of the transaction that should be charged
	// to an HSA card.
	HSAEligibleAmount string

	// EBTEligibleAmount is the amount of the transaction that should be charged
	// to an EBT card.
	EBTEligibleAmount string
}

// From creates an instance of Subtotals with values
// from a generic type.
func (r Subtotals) From(raw interface{}) (result Subtotals, ok bool) {
	ok = copyTo(raw, &r)
	return r, ok
}

// PreviousTransaction contains a reference to a previous transaction.
type PreviousTransaction struct {
	// TransactionID is the ID of the previous transaction being referenced.
	TransactionID string
}

// From creates an instance of PreviousTransaction with values
// from a generic type.
func (r PreviousTransaction) From(raw interface{}) (result PreviousTransaction, ok bool) {
	ok = copyTo(raw, &r)
	return r, ok
}

// CoreResponse contains core response fields for a transaction.
type CoreResponse struct {
	// TransactionID is the ID assigned to the transaction.
	TransactionID string

	// BatchID is the ID assigned to the batch.
	BatchID string

	// TransactionRef is the transaction reference string assigned to the
	// transaction request. If no transaction ref was assiged on the request,
	// then the gateway will randomly generate one.
	TransactionRef string

	// TransactionType is the type of transaction.
	TransactionType string

	// Timestamp is the timestamp of the transaction.
	Timestamp string

	// TickBlock is the hash of the last tick block.
	TickBlock string

	// Test indicates that the transaction was processed on the test gateway.
	Test bool

	// Sig is the ECC signature of the response. Can be used to ensure that it
	// was signed by the terminal and detect man-in-the middle attacks.
	Sig string
}

// From creates an instance of CoreResponse with values
// from a generic type.
func (r CoreResponse) From(raw interface{}) (result CoreResponse, ok bool) {
	ok = copyTo(raw, &r)
	return r, ok
}

func copyTo(from, to interface{}) (ok bool) {
	fromV := reflect.ValueOf(from)
	if fromV.Kind() == reflect.Ptr {
		fromV = fromV.Elem()
	}

	toV := reflect.ValueOf(to).Elem()

	for i := 0; i < toV.NumField(); i++ {
		val := fromV.FieldByName(toV.Type().Field(i).Name)
		if !val.IsValid() {
			continue
		}

		ok = true
		toV.Field(i).Set(val)
	}

	return
}

func clearField(ptr interface{}, field string) {
	v := reflect.ValueOf(ptr)
	if v.Kind() != reflect.Ptr {
		panic("must pass by pointer")
	}
	v = v.Elem()

	if f := v.FieldByName(field); f.IsValid() {
		f.Set(reflect.Zero(f.Type()))
	}
}
