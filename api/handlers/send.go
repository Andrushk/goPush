package handlers

import "net/http"

func Send(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("not implemented"))
}