package main

import (
	"fmt"
	"net/http"

	"github.com/jaehue/myweb"
)

func main() {
	r, _ := myweb.NewRouter()
	r.HandleFunc("/pattern[0-9]{3}", func(w http.ResponseWriter, r *http.Request, _ map[string]string) {
		fmt.Fprintln(w, "3 digit pattern")
	})
	r.HandleFunc("/pattern[0-9]{2}", func(w http.ResponseWriter, r *http.Request, _ map[string]string) {
		fmt.Fprintln(w, "2 digit pattern")
	})
	r.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request, _ map[string]string) {
		fmt.Fprintln(w, "this is about page.")
	})
	r.HandleFunc("/users/:id/images", func(w http.ResponseWriter, r *http.Request, params map[string]string) {
		fmt.Fprintln(w, "this is user("+params["id"]+")'s image page.")
	})
	r.HandleFunc("/users/:id/addresses", func(w http.ResponseWriter, r *http.Request, params map[string]string) {
		fmt.Fprintln(w, "this is user("+params["id"]+")'s address page.")
	})

	http.ListenAndServe(":8080", r)
}
