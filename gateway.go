package blockchyp

/*
TerminalRouteResponse models a terminal route response from the gateway.
*/
type TerminalRouteResponse struct {
  TerminalRoute
  Success bool `json:"success"`
}
