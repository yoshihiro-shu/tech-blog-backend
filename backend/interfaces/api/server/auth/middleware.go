package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get Token from Header
		token, err := getTokenFromHeader(r)
		if err != nil {
			fmt.Fprintf(w, "Something went wrong : %s\n", err.Error())
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userId := claims["user_id"]

		// SET User Info to Context
		ctx := context.WithValue(r.Context(), UserKey, userId)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
