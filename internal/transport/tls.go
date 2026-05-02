// GRID-SAST-003: TLS config disables verification and forces TLS1.0.
package transport

import (
	"crypto/tls"
	"net/http"
)

func NewClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,        // GRID-SAST-003
			MinVersion:         tls.VersionTLS10,
		},
	}
	return &http.Client{Transport: tr}
}
