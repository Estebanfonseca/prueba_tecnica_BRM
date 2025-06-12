package middleware

import (
	"context"
	"net/http"
	"strings"
	"api_users/api/models"
	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserIDKey contextKey = "userID"

func AuthMiddleware(secretKey string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Authorization header is required", http.StatusUnauthorized)
				return
			}

			tokenLogin := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenLogin == authHeader {
				http.Error(w, "Invalid token format", http.StatusUnauthorized)
				return
			}

			token, err := jwt.ParseWithClaims(tokenLogin, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(secretKey), nil
			})

			if err != nil || !token.Valid {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			claims, ok := token.Claims.(*models.Claims)
			if !ok {
				http.Error(w, "Invalid token claims", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}