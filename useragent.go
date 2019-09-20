package blockchyp

import (
	"fmt"
	"net/http"
)

// BuildUserAgent assembles a user agent header for use with requests to the
// gateway and terminals.
func BuildUserAgent() string {
	version := Version
	if version == "" {
		version = "Unknown"
	}

	return fmt.Sprintf("BlockChyp-Go/%s", version)
}

// AddUserAgent adds a user agent header to an http.RoundTripper.
func AddUserAgent(transport http.RoundTripper, userAgent string) http.RoundTripper {
	return &addUserAgent{
		Transport: transport,
		UserAgent: userAgent,
	}
}

type addUserAgent struct {
	Transport http.RoundTripper
	UserAgent string
}

func (t *addUserAgent) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Set("User-Agent", t.UserAgent)

	return t.Transport.RoundTrip(r)
}
