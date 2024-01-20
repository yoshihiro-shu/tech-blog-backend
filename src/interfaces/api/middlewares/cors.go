package middlewares

import (
	"net/http"

	"github.com/yoshihiro-shu/tech-blog-backend/src/internal/config"
)

func Cors(c config.Frontend) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Max-Age", "3600")
			// w.Header().Set("Access-Control-Allow-Origin", c.Domain)
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Content-Type", "application/json;charset=utf-8")

			// response for Preflight request
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
