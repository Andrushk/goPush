package middleware

import "net/http"

func (app *goPushServer) executor() http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		route := app.endpoints[req.URL.Path]
		route.Executor(w, req)
		w.Header().Set("Content-Type", "app/json")
	}
	return http.HandlerFunc(fn)
}