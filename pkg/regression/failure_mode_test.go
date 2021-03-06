// +build regression

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
	if acquirerMode {
		t.Skip("skipped for acquirer test run")
	}

	tests := map[string]struct {
		args       []interface{}
		assert     []interface{}
		validation validation

		// localOnly causes tests to be skipped when running in cloud relay
		// mode.
		localOnly bool
	}{
		"GatewayDown": {
			localOnly: true,
			args: []interface{}{
				[]string{
					"-type", "ping", "-terminal", terminalName, "-test",
				},
				`Stop the cloud stack or change the host in blockchyp.json and firmware.yml to an invalid value.

When prompted, insert a valid test card.`,
				[]string{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", amountRange(0, 100, 4000),
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
					RequestedAmount:  amount(0),
					AuthorizedAmount: amount(0),
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
			localOnly: true,
			args: []interface{}{
				[]string{
					"-type", "ping", "-terminal", terminalName, "-test",
				},
				[]string{
					"-type", "cache-expire",
				},
				`Stop the cloud stack or change the host in blockchyp.json and firmware.yml to an invalid value.

When prompted, insert a valid test card.`,
				[]string{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", amountRange(0, 100, 4000),
				},
				"Restart the cloud stack.",
			},
			assert: []interface{}{
				blockchyp.Acknowledgement{
					Success: true,
				},
				nil,
				nil,
				blockchyp.AuthorizationResponse{
					Success:          true,
					Approved:         true,
					Test:             true,
					TransactionType:  "charge",
					RequestedAmount:  amount(0),
					AuthorizedAmount: amount(0),
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
					"-type", "ping", "-terminal", terminalName, "-test",
				},
				scrambleIPs,
				10 * time.Second,
				[]string{
					"-type", "ping", "-terminal", terminalName, "-test",
				},
			},
			assert: []interface{}{
				blockchyp.Acknowledgement{
					Success: true,
				},
				nil,
				nil,
				blockchyp.Acknowledgement{
					Success: true,
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			cli := newCLI(t)
			if test.localOnly {
				cli.skipCloudRelay()
			}

			failures := make(chan string, 0)

			for i := range test.args {
				switch v := test.args[i].(type) {
				case string:
					setup(t, v, true)
				case func(*testing.T):
					v(t)
				case time.Duration:
					timer := time.NewTimer(v)
					go func() {
						<-timer.C
						failures <- "timed out while renegotiating route"
					}()
				case []string:
					cli.run(v, test.assert[i])
				}
			}

			select {
			case failure := <-failures:
				t.Error(failure)
			default:
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
