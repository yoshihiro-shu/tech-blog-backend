package middlewares

import (
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/yoshihiro-shu/draft-backend/backend/internal/config"
	"github.com/yoshihiro-shu/draft-backend/backend/internal/pkg/logger"
	"go.uber.org/zap"
)

const (
	cookieName string = "csrf_token"
	// The default HTTP request header to inspect
	headerName = "X-CSRF-Token"
)

func SetterCsrfToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodOptions:
			w.Header().Set("Access-Control-Allow-Headers", headerName)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
		case http.MethodGet:
			// X-CSRF-Tokenをフロント側で受け取れるようにする
			w.Header().Set("Access-Control-Expose-Headers", headerName)
			// フロント側のブラウザにクッキーがセットされるようにする
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set(headerName, csrf.Token(r))
		case http.MethodPost:
			// X-CSRF-Tokenをフロント側で受け取れるようにする
			w.Header().Set("Access-Control-Expose-Headers", headerName)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		next.ServeHTTP(w, r)
	})
}

func CsrfProtecter(conf config.Configs, l logger.Logger) func(h http.Handler) http.Handler {
	return csrf.Protect(
		[]byte(conf.CsrfToken.Key),
		csrf.CookieName(cookieName),
		csrf.RequestHeader(headerName),
		csrf.SameSite(csrf.SameSiteNoneMode),
		csrf.MaxAge(3600*100),
		csrf.Path("/"),
		csrf.ErrorHandler(errHandler(l)),
	)
}

func errHandler(l logger.Logger) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l.Error("CSRF攻撃の疑いのあるリクエストが発行されました", zap.Error(nil))
		w.WriteHeader(http.StatusForbidden)
	})
}
