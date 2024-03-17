package main

import (
	"log"
	"net/http"
)

func runHTTPServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// log
		log.Printf("http: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)

		switch r.URL.Path {
		case "/200", "/ok":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("ok"))
		case "/500", "/internal-server-error":
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("internal server error"))
		case "/403", "/forbidden":
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("forbidden"))
		default:
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("not found"))
		}
	})

	log.Printf("http: starting http server at %s", ":8080")
	http.ListenAndServe(":8080", nil)
}
