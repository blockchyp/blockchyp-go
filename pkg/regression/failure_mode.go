package regression

import (
	"encoding/json"
	"io"
	"net"
	"os"
	"path/filepath"
	"time"

	"github.com/blockchyp/blockchyp-go"
)

var failureModeTests = testCases{
	{
		name:  "FailureMode/GatewayDown",
		group: testGroupInteractive,
		local: true,
		sim:   true,
		operations: []operation{
			{
				msg: `Stop the cloud stack or change the host in blockchyp.json and firmware.yml to an invalid value.

  When prompted, insert a valid test card.`,
				args: []string{
					"-type", "ping", "-terminal", terminalName, "-test",
				},
				expect: blockchyp.Acknowledgement{
					Success: true,
				},
			},
			{
				args: []string{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", amountRange(0, 100, 4000),
				},
				expect: blockchyp.AuthorizationResponse{
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
			{
				msg: "Restart the cloud stack.",
				validation: &validation{
					prompt: "Was the transaction sent to the gateway and approved (May take a few minutes)?",
					expect: true,
				},
			},
		},
	},
	{
		name:  "FailureMode/ExpiredCache",
		group: testGroupInteractive,
		local: true,
		sim:   true,
		operations: []operation{
			{
				args: []string{
					"-type", "ping", "-terminal", terminalName, "-test",
				},
				expect: blockchyp.Acknowledgement{
					Success: true,
				},
			},
			{
				args: []string{
					"-type", "cache-expire",
				},
			},
			{
				msg: `Stop the cloud stack or change the host in blockchyp.json and firmware.yml to an invalid value.

When prompted, insert a valid test card.`,
				args: []string{
					"-type", "charge", "-terminal", terminalName, "-test",
					"-amount", amountRange(0, 100, 4000),
				},
				expect: blockchyp.AuthorizationResponse{
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
			{
				msg: "Restart the cloud stack.",
				validation: &validation{
					prompt: "Was the transaction sent to the gateway and approved (May take a few minutes)?",
					expect: true,
				},
			},
		},
	},
	{
		name:  "FailureMode/IPChange",
		group: testGroupNonInteractive,
		sim:   true,
		operations: []operation{
			{
				args: []string{
					"-type", "ping", "-terminal", terminalName, "-test",
				},
				expect: blockchyp.Acknowledgement{
					Success: true,
				},
				fn: scrambleIPs,
			},
			{
				args: []string{
					"-type", "ping", "-terminal", terminalName, "-test",
				},
				timeout: 10 * time.Second,
				expect: blockchyp.Acknowledgement{
					Success: true,
				},
			},
		},
	},
}

func scrambleIPs() error {
	path := filepath.Join(os.TempDir(), ".blockchyp_routes")

	f, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	var cache blockchyp.RouteCache
	if err := json.NewDecoder(f).Decode(&cache); err != nil {
		return err
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
		return err
	}

	return nil
}
