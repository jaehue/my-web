package myweb

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"sync"
)

type router struct {
	mu       sync.RWMutex
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
	r.mu.Lock()
	r.handlers[path] = handler
	r.mu.Unlock()
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

func lookup(pattern, path string) (found bool, params map[string]string) {
	patterns := strings.Split(pattern, "/")
	paths := strings.Split(path, "/")

	if len(patterns) != len(paths) {
		return
	}

	params = make(map[string]string)

	for i := 0; i < len(patterns); i++ {
		if patterns[i] == paths[i] {
			found = true
			continue
		}
		if patterns[i][0] == ':' {
			params[patterns[i][1:]] = paths[i]
			found = true
			continue
		}
		found = false
		break
	}
	return
}
