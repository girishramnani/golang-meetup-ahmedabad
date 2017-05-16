package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

type Middleware func(http.Handler) http.Handler

func Logger(l *log.Logger) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			l.Printf(r.Method + " " + r.URL.Path)
			h.ServeHTTP(w, r)
		})
	}
}
func WithHeader(key, value string) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add(key, value)
			h.ServeHTTP(w, r)
		})
	}
}
func ComposeMiddleWare(h http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}

func main() {
	mux := http.NewServeMux()
	logger := log.New(os.Stdout, "From Server:", log.Ldate)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API server running"))
	})
	mux.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(time.Now().Format(time.RFC1123)))
	})
	root := ComposeMiddleWare(mux, Logger(logger), WithHeader("Powered-By", "GoLang"))
	http.ListenAndServe(":3000", root)
}
