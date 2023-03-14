// Copyright 2019-2023 BlockChyp, Inc. All rights reserved. Use of this code
// is governed by a license that can be found in the LICENSE file.
//
// This file was generated automatically by the BlockChyp SDK Generator.
// Changes to this file will be lost every time the code is regenerated.

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
	CardTypeHealthcare
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

// CVMType designates a customer verification method.
type CVMType string

// CVMs defined.
const (
	CVMTypeSignature  CVMType = "Signature"
	CVMTypeOfflinePIN CVMType = "Offline PIN"
	CVMTypeOnlinePIN  CVMType = "Online PIN"
	CVMTypeCDCVM      CVMType = "CDCVM"
	CVMTypeNoCVM      CVMType = "No CVM"
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

// HealthcareType is a category of healthcare.
type HealthcareType string

// HealthcareTypes.
const (
	HealthcareTypeHealthcare   = "healthcare"
	HealthcareTypePrescription = "prescription"
	HealthcareTypeVision       = "vision"
	HealthcareTypeClinic       = "clinic"
	HealthcareTypeDental       = "dental"
)

// RoundingMode indicates how partial penny rounding operations should work
type RoundingMode string

// RoundingMode types
const (
	RoundingModeUp      = "up"
	RoundingModeNearest = "nearest"
	RoundingModeDown    = "down"
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

	// CVMUsed indicates the customer verification method used for the
	// transaction.
	CVMUsed CVMType `json:"cvmUsed,omitempty"`

	// Fallback indicates that a chip read failure caused the transaction to fall
	// back to the magstripe.
	Fallback bool `json:"fallback,omitempty"`

	// BatchSequence is the sequence of the transaction in the batch.
	BatchSequence int `json:"batchSequence,omitempty"`

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
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string `json:"transactionRef,omitempty"`

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool `json:"async"`

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool `json:"queue"`

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool `json:"waitForRemovedCard,omitempty"`

	// Force causes a transaction to override any in-progress transactions.
	Force bool `json:"force,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

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

	// ResetConnection forces the terminal cloud connection to be reset while a
	// transactions is in flight. This is a diagnostic settings that can be used
	// only for test transactions.
	ResetConnection bool `json:"resetConnection"`
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
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string `json:"transactionRef,omitempty"`

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool `json:"async"`

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool `json:"queue"`

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool `json:"waitForRemovedCard,omitempty"`

	// Force causes a transaction to override any in-progress transactions.
	Force bool `json:"force,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// TerminalName is the name of the target payment terminal.
	TerminalName string `json:"terminalName,omitempty"`

	// ResetConnection forces the terminal cloud connection to be reset while a
	// transactions is in flight. This is a diagnostic settings that can be used
	// only for test transactions.
	ResetConnection bool `json:"resetConnection"`
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

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// Sig is the ECC signature of the response. Can be used to ensure that it
	// was signed by the terminal and detect man-in-the middle attacks.
	Sig string `json:"sig,omitempty"`
}

// LocateRequest contains information needed to retrieve location information
// for a terminal.
type LocateRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string `json:"transactionRef,omitempty"`

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool `json:"async"`

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool `json:"queue"`

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool `json:"waitForRemovedCard,omitempty"`

	// Force causes a transaction to override any in-progress transactions.
	Force bool `json:"force,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// TerminalName is the name of the target payment terminal.
	TerminalName string `json:"terminalName,omitempty"`

	// ResetConnection forces the terminal cloud connection to be reset while a
	// transactions is in flight. This is a diagnostic settings that can be used
	// only for test transactions.
	ResetConnection bool `json:"resetConnection"`
}

// LocateResponse contains the response to a locate request.
type LocateResponse struct {
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

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// Sig is the ECC signature of the response. Can be used to ensure that it
	// was signed by the terminal and detect man-in-the middle attacks.
	Sig string `json:"sig,omitempty"`

	// TerminalName is the name assigned to the terminal at activation.
	TerminalName string `json:"terminalName"`

	// IPAddress is the local IP address of the terminal.
	IPAddress string `json:"ipAddress"`

	// CloudRelay indicates whether or not the terminal is running in cloud relay
	// mode.
	CloudRelay bool `json:"cloudRelay"`

	// PublicKey is the terminal's public key.
	PublicKey string `json:"publicKey"`
}

// MessageRequest contains a message to be displayed on the terminal screen.
type MessageRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string `json:"transactionRef,omitempty"`

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool `json:"async"`

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool `json:"queue"`

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool `json:"waitForRemovedCard,omitempty"`

	// Force causes a transaction to override any in-progress transactions.
	Force bool `json:"force,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// TerminalName is the name of the target payment terminal.
	TerminalName string `json:"terminalName,omitempty"`

	// ResetConnection forces the terminal cloud connection to be reset while a
	// transactions is in flight. This is a diagnostic settings that can be used
	// only for test transactions.
	ResetConnection bool `json:"resetConnection"`

	// Message is the message to be displayed on the terminal.
	Message string `json:"message"`
}

// BooleanPromptRequest contains a simple yes no prompt request.
type BooleanPromptRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string `json:"transactionRef,omitempty"`

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool `json:"async"`

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool `json:"queue"`

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool `json:"waitForRemovedCard,omitempty"`

	// Force causes a transaction to override any in-progress transactions.
	Force bool `json:"force,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// TerminalName is the name of the target payment terminal.
	TerminalName string `json:"terminalName,omitempty"`

	// ResetConnection forces the terminal cloud connection to be reset while a
	// transactions is in flight. This is a diagnostic settings that can be used
	// only for test transactions.
	ResetConnection bool `json:"resetConnection"`

	// YesCaption is the preferred caption for the 'yes' button.
	YesCaption string `json:"yesCaption"`

	// NoCaption is the preferred caption for the 'no' button.
	NoCaption string `json:"noCaption"`

	// Prompt is the text to be displayed on the terminal.
	Prompt string `json:"prompt"`
}

// TextPromptRequest contains a text prompt request.
type TextPromptRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string `json:"transactionRef,omitempty"`

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool `json:"async"`

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool `json:"queue"`

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool `json:"waitForRemovedCard,omitempty"`

	// Force causes a transaction to override any in-progress transactions.
	Force bool `json:"force,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// TerminalName is the name of the target payment terminal.
	TerminalName string `json:"terminalName,omitempty"`

	// ResetConnection forces the terminal cloud connection to be reset while a
	// transactions is in flight. This is a diagnostic settings that can be used
	// only for test transactions.
	ResetConnection bool `json:"resetConnection"`

	// PromptType is the prompt type (email, phone, etc).
	PromptType PromptType `json:"promptType"`
}

// CustomerRequest models a customer data request.
type CustomerRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string `json:"transactionRef,omitempty"`

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool `json:"async"`

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool `json:"queue"`

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool `json:"waitForRemovedCard,omitempty"`

	// Force causes a transaction to override any in-progress transactions.
	Force bool `json:"force,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

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
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string `json:"transactionRef,omitempty"`

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool `json:"async"`

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool `json:"queue"`

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool `json:"waitForRemovedCard,omitempty"`

	// Force causes a transaction to override any in-progress transactions.
	Force bool `json:"force,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// Query search query for customer searches.
	Query string `json:"query"`
}

// UpdateCustomerRequest models a customer data search request.
type UpdateCustomerRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string `json:"transactionRef,omitempty"`

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool `json:"async"`

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool `json:"queue"`

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool `json:"waitForRemovedCard,omitempty"`

	// Force causes a transaction to override any in-progress transactions.
	Force bool `json:"force,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

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

// TokenMetadataRequest retrieves token metadata.
type TokenMetadataRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string `json:"transactionRef,omitempty"`

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool `json:"async"`

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool `json:"queue"`

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool `json:"waitForRemovedCard,omitempty"`

	// Force causes a transaction to override any in-progress transactions.
	Force bool `json:"force,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// Token the token to retrieve.
	Token string `json:"token"`
}

// TokenMetadataResponse models a payment token metadata response.
type TokenMetadataResponse struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// Token the token metadata for a given query.
	Token CustomerToken `json:"token"`
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

	// Customers models customer records associated with a payment token.
	Customers []Customer `json:"customers"`
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
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string `json:"transactionRef,omitempty"`

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool `json:"async"`

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool `json:"queue"`

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool `json:"waitForRemovedCard,omitempty"`

	// Force causes a transaction to override any in-progress transactions.
	Force bool `json:"force,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

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

	// TerminalName is the name of the target payment terminal.
	TerminalName string `json:"terminalName,omitempty"`

	// ResetConnection forces the terminal cloud connection to be reset while a
	// transactions is in flight. This is a diagnostic settings that can be used
	// only for test transactions.
	ResetConnection bool `json:"resetConnection"`

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
	// Use only if you want to set your own exchange rate for a crypto
	// transaction.
	AltPrices map[string]string `json:"altPrices,omitempty"`

	// Customer contains customer information.
	Customer *Customer `json:"customer"`

	// RoundingMode indicates how partial pennies should be rounded for
	// calculated values like surcharges. Rounding up is the default behavior.
	RoundingMode *RoundingMode `json:"roundingMode"`

	// Healthcare contains details for HSA/FSA transactions.
	Healthcare *Healthcare `json:"healthcare,omitempty"`

	// Cryptocurrency indicates that the transaction should be a cryptocurrency
	// transaction. Value should be a crypto currency code (ETH, BTC) or ANY to
	// prompt the user to choose from supported cryptocurrencies.
	Cryptocurrency *string `json:"cryptocurrency,omitempty"`

	// CryptoNetwork is an optional parameter that can be used to force a crypto
	// transaction onto a level one or level two network. Valid values are L1 and
	// L2. Defaults to L1.
	CryptoNetwork *string `json:"cryptoNetwork,omitempty"`

	// CryptoReceiveAddress can be used to specify a specific receive address for
	// a crypto transaction. Disabled by default. This should only be used by
	// sophisticated users with access to properly configured hot wallets.
	CryptoReceiveAddress *string `json:"cryptoReceiveAddress,omitempty"`

	// PaymentRequestLabel can optionally add a label to the payment request if
	// the target cryptocurrency supports labels. Defaults to the merchant's DBA
	// Name.
	PaymentRequestLabel *string `json:"paymentRequestLabel,omitempty"`

	// PaymentRequestMessage can optionally add a message to the payment request
	// if the target cryptocurrency supports labels. Defaults to empty.
	PaymentRequestMessage *string `json:"paymentRequestMessage,omitempty"`
}

// BalanceRequest contains a request for the remaining balance on a payment
// type.
type BalanceRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string `json:"transactionRef,omitempty"`

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool `json:"async"`

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool `json:"queue"`

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool `json:"waitForRemovedCard,omitempty"`

	// Force causes a transaction to override any in-progress transactions.
	Force bool `json:"force,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

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

	// ResetConnection forces the terminal cloud connection to be reset while a
	// transactions is in flight. This is a diagnostic settings that can be used
	// only for test transactions.
	ResetConnection bool `json:"resetConnection"`
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

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

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

	// ExpMonth is the card expiration month in MM format.
	ExpMonth string `json:"expMonth,omitempty"`

	// ExpYear is the card expiration year in YY format.
	ExpYear string `json:"expYear,omitempty"`

	// AVSResponse contains address verification results if address information
	// was submitted.
	AVSResponse AVSResponse `json:"avsResponse"`

	// ReceiptSuggestions contains suggested receipt fields.
	ReceiptSuggestions ReceiptSuggestions `json:"receiptSuggestions"`

	// Customer contains customer data, if any. Preserved for reverse
	// compatibility.
	Customer *Customer `json:"customer"`

	// Customers contains customer data, if any.
	Customers []Customer `json:"customers"`

	// RemainingBalance remaining balance on the payment method.
	RemainingBalance string `json:"remainingBalance,omitempty"`
}

// RefundRequest contains a refund request.
type RefundRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string `json:"transactionRef,omitempty"`

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool `json:"async"`

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool `json:"queue"`

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool `json:"waitForRemovedCard,omitempty"`

	// Force causes a transaction to override any in-progress transactions.
	Force bool `json:"force,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

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

	// TerminalName is the name of the target payment terminal.
	TerminalName string `json:"terminalName,omitempty"`

	// ResetConnection forces the terminal cloud connection to be reset while a
	// transactions is in flight. This is a diagnostic settings that can be used
	// only for test transactions.
	ResetConnection bool `json:"resetConnection"`

	// Healthcare contains details for HSA/FSA transactions.
	Healthcare *Healthcare `json:"healthcare,omitempty"`
}

// CaptureRequest contains the information needed to capture a preauth.
type CaptureRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string `json:"transactionRef,omitempty"`

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool `json:"async"`

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool `json:"queue"`

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool `json:"waitForRemovedCard,omitempty"`

	// Force causes a transaction to override any in-progress transactions.
	Force bool `json:"force,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

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

	// AuthResponseCode is the code returned by the terminal or the card issuer
	// to indicate the disposition of the message.
	AuthResponseCode string `json:"authResponseCode,omitempty"`

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

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

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

	// ExpMonth is the card expiration month in MM format.
	ExpMonth string `json:"expMonth,omitempty"`

	// ExpYear is the card expiration year in YY format.
	ExpYear string `json:"expYear,omitempty"`

	// AVSResponse contains address verification results if address information
	// was submitted.
	AVSResponse AVSResponse `json:"avsResponse"`

	// ReceiptSuggestions contains suggested receipt fields.
	ReceiptSuggestions ReceiptSuggestions `json:"receiptSuggestions"`

	// Customer contains customer data, if any. Preserved for reverse
	// compatibility.
	Customer *Customer `json:"customer"`

	// Customers contains customer data, if any.
	Customers []Customer `json:"customers"`
}

// VoidRequest contains a void request.
type VoidRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string `json:"transactionRef,omitempty"`

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool `json:"async"`

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool `json:"queue"`

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool `json:"waitForRemovedCard,omitempty"`

	// Force causes a transaction to override any in-progress transactions.
	Force bool `json:"force,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

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

	// AuthResponseCode is the code returned by the terminal or the card issuer
	// to indicate the disposition of the message.
	AuthResponseCode string `json:"authResponseCode,omitempty"`

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

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

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

	// ExpMonth is the card expiration month in MM format.
	ExpMonth string `json:"expMonth,omitempty"`

	// ExpYear is the card expiration year in YY format.
	ExpYear string `json:"expYear,omitempty"`

	// AVSResponse contains address verification results if address information
	// was submitted.
	AVSResponse AVSResponse `json:"avsResponse"`

	// ReceiptSuggestions contains suggested receipt fields.
	ReceiptSuggestions ReceiptSuggestions `json:"receiptSuggestions"`

	// Customer contains customer data, if any. Preserved for reverse
	// compatibility.
	Customer *Customer `json:"customer"`

	// Customers contains customer data, if any.
	Customers []Customer `json:"customers"`

	// SigFile contains the hex encoded signature data.
	SigFile string `json:"sigFile,omitempty"`
}

// EnrollRequest contains the information needed to enroll a new payment
// method in the token vault.
type EnrollRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string `json:"transactionRef,omitempty"`

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool `json:"async"`

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool `json:"queue"`

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool `json:"waitForRemovedCard,omitempty"`

	// Force causes a transaction to override any in-progress transactions.
	Force bool `json:"force,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

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

	// ResetConnection forces the terminal cloud connection to be reset while a
	// transactions is in flight. This is a diagnostic settings that can be used
	// only for test transactions.
	ResetConnection bool `json:"resetConnection"`

	// EntryMethod is the method by which the payment card was entered (MSR,
	// CHIP, KEYED, etc.).
	EntryMethod string `json:"entryMethod,omitempty"`

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

	// AuthResponseCode is the code returned by the terminal or the card issuer
	// to indicate the disposition of the message.
	AuthResponseCode string `json:"authResponseCode,omitempty"`

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

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

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

	// ExpMonth is the card expiration month in MM format.
	ExpMonth string `json:"expMonth,omitempty"`

	// ExpYear is the card expiration year in YY format.
	ExpYear string `json:"expYear,omitempty"`

	// AVSResponse contains address verification results if address information
	// was submitted.
	AVSResponse AVSResponse `json:"avsResponse"`

	// ReceiptSuggestions contains suggested receipt fields.
	ReceiptSuggestions ReceiptSuggestions `json:"receiptSuggestions"`

	// Customer contains customer data, if any. Preserved for reverse
	// compatibility.
	Customer *Customer `json:"customer"`

	// Customers contains customer data, if any.
	Customers []Customer `json:"customers"`

	// SigFile contains the hex encoded signature data.
	SigFile string `json:"sigFile,omitempty"`
}

// ClearTerminalRequest contains the information needed to enroll a new
// payment method in the token vault.
type ClearTerminalRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string `json:"transactionRef,omitempty"`

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool `json:"async"`

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool `json:"queue"`

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool `json:"waitForRemovedCard,omitempty"`

	// Force causes a transaction to override any in-progress transactions.
	Force bool `json:"force,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// TerminalName is the name of the target payment terminal.
	TerminalName string `json:"terminalName,omitempty"`

	// ResetConnection forces the terminal cloud connection to be reset while a
	// transactions is in flight. This is a diagnostic settings that can be used
	// only for test transactions.
	ResetConnection bool `json:"resetConnection"`
}

// GiftActivateRequest contains the information needed to activate or recharge
// a gift card.
type GiftActivateRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string `json:"transactionRef,omitempty"`

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool `json:"async"`

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool `json:"queue"`

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool `json:"waitForRemovedCard,omitempty"`

	// Force causes a transaction to override any in-progress transactions.
	Force bool `json:"force,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

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

	// ResetConnection forces the terminal cloud connection to be reset while a
	// transactions is in flight. This is a diagnostic settings that can be used
	// only for test transactions.
	ResetConnection bool `json:"resetConnection"`
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

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

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
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string `json:"transactionRef,omitempty"`

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool `json:"async"`

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool `json:"queue"`

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool `json:"waitForRemovedCard,omitempty"`

	// Force causes a transaction to override any in-progress transactions.
	Force bool `json:"force,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// BatchID optional batch id.
	BatchID string `json:"batchId"`
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

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// Sig is the ECC signature of the response. Can be used to ensure that it
	// was signed by the terminal and detect man-in-the middle attacks.
	Sig string `json:"sig,omitempty"`

	// Batches is a collection of batches closed during the batch close
	// operation.
	Batches []BatchSummary `json:"batches"`
}

// TermsAndConditionsRequest contains the fields needed for custom Terms and
// Conditions prompts.
type TermsAndConditionsRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string `json:"transactionRef,omitempty"`

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool `json:"async"`

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool `json:"queue"`

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool `json:"waitForRemovedCard,omitempty"`

	// Force causes a transaction to override any in-progress transactions.
	Force bool `json:"force,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

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

	// ResetConnection forces the terminal cloud connection to be reset while a
	// transactions is in flight. This is a diagnostic settings that can be used
	// only for test transactions.
	ResetConnection bool `json:"resetConnection"`

	// TCAlias is an alias for a Terms and Conditions template configured in the
	// BlockChyp dashboard.
	TCAlias string `json:"tcAlias"`

	// TCName contains the name of the Terms and Conditions the user is
	// accepting.
	TCName string `json:"tcName"`

	// TCContent is the content of the terms and conditions that will be
	// presented to the user.
	TCContent string `json:"tcContent"`

	// ContentHash is a hash of the terms and conditions content that can be used
	// for caching.
	ContentHash string `json:"contentHash"`

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

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

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

	// AuthResponseCode is the code returned by the terminal or the card issuer
	// to indicate the disposition of the message.
	AuthResponseCode string `json:"authResponseCode,omitempty"`

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

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

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

	// Confirmed indicates that the transaction has met the standard criteria for
	// confirmation on the network. (For example, 6 confirmations for level one
	// bitcoin.)
	Confirmed bool `json:"confirmed"`

	// CryptoAuthorizedAmount is the amount submitted to the blockchain.
	CryptoAuthorizedAmount string `json:"cryptoAuthorizedAmount"`

	// CryptoNetworkFee is the network level fee assessed for the transaction
	// denominated in cryptocurrency. This fee goes to channel operators and
	// crypto miners, not BlockChyp.
	CryptoNetworkFee string `json:"cryptoNetworkFee"`

	// Cryptocurrency is the three letter cryptocurrency code used for the
	// transactions.
	Cryptocurrency string `json:"cryptocurrency"`

	// CryptoNetwork indicates whether or not the transaction was processed on
	// the level one or level two network.
	CryptoNetwork string `json:"cryptoNetwork"`

	// CryptoReceiveAddress the address on the crypto network the transaction was
	// sent to.
	CryptoReceiveAddress string `json:"cryptoReceiveAddress"`

	// CryptoBlock hash or other identifier that identifies the block on the
	// cryptocurrency network, if available or relevant.
	CryptoBlock string `json:"cryptoBlock"`

	// CryptoTransactionID hash or other transaction identifier that identifies
	// the transaction on the cryptocurrency network, if available or relevant.
	CryptoTransactionID string `json:"cryptoTransactionId"`

	// CryptoPaymentRequest is the payment request URI used for the transaction,
	// if available.
	CryptoPaymentRequest string `json:"cryptoPaymentRequest"`

	// CryptoStatus is used for additional status information related to crypto
	// transactions.
	CryptoStatus string `json:"cryptoStatus"`

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

	// ExpMonth is the card expiration month in MM format.
	ExpMonth string `json:"expMonth,omitempty"`

	// ExpYear is the card expiration year in YY format.
	ExpYear string `json:"expYear,omitempty"`

	// AVSResponse contains address verification results if address information
	// was submitted.
	AVSResponse AVSResponse `json:"avsResponse"`

	// ReceiptSuggestions contains suggested receipt fields.
	ReceiptSuggestions ReceiptSuggestions `json:"receiptSuggestions"`

	// Customer contains customer data, if any. Preserved for reverse
	// compatibility.
	Customer *Customer `json:"customer"`

	// Customers contains customer data, if any.
	Customers []Customer `json:"customers"`

	// SigFile contains the hex encoded signature data.
	SigFile string `json:"sigFile,omitempty"`

	// WhiteListedCard contains card BIN ranges can be whitelisted so that they
	// are read instead of being processed directly. This is useful for
	// integration with legacy gift card systems.
	WhiteListedCard *WhiteListedCard `json:"whiteListedCard"`

	// StoreAndForward indicates that the transaction was flagged for store and
	// forward due to network problems.
	StoreAndForward bool `json:"storeAndForward"`

	// Status indicates the current status of a transaction.
	Status string `json:"status"`
}

// TransactionStatusRequest models the request for updated information about a
// transaction.
type TransactionStatusRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string `json:"transactionRef,omitempty"`

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool `json:"async"`

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool `json:"queue"`

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool `json:"waitForRemovedCard,omitempty"`

	// Force causes a transaction to override any in-progress transactions.
	Force bool `json:"force,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

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

	// AuthResponseCode is the code returned by the terminal or the card issuer
	// to indicate the disposition of the message.
	AuthResponseCode string `json:"authResponseCode,omitempty"`

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

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

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

	// ExpMonth is the card expiration month in MM format.
	ExpMonth string `json:"expMonth,omitempty"`

	// ExpYear is the card expiration year in YY format.
	ExpYear string `json:"expYear,omitempty"`

	// AVSResponse contains address verification results if address information
	// was submitted.
	AVSResponse AVSResponse `json:"avsResponse"`

	// ReceiptSuggestions contains suggested receipt fields.
	ReceiptSuggestions ReceiptSuggestions `json:"receiptSuggestions"`

	// Customer contains customer data, if any. Preserved for reverse
	// compatibility.
	Customer *Customer `json:"customer"`

	// Customers contains customer data, if any.
	Customers []Customer `json:"customers"`

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

	// UnitCode is an alphanumeric code for units of measurement as used in
	// international trade.
	UnitCode string `json:"unitCode"`

	// CommodityCode is an international description code of the item.
	CommodityCode string `json:"commodityCode"`

	// ProductCode is a merchant-defined description code of the item.
	ProductCode string `json:"productCode"`

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
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string `json:"transactionRef,omitempty"`

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool `json:"async"`

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool `json:"queue"`

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool `json:"waitForRemovedCard,omitempty"`

	// Force causes a transaction to override any in-progress transactions.
	Force bool `json:"force,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// TerminalName is the name of the target payment terminal.
	TerminalName string `json:"terminalName,omitempty"`

	// ResetConnection forces the terminal cloud connection to be reset while a
	// transactions is in flight. This is a diagnostic settings that can be used
	// only for test transactions.
	ResetConnection bool `json:"resetConnection"`

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
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string `json:"transactionRef,omitempty"`

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool `json:"async"`

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool `json:"queue"`

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool `json:"waitForRemovedCard,omitempty"`

	// Force causes a transaction to override any in-progress transactions.
	Force bool `json:"force,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// TerminalName is the name of the target payment terminal.
	TerminalName string `json:"terminalName,omitempty"`

	// ResetConnection forces the terminal cloud connection to be reset while a
	// transactions is in flight. This is a diagnostic settings that can be used
	// only for test transactions.
	ResetConnection bool `json:"resetConnection"`
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

	// CardInSlot indicates whether or not a card is currently in the card slot.
	CardInSlot bool `json:"cardInSlot"`

	// Status contains the operation that the terminal is performing.
	Status string `json:"status"`

	// TransactionRef contains the transaction reference for an ongoing
	// transaction, if one was specified at request time.
	TransactionRef string `json:"transactionRef"`

	// AutogeneratedRef indicates that the transaction reference was
	// autogenerated and should be ignored for the purposes of duplicate
	// detection.
	AutogeneratedRef bool `json:"autogeneratedRef"`

	// Since is the timestamp of the last status change.
	Since time.Time `json:"since"`
}

// PaymentLinkRequest creates a payment link.
type PaymentLinkRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string `json:"transactionRef,omitempty"`

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool `json:"async"`

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool `json:"queue"`

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool `json:"waitForRemovedCard,omitempty"`

	// Force causes a transaction to override any in-progress transactions.
	Force bool `json:"force,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

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

	// ResetConnection forces the terminal cloud connection to be reset while a
	// transactions is in flight. This is a diagnostic settings that can be used
	// only for test transactions.
	ResetConnection bool `json:"resetConnection"`

	// AutoSend automatically send the link via an email.
	AutoSend bool `json:"autoSend"`

	// Enroll indicates that the payment method should be added to the token
	// vault alongside the authorization.
	Enroll bool `json:"enroll,omitempty"`

	// EnrollOnly indicates that the link should be used to enroll a token only.
	// Can only be used in cashier mode.
	EnrollOnly bool `json:"enrollOnly,omitempty"`

	// QrcodeBinary indicates that the QR Code binary should be returned.
	QrcodeBinary bool `json:"qrcodeBinary,omitempty"`

	// QrcodeSize determines the size of the qr code to be returned.
	QrcodeSize int `json:"qrcodeSize,omitempty"`

	// DaysToExpiration number of days until the payment link expires.
	DaysToExpiration int `json:"daysToExpiration,omitempty"`

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

	// Cryptocurrency indicates that the transaction should be a cryptocurrency
	// transaction. Value should be a crypto currency code (ETH, BTC) or ANY to
	// prompt the user to choose from supported cryptocurrencies.
	Cryptocurrency *string `json:"cryptocurrency"`

	// CryptoNetwork is an optional parameter that can be used to force a crypto
	// transaction onto a level one or level two network. Valid values are L1 and
	// L2. Defaults to L1.
	CryptoNetwork *string `json:"cryptoNetwork"`

	// CryptoReceiveAddress can be used to specify a specific receive address for
	// a crypto transaction. Disabled by default. This should only be used by
	// sophisticated users with access to properly configured hot wallets.
	CryptoReceiveAddress *string `json:"cryptoReceiveAddress"`

	// PaymentRequestLabel can optionally add a label to the payment request if
	// the target cryptocurrency supports labels. Defaults to the merchant's DBA
	// Name.
	PaymentRequestLabel *string `json:"paymentRequestLabel"`

	// PaymentRequestMessage can optionally add a message to the payment request
	// if the target cryptocurrency supports labels. Defaults to empty.
	PaymentRequestMessage *string `json:"paymentRequestMessage"`
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

	// LinkCode is the payment link code.
	LinkCode string `json:"linkCode"`

	// URL is the url for the payment link.
	URL string `json:"url"`

	// QrcodeURL is the url for a QR Code associated with this link.
	QrcodeURL string `json:"qrcodeUrl"`

	// QrcodeBinary is the hex encoded binary for the QR Code, if requested.
	// Encoded in PNG format.
	QrcodeBinary string `json:"qrcodeBinary"`

	// CustomerID is the customer id created or used for the payment.
	CustomerID string `json:"customerId"`
}

// CancelPaymentLinkRequest cancels a pending payment link. Payment links that
// have already been used cannot be canceled and the request will be rejected.
type CancelPaymentLinkRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string `json:"transactionRef,omitempty"`

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool `json:"async"`

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool `json:"queue"`

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool `json:"waitForRemovedCard,omitempty"`

	// Force causes a transaction to override any in-progress transactions.
	Force bool `json:"force,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// LinkCode is the payment link code to cancel.
	LinkCode string `json:"linkCode"`
}

// CancelPaymentLinkResponse indicates success or failure of a payment link
// cancellation.
type CancelPaymentLinkResponse struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`
}

// CashDiscountRequest computes the cash discount for a cash discount if
// enabled.
type CashDiscountRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string `json:"transactionRef,omitempty"`

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool `json:"async"`

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool `json:"queue"`

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool `json:"waitForRemovedCard,omitempty"`

	// Force causes a transaction to override any in-progress transactions.
	Force bool `json:"force,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

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

	// RoundingMode indicates how partial pennies should be rounded for
	// calculated values like surcharges. Rounding up is the default behavior.
	RoundingMode *RoundingMode `json:"roundingMode"`
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
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string `json:"transactionRef,omitempty"`

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool `json:"async"`

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool `json:"queue"`

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool `json:"waitForRemovedCard,omitempty"`

	// Force causes a transaction to override any in-progress transactions.
	Force bool `json:"force,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// Query optional search query. Will match amount, last 4 and customer name.
	// batchId and terminalName are not supported with this option.
	Query string `json:"query"`

	// BatchID optional batch id.
	BatchID string `json:"batchId"`

	// TerminalName optional terminal name.
	TerminalName string `json:"terminalName"`

	// StartDate optional start date filter for batch history.
	StartDate time.Time `json:"startDate"`

	// EndDate optional end date filter for batch history.
	EndDate time.Time `json:"endDate"`

	// MaxResults max results to be returned by this request.
	MaxResults int `json:"maxResults"`

	// StartIndex starting index for results to be returned.
	StartIndex int `json:"startIndex"`
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

	// Test indicates that the response came from the test gateway.
	Test bool `json:"test"`

	// BatchID batch identifier if filtered by batch.
	BatchID string `json:"batchId"`

	// TerminalName terminal name if filtered by terminal.
	TerminalName string `json:"terminalName"`

	// StartDate start date if filtered by start date.
	StartDate time.Time `json:"startDate"`

	// EndDate end date if filtered by end date.
	EndDate time.Time `json:"endDate"`

	// MaxResults max results from the original request echoed back. Defaults to
	// the system max of 250.
	MaxResults int `json:"maxResults"`

	// StartIndex starting index from the original request echoed back.
	StartIndex int `json:"startIndex"`

	// TotalResultCount total number of results accessible through paging.
	TotalResultCount int `json:"totalResultCount"`

	// Transactions matching transaction history.
	Transactions []AuthorizationResponse `json:"transactions"`
}

// BatchHistoryRequest models a batch history request.
type BatchHistoryRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string `json:"transactionRef,omitempty"`

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool `json:"async"`

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool `json:"queue"`

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool `json:"waitForRemovedCard,omitempty"`

	// Force causes a transaction to override any in-progress transactions.
	Force bool `json:"force,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// StartDate optional start date filter for batch history.
	StartDate time.Time `json:"startDate"`

	// EndDate optional end date filter for batch history.
	EndDate time.Time `json:"endDate"`

	// MaxResults max results to be returned by this request. Defaults to the
	// system max of 250.
	MaxResults int `json:"maxResults"`

	// StartIndex starting index for results to be returned.
	StartIndex int `json:"startIndex"`
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

	// Test indicates that the response came from the test gateway.
	Test bool `json:"test"`

	// StartDate start date if filtered by start date.
	StartDate time.Time `json:"startDate"`

	// EndDate end date if filtered by end date.
	EndDate time.Time `json:"endDate"`

	// Batches merchant's batch history in descending order.
	Batches []BatchSummary `json:"batches"`

	// MaxResults max results from the original request echoed back.
	MaxResults int `json:"maxResults"`

	// StartIndex starting index from the original request echoed back.
	StartIndex int `json:"startIndex"`

	// TotalResultCount total number of results accessible through paging.
	TotalResultCount int `json:"totalResultCount"`
}

// BatchSummary models high level information about a single batch.
type BatchSummary struct {
	// BatchID batch identifier.
	BatchID string `json:"batchId"`

	// EntryMethod entry method for the batch, if any.
	EntryMethod string `json:"entryMethod"`

	// DestinationAccountID merchant deposit account into which proceeds should
	// be deposited.
	DestinationAccountID string `json:"destinationAccountId"`

	// CapturedAmount is the new captured amount.
	CapturedAmount string `json:"capturedAmount"`

	// OpenPreauths is the amount of preauths opened during the batch that have
	// not been captured.
	OpenPreauths string `json:"openPreauths"`

	// CurrencyCode is the currency the batch was settled in.
	CurrencyCode string `json:"currencyCode"`

	// Open flag indicating whether or not the batch is open.
	Open bool `json:"open"`

	// OpenDate date and time of the first transaction for this batch.
	OpenDate time.Time `json:"openDate"`

	// CloseDate date and time the batch was closed.
	CloseDate time.Time `json:"closeDate"`
}

// BatchDetailsRequest models a request for details about a single batch.
type BatchDetailsRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string `json:"transactionRef,omitempty"`

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool `json:"async"`

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool `json:"queue"`

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool `json:"waitForRemovedCard,omitempty"`

	// Force causes a transaction to override any in-progress transactions.
	Force bool `json:"force,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// BatchID id for the batch to be retrieved.
	BatchID string `json:"batchId"`
}

// BatchDetailsResponse models a response for details about a single batch.
type BatchDetailsResponse struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// Test indicates that the response came from the test gateway.
	Test bool `json:"test"`

	// BatchID batch identifier.
	BatchID string `json:"batchId"`

	// EntryMethod entry method for the batch, if any.
	EntryMethod string `json:"entryMethod"`

	// DestinationAccountID merchant deposit account into which proceeds should
	// be deposited.
	DestinationAccountID string `json:"destinationAccountId"`

	// CapturedAmount is the new captured amount.
	CapturedAmount string `json:"capturedAmount"`

	// OpenPreauths preauths from this batch still open.
	OpenPreauths string `json:"openPreauths"`

	// TotalVolume is the total volume from this batch.
	TotalVolume string `json:"totalVolume"`

	// TransactionCount is the total number of transactions in this batch.
	TransactionCount int `json:"transactionCount"`

	// GiftCardsSold is the total volume of gift cards sold.
	GiftCardsSold string `json:"giftCardsSold"`

	// GiftCardVolume is the total volume of gift cards transactions.
	GiftCardVolume string `json:"giftCardVolume"`

	// ExpectedDeposit is the expected volume for this batch, usually captured
	// volume less gift card volume.
	ExpectedDeposit string `json:"expectedDeposit"`

	// Open flag indicating whether or not the batch is open.
	Open bool `json:"open"`

	// OpenDate date and time of the first transaction for this batch.
	OpenDate time.Time `json:"openDate"`

	// CloseDate date and time the batch was closed.
	CloseDate time.Time `json:"closeDate"`

	// VolumeByTerminal merchant's batch history in descending order.
	VolumeByTerminal []TerminalVolume `json:"volumeByTerminal"`
}

// TerminalVolume models transaction volume for a single terminal.
type TerminalVolume struct {
	// TerminalName is the terminal name assigned during activation.
	TerminalName string `json:"terminalName"`

	// SerialNumber is the manufacturer's serial number.
	SerialNumber string `json:"serialNumber"`

	// TerminalType is the terminal type.
	TerminalType string `json:"terminalType"`

	// CapturedAmount is the captured amount.
	CapturedAmount string `json:"capturedAmount"`

	// TransactionCount is the number of transactions run on this terminal.
	TransactionCount int `json:"transactionCount"`
}

// AddTestMerchantRequest models basic information needed to create a test
// merchant.
type AddTestMerchantRequest struct {
	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// DBAName the DBA name for the test merchant.
	DBAName string `json:"dbaName"`

	// CompanyName is the corporate name for the test merchant.
	CompanyName string `json:"companyName"`

	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`
}

// MerchantProfileRequest models a request for information about the merchant
// profile.
type MerchantProfileRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string `json:"transactionRef,omitempty"`

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool `json:"async"`

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool `json:"queue"`

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool `json:"waitForRemovedCard,omitempty"`

	// Force causes a transaction to override any in-progress transactions.
	Force bool `json:"force,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// MerchantID is the merchant id. Optional for merchant scoped requests.
	MerchantID string `json:"merchantId"`
}

// MerchantPlatformRequest models a request related to a platform
// configuration.
type MerchantPlatformRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string `json:"transactionRef,omitempty"`

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool `json:"async"`

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool `json:"queue"`

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool `json:"waitForRemovedCard,omitempty"`

	// Force causes a transaction to override any in-progress transactions.
	Force bool `json:"force,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// PlatformID is the platform configuration id.
	PlatformID string `json:"platformId"`
}

// InviteMerchantUserRequest models a request for adding a new user to a
// merchant account.
type InviteMerchantUserRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string `json:"transactionRef,omitempty"`

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool `json:"async"`

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool `json:"queue"`

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool `json:"waitForRemovedCard,omitempty"`

	// Force causes a transaction to override any in-progress transactions.
	Force bool `json:"force,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// MerchantID is the merchant id. Optional for merchant scoped requests.
	MerchantID string `json:"merchantId"`

	// Email is the email address of the user.
	Email string `json:"email"`

	// FirstName is the first name of the new user.
	FirstName string `json:"firstName"`

	// LastName is the last name of the new user.
	LastName string `json:"lastName"`

	// Roles an optional array of role codes that will be assigned to the user.
	// If omitted defaults to the default merchant role.
	Roles []string `json:"roles"`
}

// Address models a physical address.
type Address struct {
	// Address1 is the first line of the street address.
	Address1 string `json:"address1"`

	// Address2 is the second line of the street address.
	Address2 string `json:"address2"`

	// City is the city associated with the street address.
	City string `json:"city"`

	// StateOrProvince is the state or province associated with the street
	// address.
	StateOrProvince string `json:"stateOrProvince"`

	// PostalCode is the postal code associated with the street address.
	PostalCode string `json:"postalCode"`

	// CountryCode is the ISO country code associated with the street address.
	CountryCode string `json:"countryCode"`

	// Latitude is the latitude component of the address's GPS coordinates.
	Latitude float64 `json:"latitude"`

	// Longitude is the longitude component of the address's GPS coordinates.
	Longitude float64 `json:"longitude"`
}

// MerchantProfile models a merchant profile.
type MerchantProfile struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test indicates that the response came from the test gateway.
	Test bool `json:"test"`

	// MerchantID is the merchant id.
	MerchantID string `json:"merchantId"`

	// CompanyName is the merchant's company name.
	CompanyName string `json:"companyName"`

	// DBAName is the dba name of the merchant.
	DBAName string `json:"dbaName"`

	// InvoiceName is the name the merchant prefers on payment link invoices.
	InvoiceName string `json:"invoiceName"`

	// ContactName is the contact name for the merchant.
	ContactName string `json:"contactName"`

	// ContactNumber is the contact number for the merchant.
	ContactNumber string `json:"contactNumber"`

	// LocationName is the location name.
	LocationName string `json:"locationName"`

	// StoreNumber is the store number.
	StoreNumber string `json:"storeNumber"`

	// PartnerRef is the partner assigne reference for this merchant.
	PartnerRef string `json:"partnerRef"`

	// TimeZone is the merchant's local time zone.
	TimeZone string `json:"timeZone"`

	// BatchCloseTime is the batch close time in the merchant's time zone.
	BatchCloseTime string `json:"batchCloseTime"`

	// TerminalUpdateTime is the terminal firmware update time.
	TerminalUpdateTime string `json:"terminalUpdateTime"`

	// AutoBatchClose flag indicating whether or not the batch automatically
	// closes.
	AutoBatchClose bool `json:"autoBatchClose"`

	// DisableBatchEmails flag indicating whether or not batch closure emails
	// should be automatically sent.
	DisableBatchEmails bool `json:"disableBatchEmails"`

	// PINEnabled flag indicating whether or not pin entry is enabled.
	PINEnabled bool `json:"pinEnabled"`

	// CashBackEnabled flag indicating whether or not cash back is enabled.
	CashBackEnabled bool `json:"cashBackEnabled"`

	// StoreAndForwardEnabled flag indicating whether or not store and forward is
	// enabled.
	StoreAndForwardEnabled bool `json:"storeAndForwardEnabled"`

	// PartialAuthEnabled flag indicating whether or not partial authorizations
	// are supported for this merchant.
	PartialAuthEnabled bool `json:"partialAuthEnabled"`

	// SplitBankAccountsEnabled flag indicating whether or not this merchant
	// support split settlement.
	SplitBankAccountsEnabled bool `json:"splitBankAccountsEnabled"`

	// StoreAndForwardFloorLimit floor limit for store and forward transactions.
	StoreAndForwardFloorLimit string `json:"storeAndForwardFloorLimit"`

	// PublicKey is the blockchyp public key for this merchant.
	PublicKey string `json:"publicKey"`

	// Status is the underwriting/processing status for the the merchant.
	Status string `json:"status"`

	// CashDiscountEnabled enables cash discount or surcharging.
	CashDiscountEnabled bool `json:"cashDiscountEnabled"`

	// SurveyTimeout is the post transaction survey timeout in seconds.
	SurveyTimeout int `json:"surveyTimeout"`

	// CooldownTimeout is time a transaction result is displayed on a terminal
	// before the terminal is automatically cleared in seconds.
	CooldownTimeout int `json:"cooldownTimeout"`

	// TipEnabled indicates that tips are enabled for a merchant account.
	TipEnabled bool `json:"tipEnabled"`

	// PromptForTip indicates that tips should be automatically prompted for
	// after charge and preauth transactions.
	PromptForTip bool `json:"promptForTip"`

	// TipDefaults three default values for tips. Can be provided as a percentage
	// if a percent sign is provided. Otherwise the values are assumed to be
	// basis points.
	TipDefaults []string `json:"tipDefaults"`

	// CashbackPresets four default values for cashback prompts.
	CashbackPresets []string `json:"cashbackPresets"`

	// EBTEnabled indicates that EBT cards are enabled.
	EBTEnabled bool `json:"ebtEnabled"`

	// FreeRangeRefundsEnabled indicates that refunds without transaction
	// references are permitted.
	FreeRangeRefundsEnabled bool `json:"freeRangeRefundsEnabled"`

	// PINBypassEnabled indicates that pin bypass is enabled.
	PINBypassEnabled bool `json:"pinBypassEnabled"`

	// GiftCardsDisabled indicates that gift cards are disabled.
	GiftCardsDisabled bool `json:"giftCardsDisabled"`

	// TCDisabled disables terms and conditions pages in the merchant UI.
	TCDisabled bool `json:"tcDisabled"`

	// DigitalSignaturesEnabled indicates that digital signature capture is
	// enabled.
	DigitalSignaturesEnabled bool `json:"digitalSignaturesEnabled"`

	// DigitalSignatureReversal indicates that transactions should auto-reverse
	// when signatures are refused.
	DigitalSignatureReversal bool `json:"digitalSignatureReversal"`

	// BillingAddress is the address to be used for billing correspondence.
	BillingAddress Address `json:"billingAddress"`

	// ShippingAddress is the address to be used for shipping.
	ShippingAddress Address `json:"shippingAddress"`

	// Visa indicates that Visa cards are supported.
	Visa bool `json:"visa"`

	// MasterCard indicates that MasterCard is supported.
	MasterCard bool `json:"masterCard"`

	// AMEX indicates that American Express is supported.
	AMEX bool `json:"amex"`

	// Discover indicates that Discover cards are supported.
	Discover bool `json:"discover"`

	// JCB indicates that JCB (Japan Card Bureau) cards are supported.
	JCB bool `json:"jcb"`

	// UnionPay indicates that China Union Pay cards are supported.
	UnionPay bool `json:"unionPay"`

	// ContactlessEMV indicates that contactless EMV cards are supported.
	ContactlessEMV bool `json:"contactlessEmv"`

	// ManualEntryEnabled indicates that manual card entry is enabled.
	ManualEntryEnabled bool `json:"manualEntryEnabled"`

	// ManualEntryPromptZip requires a zip code to be entered for manually
	// entered transactions.
	ManualEntryPromptZip bool `json:"manualEntryPromptZip"`

	// ManualEntryPromptStreetNumber requires a street number to be entered for
	// manually entered transactions.
	ManualEntryPromptStreetNumber bool `json:"manualEntryPromptStreetNumber"`

	// GatewayOnly indicates that this merchant is boarded on BlockChyp in
	// gateway only mode.
	GatewayOnly bool `json:"gatewayOnly"`

	// BankAccounts bank accounts for split bank account merchants.
	BankAccounts []BankAccount `json:"bankAccounts"`
}

// MerchantProfileResponse models a response for a single merchant profile.
type MerchantProfileResponse struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// Test indicates that the response came from the test gateway.
	Test bool `json:"test"`

	// MerchantID is the merchant id.
	MerchantID string `json:"merchantId"`

	// CompanyName is the merchant's company name.
	CompanyName string `json:"companyName"`

	// DBAName is the dba name of the merchant.
	DBAName string `json:"dbaName"`

	// InvoiceName is the name the merchant prefers on payment link invoices.
	InvoiceName string `json:"invoiceName"`

	// ContactName is the contact name for the merchant.
	ContactName string `json:"contactName"`

	// ContactNumber is the contact number for the merchant.
	ContactNumber string `json:"contactNumber"`

	// LocationName is the location name.
	LocationName string `json:"locationName"`

	// StoreNumber is the store number.
	StoreNumber string `json:"storeNumber"`

	// PartnerRef is the partner assigne reference for this merchant.
	PartnerRef string `json:"partnerRef"`

	// TimeZone is the merchant's local time zone.
	TimeZone string `json:"timeZone"`

	// BatchCloseTime is the batch close time in the merchant's time zone.
	BatchCloseTime string `json:"batchCloseTime"`

	// TerminalUpdateTime is the terminal firmware update time.
	TerminalUpdateTime string `json:"terminalUpdateTime"`

	// AutoBatchClose flag indicating whether or not the batch automatically
	// closes.
	AutoBatchClose bool `json:"autoBatchClose"`

	// DisableBatchEmails flag indicating whether or not batch closure emails
	// should be automatically sent.
	DisableBatchEmails bool `json:"disableBatchEmails"`

	// PINEnabled flag indicating whether or not pin entry is enabled.
	PINEnabled bool `json:"pinEnabled"`

	// CashBackEnabled flag indicating whether or not cash back is enabled.
	CashBackEnabled bool `json:"cashBackEnabled"`

	// StoreAndForwardEnabled flag indicating whether or not store and forward is
	// enabled.
	StoreAndForwardEnabled bool `json:"storeAndForwardEnabled"`

	// PartialAuthEnabled flag indicating whether or not partial authorizations
	// are supported for this merchant.
	PartialAuthEnabled bool `json:"partialAuthEnabled"`

	// SplitBankAccountsEnabled flag indicating whether or not this merchant
	// support split settlement.
	SplitBankAccountsEnabled bool `json:"splitBankAccountsEnabled"`

	// StoreAndForwardFloorLimit floor limit for store and forward transactions.
	StoreAndForwardFloorLimit string `json:"storeAndForwardFloorLimit"`

	// PublicKey is the blockchyp public key for this merchant.
	PublicKey string `json:"publicKey"`

	// Status is the underwriting/processing status for the the merchant.
	Status string `json:"status"`

	// CashDiscountEnabled enables cash discount or surcharging.
	CashDiscountEnabled bool `json:"cashDiscountEnabled"`

	// SurveyTimeout is the post transaction survey timeout in seconds.
	SurveyTimeout int `json:"surveyTimeout"`

	// CooldownTimeout is time a transaction result is displayed on a terminal
	// before the terminal is automatically cleared in seconds.
	CooldownTimeout int `json:"cooldownTimeout"`

	// TipEnabled indicates that tips are enabled for a merchant account.
	TipEnabled bool `json:"tipEnabled"`

	// PromptForTip indicates that tips should be automatically prompted for
	// after charge and preauth transactions.
	PromptForTip bool `json:"promptForTip"`

	// TipDefaults three default values for tips. Can be provided as a percentage
	// if a percent sign is provided. Otherwise the values are assumed to be
	// basis points.
	TipDefaults []string `json:"tipDefaults"`

	// CashbackPresets four default values for cashback prompts.
	CashbackPresets []string `json:"cashbackPresets"`

	// EBTEnabled indicates that EBT cards are enabled.
	EBTEnabled bool `json:"ebtEnabled"`

	// FreeRangeRefundsEnabled indicates that refunds without transaction
	// references are permitted.
	FreeRangeRefundsEnabled bool `json:"freeRangeRefundsEnabled"`

	// PINBypassEnabled indicates that pin bypass is enabled.
	PINBypassEnabled bool `json:"pinBypassEnabled"`

	// GiftCardsDisabled indicates that gift cards are disabled.
	GiftCardsDisabled bool `json:"giftCardsDisabled"`

	// TCDisabled disables terms and conditions pages in the merchant UI.
	TCDisabled bool `json:"tcDisabled"`

	// DigitalSignaturesEnabled indicates that digital signature capture is
	// enabled.
	DigitalSignaturesEnabled bool `json:"digitalSignaturesEnabled"`

	// DigitalSignatureReversal indicates that transactions should auto-reverse
	// when signatures are refused.
	DigitalSignatureReversal bool `json:"digitalSignatureReversal"`

	// BillingAddress is the address to be used for billing correspondence.
	BillingAddress Address `json:"billingAddress"`

	// ShippingAddress is the address to be used for shipping.
	ShippingAddress Address `json:"shippingAddress"`

	// Visa indicates that Visa cards are supported.
	Visa bool `json:"visa"`

	// MasterCard indicates that MasterCard is supported.
	MasterCard bool `json:"masterCard"`

	// AMEX indicates that American Express is supported.
	AMEX bool `json:"amex"`

	// Discover indicates that Discover cards are supported.
	Discover bool `json:"discover"`

	// JCB indicates that JCB (Japan Card Bureau) cards are supported.
	JCB bool `json:"jcb"`

	// UnionPay indicates that China Union Pay cards are supported.
	UnionPay bool `json:"unionPay"`

	// ContactlessEMV indicates that contactless EMV cards are supported.
	ContactlessEMV bool `json:"contactlessEmv"`

	// ManualEntryEnabled indicates that manual card entry is enabled.
	ManualEntryEnabled bool `json:"manualEntryEnabled"`

	// ManualEntryPromptZip requires a zip code to be entered for manually
	// entered transactions.
	ManualEntryPromptZip bool `json:"manualEntryPromptZip"`

	// ManualEntryPromptStreetNumber requires a street number to be entered for
	// manually entered transactions.
	ManualEntryPromptStreetNumber bool `json:"manualEntryPromptStreetNumber"`

	// GatewayOnly indicates that this merchant is boarded on BlockChyp in
	// gateway only mode.
	GatewayOnly bool `json:"gatewayOnly"`

	// BankAccounts bank accounts for split bank account merchants.
	BankAccounts []BankAccount `json:"bankAccounts"`
}

// BankAccount models meta data about a merchant bank account.
type BankAccount struct {
	// ID is the account identifier to be used with authorization requests.
	ID string `json:"id"`

	// Name is the name of the account.
	Name string `json:"name"`

	// Purpose is the purpose of the account.
	Purpose string `json:"purpose"`

	// MaskedAccountNumber is the masked account number.
	MaskedAccountNumber string `json:"maskedAccountNumber"`
}

// ListQueuedTransactionsRequest returns a list of queued transactions on a
// terminal.
type ListQueuedTransactionsRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string `json:"transactionRef,omitempty"`

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool `json:"async"`

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool `json:"queue"`

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool `json:"waitForRemovedCard,omitempty"`

	// Force causes a transaction to override any in-progress transactions.
	Force bool `json:"force,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// TerminalName is the name of the target payment terminal.
	TerminalName string `json:"terminalName,omitempty"`

	// ResetConnection forces the terminal cloud connection to be reset while a
	// transactions is in flight. This is a diagnostic settings that can be used
	// only for test transactions.
	ResetConnection bool `json:"resetConnection"`
}

// ListQueuedTransactionsResponse contains a list of queued transactions on a
// terminal.
type ListQueuedTransactionsResponse struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// TransactionRefs is a list of queued transactions on the terminal.
	TransactionRefs []string `json:"transactionRefs"`
}

// DeleteQueuedTransactionRequest deletes one or all transactions from a
// terminal queue.
type DeleteQueuedTransactionRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool `json:"async"`

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool `json:"queue"`

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool `json:"waitForRemovedCard,omitempty"`

	// Force causes a transaction to override any in-progress transactions.
	Force bool `json:"force,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// TerminalName is the name of the target payment terminal.
	TerminalName string `json:"terminalName,omitempty"`

	// ResetConnection forces the terminal cloud connection to be reset while a
	// transactions is in flight. This is a diagnostic settings that can be used
	// only for test transactions.
	ResetConnection bool `json:"resetConnection"`

	// TransactionRef contains a transaction reference string of the transaction
	// to delete. Passing `*` will clear all queued transactions.
	TransactionRef string `json:"transactionRef"`
}

// DeleteQueuedTransactionResponse is the response to a delete queued
// transaction request.
type DeleteQueuedTransactionResponse struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`
}

// DeleteCustomerRequest deletes a customer record.
type DeleteCustomerRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string `json:"transactionRef,omitempty"`

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool `json:"async"`

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool `json:"queue"`

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool `json:"waitForRemovedCard,omitempty"`

	// Force causes a transaction to override any in-progress transactions.
	Force bool `json:"force,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// CustomerID the ID of the customer to delete.
	CustomerID string `json:"customerId"`
}

// DeleteCustomerResponse is the response to a delete customer request.
type DeleteCustomerResponse struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`
}

// DeleteTokenRequest deletes a payment token.
type DeleteTokenRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string `json:"transactionRef,omitempty"`

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool `json:"async"`

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool `json:"queue"`

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool `json:"waitForRemovedCard,omitempty"`

	// Force causes a transaction to override any in-progress transactions.
	Force bool `json:"force,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// Token the token to delete.
	Token string `json:"token"`
}

// DeleteTokenResponse is the response to a delete token request.
type DeleteTokenResponse struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`
}

// LinkTokenRequest links a payment token with a customer record.
type LinkTokenRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string `json:"transactionRef,omitempty"`

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool `json:"async"`

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool `json:"queue"`

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool `json:"waitForRemovedCard,omitempty"`

	// Force causes a transaction to override any in-progress transactions.
	Force bool `json:"force,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// Token the token to delete.
	Token string `json:"token"`

	// CustomerID BlockChyp assigned customer id.
	CustomerID string `json:"customerId"`
}

// UnlinkTokenRequest removes a link between a payment token with a customer
// record, if one exists.
type UnlinkTokenRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string `json:"transactionRef,omitempty"`

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool `json:"async"`

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool `json:"queue"`

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool `json:"waitForRemovedCard,omitempty"`

	// Force causes a transaction to override any in-progress transactions.
	Force bool `json:"force,omitempty"`

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string `json:"orderRef,omitempty"`

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string `json:"destinationAccount,omitempty"`

	// Token the token to delete.
	Token string `json:"token"`

	// CustomerID BlockChyp assigned customer id.
	CustomerID string `json:"customerId"`
}

// Healthcare contains fields for HSA/FSA transactions.
type Healthcare struct {
	// Types is a list of healthcare categories in the transaction.
	Types []HealthcareGroup `json:"types"`

	// IIASVerified indicates that the purchased items were verified against an
	// Inventory Information Approval System (IIAS).
	IIASVerified bool `json:"iiasVerified"`

	// IIASExempt indicates that the transaction is exempt from IIAS
	// verification.
	IIASExempt bool `json:"iiasExempt"`
}

// HealthcareGroup is a group of fields for a specific type of healthcare.
type HealthcareGroup struct {
	// Type the type of healthcare cost.
	Type HealthcareType `json:"type"`

	// Amount is the amount of this type.
	Amount string `json:"amount"`

	// ProviderID the provider ID used for Mastercard and Discover IIAS requests.
	ProviderID string `json:"providerId"`

	// ServiceTypeCode the service type code used for Mastercard and Discover
	// IIAS requests.
	ServiceTypeCode string `json:"serviceTypeCode"`

	// PayerOrCarrierID thr payer ID/carrier ID used for Mastercard and Discover
	// IIAS requests.
	PayerOrCarrierID string `json:"payerOrCarrierId"`

	// ApprovalRejectReasonCode the approval or reject reason code used for
	// Mastercard and Discover IIAS requests.
	ApprovalRejectReasonCode string `json:"approvalRejectReasonCode"`
}

// GetMerchantsRequest models a request for merchant information.
type GetMerchantsRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test indicates whether or not to return test or live merchants.
	Test bool `json:"test"`

	// MaxResults max to be returned in a single page. Defaults to the system max
	// of 250.
	MaxResults int `json:"maxResults"`

	// StartIndex starting index for paged results. Defaults to zero.
	StartIndex int `json:"startIndex"`
}

// GetMerchantsResponse contains the results for a merchant list request.
type GetMerchantsResponse struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// Test indicates whether or not these results are for test or live
	// merchants.
	Test bool `json:"test"`

	// MaxResults max to be returned in a single page. Defaults to the system max
	// of 250.
	MaxResults int `json:"maxResults"`

	// StartIndex starting index for paged results. Defaults to zero.
	StartIndex int `json:"startIndex"`

	// ResultCount total number of results accessible through paging.
	ResultCount int `json:"resultCount"`

	// Merchants merchants in the current page of results.
	Merchants []MerchantProfileResponse `json:"merchants"`
}

// MerchantUsersResponse contains the results for a merchant users list.
type MerchantUsersResponse struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// Test indicates whether or not these results are for test or live
	// merchants.
	Test bool `json:"test"`

	// Results users and pending invites associated with the merchant.
	Results []MerchantUser `json:"results"`
}

// MerchantUser contains details about a merchant user.
type MerchantUser struct {
	// Test indicates whether or not these results are for test or live
	// merchants.
	Test bool `json:"test"`

	// ID is the user's primary key.
	ID string `json:"id"`

	// FirstName is the user's first name.
	FirstName string `json:"firstName"`

	// LastName is the user's last name.
	LastName string `json:"lastName"`

	// Email is the user's email address.
	Email string `json:"email"`

	// Status is the user account status.
	Status string `json:"status"`

	// Type is the type of user account.
	Type string `json:"type"`

	// Roles are the role codes assigned to this user.
	Roles []string `json:"roles"`

	// Locked indicates whether or not this user account is locked.
	Locked bool `json:"locked"`
}

// MerchantPlatformsResponse contains the results for a merchant platforms
// inquiry.
type MerchantPlatformsResponse struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// Test indicates whether or not these results are for test or live
	// merchants.
	Test bool `json:"test"`

	// Results enumerates merchant platform settings.
	Results []MerchantPlatform `json:"results"`
}

// MerchantPlatform contains details about a merchant board platform
// configuration.
type MerchantPlatform struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// ID primary identifier for a given platform configuration.
	ID string `json:"id"`

	// Disabled indicates that a platform configuration is disabled.
	Disabled bool `json:"disabled"`

	// PlatformCode is BlockChyp's code for the boarding platform.
	PlatformCode string `json:"platformCode"`

	// Priority is the platform's priority in a multi platform setup.
	Priority int `json:"priority"`

	// RegistrationID is an optional field specifying the merchant's card brand
	// registration record.
	RegistrationID string `json:"registrationId"`

	// MerchantID is the merchant's primary identifier.
	MerchantID string `json:"merchantId"`

	// AcquirerMid specifies the merchant id assigned by the acquiring bank.
	AcquirerMid string `json:"acquirerMid"`

	// Notes free form notes description the purpose or intent behind the
	// platform configuration.
	Notes string `json:"notes"`

	// EntryMethod is the optional entry method code if a platform should only be
	// used for specific entry methods. Leave blank for 'all'.
	EntryMethod string `json:"entryMethod"`

	// DateCreated is the date the platform configuration was first created.
	DateCreated string `json:"dateCreated"`

	// LastChange is the date the platform configuration was last modified.
	LastChange string `json:"lastChange"`

	// ConfigMap is a map of configuration values specific to the boarding
	// platform. These are not published. Contact your BlockChyp rep for
	// supported values.
	ConfigMap map[string]string `json:"configMap,omitempty"`
}

// TerminalProfileRequest models a terminal profile request.
type TerminalProfileRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`
}

// TerminalProfileResponse models a terminal profile response.
type TerminalProfileResponse struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// Results enumerates all terminal profiles in the response.
	Results []TerminalProfile `json:"results"`
}

// TerminalDeactivationRequest models a terminal deactivation request.
type TerminalDeactivationRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TerminalName is the terminal name assigned to the terminal.
	TerminalName string `json:"terminalName"`

	// TerminalID is the id assigned by BlockChyp to the terminal.
	TerminalID string `json:"terminalId"`
}

// TerminalActivationRequest models a terminal activation request.
type TerminalActivationRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// MerchantID is the optional merchant id.
	MerchantID string `json:"merchantId"`

	// ActivationCode is the terminal activation code displayed on the terminal.
	ActivationCode string `json:"activationCode"`

	// TerminalName is the name to be assigned to the terminal. Must be unique
	// for the merchant account.
	TerminalName string `json:"terminalName"`

	// CloudRelay indicates that the terminal should be activated in cloud relay
	// mode.
	CloudRelay bool `json:"cloudRelay"`
}

// TerminalProfile contains details about a merchant board platform
// configuration.
type TerminalProfile struct {
	// ID primary identifier for a given terminal.
	ID string `json:"id"`

	// IPAddress is the terminal's local IP address.
	IPAddress string `json:"ipAddress"`

	// TerminalName is the name assigned to the terminal during activation.
	TerminalName string `json:"terminalName"`

	// TerminalType is the terminal type.
	TerminalType string `json:"terminalType"`

	// TerminalTypeDisplayString is the terminal type display string.
	TerminalTypeDisplayString string `json:"terminalTypeDisplayString"`

	// BlockChypFirmwareVersion is the current firmware version deployed on the
	// terminal.
	BlockChypFirmwareVersion string `json:"blockChypFirmwareVersion"`

	// CloudBased indicates whether or not the terminal is configured for cloud
	// relay.
	CloudBased bool `json:"cloudBased"`

	// PublicKey is the terminal's elliptic curve public key.
	PublicKey string `json:"publicKey"`

	// SerialNumber is the manufacturer's serial number.
	SerialNumber string `json:"serialNumber"`

	// Online indicates whether or not the terminal is currently online.
	Online bool `json:"online"`

	// Since the date and time the terminal was first brought online.
	Since string `json:"since"`

	// TotalMemory is the total memory on the terminal.
	TotalMemory int `json:"totalMemory"`

	// TotalStorage is the storage on the terminal.
	TotalStorage int `json:"totalStorage"`

	// AvailableMemory is the available (unused) memory on the terminal.
	AvailableMemory int `json:"availableMemory"`

	// AvailableStorage is the available (unused) storage on the terminal.
	AvailableStorage int `json:"availableStorage"`

	// UsedMemory is the memory currently in use on the terminal.
	UsedMemory int `json:"usedMemory"`

	// UsedStorage is the storage currently in use on the terminal.
	UsedStorage int `json:"usedStorage"`

	// BrandingPreview is the branding asset currently displayed on the terminal.
	BrandingPreview string `json:"brandingPreview"`

	// GroupID the id of the terminal group to which the terminal belongs, if
	// any.
	GroupID string `json:"groupId"`

	// GroupName the name of the terminal group to which the terminal belongs, if
	// any.
	GroupName string `json:"groupName"`
}

// TermsAndConditionsTemplate models a full terms and conditions template.
type TermsAndConditionsTemplate struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// ID primary identifier for a given template.
	ID string `json:"id"`

	// Alias is an alias or code used to refer to a template.
	Alias string `json:"alias"`

	// Name is the name of the template. Displayed as the agreement title on the
	// terminal.
	Name string `json:"name"`

	// Content is the full text of the agreement template.
	Content string `json:"content"`
}

// TermsAndConditionsTemplateRequest models a request to retrieve or
// manipulate terms and conditions data.
type TermsAndConditionsTemplateRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// TemplateID id of a single template.
	TemplateID string `json:"templateId"`
}

// TermsAndConditionsTemplateResponse models a set of templates responsive to
// a request.
type TermsAndConditionsTemplateResponse struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// Results results responsive to a request.
	Results []TermsAndConditionsTemplate `json:"results"`

	// Timeout is an optional timeout override.
	Timeout int `json:"timeout"`
}

// TermsAndConditionsLogRequest models a Terms and Conditions history request.
type TermsAndConditionsLogRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// LogEntryID is the identifier of the log entry to be returned for single
	// result requests.
	LogEntryID string `json:"logEntryId"`

	// TransactionID optional transaction id if only log entries related to a
	// transaction should be returned.
	TransactionID string `json:"transactionId"`

	// MaxResults max to be returned in a single page. Defaults to the system max
	// of 250.
	MaxResults int `json:"maxResults"`

	// StartIndex starting index for paged results. Defaults to zero.
	StartIndex int `json:"startIndex"`

	// StartDate is an optional start date for filtering response data.
	StartDate string `json:"startDate"`

	// EndDate is an optional end date for filtering response data.
	EndDate string `json:"endDate"`
}

// TermsAndConditionsLogResponse models a Terms and Conditions history
// request.
type TermsAndConditionsLogResponse struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// TransactionID optional transaction id if only log entries related to a
	// transaction should be returned.
	TransactionID string `json:"transactionId"`

	// MaxResults max to be returned in a single page. Defaults to the system max
	// of 250.
	MaxResults int `json:"maxResults"`

	// StartIndex starting index for paged results. Defaults to zero.
	StartIndex int `json:"startIndex"`

	// ResultCount total number of results accessible through paging.
	ResultCount int `json:"resultCount"`

	// Results is the full result set responsive to the original request, subject
	// to pagination limits.
	Results []TermsAndConditionsLogEntry `json:"results"`
}

// TermsAndConditionsLogEntry models a Terms and Conditions log entry.
type TermsAndConditionsLogEntry struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// ID internal id for a Terms and Conditions entry.
	ID string `json:"id"`

	// TerminalID id of the terminal that captured this terms and conditions
	// entry.
	TerminalID string `json:"terminalId"`

	// TerminalName name of the terminal that captured this terms and conditions
	// entry.
	TerminalName string `json:"terminalName"`

	// Test is a flag indicating whether or not the terminal was a test terminal.
	Test bool `json:"test"`

	// Timestamp date and time the terms and conditions acceptance occurred.
	Timestamp string `json:"timestamp"`

	// TransactionRef optional transaction ref if the terms and conditions was
	// associated with a transaction.
	TransactionRef string `json:"transactionRef"`

	// TransactionID optional transaction id if only log entries related to a
	// transaction should be returned.
	TransactionID string `json:"transactionId"`

	// Alias alias of the terms and conditions template used for this entry, if
	// any.
	Alias string `json:"alias"`

	// Name title of the document displayed on the terminal at the time of
	// capture.
	Name string `json:"name"`

	// Content full text of the document agreed to at the time of signature
	// capture.
	Content string `json:"content"`

	// ContentLeader first 32 characters of the full text. Used to support user
	// interfaces that show summaries.
	ContentLeader string `json:"contentLeader"`

	// HasSignature is a flag that indicates whether or not a signature has been
	// captured.
	HasSignature bool `json:"hasSignature"`

	// SigFormat specifies the image format to be used for returning signatures.
	SigFormat SignatureFormat `json:"sigFormat,omitempty"`

	// Signature is the base 64 encoded signature image if the format requested.
	Signature string `json:"signature"`
}

// SurveyQuestion models a survey question.
type SurveyQuestion struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// ID internal id for a survey question.
	ID string `json:"id"`

	// Ordinal ordinal number indicating the position of the survey question in
	// the post transaction sequence.
	Ordinal int `json:"ordinal"`

	// Enabled determines whether or not the question will be presented post
	// transaction.
	Enabled bool `json:"enabled"`

	// QuestionText is the full text of the transaction.
	QuestionText string `json:"questionText"`

	// QuestionType indicates the type of question. Valid values are 'yes_no' and
	// 'scaled'.
	QuestionType string `json:"questionType"`

	// TransactionCount is the total number of transactions processed during the
	// query period if results are requested.
	TransactionCount int `json:"transactionCount,omitempty"`

	// ResponseCount is the total number of responses during the query period if
	// results are requested.
	ResponseCount int `json:"responseCount,omitempty"`

	// ResponseRate is the response rate, expressed as a ratio, if results are
	// requested.
	ResponseRate float64 `json:"responseRate,omitempty"`

	// Responses is the set of response data points.
	Responses []SurveyDataPoint `json:"responses"`
}

// SurveyQuestionRequest models a request to retrieve or manipulate survey
// questions.
type SurveyQuestionRequest struct {
	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// QuestionID id of a single question.
	QuestionID string `json:"questionId"`

	// Timeout is an optional timeout override.
	Timeout int `json:"timeout"`
}

// SurveyQuestionResponse models a survey question response.
type SurveyQuestionResponse struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// Results is the full result set responsive to the original request.
	Results []SurveyQuestion `json:"results"`
}

// SurveyDataPoint models a request to retrieve or manipulate survey
// questions.
type SurveyDataPoint struct {
	// AnswerKey is a unique identifier for a specific answer type.
	AnswerKey string `json:"answerKey"`

	// AnswerDescription is a narrative description of the answer.
	AnswerDescription string `json:"answerDescription"`

	// ResponseCount is the number of responses.
	ResponseCount int `json:"responseCount"`

	// ResponsePercentage is response rate as a percentage of total transactions.
	ResponsePercentage float64 `json:"responsePercentage"`

	// AverageTransaction is the average transaction amount for a given answer.
	AverageTransaction float64 `json:"averageTransaction"`
}

// SurveyResultsRequest models a request to retrieve survey results.
type SurveyResultsRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// QuestionID id of a single question.
	QuestionID string `json:"questionId"`

	// StartDate is an optional start date for filtering response data.
	StartDate string `json:"startDate"`

	// EndDate is an optional end date for filtering response data.
	EndDate string `json:"endDate"`
}

// MediaMetadata models a request to retrieve survey results.
type MediaMetadata struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// ID id used to identify the media asset.
	ID string `json:"id"`

	// OriginalFile is the original filename assigned to the media asset.
	OriginalFile string `json:"originalFile"`

	// Name is the descriptive name of the media asset.
	Name string `json:"name"`

	// Description is a description of the media asset and its purpose.
	Description string `json:"description"`

	// Tags is an array of tags associated with a media asset.
	Tags []string `json:"tags"`

	// FileURL is the url for the full resolution versio of the media file.
	FileURL string `json:"fileUrl"`

	// ThumbnailURL is the url for to the thumbnail of an image.
	ThumbnailURL string `json:"thumbnailUrl"`

	// Video is an identifier used to flag video files.
	Video bool `json:"video"`
}

// UploadMetadata models information needed to process a file upload.
type UploadMetadata struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// UploadID optional id used to track status and progress of an upload while
	// in progress.
	UploadID string `json:"uploadId"`

	// FileSize is the size of the file to be uploaded in bytes.
	FileSize int64 `json:"fileSize"`

	// FileName is the name of file to be uploaded.
	FileName string `json:"fileName"`
}

// UploadStatus models the current status of a file upload.
type UploadStatus struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// ID id used to track status and progress of an upload while in progress.
	ID string `json:"id"`

	// MediaID is the media id assigned to the result.
	MediaID string `json:"mediaId"`

	// FileSize is the size of the file to be uploaded in bytes.
	FileSize int64 `json:"fileSize"`

	// UploadedAmount is the amount of the file already uploaded.
	UploadedAmount int64 `json:"uploadedAmount"`

	// Status is the current status of a file upload.
	Status string `json:"status"`

	// Complete indicates whether or not the upload and associated file
	// processing is complete.
	Complete bool `json:"complete"`

	// Processing indicates whether or not the file is processing. This normally
	// applied to video files undergoing format transcoding.
	Processing bool `json:"processing"`

	// Percentage current upload progress rounded to the nearest integer.
	Percentage int `json:"percentage"`

	// ThumbnailLocation is the url of a thumbnail for the file, if available.
	ThumbnailLocation string `json:"thumbnailLocation"`
}

// UploadStatusRequest is used to request the status of a file upload.
type UploadStatusRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// UploadID id used to track status and progress of an upload while in
	// progress.
	UploadID string `json:"uploadId"`
}

// MediaRequest models a request to retrieve or manipulate media assets.
type MediaRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// MediaID id used to track a media asset.
	MediaID string `json:"mediaId"`
}

// MediaLibraryResponse models a media library response.
type MediaLibraryResponse struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// MaxResults max to be returned in a single page. Defaults to the system max
	// of 250.
	MaxResults int `json:"maxResults"`

	// StartIndex starting index for paged results. Defaults to zero.
	StartIndex int `json:"startIndex"`

	// ResultCount total number of results accessible through paging.
	ResultCount int `json:"resultCount"`

	// Results enumerates all media assets available in the context.
	Results []MediaMetadata `json:"results"`
}

// Slide models a slide within a slide show.
type Slide struct {
	// MediaID is the id for the media asset to be used for this slide. Must be
	// an image.
	MediaID string `json:"mediaId"`

	// Ordinal position of the slide within the slide show.
	Ordinal int `json:"ordinal"`

	// ThumbnailURL is the fully qualified thumbnail url for the slide.
	ThumbnailURL string `json:"thumbnailUrl"`
}

// SlideShow models a media library response.
type SlideShow struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// ID is the primary id for the slide show.
	ID string `json:"id"`

	// Name is the name of the slide show.
	Name string `json:"name"`

	// Delay time between slides in seconds.
	Delay int `json:"delay"`

	// Slides enumerates all slides in the display sequence.
	Slides []*Slide `json:"slides"`
}

// SlideShowResponse models a slide show response.
type SlideShowResponse struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// MaxResults max to be returned in a single page. Defaults to the system max
	// of 250.
	MaxResults int `json:"maxResults"`

	// StartIndex starting index for paged results. Defaults to zero.
	StartIndex int `json:"startIndex"`

	// ResultCount total number of results accessible through paging.
	ResultCount int `json:"resultCount"`

	// Results enumerates all slide shows responsive to the original query.
	Results []SlideShow `json:"results"`
}

// SlideShowRequest models a request to retrieve or manipulate terminal slide
// shows.
type SlideShowRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// SlideShowID id used to track a slide show.
	SlideShowID string `json:"slideShowId"`
}

// BrandingAssetRequest models a request to retrieve or manipulate terminal
// slide shows.
type BrandingAssetRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// AssetID id used to track a branding asset.
	AssetID string `json:"assetId"`
}

// BrandingAsset models the priority and display settings for terminal media.
type BrandingAsset struct {
	// Timeout is the request timeout in seconds.
	Timeout int `json:"timeout"`

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool `json:"test"`

	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// ID id used to track a branding asset.
	ID string `json:"id"`

	// OwnerID is the id owner of the tenant who owns the branding asset.
	OwnerID string `json:"ownerId"`

	// TerminalID is the terminal id if this branding asset is specific to a
	// single terminal.
	TerminalID string `json:"terminalId"`

	// TerminalGroupID is the terminal group id if this branding asset is
	// specific to a terminal group.
	TerminalGroupID string `json:"terminalGroupId"`

	// MerchantID is the merchant id associated with this branding asset.
	MerchantID string `json:"merchantId"`

	// OrganizationID is the organization id associated with this branding asset.
	OrganizationID string `json:"organizationId"`

	// PartnerID is the partner id associated with this branding asset.
	PartnerID string `json:"partnerId"`

	// SlideShowID is the slide show associated with this branding asset, if any.
	// A branding asset can reference a slide show or media asset, but not both.
	SlideShowID string `json:"slideShowId"`

	// MediaID is the media id associated with this branding asset, if any. A
	// branding asset can reference a slide show or media asset, but not both.
	MediaID string `json:"mediaId"`

	// Padded applies standard margins to images displayed on terminals. Usually
	// the best option for logos.
	Padded bool `json:"padded"`

	// StartDate is the start date if this asset should be displayed based on a
	// schedule. Format: MM/DD/YYYY.
	StartDate string `json:"startDate"`

	// EndDate is the end date if this asset should be displayed based on a
	// schedule. Format: MM/DD/YYYY.
	EndDate string `json:"endDate"`

	// DaysOfWeek is an array of days of the week during which a branding asset
	// should be enabled. Days of the week are coded as integers starting with
	// Sunday (0) and ending with Saturday (6).
	DaysOfWeek []time.Weekday `json:"daysOfWeek"`

	// StartTime is the start date if this asset should be displayed based on a
	// schedule. Format: MM/DD/YYYY.
	StartTime string `json:"startTime"`

	// EndTime is the end date if this asset should be displayed based on a
	// schedule. Format: MM/DD/YYYY.
	EndTime string `json:"endTime"`

	// Ordinal is the ordinal number marking the position of this asset within
	// the branding stack.
	Ordinal int `json:"ordinal"`

	// Enabled enables the asset for display.
	Enabled bool `json:"enabled"`

	// Preview if true, the asset will be displayed in the merchant portal, but
	// not on merchant terminal hardware. Developers will usually want this to
	// always be false.
	Preview bool `json:"preview"`

	// UserID id of the user who created this branding asset, if applicable.
	UserID string `json:"userId"`

	// UserName name of the user who created this branding asset, if applicable.
	UserName string `json:"userName"`

	// Thumbnail the fully qualified URL of the thumbnail image for this branding
	// asset.
	Thumbnail string `json:"thumbnail"`

	// LastModified is the time and date this asset was last modified.
	LastModified string `json:"lastModified"`

	// Notes is a field for notes related to a branding asset.
	Notes string `json:"notes"`

	// Editable if true, the API credentials used to retrieve the branding asset
	// record can be used to update it.
	Editable bool `json:"editable"`

	// AssetType is the type of branding asset.
	AssetType string `json:"assetType"`

	// OwnerType is the type of user or tenant that owns this asset.
	OwnerType string `json:"ownerType"`

	// OwnerTypeCaption is a recommended caption for displaying the owner. Takes
	// into account multiple organization types.
	OwnerTypeCaption string `json:"ownerTypeCaption"`

	// OwnerName is the name of the tenant or entity that owns the branding
	// asset.
	OwnerName string `json:"ownerName"`

	// PreviewImage is the recommended image to be displayed when rendering a
	// preview of this branding asset.
	PreviewImage string `json:"previewImage"`

	// NarrativeEffectiveDates is a compact narrative string explaining the
	// effective date and time rules for a branding asset.
	NarrativeEffectiveDates string `json:"narrativeEffectiveDates"`

	// NarrativeDisplayPeriod is a compact narrative string explaining the
	// display period for a branding asset.
	NarrativeDisplayPeriod string `json:"narrativeDisplayPeriod"`
}

// BrandingAssetResponse models a branding asset response.
type BrandingAssetResponse struct {
	// Success indicates whether or not the request succeeded.
	Success bool `json:"success"`

	// Error is the error, if an error occurred.
	Error string `json:"error"`

	// ResponseDescription contains a narrative description of the transaction
	// result.
	ResponseDescription string `json:"responseDescription"`

	// OwnerID is the id owner of this branding stack.
	OwnerID string `json:"ownerId"`

	// OwnerType is the type of user or tenant that owns this branding stack.
	OwnerType string `json:"ownerType"`

	// OwnerName is the name of the entity or tenant that owns this branding
	// stack.
	OwnerName string `json:"ownerName"`

	// LevelName is the owner level currently being displayed.
	LevelName string `json:"levelName"`

	// NarrativeTime is a narrative description of the current simulate time.
	NarrativeTime string `json:"narrativeTime"`

	// ActiveAsset is the asset currently displayed on the terminal.
	ActiveAsset BrandingAsset `json:"activeAsset"`

	// Results enumerates all branding assets in a given credential scope.
	Results []BrandingAsset `json:"results"`
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

// TerminalLocateRequest contains information needed to retrieve location
// information for a terminal.
type TerminalLocateRequest struct {
	APICredentials
	Request LocateRequest `json:"request"`
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

// TerminalListQueuedTransactionsRequest returns a list of queued transactions
// on a terminal.
type TerminalListQueuedTransactionsRequest struct {
	APICredentials
	Request ListQueuedTransactionsRequest `json:"request"`
}

// TerminalDeleteQueuedTransactionRequest deletes one or all transactions from
// a terminal queue.
type TerminalDeleteQueuedTransactionRequest struct {
	APICredentials
	Request DeleteQueuedTransactionRequest `json:"request"`
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

	// ResetConnection forces the terminal cloud connection to be reset while a
	// transactions is in flight. This is a diagnostic settings that can be used
	// only for test transactions.
	ResetConnection bool
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

	// AuthResponseCode is the code returned by the terminal or the card issuer
	// to indicate the disposition of the message.
	AuthResponseCode string
}

// From creates an instance of ApprovalResponse with values
// from a generic type.
func (r ApprovalResponse) From(raw interface{}) (result ApprovalResponse, ok bool) {
	ok = copyTo(raw, &r)
	return r, ok
}

// TimeoutRequest models a low level request with a timeout and test flag.
type TimeoutRequest struct {
	// Timeout is the request timeout in seconds.
	Timeout int

	// Test specifies whether or not to route transaction to the test gateway.
	Test bool
}

// From creates an instance of TimeoutRequest with values
// from a generic type.
func (r TimeoutRequest) From(raw interface{}) (result TimeoutRequest, ok bool) {
	ok = copyTo(raw, &r)
	return r, ok
}

// CoreRequest contains core request fields for a transaction.
type CoreRequest struct {
	// TransactionRef contains a user-assigned reference that can be used to
	// recall or reverse transactions.
	TransactionRef string

	// Async defers the response to the transaction and returns immediately.
	// Callers should retrive the transaction result using the Transaction Status
	// API.
	Async bool

	// Queue adds the transaction to the queue and returns immediately. Callers
	// should retrive the transaction result using the Transaction Status API.
	Queue bool

	// WaitForRemovedCard specifies whether or not the request should block until
	// all cards have been removed from the card reader.
	WaitForRemovedCard bool

	// Force causes a transaction to override any in-progress transactions.
	Force bool

	// OrderRef is an identifier from an external point of sale system.
	OrderRef string

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string
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

	// ExpMonth is the card expiration month in MM format.
	ExpMonth string

	// ExpYear is the card expiration year in YY format.
	ExpYear string

	// AVSResponse contains address verification results if address information
	// was submitted.
	AVSResponse AVSResponse

	// ReceiptSuggestions contains suggested receipt fields.
	ReceiptSuggestions ReceiptSuggestions

	// Customer contains customer data, if any. Preserved for reverse
	// compatibility.
	Customer *Customer

	// Customers contains customer data, if any.
	Customers []Customer
}

// From creates an instance of PaymentMethodResponse with values
// from a generic type.
func (r PaymentMethodResponse) From(raw interface{}) (result PaymentMethodResponse, ok bool) {
	ok = copyTo(raw, &r)
	return r, ok
}

// CryptocurrencyResponse contains response details for a cryptocurrency
// transaction.
type CryptocurrencyResponse struct {
	// Confirmed indicates that the transaction has met the standard criteria for
	// confirmation on the network. (For example, 6 confirmations for level one
	// bitcoin.)
	Confirmed bool

	// CryptoAuthorizedAmount is the amount submitted to the blockchain.
	CryptoAuthorizedAmount string

	// CryptoNetworkFee is the network level fee assessed for the transaction
	// denominated in cryptocurrency. This fee goes to channel operators and
	// crypto miners, not BlockChyp.
	CryptoNetworkFee string

	// Cryptocurrency is the three letter cryptocurrency code used for the
	// transactions.
	Cryptocurrency string

	// CryptoNetwork indicates whether or not the transaction was processed on
	// the level one or level two network.
	CryptoNetwork string

	// CryptoReceiveAddress the address on the crypto network the transaction was
	// sent to.
	CryptoReceiveAddress string

	// CryptoBlock hash or other identifier that identifies the block on the
	// cryptocurrency network, if available or relevant.
	CryptoBlock string

	// CryptoTransactionID hash or other transaction identifier that identifies
	// the transaction on the cryptocurrency network, if available or relevant.
	CryptoTransactionID string

	// CryptoPaymentRequest is the payment request URI used for the transaction,
	// if available.
	CryptoPaymentRequest string

	// CryptoStatus is used for additional status information related to crypto
	// transactions.
	CryptoStatus string
}

// From creates an instance of CryptocurrencyResponse with values
// from a generic type.
func (r CryptocurrencyResponse) From(raw interface{}) (result CryptocurrencyResponse, ok bool) {
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

	// DestinationAccount is the settlement account for merchants with split
	// settlements.
	DestinationAccount string

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
