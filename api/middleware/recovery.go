package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (app *goPushServer) recovery(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			err := recover()
			if err != nil {
				jsonBody, _ := json.Marshal(map[string]string{
					"error": fmt.Sprint(err),
				})

				w.WriteHeader(http.StatusInternalServerError)
				w.Write(jsonBody)
			}

		}()

		next.ServeHTTP(w, r)

	})
}