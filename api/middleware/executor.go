package middleware

import "net/http"

func (app *goPushServer) executor() http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "app/json")
		route := app.endpoints[req.URL.Path]
		route.Executor(w, req)
	}
	return http.HandlerFunc(fn)
}