package request

import (
	"fmt"
	"net/http"
)

func (c Context) TestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c.Logger.Println("url:", r.URL)
		fmt.Fprintln(w, "url:", r.URL)
		next.ServeHTTP(w, r)
	})
}
