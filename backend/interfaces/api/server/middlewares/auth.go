package middlewares

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/yoshihiro-shu/draft-backend/backend/interfaces/api/server/auth"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get Token from Header
		token, err := auth.GetTokenFromHeader(r)
		if err != nil {
			fmt.Fprintf(w, "Something went wrong : %s\n", err.Error())
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		// verify expires ad
		err = claims.Valid()
		if err != nil {
			fmt.Fprintf(w, "Something went wrong : %s\n", err.Error())
			return
		}

		userId := claims["user_id"]

		// SET User Info to Context
		ctx := context.WithValue(r.Context(), auth.UserKey, userId)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
