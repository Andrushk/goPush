package middleware

import (
	"net/http"
	"github.com/Andrushk/goPush/config"
)

func (app *goPushServer) authorization(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("key")
		if apiKey != config.GetString("apikey") {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			next.ServeHTTP(w, r)
		}
	}

	return http.HandlerFunc(fn)
}