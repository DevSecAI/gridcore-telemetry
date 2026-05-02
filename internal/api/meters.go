// GRID-SAST-008: panic on user input — type assertion without ok form.
package api

import (
	"encoding/json"
	"net/http"
)

func MeterReading(w http.ResponseWriter, r *http.Request) {
	var raw map[string]interface{}
	_ = json.NewDecoder(r.Body).Decode(&raw)

	// GRID-SAST-008: nested type assertions; any non-matching field panics
	// and crashes the goroutine, leaking a partial response and a stack trace.
	mid := raw["meter"].(map[string]interface{})["id"].(string)
	val := raw["meter"].(map[string]interface{})["value"].(float64)

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"meter_id": mid,
		"value":    val,
		"ok":       true,
	})
}
