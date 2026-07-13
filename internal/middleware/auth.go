package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/adullahkapadia/auth-service/internal/config"
	"github.com/adullahkapadia/auth-service/internal/utils"
)

type contextKey string

const UserIDKey contextKey = "userID"

func Auth(cfg *config.Config) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {

				http.Error(w, "Missing Authorization Header", http.StatusUnauthorized)

				return
			}

			if !strings.HasPrefix(authHeader, "Bearer ") {

				http.Error(w, "Invalid Token Format", http.StatusUnauthorized)

				return
			}

			token := strings.TrimPrefix(authHeader, "Bearer ")

			claims, err := utils.VerifyToken(
				token,
				cfg.JWTSecret,
			)

			if err != nil {

				http.Error(w, "Invalid Token", http.StatusUnauthorized)

				return
			}

			ctx := context.WithValue(
				r.Context(),
				UserIDKey,
				claims.UserID,
			)

			next.ServeHTTP(
				w,
				r.WithContext(ctx),
			)

		})

	}

}