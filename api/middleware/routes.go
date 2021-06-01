package middleware

import (
	"net/http"
)

type Route struct {
	HTTPMethod string
	Executor   http.HandlerFunc
}

type Routes map[string]*Route

func NewRoutes() Routes {
	return Routes{}
}

func (r Routes) Get(path string, executor http.HandlerFunc) {
	r[path] = &Route{HTTPMethod: http.MethodGet, Executor: executor}
}

func (r Routes) Post(path string, executor http.HandlerFunc) {
	r[path] = &Route{HTTPMethod: http.MethodPost, Executor: executor}
}
