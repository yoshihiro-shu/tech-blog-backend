package middlewares

import (
	"net/http"

	"github.com/yoshihiro-shu/tech-blog-backend/backend/internal/logger"
	"go.uber.org/zap"
)

func Logger(logger logger.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Info("request",
				zap.String("method", r.Method),
				zap.String("url", r.URL.Path),
				zap.String("RemoteAddr", r.RemoteAddr),
				zap.String("UserAgent", r.UserAgent()),
				zap.String("Referer", r.Referer()),
				zap.String("Host", r.Host),
				zap.String("RequestURI", r.RequestURI),
			)
			next.ServeHTTP(w, r)
		})
	}
}
