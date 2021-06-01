package middleware

import (
	"net/http"
)

func setupCORS(w *http.ResponseWriter) {
	(*w).Header().Add("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func (app *goPushServer) cors(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		setupCORS(&w)
		if r.Method == http.MethodOptions {
			return
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}