package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
)

//START OMIT

type SmartMiddleware struct {
	handler http.Handler
}

func (m *SmartMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rec := httptest.NewRecorder()
	m.handler.ServeHTTP(rec, r)
	for k, v := range rec.Header() {
		w.Header()[k] = v
	}
	w.Header().Set("X-We-Modified-This", "Awesomeness :)")
	w.WriteHeader(418)
	data := []byte("Smart Middleware says hello :)")
	clen, _ := strconv.Atoi(r.Header.Get("Content-Length"))
	clen += len(data)
	r.Header.Set("Content-Length", strconv.Itoa(clen))
	w.Write(data)
	w.Write(rec.Body.Bytes())
}

//END OMIT
