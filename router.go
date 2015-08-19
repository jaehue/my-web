package myweb

import (
	"fmt"
	"net/http"
	"regexp"
)

type router struct {
	handlers map[string]http.HandlerFunc
}

func NewRouter() (*router, error) {
	r := &router{
		handlers: make(map[string]http.HandlerFunc),
	}
	return r, nil
}

func (r *router) HandleFunc(path string, handler http.HandlerFunc) {
	if len(path) == 0 {
		panic("path must not be nil")
	}
	if path[0] != '/' {
		panic(fmt.Sprintf("path must begin with '/' in path '%s'", path))
	}
	r.handlers[path] = handler
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path

	for pattern, handler := range r.handlers {
		if m, _ := regexp.MatchString(pattern, path); m {
			handler(w, req)
			return
		}
	}
	http.Error(w, fmt.Sprintf("'%s' is invalid path.", path), http.StatusNotFound)
}
