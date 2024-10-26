package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		bytes, _ := io.ReadAll(r.Body)
		log.Print(r.URL.RawQuery)
		log.Print(string(bytes))
	})

	http.ListenAndServe(":80", nil)
}
