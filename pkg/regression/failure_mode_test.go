package regression

import (
	"encoding/json"
	"io"
	"net"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/blockchyp/blockchyp-go"
)

func TestFailureModes(t *testing.T) {
	tests := map[string]struct {
		args       []interface{}
		assert     []interface{}
		validation validation
	}{
		"GatewayDown": {
			args: []interface{}{
				[]string{
					"-type", "ping", "-terminal", "Test Terminal", "-test",
				},
				"Stop the cloud stack or change the host in blockchyp.json and firmware.yml to an invalid value.",
				[]string{
					"-type", "charge", "-terminal", "Test Terminal", "-test",
					"-amount", "41.00",
				},
				"Restart the cloud stack.",
			},
			assert: []interface{}{
				blockchyp.Acknowledgement{
					Success: true,
				},
				nil,
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "charge",
					RequestedAmount:  "41.00",
					AuthorizedAmount: "41.00",
					PaymentType:      notEmpty,
					MaskedPAN:        notEmpty,
					StoreAndForward:  true,
				},
			},
			validation: validation{
				prompt: "Was the transaction sent to the gateway and approved (May take a few minutes)?",
				expect: true,
			},
		},
		"ExpiredCache": {
			args: []interface{}{
				[]string{
					"-type", "ping", "-terminal", "Test Terminal", "-test",
				},
				[]string{
					"-type", "cache-expire",
				},
				"Stop the cloud stack or change the host in blockchyp.json and firmware.yml to an invalid value.",
				[]string{
					"-type", "charge", "-terminal", "Test Terminal", "-test",
					"-amount", "41.00",
				},
				"Restart the cloud stack.",
			},
			assert: []interface{}{
				blockchyp.Acknowledgement{
					Success: true,
				},
				nil,
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "charge",
					RequestedAmount:  "42.00",
					AuthorizedAmount: "42.00",
					PaymentType:      notEmpty,
					MaskedPAN:        notEmpty,
					StoreAndForward:  true,
				},
			},
			validation: validation{
				prompt: "Was the transaction sent to the gateway and approved (May take a few minutes)?",
				expect: true,
			},
		},
		"IPChange": {
			args: []interface{}{
				[]string{
					"-type", "ping", "-terminal", "Test Terminal", "-test",
				},
				scrambleIPs,
				[]string{
					"-type", "charge", "-terminal", "Test Terminal", "-test",
					"-amount", "42.00",
				},
			},
			assert: []interface{}{
				blockchyp.Acknowledgement{
					Success: true,
				},
				nil,
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					RequestedAmount:  "42.00",
					AuthorizedAmount: "42.00",
				},
			},
		},
	}

	cli := newCLI(t)

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {

			for i := range test.args {
				switch v := test.args[i].(type) {
				case string:
					setup(t, v, true)
				case func(*testing.T):
					v(t)
					time.Sleep(10 * time.Second)
				case []string:
					cli.run(v, test.assert[i])
				}
			}

			validate(t, test.validation)
		})
	}
}

func scrambleIPs(t *testing.T) {
	path := filepath.Join(os.TempDir(), ".blockchyp_routes")

	f, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	var cache blockchyp.RouteCache
	if err := json.NewDecoder(f).Decode(&cache); err != nil {
		t.Fatal(err)
	}

	for k, v := range cache.Routes {
		ip := net.ParseIP(v.Route.IPAddress)
		ip[len(ip)-1] = 0x0
		v.Route.IPAddress = ip.String()
		cache.Routes[k] = v
	}

	f.Truncate(0)
	f.Seek(0, io.SeekStart)
	if err := json.NewEncoder(f).Encode(cache); err != nil {
		t.Fatal(err)
	}
}
