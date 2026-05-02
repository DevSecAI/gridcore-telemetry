// GRID-SAST-006: SSRF — caller-supplied substation API URL fetched server-side.
package dr

import (
	"io"
	"net/http"
)

// DispatchCurtailment posts a load-curtailment instruction to a substation
// at `substationURL`, which is currently sourced from the request body.
func DispatchCurtailment(substationURL string, body io.Reader) (int, error) {
	req, err := http.NewRequest("POST", substationURL, body)
	if err != nil {
		return 0, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	return resp.StatusCode, nil
}
