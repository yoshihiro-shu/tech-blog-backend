package middlewares

import (
	"net/http"

	"github.com/yoshihiro-shu/draft-backend/backend/interfaces/api/server/request"
	"github.com/yoshihiro-shu/draft-backend/backend/internal/config"
)

func Cors(c config.Frontend) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := request.NewHeader(w)

			header.ADD("Access-Control-Allow-Origin", c.Domain)
			header.ADD("Access-Control-Allow-Headers", "*")
			header.ADD("Access-Control-Allow-Methods", http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions)
			header.ADD("Content-Type", "application/json;charset=utf-8")
			// w.Header().Set("Access-Control-Allow-Credentials", "true")

			// response for Preflight request
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
