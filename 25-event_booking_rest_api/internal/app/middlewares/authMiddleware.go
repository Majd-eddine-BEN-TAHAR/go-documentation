package middlewares

import (
	"context"
	"net/http"
	"strings"

	"event_booking_api/pkg/auth"
	"event_booking_api/pkg/errors"
)

// AuthMiddleware is a middleware that checks if the user is authenticated
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Step 1: Extract the token from the Authorization header
        tokenHeader := r.Header.Get("Authorization")

        // Step 2: Check if the Authorization header is empty
        if tokenHeader == "" {
            // If empty, respond with a missing token error (401 Unauthorized)
            errors.HandleError(w, errors.NewAPIError(nil, "Missing auth token", http.StatusUnauthorized))
            return
        }

        // Step 3: Extract the token from "Bearer " prefix
        token := strings.TrimPrefix(tokenHeader, "Bearer ")
        if token == tokenHeader {
            // Respond with an invalid token format error (401 Unauthorized)
            errors.HandleError(w, errors.NewAPIError(nil, "Invalid token format", http.StatusUnauthorized))
            return
        }

        // Step 4: Validate the token and get the user ID
        userID, err := auth.ValidateToken(token)
        if err != nil {
            // If validation fails (e.g., invalid or expired token), respond with an error (401 Unauthorized)
            errors.HandleError(w, errors.NewAPIError(err, err.Error(), http.StatusUnauthorized))
            return
        }

        // Step 5: Add the userID to the request context to use it within the request.
        ctx := context.WithValue(r.Context(), "userID", userID)
        r = r.WithContext(ctx)

        // Step 6: Token is valid, proceed with the request
        next.ServeHTTP(w, r)
    }
}