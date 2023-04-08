package middlewares

import (
	"net/http"

	"github.com/yoshihiro-shu/draft-backend/backend/internal/config"
)

func Cors(c config.Frontend) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", c.Domain)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Content-Type", "application/json;charset=utf-8")

			// CSRF Token
			w.Header().Set("Access-Control-Expose-Headers", headerName)
			w.Header().Set("Access-Control-Allow-Credentials", "true")

			// response for Preflight request
			if r.Method == http.MethodOptions {
				// CSRF Token
				w.Header().Set("Access-Control-Allow-Headers", headerName)

				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
