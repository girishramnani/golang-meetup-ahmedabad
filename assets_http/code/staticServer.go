package main

import (
	"net/http"
)

func main() {
	// Simple static webserver:
	http.Handle("/Miniconda3/", http.StripPrefix("/Miniconda3/", http.FileServer(http.Dir("C:\\"))))
	http.ListenAndServe(":8089", nil)
}
