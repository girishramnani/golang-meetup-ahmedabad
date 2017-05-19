package main

import (
	"fmt"
	"net/http"
)

type Adapter func(http.Handler) http.Handler

func RequestMethod(method string) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == method {
				h.ServeHTTP(w, r)
			} else {
				w.WriteHeader(http.StatusMethodNotAllowed)
			}
		})
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hey there " + r.Method))
}

func main() {
	myHandler := http.HandlerFunc(IndexHandler)
	get := RequestMethod("POST")
	myHandlerGet := get(myHandler)
	http.Handle("/", myHandlerGet)
	fmt.Println("Server started listening on :8081")
	http.ListenAndServe(":8081", nil)
}
