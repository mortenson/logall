package main

import (
	"io"
	"log"
	"net/http"
	"strings"
)

type LogHandler struct {
}

func (h *LogHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !strings.Contains(r.URL.Path, "logall") {
		http.NotFound(w, r)
		return
	}
	bytes, _ := io.ReadAll(r.Body)
	log.Printf("path: %s | query: %s | body: %s", r.URL.Path, r.URL.RawQuery, string(bytes))
}

func main() {
	handler := &LogHandler{}
	http.ListenAndServe(":80", handler)
}
