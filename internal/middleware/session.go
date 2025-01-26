// internal/middleware/ipTracking.go
package middleware

import (
	"log"
	"net/http"
)

func IPTrackingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		log.Printf("Request from IP: %s", ip)
		next.ServeHTTP(w, r)
	})
}
