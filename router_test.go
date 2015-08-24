package myweb

import (
	"net/http"
	"testing"
)

type mockResponseWriter struct{}

func (m *mockResponseWriter) Header() (h http.Header)                 { return http.Header{} }
func (m *mockResponseWriter) Write(p []byte) (n int, err error)       { return len(p), nil }
func (m *mockResponseWriter) WriteString(s string) (n int, err error) { return len(s), nil }
func (m *mockResponseWriter) WriteHeader(int)                         {}

func TestRouter(t *testing.T) {
	r, _ := NewRouter()

	routed := false
	r.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		routed = true
	})

	w := new(mockResponseWriter)

	req, _ := http.NewRequest("GET", "/user", nil)
	r.ServeHTTP(w, req)

	if !routed {
		t.Fatal("routing failed")
	}
}

func TestRouterConcurrent(t *testing.T) {
	r, _ := NewRouter()

	routed1 := false
	routed2 := false

	done := make(chan struct{})
	go func() {
		r.HandleFunc("/user1", func(w http.ResponseWriter, r *http.Request) {
			routed1 = true
		})
		done <- struct{}{}
	}()

	r.HandleFunc("/user2", func(w http.ResponseWriter, r *http.Request) {
		routed2 = true
	})
	<-done

	w := new(mockResponseWriter)

	req, _ := http.NewRequest("GET", "/user1", nil)
	r.ServeHTTP(w, req)

	req, _ = http.NewRequest("GET", "/user2", nil)
	r.ServeHTTP(w, req)

	if !routed1 || !routed2 {
		t.Fatal("routing failed1")
	}
}
