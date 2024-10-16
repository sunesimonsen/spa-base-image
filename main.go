package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

var fileHandler = http.FileServer(http.Dir("./public"))

type default404 struct {
	w http.ResponseWriter
}

func (d default404) WriteHeader(code int) {
	if code == http.StatusNotFound {
		indexRequest, err := http.NewRequest("GET", "", nil)
		if err != nil {
			log.Fatal(err)
		}
		d.Header().Set("Content-type", "text/html; charset=utf-8")
		fileHandler.ServeHTTP(d.w, indexRequest)
	} else {
		d.w.WriteHeader(code)
	}
}

func (d default404) Header() http.Header {
	return d.w.Header()
}

func (d default404) Write(data []byte) (int, error) {
	return d.w.Write(data)
}

// Default 404 to index.html
func withDefault404(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(default404{w}, r)
	})
}

func withLongTermCachhing(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "public, max-age=31556952")
		handler.ServeHTTP(w, r)
	})
}

func main() {
	http.HandleFunc("GET /up", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})

	// Serve all files from public/assets with long term caching
	http.Handle("GET /assets/", withLongTermCachhing(fileHandler))

	// Serve all files from public
	http.Handle("GET /", withDefault404(fileHandler))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
