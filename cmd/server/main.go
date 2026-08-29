package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /hello", hello)

	addr := ":8080"
	log.Printf("сервер слушает http://localhost%s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintln(w, "Привет! Это простой HTTP-сервер на Go.")
	fmt.Fprintln(w, "Попробуйте /hello?name=Иван")
}

func hello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "гость"
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "Привет, %s!\n", name)
}
