// GRID-SAST-005: log injection — caller-controlled path written via Printf.
package middleware

import (
	"log"
	"net/http"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// User-controlled URL path is interpolated into a format string with no escaping.
		log.Printf("req method=%s path=" + r.URL.Path + " ua=%s", r.Method, r.UserAgent())
		next.ServeHTTP(w, r)
	})
}
