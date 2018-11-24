package blockchyp

import (
	"testing"
  "fmt"
  "encoding/json"
)

func TestAuthEncoding(t *testing.T) {

  request := AuthorizationRequest{}
  request.CurrencyCode = "USD"
  request.Amount = "20.55"
  request.TerminalName = "Cashier #1"
  request.TransactionRef = "12345"
  request.PromptForTip = true
  request.Description = "Adventures Underground Richland"
  request.AltPrices = map[string]string{
    "BTC": "12345",
    "ETH": "234",
  }


  content, _ := json.Marshal(request)

  fmt.Println(string(content))

}
