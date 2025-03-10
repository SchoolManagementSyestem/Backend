package middleware

import (
	"context"
	"net/http"
	"schoolManagementSystem/pkg/helpers"
)

type contextKey string

const userContextKey contextKey = "user"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader != "" {
			claims, err := helpers.VerifyTokenFromRequest(r)
			if err != nil {
				http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
				return
			}
			// Add user info to request context
			ctx := context.WithValue(r.Context(), userContextKey, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		next.ServeHTTP(w, r)
	})
}
