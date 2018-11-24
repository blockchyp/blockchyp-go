package blockchyp

/*
TerminalRequest adds API credentials to auth requests for use in
direct terminal transactions.
*/
type TerminalRequest struct {
  APICredentials
  Request AuthorizationRequest `json:"request"`
}
