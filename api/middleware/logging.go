package middleware

import (
	"log"
	"net/http"
	"time"
)

type StatusRecorder struct {
    http.ResponseWriter
    Status int
}

func (r *StatusRecorder) WriteHeader(status int) {
    r.Status = status
    r.ResponseWriter.WriteHeader(status)
}

func (app *goPushServer) logging(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		recorder := &StatusRecorder{
            ResponseWriter: w,
            Status:         200,
        }
		next.ServeHTTP(recorder, r)
		t2 := time.Now()
		log.Printf("[%v] %q, status: %s, %v\n", r.Method, r.URL, http.StatusText(recorder.Status), t2.Sub(t1))
	}

	return http.HandlerFunc(fn)
}